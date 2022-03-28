package router

import (
	"github.com/gorilla/mux"
	"test_dmarkham/blueprint/jokes/handlers"
	"test_dmarkham/cmd/kernel"
)

func RegisterJokesRoutes(apiRouter *mux.Router, httpKernel *kernel.HttpKernel) {
	jokesHandler := handlers.NewJokesHandler(
		httpKernel.ApplicationServices.RandomNameService,
		httpKernel.ApplicationServices.RandomJokeService,
		httpKernel.AppLogger)

	subRouter := apiRouter.PathPrefix("/api/joke").Subrouter()

	subRouter.HandleFunc("/", jokesHandler.GetJoke).Methods("GET")
}
