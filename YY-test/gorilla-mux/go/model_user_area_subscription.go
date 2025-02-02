// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * ETSI GS MEC 013 - Location API
 *
 * The ETSI MEC ISG MEC013 Location API described using OpenAPI.
 *
 * API version: 3.1.1
 */

package location-api




type UserAreaSubscription struct {

	Links Links `json:"_links,omitempty"`

	// List of the users to be monitored. 
	AddressList []string `json:"addressList"`

	AreaDefine AreaInfo `json:"areaDefine"`

	// URI exposed by the client on which to receive notifications via HTTP. See note 1.
	CallbackReference string `json:"callbackReference,omitempty"`

	// A correlator that the client can use to tag this particular resource representation during a request to create a resource on the server. See note 2.
	ClientCorrelator string `json:"clientCorrelator,omitempty"`

	ExpiryDeadline TimeStamp `json:"expiryDeadline,omitempty"`

	// List of user event values to generate notifications for (these apply to address specified). 
	LocationEventCriteria []LocationEventType `json:"locationEventCriteria,omitempty"`

	ReportingCtrl ReportingCtrl `json:"reportingCtrl,omitempty"`

	// This IE shall be set to true if a location estimate is required for each event report.
	ReportingLocationReq bool `json:"reportingLocationReq,omitempty"`

	// Set to TRUE by the service consumer to request a test notification via HTTP on the callbackReference URI, as specified in ETSI GS MEC 009 [4], clause 6.12a.
	RequestTestNotification bool `json:"requestTestNotification,omitempty"`

	// Shall be set to \"UserAreaSubscription\".
	SubscriptionType string `json:"subscriptionType"`

	// Number of meters of acceptable error.
	TrackingAccuracy float32 `json:"trackingAccuracy"`

	WebsockNotifConfig WebsockNotifConfig `json:"websockNotifConfig,omitempty"`
}

// AssertUserAreaSubscriptionRequired checks if the required fields are not zero-ed
func AssertUserAreaSubscriptionRequired(obj UserAreaSubscription) error {
	elements := map[string]interface{}{
		"addressList": obj.AddressList,
		"areaDefine": obj.AreaDefine,
		"subscriptionType": obj.SubscriptionType,
		"trackingAccuracy": obj.TrackingAccuracy,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertLinksRequired(obj.Links); err != nil {
		return err
	}
	if err := AssertAreaInfoRequired(obj.AreaDefine); err != nil {
		return err
	}
	if err := AssertTimeStampRequired(obj.ExpiryDeadline); err != nil {
		return err
	}
	if err := AssertReportingCtrlRequired(obj.ReportingCtrl); err != nil {
		return err
	}
	if err := AssertWebsockNotifConfigRequired(obj.WebsockNotifConfig); err != nil {
		return err
	}
	return nil
}

// AssertUserAreaSubscriptionConstraints checks if the values respects the defined constraints
func AssertUserAreaSubscriptionConstraints(obj UserAreaSubscription) error {
	if err := AssertLinksConstraints(obj.Links); err != nil {
		return err
	}
	if err := AssertAreaInfoConstraints(obj.AreaDefine); err != nil {
		return err
	}
	if err := AssertTimeStampConstraints(obj.ExpiryDeadline); err != nil {
		return err
	}
	if err := AssertReportingCtrlConstraints(obj.ReportingCtrl); err != nil {
		return err
	}
	if err := AssertWebsockNotifConfigConstraints(obj.WebsockNotifConfig); err != nil {
		return err
	}
	return nil
}
