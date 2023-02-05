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
		requestBody := new(entities.CornellNoteRequest)

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

		var subjects []entities.Subject
		var categories []entities.Category

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

			requestBody.Subjects[i].Description = strings.TrimSpace(requestBody.Subjects[i].Description)

			currSubject := entities.Subject{
				Name:        requestBody.Subjects[i].Name,
				Description: requestBody.Subjects[i].Description,
			}

			subjects = append(subjects, currSubject)
		}

		clientDB.
			Create(&subjects)

		fd := requestBody.Categories

		for i, categ := range fd {
			o := categ

			o.Name = strings.TrimSpace(
				o.Name,
			)

			o.Name = strings.ToLower(
				o.Name,
			)

			if len(o.Name) == 0 {
				return context.Status(fiber.StatusBadRequest).
					JSON(fiber.Map{
						"success": false,
						"content": nil,
						"error":   fmt.Sprintf("Category in position %v has an empty Name", i+1),
					})
			}

			o.Description = strings.TrimSpace(o.Description)

			currCategory := entities.Category{
				Name:        o.Name,
				Description: o.Description,
			}
			categories = append(categories, currCategory)
		}

		clientDB.Create(&categories)

		var ts []entities.Subject
		var cs []entities.Category

		for _, t := range subjects {
			clientDB.First(&t, "name=?", t.Name)

			ts = append(ts, t)
		}

		for _, c := range categories {
			clientDB.First(&c, "name=?", c.Name)

			cs = append(cs, c)
		}

		cornellNote := entities.CornellNote{
			Topic:      requestBody.Topic,
			Subjects:   ts,
			Categories: cs,
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
