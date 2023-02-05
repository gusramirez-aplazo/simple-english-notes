package models

import (
	"gorm.io/gorm"
	"log"
)

type RecallPrompt struct {
	gorm.Model

	PromptID uint   `gorm:"primaryKey;uniqueIndex;autoIncrement;not null" json:"promptId"`
	Title    string `gorm:"not null;unique;type:varchar(100)" validate:"required" json:"title"`
	Notes    []Note `json:"notes"`
}

func (prompt RecallPrompt) RunMigration(clientDB *gorm.DB) {
	recallPromptErr := clientDB.AutoMigrate(RecallPrompt{})

	if recallPromptErr != nil {
		log.Fatal(recallPromptErr)
	}
}
