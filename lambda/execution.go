package lambda

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/MauricioAntonioMartinez/aws/container"
	"github.com/MauricioAntonioMartinez/aws/helpers"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewHash(val string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(val), 10)
	if err != nil {
		return "", err
	}
	return string(hash), err
}

type LambdaExecutionManager struct {
	CapPerInvoke int
}

type LambdaExecution struct {
	ID                string
	ValHash           string
	Events            chan string
	Hash              string
	LastReceive       time.Time
	CurrentExecutions int
	mt                sync.RWMutex
}

var (
	running map[string][]*LambdaExecution = make(map[string][]*LambdaExecution)
)

func (l *LambdaExecutionManager) getExecutionByHash(hash string) *LambdaExecution {
	for _, exe := range running {
		for _, lx := range exe {
			lx.mt.RLock()
			if lx.Hash == hash && bcrypt.CompareHashAndPassword([]byte(hash), []byte(lx.ValHash)) == nil {
				lx.mt.RUnlock()
				return lx
			}
			lx.mt.RUnlock()
		}
	}
	return nil
}

func (l *LambdaExecutionManager) newLambdaExecution(ID string) (*LambdaExecution, error) {
	valHash := helpers.RandString(20)
	hs, err := NewHash(valHash)

	if err != nil {
		return nil, err
	}

	lx := LambdaExecution{
		ID:                ID,
		Events:            make(chan string),
		Hash:              hs,
		ValHash:           valHash,
		mt:                sync.RWMutex{},
		CurrentExecutions: 1,
		LastReceive:       time.Now(),
	}

	running[ID] = append(running[ID], &lx)

	return &lx, nil
}

func (l *LambdaExecutionManager) getExecutionId(id string) *LambdaExecution {

	lambdaExecutors, ok := running[id]

	if !ok {
		return nil
	}

	var lambdaExecutor *LambdaExecution

	for _, lx := range lambdaExecutors {
		lx.mt.Lock()
		if lx.CurrentExecutions < l.CapPerInvoke {
			lx.CurrentExecutions++
			lambdaExecutor = lx
			lx.mt.Unlock()
			break
		}
		lx.mt.Unlock()
	}

	if lambdaExecutor != nil {
		fmt.Println("Executor available hash", lambdaExecutor.Hash)
	}

	return lambdaExecutor
}

type LambdaExecutionConfig struct {
	ID      string
	image   string
	ram     int64
	data    string
	handler string
}

func (exe *LambdaExecutionManager) Run(config LambdaExecutionConfig) error {

	freeExecution := exe.getExecutionId(config.ID)

	if freeExecution == nil {
		lx, err := exe.newLambdaExecution(config.ID)

		if err != nil {
			return status.Errorf(codes.Internal, "Something went wrong running the lambda")
		}
		env := map[string]interface{}{
			"EVENT_DATA": config.data,
			"HANDLER":    config.handler,
			"HOST":       "localhost:6003",
			"HASH":       lx.Hash,
			"LISTEN":     true,
		}

		dispatcher.RunContainer(container.RunContainerOptions{
			Image:       config.image,
			Ram:         config.ram,
			Environment: env,
			Name:        strconv.Itoa(int(time.Now().Unix())),
		})
	} else {
		go func() {
			freeExecution.Events <- config.data
		}()
	}

	return nil

}
