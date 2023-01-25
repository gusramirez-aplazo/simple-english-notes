package routes

import (
	"github.com/gorilla/mux"
	"github.com/gusramirez-aplazo/simple-english-notes/pakages/controllers"
)

func Start(router *mux.Router, controller *controllers.Controller) {
	const apiPrefix = "/api/v1"

	router.HandleFunc("/", controller.HomeController)

	//subrouter := router.PathPrefix(apiPrefix).Subrouter()

	//subrouter.HandleFunc("/category")
}
