package controllers

import (
	"encoding/json"
	"github.com/gusramirez-aplazo/simple-english-notes/pakages/shared"
	"net/http"
)

type Controller struct {
}

func (controller Controller) HomeController(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	_ = json.NewEncoder(response).Encode(shared.ApiResponse{
		Ok:      true,
		Status:  200,
		Message: "",
		Content: "Hello World!",
	})
}
