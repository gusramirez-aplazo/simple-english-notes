package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gusramirez-aplazo/simple-english-notes/pakages/controllers"
	"github.com/gusramirez-aplazo/simple-english-notes/pakages/database"
	"github.com/gusramirez-aplazo/simple-english-notes/pakages/models"
	"github.com/gusramirez-aplazo/simple-english-notes/pakages/routes"
	"log"
	"net/http"
)

func init() {
	database.Connect()
	models.RunMigrations(database.GetDbClient())
}

func main() {
	const port = 3000

	router := mux.NewRouter()

	controller := &controllers.Controller{}

	routes.Start(router, controller, database.GetDbClient())

	fmt.Printf("Server listen on port: %v", port)

	address := fmt.Sprintf(":%v", port)

	log.Fatal(http.ListenAndServe(address, router))
}
