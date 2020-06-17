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
// OptionOrder struct for OptionOrder
type OptionOrder struct {
	Action OrderSide `json:"action,omitempty"`
	Quantity float32 `json:"quantity,omitempty"`
	TickerId float32 `json:"tickerId,omitempty"`
	TickerType string `json:"tickerType,omitempty"`
	TotalQuantity float32 `json:"totalQuantity,omitempty"`
}
