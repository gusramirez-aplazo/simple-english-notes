package cornellNote

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gusramirez-aplazo/simple-english-notes/modules/category"
	"github.com/gusramirez-aplazo/simple-english-notes/modules/note"
	"github.com/gusramirez-aplazo/simple-english-notes/modules/shared/domain"
	"github.com/gusramirez-aplazo/simple-english-notes/modules/subject"
	"github.com/gusramirez-aplazo/simple-english-notes/modules/topic"
	"gorm.io/gorm"
	"log"
)

func Start(
	clientDB *gorm.DB,
	router fiber.Router,
) {
	migrationErr := clientDB.AutoMigrate(domain.CornellNote{})

	if migrationErr != nil {
		log.Fatal(migrationErr)
	}

	const basePath = "/cornell"

	router.Post(basePath, creationControllerFactory(
		clientDB,
		topic.GetRepository(clientDB),
		subject.GetRepository(clientDB),
		category.GetRepository(clientDB),
		note.GetRepository(clientDB),
	))

	router.Get(basePath, getAllCornellNoteControllerFactory(clientDB))
}
