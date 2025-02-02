// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * ETSI GS MEC 013 - Location API
 *
 * The ETSI MEC ISG MEC013 Location API described using OpenAPI.
 *
 * API version: 3.1.1
 */

package location-api




// Links - Hyperlink related to the resource. This shall be only included in the HTTP responses and in HTTP PUT requests.
type Links struct {

	Self LinkType `json:"self"`
}

// AssertLinksRequired checks if the required fields are not zero-ed
func AssertLinksRequired(obj Links) error {
	elements := map[string]interface{}{
		"self": obj.Self,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertLinkTypeRequired(obj.Self); err != nil {
		return err
	}
	return nil
}

// AssertLinksConstraints checks if the values respects the defined constraints
func AssertLinksConstraints(obj Links) error {
	if err := AssertLinkTypeConstraints(obj.Self); err != nil {
		return err
	}
	return nil
}
