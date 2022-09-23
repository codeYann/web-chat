// Package controllers must interact with models and export its data in JSON format.
package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/codeYann/web-chat/models"
	"github.com/gorilla/mux"
)

// Users returns all users in JSON format.
// It uses models.FindAll() function to get every possible user row in the database.
func Users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	allUsers := models.FindAll()

	if err := json.NewEncoder(w).Encode(allUsers); err != nil {
		log.Fatal("It's not possible to send json")
	}

	w.WriteHeader(http.StatusOK)
}

// UserByID returns a single user in JSON format.
// It uses models.FindOne() to get a this single user.
func UserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, _ := strconv.Atoi(params["ID"])

	user := models.FindOne(uint64(userID))

	if err := json.NewEncoder(w).Encode(user); err != nil {
		log.Fatal("It's not possible to send json")
	}

	w.WriteHeader(http.StatusOK)
}

// StoreUser insert a new user to the database.
// It uses models.Save() to store a new user.
func StoreUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Fatal("Error on decode body params")
	}

	models.Save(user)

	if err := json.NewEncoder(w).Encode(user); err != nil {
		log.Fatal("Error on Encode User data")
	}

	w.WriteHeader(http.StatusOK)
}
