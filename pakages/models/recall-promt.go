package models

import (
	"gorm.io/gorm"
	"log"
)

type RecallPrompt struct {
	gorm.Model

	ID    uint   `gorm:"primaryKey; autoIncrement; not null; unique_index"`
	Title string `gorm:"not null;unique;type:varchar(100)"`
	Notes []Note `gorm:"foreignKey:ID"`
}

func (prompt RecallPrompt) RunMigration(clientDB *gorm.DB) {
	recallPromptErr := clientDB.AutoMigrate(RecallPrompt{})

	if recallPromptErr != nil {
		log.Fatal(recallPromptErr)
	}
}
