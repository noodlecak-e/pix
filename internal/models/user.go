package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID        string         `json:"id"`
	EventID   sql.NullString `json:"event_id"`
	Name      sql.NullString `json:"name"`
	CreatedAt time.Time      `json:"created_at,omitempty"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
	DeletedAt sql.NullTime   `json:"deleted_at,omitempty"`
}
