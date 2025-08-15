package utils

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func NewLogger(level, environment string) zerolog.Logger {
	logLevel, err := zerolog.ParseLevel(level)
	if err != nil {
		logLevel = zerolog.InfoLevel
	}
	zerolog.SetGlobalLevel(logLevel)

	var logger zerolog.Logger
	if environment == "development" {
		logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	} else {
		logger = log.Logger
	}

	return logger
}
