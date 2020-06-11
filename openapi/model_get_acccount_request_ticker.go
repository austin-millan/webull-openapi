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
// GetAcccountRequestTicker struct for GetAcccountRequestTicker
type GetAcccountRequestTicker struct {
	CurrencyCode string `json:"currencyCode,omitempty"`
	CurrencyId float32 `json:"currencyId,omitempty"`
	DisExchangeCode string `json:"disExchangeCode,omitempty"`
	DisSymbol string `json:"disSymbol,omitempty"`
	ExchangeCode string `json:"exchangeCode,omitempty"`
	ExchangeId float32 `json:"exchangeId,omitempty"`
	ExtType []map[string]interface{} `json:"extType,omitempty"`
	ListStatus float32 `json:"listStatus,omitempty"`
	Name string `json:"name,omitempty"`
	RegionId float32 `json:"regionId,omitempty"`
	RegionIsoCode string `json:"regionIsoCode,omitempty"`
	RegionName string `json:"regionName,omitempty"`
	SecType []float32 `json:"secType,omitempty"`
	Symbol string `json:"symbol,omitempty"`
	TickerId float32 `json:"tickerId,omitempty"`
	TinyName string `json:"tinyName,omitempty"`
	Type float32 `json:"type,omitempty"`
}
