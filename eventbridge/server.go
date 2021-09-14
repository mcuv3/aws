package eventbridge

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
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
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
)

var (
	repo          EventBridgeRepo
	PublicMethods = []string{}
)

const (
	ServerPrefix = "/eventbridge.EventBridgeService/"
)

type EventBridgeService struct {
	Name       string
	logger     zerolog.Logger
	region     string
	grpc       *grpc.Server
	grpcweb    http.Server
	dispatcher *EventBridgeDispatcher
}

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
	go listenEvents(cmd)
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
	brokers := strings.Split(cmd.Brokers, ",")
	auditInterceptor := interceptors.NewAuditInterceptor(interceptors.AuditInterceptorConfig{
		Brokers: brokers,
		Topic:   eventbus.Audit,
		Verbose: true,
		Sid:     cmd.Name(),
		Region:  cmd.Region,
	})
	authInterceptor := interceptors.NewAuthInterceptor(interceptors.AuthInterceptorConfig{
		Issuer:        cmd.Name(),
		ServerPrefix:  ServerPrefix,
		PublicMethods: PublicMethods,
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

func listenEvents(cmd cli.EventBridgeCmd) {
	brokers := strings.Split(cmd.Brokers, ",")
	dispatcher := newEventBridgeDispatcher(EventBridgeDispatcherConfig{
		Identifier: cmd.Name(),
		Verbose:    false,
		Topic:      eventbus.Audit,
		Brokers:    brokers,
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
