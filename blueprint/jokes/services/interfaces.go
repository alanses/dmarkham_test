package services

import "test_dmarkham/blueprint/jokes/signatures"

type RandomNameInterface interface {
	GetRandomName() (*signatures.FullNameInfo, error)
}

type RandomJokeInterface interface {
	GetRandomJoke(firstName *string, lastName *string) (*signatures.RandomJoke, error)
}
