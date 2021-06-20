package iam

import (
	"context"
	"net"

	aws "github.com/MauricioAntonioMartinez/aws/proto"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)



type IAMService struct {
	aws.UnimplementedIAMServiceServer
}


func (*IAMService) CreateUser(ctx context.Context,req *aws.CreateUserRequest)(*aws.CreateUserResponse,error) {
	return nil,nil
}


func Run(l zerolog.Logger) error { 
	
	lis,err := net.Listen("tcp",":50051")

	if err !=nil { 
		return err
	}

	s := grpc.NewServer()

	aws.RegisterIAMServiceServer(s,&IAMService{})
	reflection.Register(s)

	l.Info().Str("server","iam").Msg("Staring server on port :50051")

	if err:= s.Serve(lis); err !=nil { 
		return nil
	}

	return nil
}