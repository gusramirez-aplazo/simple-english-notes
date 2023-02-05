package entities

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model

	CategoryID   uint          `gorm:"type:bigserial;primaryKey;not null;autoIncrement;uniqueIndex" json:"categoryId"`
	Name         string        `gorm:"not null;unique;type:varchar(100)" json:"name"`
	Description  string        `json:"description"`
	CornellNotes []CornellNote `gorm:"many2many:cornell_categories"`
}
