package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gusramirez-aplazo/simple-english-notes/pakages/controllers"
	"gorm.io/gorm"
)

// TODO: handle not found routes
func Start(fiberApp *fiber.App, controller *controllers.Controller, clientDB *gorm.DB) {
	const apiPrefix = "/api/v1"

	fiberApp.Get("/", controller.HomeController)
	//router.HandleFunc("/", controller.HomeController).Methods("GET")
	//
	//subRouter := router.PathPrefix(apiPrefix).Subrouter()
	//
	//createTopicController := controller.CreateTopicControllerFactory(clientDB)
	//
	//subRouter.HandleFunc("/topic", createTopicController).Methods("POST")
}
