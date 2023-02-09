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

	router.Get(basePath+"/:categoryId", getCategoryByIdControllerFactory(repo))

	router.Get(basePath+"/name/:categoryName", getCategoryByNameControllerFactory(repo))

	router.Put(basePath+"/:categoryId", updateCategoryByIdControllerFactory(repo))

	router.Delete(basePath+"/:categoryId", deleteCategoryByIdControllerFactory(repo))
}
