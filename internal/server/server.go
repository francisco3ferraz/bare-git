package server

import (
	"database/sql"
	"net/http"

	"github.com/francisco3ferraz/bare-git/internal/config"
	"github.com/rs/zerolog"
)

type Server struct {
	config *config.Config
	db     *sql.DB
	logger *zerolog.Logger
}

func NewServer(cfg *config.Config, db *sql.DB, logger *zerolog.Logger) *Server {
	return &Server{
		config: cfg,
		db:     db,
		logger: logger,
	}
}

func (srv *Server) SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/login", srv.login)
	mux.HandleFunc("/register", srv.register)

	return mux
}
