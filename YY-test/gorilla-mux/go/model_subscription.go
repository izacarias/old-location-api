// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * ETSI GS MEC 013 - Location API
 *
 * The ETSI MEC ISG MEC013 Location API described using OpenAPI.
 *
 * API version: 3.1.1
 */

package location-api




// Subscription - 
type Subscription struct {

	// The URI referring to the subscription.
	Href string `json:"href"`

	// Type of the subscription. The string shall be set according to the \"subscriptionType\" attribute of the associated subscription data type defined in clauses 6.3.4, 6.3.5, 6.3.6, 6.3.7 6.3.8 and 6.3.9: \"UserLocationEventSubscription\" \"UserLocationPeriodicSubscription\" \"ZoneLocationEventSubscription\" \"ZoneStatusSubscription\" \"UserAreaSubscription\" \"UserDistanceSubscription\"
	SubscriptionType string `json:"subscriptionType"`
}

// AssertSubscriptionRequired checks if the required fields are not zero-ed
func AssertSubscriptionRequired(obj Subscription) error {
	elements := map[string]interface{}{
		"href": obj.Href,
		"subscriptionType": obj.SubscriptionType,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertSubscriptionConstraints checks if the values respects the defined constraints
func AssertSubscriptionConstraints(obj Subscription) error {
	return nil
}