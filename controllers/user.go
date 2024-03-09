package controllers

import (
	"encoding/json"
	"net/http"
	"rest-api-challenge/dtos"
	"rest-api-challenge/services"
	"strconv"
)

func CreateUser(response http.ResponseWriter, request *http.Request) {
	var createUserDto dtos.CreateUserDto
	err := json.NewDecoder(request.Body).Decode(&createUserDto)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	createdUser, err := services.CreateUser(createUserDto)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(response).Encode(createdUser)
}

func UpdateUser(response http.ResponseWriter, request *http.Request) {
	userIdStr := request.PathValue("id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	var updateUserDto dtos.UpdateUserDto
	err = json.NewDecoder(request.Body).Decode(&updateUserDto)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	updatedUser, err := services.UpdateUser(userId, updateUserDto)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(response).Encode(updatedUser)
}

func GetUser(response http.ResponseWriter, request *http.Request) {
	userIdStr := request.PathValue("id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := services.GetUser(userId)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(response).Encode(user)
}

func GetUsers(response http.ResponseWriter, request *http.Request) {
	users, err := services.GetUsers()
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(response).Encode(users)
}

func DeleteUser(response http.ResponseWriter, request *http.Request) {
	userIdStr := request.PathValue("id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	services.DeleteUser(userId)
}
