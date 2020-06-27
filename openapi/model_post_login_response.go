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
// PostLoginResponse struct for PostLoginResponse
type PostLoginResponse struct {
	ExtInfo PostLoginResponseExtInfo `json:"extInfo,omitempty"`
	AccessToken string `json:"accessToken,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	RefreshToken string `json:"refreshToken,omitempty"`
	TokenExpireTime string `json:"tokenExpireTime,omitempty"`
	FirstTimeOfThird bool `json:"firstTimeOfThird,omitempty"`
	RegisterAddress int32 `json:"registerAddress,omitempty"`
	Settings PostLoginResponseSettings `json:"settings,omitempty"`
}
