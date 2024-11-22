/*
 * ETSI GS MEC 013 - Location API
 *
 * The ETSI MEC ISG MEC013 Location API described using OpenAPI.
 *
 * API version: 3.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package location-api

type UserSubGet200Response struct {

	UserLocationEventSubscription UserLocationEventSubscription `json:"userLocationEventSubscription,omitempty"`

	UserLocationPeriodicSubscription UserLocationPeriodicSubscription `json:"userLocationPeriodicSubscription,omitempty"`
}