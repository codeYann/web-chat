// Package models define structs to map database tables.
package models

// User struct defines all information User must have.
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
}

// Users type is a list of user.
type Users []User

// UserRepository defines all operations User must have.
type UserRepository interface {
	FindAll() (Users, error)
	FindOne(ID uint64) (User, error)
	SaveOne(user User) (User, error)
	UpdateOne(ID uint64, nickname string) (User, error)
	DeleteOne(ID uint64) (User, error)
}
