package server

import "github.com/francisco3ferraz/bare-git/internal/models"

func (s *Server) getUserByUsername(username string) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, username, email, password_hash, is_admin, created_at, updated_at 
              FROM users WHERE username = $1`

	err := s.db.QueryRow(query, username).Scan(
		&user.ID, &user.Username, &user.Email, &user.Password,
		&user.IsAdmin, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Server) getUserByID(id int) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, username, email, password_hash, is_admin, created_at, updated_at 
              FROM users WHERE id = $1`

	err := s.db.QueryRow(query, id).Scan(
		&user.ID, &user.Username, &user.Email, &user.Password,
		&user.IsAdmin, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Server) createUser(user *models.User) error {
	query := `INSERT INTO users (username, email, password_hash, is_admin) 
              VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at`

	return s.db.QueryRow(query, user.Username, user.Email, user.Password, user.IsAdmin).
		Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
}
