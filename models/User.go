// Package models define structs to map database tables.
package models

import (
	"fmt"
	"log"

	"github.com/codeYann/web-chat/database"
)

// User struct defines all information User must have.
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
}

// Users type is a list of user
type Users []User

// CreateUser return a pointer to a new user
func CreateUser(id int, name, email, password, nickame string) *User {
	return &User{
		ID:       id,
		Name:     name,
		Email:    email,
		Password: password,
		Nickname: nickame,
	}
}

// FindAll returns a list of users gets by 'SELECT * FROM users' query.
func FindAll() Users {
	var usersList Users

	connection, err := database.OpenConnection()
	if err != nil {
		log.Fatal("Unable to connect to the database")
	}

	defer connection.Close()

	response, err := connection.Query(`SELECT * FROM users`)
	if err != nil {
		log.Fatal("It isn't possible to run 'SELECT * FROM useres' ")
	}

	defer response.Close()

	for response.Next() {
		var user User

		err := response.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Password,
			&user.Nickname,
		)
		if err != nil {
			log.Fatal("Error on iterate over rows returned by 'SELECT * FROM users' query.")
		}

		usersList = append(usersList, user)
	}
	return usersList
}

// FindOne returns a single user get by "SELECT * FROM users WHERE users.id = id"
func FindOne(ID uint64) User {
	var user User

	connection, err := database.OpenConnection()
	if err != nil {
		log.Fatal("Unable to connect to the database")
	}

	defer connection.Close()

	query := fmt.Sprintf(`SELECT * FROM users WHERE users.id = %d`, ID)

	response, err := connection.Query(query)
	if err != nil {
		log.Fatal("It's not possible to run 'SELECT * FROM useres WHERE users.id = /users/{id}'")
	}

	defer response.Close()

	for response.Next() {

		err := response.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Password,
			&user.Nickname,
		)
		if err != nil {
			log.Fatal("Error on iterate over rows returned by 'SELECT * FROM users' query.")
		}

	}

	return user
}

// Save function creates a new user in the database.
func Save(user User) {
	connection, err := database.OpenConnection()
	if err != nil {
		log.Fatal("Failed on connection to the database")
	}

	defer connection.Close()

	query := fmt.Sprintf(
		`INSERT INTO users
	       (name, email, password, nickname)
     VALUES
	       ('%s', '%s', '%x', '%s')`,
		user.Name,
		user.Email,
		user.Password,
		user.Nickname,
	)

	response, err := connection.Query(query)
	log.Println(err)
	if err != nil {
		log.Fatal(`It cannot run INSERT INTO query`)
	}

	defer response.Close()
}
