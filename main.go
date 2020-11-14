package main

import (
	"mes/routes"

	"github.com/gin-gonic/gin"
)

func setupRoutes(engine *gin.Engine) {
	engine.GET("/mesg/:language/services", routes.AllServices)
	engine.GET("/mesg/:language/service/:id", routes.OneService)
	engine.GET("/mesg/:language/services/search", routes.SearchService)
}

func main() {

	// Creates the Gin Engine
	engine := gin.Default()

	// Setup the API Routes
	setupRoutes(engine)

	// Run the engine
	engine.Run()
}
