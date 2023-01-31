package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gusramirez-aplazo/simple-english-notes/pakages/controllers"
	"gorm.io/gorm"
)

func initNoteRoutes(
	controller *controllers.Controller,
	router fiber.Router,
	clientDB *gorm.DB,
) {

	createNoteController := controller.CreateNoteControllerFactory(clientDB)

	router.Post("/note", createNoteController)

	getAllNotesController := controller.GetAllNotesControllerFactory(clientDB)

	router.Get("/note", getAllNotesController)

	getNoteByIdController := controller.GetNoteByIdControllerFactory(clientDB)

	router.Get("/note/:noteId", getNoteByIdController)

	deleteNoteByIdController := controller.DeleteNoteByIdControllerFactory(clientDB)

	router.Delete("/note/:noteId", deleteNoteByIdController)

	updateNoteByIdController := controller.UpdateNoteByIdControllerFactory(clientDB)

	router.Put("/note/:noteId", updateNoteByIdController)
}
