// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * ETSI GS MEC 013 - Location API
 *
 * The ETSI MEC ISG MEC013 Location API described using OpenAPI.
 *
 * API version: 3.1.1
 */

package location-api




type TestNotification struct {

	Links Links `json:"_links"`

	// Shall be set to \"TestNotification\".
	NotificationType string `json:"notificationType"`
}

// AssertTestNotificationRequired checks if the required fields are not zero-ed
func AssertTestNotificationRequired(obj TestNotification) error {
	elements := map[string]interface{}{
		"_links": obj.Links,
		"notificationType": obj.NotificationType,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertLinksRequired(obj.Links); err != nil {
		return err
	}
	return nil
}

// AssertTestNotificationConstraints checks if the values respects the defined constraints
func AssertTestNotificationConstraints(obj TestNotification) error {
	if err := AssertLinksConstraints(obj.Links); err != nil {
		return err
	}
	return nil
}
