package controllers

import (
	"encoding/json"
	"net/http"
)

func (controller Controller) HomeController(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(response).Encode(NewSuccessResponse("Hello From Server!!"))
}
