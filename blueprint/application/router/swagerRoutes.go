package router

import (
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"test_dmarkham/cmd/kernel"
	_ "test_dmarkham/docs"
)

// RegisterSwaggerRoutes @title Jokes API
// @version 1.0
// @description This is documentation for Jokes
// @host localhost:81
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /
func RegisterSwaggerRoutes(router *mux.Router, httpKernel *kernel.HttpKernel) {
	swaggerConfig := httpKernel.ApplicationConfig.SwaggerConfig

	url := swaggerConfig.SwaggerHost + swaggerConfig.SwaggerUrl

	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL(url), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("#swagger-ui"),
	))
}
