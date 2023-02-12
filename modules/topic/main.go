package topic

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
	migrationErr := clientDB.AutoMigrate(domain.Topic{})

	if migrationErr != nil {
		log.Fatal(migrationErr)
	}

	const basePath = "/topic"

	repo := GetRepository(clientDB)

	controller := GetController(repo)

	router.Post(basePath, controller.createOne)

	router.Get(basePath, controller.getAll)

	router.Get(basePath+"/name", controller.getOneByUniqueParam)

	router.Get(basePath+"/:topicId", controller.getOneById)

	router.Delete(basePath+"/:topicId", controller.deleteOne)

	router.Put(basePath+"/:topicId", controller.updateOne)
}
