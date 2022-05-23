package main

import (
	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func setupHealthCheckRoutes(router *gin.Engine) {
}

func setupApplicationRoutes(router *gin.Engine) {
}

func setupSwaggerDocumentationRoutes(router *gin.Engine) {
}
