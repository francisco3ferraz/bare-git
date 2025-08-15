package main

import (
	"os"

	"github.com/francisco3ferraz/bare-git/internal/config"
	"github.com/francisco3ferraz/bare-git/internal/server"
	"github.com/rs/zerolog"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic("Failed to load configuration: " + err.Error())
	}

	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	// TODO: Manage database connection lifecycle

	srv := server.NewServer(cfg, nil, &logger)
}
