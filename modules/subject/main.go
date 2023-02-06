package subject

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gusramirez-aplazo/simple-english-notes/modules/shared/entities"
	"gorm.io/gorm"
	"log"
)

func StartSubjectModule(
	clientDB *gorm.DB,
	router fiber.Router,
) {
	migrationErr := clientDB.AutoMigrate(entities.Subject{})

	if migrationErr != nil {
		log.Fatal(migrationErr)
	}

	const basePath = "/subject"

	repo := GetSubjectRepository(clientDB)

	router.Post(basePath, createSubjectControllerFactory(repo))

	router.Get(basePath, getSubjectsControllerFactory(repo))

	router.Get(basePath+"/:topicId", getSubjectByIdControllerFactory(repo))

	router.Put(basePath+"/:topicId", updateSubjectByIdControllerFactory(repo))

	router.Delete(basePath+"/:topicId", deleteSubjectByIdControllerFactory(repo))
}
