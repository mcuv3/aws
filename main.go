package main

import (
	"os"
	"sync"

	"github.com/MauricioAntonioMartinez/aws/iam"
	"github.com/MauricioAntonioMartinez/aws/sqs"
	"github.com/qinains/fastergoding"
	"github.com/rs/zerolog"
)

func main() {
	fastergoding.Run()
	l := logger()
	
	var wg sync.WaitGroup
	wg.Add(2)

	go func(){
		defer wg.Done()
		if err := sqs.Run(l); err !=nil {
			l.Fatal().Msg("Unable to start sqs service.")
		}
	}()

	go func (){
		defer wg.Done()
		if err := iam.Run(l); err !=nil {
			l.Fatal().Msg("Unable to start iam service.")
		}

	}()
	wg.Wait()
}

func logger() zerolog.Logger { 
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	w := zerolog.ConsoleWriter{Out: os.Stderr}
	l := zerolog.New(w).With().Timestamp().Caller().Logger()
	return l 
}

