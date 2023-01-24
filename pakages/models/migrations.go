package models

import (
	"github.com/gusramirez-aplazo/simple-english-notes/pakages/database"
	"log"
)

func RunMigrations() {
	topicMigrationErr := database.Client.AutoMigrate(Topic{})

	if topicMigrationErr != nil {
		log.Fatal(topicMigrationErr)
	}

	recallPromptErr := database.Client.AutoMigrate(RecallPrompt{})

	if recallPromptErr != nil {
		log.Fatal(recallPromptErr)
	}

	relevantQuestionErr := database.Client.AutoMigrate(RelevantQuestion{})

	if relevantQuestionErr != nil {
		log.Fatal(relevantQuestionErr)
	}
}