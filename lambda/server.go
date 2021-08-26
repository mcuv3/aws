package lambda

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
	"github.com/MauricioAntonioMartinez/aws/docker"
	"github.com/MauricioAntonioMartinez/aws/helpers"
	"github.com/MauricioAntonioMartinez/aws/model"
	aws "github.com/MauricioAntonioMartinez/aws/proto"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
)

type LambdaService struct {
	auth          *auth.AuthInterceptor
	logger        zerolog.Logger
	docker        *docker.ContainerDispatcher
	db            *gorm.DB
	region        string
	CapPerInvoque int
	grpc          *grpc.Server
	grpcweb       http.Server
	LambdaExecutionManager
}

func Run(cmd cli.LambdaCmd, l zerolog.Logger) error {

	db, err := database.New(cmd.DbUrl)
	if err != nil {
		l.Fatal().Err(err).Msg("Filed to connect the database")
		return err
	}
	runMigrations(db)

	l.Info().Str("name", db.Dialector.Name()).Msgf("databse %s", db.Debug().Name())

	service := newLambdaService(cmd, db)
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
	db.AutoMigrate(&model.Runtime{})
	db.AutoMigrate(&model.Function{})
}

func newLambdaService(cmd cli.LambdaCmd, db *gorm.DB) LambdaService {
	l := helpers.NewLogger()

	service := LambdaService{
		auth: &auth.AuthInterceptor{Issuer: cmd.Name(), Logger: l,
			ServerPrefix:  "/lambda.LambdaService/",
			PublicMethods: []string{"ReceiveEvents"},
			Mannager:      &auth.JWTMannger{SecretKey: cmd.Secret, Duration: time.Hour}},
		logger: l,

		db:     db,
		docker: docker.NewContainerDispatcher(cmd.Workers, &docker.DockerRuntime{}),
		region: cmd.Region,
		LambdaExecutionManager: LambdaExecutionManager{
			CapPerInvoque: 10,
		},
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(service.auth.Unary()),
		grpc.StreamInterceptor(service.auth.Stream()))
	service.grpc = s
	grpcweb := helpers.NewGrpcWeb(s, cmd.PortWeb, cmd.Origin)
	service.grpcweb = *grpcweb

	service.setRuntimes()
	service.docker.Start()

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
