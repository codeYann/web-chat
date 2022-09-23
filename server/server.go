// Package server implements a http server using the go standard lib.
package server

import (
	"github.com/codeYann/web-chat/routes"
)

// Run function exec the web server.
func Run() {
	server := routes.CreateRoutes()
	server.Init()
}
