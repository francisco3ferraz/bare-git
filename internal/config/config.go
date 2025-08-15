package config

import (
	"errors"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Environment   string `json:"environment"`
	Port          int    `json:"port"`
	DatabaseURL   string `json:"database_url"`
	JWTSecret     string `json:"jwt_secret"`
	SessionSecret string `json:"session_secret"`
	GitReposPath  string `json:"git_repos_path"`
	MaxRepoSize   int    `json:"max_repo_size"`
	LogLevel      string `json:"log_level"`
}

func LoadConfig() (*Config, error) {
	var cfg Config

	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	cfg.Environment = getEnvString("ENVIRONMENT", "development")
	cfg.Port = getEnvInt("PORT", 8080)
	cfg.DatabaseURL = getEnvString("DATABASE_URL", "")
	cfg.JWTSecret = getEnvString("JWT_SECRET", "")
	cfg.SessionSecret = getEnvString("SESSION_SECRET", "")
	cfg.GitReposPath = getEnvString("GIT_REPOS_PATH", "./repositories")
	cfg.MaxRepoSize = getEnvInt("MAX_REPO_SIZE", 1024*1024*1024)
	cfg.LogLevel = getEnvString("LOG_LEVEL", "info")

	if err := cfg.validateConfig(); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (cfg *Config) validateConfig() error {
	if cfg.DatabaseURL == "" {
		return errors.New("DATABASE_URL is required")
	}

	if cfg.JWTSecret == "" {
		return errors.New("JWT_SECRET is required")
	}

	return nil
}

func getEnvString(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getEnvInt(key string, defaultValue int) int {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultValue
	}
	return value
}
