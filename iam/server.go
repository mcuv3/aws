package iam

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

type IAMService struct {
	aws.UnimplementedIAMServiceServer
	storage *IamRepository
	auth    *auth.AuthInterceptor
	logger  zerolog.Logger
	grpc    *grpc.Server
	webgrpc http.Server
}

func Run(cmd cli.IamCmd, logger zerolog.Logger) error {

	db, err := database.New(cmd.DbUrl)
	if err != nil {
		logger.Fatal().Err(err).Msg("Unable to connect to the database.")
	}
	runMigrations(db)

	service := newIamService(cmd, db)
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

func (s *IAMService) Error(err error, code codes.Code, msg string) error {
	s.logger.Err(err)
	return status.Errorf(code, msg)
}

func runMigrations(db *gorm.DB) {
	db.AutoMigrate(model.RootUser{}, model.Group{}, model.Role{}, model.Policy{}, model.User{}, model.AccessKey{})
}

func newIamService(cmd cli.IamCmd, db *gorm.DB) IAMService {
	logger := helpers.NewLogger()

	authInt := auth.AuthInterceptor{Issuer: cmd.Name(), Logger: logger,
		ServerPrefix:  "/iam.IAMService/",
		PublicMethods: []string{"RootUserLogin", "SignUp", "UserLogin"},
		Mannager:      &auth.JWTMannger{SecretKey: cmd.Secret, Duration: time.Hour}}

	s := grpc.NewServer(grpc.ChainUnaryInterceptor(authInt.Unary()))
	webgrpc := helpers.NewGrpcWeb(s, cmd.PortWeb, cmd.Origin)

	return IAMService{storage: &IamRepository{db: db}, logger: logger, auth: &authInt, grpc: s, webgrpc: *webgrpc}
}

func (s *IAMService) RegisterGRPC() {
	aws.RegisterIAMServiceServer(s.grpc, s)
	reflection.Register(s.grpc)
}

func (s *IAMService) ServeGrpc(port string, serviceName string) error {
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

func (s *IAMService) ServeWeb(port string, serviceName string) error {
	s.logger.Info().Str("web", serviceName).Msgf("Staring web server on port :%s", port)

	if err := s.webgrpc.ListenAndServe(); err != nil {
		fmt.Println(err)
		return nil
	}

	return nil
}
