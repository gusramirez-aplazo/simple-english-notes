package category

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gusramirez-aplazo/simple-english-notes/modules/shared/entities"
	"github.com/gusramirez-aplazo/simple-english-notes/modules/shared/infra"
	"strings"
)

func getAllCategoriesControllerFactory(
	repository *Repository,
) func(*fiber.Ctx) error {
	return func(context *fiber.Ctx) error {
		var categories []entities.Category

		if err := repository.GetAllItems(&categories); err != nil {
			return context.Status(fiber.StatusInternalServerError).
				JSON(fiber.Map{
					"success": false,
					"content": nil,
					"error":   err.Error(),
				})
		}

		if len(categories) == 0 {
			return context.Status(fiber.StatusOK).JSON(fiber.Map{
				"success": true,
				"content": []fiber.Map{},
				"error":   nil,
			})
		}

		return context.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"content": categories,
			"error":   nil,
		})
	}
}

func createCategoryControllerFactory(
	repository *Repository,
) func(*fiber.Ctx) error {
	return func(context *fiber.Ctx) error {
		category := new(entities.Category)

		if err := context.BodyParser(category); err != nil {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   err.Error(),
			})
		}

		category.Name = strings.TrimSpace(category.Name)
		category.Name = strings.ToLower(category.Name)

		if err := repository.GetItemOrCreate(category); err != nil {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   err.Error(),
			})
		}

		return context.Status(fiber.StatusCreated).JSON(fiber.Map{
			"success": true,
			"content": category,
			"error":   nil,
		})
	}
}

func getCategoryByNameControllerFactory(
	repository *Repository,
) func(ctx *fiber.Ctx) error {
	return func(context *fiber.Ctx) error {
		name := context.Params("categoryName")

		trimmedName := strings.TrimSpace(name)
		lowerName := strings.ToLower(trimmedName)

		if len(lowerName) == 0 {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   "The category name is required",
			})
		}

		category := entities.Category{Name: lowerName}

		repository.GetItemByName(&category)

		if category.CategoryID == 0 {
			return context.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   "Category not found",
			})
		}

		return context.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"content": category,
			"error":   nil,
		})
	}
}

func getCategoryByIdControllerFactory(
	repository *Repository,
) func(ctx *fiber.Ctx) error {
	return func(context *fiber.Ctx) error {
		id := context.Params("categoryId")

		parsedId, parsedIdErr := infra.ParseID(id)

		if parsedIdErr != nil {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   parsedIdErr.Error(),
			})
		}

		var category = entities.Category{CategoryID: parsedId}

		repository.GetItem(&category)

		if category.Name == "" {
			return context.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   fmt.Sprintf("ID %v not found", id),
			})
		}

		return context.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"content": category,
			"error":   nil,
		})
	}
}

func deleteCategoryByIdControllerFactory(
	repository *Repository,
) func(ctx *fiber.Ctx) error {
	return func(context *fiber.Ctx) error {
		id := context.Params("categoryId")

		parsedId, parsedIdErr := infra.ParseID(id)

		if parsedIdErr != nil {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   parsedIdErr.Error(),
			})
		}

		var category = entities.Category{CategoryID: parsedId}

		repository.GetItem(&category)

		if category.Name == "" {
			return context.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   fmt.Sprintf("ID %v not found", id),
			})
		}

		repository.DeleteItem(&category)

		return context.Status(fiber.StatusAccepted).JSON(fiber.Map{
			"success": true,
			"content": category,
			"error":   nil,
		})
	}
}

func updateCategoryByIdControllerFactory(
	repository *Repository,
) func(ctx *fiber.Ctx) error {
	return func(context *fiber.Ctx) error {
		id := context.Params("categoryId")

		parsedId, parsedIdErr := infra.ParseID(id)

		if parsedIdErr != nil {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   parsedIdErr.Error(),
			})
		}

		var category = entities.Category{CategoryID: parsedId}

		repository.GetItem(&category)

		if category.Name == "" {
			return context.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   fmt.Sprintf("ID %v not found", id),
			})
		}

		proposedCategory := new(entities.Category)

		if err := context.BodyParser(proposedCategory); err != nil {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   err.Error(),
			})
		}

		proposedCategory.Name = strings.TrimSpace(proposedCategory.Name)
		proposedCategory.Name = strings.ToLower(proposedCategory.Name)

		if proposedCategory.Name == "" {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"content": nil,
				"error":   "Category Name is required",
			})
		}

		if category.Name != proposedCategory.Name {
			category.Name = proposedCategory.Name
		}

		repository.UpdateItem(&category)

		return context.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"content": category,
			"error":   nil,
		})

	}
}
