package eventbridge

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"

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
	Name       string
	logger     zerolog.Logger
	region     string
	grpc       *grpc.Server
	grpcweb    http.Server
	dispatcher *EventBridgeDispatcher
}

var (
	repo EventBridgeRepo
)

func Run(cmd cli.EventBridgeCmd, l zerolog.Logger) error {

	db, err := database.New(cmd.DbUrl)

	if err != nil {
		l.Fatal().Err(err).Msgf("Error connecting the database on servie %s", cmd.Name())
		return err
	}
	runMigrations(db)
	repo = EventBridgeRepo{db}

	l.Info().Str("name", db.Dialector.Name()).Msgf("database %s", db.Debug().Name())

	service := newEventBridgeService(cmd, db)
	service.registerGRPC()

	// if err := service.Seed(); err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }
	go listenEvents()
	go service.closeListener()

	if cmd.EnableGrpc && cmd.EnableWeb {
		go service.serveGrpc(cmd.PortGrpc, cmd.Name())
		return service.serveWeb(cmd.PortWeb, cmd.Name())
	}

	if cmd.EnableGrpc {
		return service.serveGrpc(cmd.PortGrpc, cmd.Name())
	}

	if cmd.EnableWeb {
		return service.serveWeb(cmd.PortWeb, cmd.Name())
	}

	return errors.New("enable either web or grpc or both")
}

func runMigrations(db *gorm.DB) {
	db.AutoMigrate(model.ServiceEvent{}, model.Rule{}, model.Target{})
}

func newEventBridgeService(cmd cli.EventBridgeCmd, db *gorm.DB) EventBridgeService {
	unary, stream := getInterceptors(cmd)
	service := EventBridgeService{
		Name:   cmd.Name(),
		logger: helpers.NewLogger(),
		region: cmd.Region,
	}
	service.grpc = grpc.NewServer(
		grpc.ChainUnaryInterceptor(unary...),
		grpc.ChainStreamInterceptor(stream...),
	)
	service.grpcweb = *helpers.NewGrpcWeb(service.grpc, cmd.PortWeb, cmd.Origin)
	return service
}

func (s *EventBridgeService) registerGRPC() {
	aws.RegisterEventBridgeServiceServer(s.grpc, s)
	reflection.Register(s.grpc)
}

func getInterceptors(cmd cli.EventBridgeCmd) ([]grpc.UnaryServerInterceptor, []grpc.StreamServerInterceptor) {
	auditInterceptor := interceptors.NewAuditInterceptor(interceptors.AuditInterceptorConfig{
		Brokers: []string{"broker:29092"},
		Topic:   "audit",
		Verbose: true,
	})
	authInterceptor := interceptors.NewAuthInterceptor(interceptors.AuthInterceptorConfig{
		Issuer:        cmd.Name(),
		ServerPrefix:  "/eventbridge.EventBridgeService/",
		PublicMethods: []string{},
		Logger:        helpers.NewLogger(),
		SecretKey:     cmd.Secret,
	})

	return []grpc.UnaryServerInterceptor{authInterceptor.Unary(), auditInterceptor.Unary()},
		[]grpc.StreamServerInterceptor{auditInterceptor.Stream(), auditInterceptor.Stream()}
}

func (s *EventBridgeService) serveGrpc(port, serviceName string) error {
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

func (s *EventBridgeService) serveWeb(port, serviceName string) error {
	s.logger.Info().Str("web", serviceName).Msgf("Staring web server proxy for gRPC on port %s", port)

	if err := s.grpcweb.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

func listenEvents() {
	dispatcher := newEventBridgeDispatcher(EventBridgeDispatcherConfig{
		Identifier: "eventBridge",
		Verbose:    false,
		Topic:      "audit",
		Brokers:    []string{"broker:29092"},
	})
	dispatcher.Start()
}

func (s *EventBridgeService) closeListener() {
	signal := make(chan os.Signal, 1)

	sig := <-signal
	fmt.Println(sig)

	s.logger.Info().Msg("Received SIGINT, shutting down")
	s.dispatcher.Stop()
	s.grpc.Stop()
	s.grpcweb.Close()
	os.Exit(0)

}
