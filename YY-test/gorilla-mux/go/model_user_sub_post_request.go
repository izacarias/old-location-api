// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * ETSI GS MEC 013 - Location API
 *
 * The ETSI MEC ISG MEC013 Location API described using OpenAPI.
 *
 * API version: 3.1.1
 */

package location-api




type UserSubPostRequest struct {

	UserLocationEventSubscription UserLocationEventSubscription `json:"userLocationEventSubscription,omitempty"`

	UserLocationPeriodicSubscription UserLocationPeriodicSubscription `json:"userLocationPeriodicSubscription,omitempty"`
}

// AssertUserSubPostRequestRequired checks if the required fields are not zero-ed
func AssertUserSubPostRequestRequired(obj UserSubPostRequest) error {
	if err := AssertUserLocationEventSubscriptionRequired(obj.UserLocationEventSubscription); err != nil {
		return err
	}
	if err := AssertUserLocationPeriodicSubscriptionRequired(obj.UserLocationPeriodicSubscription); err != nil {
		return err
	}
	return nil
}

// AssertUserSubPostRequestConstraints checks if the values respects the defined constraints
func AssertUserSubPostRequestConstraints(obj UserSubPostRequest) error {
	if err := AssertUserLocationEventSubscriptionConstraints(obj.UserLocationEventSubscription); err != nil {
		return err
	}
	if err := AssertUserLocationPeriodicSubscriptionConstraints(obj.UserLocationPeriodicSubscription); err != nil {
		return err
	}
	return nil
}
