package routes

import (
	"github.com/gorilla/mux"
	"github.com/gusramirez-aplazo/simple-english-notes/pakages/controllers"
	"gorm.io/gorm"
)

func Start(router *mux.Router, controller *controllers.Controller, clientDB *gorm.DB) {
	const apiPrefix = "/api/v1"

	router.HandleFunc("/", controller.HomeController)

	//subrouter := router.PathPrefix(apiPrefix).Subrouter()

	//subrouter.HandleFunc("/category")
}
