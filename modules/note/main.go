package note

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gusramirez-aplazo/simple-english-notes/modules/shared/domain"
	"gorm.io/gorm"
	"log"
)

func Start(
	clientDB *gorm.DB,
	router fiber.Router,
) {
	migrationErr := clientDB.AutoMigrate(domain.Note{})

	if migrationErr != nil {
		log.Fatal(migrationErr)
	}

	const basePath = "/note"

	repo := GetRepository(clientDB)

	controller := GetController(repo)

	router.Post(basePath, controller.createOne)

	router.Get(basePath, controller.getAll)

	router.Get(basePath+"/:noteId", controller.getOneById)

	router.Delete(basePath+"/:noteId", controller.deleteOne)

	router.Put(basePath+"/:noteId", controller.updateOne)

}
