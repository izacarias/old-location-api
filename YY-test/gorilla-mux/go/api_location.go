// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * ETSI GS MEC 013 - Location API
 *
 * The ETSI MEC ISG MEC013 Location API described using OpenAPI.
 *
 * API version: 3.1.1
 */

package location-api

import (
	"encoding/json"
	"net/http"
	"strings"
	"reflect"

	"github.com/gorilla/mux"
)

// LocationAPIController binds http requests to an api service and writes the service results to the http response
type LocationAPIController struct {
	service LocationAPIServicer
	errorHandler ErrorHandler
}

// LocationAPIOption for how the controller is set up.
type LocationAPIOption func(*LocationAPIController)

// WithLocationAPIErrorHandler inject ErrorHandler into controller
func WithLocationAPIErrorHandler(h ErrorHandler) LocationAPIOption {
	return func(c *LocationAPIController) {
		c.errorHandler = h
	}
}

// NewLocationAPIController creates a default api controller
func NewLocationAPIController(s LocationAPIServicer, opts ...LocationAPIOption) *LocationAPIController {
	controller := &LocationAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the LocationAPIController
func (c *LocationAPIController) Routes() Routes {
	return Routes{
		"AccessPointGetById": Route{
			strings.ToUpper("Get"),
			"/location/v2/queries/zones/{zoneId}/accessPoints/{accessPointId}",
			c.AccessPointGetById,
		},
		"AccessPointsGET": Route{
			strings.ToUpper("Get"),
			"/location/v2/queries/zones/{zoneId}/accessPoints",
			c.AccessPointsGET,
		},
		"AreaSubDELETE": Route{
			strings.ToUpper("Delete"),
			"/location/v2/subscriptions/area/{subscriptionId}",
			c.AreaSubDELETE,
		},
		"AreaSubGET": Route{
			strings.ToUpper("Get"),
			"/location/v2/subscriptions/area/{subscriptionId}",
			c.AreaSubGET,
		},
		"AreaSubListGET": Route{
			strings.ToUpper("Get"),
			"/location/v2/subscriptions/area",
			c.AreaSubListGET,
		},
		"AreaSubPOST": Route{
			strings.ToUpper("Post"),
			"/location/v2/subscriptions/area",
			c.AreaSubPOST,
		},
		"AreaSubPUT": Route{
			strings.ToUpper("Put"),
			"/location/v2/subscriptions/area/{subscriptionId}",
			c.AreaSubPUT,
		},
		"DistanceGET": Route{
			strings.ToUpper("Get"),
			"/location/v2/queries/distance",
			c.DistanceGET,
		},
		"DistanceSubDELETE": Route{
			strings.ToUpper("Delete"),
			"/location/v2/subscriptions/distance/{subscriptionId}",
			c.DistanceSubDELETE,
		},
		"DistanceSubGET": Route{
			strings.ToUpper("Get"),
			"/location/v2/subscriptions/distance/{subscriptionId}",
			c.DistanceSubGET,
		},
		"DistanceSubListGET": Route{
			strings.ToUpper("Get"),
			"/location/v2/subscriptions/distance",
			c.DistanceSubListGET,
		},
		"DistanceSubPOST": Route{
			strings.ToUpper("Post"),
			"/location/v2/subscriptions/distance",
			c.DistanceSubPOST,
		},
		"DistanceSubPUT": Route{
			strings.ToUpper("Put"),
			"/location/v2/subscriptions/distance/{subscriptionId}",
			c.DistanceSubPUT,
		},
		"UserSubDELETE": Route{
			strings.ToUpper("Delete"),
			"/location/v2/subscriptions/users/{subscriptionId}",
			c.UserSubDELETE,
		},
		"UserSubGET": Route{
			strings.ToUpper("Get"),
			"/location/v2/subscriptions/users/{subscriptionId}",
			c.UserSubGET,
		},
		"UserSubListGET": Route{
			strings.ToUpper("Get"),
			"/location/v2/subscriptions/users",
			c.UserSubListGET,
		},
		"UserSubPOST": Route{
			strings.ToUpper("Post"),
			"/location/v2/subscriptions/users",
			c.UserSubPOST,
		},
		"UserSubPUT": Route{
			strings.ToUpper("Put"),
			"/location/v2/subscriptions/users/{subscriptionId}",
			c.UserSubPUT,
		},
		"UsersGET": Route{
			strings.ToUpper("Get"),
			"/location/v2/queries/users",
			c.UsersGET,
		},
		"ZoneGetById": Route{
			strings.ToUpper("Get"),
			"/location/v2/queries/zones/{zoneId}",
			c.ZoneGetById,
		},
		"ZoneSubDELETE": Route{
			strings.ToUpper("Delete"),
			"/location/v2/subscriptions/zones/{subscriptionId}",
			c.ZoneSubDELETE,
		},
		"ZoneSubGET": Route{
			strings.ToUpper("Get"),
			"/location/v2/subscriptions/zones/{subscriptionId}",
			c.ZoneSubGET,
		},
		"ZoneSubListGET": Route{
			strings.ToUpper("Get"),
			"/location/v2/subscriptions/zones",
			c.ZoneSubListGET,
		},
		"ZoneSubPOST": Route{
			strings.ToUpper("Post"),
			"/location/v2/subscriptions/zones",
			c.ZoneSubPOST,
		},
		"ZoneSubPUT": Route{
			strings.ToUpper("Put"),
			"/location/v2/subscriptions/zones/{subscriptionId}",
			c.ZoneSubPUT,
		},
		"ZonesGET": Route{
			strings.ToUpper("Get"),
			"/location/v2/queries/zones",
			c.ZonesGET,
		},
	}
}

// AccessPointGetById - Query information about a specific access point under a zone.
func (c *LocationAPIController) AccessPointGetById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	zoneIdParam := params["zoneId"]
	if zoneIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"zoneId"}, nil)
		return
	}
	accessPointIdParam := params["accessPointId"]
	if accessPointIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"accessPointId"}, nil)
		return
	}
	result, err := c.service.AccessPointGetById(r.Context(), zoneIdParam, accessPointIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// AccessPointsGET - Query information about a specific access point or a list of access points under a zone
func (c *LocationAPIController) AccessPointsGET(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	zoneIdParam := params["zoneId"]
	if zoneIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"zoneId"}, nil)
		return
	}
	var accessPointIdParam []string
	if query.Has("accessPointId") {
		accessPointIdParam = strings.Split(query.Get("accessPointId"), ",")
	}
	result, err := c.service.AccessPointsGET(r.Context(), zoneIdParam, accessPointIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// AreaSubDELETE - Cancel a subscription
func (c *LocationAPIController) AreaSubDELETE(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	subscriptionIdParam := params["subscriptionId"]
	if subscriptionIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"subscriptionId"}, nil)
		return
	}
	result, err := c.service.AreaSubDELETE(r.Context(), subscriptionIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// AreaSubGET - Retrieve subscription information
func (c *LocationAPIController) AreaSubGET(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	subscriptionIdParam := params["subscriptionId"]
	if subscriptionIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"subscriptionId"}, nil)
		return
	}
	result, err := c.service.AreaSubGET(r.Context(), subscriptionIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// AreaSubListGET - Retrieves information about the subscriptions for this requestor.
func (c *LocationAPIController) AreaSubListGET(w http.ResponseWriter, r *http.Request) {
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	var subscriptionTypeParam string
	if query.Has("subscription_type") {
		param := query.Get("subscription_type")

		subscriptionTypeParam = param
	} else {
	}
	result, err := c.service.AreaSubListGET(r.Context(), subscriptionTypeParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// AreaSubPOST - Creates subscription to area notifications.
func (c *LocationAPIController) AreaSubPOST(w http.ResponseWriter, r *http.Request) {
	areaSubPostRequestParam := AreaSubPostRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&areaSubPostRequestParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertAreaSubPostRequestRequired(areaSubPostRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertAreaSubPostRequestConstraints(areaSubPostRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.AreaSubPOST(r.Context(), areaSubPostRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// AreaSubPUT - Updates a subscription information
func (c *LocationAPIController) AreaSubPUT(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	subscriptionIdParam := params["subscriptionId"]
	if subscriptionIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"subscriptionId"}, nil)
		return
	}
	areaNotificationPostRequestParam := AreaNotificationPostRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&areaNotificationPostRequestParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertAreaNotificationPostRequestRequired(areaNotificationPostRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertAreaNotificationPostRequestConstraints(areaNotificationPostRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.AreaSubPUT(r.Context(), subscriptionIdParam, areaNotificationPostRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// DistanceGET - Query information about distance from a user to a location or between two users
func (c *LocationAPIController) DistanceGET(w http.ResponseWriter, r *http.Request) {
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	var addressParam []string
	if query.Has("address") {
		addressParam = strings.Split(query.Get("address"), ",")
	}
	var locationParam DistanceGetLocationParameter
	if query.Has("location") {
		param := DistanceGetLocationParameter(query.Get("location"))

		locationParam = param
	} else {
	}
	result, err := c.service.DistanceGET(r.Context(), addressParam, locationParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// DistanceSubDELETE - Cancel a subscription
func (c *LocationAPIController) DistanceSubDELETE(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	subscriptionIdParam := params["subscriptionId"]
	if subscriptionIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"subscriptionId"}, nil)
		return
	}
	result, err := c.service.DistanceSubDELETE(r.Context(), subscriptionIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// DistanceSubGET - Retrieve user distance subscription information
func (c *LocationAPIController) DistanceSubGET(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	subscriptionIdParam := params["subscriptionId"]
	if subscriptionIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"subscriptionId"}, nil)
		return
	}
	result, err := c.service.DistanceSubGET(r.Context(), subscriptionIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// DistanceSubListGET - Retrieves all active subscriptions to distance change notifications
func (c *LocationAPIController) DistanceSubListGET(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.DistanceSubListGET(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// DistanceSubPOST - Creates a subscription for distance change notification
func (c *LocationAPIController) DistanceSubPOST(w http.ResponseWriter, r *http.Request) {
	distanceSubPostRequestParam := DistanceSubPostRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&distanceSubPostRequestParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertDistanceSubPostRequestRequired(distanceSubPostRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertDistanceSubPostRequestConstraints(distanceSubPostRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.DistanceSubPOST(r.Context(), distanceSubPostRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// DistanceSubPUT - Updates a user distance subscription information
func (c *LocationAPIController) DistanceSubPUT(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	subscriptionIdParam := params["subscriptionId"]
	if subscriptionIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"subscriptionId"}, nil)
		return
	}
	distanceSubPostRequestParam := DistanceSubPostRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&distanceSubPostRequestParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertDistanceSubPostRequestRequired(distanceSubPostRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertDistanceSubPostRequestConstraints(distanceSubPostRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.DistanceSubPUT(r.Context(), subscriptionIdParam, distanceSubPostRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// UserSubDELETE - Cancel a subscription
func (c *LocationAPIController) UserSubDELETE(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	subscriptionIdParam := params["subscriptionId"]
	if subscriptionIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"subscriptionId"}, nil)
		return
	}
	result, err := c.service.UserSubDELETE(r.Context(), subscriptionIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// UserSubGET - Retrieve subscription information
func (c *LocationAPIController) UserSubGET(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	subscriptionIdParam := params["subscriptionId"]
	if subscriptionIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"subscriptionId"}, nil)
		return
	}
	result, err := c.service.UserSubGET(r.Context(), subscriptionIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// UserSubListGET - Retrieves information about the subscriptions for the requestor
func (c *LocationAPIController) UserSubListGET(w http.ResponseWriter, r *http.Request) {
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	var subscriptionTypeParam string
	if query.Has("subscription_type") {
		param := query.Get("subscription_type")

		subscriptionTypeParam = param
	} else {
	}
	var addressParam string
	if query.Has("address") {
		param := query.Get("address")

		addressParam = param
	} else {
	}
	result, err := c.service.UserSubListGET(r.Context(), subscriptionTypeParam, addressParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// UserSubPOST - Create subscription to UE location notifications.
func (c *LocationAPIController) UserSubPOST(w http.ResponseWriter, r *http.Request) {
	userSubPostRequestParam := UserSubPostRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&userSubPostRequestParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertUserSubPostRequestRequired(userSubPostRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertUserSubPostRequestConstraints(userSubPostRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UserSubPOST(r.Context(), userSubPostRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// UserSubPUT - Updates a subscription information
func (c *LocationAPIController) UserSubPUT(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	subscriptionIdParam := params["subscriptionId"]
	if subscriptionIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"subscriptionId"}, nil)
		return
	}
	userSubPostRequestParam := UserSubPostRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&userSubPostRequestParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertUserSubPostRequestRequired(userSubPostRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertUserSubPostRequestConstraints(userSubPostRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UserSubPUT(r.Context(), subscriptionIdParam, userSubPostRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// UsersGET - Query location information about a specific UE or a group of Ues
func (c *LocationAPIController) UsersGET(w http.ResponseWriter, r *http.Request) {
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	var zoneIdParam []string
	if query.Has("zoneId") {
		zoneIdParam = strings.Split(query.Get("zoneId"), ",")
	}
	var accessPointIdParam []string
	if query.Has("accessPointId") {
		accessPointIdParam = strings.Split(query.Get("accessPointId"), ",")
	}
	var addressParam []string
	if query.Has("address") {
		addressParam = strings.Split(query.Get("address"), ",")
	}
	result, err := c.service.UsersGET(r.Context(), zoneIdParam, accessPointIdParam, addressParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// ZoneGetById - Query information about a specific zone
func (c *LocationAPIController) ZoneGetById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	zoneIdParam := params["zoneId"]
	if zoneIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"zoneId"}, nil)
		return
	}
	result, err := c.service.ZoneGetById(r.Context(), zoneIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// ZoneSubDELETE - Cancel a zone subscription
func (c *LocationAPIController) ZoneSubDELETE(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	subscriptionIdParam := params["subscriptionId"]
	if subscriptionIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"subscriptionId"}, nil)
		return
	}
	result, err := c.service.ZoneSubDELETE(r.Context(), subscriptionIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// ZoneSubGET - Retrieve zone subscription information
func (c *LocationAPIController) ZoneSubGET(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	subscriptionIdParam := params["subscriptionId"]
	if subscriptionIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"subscriptionId"}, nil)
		return
	}
	result, err := c.service.ZoneSubGET(r.Context(), subscriptionIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// ZoneSubListGET - Retrieves all active subscriptions to zone notifications
func (c *LocationAPIController) ZoneSubListGET(w http.ResponseWriter, r *http.Request) {
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	var subscriptionTypeParam string
	if query.Has("subscription_type") {
		param := query.Get("subscription_type")

		subscriptionTypeParam = param
	} else {
	}
	var zoneIdParam string
	if query.Has("zoneId") {
		param := query.Get("zoneId")

		zoneIdParam = param
	} else {
	}
	result, err := c.service.ZoneSubListGET(r.Context(), subscriptionTypeParam, zoneIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// ZoneSubPOST - Creates a subscription to zone notifications
func (c *LocationAPIController) ZoneSubPOST(w http.ResponseWriter, r *http.Request) {
	zoneSubPostRequestParam := ZoneSubPostRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&zoneSubPostRequestParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertZoneSubPostRequestRequired(zoneSubPostRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertZoneSubPostRequestConstraints(zoneSubPostRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.ZoneSubPOST(r.Context(), zoneSubPostRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// ZoneSubPUT - Updates a zone subscription information
func (c *LocationAPIController) ZoneSubPUT(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	subscriptionIdParam := params["subscriptionId"]
	if subscriptionIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"subscriptionId"}, nil)
		return
	}
	zoneSubPostRequestParam := ZoneSubPostRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&zoneSubPostRequestParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertZoneSubPostRequestRequired(zoneSubPostRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertZoneSubPostRequestConstraints(zoneSubPostRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.ZoneSubPUT(r.Context(), subscriptionIdParam, zoneSubPostRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// ZonesGET - Query the information about one or more specific zones or a list of zones.
func (c *LocationAPIController) ZonesGET(w http.ResponseWriter, r *http.Request) {
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	var zoneIdParam []string
	if query.Has("zoneId") {
		zoneIdParam = strings.Split(query.Get("zoneId"), ",")
	}
	result, err := c.service.ZonesGET(r.Context(), zoneIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}
