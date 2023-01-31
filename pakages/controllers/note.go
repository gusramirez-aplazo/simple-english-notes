package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gusramirez-aplazo/simple-english-notes/pakages/models"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

func (controller Controller) GetAllNotesControllerFactory(
	clientDB *gorm.DB,
) func(*fiber.Ctx) error {
	return func(context *fiber.Ctx) error {
		var notes []models.Note

		clientDB.Find(&notes)

		var formattedNotes []fiber.Map

		for i := 0; i < len(notes); i++ {
			formattedNotes = append(formattedNotes, fiber.Map{
				"id":        notes[i].ID,
				"content":   notes[i].Content,
				"createdAt": notes[i].CreatedAt,
			})
		}

		return context.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"content": formattedNotes,
			"error":   nil,
		})
	}
}

func (controller Controller) CreateNoteControllerFactory(
	clientDB *gorm.DB,
) func(*fiber.Ctx) error {
	return func(context *fiber.Ctx) error {
		note := new(models.Note)

		if err := context.BodyParser(note); err != nil {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   err.Error(),
			})
		}

		note.Content = strings.TrimSpace(note.Content)
		note.Content = strings.ToLower(note.Content)

		if note.Content == "" {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   "There is no content to create",
			})
		}

		dbCreationResponse := clientDB.Create(&note)

		if dbCreationResponse.Error != nil {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   dbCreationResponse.Error.Error(),
			})
		}

		return context.Status(fiber.StatusCreated).JSON(fiber.Map{
			"success": true,
			"content": fiber.Map{
				"id":        note.ID,
				"content":   note.Content,
				"createdAt": note.CreatedAt,
			},
			"error": nil,
		})
	}
}

func (controller Controller) GetNoteByIdControllerFactory(
	clientDB *gorm.DB,
) func(ctx *fiber.Ctx) error {
	return func(context *fiber.Ctx) error {
		noteId := context.Params("noteId")

		if len(noteId) == 0 {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   "Note ID is empty",
			})
		}

		ui64, parseErr := strconv.ParseUint(noteId, 10, 64)

		if parseErr != nil {
			return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"success": false,
				"content": nil,
				"error":   fmt.Sprintf("Ensure the ID %v is numeric only", noteId),
			})
		}

		var note = models.Note{ID: uint(ui64)}

		clientDB.First(&note)

		if note.Content == "" {
			return context.Status(fiber.StatusNotFound).JSON(&fiber.Map{
				"success": false,
				"content": nil,
				"error":   fmt.Sprintf("ID %v not found", noteId),
			})
		}

		return context.Status(fiber.StatusOK).JSON(&fiber.Map{
			"success": true,
			"content": fiber.Map{
				"id":        note.ID,
				"content":   note.Content,
				"createdAt": note.CreatedAt,
			},
			"error": nil,
		})
	}
}

func (controller Controller) DeleteNoteByIdControllerFactory(
	clientDB *gorm.DB,
) func(ctx *fiber.Ctx) error {
	return func(context *fiber.Ctx) error {
		noteId := context.Params("noteId")

		if len(noteId) == 0 {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   "Note ID is empty",
			})
		}

		ui64, parseErr := strconv.ParseUint(noteId, 10, 64)

		if parseErr != nil {
			return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"success": false,
				"content": nil,
				"error":   fmt.Sprintf("Ensure the ID %v is numeric only", noteId),
			})
		}

		var note = models.Note{ID: uint(ui64)}

		clientDB.First(&note)

		if note.Content == "" {
			return context.Status(fiber.StatusNotFound).JSON(&fiber.Map{
				"success": false,
				"content": nil,
				"error":   fmt.Sprintf("ID %v not found", noteId),
			})
		}

		clientDB.Delete(&note)

		return context.Status(fiber.StatusAccepted).JSON(&fiber.Map{
			"success": true,
			"content": fiber.Map{
				"id":        note.ID,
				"content":   note.Content,
				"deletedAt": note.DeletedAt,
			},
			"error": nil,
		})
	}
}

func (controller Controller) UpdateNoteByIdControllerFactory(
	clientDB *gorm.DB,
) func(ctx *fiber.Ctx) error {
	return func(context *fiber.Ctx) error {
		noteId := context.Params("noteId")

		if len(noteId) == 0 {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   "Note ID is empty",
			})
		}

		ui64, parseErr := strconv.ParseUint(noteId, 10, 64)

		if parseErr != nil {
			return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"success": false,
				"content": nil,
				"error":   fmt.Sprintf("Ensure the ID %v is numeric only", noteId),
			})
		}

		var note = models.Note{ID: uint(ui64)}

		clientDB.First(&note)

		if note.Content == "" {
			return context.Status(fiber.StatusNotFound).JSON(&fiber.Map{
				"success": false,
				"content": nil,
				"error":   fmt.Sprintf("ID %v not found", noteId),
			})
		}

		proposedNote := new(models.Note)

		if err := context.BodyParser(proposedNote); err != nil {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   err.Error(),
			})
		}

		proposedNote.Content = strings.TrimSpace(proposedNote.Content)
		proposedNote.Content = strings.ToLower(proposedNote.Content)

		if proposedNote.Content == "" {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   "The note has no content!!",
			})
		}

		if note.Content != proposedNote.Content {
			note.Content = proposedNote.Content
		}

		clientDB.Save(&note)

		return context.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"content": fiber.Map{
				"id":        note.ID,
				"content":   note.Content,
				"updatedAt": note.UpdatedAt,
			},
			"error": nil,
		})

	}
}
