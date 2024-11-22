// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * ETSI GS MEC 013 - Location API
 *
 * The ETSI MEC ISG MEC013 Location API described using OpenAPI.
 *
 * API version: 3.1.1
 */

package location-api




type LocationInfo struct {

	// Horizontal accuracy/(semi-major) uncertainty of location provided in meters, as defined in [14]. Present only if \"shape\" equals 4, 5 or 6.
	Accuracy int32 `json:"accuracy,omitempty"`

	// Altitude accuracy/uncertainty of location provided in meters, as defined in [14]. Present only if \"shape\" equals 3 or 4.
	AccuracyAltitude int32 `json:"accuracyAltitude,omitempty"`

	// Horizontal accuracy/(semi-major) uncertainty of location provided in meters, as defined in [14]. Present only if \"shape\" equals 4, 5 or 6.
	AccuracySemiMinor int32 `json:"accuracySemiMinor,omitempty"`

	// Location altitude relative to the WGS84 ellipsoid surface.
	Altitude float32 `json:"altitude,omitempty"`

	// Confidence by which the position of a target entity is known to be within the shape description, expressed as a percentage and defined in [14]. Present only if \"shape\" equals 1, 4 or 6.
	Confidence int32 `json:"confidence,omitempty"`

	// Present only if \"shape\" equals 6.
	IncludedAngle int32 `json:"includedAngle,omitempty"`

	// Present only if \"shape\" equals 6.
	InnerRadius int32 `json:"innerRadius,omitempty"`

	// Location latitude, expressed in the range -90° to +90°. Cardinality greater than one only if \"shape\" equals 7.
	Latitude []float32 `json:"latitude"`

	// Location longitude, expressed in the range -180° to +180°. Cardinality greater than one only if \"shape\" equals 7.
	Longitude []float32 `json:"longitude"`

	// Present only if \"shape\" equals 6.
	OffsetAngle int32 `json:"offsetAngle,omitempty"`

	// Angle of orientation of the major axis, expressed in the range 0° to 180°, as defined in [14]. Present only if \"shape\" equals 4 or 6.
	OrientationMajorAxis int32 `json:"orientationMajorAxis,omitempty"`

	// Shape information, as detailed in [14], associated with the reported location coordinate: 1 = Ellipsoid_Arc 2 = ellipsoid_Point 3 = ellipsoid_Point_Altitude 4 = ellipsoid_Point_Altitude_Uncert_Ellipsoid 5 = ellipsoid_Point_Uncert_Circle 6 = ellipsoid_Point_Uncert_Ellipse 7 = polygon
	Shape int32 `json:"shape"`

	// Present only if \"shape\" equals 6.
	UncertaintyRadius int32 `json:"uncertaintyRadius,omitempty"`

	Velocity Velocity `json:"velocity,omitempty"`
}

// AssertLocationInfoRequired checks if the required fields are not zero-ed
func AssertLocationInfoRequired(obj LocationInfo) error {
	elements := map[string]interface{}{
		"latitude": obj.Latitude,
		"longitude": obj.Longitude,
		"shape": obj.Shape,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertVelocityRequired(obj.Velocity); err != nil {
		return err
	}
	return nil
}

// AssertLocationInfoConstraints checks if the values respects the defined constraints
func AssertLocationInfoConstraints(obj LocationInfo) error {
	if err := AssertVelocityConstraints(obj.Velocity); err != nil {
		return err
	}
	return nil
}