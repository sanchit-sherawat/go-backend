package services

import (
	"fmt"

	"github.com/sanchit-sherawat/go-backend.git/models"
)

var users = []models.User{
	{Username: "John Doe", Password: "password", Email: "john@example.com"},
}

func GetAllUsers() []models.User {
	return users
}

func CreateUser(user models.User) []models.User {
	users = append(users, user)
	fmt.Println(users)
	return users
}
