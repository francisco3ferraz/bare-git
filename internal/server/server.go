package server

import (
	"database/sql"
	"net/http"

	"github.com/francisco3ferraz/bare-git/internal/auth"
	"github.com/francisco3ferraz/bare-git/internal/config"
	"github.com/rs/zerolog"
)

type Server struct {
	config     *config.Config
	db         *sql.DB
	jwtManager *auth.JWTManager
	logger     *zerolog.Logger
}

func NewServer(cfg *config.Config, db *sql.DB, logger *zerolog.Logger, jwtManager *auth.JWTManager) *Server {
	return &Server{
		config:     cfg,
		db:         db,
		jwtManager: jwtManager,
		logger:     logger,
	}
}

func (srv *Server) SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/login", srv.login)
	mux.HandleFunc("/register", srv.register)

	return mux
}
