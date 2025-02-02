// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * ETSI GS MEC 013 - Location API
 *
 * The ETSI MEC ISG MEC013 Location API described using OpenAPI.
 *
 * API version: 3.1.1
 */

package location-api




type DistanceSubGet200Response struct {

	UserDistanceSubscription UserDistanceSubscription `json:"userDistanceSubscription,omitempty"`
}

// AssertDistanceSubGet200ResponseRequired checks if the required fields are not zero-ed
func AssertDistanceSubGet200ResponseRequired(obj DistanceSubGet200Response) error {
	if err := AssertUserDistanceSubscriptionRequired(obj.UserDistanceSubscription); err != nil {
		return err
	}
	return nil
}

// AssertDistanceSubGet200ResponseConstraints checks if the values respects the defined constraints
func AssertDistanceSubGet200ResponseConstraints(obj DistanceSubGet200Response) error {
	if err := AssertUserDistanceSubscriptionConstraints(obj.UserDistanceSubscription); err != nil {
		return err
	}
	return nil
}
