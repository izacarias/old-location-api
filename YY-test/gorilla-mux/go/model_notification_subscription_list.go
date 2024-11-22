// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * ETSI GS MEC 013 - Location API
 *
 * The ETSI MEC ISG MEC013 Location API described using OpenAPI.
 *
 * API version: 3.1.1
 */

package location-api




// NotificationSubscriptionList - This type contains a list of subscriptions.
type NotificationSubscriptionList struct {

	ResourceURL LinkType `json:"resourceURL"`

	Subscription []Subscription `json:"subscription,omitempty"`
}

// AssertNotificationSubscriptionListRequired checks if the required fields are not zero-ed
func AssertNotificationSubscriptionListRequired(obj NotificationSubscriptionList) error {
	elements := map[string]interface{}{
		"resourceURL": obj.ResourceURL,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertLinkTypeRequired(obj.ResourceURL); err != nil {
		return err
	}
	for _, el := range obj.Subscription {
		if err := AssertSubscriptionRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertNotificationSubscriptionListConstraints checks if the values respects the defined constraints
func AssertNotificationSubscriptionListConstraints(obj NotificationSubscriptionList) error {
	if err := AssertLinkTypeConstraints(obj.ResourceURL); err != nil {
		return err
	}
	for _, el := range obj.Subscription {
		if err := AssertSubscriptionConstraints(el); err != nil {
			return err
		}
	}
	return nil
}
