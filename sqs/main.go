package sqs

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/MauricioAntonioMartinez/aws/auth"
	database "github.com/MauricioAntonioMartinez/aws/db"
	"github.com/MauricioAntonioMartinez/aws/model"
	aws "github.com/MauricioAntonioMartinez/aws/proto"
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
	db *gorm.DB
)

var ctx = context.Background()




type SQSServer struct {
	aws.UnimplementedSQSServiceServer
	auth *auth.AuthInterceptor 
 }


func (s *SQSServer) CreateQueue(ctx context.Context,req * aws.CreateQueueRequest) (*aws.CreateQueueResponse,error){ 
	queueName := req.GetName();

	fmt.Println(queueName)
	us ,err := s.auth.GetUserMetadata(ctx)

	if err !=nil { 
		return nil,grpc.Errorf(codes.Unauthenticated,"Unable to authenticate")
	}

	fmt.Println(us)


	// if res := db.First(&user, "account_id = ?", accountId); res.Error != nil {
	// 	return c.Status(400).JSON(model.ErrorResponse{Status: 400, Message: "Cannot the user."})
	// }

	// fmt.Println(body.Configuration)

	// if body.Configuration.DeliveryDelayTime == "" || body.Configuration.MessageRetentionTime == "" || body.Configuration.VisibilityTime == "" {
	// 	return c.Status(400).JSON(model.ErrorResponse{Status: 400, Message: "Invalid configuration please check."})
	// }

	// name := strings.ReplaceAll(body.Name, " ", "")

	// queue := model.Queue{
	// 	Url:           model.BaseUrl + "/" + user.AccountId + "/" + name,
	// 	Configuration: body.Configuration,
	// 	AccountId:     user.AccountId,
	// 	Name:          name,
	// }

	// res := db.Create(&queue)

	// if res.Error != nil {
	// 	return c.Status(400).JSON(model.ErrorResponse{Status: 400, Message: "Couldn't create queue."})
	// }

	// return &aws.CreateQueueResponse{Queue: queue},nil
	return &aws.CreateQueueResponse{Queue: &aws.Queue{Id: "1",Name: queueName}},nil
}

func (*SQSServer) SendMessage(ctx context.Context,req * aws.SendMessageRequest)(*aws.SendMessageResponse,error){ 

	return nil,nil
}

func (*SQSServer) ReceiveMessage(req *aws.ReceiveMessageRequest,stream aws.SQSService_ReceiveMessageServer)error{ 
	fmt.Println("This should hot reload")
	return nil
}




func authConfig() *auth.AuthInterceptor {
	 m := auth.NewJWTMannager("mysecret",time.Hour)
	 return auth.NewAuthInterceptor(m,"sqs")
}

// TODO: migrate from rest api to gRPC
func Run(l zerolog.Logger) error {


	db,err :=database.New()

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
	})
	
	reflection.Register(s)
	

	l.Info().Str("server","sqs").Msg("Staring server on port :50051")

	if err := s.Serve(lis); err !=nil { 
		l.Fatal().Msg(fmt.Sprint("Failed to serve %w",err))
	}

	return nil
}





