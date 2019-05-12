package model

import "github.com/reviewsys/backend/app/interface/rpc/v1.0/protocol"

type User struct {
	protocol.UserORM
	// Base
	// TeamID  int64  `json:"team_id"`
	// Name    string `json:"name"`
	// IsAdmin bool   `gorm:";not null;default:false" json:"is_admin"`
}
