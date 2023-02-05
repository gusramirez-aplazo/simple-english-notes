package cornellNote

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gusramirez-aplazo/simple-english-notes/modules/category"
	"github.com/gusramirez-aplazo/simple-english-notes/modules/shared/entities"
	"github.com/gusramirez-aplazo/simple-english-notes/modules/subject"
	"gorm.io/gorm"
	"log"
)

func StartCornellNoteModule(
	clientDB *gorm.DB,
	router fiber.Router,
) {
	migrationErr := clientDB.AutoMigrate(entities.CornellNote{})

	if migrationErr != nil {
		log.Fatal(migrationErr)
	}

	const basePath = "/cornell"

	router.Post(basePath, createCornellNoteControllerFactory(
		clientDB,
		subject.GetSubjectRepository(clientDB),
		category.GetCategoryRepository(clientDB),
	))

	router.Get(basePath, getAllCornellNoteControllerFactory(clientDB))
}
