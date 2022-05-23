package main

func main() {

	router := initRouter()

	setupGlobalMiddlewares(router)
	setupHealthCheckRoutes(router)
	setupApplicationRoutes(router)
	setupSwaggerDocumentationRoutes(router)

	router.Run()

}
