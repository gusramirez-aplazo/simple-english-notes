package models

import (
	"gorm.io/gorm"
	"log"
)

type Note struct {
	gorm.Model

	ID      uint   `gorm:"primaryKey; autoIncrement; not null; unique_index"`
	Content string `gorm:"not null"`
}

func (note Note) RunMigration(clientDB *gorm.DB) {
	noteErr := clientDB.AutoMigrate(Note{})

	if noteErr != nil {
		log.Fatal(noteErr)
	}
}
