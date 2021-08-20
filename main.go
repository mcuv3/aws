package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/MauricioAntonioMartinez/aws/iam"
	"github.com/MauricioAntonioMartinez/aws/lambda"
	"github.com/MauricioAntonioMartinez/aws/sqs"
	"github.com/rs/zerolog"
)


var (
	service =  flag.String("svc","","The aws service to start")

)

func main() {
	flag.Parse()
 
	fmt.Println(*service)
	// fastergoding.Run()


	
	svc := *service 

	if svc == "" {
		svc = os.Getenv("SERVICE")
	} 

	l := logger()
	l.Info().Str("svc",os.Getenv("SERVICE"))
	
		
	switch(svc) {  
	case "sqs":
		if err := sqs.Run(l); err !=nil {
			l.Fatal().Msg("Unable to start sqs service.")
		}
		break
	case "iam":
		if err := iam.Run(l); err !=nil {
			l.Fatal().Msg("Unable to start iam service.")
		}
	case "lambda":
		if err := lambda.Run(l); err !=nil {
			l.Fatal().Msg("Unable to start lambda service.")
		}
	}
}

func logger() zerolog.Logger { 
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	w := zerolog.ConsoleWriter{Out: os.Stderr}
	l := zerolog.New(w).With().Timestamp().Caller().Logger()
	return l 
}

