package main

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gusramirez-aplazo/simple-english-notes/modules/category"
	"github.com/gusramirez-aplazo/simple-english-notes/modules/cornell-note"
	"github.com/gusramirez-aplazo/simple-english-notes/modules/subject"
	"github.com/gusramirez-aplazo/simple-english-notes/pakages/database"
	"log"
)

func init() {
	database.Connect()
}

func main() {
	const port = 3000

	fiberConfig := fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	}

	app := fiber.New(fiberConfig)

	app.Get("/", func(context *fiber.Ctx) error {
		return context.JSON(&fiber.Map{
			"success": true,
			"content": "Hello World with Fiber!!",
		})
	})

	const apiPrefix = "/api"
	const versionPrefix = "/v1"

	apiRoutes := app.Group(apiPrefix)

	currentVersionedRoutes := apiRoutes.Group(versionPrefix)

	subject.StartSubjectModule(database.GetDbClient(), currentVersionedRoutes)

	category.StartCategoryModule(database.GetDbClient(), currentVersionedRoutes)

	cornellNote.StartCornellNoteModule(database.GetDbClient(), currentVersionedRoutes)

	log.Printf("Server listen on port: %v", port)

	address := fmt.Sprintf(":%v", port)

	log.Fatal(app.Listen(address))
}
