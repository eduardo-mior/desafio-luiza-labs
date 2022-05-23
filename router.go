package main

import (
	"desafio_luizalabs/middlewares"

	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func setupGlobalMiddlewares(router *gin.Engine) {
	router.NoRoute(middlewares.NotFound())
}

func setupHealthCheckRoutes(router *gin.Engine) {
}

func setupApplicationRoutes(router *gin.Engine) {
}

func setupSwaggerDocumentationRoutes(router *gin.Engine) {
}
