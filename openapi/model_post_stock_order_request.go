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
// PostStockOrderRequest struct for PostStockOrderRequest
type PostStockOrderRequest struct {
	// action
	Action string `json:"action,omitempty"`
	ComboType string `json:"comboType,omitempty"`
	// limit price
	LmtPrice string `json:"lmtPrice,omitempty"`
	OrderType string `json:"orderType,omitempty"`
	OutsideRegularTradingHour bool `json:"outsideRegularTradingHour,omitempty"`
	Quantity float32 `json:"quantity,omitempty"`
	SerialId string `json:"serialId,omitempty"`
	TickerId float32 `json:"tickerId,omitempty"`
	TimeInForce string `json:"timeInForce,omitempty"`
}
