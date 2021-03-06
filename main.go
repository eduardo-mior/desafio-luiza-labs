package main

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {

	loadEnviromentVariables()

	initSentry()

	router := initRouter()

	enableCORS(router)
	setupSentryMiddleware(router)
	setupGlobalMiddlewares(router)
	setupHealthCheckRoutes(router)
	setupApplicationRoutes(router)
	setupSwaggerDocumentationRoutes(router)

	router.Run()

}
