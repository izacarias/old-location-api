/*
 * ETSI GS MEC 013 - Location API
 *
 * The ETSI MEC ISG MEC013 Location API described using OpenAPI.
 *
 * API version: 3.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package location-api

type RetrievalStatus string

// List of RetrievalStatus
const (
	RETRIEVED RetrievalStatus = "Retrieved"
	NOT_RETRIEVED RetrievalStatus = "NotRetrieved"
	ERROR RetrievalStatus = "Error"
)