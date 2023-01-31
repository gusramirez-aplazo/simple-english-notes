package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
)

type CornellNote struct {
	gorm.Model

	ID       uuid.UUID        `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Title    string           `gorm:"not null" validate:"required"`
	Topic    Topic            `gorm:"not null;foreignKey:ID"`
	Category Category         `gorm:"not null;foreignKey:ID"`
	Question RelevantQuestion `gorm:"foreignKey:ID" validate:"required"`
	Prompt   RecallPrompt     `gorm:"foreignKey:ID" validate:"required"`
}

func (note CornellNote) RunMigration(clientDB *gorm.DB) {
	cornellNoteErr := clientDB.AutoMigrate(CornellNote{})

	if cornellNoteErr != nil {
		log.Fatal(cornellNoteErr)
	}
}
