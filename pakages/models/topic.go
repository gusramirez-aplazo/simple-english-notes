package models

import (
	"gorm.io/gorm"
	"log"
)

type Topic struct {
	gorm.Model

	ID          uint   `gorm:"primaryKey; autoIncrement; not null; unique_index" json:"id"`
	Name        string `gorm:"not null;unique;type:varchar(100)" validate:"required" json:"name"`
	Description string `json:"description"`
}

func (topic Topic) RunMigration(clientDB *gorm.DB) {
	topicMigrationErr := clientDB.AutoMigrate(Topic{})

	if topicMigrationErr != nil {
		log.Fatal(topicMigrationErr)
	}
}
