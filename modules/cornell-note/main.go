package cornellNote

import (
	"github.com/gofiber/fiber/v2"
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
	err := clientDB.SetupJoinTable(&domain.CornellNote{}, "Subjects", &domain.CornellSubjects{})
	if err != nil {
		log.Fatal(err)
	}

	migrationErr := clientDB.AutoMigrate(domain.CornellNote{})

	if migrationErr != nil {
		log.Fatal(migrationErr)
	}

	repo := GetRepository(clientDB)
	topicRepo := topic.GetRepository(clientDB)
	subjectRepo := subject.GetRepository(clientDB)
	noteRepo := note.GetRepository(clientDB)
	controller := GetController(repo, topicRepo, subjectRepo, noteRepo)

	const basePath = "/cornell"

	router.Post(basePath, controller.createOne)

	router.Get(basePath, controller.getAll)

	router.Get(basePath+"/:id", controller.getOneById)
}
