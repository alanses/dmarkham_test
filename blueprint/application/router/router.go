package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	jokesRouter "test_dmarkham/blueprint/jokes/router"
	"test_dmarkham/cmd/kernel"
)

type ApiRouter struct{}

func (appRouting *ApiRouter) RegisterRouters(httpKernel *kernel.HttpKernel) {
	apiRouter := mux.NewRouter()

	appRouting.registerApplicationRoutes(apiRouter, httpKernel)
	appRouting.registerJokesRoutes(apiRouter, httpKernel)

	fmt.Println("Server is listening...")

	http.ListenAndServe(":81", apiRouter)
}

func (appRouting *ApiRouter) registerApplicationRoutes(apiRouter *mux.Router, httpKernel *kernel.HttpKernel) {
	RegisterSwaggerRoutes(apiRouter, httpKernel)
}

func (appRouting ApiRouter) registerJokesRoutes(apiRouter *mux.Router, httpKernel *kernel.HttpKernel) {
	jokesRouter.RegisterJokesRoutes(apiRouter, httpKernel)
}