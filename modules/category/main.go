package category

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gusramirez-aplazo/simple-english-notes/modules/shared/entities"
	"gorm.io/gorm"
	"log"
)

func StartCategoryModule(
	clientDB *gorm.DB,
	router fiber.Router,
) {
	migrationErr := clientDB.AutoMigrate(entities.Category{})

	if migrationErr != nil {
		log.Fatal(migrationErr)
	}

	repo := GetCategoryRepository(clientDB)

	const basePath = "/category"

	router.Post(basePath, createCategoryControllerFactory(repo))

	router.Get(basePath, getAllCategoriesControllerFactory(repo))

	router.Get(basePath+"/:topicId", getCategoryByIdControllerFactory(repo))

	router.Put(basePath+"/:topicId", updateCategoryByIdControllerFactory(repo))

	router.Delete(basePath+"/:topicId", deleteCategoryByIdControllerFactory(repo))
}
