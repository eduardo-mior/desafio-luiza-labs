package main

import (
	auth_controller "desafio_luizalabs/controllers/auth"
	cep_controller "desafio_luizalabs/controllers/cep"
	healthcheck_controller "desafio_luizalabs/controllers/healthcheck"
	docs "desafio_luizalabs/docs"
	"desafio_luizalabs/middlewares"
	"desafio_luizalabs/repositories"
	"desafio_luizalabs/services"

	swaggerFiles "github.com/swaggo/files"
	swaggerGin "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	return router
}

func setupGlobalMiddlewares(router *gin.Engine) {
	router.NoRoute(middlewares.NotFound())
	router.Use(middlewares.PanicRecovery())
}

func setupHealthCheckRoutes(router *gin.Engine) {
	router.GET("/health-check", healthcheck_controller.HealthCheck())
}

func setupApplicationRoutes(router *gin.Engine) {

	jwtService := services.ServiceAuth{}
	router.GET("/gerar-token-jwt", auth_controller.GerarTokenJWT(&jwtService))

	cepRepository := repositories.RepositoryCEP{}
	cepService := services.ServiceCEP{IRepositoryCEP: &cepRepository}
	router.GET("/buscar-cep/:cep", middlewares.AuthJWT(), cep_controller.BuscarCEP(cepService))

}

func setupSwaggerDocumentationRoutes(router *gin.Engine) {

	docs.SwaggerInfo.Title = "Desafio Luiza Labs"
	docs.SwaggerInfo.Description = "Documentação da API de Consulta de CEP"
	docs.SwaggerInfo.Version = "1.0.0"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	swagger := router.Group("/docs")
	swagger.GET("/*any", swaggerGin.WrapHandler(swaggerFiles.Handler))

}
