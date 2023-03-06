package category

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
	migrationErr := clientDB.AutoMigrate(domain.Category{})

	if migrationErr != nil {
		log.Fatal(migrationErr)
	}

	repo := GetRepository(clientDB)

	controller := GetController(repo)

	const basePath = "/category"

	router.Post(basePath, controller.createOne)

	router.Get(basePath, controller.getAll)

	router.Get(basePath+"/name", controller.getOneByUniqueParam)

	router.Get(basePath+"/:categoryId", controller.getOneById)

	router.Put(basePath+"/:categoryId", controller.updateOne)

	router.Delete(basePath+"/:categoryId", controller.deleteOne)
}
