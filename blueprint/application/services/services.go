package services

import "test_dmarkham/blueprint/jokes/services"

type ApplicationServices struct {
	RandomNameService services.RandomNameInterface
	RandomJokeService services.RandomJokeInterface
}

func NewApplicationServices() *ApplicationServices {
	return &ApplicationServices{
		RandomNameService: services.NewRandomNameService(),
		RandomJokeService: services.NewRandomJokeService(),
	}
}
