// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * ETSI GS MEC 013 - Location API
 *
 * The ETSI MEC ISG MEC013 Location API described using OpenAPI.
 *
 * API version: 3.1.1
 */

package location-api




type AreaSubGet200Response struct {

	UserAreaNotification UserAreaNotification `json:"userAreaNotification,omitempty"`
}

// AssertAreaSubGet200ResponseRequired checks if the required fields are not zero-ed
func AssertAreaSubGet200ResponseRequired(obj AreaSubGet200Response) error {
	if err := AssertUserAreaNotificationRequired(obj.UserAreaNotification); err != nil {
		return err
	}
	return nil
}

// AssertAreaSubGet200ResponseConstraints checks if the values respects the defined constraints
func AssertAreaSubGet200ResponseConstraints(obj AreaSubGet200Response) error {
	if err := AssertUserAreaNotificationConstraints(obj.UserAreaNotification); err != nil {
		return err
	}
	return nil
}