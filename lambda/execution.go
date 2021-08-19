package lambda

import (
	"sync"
	"time"

	"github.com/MauricioAntonioMartinez/aws/helpers"
	"golang.org/x/crypto/bcrypt"
)




func NewHash(val string) (string,error) {
	hash,err := bcrypt.GenerateFromPassword([]byte(val),10)

	if err != nil {
		return "",err
	}

	return string(hash),err
}

type LambdaExecutionManager struct {
	CapPerInvoque int
}

type LambdaExecution struct {
	FuncArn string
	ValHash string
	Events chan string
	Hash string
	LastReceive time.Time
	CurrentExecutions int
	mt sync.RWMutex
}




var (
	running map[string][]*LambdaExecution = make(map[string][]*LambdaExecution)
)

func (l *LambdaExecutionManager) getExection(hash string) *LambdaExecution {
	for _,exe := range running {
		for _,lx := range exe {
			lx.mt.RLock()
			if lx.Hash == hash  && bcrypt.CompareHashAndPassword([]byte(hash),[]byte(lx.ValHash)) == nil{
				lx.mt.RUnlock()
				return lx
			}
			lx.mt.RUnlock()
		}
	}
	return nil
}

func (l *LambdaExecutionManager) newLambdaExecution(arn string) error {
	valHash := helpers.RandString(20)
	hs,err := NewHash(valHash)

	if err !=nil {
		return err
	}

	 running[arn] = append(running[arn], &LambdaExecution{
		FuncArn: arn,
		Events: make(chan string),
		Hash: hs,
		ValHash: valHash,
		mt: sync.RWMutex{},
		CurrentExecutions: 1,
		LastReceive: time.Now(),
	 })

	 return nil
}

func (l *LambdaExecutionManager) getAvaibleExecution(funcArn string) *LambdaExecution {

	lambdaExecutors,ok := running[funcArn]

	if !ok { return nil }

	var lambdaExecutor *LambdaExecution

	for _,lx := range lambdaExecutors {
		lx.mt.Lock()
		if lx.CurrentExecutions < l.CapPerInvoque {
			lx.CurrentExecutions++
			lambdaExecutor = lx
			break
		}
		lx.mt.Unlock()
	}

	return lambdaExecutor
}

