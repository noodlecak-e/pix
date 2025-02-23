package models

import (
	"database/sql"
	"time"
)

type Picture struct {
	ID        string       `json:"id"`
	EventID   string       `json:"event_id"`
	UserID    string       `json:"user_id"`
	ImagePath string       `json:"image_path"`
	CreatedAt time.Time    `json:"created_at,omitempty"`
	UpdatedAt time.Time    `json:"updated_at,omitempty"`
	DeletedAt sql.NullTime `json:"deleted_at,omitempty"`
}
