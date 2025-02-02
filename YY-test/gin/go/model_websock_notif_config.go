/*
 * ETSI GS MEC 013 - Location API
 *
 * The ETSI MEC ISG MEC013 Location API described using OpenAPI.
 *
 * API version: 3.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package location-api

type WebsockNotifConfig struct {

	// Set to true by the service consumer to indicate that Websocket delivery is requested.
	RequestWebsocketUri bool `json:"requestWebsocketUri,omitempty"`

	// Set by location server to indicate to the service consumer the Websocket URI to be used for delivering notifications.
	WebsocketUri string `json:"websocketUri,omitempty"`
}
