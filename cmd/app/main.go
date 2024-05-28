package main

import (
	"log"
	"net/http"

	"git.rz.tu-bs.de/i.zacarias/location-api/docs"
	"git.rz.tu-bs.de/i.zacarias/location-api/internal/config"
	"git.rz.tu-bs.de/i.zacarias/location-api/internal/controllers"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	// swagger embed files
)

// Main
func main() {
	// reads configuration
	var config, err = config.LoadConfig("../../internal/config")
	if err != nil {
		log.Fatal("Cannot read configuration: ", err)
	}

	router := gin.Default()

	// Registering routes
	apiV2queries := router.Group("/location/v2")
	apiV2queries.GET("/ping", PingHandler)
	apiV2queries.GET("/queries/zones", controllers.GetZones)
	apiV2queries.GET("/queries/zones/:id", controllers.GetZoneById)

	// Serving Swagger documentation
	docs.SwaggerInfo.BasePath = "/location/v2/"
	router.GET("/location/v2/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Run the server
	router.Run(config.Server.Host + ":" + config.Server.Port)
}

// @BasePath /ping
// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Produce json
// @Success 200 {string} This is the homepage
// @Router /ping [get]
func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "This is the homepage"})
}
