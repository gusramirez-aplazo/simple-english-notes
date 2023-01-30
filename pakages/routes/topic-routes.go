package routes

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gusramirez-aplazo/simple-english-notes/pakages/controllers"
	"gorm.io/gorm"
)

func initTopicRoutes(
	controller *controllers.Controller,
	router fiber.Router,
	clientDB *gorm.DB,
	validate *validator.Validate,
) {

	createTopicController := controller.CreateTopicControllerFactory(clientDB, validate)

	router.Post("/topic", createTopicController)

	getTopicsController := controller.GetTopicsControllerFactory(clientDB)

	router.Get("/topic", getTopicsController)

	getTopicByIdController := controller.GetTopicByIdControllerFactory(clientDB)

	router.Get("/topic/:topicId", getTopicByIdController)

	deleteTopicByIdController := controller.DeleteTopicByIdControllerFactory(clientDB)

	router.Delete("/topic/:topicId", deleteTopicByIdController)

	updateTopicByIdController := controller.UpdateTopicByIdControllerFactory(clientDB, validate)

	router.Put("/topic/:topicId", updateTopicByIdController)
}
