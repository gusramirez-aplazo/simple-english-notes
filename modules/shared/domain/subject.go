package domain

import (
	"gorm.io/gorm"
	"time"
)

type Subject struct {
	ID        uint           `gorm:"type:bigserial;primaryKey;not null;autoIncrement;unique;uniqueIndex" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	Name      string         `gorm:"not null;unique;type:varchar(100)"  json:"name"`
}
