package sqs

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/MauricioAntonioMartinez/aws/auth"
	"github.com/MauricioAntonioMartinez/aws/cli"
	database "github.com/MauricioAntonioMartinez/aws/db"
	"github.com/MauricioAntonioMartinez/aws/helpers"
	"github.com/MauricioAntonioMartinez/aws/model"
	aws "github.com/MauricioAntonioMartinez/aws/proto"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func (s *SQSService) Error(err error, code codes.Code, msg string) error {
	s.logger.Err(err)
	return status.Errorf(code, msg)
}

type SQSService struct {
	auth    *auth.AuthInterceptor
	logger  zerolog.Logger
	db      *gorm.DB
	grpc    *grpc.Server
	grpcweb http.Server
}

func Run(cmd cli.SqsCmd, l zerolog.Logger) error {
	fmt.Println(cmd.EnableWeb)
	db, err := database.New(cmd.DbUrl)
	if err != nil {
		err = fmt.Errorf("Failed to connect database: %w", err)
		l.Fatal().Err(err).Msg("Filed to connect the database")
		return err
	}
	runMigrations(db)

	l.Info().Str("name", db.Dialector.Name()).Str("database", db.Debug().Name()).Msg("Succeeded to connect to the database")
	database.NewRedis()

	service := newSqsService(cmd, db)
	service.RegisterGRPC()

	if cmd.EnableGrpc && cmd.EnableWeb {
		go service.ServeGrpc(cmd.PortGrpc, cmd.Name())
		return service.ServeWeb(cmd.PortWeb, cmd.Name())
	}

	if cmd.EnableGrpc {
		return service.ServeGrpc(cmd.PortGrpc, cmd.Name())
	}

	if cmd.EnableWeb {
		return service.ServeWeb(cmd.PortWeb, cmd.Name())
	}

	return errors.New("Enable either web or grpc or both.")
}

func runMigrations(db *gorm.DB) {
	db.AutoMigrate(&model.Queue{},
		&model.User{}, &model.Message{})
}

func newSqsService(cmd cli.SqsCmd, db *gorm.DB) SQSService {
	l := helpers.NewLogger()
	authInterceptor := auth.AuthInterceptor{Issuer: cmd.Name(), Logger: l,
		ServerPrefix:  "/iam.SQSService/",
		PublicMethods: []string{},
		Mannager:      &auth.JWTMannger{SecretKey: cmd.Secret, Duration: time.Hour}}

	s := grpc.NewServer(grpc.UnaryInterceptor(authInterceptor.Unary()), grpc.StreamInterceptor(authInterceptor.Stream()))

	grpcweb := helpers.NewGrpcWeb(s, cmd.PortWeb, cmd.Origin)

	return SQSService{
		auth:    &authInterceptor,
		logger:  l,
		grpc:    s,
		grpcweb: *grpcweb,
		db:      db,
	}

}

func (s *SQSService) RegisterGRPC() {
	aws.RegisterSQSServiceServer(s.grpc, s)
	reflection.Register(s.grpc)
}

func (s *SQSService) ServeGrpc(port string, serviceName string) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))

	if err != nil {
		return err
	}

	s.logger.Info().Str("grpc", serviceName).Msgf("Staring gRPC server on port :%s", port)

	if err := s.grpc.Serve(lis); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (s *SQSService) ServeWeb(port string, serviceName string) error {
	s.logger.Info().Str("web", serviceName).Msgf("Staring web server on port :%s", port)

	if err := s.grpcweb.ListenAndServe(); err != nil {
		fmt.Println(err)
		return nil
	}

	return nil
}
