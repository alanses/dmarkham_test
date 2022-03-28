package services

import (
	"encoding/json"
	"net/http"
	"test_dmarkham/blueprint/jokes/signatures"
)

type RandomNameService struct {}

func NewRandomNameService() RandomNameInterface {
	return &RandomNameService{}
}

func (randomNameService *RandomNameService) GetRandomName() (*signatures.FullNameInfo, error) {
	resp, responseError := http.Get("https://names.mcquay.me/api/v0/")

	if responseError != nil {
		return nil, responseError
	}

	fullNameInfo := &signatures.FullNameInfo{}
	readBodyError := json.NewDecoder(resp.Body).Decode(fullNameInfo)

	if readBodyError != nil {
		return nil, readBodyError
	}

	return fullNameInfo, nil
}