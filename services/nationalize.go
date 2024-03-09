package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"rest-api-challenge/dtos"
)

func GetNationalityByName(personName string) (dtos.NationalizeDTO, error) {
	url := fmt.Sprintf("https://api.nationalize.io/?name=%s", personName)

	resp, err := http.Get(url)
	if err != nil {
		return dtos.NationalizeDTO{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return dtos.NationalizeDTO{}, err
	}

	var result dtos.NationalizeDTO
	err = json.Unmarshal(body, &result)
	if err != nil {
		return dtos.NationalizeDTO{}, err
	}

	return result, nil
}
