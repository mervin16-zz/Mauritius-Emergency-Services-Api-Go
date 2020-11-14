package main

import (
	"mes/routes"

	"github.com/gin-gonic/gin"
)

func setupMainRoutes(engine *gin.Engine) {
	engine.GET("/mesg/:language/services", routes.AllServices)
	engine.GET("/mesg/:language/service/:id", routes.OneService)
	engine.GET("/mesg/:language/services/search", routes.SearchService)
}

func setupErrorRoutes(engine *gin.Engine) {
	engine.NoRoute(routes.Error404Handler)
}

func main() {

	// Creates the Gin Engine
	engine := gin.Default()

	// Setup the API Routes
	setupMainRoutes(engine)

	// Setup Error Routes
	setupErrorRoutes(engine)

	// Run the engine
	engine.Run()
}
