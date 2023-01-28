package models

import (
	"gorm.io/gorm"
)

func RunMigrations(clientDB *gorm.DB) {
	Topic{}.RunMigration(clientDB)

	Category{}.RunMigration(clientDB)

	Note{}.RunMigration(clientDB)

	RecallPrompt{}.RunMigration(clientDB)

	RelevantQuestion{}.RunMigration(clientDB)

	CornellNote{}.RunMigration(clientDB)
}
