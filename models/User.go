package models

import (
	"github.com/codeYann/web-chat/utils"
)

type UserInterface interface {
	GetAllUsers() []User
	GetUserById(ID int) User
}

type User struct {
	Name     string `json:"Name"`
	Nickname string `json:"Nickname"`
	Password string `json:"-"`
	ID       int    `json:"ID"`
}

func CreateUser(name, nickName, password string) *User {
	return &User{
		Name:     name,
		Nickname: name,
		Password: utils.GenerateHashString(password),
	}
}
