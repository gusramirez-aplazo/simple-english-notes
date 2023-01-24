package models

import "gorm.io/gorm"

type Note struct {
	gorm.Model

	ID          uint   `gorm:"primaryKey; autoIncrement; not null; unique_index"`
	Title       string `gorm:"type:varchar(100)"`
	Description string
}
