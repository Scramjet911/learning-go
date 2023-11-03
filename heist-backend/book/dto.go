package book

import "gorm.io/gorm"

type BookInput struct {
	gorm.Model
	ID     string `gorm:"" json:"id"`
	Isbn   string `gorm:"" json:"isbn"`
	Title  string `gorm:"" json:"title"`
	Author string `gorm:"" json:"author"`
}

type Author struct {
	FirstName string `json:"firstName"`
}
