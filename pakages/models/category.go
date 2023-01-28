package models

import (
	"gorm.io/gorm"
	"log"
)

type Category struct {
	gorm.Model

	ID          uint   `gorm:"primaryKey;autoIncrement;not null;unique_index"`
	Name        string `gorm:"not null;unique;type:varchar(100)"`
	Description string
}

func (category Category) RunMigration(clientDB *gorm.DB) {
	categoryErr := clientDB.AutoMigrate(Category{})

	if categoryErr != nil {
		log.Fatal(categoryErr)
	}
}
