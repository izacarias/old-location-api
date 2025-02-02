// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * ETSI GS MEC 013 - Location API
 *
 * The ETSI MEC ISG MEC013 Location API described using OpenAPI.
 *
 * API version: 3.1.1
 */

package location-api




type AccessPointGetById200Response struct {

	AccessPointInfo AccessPointInfo `json:"accessPointInfo,omitempty"`
}

// AssertAccessPointGetById200ResponseRequired checks if the required fields are not zero-ed
func AssertAccessPointGetById200ResponseRequired(obj AccessPointGetById200Response) error {
	if err := AssertAccessPointInfoRequired(obj.AccessPointInfo); err != nil {
		return err
	}
	return nil
}

// AssertAccessPointGetById200ResponseConstraints checks if the values respects the defined constraints
func AssertAccessPointGetById200ResponseConstraints(obj AccessPointGetById200Response) error {
	if err := AssertAccessPointInfoConstraints(obj.AccessPointInfo); err != nil {
		return err
	}
	return nil
}
