// Package server implements a http server using the go standard lib.
package server

import (
	"log"
	"net/http"

	"github.com/codeYann/web-chat/controllers"
	"github.com/codeYann/web-chat/settings"
)

var apiSettings *settings.APIConfig = settings.ExportAPIConfig()

func createHTTPServer() {
	http.HandleFunc("/", controllers.Users)
	http.ListenAndServe(apiSettings.Port, nil)
}

// Run function exec the web server.
func Run() {
	log.Printf("Server Running on Localhost%s", apiSettings.Port)
	createHTTPServer()
}
