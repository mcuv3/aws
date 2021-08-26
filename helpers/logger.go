package helpers

import (
	"os"

	"github.com/rs/zerolog"
)

func NewLogger() zerolog.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	w := zerolog.ConsoleWriter{Out: os.Stderr}
	l := zerolog.New(w).With().Timestamp().Caller().Logger()
	return l
}
