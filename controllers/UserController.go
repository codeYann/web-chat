// Package controllers must interact with models and export its data in JSON format.
package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/codeYann/web-chat/database"
	"github.com/codeYann/web-chat/repositories"
	"github.com/codeYann/web-chat/services"
	"github.com/gorilla/mux"
)

func repositoryWrapper(connection *sql.DB) *services.UserService {
	repository := repositories.CreateIPostgresRepository(connection)
	return services.CreateUserServices(repository)
}

// Users returns all users in JSON format.
// It uses services.GetAllUsers() function to get every possible user.
func Users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	connection, _ := database.OpenConnection()
	userServices := repositoryWrapper(connection)

	response, _ := userServices.GetAllUsers()

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Fatal("It's not possible to send json", err.Error())
	}
}

// UserByID returns a single user in JSON format.
// It uses services.GetUser() to an user.
func UserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, _ := strconv.Atoi(params["ID"])

	connection, _ := database.OpenConnection()
	userServices := repositoryWrapper(connection)

	response, _ := userServices.GetUser(uint64(userID))

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Fatal("It's not possible to send json", err.Error())
	}
}

// StoreUser insert a new user to the database.
// It uses models.Save() to store a new user.
func StoreUser(w http.ResponseWriter, r *http.Request) {
	var user struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Nickname string `json:"nickname"`
	}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Fatal("Error on decode body params", err.Error())
	}

	connection, _ := database.OpenConnection()
	userServices := repositoryWrapper(connection)

	response, _ := userServices.CreateUser(user)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Fatal("Error on Encode User data", err.Error())
	}
}

// UpdateUser updates a user nickname
// It uses repository.UpdateUser and return this user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var userInfo struct {
		ID       uint64
		Nickname string
	}

	if err := json.NewDecoder(r.Body).Decode(&userInfo); err != nil {
		log.Fatal("Error on decode body params", err.Error())
	}

	connection, _ := database.OpenConnection()
	userServices := repositoryWrapper(connection)

	response, _ := userServices.UpdateUser(userInfo.ID, userInfo.Nickname)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Fatal("Error on encode user data", err.Error())
	}
}

// RemoveUser removes a user.
// It ueses repository.RemoveOne and return this user
func RemoveUser(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)
	id, _ := strconv.Atoi(userID["ID"])

	connection, _ := database.OpenConnection()
	userServices := repositoryWrapper(connection)

	response, _ := userServices.RemoveUser(uint64(id))

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Fatal("Error on encode user data", err.Error())
	}
}
