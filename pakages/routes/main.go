package routes

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gusramirez-aplazo/simple-english-notes/pakages/controllers"
	"gorm.io/gorm"
)

// TODO: handle not found routes
func Start(
	fiberApp *fiber.App,
	controller *controllers.Controller,
	clientDB *gorm.DB,
	validate *validator.Validate,
) {
	const apiPrefix = "/api"
	const v1Prefix = "/v1"

	fiberApp.Get("/", controller.HomeController)

	apiRoutes := fiberApp.Group(apiPrefix)
	version1Routes := apiRoutes.Group(v1Prefix)

	createTopicController := controller.CreateTopicControllerFactory(clientDB, validate)

	version1Routes.Post("/topic", createTopicController)

	getTopicsController := controller.GetTopicsControllerFactory(clientDB)

	version1Routes.Get("/topic", getTopicsController)

	getTopicByIdController := controller.GetTopicByIdControllerFactory(clientDB)

	version1Routes.Get("/topic/:topicId", getTopicByIdController)

	deleteTopicByIdController := controller.DeleteTopicByIdControllerFactory(clientDB)

	version1Routes.Delete("/topic/:topicId", deleteTopicByIdController)
}
