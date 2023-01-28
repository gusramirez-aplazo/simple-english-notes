package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CornellNote struct {
	gorm.Model

	ID       uuid.UUID        `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Title    string           `gorm:"not null"`
	Topic    Topic            `gorm:"not null;foreignKey:ID"`
	Category Category         `gorm:"not null;foreignKey:ID"`
	Question RelevantQuestion `gorm:"foreignKey:ID"`
	Prompt   RecallPrompt     `gorm:"foreignKey:ID"`
}
