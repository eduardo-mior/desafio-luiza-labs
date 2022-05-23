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
	router.Use(middlewares.PanicRecovery())
}

func setupHealthCheckRoutes(router *gin.Engine) {
}

func setupApplicationRoutes(router *gin.Engine) {
	_ = router.Group("", middlewares.AuthJWT()) // Grupo de rotas que serão protegidos por autenticação JWT
}

func setupSwaggerDocumentationRoutes(router *gin.Engine) {
}
