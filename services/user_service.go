// Package services export a way to consume data from a repository and pass it to a controller
package services

import (
	"log"

	"github.com/codeYann/web-chat/models"
)

// UserService defines a struct that contains a repository that implements models.UserRepository.
type UserService struct {
	// Repository is a UserRepository interface
	Repository models.UserRepository
}

// CreateUserServices returns a new instance of UserService.
func CreateUserServices(repository models.UserRepository) *UserService {
	return &UserService{
		Repository: repository,
	}
}

// GetAllUsers returns a list of users.
// It uses a Repository.FindAll() to get those users
func (u *UserService) GetAllUsers() (models.Users, error) {
	response, err := u.Repository.FindAll()
	if err != nil {
		log.Fatal("Unable to return all users.", err.Error())
	}
	return response, err
}

// GetUser returns a single user.
// It uses a Repository.FindOne() to get this user.
func (u *UserService) GetUser(ID uint64) (models.User, error) {
	response, err := u.Repository.FindOne(ID)
	if err != nil {
		log.Fatal("Unable to return a user", err.Error())
	}
	return response, err
}

// CreateUser returns a new user.
// It uses a Repository.SaveOne() to create this user.
func (u *UserService) CreateUser(user models.User) (models.User, error) {
	response, err := u.Repository.SaveOne(user)
	if err != nil {
		log.Fatal("Unable to return a new user", err.Error())
	}
	return response, err
}

// RemoveUser removes a single user.
// It uses a Repository.DeleteOne to remove this user.
func (u *UserService) RemoveUser(ID uint64) (models.User, error) {
	response, err := u.Repository.DeleteOne(ID)
	if err != nil {
		log.Fatal("Unable to delete a new user", err.Error())
	}
	return response, err
}

// UpdateUser updates a single user.
// It uses a Repository.UpdateOne() to update this user.
func (u *UserService) UpdateUser(ID uint64, nickname string) (models.User, error) {
	response, err := u.Repository.UpdateOne(ID, nickname)
	if err != nil {
		log.Fatal("Unable to update an user", err.Error())
	}
	return response, err
}
