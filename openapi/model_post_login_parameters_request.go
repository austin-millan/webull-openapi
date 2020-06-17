/*
 * Webull API
 *
 * Webull API Documentation
 *
 * API version: 3.0.1
 * Contact: austin.millan@gmail.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// PostLoginParametersRequest struct for PostLoginParametersRequest
type PostLoginParametersRequest struct {
	// username for your account
	Account string `json:"account,omitempty"`
	AccountType AccountType `json:"accountType,omitempty"`
	DeviceId string `json:"deviceId,omitempty"`
	// device name
	DeviceName string `json:"deviceName,omitempty"`
	ExtInfo PostLoginParametersRequestExtInfo `json:"extInfo,omitempty"`
	Grade float32 `json:"grade,omitempty"`
	// pwd = md5(passwordHash + password)
	Pwd string `json:"pwd,omitempty"`
	RegionId float32 `json:"regionId,omitempty"`
}
