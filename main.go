package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

type serviceStruct struct {
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Icon       string `json:"icon"`
	Number     int    `json:"number"`
}

func getServiceFileByLanguage(language string) string {
	// Get the correct json file
	// according to appropriate language
	// If language doesn't exists, returns default (english)
	switch language {
	case "en":
		return "services_en.json"
	case "fr":
		return "services_fr.json"
	default:
		return "services_en.json"
	}
}

func oneService(context *gin.Context) {

	// Fetches the services json file per
	// appropriate language
	servicesFile := getServiceFileByLanguage(context.Param("language"))

	// Reads the file
	file, _ := ioutil.ReadFile("data/" + servicesFile)

	// Initialize array of services struct
	var services []serviceStruct

	// Unmarshal JSON data to struct
	json.Unmarshal([]byte(file), &services)

	// Iterate through each services to get
	// by id
	id := context.Param("id")
	for _, service := range services {
		if service.Identifier == id {
			// Loads the data
			context.JSON(200, gin.H{"services": []serviceStruct{service}, "message": "", "success": true})
			return
		}
	}

	// If none found
	// Returns error
	context.JSON(404, gin.H{"services": []serviceStruct{}, "message": "No services found under id " + id, "success": false})
}

func allServices(context *gin.Context) {

	// Fetches the services json file per
	// appropriate language
	servicesFile := getServiceFileByLanguage(context.Param("language"))

	// Reads the file
	file, _ := ioutil.ReadFile("data/" + servicesFile)

	// Initialize array of services struct
	var services []serviceStruct

	// Unmarshal JSON data to struct
	json.Unmarshal([]byte(file), &services)

	// Loads the data
	context.JSON(200, gin.H{"services": services, "message": "", "success": true})
}

func setupRoutes(engine *gin.Engine) {
	engine.GET("/mesg/:language/services", allServices)
	engine.GET("/mesg/:language/service/:id", oneService)
}

func main() {

	// Creates the Gin Engine
	engine := gin.Default()

	// Setup the API Routes
	setupRoutes(engine)

	// Run the engine
	engine.Run()
}
