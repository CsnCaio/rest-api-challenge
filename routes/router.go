package routes

import (
	"net/http"
	"rest-api-challenge/controllers"
)

func HandleRoutes() {
	http.Handle("GET /zip/{zip}", http.HandlerFunc(controllers.GetLocationInfoByZipCode))
	http.Handle("GET /nationality/{name}", http.HandlerFunc(controllers.GetNationalityByName))

	http.Handle("POST /users", http.HandlerFunc(controllers.CreateUser))
	http.Handle("PUT /users/{id}", http.HandlerFunc(controllers.UpdateUser))
	http.Handle("GET /users/{id}", http.HandlerFunc(controllers.GetUser))
	http.Handle("GET /users", http.HandlerFunc(controllers.GetUsers))
	http.Handle("DELETE /user/{id}", http.HandlerFunc(controllers.DeleteUser))
}
