package sqs

import (
	"context"
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/MauricioAntonioMartinez/aws/auth"
	database "github.com/MauricioAntonioMartinez/aws/db"
	"github.com/MauricioAntonioMartinez/aws/model"
	aws "github.com/MauricioAntonioMartinez/aws/proto"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

var (
	Secret string
)

var ctx = context.Background()

func (s *SQSServer) Error(err error, code codes.Code, msg string) error {
	s.logger.Err(err)
	return grpc.Errorf(code, msg)
}


type SQSServer struct {
	aws.UnimplementedSQSServiceServer
	auth *auth.AuthInterceptor 
	logger zerolog.Logger
	db *gorm.DB
 }


func (s *SQSServer) CreateQueue(ctx context.Context,req * aws.CreateQueueRequest) (*aws.CreateQueueResponse,error){ 
	queueName := req.GetName();

	us ,err := s.auth.GetUserMetadata(ctx)

	if err !=nil { 
		return nil,s.Error(err,codes.Unauthenticated,"Unable to authenticate")
	}

 
	if res := s.db.Where("account_id = ? AND name =  ?", us.AccountId,queueName).First(&model.Queue{}); res.Error == nil  {
		return nil,grpc.Errorf(codes.NotFound,"Something went wrong %s",queueName)
	}

	conf := req.GetConf()

	
	if string(conf.DeliveryDelayTime) == "" || string(conf.MessageRetentionTime) == "" || string(conf.VisibilityTime) == "" {
		return nil,s.Error(errors.New("Invalid configuration"),codes.InvalidArgument,"Invalid configuration please check.")
	}


	arn, err := auth.NewArn(auth.SQS,auth.US_EAST_1,us.AccountId,fmt.Sprintf("/queue/%s",queueName))

	if err != nil {
		return nil,s.Error(err,codes.InvalidArgument,"Invalid ARN")
	}


	queue := model.Queue{
		Name:          queueName, 
		AccountId:     us.AccountId,
		Arn: arn.String(),
		Configuration: model.ConfigurationQueue{
			MessageRetentionTime: conf.MessageRetentionTime.String(),
			VisibilityTimeout: int(conf.VisibilityTimeout),
			VisibilityTime: conf.VisibilityTime.String(),
			MessageRetention: int(conf.MessageRetention),
			DeliveryDelay: int(conf.DeliveryDelayTime),
			DeliveryDelayTime: conf.DeliveryDelayTime.String(),
		},
	}


	res := s.db.Create(&queue)

	if res.Error != nil {
		return nil,s.Error(res.Error,codes.Aborted,"Couldn't create queue.")
	}

	return &aws.CreateQueueResponse{Queue: &aws.Queue{
		Name: queueName,
		Arn: arn.String(),
		}},nil
}

func (s *SQSServer) SendMessage(ctx context.Context,req * aws.SendMessageRequest)(*aws.SendMessageResponse,error){ 

	us ,err := s.auth.GetUserMetadata(ctx)

	if err != nil {
		return nil,s.Error(err,codes.Unauthenticated,"Unable to authenticate")
	}

	queue:= model.Queue{}

	res := s.db.Where("account_id = ? AND name = ?",us.AccountId,req.GetQueueName()).First(&queue)

	if res.Error !=nil {
		return nil,s.Error(res.Error,codes.NotFound,"Queue not found")
	}


	msg := model.Message{
		Message: req.GetMessage(),
		QueueID: queue.ID,
		UserMessageId: model.UserMessageId(uuid.NewString()),
	}

	
	res = s.db.Create(&msg) 
 
	if res.Error != nil {	
		return nil,s.Error(res.Error,codes.InvalidArgument,"Unable to send the message.")
	}

	//	err = redis.Set(ctx, msg.Key(), m, 0).Err()

		// if res.Error != nil || err != nil {
		// 	return nil,s.Error() : "Could not save the message", Status: 400})
		// }
	return &aws.SendMessageResponse{
		Message: &aws.Message{
			Id: string(msg.UserMessageId),
			Message: msg.Message,
			QueueName: queue.Name,
		},
	},nil
}

func (s *SQSServer) ReceiveMessage(req *aws.ReceiveMessageRequest,stream aws.SQSService_ReceiveMessageServer)error{ 
	
	us ,err := s.auth.GetUserMetadata(ctx)
	
	if err != nil {	
		s.logger.Err(err).Str("error:",err.Error())
		return s.Error(err,codes.Unauthenticated,"Unable to authenticate")
	}
	
	queue := model.Queue{}

	res := s.db.Where("account_id = ? AND name = ?",us.AccountId,req.GetQueueName()).First(&queue)

	if res.Error !=nil {
		return s.Error(res.Error,codes.NotFound,"Queue not found")
	}



	// msgs := make(chan []model.Message)
	errChan := make(chan error)

		go func() {

			messages := make([]model.Message, 0)
			var n int

			for {
				m := model.Message{}
 
				// keys, err := redis.Keys(ctx, queue.Pattern()).Result()

				 res := s.db.Where("queue_id = ?",queue.ID).Find(&m) 

				 if res.Error !=nil {
					  errChan <- s.Error(res.Error,codes.NotFound,"Couldn't find the message")
				}

				stream.SendMsg( &aws.ReceiveMessageResponse{Message: 
					&aws.Message{Id:  string(m.UserMessageId), Message: m.Message, QueueName: queue.Name}})
				// if err == nil {
				// 	for _, k := range keys {
				// 		val, _ := redis.Get(ctx, k).Result()
				// 		redis.Del(ctx, k)
				// 		json.Unmarshal([]byte(val), &m)

				// 		messages = append(messages, m)
				// 	}
				// }

				time.Sleep(time.Second)
				n++

				if int(req.GetBatchLimit()) == len(messages) || int(req.GetWaitLimit()) <= n {
					break
				}

			}

			errChan <- nil
		}()
		
	return <- errChan
}




func authConfig() *auth.AuthInterceptor {
	 m := auth.NewJWTMannager("mysecret",time.Hour)
	 return auth.NewAuthInterceptor(m,"sqs")
}

// TODO: migrate from rest api to gRPC
func Run(l zerolog.Logger) error {


	db,err := database.New()

	if err !=nil {
		err = fmt.Errorf("Failed to connect database: %w",err)
		l.Fatal().Err(err).Msg("Filed to connect the database")
		return err
	}

	l.Info().Str("name",db.Dialector.Name()).Str("database",db.Debug().Name()).Msg("Succeeded to connect to the database")
	
	 database.NewRedis() 

	db.AutoMigrate(&model.Queue{},
		&model.User{},&model.Message{})
	// db.AutoMigrate(db)


	lis,err := net.Listen("tcp",":50051")

	if err !=nil { 
		l.Err(err).Str("listener",lis.Addr().String()).Msg("Unable to listen")
		return err
	}
 
	authManager := auth.NewAuthInterceptor(
		auth.NewJWTMannager("supersecret",time.Hour),
		"iam",
	)
 
	s := grpc.NewServer(grpc.ChainUnaryInterceptor(authManager.Unary()))

	aws.RegisterSQSServiceServer(s,&SQSServer{
		auth: authManager,
		db: db,
		logger: l,
	})
	
	reflection.Register(s)
	

	l.Info().Str("server","sqs").Msg("Staring server on port :50051")

	if err := s.Serve(lis); err !=nil { 
		l.Fatal().Msg(fmt.Sprint("Failed to serve %w",err))
	}

	return nil
}





