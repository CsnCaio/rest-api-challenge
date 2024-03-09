package main

import (
	"fmt"
	"net/http"
	"rest-api-challenge/config"
	"rest-api-challenge/routes"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("============ FATEC CHALLENGE ============")
	fmt.Println("Welcome to the FATEC Challenge!")

	routes.HandleRoutes()
	config.LoadEnv()
	config.InitDatabase()

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
