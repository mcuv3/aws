package iam

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/MauricioAntonioMartinez/aws/cli"
	database "github.com/MauricioAntonioMartinez/aws/db"
	"github.com/MauricioAntonioMartinez/aws/eventbus"
	"github.com/MauricioAntonioMartinez/aws/helpers"
	"github.com/MauricioAntonioMartinez/aws/interceptors"
	"github.com/MauricioAntonioMartinez/aws/model"
	aws "github.com/MauricioAntonioMartinez/aws/proto"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

const (
	ServerPrefix = "/iam.IAMService/"
)

var (
	PublicMethods = []string{"RootUserLogin", "SignUp", "UserLogin"}
)

type IAMService struct {
	aws.UnimplementedIAMServiceServer
	storage *IamRepository
	logger  zerolog.Logger
	jwt     JWTMannger
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

	return errors.New("enable either web or grpc or both")
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
	brokers := strings.Split(cmd.Brokers, ",")
	auditInterceptor := interceptors.NewAuditInterceptor(interceptors.AuditInterceptorConfig{
		Brokers: brokers,
		Topic:   eventbus.Audit,
		Verbose: true,
	})

	authInterceptor := interceptors.NewAuthInterceptor(interceptors.AuthInterceptorConfig{Issuer: cmd.Name(), Logger: logger,
		ServerPrefix:  ServerPrefix,
		PublicMethods: PublicMethods,
		SecretKey:     cmd.Secret,
	})
	s := grpc.NewServer(grpc.ChainUnaryInterceptor(authInterceptor.Unary(), auditInterceptor.Unary()))
	webgrpc := helpers.NewGrpcWeb(s, cmd.PortWeb, cmd.Origin)

	return IAMService{storage: &IamRepository{db: db}, jwt: JWTMannger{
		SecretKey: cmd.Secret,
		Duration:  time.Hour,
	}, logger: logger, grpc: s, webgrpc: *webgrpc}
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

func (s *IAMService) Dev(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
