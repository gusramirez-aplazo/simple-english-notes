package category

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
	migrationErr := clientDB.AutoMigrate(entities.Category{})

	if migrationErr != nil {
		log.Fatal(migrationErr)
	}

	repo := GetRepository(clientDB)

	const basePath = "/category"

	router.Post(basePath, creationControllerFactory(repo))

	router.Get(basePath, getAllItemsControllerFactory(repo))

	router.Get(basePath+"/:categoryId", getItemByIdControllerFactory(repo))

	router.Get(basePath+"/name/:categoryName", getItemByNameControllerFactory(repo))

	router.Put(basePath+"/:categoryId", updateItemByIdControllerFactory(repo))

	router.Delete(basePath+"/:categoryId", deleteItemByIdControllerFactory(repo))
}
