package controllers

import (
	"encoding/json"
	"net/http"
	"rest-api-challenge/services"
	"unicode"
)

func sanitizeZipCode(zipCode string) string {
	sanitized := ""
	for _, char := range zipCode {
		if unicode.IsDigit(char) {
			sanitized += string(char)
		}
	}

	return sanitized
}

func GetLocationInfoByZipCode(response http.ResponseWriter, request *http.Request) {
	zipCode := request.PathValue("zip")
	
	zipCode = sanitizeZipCode(zipCode)
	if len(zipCode) != 8 {
		http.Error(response, "Invalid zip code", http.StatusBadRequest)
		return
	}

	locationInfo, err := services.GetLocationInfoByZipCode(zipCode)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(response).Encode(locationInfo)
}