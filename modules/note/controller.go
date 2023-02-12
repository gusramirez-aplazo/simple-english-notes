package note

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
	requestBody := new(domain.Note)

	if err := context.BodyParser(requestBody); err != nil {
		return infra.CustomResponse(
			context,
			fiber.StatusBadRequest,
			false,
			nil,
			err.Error(),
		)
	}

	requestBody.Content = strings.TrimSpace(requestBody.Content)
	requestBody.Cue = strings.TrimSpace(requestBody.Cue)

	if len(requestBody.Content) == 0 {
		return infra.CustomResponse(
			context,
			fiber.StatusBadRequest,
			false,
			nil,
			"note content is empty",
		)
	}

	newItem, createErr := getRepo().
		CreateOne(
			requestBody.Content,
			requestBody.Cue,
		)

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

func (controller *Controller) getOneById(
	context *fiber.Ctx,
) error {
	id := context.Params("noteId")

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
	id := context.Params("noteId")

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

func (controller *Controller) updateOne(
	context *fiber.Ctx,
) error {
	id := context.Params("noteId")

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

	var proposedItem *domain.Note

	if err := context.BodyParser(&proposedItem); err != nil {
		return infra.CustomResponse(
			context,
			fiber.StatusInternalServerError,
			false,
			nil,
			err.Error(),
		)
	}

	proposedItem.Content = strings.TrimSpace(proposedItem.Content)
	proposedItem.Cue = strings.TrimSpace(proposedItem.Cue)

	if proposedItem.Content == "" {
		emptyErr := errors.New("note content is empty")

		return infra.CustomResponse(
			context,
			fiber.StatusBadRequest,
			false,
			nil,
			emptyErr.Error(),
		)
	}

	item, updateErr := getRepo().
		UpdateOne(
			parsedId,
			proposedItem.Content,
			proposedItem.Cue,
		)

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
