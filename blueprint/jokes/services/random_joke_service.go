package services

import (
	"encoding/json"
	"net/http"
	"test_dmarkham/blueprint/application/helpers"
	"test_dmarkham/blueprint/jokes/signatures"
)

type RandomJokeService struct {}

func (randomJokeService *RandomJokeService) GetRandomJoke(firstName *string, lastName *string) (*signatures.RandomJoke, error) {
	request, requestError := http.NewRequest("GET","http://api.icndb.com/jokes/random", nil)

	if requestError != nil {
		return nil, requestError
	}

	query := request.URL.Query()
	query.Add("firstName", helpers.ToString(firstName))
	query.Add("lastName", helpers.ToString(lastName))
	query.Add("limitTo", "nerdy")

	request.URL.RawQuery = query.Encode()

	client := &http.Client{}

	resp, responseError := client.Do(request)

	if responseError != nil {
		return nil, responseError
	}

	randomJoke := &signatures.RandomJoke{}
	readBodyError := json.NewDecoder(resp.Body).Decode(randomJoke)

	if readBodyError != nil {
		return nil, readBodyError
	}

	return randomJoke, nil
}

func NewRandomJokeService() RandomJokeInterface {
	return &RandomJokeService{}
}