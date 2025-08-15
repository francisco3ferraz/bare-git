package main

import (
	"github.com/francisco3ferraz/bare-git/internal/config"
	"github.com/francisco3ferraz/bare-git/internal/server"
	"github.com/francisco3ferraz/bare-git/internal/utils"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic("Failed to load configuration: " + err.Error())
	}

	logger := utils.NewLogger(cfg.LogLevel, cfg.Environment)

	// TODO: Manage database connection lifecycle

	srv := server.NewServer(cfg, nil, &logger)
}
