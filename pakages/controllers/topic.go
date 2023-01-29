package controllers

import (
	"encoding/json"
	"github.com/gusramirez-aplazo/simple-english-notes/pakages/models"
	"gorm.io/gorm"
	"log"
	"net/http"
)

// TODO: sanitize request
// TODO: testing
// TODO: ensure correct capitlazation (all lower or upper)[only for names, titles and key parameters]
func (controller Controller) CreateTopicControllerFactory(clientDB *gorm.DB) func(http.ResponseWriter, *http.Request) {
	return func(response http.ResponseWriter, request *http.Request) {
		var topic models.Topic
		response.Header().Set("Content-Type", "application/json")

		if err := json.NewDecoder(request.Body).Decode(&topic); err != nil {

			response.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(response).Encode(NewErrorResponse("Bad request", 400))
			return
		}

		newTopic := clientDB.Create(&topic)

		if newTopic.Error != nil {

			response.WriteHeader(http.StatusBadRequest)
			// TODO: pass a sanitized error maybe with a global error factory
			_ = json.NewEncoder(response).Encode(NewErrorResponse("topic not created", 400))
			return
		}

		// TODO: handle nested structs
		res, err := json.Marshal(newTopic)

		if err == nil {
			log.Fatal("json parsing runtime error")
		}

		response.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(response).Encode(NewSuccessResponse(res))
	}
}
