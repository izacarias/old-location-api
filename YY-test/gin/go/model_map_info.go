/*
 * ETSI GS MEC 013 - Location API
 *
 * The ETSI MEC ISG MEC013 Location API described using OpenAPI.
 *
 * API version: 3.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package location-api

type MapInfo struct {

	// Ancillary map information may be used to convert coordinates between different coordinate systems.
	AncillaryMapInfo map[string]interface{} `json:"ancillaryMapInfo,omitempty"`

	// Indicates the ID of the map. 
	MapId string `json:"mapId"`

	Origin Origin `json:"origin,omitempty"`
}
