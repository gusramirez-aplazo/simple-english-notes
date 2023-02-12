package cornellNote

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gusramirez-aplazo/simple-english-notes/modules/category"
	"github.com/gusramirez-aplazo/simple-english-notes/modules/note"
	"github.com/gusramirez-aplazo/simple-english-notes/modules/shared/domain"
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
		requestBody := new(domain.CornellNote)

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
			return infra.CustomResponse(
				context,
				fiber.StatusBadRequest,
				false,
				nil,
				"Topic is required",
			)
		}

		_, findErr := topicRepo.GetItemByUniqueParam(requestBody.Topic.Name)

		hasExistingItem := findErr == nil

		if hasExistingItem {
			return infra.CustomResponse(
				context,
				fiber.StatusBadRequest,
				false,
				nil,
				"topic is already taken, try with another one")
		}

		if len(requestBody.Subjects) == 0 {
			return infra.CustomResponse(
				context,
				fiber.StatusBadRequest,
				false,
				nil,
				"Add at least 1 subject",
			)
		}

		if len(requestBody.Categories) == 0 {
			return infra.CustomResponse(
				context,
				fiber.StatusBadRequest,
				false,
				nil,
				"Add at least 1 category",
			)
		}

		if len(requestBody.Notes) == 0 {
			return infra.CustomResponse(
				context,
				fiber.StatusBadRequest,
				false,
				nil,
				"Add at least 1 note",
			)
		}

		for i := 0; i < len(requestBody.Subjects); i++ {
			if requestBody.Subjects[i].ID == 0 {
				return infra.CustomResponse(
					context,
					fiber.StatusBadRequest,
					false,
					nil,
					fmt.Sprintf(
						"Subject %v does not have ID",
						requestBody.Subjects[i].Name,
					),
				)
			}

			requestBody.Subjects[i].Name = strings.TrimSpace(
				requestBody.Subjects[i].Name,
			)

			requestBody.Subjects[i].Name = strings.ToLower(
				requestBody.Subjects[i].Name,
			)

			if len(requestBody.Subjects[i].Name) == 0 {
				return infra.CustomResponse(
					context,
					fiber.StatusBadRequest,
					false,
					nil,
					fmt.Sprintf(
						"Subject in position %v has an empty Name",
						i+1,
					),
				)
			}

			item, findErr := subjectRepo.
				GetItemById(requestBody.Subjects[i].ID)

			if findErr != nil {
				notFoundErr := errors.New(
					fmt.Sprintf(
						"subject at position %v not found",
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

			requestBody.Subjects[i] = item
		}

		for i := 0; i < len(requestBody.Categories); i++ {
			if requestBody.Categories[i].ID == 0 {
				return infra.CustomResponse(
					context,
					fiber.StatusBadRequest,
					false,
					nil,
					fmt.Sprintf(
						"Category %v does not have ID",
						requestBody.Categories[i].Name,
					),
				)
			}

			requestBody.Categories[i].Name = strings.TrimSpace(
				requestBody.Categories[i].Name,
			)
			requestBody.Categories[i].Name = strings.ToLower(
				requestBody.Categories[i].Name,
			)

			if len(requestBody.Categories[i].Name) == 0 {
				return infra.CustomResponse(
					context,
					fiber.StatusBadRequest,
					false,
					nil,
					fmt.Sprintf(
						"Category in position %v has an empty Name",
						i+1,
					),
				)
			}

			item, findErr := categoryRepo.
				GetItemById(requestBody.Categories[i].ID)

			if findErr != nil {
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
			if requestBody.Notes[i].ID == 0 {
				return infra.CustomResponse(
					context,
					fiber.StatusBadRequest,
					false,
					nil,
					fmt.Sprintf(
						"Note at position %v does not have ID",
						i+1,
					),
				)
			}

			requestBody.Notes[i].Content = strings.TrimSpace(
				requestBody.Notes[i].Content,
			)
			requestBody.Notes[i].Cue = strings.TrimSpace(
				requestBody.Notes[i].Cue,
			)

			if len(requestBody.Notes[i].Content) == 0 {
				return infra.CustomResponse(
					context,
					fiber.StatusBadRequest,
					false,
					nil,
					fmt.Sprintf(
						"Note content at position %v is empty",
						i+1,
					),
				)
			}

			item, findErr := noteRepo.
				GetItemById(requestBody.Notes[i].ID)

			if findErr != nil {
				notFoundErr := errors.New(
					fmt.Sprintf(
						"Note at position %v not found",
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

			requestBody.Notes[i] = item
		}

		cornellNote := domain.CornellNote{
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
			return infra.CustomResponse(
				context,
				fiber.StatusBadRequest,
				false,
				nil,
				dbCreationResponse.Error.Error(),
			)
		}

		return infra.CustomResponse(
			context,
			fiber.StatusCreated,
			true,
			cornellNote,
			"",
		)
	}
}

func getAllCornellNoteControllerFactory(
	clientDB *gorm.DB,
) func(*fiber.Ctx) error {
	return func(context *fiber.Ctx) error {
		var cornellNotes []domain.CornellNote

		clientDB.
			Preload(clause.Associations).
			Find(&cornellNotes)

		return infra.CustomResponse(
			context,
			fiber.StatusOK,
			true,
			cornellNotes,
			"",
		)
	}
}
