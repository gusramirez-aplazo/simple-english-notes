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
	items, getAllErr := getRepo().
		GetAllItems()

	if getAllErr != nil {
		return infra.CustomResponse(
			context,
			fiber.StatusInternalServerError,
			false,
			nil,
			getAllErr.Error(),
		)
	}

	if len(items) == 0 {
		return infra.CustomResponse(
			context,
			fiber.StatusOK,
			true,
			[]fiber.Map{},
			"",
		)
	}

	return infra.CustomResponse(
		context,
		fiber.StatusOK,
		true,
		items,
		"",
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
			err.Error(),
		)
	}

	requestBody.Name = strings.TrimSpace(requestBody.Name)
	requestBody.Name = strings.ToLower(requestBody.Name)

	if len(requestBody.Name) == 0 {
		emptyNameErr := errors.New("category name is empty")
		return infra.CustomResponse(
			context,
			fiber.StatusBadRequest,
			false,
			nil,
			emptyNameErr.Error(),
		)
	}

	_, findErr := getRepo().
		GetItemByUniqueParam(requestBody.Name)

	isItemFounded := findErr == nil

	if isItemFounded {
		return infra.CustomResponse(
			context,
			fiber.StatusBadRequest,
			true,
			nil,
			"the category is already created",
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
			createErr.Error(),
		)
	}

	return infra.CustomResponse(
		context,
		fiber.StatusCreated,
		true,
		newItem,
		"",
	)
}

func (controller *Controller) getOneByUniqueParam(
	context *fiber.Ctx,
) error {
	type queryParams struct {
		Name string `query:"name"`
	}

	q := new(queryParams)

	if err := context.QueryParser(q); err != nil {
		return infra.CustomResponse(
			context,
			fiber.StatusInternalServerError,
			false,
			nil,
			err.Error(),
		)
	}

	q.Name = strings.TrimSpace(q.Name)
	q.Name = strings.ToLower(q.Name)

	if len(q.Name) == 0 {
		emptyNameErr := errors.New("name is empty")
		return infra.CustomResponse(
			context,
			fiber.StatusBadRequest,
			false,
			nil,
			emptyNameErr.Error(),
		)
	}

	item, findErr := getRepo().
		GetItemByUniqueParam(q.Name)

	itemNotFound := findErr != nil

	if itemNotFound {
		return infra.CustomResponse(
			context,
			fiber.StatusNotFound,
			false,
			nil,
			findErr.Error(),
		)
	}

	return infra.CustomResponse(
		context,
		fiber.StatusOK,
		true,
		item,
		"",
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
			parsedIdErr.Error(),
		)
	}

	item, findErr := getRepo().
		GetItemById(parsedId)

	if findErr != nil {
		return infra.CustomResponse(
			context,
			fiber.StatusNotFound,
			false,
			nil,
			findErr.Error(),
		)
	}

	if item.ID == 0 {
		notFoundErr := errors.New("item not found")
		return infra.CustomResponse(
			context,
			fiber.StatusNotFound,
			false,
			nil,
			notFoundErr.Error(),
		)
	}

	return infra.CustomResponse(
		context,
		fiber.StatusOK,
		true,
		item,
		"",
	)
}

func (controller *Controller) deleteOne(
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
			parsedIdErr.Error(),
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
			deleteErr.Error(),
		)
	}

	if deletedItem.ID == 0 {
		notFoundErr := errors.New(
			fmt.Sprintf(
				"item with ID <%v> not found",
				parsedId,
			),
		)
		return infra.CustomResponse(
			context,
			fiber.StatusNotFound,
			false,
			nil,
			notFoundErr.Error(),
		)
	}

	return infra.CustomResponse(
		context,
		fiber.StatusAccepted,
		true,
		deletedItem,
		"",
	)
}

func (controller *Controller) updateOne(context *fiber.Ctx) error {
	id := context.Params("categoryId")

	parsedId, parsedIdErr := infra.ParseID(id)

	if parsedIdErr != nil {
		return infra.CustomResponse(
			context,
			fiber.StatusBadRequest,
			false,
			nil,
			parsedIdErr.Error(),
		)
	}

	var proposedCategory *domain.Category

	if err := context.BodyParser(&proposedCategory); err != nil {
		return infra.CustomResponse(
			context,
			fiber.StatusInternalServerError,
			false,
			nil,
			err.Error(),
		)
	}

	proposedCategory.Name = strings.TrimSpace(proposedCategory.Name)
	proposedCategory.Name = strings.ToLower(proposedCategory.Name)

	if proposedCategory.Name == "" {
		emptyErr := errors.New("name is empty")

		return infra.CustomResponse(
			context,
			fiber.StatusBadRequest,
			false,
			nil,
			emptyErr.Error(),
		)
	}

	item, updateErr := getRepo().
		UpdateOne(parsedId, proposedCategory.Name)

	if updateErr != nil {
		return infra.CustomResponse(
			context,
			fiber.StatusBadRequest,
			false,
			nil,
			updateErr.Error(),
		)
	}

	return infra.CustomResponse(
		context,
		fiber.StatusAccepted,
		true,
		item,
		"",
	)
}
