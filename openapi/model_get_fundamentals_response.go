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
// GetFundamentalsResponse struct for GetFundamentalsResponse
type GetFundamentalsResponse struct {
	Analysis GetFundamentalsResponseAnalysis `json:"analysis,omitempty"`
	Forecast []GetFundamentalsResponseForecast `json:"forecast,omitempty"`
	Remind GetFundamentalsResponseRemind `json:"remind,omitempty"`
	SimpleStatement []GetFundamentalsResponseSimpleStatement `json:"simpleStatement,omitempty"`
}
