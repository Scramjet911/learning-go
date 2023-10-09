package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	ID     string `gorm:"primaryKey;" json:"id"`
	Isbn   string `json:"isbn"`
	Title  string `json:"title"`
	Author string `json:"author"`
}
