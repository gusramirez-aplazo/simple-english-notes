package category

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gusramirez-aplazo/simple-english-notes/modules/shared/domain"
	"gorm.io/gorm"
	"log"
)

//type CategoryContainer struct {
//	repo *domain.BaseRepository
//	cont *domain.BaseController
//	db   *gorm.DB
//}
//
//func (category *CategoryContainer) Init(
//	repository *domain.BaseRepository,
//	controller *domain.BaseController,
//	clientDB *gorm.DB,
//) {
//
//}

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

	router.Get(basePath+"/:categoryId", controller.getOneById)

	router.Get(basePath+"/name/:categoryName", controller.getOneByUniqueParam)

	router.Put(basePath+"/:categoryId", controller.updateOne)

	router.Delete(basePath+"/:categoryId", controller.DeleteOne)
}
