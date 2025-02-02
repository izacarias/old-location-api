/*
 * ETSI GS MEC 013 - Location API
 *
 * The ETSI MEC ISG MEC013 Location API described using OpenAPI.
 *
 * API version: 3.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package location-api

// UserInfo - This type represents the information related to a user attached to an access point associated to the MEC host, such access point is in scope of the Location Service instance.
type UserInfo struct {

	// Address of user (e.g. 'sip' URI, 'tel' URI, 'acr' URI) currently on the access point, see note 1.
	Address string `json:"address"`

	// The identity of the access point the user is currently on, see note 1.
	AccessPointId string `json:"AccessPointId,omitempty"`

	// The identity of the zone the user is currently within, see note 1.
	ZoneId string `json:"zoneId"`

	// Self-referring URL, see note 1.
	ResourceURL string `json:"resourceURL"`

	Timestamp TimeStamp `json:"timestamp"`

	LocationInfo LocationInfo `json:"locationInfo,omitempty"`

	CivicInfo CivicAddress `json:"civicInfo,omitempty"`

	// Reserved for future use.
	AncillaryInfo string `json:"ancillaryInfo,omitempty"`

	RelativeLocationInfo RelativeLocationInfo `json:"relativeLocationInfo,omitempty"`
}
