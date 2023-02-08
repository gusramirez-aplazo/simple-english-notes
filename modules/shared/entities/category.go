package entities

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model

	CategoryID uint   `gorm:"type:bigserial;primaryKey;not null;autoIncrement;uniqueIndex" json:"id"`
	Name       string `gorm:"not null;unique;type:varchar(100)" json:"name"`
}

type CategoryResponse struct {
	CategoryID uint   `json:"id"`
	Name       string `json:"name"`
}
