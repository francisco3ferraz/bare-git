package main

import (
	"github.com/francisco3ferraz/bare-git/internal/auth"
	"github.com/francisco3ferraz/bare-git/internal/config"
	"github.com/francisco3ferraz/bare-git/internal/database"
	"github.com/francisco3ferraz/bare-git/internal/server"
	"github.com/francisco3ferraz/bare-git/internal/utils"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic("Failed to load configuration: " + err.Error())
	}

	logger := utils.NewLogger(cfg.LogLevel, cfg.Environment)

	database, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to connect to the database")
	}

	jwtManager := auth.NewJWTManager(cfg.JWTSecret)

	srv := server.NewServer(cfg, database, &logger, jwtManager)
}
