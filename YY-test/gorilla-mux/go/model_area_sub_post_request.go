// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * ETSI GS MEC 013 - Location API
 *
 * The ETSI MEC ISG MEC013 Location API described using OpenAPI.
 *
 * API version: 3.1.1
 */

package location-api




type AreaSubPostRequest struct {

	UserAreaSubscription UserAreaSubscription `json:"userAreaSubscription,omitempty"`
}

// AssertAreaSubPostRequestRequired checks if the required fields are not zero-ed
func AssertAreaSubPostRequestRequired(obj AreaSubPostRequest) error {
	if err := AssertUserAreaSubscriptionRequired(obj.UserAreaSubscription); err != nil {
		return err
	}
	return nil
}

// AssertAreaSubPostRequestConstraints checks if the values respects the defined constraints
func AssertAreaSubPostRequestConstraints(obj AreaSubPostRequest) error {
	if err := AssertUserAreaSubscriptionConstraints(obj.UserAreaSubscription); err != nil {
		return err
	}
	return nil
}
