package iam

import (
	"net"
	"time"

	"github.com/MauricioAntonioMartinez/aws/auth"
	database "github.com/MauricioAntonioMartinez/aws/db"
	"github.com/MauricioAntonioMartinez/aws/model"
	aws "github.com/MauricioAntonioMartinez/aws/proto"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
)

type IAMService struct {
	aws.UnimplementedIAMServiceServer
	storage *IamRepository
	auth    *auth.AuthInterceptor
	logger  zerolog.Logger
}

func Run(logger zerolog.Logger) error {

	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		return err
	}

	db, err := database.New()

	if err != nil {
		logger.Fatal().Err(err).Msg("Unable to connect to the database.")
	}

	runMigrations(db)

	authInt := auth.AuthInterceptor{Issuer: "iam", Logger: logger,
		Mannager: &auth.JWTMannger{SecretKey: "supersecret", Duration: time.Hour}}

	s := grpc.NewServer(grpc.ChainUnaryInterceptor(authInt.Unary()))

	aws.RegisterIAMServiceServer(s, &IAMService{storage: &IamRepository{db: db}, logger: logger, auth: &authInt})

	reflection.Register(s)

	logger.Info().Str("server", "iam").Msg("Staring server on port :50051")

	if err := s.Serve(lis); err != nil {
		return nil
	}
	return err
}

func (s *IAMService) Error(err error, code codes.Code, msg string) error {
	s.logger.Err(err)
	return grpc.Errorf(code, msg)
}

func runMigrations(db *gorm.DB) {
	db.AutoMigrate(model.AwsUser{}, model.AccessKey{}, model.Group{}, model.UserIam{}, model.Role{}, model.Policy{})
}
