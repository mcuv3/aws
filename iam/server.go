package iam

import (
	"net"

	database "github.com/MauricioAntonioMartinez/aws/db"
	aws "github.com/MauricioAntonioMartinez/aws/proto"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)



type IAMService struct {
	aws.UnimplementedIAMServiceServer
	storage *IamRepository
}



func Run(l zerolog.Logger) error { 


	lis,err := net.Listen("tcp",":50051")

	if err !=nil { 
		return err
	}

	db,err := database.New()

	if err !=nil { 
		l.Fatal().Err(err).Msg("Unable to connect to the database.")
	}


	s := grpc.NewServer()
	aws.RegisterIAMServiceServer(s,&IAMService{storage: &IamRepository{db: db}})
	
	reflection.Register(s)

	l.Info().Str("server","iam").Msg("Staring server on port :50051")

		if err:= s.Serve(lis); err !=nil { 
			return nil
		}
		return err

}