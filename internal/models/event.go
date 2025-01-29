package models

import (
	"database/sql"
	"time"
)

type Event struct {
	ID        string       `json:"id"`
	Name      string       `json:"name"`
	CreatedAt time.Time    `json:"created_at,omitempty"`
	UpdatedAt time.Time    `json:"updated_at,omitempty"`
	DeletedAt sql.NullTime `json:"deleted_at,omitempty"`
}
