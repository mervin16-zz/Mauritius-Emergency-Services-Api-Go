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

func setupServer() *gin.Engine {
	// Gin Mode
	gin.SetMode(gin.ReleaseMode)

	// Creates the Gin Engine
	engine := gin.New()

	// Setup the API Routes
	setupMainRoutes(engine)

	// Setup Error Routes
	setupErrorRoutes(engine)

	// Return engine
	return engine
}

func main() {
	// Run the engine
	setupServer().Run()
}
