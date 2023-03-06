package cornellNote

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gusramirez-aplazo/simple-english-notes/modules/note"
	"github.com/gusramirez-aplazo/simple-english-notes/modules/shared/domain"
	"github.com/gusramirez-aplazo/simple-english-notes/modules/shared/infra"
	"github.com/gusramirez-aplazo/simple-english-notes/modules/subject"
	"github.com/gusramirez-aplazo/simple-english-notes/modules/topic"
	"strings"
)

type Controller struct {
	repository  *Repository
	topicRepo   *topic.Repository
	subjectRepo *subject.Repository
	noteRepo    *note.Repository
}

var singleController *Controller

func GetController(
	repo *Repository,
	topicRepo *topic.Repository,
	subjectRepo *subject.Repository,
	noteRepo *note.Repository,
) *Controller {
	if singleController == nil {
		singleController = &Controller{
			repository:  repo,
			topicRepo:   topicRepo,
			subjectRepo: subjectRepo,
			noteRepo:    noteRepo,
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

func (controller *Controller) getOneById(
	context *fiber.Ctx,
) error {
	id := context.Params("id")

	parseId, parseIdErr := infra.ParseID(id)

	if parseIdErr != nil {
		return infra.CustomResponse(
			context,
			fiber.StatusBadRequest,
			false,
			nil,
			parseIdErr.Error(),
		)
	}

	item, findErr := getRepo().
		GetItemById(parseId)

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
		"")
}

func (controller *Controller) createOne(
	context *fiber.Ctx,
) error {
	requestBody := new(domain.CornellNote)

	if err := context.BodyParser(requestBody); err != nil {
		return infra.CustomResponse(
			context,
			fiber.StatusBadRequest,
			false,
			nil,
			err.Error(),
		)
	}

	requestBody.Topic.Name = strings.TrimSpace(requestBody.Topic.Name)
	requestBody.Topic.Name = strings.ToLower(requestBody.Topic.Name)

	if requestBody.Topic.Name == "" {
		return infra.CustomResponse(
			context,
			fiber.StatusBadRequest,
			false,
			nil,
			"Topic is required",
		)
	}

	_, findErr := controller.topicRepo.GetItemByUniqueParam(requestBody.Topic.Name)

	hasExistingItem := findErr == nil

	if hasExistingItem {
		return infra.CustomResponse(
			context,
			fiber.StatusBadRequest,
			false,
			nil,
			"topic is already taken, try with another one")
	}

	if len(requestBody.Subjects) == 0 {
		return infra.CustomResponse(
			context,
			fiber.StatusBadRequest,
			false,
			nil,
			"Add at least 1 subject",
		)
	}

	if len(requestBody.Notes) == 0 {
		return infra.CustomResponse(
			context,
			fiber.StatusBadRequest,
			false,
			nil,
			"Add at least 1 note",
		)
	}

	for i := 0; i < len(requestBody.Subjects); i++ {
		if requestBody.Subjects[i].ID == 0 {
			return infra.CustomResponse(
				context,
				fiber.StatusBadRequest,
				false,
				nil,
				fmt.Sprintf(
					"Subject %v does not have ID",
					requestBody.Subjects[i].Name,
				),
			)
		}

		requestBody.Subjects[i].Name = strings.TrimSpace(
			requestBody.Subjects[i].Name,
		)

		requestBody.Subjects[i].Name = strings.ToLower(
			requestBody.Subjects[i].Name,
		)

		if len(requestBody.Subjects[i].Name) == 0 {
			return infra.CustomResponse(
				context,
				fiber.StatusBadRequest,
				false,
				nil,
				fmt.Sprintf(
					"Subject in position %v has an empty Name",
					i+1,
				),
			)
		}

		item, findErr := controller.subjectRepo.
			GetItemById(requestBody.Subjects[i].ID)

		if findErr != nil {
			notFoundErr := errors.New(
				fmt.Sprintf(
					"subject at position %v not found",
					i+1,
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
		requestBody.Subjects[i] = item
	}

	for i := 0; i < len(requestBody.Notes); i++ {

		requestBody.Notes[i].Content = strings.TrimSpace(
			requestBody.Notes[i].Content,
		)
		requestBody.Notes[i].Cue = strings.TrimSpace(
			requestBody.Notes[i].Cue,
		)

		if len(requestBody.Notes[i].Content) == 0 {
			return infra.CustomResponse(
				context,
				fiber.StatusBadRequest,
				false,
				nil,
				fmt.Sprintf(
					"Note content at position %v is empty",
					i+1,
				),
			)
		}

		item, err := controller.noteRepo.CreateOne(
			requestBody.Notes[i].Content,
			requestBody.Notes[i].Cue,
		)

		if err != nil {
			return infra.CustomResponse(
				context,
				fiber.StatusBadRequest,
				false,
				nil,
				err.Error(),
			)
		}

		requestBody.Notes[i] = item
	}

	cornellNote, createNoteErr := getRepo().CreateOne(
		requestBody.Topic,
		requestBody.Subjects,
		requestBody.Notes,
	)

	if createNoteErr != nil {
		return infra.CustomResponse(
			context,
			fiber.StatusBadRequest,
			false,
			nil,
			createNoteErr.Error(),
		)
	}

	return infra.CustomResponse(
		context,
		fiber.StatusCreated,
		true,
		cornellNote,
		"",
	)
}
