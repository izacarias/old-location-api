// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * ETSI GS MEC 013 - Location API
 *
 * The ETSI MEC ISG MEC013 Location API described using OpenAPI.
 *
 * API version: 3.1.1
 */

package location-api




type ProblemDetails struct {

	// A human-readable explanation specific to this occurrence of the problem
	Detail string `json:"detail,omitempty"`

	// A URI reference that identifies the specific occurrence of the problem
	Instance string `json:"instance,omitempty"`

	// The HTTP status code for this occurrence of the problem
	Status int32 `json:"status,omitempty"`

	// A short, human-readable summary of the problem type
	Title string `json:"title,omitempty"`

	// A URI reference according to IETF RFC 3986 that identifies the problem type
	Type string `json:"type,omitempty"`
}

// AssertProblemDetailsRequired checks if the required fields are not zero-ed
func AssertProblemDetailsRequired(obj ProblemDetails) error {
	return nil
}

// AssertProblemDetailsConstraints checks if the values respects the defined constraints
func AssertProblemDetailsConstraints(obj ProblemDetails) error {
	return nil
}