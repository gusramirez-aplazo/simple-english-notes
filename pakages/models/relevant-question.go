package models

import (
	"gorm.io/gorm"
	"log"
)

type RelevantQuestion struct {
	gorm.Model

	ID    uint   `gorm:"primaryKey; autoIncrement; not null; unique_index"`
	Title string `gorm:"not null"`
	Notes []Note `gorm:"foreignKey:ID"`
}

func (question RelevantQuestion) RunMigration(clientDB *gorm.DB) {
	relevantQuestionErr := clientDB.AutoMigrate(RelevantQuestion{})

	if relevantQuestionErr != nil {
		log.Fatal(relevantQuestionErr)
	}
}
