package services

import (
	"rest-api-challenge/dtos"
	"rest-api-challenge/repositories"
)

func CreateUser(createUserDto dtos.CreateUserDto) (dtos.UserDto, error) {
	return repositories.CreateUser(createUserDto)
}

func UpdateUser(id int, updateUserDto dtos.UpdateUserDto) (dtos.UserDto, error) {
	return repositories.UpdateUser(id, updateUserDto)
}

func GetUser(id int) (dtos.UserDto, error) {
	return repositories.GetUser(id)
}

func GetUsers() ([]dtos.UserDto, error) {
	return repositories.GetUsers()
}

func DeleteUser(id int) {
	repositories.DeleteUser(id)
}
