package models

import "gorm.io/gorm"

type RelevantQuestion struct {
	gorm.Model

	ID    uint   `gorm:"primaryKey; autoIncrement; not null; unique_index"`
	Title string `gorm:"not null"`
	Notes []Note `gorm:"foreignKey:ID"`
}
