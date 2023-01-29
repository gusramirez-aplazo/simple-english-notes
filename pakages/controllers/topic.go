package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gusramirez-aplazo/simple-english-notes/pakages/models"
	"gorm.io/gorm"
	"strings"
)

// TODO: testing
func (controller Controller) CreateTopicControllerFactory(clientDB *gorm.DB, validate *validator.Validate) func(*fiber.Ctx) error {
	return func(context *fiber.Ctx) error {
		context.Accepts("application/json")

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
