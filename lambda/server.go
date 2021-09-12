package lambda

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/MauricioAntonioMartinez/aws/cli"
	"github.com/MauricioAntonioMartinez/aws/container"
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

const (
	ServerPrefix = "/lambda.LambdaService/"
)

type LambdaService struct {
	logger  zerolog.Logger
	db      *gorm.DB
	region  string
	grpc    *grpc.Server
	grpcweb http.Server
}

var (
	repo          LambdaRepo
	execution     LambdaExecutionManager
	dispatcher    *container.ContainerDispatcher
	PublicMethods = []string{"ReceiveEvents"}
)

func Run(cmd cli.LambdaCmd, l zerolog.Logger) error {

	db, err := database.New(cmd.DbUrl)
	if err != nil {
		l.Fatal().Err(err).Msg("Filed to connect the database")
		return err
	}
	runMigrations(db)
	repo = LambdaRepo{db}

	l.Info().Str("name", db.Dialector.Name()).Msgf("database %s", db.Debug().Name())

	service := newLambdaService(cmd, db)
	service.RegisterGRPC()
	go listenEvents()
	startLambdaExecutions(cmd)

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

	return errors.New("enable either web or grpc or both ")
}

func runMigrations(db *gorm.DB) {
	db.AutoMigrate(&model.Runtime{})
	db.AutoMigrate(&model.Function{})
}

func newLambdaService(cmd cli.LambdaCmd, db *gorm.DB) LambdaService {
	l := helpers.NewLogger()
	brokers := strings.Split(cmd.Brokers, ",")
	auditInterceptor := interceptors.NewAuditInterceptor(interceptors.AuditInterceptorConfig{
		Brokers: brokers,
		Topic:   eventbus.Audit,
		Verbose: true,
	})
	auth := interceptors.NewAuthInterceptor(interceptors.AuthInterceptorConfig{Issuer: cmd.Name(), Logger: l,
		ServerPrefix:  ServerPrefix,
		PublicMethods: PublicMethods,
		SecretKey:     cmd.Secret,
	})

	service := LambdaService{
		logger: l,
		db:     db,
		region: cmd.Region,
	}

	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(auth.Unary(), auditInterceptor.Unary()),
		grpc.StreamInterceptor(auth.Stream()))
	service.grpc = s
	grpcweb := helpers.NewGrpcWeb(s, cmd.PortWeb, cmd.Origin)
	service.grpcweb = *grpcweb

	service.setRuntimes()

	return service
}

func (s *LambdaService) RegisterGRPC() {
	aws.RegisterLambdaServiceServer(s.grpc, s)
	reflection.Register(s.grpc)
}

func (s *LambdaService) ServeGrpc(port string, serviceName string) error {
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

func (s *LambdaService) ServeWeb(port string, serviceName string) error {
	s.logger.Info().Str("web", serviceName).Msgf("Staring web server on port :%s", port)

	if err := s.grpcweb.ListenAndServe(); err != nil {
		fmt.Println(err)
		return nil
	}

	return nil
}

func listenEvents() {
	lis := newLambdaListener()
	lis.Start()
}

func startLambdaExecutions(cmd cli.LambdaCmd) {
	dispatcher = container.NewContainerDispatcher(cmd.Workers, &container.DockerRuntime{})
	execution = LambdaExecutionManager{
		CapPerInvoke: cmd.EventsPerInvoke,
	}
	dispatcher.Start()
}
