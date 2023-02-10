package topic

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

	migrationErr := clientDB.AutoMigrate(entities.Topic{})

	if migrationErr != nil {
		log.Fatal(migrationErr)
	}

	repo := GetRepository(clientDB)

	const basePath = "/topic"

	router.Post(basePath, creationControllerFactory(repo))

	router.Get(basePath, getAllItemsControllerFactory(repo))

	router.Get(basePath+"/:topicId", getItemByIdControllerFactory(repo))

	router.Delete(basePath+"/:topicId", deleteItemByIdControllerFactory(repo))

	router.Put(basePath+"/:topicId", updateItemByIdControllerFactory(repo))
}
