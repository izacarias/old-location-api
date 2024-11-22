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
	"net/http"

	"github.com/gin-gonic/gin"
)

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name		string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method		string
	// Pattern is the pattern of the URI.
	Pattern	 	string
	// HandlerFunc is the handler function of this route.
	HandlerFunc	gin.HandlerFunc
}

// NewRouter returns a new router.
func NewRouter(handleFunctions ApiHandleFunctions) *gin.Engine {
	return NewRouterWithGinEngine(gin.Default(), handleFunctions)
}

// NewRouter add routes to existing gin engine.
func NewRouterWithGinEngine(router *gin.Engine, handleFunctions ApiHandleFunctions) *gin.Engine {
	for _, route := range getRoutes(handleFunctions) {
		if route.HandlerFunc == nil {
			route.HandlerFunc = DefaultHandleFunc
		}
		switch route.Method {
		case http.MethodGet:
			router.GET(route.Pattern, route.HandlerFunc)
		case http.MethodPost:
			router.POST(route.Pattern, route.HandlerFunc)
		case http.MethodPut:
			router.PUT(route.Pattern, route.HandlerFunc)
		case http.MethodPatch:
			router.PATCH(route.Pattern, route.HandlerFunc)
		case http.MethodDelete:
			router.DELETE(route.Pattern, route.HandlerFunc)
		}
	}

	return router
}

// Default handler for not yet implemented routes
func DefaultHandleFunc(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

type ApiHandleFunctions struct {

	// Routes for the LocationAPI part of the API
	LocationAPI LocationAPI
}

func getRoutes(handleFunctions ApiHandleFunctions) []Route {
	return []Route{ 
		{
			"AccessPointGetById",
			http.MethodGet,
			"/location/v2/queries/zones/:zoneId/accessPoints/:accessPointId",
			handleFunctions.LocationAPI.AccessPointGetById,
		},
		{
			"AccessPointsGET",
			http.MethodGet,
			"/location/v2/queries/zones/:zoneId/accessPoints",
			handleFunctions.LocationAPI.AccessPointsGET,
		},
		{
			"AreaSubDELETE",
			http.MethodDelete,
			"/location/v2/subscriptions/area/:subscriptionId",
			handleFunctions.LocationAPI.AreaSubDELETE,
		},
		{
			"AreaSubGET",
			http.MethodGet,
			"/location/v2/subscriptions/area/:subscriptionId",
			handleFunctions.LocationAPI.AreaSubGET,
		},
		{
			"AreaSubListGET",
			http.MethodGet,
			"/location/v2/subscriptions/area",
			handleFunctions.LocationAPI.AreaSubListGET,
		},
		{
			"AreaSubPOST",
			http.MethodPost,
			"/location/v2/subscriptions/area",
			handleFunctions.LocationAPI.AreaSubPOST,
		},
		{
			"AreaSubPUT",
			http.MethodPut,
			"/location/v2/subscriptions/area/:subscriptionId",
			handleFunctions.LocationAPI.AreaSubPUT,
		},
		{
			"DistanceGET",
			http.MethodGet,
			"/location/v2/queries/distance",
			handleFunctions.LocationAPI.DistanceGET,
		},
		{
			"DistanceSubDELETE",
			http.MethodDelete,
			"/location/v2/subscriptions/distance/:subscriptionId",
			handleFunctions.LocationAPI.DistanceSubDELETE,
		},
		{
			"DistanceSubGET",
			http.MethodGet,
			"/location/v2/subscriptions/distance/:subscriptionId",
			handleFunctions.LocationAPI.DistanceSubGET,
		},
		{
			"DistanceSubListGET",
			http.MethodGet,
			"/location/v2/subscriptions/distance",
			handleFunctions.LocationAPI.DistanceSubListGET,
		},
		{
			"DistanceSubPOST",
			http.MethodPost,
			"/location/v2/subscriptions/distance",
			handleFunctions.LocationAPI.DistanceSubPOST,
		},
		{
			"DistanceSubPUT",
			http.MethodPut,
			"/location/v2/subscriptions/distance/:subscriptionId",
			handleFunctions.LocationAPI.DistanceSubPUT,
		},
		{
			"UserSubDELETE",
			http.MethodDelete,
			"/location/v2/subscriptions/users/:subscriptionId",
			handleFunctions.LocationAPI.UserSubDELETE,
		},
		{
			"UserSubGET",
			http.MethodGet,
			"/location/v2/subscriptions/users/:subscriptionId",
			handleFunctions.LocationAPI.UserSubGET,
		},
		{
			"UserSubListGET",
			http.MethodGet,
			"/location/v2/subscriptions/users",
			handleFunctions.LocationAPI.UserSubListGET,
		},
		{
			"UserSubPOST",
			http.MethodPost,
			"/location/v2/subscriptions/users",
			handleFunctions.LocationAPI.UserSubPOST,
		},
		{
			"UserSubPUT",
			http.MethodPut,
			"/location/v2/subscriptions/users/:subscriptionId",
			handleFunctions.LocationAPI.UserSubPUT,
		},
		{
			"UsersGET",
			http.MethodGet,
			"/location/v2/queries/users",
			handleFunctions.LocationAPI.UsersGET,
		},
		{
			"ZoneGetById",
			http.MethodGet,
			"/location/v2/queries/zones/:zoneId",
			handleFunctions.LocationAPI.ZoneGetById,
		},
		{
			"ZoneSubDELETE",
			http.MethodDelete,
			"/location/v2/subscriptions/zones/:subscriptionId",
			handleFunctions.LocationAPI.ZoneSubDELETE,
		},
		{
			"ZoneSubGET",
			http.MethodGet,
			"/location/v2/subscriptions/zones/:subscriptionId",
			handleFunctions.LocationAPI.ZoneSubGET,
		},
		{
			"ZoneSubListGET",
			http.MethodGet,
			"/location/v2/subscriptions/zones",
			handleFunctions.LocationAPI.ZoneSubListGET,
		},
		{
			"ZoneSubPOST",
			http.MethodPost,
			"/location/v2/subscriptions/zones",
			handleFunctions.LocationAPI.ZoneSubPOST,
		},
		{
			"ZoneSubPUT",
			http.MethodPut,
			"/location/v2/subscriptions/zones/:subscriptionId",
			handleFunctions.LocationAPI.ZoneSubPUT,
		},
		{
			"ZonesGET",
			http.MethodGet,
			"/location/v2/queries/zones",
			handleFunctions.LocationAPI.ZonesGET,
		},
	}
}