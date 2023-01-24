package main

import (
	"fmt"
	"github.com/gusramirez-aplazo/simple-english-notes/pakages/database"
	"github.com/gusramirez-aplazo/simple-english-notes/pakages/models"
	"log"
)

func init() {
	database.Connect()

	topicMigrationErr := database.Client.AutoMigrate(models.Topic{})

	if topicMigrationErr != nil {
		log.Fatal(topicMigrationErr)
	}

	recallPromptErr := database.Client.AutoMigrate(models.RecallPrompt{})

	if recallPromptErr != nil {
		log.Fatal(recallPromptErr)
	}

	relevantQuestionErr := database.Client.AutoMigrate(models.RelevantQuestion{})

	if relevantQuestionErr != nil {
		log.Fatal(relevantQuestionErr)
	}
}

func main() {
	fmt.Println("here")
}
