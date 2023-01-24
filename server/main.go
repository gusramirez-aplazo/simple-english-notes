package main

import (
	"fmt"
	"github.com/gusramirez-aplazo/simple-english-notes/pakages/database"
	"github.com/gusramirez-aplazo/simple-english-notes/pakages/models"
)

func init() {
	database.Connect()
	models.RunMigrations()
}

func main() {
	fmt.Println("here")
}
