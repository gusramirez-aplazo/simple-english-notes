package models

import (
	"gorm.io/gorm"
	"log"
)

type Topic struct {
	gorm.Model

	ID          uint   `gorm:"primaryKey; autoIncrement; not null; unique_index"`
	Name        string `gorm:"not null;unique;type:varchar(100)" validate:"required"`
	Description string `validate:"required"`
}

func (topic Topic) RunMigration(clientDB *gorm.DB) {
	topicMigrationErr := clientDB.AutoMigrate(Topic{})

	if topicMigrationErr != nil {
		log.Fatal(topicMigrationErr)
	}
}
