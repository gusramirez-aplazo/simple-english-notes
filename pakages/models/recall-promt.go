package models

import "gorm.io/gorm"

type RecallPrompt struct {
	gorm.Model

	ID    uint   `gorm:"primaryKey; autoIncrement; not null; unique_index"`
	Title string `gorm:"not null; unique; var(100)"`
}
