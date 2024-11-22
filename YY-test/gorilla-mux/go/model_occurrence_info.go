// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * ETSI GS MEC 013 - Location API
 *
 * The ETSI MEC ISG MEC013 Location API described using OpenAPI.
 *
 * API version: 3.1.1
 */

package location-api


import (
	"fmt"
)


// OccurrenceInfo : The enumeration OccurrenceInfo indicates whether event reporting is one time.
type OccurrenceInfo string

// List of OccurrenceInfo
const (
	ONE_TIME_EVENT OccurrenceInfo = "ONE_TIME_EVENT"
	MULTIPLE_TIME_EVENT OccurrenceInfo = "MULTIPLE_TIME_EVENT"
)

// AllowedOccurrenceInfoEnumValues is all the allowed values of OccurrenceInfo enum
var AllowedOccurrenceInfoEnumValues = []OccurrenceInfo{
	"ONE_TIME_EVENT",
	"MULTIPLE_TIME_EVENT",
}

// validOccurrenceInfoEnumValue provides a map of OccurrenceInfos for fast verification of use input
var validOccurrenceInfoEnumValues = map[OccurrenceInfo]struct{}{
	"ONE_TIME_EVENT": {},
	"MULTIPLE_TIME_EVENT": {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OccurrenceInfo) IsValid() bool {
	_, ok := validOccurrenceInfoEnumValues[v]
	return ok
}

// NewOccurrenceInfoFromValue returns a pointer to a valid OccurrenceInfo
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOccurrenceInfoFromValue(v string) (OccurrenceInfo, error) {
	ev := OccurrenceInfo(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for OccurrenceInfo: valid values are %v", v, AllowedOccurrenceInfoEnumValues)
}



// AssertOccurrenceInfoRequired checks if the required fields are not zero-ed
func AssertOccurrenceInfoRequired(obj OccurrenceInfo) error {
	return nil
}

// AssertOccurrenceInfoConstraints checks if the values respects the defined constraints
func AssertOccurrenceInfoConstraints(obj OccurrenceInfo) error {
	return nil
}
