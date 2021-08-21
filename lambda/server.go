package lambda

import (
	"fmt"
	"net"
	"time"

	"github.com/MauricioAntonioMartinez/aws/auth"
	database "github.com/MauricioAntonioMartinez/aws/db"
	"github.com/MauricioAntonioMartinez/aws/docker"
	"github.com/MauricioAntonioMartinez/aws/model"
	aws "github.com/MauricioAntonioMartinez/aws/proto"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
)
 
type LambdaServer struct {
	auth *auth.AuthInterceptor
	logger zerolog.Logger
	docker *docker.ContainerDispatcher
	db *gorm.DB
	region string
	CapPerInvoque int
	LambdaExecutionManager
}

var region string



func Run(l zerolog.Logger) error {

	db,err := database.New()

	if err !=nil {
		err = fmt.Errorf("Failed to connect database: %w",err)
		l.Fatal().Err(err).Msg("Filed to connect the database")
		return err
	}

	db.AutoMigrate(&model.Runtime{})
	db.AutoMigrate(&model.Function{})

	l.Info().Str("name",db.Dialector.Name()).Str("databse",db.Debug().Name())

	
	lis,err := net.Listen("tcp",":50051")

	if err !=nil {
		l.Err(err).Str("listener",lis.Addr().String()).Msg("Failed to listen")
		return err
	} 

	authInterceptor := auth.AuthInterceptor{Issuer: "lambda",Logger: l,
	ServerPrefix: "/lambda.LambdaService/",
	PublicMethods: []string{"ReceiveEvents"},
	Mannager: &auth.JWTMannger{SecretKey: "supersecret", Duration: time.Hour}}

	s := grpc.NewServer(grpc.UnaryInterceptor(authInterceptor.Unary()),grpc.StreamInterceptor(authInterceptor.Stream()))


	d := docker.NewContainerDispatcher(3,&docker.DockerRuntime{})
	d.Start()
	
	region = "us-east-1"
	setRegion(region)

	aws.RegisterLambdaServiceServer(s,&LambdaServer{
		auth: &authInterceptor,
		logger: l,
		db: db,
		docker: d,
		region: "us-east-1", 
		LambdaExecutionManager: LambdaExecutionManager{
			CapPerInvoque:10,
		},
	})

	reflection.Register(s)


	l.Info().Str("server","lambda").Msg("Staring lambda server on port :50051")

	if err := s.Serve(lis); err !=nil {
		l.Fatal().Msg(fmt.Sprint("Failed to serve %w",err))
		return err
	}

	return nil
}


