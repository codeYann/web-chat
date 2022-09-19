// Package server implements a http server using the go standard lib.
package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/codeYann/web-chat/models"
	"github.com/codeYann/web-chat/settings"
)

var apiSettings *settings.APIConfig = settings.ExportAPIConfig()

func createHTTPServer() {
	log.Println(models.GetAllUsers())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s", "this is my first web server "+r.URL.Path)
	})

	http.ListenAndServe(apiSettings.Port, nil)
}

// Run function exec the web server.
func Run() {
	log.Printf("Server Running on Localhost%s", apiSettings.Port)
	createHTTPServer()
}
