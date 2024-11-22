// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * ETSI GS MEC 013 - Location API
 *
 * The ETSI MEC ISG MEC013 Location API described using OpenAPI.
 *
 * API version: 3.1.1
 */

package location-api




type LinkType struct {

	// URI referring to a resource.
	Href string `json:"href"`
}

// AssertLinkTypeRequired checks if the required fields are not zero-ed
func AssertLinkTypeRequired(obj LinkType) error {
	elements := map[string]interface{}{
		"href": obj.Href,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertLinkTypeConstraints checks if the values respects the defined constraints
func AssertLinkTypeConstraints(obj LinkType) error {
	return nil
}
