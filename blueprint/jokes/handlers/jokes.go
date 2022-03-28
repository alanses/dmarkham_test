package handlers

import (
	"net/http"
	"test_dmarkham/blueprint/application/errors_handlers"
	"test_dmarkham/blueprint/application/logger"
	"test_dmarkham/blueprint/jokes/responses"
	"test_dmarkham/blueprint/jokes/services"
)

type JokesHandler struct {
	RandomNameService services.RandomNameInterface
	RandomJokeService services.RandomJokeInterface
	AppLogger         *logger.AppLogger
}

func NewJokesHandler(randomNameService services.RandomNameInterface, randomJokeService services.RandomJokeInterface, appLogger *logger.AppLogger) *JokesHandler {
	return &JokesHandler{
		RandomNameService: randomNameService,
		RandomJokeService: randomJokeService,
		AppLogger:         appLogger,
	}
}

// GetJoke GetJoke
// @Summary Get Joke
// @Produce  json
// @tags Jokes
// @Security ApiKeyAuth
// @Success 200 {object} responses.JokeResponse
// @Router /api/joke/ [get]
func (jokesHandler *JokesHandler) GetJoke(writer http.ResponseWriter, request *http.Request) {
	fullNameInfo, randomNameError := jokesHandler.RandomNameService.GetRandomName()

	if randomNameError != nil {
		errors_handlers.HandleError(jokesHandler.AppLogger.Logger, randomNameError)
		return
	}

	var fullName, lastName *string

	if fullNameInfo != nil {
		fullName = &fullNameInfo.FirstName
		lastName = &fullNameInfo.LastName
	}

	joke, randomJokeError := jokesHandler.RandomJokeService.GetRandomJoke(fullName, lastName)

	if randomJokeError != nil {
		errors_handlers.HandleError(jokesHandler.AppLogger.Logger, randomJokeError)
		return
	}

	response := responses.NewJokeResponse(fullNameInfo, joke)

	response.ToJson(writer)
}
