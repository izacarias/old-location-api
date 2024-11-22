// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * ETSI GS MEC 013 - Location API
 *
 * The ETSI MEC ISG MEC013 Location API described using OpenAPI.
 *
 * API version: 3.1.1
 */

package location-api




type ZoneStatusNotification struct {

	Links Links `json:"_links"`

	// Identifier of an access point (e.g. ap01). Shall be included when userNumEvent related with access point or operationStatus is included.
	AccessPointId string `json:"accessPointId,omitempty"`

	// Shall be set to \"ZoneStatusNotification\".
	NotificationType string `json:"notificationType"`

	OperationStatus OperationStatus `json:"operationStatus,omitempty"`

	TimeStamp TimeStamp `json:"timeStamp,omitempty"`

	// Shall be present when ZoneStatusSubscription includes upperNumberOfUsersZoneThreshold, lowerNumberOfUsersZoneThreshold, upperNumberOfUsersAPThreshold or lowerNumberOfUsersAPThreshold, and the number of users in a zone or an access point crosses the threshold defined in the subscription: 1 = OVER_ZONE_UPPER_THD. 2 = UNDER_ZONE_LOWER_THD. 3 = OVER_AP_UPPER_THD. 4 = UNDER_AP_LOWER_THD.
	UserNumEvent int32 `json:"userNumEvent,omitempty"`

	// The identity of the zone. 
	ZoneId string `json:"zoneId"`
}

// AssertZoneStatusNotificationRequired checks if the required fields are not zero-ed
func AssertZoneStatusNotificationRequired(obj ZoneStatusNotification) error {
	elements := map[string]interface{}{
		"_links": obj.Links,
		"notificationType": obj.NotificationType,
		"zoneId": obj.ZoneId,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertLinksRequired(obj.Links); err != nil {
		return err
	}
	if err := AssertTimeStampRequired(obj.TimeStamp); err != nil {
		return err
	}
	return nil
}

// AssertZoneStatusNotificationConstraints checks if the values respects the defined constraints
func AssertZoneStatusNotificationConstraints(obj ZoneStatusNotification) error {
	if err := AssertLinksConstraints(obj.Links); err != nil {
		return err
	}
	if err := AssertTimeStampConstraints(obj.TimeStamp); err != nil {
		return err
	}
	return nil
}
