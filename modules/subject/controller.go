package subject

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gusramirez-aplazo/simple-english-notes/modules/shared/entities"
	"github.com/gusramirez-aplazo/simple-english-notes/modules/shared/infra"
	"strings"
)

// TODO: testing
func createSubjectControllerFactory(
	repository *Repository,
) func(*fiber.Ctx) error {
	return func(context *fiber.Ctx) error {
		subject := new(entities.Subject)

		if err := context.BodyParser(subject); err != nil {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   err.Error(),
			})
		}

		subject.Name = strings.TrimSpace(subject.Name)
		subject.Name = strings.ToLower(subject.Name)

		if subject.Name == "" {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   "Subject Name is required",
			})
		}

		if err := repository.GetItemOrCreate(subject); err != nil {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   err.Error(),
			})
		}

		return context.Status(fiber.StatusCreated).JSON(fiber.Map{
			"success": true,
			"content": subject,
			"error":   nil,
		})
	}
}

func getSubjectsControllerFactory(
	repository *Repository,
) func(*fiber.Ctx) error {
	return func(context *fiber.Ctx) error {
		var subjects []entities.Subject

		if err := repository.GetAllItems(&subjects); err != nil {
			return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   err,
			})
		}

		if len(subjects) == 0 {
			return context.Status(fiber.StatusOK).JSON(fiber.Map{
				"success": true,
				"content": []fiber.Map{},
				"error":   nil,
			})
		}

		return context.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"content": subjects,
			"error":   nil,
		})
	}
}

func getSubjectByIdControllerFactory(
	repository *Repository,
) func(ctx *fiber.Ctx) error {
	return func(context *fiber.Ctx) error {
		subjectId := context.Params("subjectId")

		parsedId, parsedIdErr := infra.ParseID(subjectId)

		if parsedIdErr != nil {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   parsedIdErr.Error(),
			})
		}

		var subject = entities.Subject{SubjectID: parsedId}

		repository.GetItem(&subject)

		if subject.Name == "" {
			return context.Status(fiber.StatusNotFound).JSON(&fiber.Map{
				"success": false,
				"content": nil,
				"error":   fmt.Sprintf("ID %v not found", subjectId),
			})
		}

		return context.Status(fiber.StatusOK).JSON(&fiber.Map{
			"success": true,
			"content": subject,
			"error":   nil,
		})
	}
}

func getSubjectByNameControllerFactory(
	repository *Repository,
) func(ctx *fiber.Ctx) error {
	return func(context *fiber.Ctx) error {
		name := context.Params("subjectName")

		trimmedName := strings.TrimSpace(name)
		lowerName := strings.ToLower(trimmedName)

		if len(lowerName) == 0 {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   "The subject name is required",
			})
		}

		subject := entities.Subject{Name: lowerName}

		repository.GetItemByName(&subject)

		if subject.SubjectID == 0 {
			return context.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   "Subject not found",
			})
		}

		return context.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"content": subject,
			"error":   nil,
		})
	}
}

// TODO: prevent delete when at least 1 Cornell Note is an Owner
func deleteSubjectByIdControllerFactory(
	repository *Repository,
) func(ctx *fiber.Ctx) error {
	return func(context *fiber.Ctx) error {
		subjectId := context.Params("subjectId")

		parsedId, parsedIdErr := infra.ParseID(subjectId)

		if parsedIdErr != nil {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   parsedIdErr.Error(),
			})
		}

		var subject = entities.Subject{SubjectID: parsedId}

		repository.GetItem(&subject)

		if subject.Name == "" {
			return context.Status(fiber.StatusNotFound).JSON(&fiber.Map{
				"success": false,
				"content": nil,
				"error":   fmt.Sprintf("ID %v not found", subjectId),
			})
		}

		repository.DeleteItem(&subject)

		return context.Status(fiber.StatusAccepted).JSON(&fiber.Map{
			"success": true,
			"content": subject,
			"error":   nil,
		})
	}
}

// TODO: In which case don't allow update name ??
func updateSubjectByIdControllerFactory(
	repository *Repository,
) func(ctx *fiber.Ctx) error {
	return func(context *fiber.Ctx) error {
		subjectId := context.Params("subjectId")

		parsedId, parsedIdErr := infra.ParseID(subjectId)

		if parsedIdErr != nil {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   parsedIdErr.Error(),
			})
		}

		var subject = entities.Subject{SubjectID: parsedId}

		repository.GetItem(&subject)

		if subject.Name == "" {
			return context.Status(fiber.StatusNotFound).JSON(&fiber.Map{
				"success": false,
				"content": nil,
				"error":   fmt.Sprintf("ID %v not found", subjectId),
			})
		}

		proposedSubject := new(entities.Subject)

		if err := context.BodyParser(proposedSubject); err != nil {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   err.Error(),
			})
		}

		proposedSubject.Name = strings.TrimSpace(proposedSubject.Name)
		proposedSubject.Name = strings.ToLower(proposedSubject.Name)

		if proposedSubject.Name == "" {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   "Subject Name is required",
			})
		}

		if subject.Name != proposedSubject.Name {
			subject.Name = proposedSubject.Name
		}

		repository.UpdateItem(&subject)

		return context.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"content": subject,
			"error":   nil,
		})

	}
}
