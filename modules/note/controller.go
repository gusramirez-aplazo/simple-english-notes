package note

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gusramirez-aplazo/simple-english-notes/modules/shared/entities"
	"github.com/gusramirez-aplazo/simple-english-notes/modules/shared/infra"
	"strings"
)

func creationControllerFactory(
	repository *Repository,
) func(*fiber.Ctx) error {
	return func(context *fiber.Ctx) error {
		note := new(entities.Note)

		if err := context.BodyParser(note); err != nil {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   err.Error(),
			})
		}

		note.Content = strings.TrimSpace(note.Content)
		note.Cue = strings.TrimSpace(note.Cue)

		if note.Content == "" {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   "Content is required",
			})
		}

		if err := repository.CreateItem(note); err != nil {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   err.Error(),
			})
		}

		return context.Status(fiber.StatusCreated).JSON(fiber.Map{
			"success": true,
			"content": note,
			"error":   nil,
		})

	}
}

func getAllItemsControllerFactory(
	repository *Repository,
) func(*fiber.Ctx) error {
	return func(context *fiber.Ctx) error {
		var notes []entities.Note

		if err := repository.GetAllItems(&notes); err != nil {
			return context.Status(fiber.StatusInternalServerError).
				JSON(fiber.Map{
					"success": false,
					"content": nil,
					"error":   err.Error(),
				})
		}

		if len(notes) == 0 {
			return context.Status(fiber.StatusOK).JSON(fiber.Map{
				"success": true,
				"content": []fiber.Map{},
				"error":   nil,
			})
		}

		return context.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"content": notes,
			"error":   nil,
		})
	}
}

func getItemByIdControllerFactory(
	repository *Repository,
) func(ctx *fiber.Ctx) error {
	return func(context *fiber.Ctx) error {
		id := context.Params("noteId")

		parsedId, parsedIdErr := infra.ParseID(id)

		if parsedIdErr != nil {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   parsedIdErr.Error(),
			})
		}

		var note = entities.Note{NoteID: parsedId}

		repository.GetItem(&note)

		if note.Content == "" {
			return context.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   fmt.Sprintf("ID %v not found", id),
			})
		}

		return context.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"content": note,
			"error":   nil,
		})
	}
}

func deleteItemByIdControllerFactory(
	repository *Repository,
) func(ctx *fiber.Ctx) error {
	return func(context *fiber.Ctx) error {
		id := context.Params("noteId")

		parsedId, parsedIdErr := infra.ParseID(id)

		if parsedIdErr != nil {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   parsedIdErr.Error(),
			})
		}

		var note = entities.Note{NoteID: parsedId}

		repository.GetItem(&note)

		if note.Content == "" {
			return context.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   fmt.Sprintf("ID %v not found", id),
			})
		}

		repository.DeleteItem(&note)

		return context.Status(fiber.StatusAccepted).JSON(fiber.Map{
			"success": true,
			"content": note,
			"error":   nil,
		})
	}
}

func updateItemByIdControllerFactory(
	repository *Repository,
) func(ctx *fiber.Ctx) error {
	return func(context *fiber.Ctx) error {
		id := context.Params("noteId")

		parsedId, parsedIdErr := infra.ParseID(id)

		if parsedIdErr != nil {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   parsedIdErr.Error(),
			})
		}

		var note = entities.Note{NoteID: parsedId}

		repository.GetItem(&note)

		if note.Content == "" {
			return context.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   fmt.Sprintf("ID %v not found", id),
			})
		}

		proposedItem := new(entities.Note)

		if err := context.BodyParser(proposedItem); err != nil {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   err.Error(),
			})
		}

		proposedItem.Content = strings.TrimSpace(proposedItem.Content)
		proposedItem.Cue = strings.TrimSpace(proposedItem.Cue)

		if proposedItem.Content == "" {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   "Content is required",
			})
		}

		if note.Content != proposedItem.Content {
			note.Content = proposedItem.Content
		}

		if note.Cue != proposedItem.Cue {
			note.Cue = proposedItem.Cue
		}

		repository.UpdateItem(&note)

		return context.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"content": note,
			"error":   nil,
		})

	}
}
