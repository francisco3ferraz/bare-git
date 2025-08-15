package server

import (
	"encoding/json"
	"net/http"

	"github.com/francisco3ferraz/bare-git/internal/models"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type CreateRepositoryRequest struct {
	Name        string `json:"name" binding:"required,min=1,max=100"`
	Description string `json:"description"`
	IsPrivate   bool   `json:"is_private"`
}

func (srv *Server) login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := srv.getUserByUsername(req.Username)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	if isEqual := user.CheckPassword(req.Password); !isEqual {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	token, err := srv.jwtManager.GenerateToken(user.ID, user.Username, user.IsAdmin)
	if err != nil {
		srv.logger.Error().Err(err).Msg("Failed to generate token")
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"token": token,
		"user":  user,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		srv.logger.Error().Err(err).Msg("Failed to encode response")
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (srv *Server) register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	existingUser, _ := srv.getUserByUsername(req.Username)
	if existingUser != nil {
		http.Error(w, "Username already exists", http.StatusConflict)
		return
	}

	user := &models.User{
		Username: req.Username,
		Password: req.Password,
	}

	if err := user.HashPassword(req.Password); err != nil {
		srv.logger.Error().Err(err).Msg("Failed to hash password")
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	if err := srv.createUser(user); err != nil {
		srv.logger.Error().Err(err).Msg("Failed to create user")
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	token, err := srv.jwtManager.GenerateToken(user.ID, user.Username, user.IsAdmin)
	if err != nil {
		srv.logger.Error().Err(err).Msg("Failed to generate token")
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"token": token,
		"user":  user,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		srv.logger.Error().Err(err).Msg("Failed to encode response")
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
