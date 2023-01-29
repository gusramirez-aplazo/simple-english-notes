package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gusramirez-aplazo/simple-english-notes/pakages/controllers"
	"github.com/gusramirez-aplazo/simple-english-notes/pakages/database"
	"github.com/gusramirez-aplazo/simple-english-notes/pakages/models"
	"github.com/gusramirez-aplazo/simple-english-notes/pakages/routes"
	"log"
)

func init() {
	database.Connect()
	models.RunMigrations(database.GetDbClient())
}

func main() {
	const port = 3000

	app := fiber.New()

	controller := &controllers.Controller{}

	routes.Start(app, controller, database.GetDbClient())

	fmt.Printf("Server listen on port: %v", port)

	address := fmt.Sprintf(":%v", port)

	log.Fatal(app.Listen(address))
}
