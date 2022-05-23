package main

func main() {

	router := initRouter()

	setupHealthCheckRoutes(router)
	setupApplicationRoutes(router)
	setupSwaggerDocumentationRoutes(router)

	router.Run()

}
