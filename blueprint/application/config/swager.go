package config

import "os"

type SwaggerConfig struct {
	SwaggerHost string
	SwaggerUrl  string
}

func registerSwaggerConfig() *SwaggerConfig {
	swaggerConfig := new(SwaggerConfig)
	swaggerConfig.SwaggerHost = os.Getenv("SWAGGER_HOST")
	swaggerConfig.SwaggerUrl = os.Getenv("SWAGGER_URL")

	return swaggerConfig
}
