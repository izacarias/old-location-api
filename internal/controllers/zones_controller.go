package controllers

import (
	"fmt"
	"log"
	"net/http"

	service "git.rz.tu-bs.de/i.zacarias/location-api/internal/services"
	"github.com/gin-gonic/gin"
)

// Returns data about a specific zone
// @Summary 	Gets information about one specific zone
// @Description Gets information about one specific zone identified by its ID
// @Tags zones
// @Accept		json
// @Produce 	json
// @Param 		id		path 		string			true 	"Zone ID"
// @Success 	200 	{object} 	domains.ZoneResponse
// @Failure		404
// @Failure		500		{object}
// @Router /queries/zones/{id} [get]
func GetZoneById(c *gin.Context) {
	id := c.Params.ByName("id")
	log.Printf("Processing GetZoneById with id %v", id)
	zone, err := service.GetZoneById(id)
	zone.ResourceURL = c.Request.Host + c.Request.URL.String()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": fmt.Sprintf("Zone %v not found", id)})
		return
	} else {
		c.JSON(http.StatusOK, zone)
	}
}

// @Summary Get list of all zones
// @Schemes
// @Description
// @Tags Zones
// @Produce json
// @Success 200 {string} List of zones
// @BasePath /location/queries
// @Router /queries/zones [get]
func GetZones(c *gin.Context) {
	log.Printf("Processing GetZones to list all zones")
	zones, err := service.ListZones()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, zones)
}
