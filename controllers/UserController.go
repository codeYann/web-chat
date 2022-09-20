// Package controllers must interact with models and export its data in JSON format.
package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/codeYann/web-chat/models"
)

// Users returns all users in JSON format.
// It uses models.GetAllUsers function to get every possible user row in the database.
func Users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	allUsers := models.GetAllUsers()

	err := json.NewEncoder(w).Encode(allUsers)
	if err != nil {
		log.Fatal("It's not possible to send json")
	}
}
