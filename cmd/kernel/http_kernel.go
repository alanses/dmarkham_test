package kernel

import (
	"go.uber.org/zap/zapcore"
	"test_dmarkham/blueprint/application/config"
	"test_dmarkham/blueprint/application/logger"
	"test_dmarkham/blueprint/application/services"
)

type HttpKernel struct {
	ApplicationConfig   *config.ApplicationConfig
	AppLogger           *logger.AppLogger
	ApplicationServices *services.ApplicationServices
}

func NewHttpKernel() *HttpKernel {
	return &HttpKernel{}
}

func (httpKernel *HttpKernel) RegisterConfig() error {
	applicationConfig, err := config.NewApplicationConfig()

	httpKernel.ApplicationConfig = applicationConfig

	return err
}

func (httpKernel *HttpKernel) RegisterLogger() {
	httpKernel.AppLogger = logger.InitLogger(zapcore.DebugLevel)
}

func (httpKernel *HttpKernel) RegisterApplicationServices()  {
	httpKernel.ApplicationServices = services.NewApplicationServices()
}
