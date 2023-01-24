package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model

	ID          uint   `gorm:"primaryKey; autoIncrement; not null; unique_index"`
	Name        string `gorm:"not null; unique; var(100)"`
	Description string `gorm:"string"`
}
