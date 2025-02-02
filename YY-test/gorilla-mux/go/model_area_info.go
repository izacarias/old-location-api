// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * ETSI GS MEC 013 - Location API
 *
 * The ETSI MEC ISG MEC013 Location API described using OpenAPI.
 *
 * API version: 3.1.1
 */

package location-api




type AreaInfo struct {

	// Shall include one point if the shape is CIRCLE. Shall include 3-15 points if the shape is POLYGON.
	Points []Point `json:"points"`

	// Shall be present if the shape is CIRCLE.
	Radius int32 `json:"radius,omitempty"`

	// The shape of the area monitored: 1 = CIRCLE. 2 = POLYGON
	Shape int32 `json:"shape"`
}

// AssertAreaInfoRequired checks if the required fields are not zero-ed
func AssertAreaInfoRequired(obj AreaInfo) error {
	elements := map[string]interface{}{
		"points": obj.Points,
		"shape": obj.Shape,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.Points {
		if err := AssertPointRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertAreaInfoConstraints checks if the values respects the defined constraints
func AssertAreaInfoConstraints(obj AreaInfo) error {
	for _, el := range obj.Points {
		if err := AssertPointConstraints(el); err != nil {
			return err
		}
	}
	return nil
}
