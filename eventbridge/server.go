package eventbridge

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/MauricioAntonioMartinez/aws/auth"
	"github.com/MauricioAntonioMartinez/aws/cli"
	database "github.com/MauricioAntonioMartinez/aws/db"
	"github.com/MauricioAntonioMartinez/aws/helpers"
	"github.com/MauricioAntonioMartinez/aws/interceptors"
	"github.com/MauricioAntonioMartinez/aws/model"
	aws "github.com/MauricioAntonioMartinez/aws/proto"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
)

type EventBridgeService struct {
	auth    *auth.AuthInterceptor
	Name    auth.Service
	logger  zerolog.Logger
	db      *gorm.DB
	region  string
	grpc    *grpc.Server
	grpcweb http.Server
	repo    EventBridgeRepo
}

func Run(cmd cli.EventBridgeCmd, l zerolog.Logger) error {

	db, err := database.New(cmd.DbUrl)

	if err != nil {
		l.Fatal().Err(err).Msgf("Error connecting the database on servie %s", cmd.Name())
		return err
	}
	runMigrations(db)

	l.Info().Str("name", db.Dialector.Name()).Msgf("database %s", db.Debug().Name())

	service := newEventBridgeService(cmd, db)
	service.RegisterGRPC()

	// if err := service.Seed(); err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }

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

func runMigrations(db *gorm.DB) {
	db.AutoMigrate(model.ServiceEvent{}, model.Rule{}, model.Target{})
}

func newEventBridgeService(cmd cli.EventBridgeCmd, db *gorm.DB) EventBridgeService {
	l := helpers.NewLogger()
	auditor := interceptors.NewAuditInterceptor("kafka:9092", "audit", 1)

	service := EventBridgeService{
		Name:   auth.EventBridge,
		logger: l,
		repo:   EventBridgeRepo{db},
		auth: &auth.AuthInterceptor{
			Issuer: cmd.Name(),
			Mannager: &auth.JWTMannger{
				SecretKey: cmd.Secret,
				Duration:  time.Hour,
			},
			ServerPrefix:  "/eventbridge.EventBridgeService/",
			PublicMethods: []string{},
			Logger:        l,
		},
		db:     db,
		region: cmd.Region,
	}
	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(service.auth.Unary(), auditor.Unary()),
		grpc.StreamInterceptor(auditor.Stream()),
	)
	service.grpc = s
	grpcweb := helpers.NewGrpcWeb(s, cmd.PortWeb, cmd.Origin)
	service.grpcweb = *grpcweb

	return service
}

func (s *EventBridgeService) RegisterGRPC() {
	aws.RegisterEventBridgeServiceServer(s.grpc, s)
	reflection.Register(s.grpc)
}

func (s *EventBridgeService) ServeGrpc(port, serviceName string) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))

	if err != nil {
		return err
	}

	s.logger.Info().Str("grpc", serviceName).Msgf("Starting %s gRPC server on port %s ", serviceName, port)

	if err := s.grpc.Serve(lis); err != nil {
		return err
	}
	return nil
}

func (s *EventBridgeService) ServeWeb(port, serviceName string) error {
	s.logger.Info().Str("web", serviceName).Msgf("Staring web server proxy for gRPC on port %s", port)

	if err := s.grpcweb.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
