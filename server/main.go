package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gusramirez-aplazo/simple-english-notes/pakages/controllers"
	"github.com/gusramirez-aplazo/simple-english-notes/pakages/database"
	"github.com/gusramirez-aplazo/simple-english-notes/pakages/models"
	"github.com/gusramirez-aplazo/simple-english-notes/pakages/routes"
	"log"
)

var validate = validator.New()

func init() {
	database.Connect()
	models.RunMigrations(database.GetDbClient())
}

func main() {
	const port = 3000

	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	controller := controllers.GetController()

	routes.Start(app, controller, database.GetDbClient(), validate)

	log.Printf("Server listen on port: %v", port)

	address := fmt.Sprintf(":%v", port)

	log.Fatal(app.Listen(address))
}
