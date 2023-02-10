package note

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gusramirez-aplazo/simple-english-notes/modules/shared/entities"
	"gorm.io/gorm"
	"log"
)

func Start(
	clientDB *gorm.DB,
	router fiber.Router,
) {
	migrationErr := clientDB.AutoMigrate(entities.Note{})

	if migrationErr != nil {
		log.Fatal(migrationErr)
	}

	repo := GetRepository(clientDB)

	const basePath = "/note"

	router.Post(basePath, creationControllerFactory(repo))

	router.Get(basePath, getAllItemsControllerFactory(repo))

	router.Get(basePath+"/:noteId", getItemByIdControllerFactory(repo))

	router.Delete(basePath+"/:noteId", deleteItemByIdControllerFactory(repo))

	router.Put(basePath+"/:noteId", updateItemByIdControllerFactory(repo))

}
