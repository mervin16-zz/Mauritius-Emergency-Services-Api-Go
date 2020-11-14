package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

var singleJSONTest string = `{
	"identifier": "security-police-direct-1",
	"name": "Police Direct Line 1",
	"type": "SECURITY",
	"icon": "https://img.icons8.com/fluent/100/000000/policeman-male.png",
	"number": 999
}`

var multipleJSONTest string = `[
	{
		"identifier": "security-police-direct-1",
		"name": "Police Direct Line 1",
		"type": "SECURITY",
		"icon": "https://img.icons8.com/fluent/100/000000/policeman-male.png",
		"number": 999
	},
	{
		"identifier": "health-samu",
		"name": "Samu",
		"type": "HEALTH",
		"icon": "https://img.icons8.com/bubbles/100/000000/ambulance.png",
		"number": 114
	}
]`

type serviceStruct struct {
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Icon       string `json:"icon"`
	Number     int    `json:"number"`
}

func oneService(c *gin.Context) {

	//language := c.Param("language")

	var service serviceStruct

	json.Unmarshal([]byte(singleJSONTest), &service)

	c.JSON(200, service)

	//fmt.Printf("%T\n", singleJSONTest)
}

func allServices(c *gin.Context) {

	//language := c.Param("language")

	file, _ := ioutil.ReadFile("data/services_en.json")

	var keys []serviceStruct

	json.Unmarshal([]byte(file), &keys)

	c.JSON(200, gin.H{"services": keys, "message": "", "success": true})
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
