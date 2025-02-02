/*
 * ETSI GS MEC 013 - Location API
 *
 * The ETSI MEC ISG MEC013 Location API described using OpenAPI.
 *
 * API version: 3.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package location-api

// PeriodicEventInfo - NOTE: reportingAmount x reportingInterval shall not exceed 8639999 (99 days, 23 hours, 59 minutes and 59 seconds) for compatibility with OMA MLP and RLP.
type PeriodicEventInfo struct {

	// Number of event reports
	ReportingAmount float32 `json:"reportingAmount"`

	// Interval of event reports
	ReportingInterval float32 `json:"reportingInterval"`
}
