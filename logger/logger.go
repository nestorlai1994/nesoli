package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.UnixDate})
}

func Info(msg string) {
	log.Info().Msg(msg)
}

func Error(msg string, err error) {
	log.Error().Err(err).Msg(msg)
}
