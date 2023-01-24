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

}

func main() {
	fmt.Println("here")
}
