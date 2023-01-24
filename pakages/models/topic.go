package models

import "gorm.io/gorm"

type Topic struct {
	gorm.Model

	ID          uint   `gorm:"primaryKey; autoIncrement; not null; unique_index"`
	Name        string `gorm:"not null;unique;type:varchar(100)"`
	Description string
}
