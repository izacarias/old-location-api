package main

import (
	"log"
	"net/http"

	"git.rz.tu-bs.de/i.zacarias/location-api/internal/config"
	"git.rz.tu-bs.de/i.zacarias/location-api/pkg/zones"
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

	// Create gin router
	router := gin.Default()

	// Instatiate zones Handler and provide a data store implementation
	zStore := zones.NewMemStore()
	zHandler := NewZonesHandler(zStore)

	// Add dummy date to Zones
	zStore.AddDummyData()

	// Register routes
	router.GET("/", homePage)

	// routes for Zones
	router.GET("location/v2/queries/zones", zHandler.ListZones)

	// Run the server
	router.Run(config.Server.Host + ":" + config.Server.Port)
}

// handler for root path "/"
func homePage(c *gin.Context) {
	c.String(http.StatusOK,
		"This is the home page",
	)
}

type ZonesHandler struct {
	store zoneStore
}

func NewZonesHandler(s zoneStore) *ZonesHandler {
	return &ZonesHandler{
		store: s,
	}
}

type zoneStore interface {
	List() (map[int]zones.Zone, error)
}

func (h ZonesHandler) ListZones(c *gin.Context) {
	r, err := h.store.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, r)
}

func (h ZonesHandler) GetZone(c *gin.Context) {}
