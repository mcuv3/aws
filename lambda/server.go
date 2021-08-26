package lambda

import (
	"fmt"
	"net"
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
	LambdaExecutionManager
}

func Run(cmd cli.LambdaCmd, l zerolog.Logger) error {

	db, err := database.New(cmd.DbUrl)
	if err != nil {
		l.Fatal().Err(err).Msg("Filed to connect the database")
		return err
	}
	db.AutoMigrate(&model.Runtime{})
	db.AutoMigrate(&model.Function{})

	l.Info().Str("name", db.Dialector.Name()).Msgf("databse %s", db.Debug().Name())

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cmd.PortGrpc))
	if err != nil {
		l.Err(err).Str("listener", lis.Addr().String()).Msg("Failed to listen")
		return err
	}

	service, s := newLambdaService(cmd, db)
	aws.RegisterLambdaServiceServer(s, &service)
	reflection.Register(s)

	l.Info().Str("service", "lambda").Msgf("Staring lambda service on port :%s", cmd.PortGrpc)

	if err := s.Serve(lis); err != nil {
		l.Fatal().Msg(fmt.Sprint("Failed to serve %w", err))
		return err
	}

	return nil
}

func newLambdaService(cmd cli.LambdaCmd, db *gorm.DB) (LambdaService, *grpc.Server) {
	l := helpers.NewLogger()

	server := LambdaService{
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
		grpc.UnaryInterceptor(server.auth.Unary()),
		grpc.StreamInterceptor(server.auth.Stream()))

	server.setRuntimes()
	server.docker.Start()

	return server, s
}
