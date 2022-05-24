package main

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {

	router := initRouter()

	enableCORS(router)
	setupGlobalMiddlewares(router)
	setupHealthCheckRoutes(router)
	setupApplicationRoutes(router)
	setupSwaggerDocumentationRoutes(router)

	router.Run()

}
