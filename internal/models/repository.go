package models

import "time"

type Repository struct {
	ID          int       `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	OwnerID     int       `json:"owner_id" db:"owner_id"`
	Owner       *User     `json:"owner,omitempty"`
	IsPrivate   bool      `json:"is_private" db:"is_private"`
	CloneURL    string    `json:"clone_url"`
	Size        int64     `json:"size" db:"size"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
