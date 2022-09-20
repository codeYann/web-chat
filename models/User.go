// Package models define structs to map database tables.
package models

import (
	"log"

	"github.com/codeYann/web-chat/database"
)

// Users struct defines all information Users must have.
type Users struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	password string `json:"-"`
	Nickname string `json:"nickname"`
}

// CreateUsers return a pointer to a new user
func CreateUsers(id int, name, email, passowrd, nickame string) *Users {
	return &Users{
		ID:       id,
		Name:     name,
		Email:    email,
		password: passowrd,
		Nickname: nickame,
	}
}

// GetAllUsers returns a list of users gets by 'SELECT * FROM users' query.
func GetAllUsers() []Users {
	usersList := make([]Users, 0)

	connection, err := database.OpenConnection()
	if err != nil {
		log.Fatal("Unable to connect to the database")
	}

	var (
		ID       int
		Name     string
		Email    string
		Password string
		Nickname string
	)

	defer connection.Close()

	response, err := connection.Query(`SELECT * FROM users`)
	if err != nil {
		log.Fatal("It isn't possible to run 'SELECT * FROM useres' ")
	}

	defer response.Close()

	for response.Next() {
		err := response.Scan(
			&ID,
			&Name,
			&Email,
			&Password,
			&Nickname,
		)
		if err != nil {
			log.Fatal("Error on iterate over rows returned by 'SELECT * FROM users' query.")
		}
		user := CreateUsers(ID, Name, Email, Password, Nickname)

		usersList = append(usersList, *user)
	}
	return usersList
}
