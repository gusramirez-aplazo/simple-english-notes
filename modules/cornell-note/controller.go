package cornellNote

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gusramirez-aplazo/simple-english-notes/modules/category"
	"github.com/gusramirez-aplazo/simple-english-notes/modules/shared/entities"
	"github.com/gusramirez-aplazo/simple-english-notes/modules/subject"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strings"
)

func createCornellNoteControllerFactory(
	clientDB *gorm.DB,
	subjectRepo *subject.Repository,
	categoryRepo *category.Repository,
) func(*fiber.Ctx) error {
	return func(context *fiber.Ctx) error {
		requestBody := new(entities.CornellNote)

		if err := context.BodyParser(requestBody); err != nil {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   err.Error(),
			})

		}

		requestBody.Topic = strings.TrimSpace(requestBody.Topic)
		requestBody.Topic = strings.ToLower(requestBody.Topic)

		if requestBody.Topic == "" {
			return context.Status(fiber.StatusBadRequest).
				JSON(fiber.Map{
					"success": false,
					"content": nil,
					"error":   "Topic/Title is required",
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

			clientDB.First(
				&requestBody.Subjects[i],
				"name=?",
				requestBody.Subjects[i].Name,
			)

		}

		clientDB.
			Create(&requestBody.Subjects)

		for i := 0; i < len(requestBody.Categories); i++ {

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

			clientDB.First(
				&requestBody.Categories[i],
				"name=?",
				requestBody.Categories[i].Name,
			)
		}

		clientDB.Create(&requestBody.Categories)

		for i := 0; i < len(requestBody.Notes); i++ {
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
						"error":   fmt.Sprintf("Note content in position %v is empty", i+1),
					})
			}
		}

		clientDB.
			Create(&requestBody.Notes)

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
