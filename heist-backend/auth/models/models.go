package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       string `gorm:"primaryKey"`
	TeamName string `json:"name"`
	NameHash string `json:"hash"`
	Points   string `json:"totalPoints"`
}
