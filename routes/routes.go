// Package routes export all routes that application use
package routes

import (
	"log"
	"net/http"

	"github.com/codeYann/web-chat/controllers"
	"github.com/codeYann/web-chat/settings"
	"github.com/gorilla/mux"
)

// Routes struct export a mux.Router to handle all routes
type Routes struct {
	Router *mux.Router
}

// CreateRoutes is a factory function to return a new Routes
func CreateRoutes() *Routes {
	return &Routes{
		Router: mux.NewRouter(),
	}
}

// Init runs http server
func (r Routes) Init() {
	apiSettings := settings.ExportAPIConfig()

	r.Router.HandleFunc("/users", controllers.Users).Methods("GET")
	r.Router.HandleFunc("/users/{ID}", controllers.UserByID).Methods("GET")
  r.Router.HandleFunc("/users", controllers.StoreUser).Methods("POST")

	log.Printf("Server Running on Localhost%s", apiSettings.Port)
	log.Fatal(http.ListenAndServe(apiSettings.Port, r.Router))
}
