package models

import (
	"time"
)

type User struct {
	ID        int64     `json:"id"`
	TeamID    int64     `json:"team_id"`
	Name      string    `json:"name"`
	IsAdmin   bool      `json:"is_admin"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}
