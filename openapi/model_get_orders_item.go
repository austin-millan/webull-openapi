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
// GetOrdersItem struct for GetOrdersItem
type GetOrdersItem struct {
	ComboTickerType string `json:"comboTickerType,omitempty"`
	Orders []GetOrdersItemOrders `json:"orders,omitempty"`
	OutsideRegularTradingHour bool `json:"outsideRegularTradingHour,omitempty"`
}
