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

		subject.Description = strings.TrimSpace(subject.Description)

		if err := repository.GetItemOrCreate(subject); err != nil {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   err.Error(),
			})
		}

		return context.Status(fiber.StatusCreated).JSON(fiber.Map{
			"success": true,
			"content": fiber.Map{
				"subjectId":   subject.SubjectID,
				"name":        subject.Name,
				"description": subject.Description,
				"createdAt":   subject.CreatedAt,
			},
			"error": nil,
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

		var formattedSubjects []fiber.Map

		for i := 0; i < len(subjects); i++ {
			formattedSubjects = append(formattedSubjects, fiber.Map{
				"subjectId":   subjects[i].SubjectID,
				"name":        subjects[i].Name,
				"description": subjects[i].Description,
				"createdAt":   subjects[i].CreatedAt,
			})
		}

		if len(formattedSubjects) == 0 {
			return context.Status(fiber.StatusOK).JSON(fiber.Map{
				"success": true,
				"content": []fiber.Map{},
				"error":   nil,
			})
		}

		return context.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"content": formattedSubjects,
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
			"content": fiber.Map{
				"subjectId":   subject.SubjectID,
				"name":        subject.Name,
				"description": subject.Description,
				"createdAt":   subject.CreatedAt,
			},
			"error": nil,
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
			"content": fiber.Map{
				"subjectId":   subject.SubjectID,
				"name":        subject.Name,
				"description": subject.Description,
				"deletedAt":   subject.DeletedAt,
			},
			"error": nil,
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
		proposedSubject.Description = strings.TrimSpace(proposedSubject.Description)

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

		if proposedSubject.Description != "" &&
			subject.Description != proposedSubject.Description {
			subject.Description = proposedSubject.Description
		}

		repository.UpdateItem(&subject)

		return context.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"content": fiber.Map{
				"id":          subject.SubjectID,
				"name":        subject.Name,
				"description": subject.Description,
				"updatedAt":   subject.UpdatedAt,
			},
			"error": nil,
		})

	}
}
