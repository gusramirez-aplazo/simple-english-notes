package subject

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
	migrationErr := clientDB.AutoMigrate(entities.Subject{})

	if migrationErr != nil {
		log.Fatal(migrationErr)
	}

	const basePath = "/subject"

	repo := GetRepository(clientDB)

	router.Post(basePath, creationControllerFactory(repo))

	router.Get(basePath, getAllItemsControllerFactory(repo))

	router.Get(basePath+"/:subjectId", getItemByIdControllerFactory(repo))

	router.Get(basePath+"/name/:subjectName", getItemByNameControllerFactory(repo))

	router.Put(basePath+"/:subjectId", updateItemByIdControllerFactory(repo))

	router.Delete(basePath+"/:subjectId", deleteItemByIdControllerFactory(repo))
}
