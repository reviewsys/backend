package model

type User struct {
	Base
	TeamID  int64  `json:"team_id"`
	Name    string `json:"name"`
	IsAdmin bool   `gorm:";not null;default:false" json:"is_admin"`
}
