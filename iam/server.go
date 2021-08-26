package iam

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/MauricioAntonioMartinez/aws/auth"
	"github.com/MauricioAntonioMartinez/aws/cli"
	database "github.com/MauricioAntonioMartinez/aws/db"
	"github.com/MauricioAntonioMartinez/aws/model"
	aws "github.com/MauricioAntonioMartinez/aws/proto"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
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

func Run(cmd cli.IamCmd, logger zerolog.Logger) error {

	lis, err := net.Listen("tcp", ":6000")

	if err != nil {
		return err
	}

	db, err := database.New(cmd.DbUrl)

	if err != nil {
		logger.Fatal().Err(err).Msg("Unable to connect to the database.")
	}

	runMigrations(db)

	authInt := auth.AuthInterceptor{Issuer: "iam", Logger: logger,
		ServerPrefix:  "/iam.IAMService/",
		PublicMethods: []string{"RootUserLogin", "SignUp", "UserLogin"},
		Mannager:      &auth.JWTMannger{SecretKey: "supersecret", Duration: time.Hour}}

	s := grpc.NewServer(grpc.ChainUnaryInterceptor(authInt.Unary()))

	aws.RegisterIAMServiceServer(s, &IAMService{storage: &IamRepository{db: db}, logger: logger, auth: &authInt})

	reflection.Register(s)

	logger.Info().Str("server", "iam").Msg("Staring server on port :6000")

	wrapped := grpcweb.WrapServer(s, grpcweb.WithOriginFunc(func(origin string) bool {
		logger.Info().Str("origin", origin).Msg("Origin: " + origin)
		return true
	}))

	handler := func(resp http.ResponseWriter, req *http.Request) {
		wrapped.ServeHTTP(resp, req)
	}

	httpsServer := http.Server{
		Addr:    ":7000",
		Handler: http.HandlerFunc(handler),
	}

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Println(err)
		}
	}()

	if err := httpsServer.ListenAndServe(); err != nil {
		fmt.Println(err)
		return nil
	}
	return err
}

func (s *IAMService) Error(err error, code codes.Code, msg string) error {
	s.logger.Err(err)
	return grpc.Errorf(code, msg)
}

func runMigrations(db *gorm.DB) {
	db.AutoMigrate(model.RootUser{}, model.Group{}, model.Role{}, model.Policy{}, model.User{}, model.AccessKey{})
}
