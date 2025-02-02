// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * ETSI GS MEC 013 - Location API
 *
 * The ETSI MEC ISG MEC013 Location API described using OpenAPI.
 *
 * API version: 3.1.1
 */

package location-api




type AreaSubListGet200Response struct {

	NotificationSubscriptionList NotificationSubscriptionList `json:"notificationSubscriptionList"`
}

// AssertAreaSubListGet200ResponseRequired checks if the required fields are not zero-ed
func AssertAreaSubListGet200ResponseRequired(obj AreaSubListGet200Response) error {
	elements := map[string]interface{}{
		"notificationSubscriptionList": obj.NotificationSubscriptionList,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertNotificationSubscriptionListRequired(obj.NotificationSubscriptionList); err != nil {
		return err
	}
	return nil
}

// AssertAreaSubListGet200ResponseConstraints checks if the values respects the defined constraints
func AssertAreaSubListGet200ResponseConstraints(obj AreaSubListGet200Response) error {
	if err := AssertNotificationSubscriptionListConstraints(obj.NotificationSubscriptionList); err != nil {
		return err
	}
	return nil
}
