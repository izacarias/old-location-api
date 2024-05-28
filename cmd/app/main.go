package main

import (
	"log"
	"net/http"

	"git.rz.tu-bs.de/i.zacarias/location-api/internal/config"
	"git.rz.tu-bs.de/i.zacarias/location-api/internal/controllers"
	"github.com/gin-gonic/gin"
)

const API_NAME = "location"
const API_VER = "v2"
const API_Q_PREFIX = API_NAME + "/" + API_VER + "/queries"

// Main
func main() {
	// reads configuration
	var config, err = config.LoadConfig("../../internal/config")
	if err != nil {
		log.Fatal("Cannot read configuration: ", err)
	}

	router := gin.Default()

	// Registering routes
	apiV1 := router.Group("/v1")
	apiV2queries := router.Group("/location/v2/queries")

	apiV1.GET("/", HomePageHandler)

	apiV2queries.GET("/zones", controllers.GetZones)
	apiV2queries.GET("/zones/:id", controllers.GetZoneById)

	// Run the server
	router.Run(config.Server.Host + ":" + config.Server.Port)
}

// handler for root path "/"
func HomePageHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "This is the homepage"})
}
