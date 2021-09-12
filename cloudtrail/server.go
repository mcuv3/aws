package cloudtrail

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

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
	"gorm.io/gorm"
)

const (
	ServerPrefix = "/cloudtrail.CloudTrailService/"
)

var (
	PublicMethods = []string{}
)

type CloudTrailService struct {
	logger  zerolog.Logger
	grpc    *grpc.Server
	webgrpc http.Server
}

func Run(cmd cli.CloudTrailCmd, logger zerolog.Logger) error {

	db, err := database.New(cmd.DbUrl)
	if err != nil {
		logger.Fatal().Err(err).Msg("Unable to connect to the database.")
	}
	runMigrations(db)

	service := newCloudTrailService(cmd, db)
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

func (s *CloudTrailService) Error(err error, code codes.Code, msg string) error {
	s.logger.Err(err)
	return status.Errorf(code, msg)
}

func runMigrations(db *gorm.DB) {
	db.AutoMigrate(model.RootUser{}, model.Group{}, model.Role{}, model.Policy{}, model.User{}, model.AccessKey{})
}

func newCloudTrailService(cmd cli.CloudTrailCmd, db *gorm.DB) CloudTrailService {
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

	return CloudTrailService{logger: logger, grpc: s, webgrpc: *webgrpc}
}

func (s *CloudTrailService) RegisterGRPC() {
	aws.RegisterCloudTrailServiceServer(s.grpc, s)
	reflection.Register(s.grpc)
}

func (s *CloudTrailService) ServeGrpc(port string, serviceName string) error {
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

func (s *CloudTrailService) ServeWeb(port string, serviceName string) error {
	s.logger.Info().Str("web", serviceName).Msgf("Staring web server on port :%s", port)

	if err := s.webgrpc.ListenAndServe(); err != nil {
		fmt.Println(err)
		return nil
	}

	return nil
}
