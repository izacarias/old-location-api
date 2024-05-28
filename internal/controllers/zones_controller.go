package controllers

import (
	"log"
	"net/http"

	service "git.rz.tu-bs.de/i.zacarias/location-api/internal/services"
	"github.com/gin-gonic/gin"
)

func GetZoneById(c *gin.Context) {
	id := c.Params.ByName("id")
	log.Printf("Processing GetZoneById with id %v", id)
	zone, err := service.GetZoneById(id)
	if err != nil {
		log.Printf("Could find the zone with id %v", id)
		return
	} else {
		c.JSON(http.StatusOK, zone)
	}
}

func GetZones(c *gin.Context) {
	c.JSON(http.StatusOK, "V2: List of Zones")
	log.Printf("Processing GetZones to list all zones")
}
