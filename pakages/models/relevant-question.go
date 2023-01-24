package models

import "gorm.io/gorm"

type RelevantQuestion struct {
	gorm.Model

	ID      uint   `gorm:"primaryKey; autoIncrement; not null; unique_index"`
	Content string `gorm:"not null; text"`
	Notes   []Note `gorm:"foreignKey:ID"`
}
