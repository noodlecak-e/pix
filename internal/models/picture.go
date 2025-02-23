package models

import (
	"database/sql"
	"time"
)

type Picture struct {
	ID        string         `json:"id"`
	EventID   sql.NullString `json:"event_id"`
	UserID    sql.NullString `json:"user_id"`
	ImagePath sql.NullString `json:"image_path"`
	CreatedAt time.Time      `json:"created_at,omitempty"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
	DeletedAt sql.NullTime   `json:"deleted_at,omitempty"`
}
