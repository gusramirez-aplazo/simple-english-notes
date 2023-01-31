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

func (controller Controller) GetAllCategoriesControllerFactory(
	clientDB *gorm.DB,
) func(*fiber.Ctx) error {
	return func(context *fiber.Ctx) error {
		var categories []models.Category

		clientDB.Find(&categories)

		var formattedCategories []fiber.Map

		for i := 0; i < len(categories); i++ {
			formattedCategories = append(formattedCategories, fiber.Map{
				"id":          categories[i].ID,
				"name":        categories[i].Name,
				"description": categories[i].Description,
				"createdAt":   categories[i].CreatedAt,
			})
		}

		return context.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"content": formattedCategories,
			"error":   nil,
		})
	}
}

func (controller Controller) CreateCategoryControllerFactory(
	clientDB *gorm.DB,
	validate *validator.Validate,
) func(*fiber.Ctx) error {
	return func(context *fiber.Ctx) error {
		category := new(models.Category)

		if err := context.BodyParser(category); err != nil {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   err.Error(),
			})
		}

		category.Name = strings.TrimSpace(category.Name)
		category.Name = strings.ToLower(category.Name)
		category.Description = strings.TrimSpace(category.Description)

		validationErrors := ValidateStruct(*category, validate)

		if validationErrors != nil {
			return context.Status(fiber.StatusBadRequest).
				JSON(fiber.Map{
					"success": false,
					"content": nil,
					"error":   validationErrors,
				})
		}

		dbCreationResponse := clientDB.Create(&category)

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
				"id":          category.ID,
				"name":        category.Name,
				"description": category.Description,
				"createdAt":   category.CreatedAt,
			},
			"error": nil,
		})
	}
}

func (controller Controller) GetCategoryByIdControllerFactory(
	clientDB *gorm.DB,
) func(ctx *fiber.Ctx) error {
	return func(context *fiber.Ctx) error {
		categoryId := context.Params("categoryId")

		if len(categoryId) == 0 {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   "Category id is empty",
			})
		}

		ui64, parseErr := strconv.ParseUint(categoryId, 10, 64)

		if parseErr != nil {
			return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"success": false,
				"content": nil,
				"error":   fmt.Sprintf("Ensure the ID %v is numeric only", categoryId),
			})
		}

		var category = models.Category{ID: uint(ui64)}

		clientDB.First(&category)

		if category.Name == "" {
			return context.Status(fiber.StatusNotFound).JSON(&fiber.Map{
				"success": false,
				"content": nil,
				"error":   fmt.Sprintf("ID %v not found", categoryId),
			})
		}

		return context.Status(fiber.StatusOK).JSON(&fiber.Map{
			"success": true,
			"content": fiber.Map{
				"id":          category.ID,
				"name":        category.Name,
				"description": category.Description,
				"createdAt":   category.CreatedAt,
			},
			"error": nil,
		})
	}
}

func (controller Controller) DeleteCategoryByIdControllerFactory(
	clientDB *gorm.DB,
) func(ctx *fiber.Ctx) error {
	return func(context *fiber.Ctx) error {
		categoryId := context.Params("categoryId")

		if len(categoryId) == 0 {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   "Category id is empty",
			})
		}

		ui64, parseErr := strconv.ParseUint(categoryId, 10, 64)

		if parseErr != nil {
			return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"success": false,
				"content": nil,
				"error":   fmt.Sprintf("Ensure the ID %v is numeric only", categoryId),
			})
		}

		var category = models.Category{ID: uint(ui64)}

		clientDB.First(&category)

		if category.Name == "" {
			return context.Status(fiber.StatusNotFound).JSON(&fiber.Map{
				"success": false,
				"content": nil,
				"error":   fmt.Sprintf("ID %v not found", categoryId),
			})
		}

		clientDB.Delete(&category)

		return context.Status(fiber.StatusAccepted).JSON(&fiber.Map{
			"success": true,
			"content": fiber.Map{
				"id":          category.ID,
				"name":        category.Name,
				"description": category.Description,
				"deletedAt":   category.DeletedAt,
			},
			"error": nil,
		})
	}
}

func (controller Controller) UpdateCategoryByIdControllerFactory(
	clientDB *gorm.DB,
	validate *validator.Validate,
) func(ctx *fiber.Ctx) error {
	return func(context *fiber.Ctx) error {
		categoryId := context.Params("categoryId")

		if len(categoryId) == 0 {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   "Category id is empty",
			})
		}

		ui64, parseErr := strconv.ParseUint(categoryId, 10, 64)

		if parseErr != nil {
			return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"success": false,
				"content": nil,
				"error":   fmt.Sprintf("Ensure the ID %v is numeric only", categoryId),
			})
		}

		var category = models.Category{ID: uint(ui64)}

		clientDB.First(&category)

		if category.Name == "" {
			return context.Status(fiber.StatusNotFound).JSON(&fiber.Map{
				"success": false,
				"content": nil,
				"error":   fmt.Sprintf("ID %v not found", categoryId),
			})
		}

		proposedCategory := new(models.Category)

		if err := context.BodyParser(proposedCategory); err != nil {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   err.Error(),
			})
		}

		proposedCategory.Name = strings.TrimSpace(proposedCategory.Name)
		proposedCategory.Name = strings.ToLower(proposedCategory.Name)
		proposedCategory.Description = strings.TrimSpace(proposedCategory.Description)

		validationErrors := ValidateStruct(*proposedCategory, validate)

		if validationErrors != nil {
			return context.Status(fiber.StatusBadRequest).
				JSON(fiber.Map{
					"success": false,
					"content": nil,
					"error":   validationErrors,
				})
		}

		if category.Name != proposedCategory.Name {
			category.Name = proposedCategory.Name
		}

		if category.Description != proposedCategory.Description {
			category.Description = proposedCategory.Description
		}

		clientDB.Save(&category)

		return context.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"content": fiber.Map{
				"id":          category.ID,
				"name":        category.Name,
				"description": category.Description,
				"updatedAt":   category.UpdatedAt,
			},
			"error": nil,
		})

	}
}
