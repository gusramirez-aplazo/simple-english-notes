package domain

import (
	"gorm.io/gorm"
	"time"
)

type Note struct {
	ID        uint           `gorm:"type:bigserial;primaryKey;not null;autoIncrement;unique;uniqueIndex" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	Cue       string         `gorm:"type:varchar(512)" json:"cue"`
	Content   string         `gorm:"not null" json:"content"`
}
