package model

import (
	"time"

	"github.com/infobloxopen/atlas-app-toolkit/rpc/resource"
)

// Base contains common columns for all tables.
type Base struct {
	ID        *resource.Identifier `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time            `json:"created_at"`
	UpdatedAt time.Time            `json:"update_at"`
	DeletedAt *time.Time           `sql:"index" json:"deleted_at"`
}
