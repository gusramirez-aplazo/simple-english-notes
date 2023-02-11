package category

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gusramirez-aplazo/simple-english-notes/modules/shared/domain"
	"github.com/gusramirez-aplazo/simple-english-notes/modules/shared/infra"
	"strings"
)

type Controller struct {
	repository *Repository
}

var singleController *Controller

func GetController(
	repo *Repository,
) *Controller {
	if singleController == nil {
		singleController = &Controller{
			repository: repo,
		}
	}
	return singleController
}

func getRepo() *Repository {
	return singleController.repository
}

func (controller *Controller) getAll(
	context *fiber.Ctx,
) error {
	items, err := getRepo().
		GetAllItems()

	if err != nil {
		return infra.CustomResponse(
			context,
			fiber.StatusInternalServerError,
			false,
			nil,
			err,
		)
	}

	if len(items) == 0 {
		return infra.CustomResponse(
			context,
			fiber.StatusOK,
			true,
			[]fiber.Map{},
			nil,
		)
	}

	return infra.CustomResponse(
		context,
		fiber.StatusOK,
		true,
		items,
		nil,
	)
}

func (controller *Controller) createOne(
	context *fiber.Ctx,
) error {
	requestBody := new(domain.Category)

	if err := context.BodyParser(requestBody); err != nil {
		return infra.CustomResponse(
			context,
			fiber.StatusBadRequest,
			false,
			nil,
			err,
		)
	}

	requestBody.Name = strings.TrimSpace(requestBody.Name)
	requestBody.Name = strings.ToLower(requestBody.Name)

	if len(requestBody.Name) == 0 {
		return infra.CustomResponse(
			context,
			fiber.StatusBadRequest,
			false,
			nil,
			errors.New("category name is empty"),
		)
	}

	item, findErr := getRepo().
		GetItemByUniqueParam(requestBody.Name)

	isItemFounded := item.ID != 0

	if findErr != nil && isItemFounded {
		return infra.CustomResponse(
			context,
			fiber.StatusOK,
			true,
			item,
			nil,
		)
	}

	if findErr != nil && !isItemFounded {
		return infra.CustomResponse(
			context,
			fiber.StatusInternalServerError,
			false,
			nil,
			findErr,
		)
	}

	newItem, createErr := getRepo().
		CreateOne(requestBody.Name)

	if createErr != nil {
		return infra.CustomResponse(
			context,
			fiber.StatusInternalServerError,
			false,
			nil,
			createErr,
		)
	}

	return infra.CustomResponse(
		context,
		fiber.StatusCreated,
		true,
		newItem,
		nil,
	)
}

func (controller *Controller) getOneByUniqueParam(
	context *fiber.Ctx,
) error {
	name := context.Params("categoryName")

	trimmedName := strings.TrimSpace(name)
	lowerName := strings.ToLower(trimmedName)

	if len(lowerName) == 0 {
		return infra.CustomResponse(
			context,
			fiber.StatusBadRequest,
			false,
			nil,
			errors.New("name is empty"),
		)
	}

	item, findErr := getRepo().
		GetItemByUniqueParam(name)

	isItemFounded := item.ID != 0

	if findErr != nil && !isItemFounded {
		return infra.CustomResponse(
			context,
			fiber.StatusInternalServerError,
			false,
			nil,
			findErr,
		)
	}

	return infra.CustomResponse(
		context,
		fiber.StatusOK,
		true,
		item,
		nil,
	)
}

func (controller *Controller) getOneById(
	context *fiber.Ctx,
) error {
	id := context.Params("categoryId")

	parsedId, parsedIdErr := infra.ParseID(id)

	if parsedIdErr != nil {
		return infra.CustomResponse(
			context,
			fiber.StatusBadRequest,
			false,
			nil,
			parsedIdErr,
		)
	}

	item, findErr := getRepo().
		GetItemById(parsedId)

	if findErr != nil {
		return infra.CustomResponse(
			context,
			fiber.StatusInternalServerError,
			false,
			nil,
			findErr,
		)
	}

	if item.ID == 0 {
		return infra.CustomResponse(
			context,
			fiber.StatusNotFound,
			false,
			nil,
			errors.New("item not found"))
	}

	return infra.CustomResponse(
		context,
		fiber.StatusOK,
		true,
		item,
		nil,
	)
}

func (controller *Controller) DeleteOne(
	context *fiber.Ctx,
) error {
	id := context.Params("categoryId")

	parsedId, parsedIdErr := infra.ParseID(id)

	if parsedIdErr != nil {
		return infra.CustomResponse(
			context,
			fiber.StatusBadRequest,
			false,
			nil,
			parsedIdErr,
		)
	}

	deletedItem, deleteErr := getRepo().
		DeleteOne(parsedId)

	if deleteErr != nil {
		return infra.CustomResponse(
			context,
			fiber.StatusBadRequest,
			false,
			nil,
			deleteErr,
		)
	}

	if deletedItem.ID == 0 {
		return infra.CustomResponse(
			context,
			fiber.StatusNotFound,
			false,
			nil,
			errors.New(
				fmt.Sprintf(
					"item with ID <%v> not found",
					parsedId,
				),
			),
		)
	}

	return infra.CustomResponse(
		context,
		fiber.StatusAccepted,
		true,
		deletedItem,
		nil,
	)
}

func (controller *Controller) updateOne(context *fiber.Ctx) error {
	id := context.Params("categoryId")

	parsedId, parsedIdErr := infra.ParseID(id)

	if parsedIdErr != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"content": nil,
			"error":   parsedIdErr.Error(),
		})
	}

	var proposedCategory *domain.Category

	if err := context.BodyParser(&proposedCategory); err != nil {
		return infra.CustomResponse(
			context,
			fiber.StatusInternalServerError,
			false,
			nil,
			err,
		)
	}

	proposedCategory.Name = strings.TrimSpace(proposedCategory.Name)
	proposedCategory.Name = strings.ToLower(proposedCategory.Name)

	if proposedCategory.Name == "" {
		return infra.CustomResponse(
			context,
			fiber.StatusBadRequest,
			false,
			nil,
			errors.New("name is empty"))
	}

	item, updateErr := getRepo().
		UpdateOne(parsedId, proposedCategory.Name)

	if updateErr != nil {
		return infra.CustomResponse(
			context,
			fiber.StatusNotFound,
			false,
			nil,
			errors.New("item not found"),
		)
	}

	return infra.CustomResponse(
		context,
		fiber.StatusAccepted,
		true,
		item,
		nil,
	)
}
