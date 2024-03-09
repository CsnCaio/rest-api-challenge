package controllers

import (
	"encoding/json"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"rest-api-challenge/services"
)

func sanitizePersonName(name string) string {
	// Transform characters into url-form-encoded ones
	sanitizedName := url.QueryEscape(name)

	// Remove numbers
	sanitizedName = regexp.MustCompile("[0-9]").ReplaceAllString(sanitizedName, "")

	// Replace -, ., and _ with " "
	sanitizedName = strings.ReplaceAll(sanitizedName, "-", " ")
	sanitizedName = strings.ReplaceAll(sanitizedName, ".", " ")
	sanitizedName = strings.ReplaceAll(sanitizedName, "_", " ")

	return sanitizedName
}

func GetNationalityByName(response http.ResponseWriter, request *http.Request) {
	personName := sanitizePersonName(request.PathValue("name"))

	personInfo, err := services.GetNationalityByName(personName)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(response).Encode(personInfo)
}
