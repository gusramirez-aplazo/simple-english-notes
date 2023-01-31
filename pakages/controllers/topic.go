package controllers

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gusramirez-aplazo/simple-english-notes/pakages/models"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

// TODO: testing
func (controller Controller) CreateTopicControllerFactory(
	clientDB *gorm.DB,
	validate *validator.Validate,
) func(*fiber.Ctx) error {
	return func(context *fiber.Ctx) error {
		topic := new(models.Topic)

		if err := context.BodyParser(topic); err != nil {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   err.Error(),
			})
		}

		topic.Name = strings.TrimSpace(topic.Name)
		topic.Name = strings.ToLower(topic.Name)
		topic.Description = strings.TrimSpace(topic.Description)

		validationErrors := ValidateStruct(*topic, validate)

		if validationErrors != nil {
			return context.Status(fiber.StatusBadRequest).
				JSON(fiber.Map{
					"success": false,
					"content": nil,
					"error":   validationErrors,
				})
		}

		dbCreationResponse := clientDB.Create(&topic)

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
				"id":          topic.ID,
				"name":        topic.Name,
				"description": topic.Description,
				"createdAt":   topic.CreatedAt,
			},
			"error": nil,
		})
	}
}

func (controller Controller) GetTopicsControllerFactory(
	clientDB *gorm.DB,
) func(*fiber.Ctx) error {
	return func(context *fiber.Ctx) error {
		var topics []models.Topic

		clientDB.Find(&topics)

		var formattedTopics []fiber.Map

		for i := 0; i < len(topics); i++ {
			formattedTopics = append(formattedTopics, fiber.Map{
				"id":          topics[i].ID,
				"name":        topics[i].Name,
				"description": topics[i].Description,
				"createdAt":   topics[i].CreatedAt,
			})
		}

		return context.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"content": formattedTopics,
			"error":   nil,
		})
	}
}

func (controller Controller) GetTopicByIdControllerFactory(
	clientDB *gorm.DB,
) func(ctx *fiber.Ctx) error {
	return func(context *fiber.Ctx) error {
		topicId := context.Params("topicId")

		if len(topicId) == 0 {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   "Topic id is empty",
			})
		}

		ui64, parseErr := strconv.ParseUint(topicId, 10, 64)

		if parseErr != nil {
			return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"success": false,
				"content": nil,
				"error":   fmt.Sprintf("Ensure the ID %v is numeric only", topicId),
			})
		}

		var topic = models.Topic{ID: uint(ui64)}

		clientDB.First(&topic)

		if topic.Name == "" {
			return context.Status(fiber.StatusNotFound).JSON(&fiber.Map{
				"success": false,
				"content": nil,
				"error":   fmt.Sprintf("ID %v not found", topicId),
			})
		}

		return context.Status(fiber.StatusOK).JSON(&fiber.Map{
			"success": true,
			"content": fiber.Map{
				"id":          topic.ID,
				"name":        topic.Name,
				"description": topic.Description,
				"createdAt":   topic.CreatedAt,
			},
			"error": nil,
		})
	}
}

func (controller Controller) DeleteTopicByIdControllerFactory(
	clientDB *gorm.DB,
) func(ctx *fiber.Ctx) error {
	return func(context *fiber.Ctx) error {
		topicId := context.Params("topicId")

		if len(topicId) == 0 {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   "Topic id is empty",
			})
		}

		ui64, parseErr := strconv.ParseUint(topicId, 10, 64)

		if parseErr != nil {
			return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"success": false,
				"content": nil,
				"error":   fmt.Sprintf("Ensure the ID %v is numeric only", topicId),
			})
		}

		var topic = models.Topic{ID: uint(ui64)}

		clientDB.First(&topic)

		if topic.Name == "" {
			return context.Status(fiber.StatusNotFound).JSON(&fiber.Map{
				"success": false,
				"content": nil,
				"error":   fmt.Sprintf("ID %v not found", topicId),
			})
		}

		clientDB.Delete(&topic)

		return context.Status(fiber.StatusAccepted).JSON(&fiber.Map{
			"success": true,
			"content": fiber.Map{
				"id":          topic.ID,
				"name":        topic.Name,
				"description": topic.Description,
				"deletedAt":   topic.DeletedAt,
			},
			"error": nil,
		})
	}
}

func (controller Controller) UpdateTopicByIdControllerFactory(
	clientDB *gorm.DB,
	validate *validator.Validate,
) func(ctx *fiber.Ctx) error {
	return func(context *fiber.Ctx) error {
		topicId := context.Params("topicId")

		if len(topicId) == 0 {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   "Topic id is empty",
			})
		}

		ui64, parseErr := strconv.ParseUint(topicId, 10, 64)

		if parseErr != nil {
			return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"success": false,
				"content": nil,
				"error":   fmt.Sprintf("Ensure the ID %v is numeric only", topicId),
			})
		}

		var topic = models.Topic{ID: uint(ui64)}

		clientDB.First(&topic)

		if topic.Name == "" {
			return context.Status(fiber.StatusNotFound).JSON(&fiber.Map{
				"success": false,
				"content": nil,
				"error":   fmt.Sprintf("ID %v not found", topicId),
			})
		}

		proposedTopic := new(models.Topic)

		if err := context.BodyParser(proposedTopic); err != nil {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   err.Error(),
			})
		}

		proposedTopic.Name = strings.TrimSpace(proposedTopic.Name)
		proposedTopic.Name = strings.ToLower(proposedTopic.Name)
		proposedTopic.Description = strings.TrimSpace(proposedTopic.Description)

		validationErrors := ValidateStruct(*proposedTopic, validate)

		if validationErrors != nil {
			return context.Status(fiber.StatusBadRequest).
				JSON(fiber.Map{
					"success": false,
					"content": nil,
					"error":   validationErrors,
				})
		}

		if topic.Name != proposedTopic.Name {
			topic.Name = proposedTopic.Name
		}

		if proposedTopic.Description != "" &&
			topic.Description != proposedTopic.Description {
			topic.Description = proposedTopic.Description
		}

		clientDB.Save(&topic)

		return context.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"content": fiber.Map{
				"id":          topic.ID,
				"name":        topic.Name,
				"description": topic.Description,
				"updatedAt":   topic.UpdatedAt,
			},
			"error": nil,
		})

	}
}
