package cornellNote

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gusramirez-aplazo/simple-english-notes/modules/category"
	"github.com/gusramirez-aplazo/simple-english-notes/modules/note"
	"github.com/gusramirez-aplazo/simple-english-notes/modules/shared/entities"
	"github.com/gusramirez-aplazo/simple-english-notes/modules/shared/infra"
	"github.com/gusramirez-aplazo/simple-english-notes/modules/subject"
	"github.com/gusramirez-aplazo/simple-english-notes/modules/topic"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strings"
)

func creationControllerFactory(
	clientDB *gorm.DB,
	topicRepo *topic.Repository,
	subjectRepo *subject.Repository,
	categoryRepo *category.Repository,
	noteRepo *note.Repository,
) func(*fiber.Ctx) error {
	return func(context *fiber.Ctx) error {
		requestBody := new(entities.CornellNote)

		if err := context.BodyParser(requestBody); err != nil {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   fmt.Sprintf("Ensure the correct data is sent to the server. Error: %v", err.Error()),
			})
		}

		requestBody.Topic.Name = strings.TrimSpace(requestBody.Topic.Name)
		requestBody.Topic.Name = strings.ToLower(requestBody.Topic.Name)

		if requestBody.Topic.Name == "" {
			return context.Status(fiber.StatusBadRequest).
				JSON(fiber.Map{
					"success": false,
					"content": nil,
					"error":   "Topic is required",
				})
		}

		topicRepo.GetItemByName(&requestBody.Topic)

		if requestBody.Topic.TopicID != 0 {
			return context.Status(fiber.StatusBadRequest).
				JSON(fiber.Map{
					"success": false,
					"content": nil,
					"error":   "The requested topic name is already created, try to update it instead",
				})
		}

		if len(requestBody.Subjects) == 0 {
			return context.Status(fiber.StatusBadRequest).
				JSON(fiber.Map{
					"success": false,
					"content": nil,
					"error":   "Add at least 1 subject",
				})
		}

		if len(requestBody.Categories) == 0 {
			return context.Status(fiber.StatusBadRequest).
				JSON(fiber.Map{
					"success": false,
					"content": nil,
					"error":   "Add at least 1 category",
				})
		}

		if len(requestBody.Notes) == 0 {
			return context.Status(fiber.StatusBadRequest).
				JSON(fiber.Map{
					"success": false,
					"content": nil,
					"error":   "Add at least 1 note",
				})
		}

		for i := 0; i < len(requestBody.Subjects); i++ {
			if requestBody.Subjects[i].SubjectID == 0 {
				return context.Status(fiber.StatusBadRequest).
					JSON(fiber.Map{
						"success": false,
						"content": nil,
						"error": fmt.Sprintf(
							"Subject %v does not have ID",
							requestBody.Subjects[i].Name,
						),
					})
			}

			requestBody.Subjects[i].Name = strings.TrimSpace(
				requestBody.Subjects[i].Name,
			)

			requestBody.Subjects[i].Name = strings.ToLower(
				requestBody.Subjects[i].Name,
			)

			if len(requestBody.Subjects[i].Name) == 0 {
				return context.Status(fiber.StatusBadRequest).
					JSON(fiber.Map{
						"success": false,
						"content": nil,
						"error":   fmt.Sprintf("Subject in position %v has an empty Name", i+1),
					})
			}

			subjectRepo.GetItemById(&requestBody.Subjects[i])
		}

		for i := 0; i < len(requestBody.Categories); i++ {
			if requestBody.Categories[i].ID == 0 {
				return context.Status(fiber.StatusBadRequest).
					JSON(fiber.Map{
						"success": false,
						"content": nil,
						"error": fmt.Sprintf(
							"Category %v does not have ID",
							requestBody.Categories[i].Name,
						),
					})
			}
			requestBody.Categories[i].Name = strings.TrimSpace(
				requestBody.Categories[i].Name,
			)

			requestBody.Categories[i].Name = strings.ToLower(
				requestBody.Categories[i].Name,
			)

			if len(requestBody.Categories[i].Name) == 0 {
				return context.Status(fiber.StatusBadRequest).
					JSON(fiber.Map{
						"success": false,
						"content": nil,
						"error":   fmt.Sprintf("Category in position %v has an empty Name", i+1),
					})
			}

			item, findErr := categoryRepo.
				GetItemById(requestBody.Categories[i].ID)

			if findErr != nil {
				return infra.CustomResponse(
					context,
					fiber.StatusInternalServerError,
					false,
					nil,
					findErr.Error(),
				)
			}

			if item.ID == 0 {
				notFoundErr := errors.New(
					fmt.Sprintf(
						"Category at position %v not found",
						i+1,
					),
				)

				return infra.CustomResponse(
					context,
					fiber.StatusNotFound,
					false,
					nil,
					notFoundErr.Error(),
				)
			}

			requestBody.Categories[i] = item
		}

		for i := 0; i < len(requestBody.Notes); i++ {
			if requestBody.Notes[i].NoteID == 0 {
				return context.Status(fiber.StatusBadRequest).
					JSON(fiber.Map{
						"success": false,
						"content": nil,
						"error": fmt.Sprintf(
							"Note at position %v does not have ID",
							i+1,
						),
					})
			}

			requestBody.Notes[i].Content = strings.TrimSpace(
				requestBody.Notes[i].Content,
			)

			requestBody.Notes[i].Cue = strings.TrimSpace(
				requestBody.Notes[i].Cue,
			)

			if len(requestBody.Notes[i].Content) == 0 {
				return context.Status(fiber.StatusBadRequest).
					JSON(fiber.Map{
						"success": false,
						"content": nil,
						"error": fmt.Sprintf(
							"Note content at position %v is empty",
							i+1,
						),
					})
			}

			noteRepo.GetItem(&requestBody.Notes[i])
		}

		cornellNote := entities.CornellNote{
			Topic:      requestBody.Topic,
			Subjects:   requestBody.Subjects,
			Categories: requestBody.Categories,
			Notes:      requestBody.Notes,
		}

		dbCreationResponse := clientDB.
			Create(
				&cornellNote,
			)

		if dbCreationResponse.Error != nil {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   dbCreationResponse.Error.Error(),
			})
		}

		return context.Status(fiber.StatusCreated).JSON(fiber.Map{
			"success": true,
			"content": cornellNote,
			"error":   nil,
		})
	}
}

func getAllCornellNoteControllerFactory(
	clientDB *gorm.DB,
) func(*fiber.Ctx) error {
	return func(context *fiber.Ctx) error {
		var cornellNotes []entities.CornellNote

		clientDB.
			Preload(clause.Associations).
			Find(&cornellNotes)

		return context.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"content": cornellNotes,
			"error":   nil,
		})
	}
}
