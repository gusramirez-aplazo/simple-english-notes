package models

import "gorm.io/gorm"

type Note struct {
	gorm.Model

	ID    uint   `gorm:"primaryKey; autoIncrement; not null; unique_index"`
	Title string `gorm:"var(100)"`
	Text  string `gorm:"string"`
}
