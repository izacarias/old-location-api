/*
 * ETSI GS MEC 013 - Location API
 *
 * The ETSI MEC ISG MEC013 Location API described using OpenAPI.
 *
 * API version: 3.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package location-api

import (
	"github.com/gin-gonic/gin"
)

type LocationAPI struct {
}

// Get /location/v2/queries/zones/:zoneId/accessPoints/:accessPointId
// Query information about a specific access point under a zone. 
func (api *LocationAPI) AccessPointGetById(c *gin.Context) {
	// Your handler implementation
	c.JSON(200, gin.H{"status": "OK"})
}

// Get /location/v2/queries/zones/:zoneId/accessPoints
// Query information about a specific access point or a list of access points under a zone 
func (api *LocationAPI) AccessPointsGET(c *gin.Context) {
	// Your handler implementation
	c.JSON(200, gin.H{"status": "OK"})
}

// Delete /location/v2/subscriptions/area/:subscriptionId
// Cancel a subscription 
func (api *LocationAPI) AreaSubDELETE(c *gin.Context) {
	// Your handler implementation
	c.JSON(200, gin.H{"status": "OK"})
}

// Get /location/v2/subscriptions/area/:subscriptionId
// Retrieve subscription information 
func (api *LocationAPI) AreaSubGET(c *gin.Context) {
	// Your handler implementation
	c.JSON(200, gin.H{"status": "OK"})
}

// Get /location/v2/subscriptions/area
// Retrieves information about the subscriptions for this requestor. 
func (api *LocationAPI) AreaSubListGET(c *gin.Context) {
	// Your handler implementation
	c.JSON(200, gin.H{"status": "OK"})
}

// Post /location/v2/subscriptions/area
// Creates subscription to area notifications. 
func (api *LocationAPI) AreaSubPOST(c *gin.Context) {
	// Your handler implementation
	c.JSON(200, gin.H{"status": "OK"})
}

// Put /location/v2/subscriptions/area/:subscriptionId
// Updates a subscription information 
func (api *LocationAPI) AreaSubPUT(c *gin.Context) {
	// Your handler implementation
	c.JSON(200, gin.H{"status": "OK"})
}

// Get /location/v2/queries/distance
// Query information about distance from a user to a location or between two users 
func (api *LocationAPI) DistanceGET(c *gin.Context) {
	// Your handler implementation
	c.JSON(200, gin.H{"status": "OK"})
}

// Delete /location/v2/subscriptions/distance/:subscriptionId
// Cancel a subscription 
func (api *LocationAPI) DistanceSubDELETE(c *gin.Context) {
	// Your handler implementation
	c.JSON(200, gin.H{"status": "OK"})
}

// Get /location/v2/subscriptions/distance/:subscriptionId
// Retrieve user distance subscription information 
func (api *LocationAPI) DistanceSubGET(c *gin.Context) {
	// Your handler implementation
	c.JSON(200, gin.H{"status": "OK"})
}

// Get /location/v2/subscriptions/distance
// Retrieves all active subscriptions to distance change notifications 
func (api *LocationAPI) DistanceSubListGET(c *gin.Context) {
	// Your handler implementation
	c.JSON(200, gin.H{"status": "OK"})
}

// Post /location/v2/subscriptions/distance
// Creates a subscription for distance change notification 
func (api *LocationAPI) DistanceSubPOST(c *gin.Context) {
	// Your handler implementation
	c.JSON(200, gin.H{"status": "OK"})
}

// Put /location/v2/subscriptions/distance/:subscriptionId
// Updates a user distance subscription information 
func (api *LocationAPI) DistanceSubPUT(c *gin.Context) {
	// Your handler implementation
	c.JSON(200, gin.H{"status": "OK"})
}

// Delete /location/v2/subscriptions/users/:subscriptionId
// Cancel a subscription 
func (api *LocationAPI) UserSubDELETE(c *gin.Context) {
	// Your handler implementation
	c.JSON(200, gin.H{"status": "OK"})
}

// Get /location/v2/subscriptions/users/:subscriptionId
// Retrieve subscription information 
func (api *LocationAPI) UserSubGET(c *gin.Context) {
	// Your handler implementation
	c.JSON(200, gin.H{"status": "OK"})
}

// Get /location/v2/subscriptions/users
// Retrieves information about the subscriptions for the requestor 
func (api *LocationAPI) UserSubListGET(c *gin.Context) {
	// Your handler implementation
	c.JSON(200, gin.H{"status": "OK"})
}

// Post /location/v2/subscriptions/users
// Create subscription to UE location notifications. 
func (api *LocationAPI) UserSubPOST(c *gin.Context) {
	// Your handler implementation
	c.JSON(200, gin.H{"status": "OK"})
}

// Put /location/v2/subscriptions/users/:subscriptionId
// Updates a subscription information 
func (api *LocationAPI) UserSubPUT(c *gin.Context) {
	// Your handler implementation
	c.JSON(200, gin.H{"status": "OK"})
}

// Get /location/v2/queries/users
// Query location information about a specific UE or a group of Ues 
func (api *LocationAPI) UsersGET(c *gin.Context) {
	// Your handler implementation
	c.JSON(200, gin.H{"status": "OK"})
}

// Get /location/v2/queries/zones/:zoneId
// Query information about a specific zone 
func (api *LocationAPI) ZoneGetById(c *gin.Context) {
	// Your handler implementation
	c.JSON(200, gin.H{"status": "OK"})
}

// Delete /location/v2/subscriptions/zones/:subscriptionId
// Cancel a zone subscription 
func (api *LocationAPI) ZoneSubDELETE(c *gin.Context) {
	// Your handler implementation
	c.JSON(200, gin.H{"status": "OK"})
}

// Get /location/v2/subscriptions/zones/:subscriptionId
// Retrieve zone subscription information 
func (api *LocationAPI) ZoneSubGET(c *gin.Context) {
	// Your handler implementation
	c.JSON(200, gin.H{"status": "OK"})
}

// Get /location/v2/subscriptions/zones
// Retrieves all active subscriptions to zone notifications 
func (api *LocationAPI) ZoneSubListGET(c *gin.Context) {
	// Your handler implementation
	c.JSON(200, gin.H{"status": "OK"})
}

// Post /location/v2/subscriptions/zones
// Creates a subscription to zone notifications 
func (api *LocationAPI) ZoneSubPOST(c *gin.Context) {
	// Your handler implementation
	c.JSON(200, gin.H{"status": "OK"})
}

// Put /location/v2/subscriptions/zones/:subscriptionId
// Updates a zone subscription information 
func (api *LocationAPI) ZoneSubPUT(c *gin.Context) {
	// Your handler implementation
	c.JSON(200, gin.H{"status": "OK"})
}

// Get /location/v2/queries/zones
// Query the information about one or more specific zones or a list of zones. 
func (api *LocationAPI) ZonesGET(c *gin.Context) {
	// Your handler implementation
	c.JSON(200, gin.H{"status": "OK"})
}
