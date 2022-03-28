package services

import (
	"fmt"
	"testing"
)

func TestGetRandomJoke(t *testing.T) {
	randJokeService := &RandomJokeService{}

	result, err := randJokeService.GetRandomJoke("slavik", "semeniv")

	if err != nil {
		t.Error(err, "")
	}

	fmt.Println(result.Value)
}
