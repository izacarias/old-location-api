/*
 * ETSI GS MEC 013 - Location API
 *
 * The ETSI MEC ISG MEC013 Location API described using OpenAPI.
 *
 * API version: 3.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package location-api

type ReportingCtrl struct {

	// Maximum number of notifications. For no maximum, either do not include this element or specify a value of zero. Default value is 0.
	MaximumCount int32 `json:"maximumCount,omitempty"`

	// Maximum frequency (in seconds) of notifications per subscription.
	MaximumFrequency int32 `json:"maximumFrequency,omitempty"`

	// Minimum interval between reports in case frequently reporting. Unit is second.
	MinimumInterval int32 `json:"minimumInterval,omitempty"`
}
