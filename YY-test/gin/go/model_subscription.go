/*
 * ETSI GS MEC 013 - Location API
 *
 * The ETSI MEC ISG MEC013 Location API described using OpenAPI.
 *
 * API version: 3.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package location-api

// Subscription - 
type Subscription struct {

	// The URI referring to the subscription.
	Href string `json:"href"`

	// Type of the subscription. The string shall be set according to the \"subscriptionType\" attribute of the associated subscription data type defined in clauses 6.3.4, 6.3.5, 6.3.6, 6.3.7 6.3.8 and 6.3.9: \"UserLocationEventSubscription\" \"UserLocationPeriodicSubscription\" \"ZoneLocationEventSubscription\" \"ZoneStatusSubscription\" \"UserAreaSubscription\" \"UserDistanceSubscription\"
	SubscriptionType string `json:"subscriptionType"`
}
