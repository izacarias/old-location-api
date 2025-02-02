// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * ETSI GS MEC 013 - Location API
 *
 * The ETSI MEC ISG MEC013 Location API described using OpenAPI.
 *
 * API version: 3.1.1
 */

package location-api




type UserLocationEventNotification struct {

	Links Links `json:"_links"`

	// The identity of the access point. For the events of \"ENTERING_AREA_EVENT\", it indicates the access point that the user is currently within.  For the event of \"LEAVING_AREA_EVENT\", it indicates the access point that the user used to be within. See note 2.
	AccessPointId string `json:"accessPointId,omitempty"`

	// Address of user (e.g. ‘sip’ URI, ‘tel’ URI, ‘acr’ URI).
	Address string `json:"address,omitempty"`

	CivicInfo CivicAddress `json:"civicInfo,omitempty"`

	LocationInfo LocationInfo `json:"locationInfo,omitempty"`

	// Shall be set to \"UserLocationEventNotification\".
	NotificationType string `json:"notificationType"`

	RelativeLocationInfo RelativeLocationInfo `json:"relativeLocationInfo,omitempty"`

	TimeStamp TimeStamp `json:"timeStamp,omitempty"`

	UserLocationEvent LocationEventType `json:"userLocationEvent"`

	// The identity of the zone.  For the events of \"ENTERING_AREA_EVENT\", it is the zone that the user is currently within.  For the event of \"LEAVING_AREA_EVENT\", it is the zone that the user used to be within. See note 2.
	ZoneId string `json:"zoneId,omitempty"`
}

// AssertUserLocationEventNotificationRequired checks if the required fields are not zero-ed
func AssertUserLocationEventNotificationRequired(obj UserLocationEventNotification) error {
	elements := map[string]interface{}{
		"_links": obj.Links,
		"notificationType": obj.NotificationType,
		"userLocationEvent": obj.UserLocationEvent,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertLinksRequired(obj.Links); err != nil {
		return err
	}
	if err := AssertCivicAddressRequired(obj.CivicInfo); err != nil {
		return err
	}
	if err := AssertLocationInfoRequired(obj.LocationInfo); err != nil {
		return err
	}
	if err := AssertRelativeLocationInfoRequired(obj.RelativeLocationInfo); err != nil {
		return err
	}
	if err := AssertTimeStampRequired(obj.TimeStamp); err != nil {
		return err
	}
	return nil
}

// AssertUserLocationEventNotificationConstraints checks if the values respects the defined constraints
func AssertUserLocationEventNotificationConstraints(obj UserLocationEventNotification) error {
	if err := AssertLinksConstraints(obj.Links); err != nil {
		return err
	}
	if err := AssertCivicAddressConstraints(obj.CivicInfo); err != nil {
		return err
	}
	if err := AssertLocationInfoConstraints(obj.LocationInfo); err != nil {
		return err
	}
	if err := AssertRelativeLocationInfoConstraints(obj.RelativeLocationInfo); err != nil {
		return err
	}
	if err := AssertTimeStampConstraints(obj.TimeStamp); err != nil {
		return err
	}
	return nil
}
