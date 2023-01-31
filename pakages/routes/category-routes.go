package routes

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gusramirez-aplazo/simple-english-notes/pakages/controllers"
	"gorm.io/gorm"
)

func initCategoryRoutes(
	controller *controllers.Controller,
	router fiber.Router,
	clientDB *gorm.DB,
	validate *validator.Validate,
) {

	createCategoryController := controller.CreateCategoryControllerFactory(clientDB, validate)

	router.Post("/category", createCategoryController)

	getAllCategoriesController := controller.GetAllCategoriesControllerFactory(clientDB)

	router.Get("/category", getAllCategoriesController)

	getCategoryByIdController := controller.GetCategoryByIdControllerFactory(clientDB)

	router.Get("/category/:categoryId", getCategoryByIdController)

	deleteCategoryByIdController := controller.DeleteCategoryByIdControllerFactory(clientDB)

	router.Delete("/category/:categoryId", deleteCategoryByIdController)

	updateCategoryByIdController := controller.UpdateCategoryByIdControllerFactory(clientDB, validate)

	router.Put("/category/:categoryId", updateCategoryByIdController)
}
