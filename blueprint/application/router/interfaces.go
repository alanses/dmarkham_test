package router

import "test_dmarkham/cmd/kernel"

type ApiRouterInterface interface {
	RegisterRouters(httpKernel *kernel.HttpKernel)
}
