package topic

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
		topic := new(entities.Topic)

		if err := context.BodyParser(topic); err != nil {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   err.Error(),
			})
		}

		topic.Name = strings.TrimSpace(topic.Name)
		topic.Name = strings.ToLower(topic.Name)

		if topic.Name == "" {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   "Topic name is required",
			})
		}

		repository.GetItemByName(topic)

		if topic.TopicID != 0 {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   "Topic already created",
			})
		}

		if err := repository.CreateItem(topic); err != nil {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   err.Error(),
			})
		}

		return context.Status(fiber.StatusCreated).JSON(fiber.Map{
			"success": true,
			"content": topic,
			"error":   nil,
		})

	}
}

func getAllItemsControllerFactory(
	repository *Repository,
) func(*fiber.Ctx) error {
	return func(context *fiber.Ctx) error {
		var topics []entities.Topic

		if err := repository.GetAllItems(&topics); err != nil {
			return context.Status(fiber.StatusInternalServerError).
				JSON(fiber.Map{
					"success": false,
					"content": nil,
					"error":   err.Error(),
				})
		}

		if len(topics) == 0 {
			return context.Status(fiber.StatusOK).JSON(fiber.Map{
				"success": true,
				"content": []fiber.Map{},
				"error":   nil,
			})
		}

		return context.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"content": topics,
			"error":   nil,
		})
	}
}

func getItemByIdControllerFactory(
	repository *Repository,
) func(ctx *fiber.Ctx) error {
	return func(context *fiber.Ctx) error {
		id := context.Params("topicId")

		parsedId, parsedIdErr := infra.ParseID(id)

		if parsedIdErr != nil {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   parsedIdErr.Error(),
			})
		}

		var topic = entities.Topic{TopicID: parsedId}

		repository.GetItemById(&topic)

		if topic.Name == "" {
			return context.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   fmt.Sprintf("ID %v not found", id),
			})
		}

		return context.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"content": topic,
			"error":   nil,
		})
	}
}

func deleteItemByIdControllerFactory(
	repository *Repository,
) func(ctx *fiber.Ctx) error {
	return func(context *fiber.Ctx) error {
		id := context.Params("topicId")

		parsedId, parsedIdErr := infra.ParseID(id)

		if parsedIdErr != nil {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   parsedIdErr.Error(),
			})
		}

		var topic = entities.Topic{TopicID: parsedId}

		repository.GetItemById(&topic)

		if topic.Name == "" {
			return context.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   fmt.Sprintf("ID %v not found", id),
			})
		}

		repository.DeleteItem(&topic)

		return context.Status(fiber.StatusAccepted).JSON(fiber.Map{
			"success": true,
			"content": topic,
			"error":   nil,
		})
	}
}

func updateItemByIdControllerFactory(
	repository *Repository,
) func(ctx *fiber.Ctx) error {
	return func(context *fiber.Ctx) error {
		id := context.Params("topicId")

		parsedId, parsedIdErr := infra.ParseID(id)

		if parsedIdErr != nil {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   parsedIdErr.Error(),
			})
		}

		var topic = entities.Topic{TopicID: parsedId}

		repository.GetItemById(&topic)

		if topic.Name == "" {
			return context.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   fmt.Sprintf("ID %v not found", id),
			})
		}

		proposedItem := new(entities.Topic)

		if err := context.BodyParser(proposedItem); err != nil {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   err.Error(),
			})
		}

		topic.Name = strings.TrimSpace(topic.Name)
		topic.Name = strings.ToLower(topic.Name)

		if proposedItem.Name == "" {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   "Topic name is required",
			})
		}

		if topic.Name != proposedItem.Name {
			topic.Name = proposedItem.Name
		}

		repository.UpdateItem(&topic)

		return context.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"content": topic,
			"error":   nil,
		})

	}
}
