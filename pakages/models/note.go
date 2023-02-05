package models

import (
	"gorm.io/gorm"
	"log"
)

type Note struct {
	gorm.Model

	NoteID  uint   `gorm:"primaryKey; autoIncrement; not null; unique_index" json:"noteId"`
	Content string `gorm:"not null" json:"content"`
}

func (note Note) RunMigration(clientDB *gorm.DB) {
	noteErr := clientDB.AutoMigrate(Note{})

	if noteErr != nil {
		log.Fatal(noteErr)
	}
}
