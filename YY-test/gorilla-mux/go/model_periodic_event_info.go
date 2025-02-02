// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * ETSI GS MEC 013 - Location API
 *
 * The ETSI MEC ISG MEC013 Location API described using OpenAPI.
 *
 * API version: 3.1.1
 */

package location-api




// PeriodicEventInfo - NOTE: reportingAmount x reportingInterval shall not exceed 8639999 (99 days, 23 hours, 59 minutes and 59 seconds) for compatibility with OMA MLP and RLP.
type PeriodicEventInfo struct {

	// Number of event reports
	ReportingAmount float32 `json:"reportingAmount"`

	// Interval of event reports
	ReportingInterval float32 `json:"reportingInterval"`
}

// AssertPeriodicEventInfoRequired checks if the required fields are not zero-ed
func AssertPeriodicEventInfoRequired(obj PeriodicEventInfo) error {
	elements := map[string]interface{}{
		"reportingAmount": obj.ReportingAmount,
		"reportingInterval": obj.ReportingInterval,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertPeriodicEventInfoConstraints checks if the values respects the defined constraints
func AssertPeriodicEventInfoConstraints(obj PeriodicEventInfo) error {
	return nil
}
