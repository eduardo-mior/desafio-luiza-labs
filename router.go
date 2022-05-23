package main

import (
	auth_controller "desafio_luizalabs/controllers/auth"
	cep_controller "desafio_luizalabs/controllers/cep"
	healthcheck_controller "desafio_luizalabs/controllers/healthcheck"
	"desafio_luizalabs/middlewares"
	"desafio_luizalabs/repositories"
	"desafio_luizalabs/services"

	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
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
	router.GET("/buscar-cep/:cep", cep_controller.BuscarCEP(cepService))

}

func setupSwaggerDocumentationRoutes(router *gin.Engine) {
}
