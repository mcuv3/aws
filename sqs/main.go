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
)

var ctx = context.Background()

func (s *SQSServer) Error(err error, code codes.Code, msg string) error {
	s.logger.Err(err)
	return grpc.Errorf(code, msg)
}


type SQSServer struct {
	// aws.UnimplementedSQSServiceServer
	auth *auth.AuthInterceptor 
	logger zerolog.Logger
	db *gorm.DB
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
 
 
	authInterceptor := auth.AuthInterceptor{Issuer: "iam", Logger: l,
	ServerPrefix: "/iam.SQSService/",
	PublicMethods: []string{},
		Mannager: &auth.JWTMannger{SecretKey: "supersecret", Duration: time.Hour}}
 
	s := grpc.NewServer(grpc.UnaryInterceptor(authInterceptor.Unary()),grpc.StreamInterceptor(authInterceptor.Stream()))

	aws.RegisterSQSServiceServer(s,&SQSServer{
		auth: &authInterceptor,
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





