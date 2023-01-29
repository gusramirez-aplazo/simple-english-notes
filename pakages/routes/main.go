package routes

import (
	"github.com/gorilla/mux"
	"github.com/gusramirez-aplazo/simple-english-notes/pakages/controllers"
	"gorm.io/gorm"
)

// TODO: handle not found routes
func Start(router *mux.Router, controller *controllers.Controller, clientDB *gorm.DB) {
	const apiPrefix = "/api/v1"

	router.HandleFunc("/", controller.HomeController).Methods("GET")

	subRouter := router.PathPrefix(apiPrefix).Subrouter()

	createTopicController := controller.CreateTopicControllerFactory(clientDB)

	subRouter.HandleFunc("/topic", createTopicController).Methods("POST")
}
