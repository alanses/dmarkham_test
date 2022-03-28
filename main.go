package main

import (
	"test_dmarkham/blueprint/application/errors_handlers"
	"test_dmarkham/cmd/api"
	"test_dmarkham/cmd/kernel"
)

func main()  {
	var err error
	httpKernel := kernel.NewHttpKernel()

	httpKernel.RegisterLogger()
	httpKernel.RegisterApplicationServices()

	err = httpKernel.RegisterConfig()

	if err != nil {
		errors_handlers.HandleError(httpKernel.AppLogger.Logger, err)
	}

	application := api.NewHttpApplication()

	application.RunApplication(httpKernel)
}
