package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"rest-api-challenge/dtos"
)

func GetLocationInfoByZipCode(zipCode string) (dtos.ViaCepDTO, error) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", zipCode)

	resp, err := http.Get(url)
	if err != nil {
		return dtos.ViaCepDTO{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return dtos.ViaCepDTO{}, err
	}

	var result dtos.ViaCepDTO
	err = json.Unmarshal(body, &result)
	if err != nil {
		return dtos.ViaCepDTO{}, err
	}

	return result, nil
}
