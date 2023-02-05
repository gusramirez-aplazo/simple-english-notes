package models

import (
	"gorm.io/gorm"
	"log"
)

type RelevantQuestion struct {
	gorm.Model

	QuestionID uint   `gorm:"primaryKey;uniqueIndex;autoIncrement;not null" json:"questionId"`
	Title      string `gorm:"not null" validate:"required" json:"title"`
	Notes      []Note `json:"notes"`
}

func (question RelevantQuestion) RunMigration(clientDB *gorm.DB) {
	relevantQuestionErr := clientDB.AutoMigrate(RelevantQuestion{})

	if relevantQuestionErr != nil {
		log.Fatal(relevantQuestionErr)
	}
}
