package api

import (
	"test_dmarkham/blueprint/application/router"
	"test_dmarkham/cmd/kernel"
)

type Application struct {
	ApiRouter router.ApiRouterInterface
}

func NewHttpApplication() *Application {
	return &Application{
		ApiRouter: &router.ApiRouter{},
	}
}

func (application *Application) RunApplication(httpKernel *kernel.HttpKernel) {
	application.ApiRouter.RegisterRouters(httpKernel)
}