package config

type ApplicationConfig struct {
	SwaggerConfig  *SwaggerConfig
}

func NewApplicationConfig() (*ApplicationConfig, error) {
	var err error
	applicationConfig := &ApplicationConfig{}

	applicationConfig.registerSwaggerConfig()

	return applicationConfig, err
}

func (ApplicationConfig *ApplicationConfig) registerSwaggerConfig() {
	ApplicationConfig.SwaggerConfig = registerSwaggerConfig()
}