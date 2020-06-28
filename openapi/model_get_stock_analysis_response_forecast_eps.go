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
// GetStockAnalysisResponseForecastEps struct for GetStockAnalysisResponseForecastEps
type GetStockAnalysisResponseForecastEps struct {
	Id string `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
	CurrencyName string `json:"currencyName,omitempty"`
	Points []GetStockAnalysisResponseForecastEpsPoints `json:"points,omitempty"`
}
