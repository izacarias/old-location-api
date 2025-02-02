// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * ETSI GS MEC 013 - Location API
 *
 * The ETSI MEC ISG MEC013 Location API described using OpenAPI.
 *
 * API version: 3.1.1
 */

package location-api




type UsersGet200Response struct {

	UserList UserList `json:"userList,omitempty"`
}

// AssertUsersGet200ResponseRequired checks if the required fields are not zero-ed
func AssertUsersGet200ResponseRequired(obj UsersGet200Response) error {
	if err := AssertUserListRequired(obj.UserList); err != nil {
		return err
	}
	return nil
}

// AssertUsersGet200ResponseConstraints checks if the values respects the defined constraints
func AssertUsersGet200ResponseConstraints(obj UsersGet200Response) error {
	if err := AssertUserListConstraints(obj.UserList); err != nil {
		return err
	}
	return nil
}
