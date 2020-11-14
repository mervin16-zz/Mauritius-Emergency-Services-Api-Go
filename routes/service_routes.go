package routes

import (
	"encoding/json"
	"io/ioutil"
	"mes/helpers"
	structure "mes/models"
	"strings"

	"github.com/gin-gonic/gin"
)

func AllServices(context *gin.Context) {

	// Fetches the services json file per
	// appropriate language
	servicesFile := helpers.GetServiceFileByLanguage(context.Param("language"))

	// Reads the file
	file, _ := ioutil.ReadFile("data/" + servicesFile)

	// Initialize array of services struct
	var services []structure.Service

	// Unmarshal JSON data to struct
	json.Unmarshal([]byte(file), &services)

	// Loads the data
	context.JSON(200, gin.H{"services": services, "message": "", "success": true})
}

func OneService(context *gin.Context) {

	// Fetches the services json file per
	// appropriate language
	servicesFile := helpers.GetServiceFileByLanguage(context.Param("language"))

	// Reads the file
	file, _ := ioutil.ReadFile("data/" + servicesFile)

	// Initialize array of services struct
	var services []structure.Service

	// Unmarshal JSON data to struct
	json.Unmarshal([]byte(file), &services)

	// Iterate through each services to get
	// by id
	id := context.Param("id")
	for _, service := range services {
		if service.Identifier == id {
			// Loads the data
			context.JSON(200, gin.H{"services": []structure.Service{service}, "message": "", "success": true})
			return
		}
	}

	// If none found
	// Returns error
	context.JSON(404, gin.H{"services": []structure.Service{}, "message": "No services found under id " + id, "success": false})
}

func SearchService(context *gin.Context) {

	searchString := strings.ToLower(context.Query("query"))

	// Fetches the services json file per
	// appropriate language
	servicesFile := helpers.GetServiceFileByLanguage(context.Param("language"))

	// Reads the file
	file, _ := ioutil.ReadFile("data/" + servicesFile)

	// Initialize array of services struct
	var services []structure.Service

	// Unmarshal JSON data to struct
	json.Unmarshal([]byte(file), &services)

	// Iterate through each services
	// and check query that matches the name
	filteredServices := []structure.Service{}
	for _, service := range services {
		if strings.Contains(strings.ToLower(service.Name), searchString) {
			filteredServices = append(filteredServices, service)
		}
	}

	// If none found
	// Returns error
	context.JSON(200, gin.H{"services": filteredServices, "message": "", "success": true})
}
