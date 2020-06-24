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
// GetSecurityAccountsResponseTickerTypes struct for GetSecurityAccountsResponseTickerTypes
type GetSecurityAccountsResponseTickerTypes struct {
	RegionId float32 `json:"regionId,omitempty"`
	TickerType float32 `json:"tickerType,omitempty"`
	OrderTypes []string `json:"orderTypes,omitempty"`
}
