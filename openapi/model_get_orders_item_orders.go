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
// GetOrdersItemOrders struct for GetOrdersItemOrders
type GetOrdersItemOrders struct {
	Action OrderSide `json:"action,omitempty"`
	AssetType string `json:"assetType,omitempty"`
	AvgFilledPrice string `json:"avgFilledPrice,omitempty"`
	BrokerId int32 `json:"brokerId,omitempty"`
	CanCancel bool `json:"canCancel,omitempty"`
	CanModify bool `json:"canModify,omitempty"`
	CreateTime string `json:"createTime,omitempty"`
	CreateTime0 int32 `json:"createTime0,omitempty"`
	FilledQuantity string `json:"filledQuantity,omitempty"`
	FilledTime string `json:"filledTime,omitempty"`
	FilledTime0 int32 `json:"filledTime0,omitempty"`
	FilledValue string `json:"filledValue,omitempty"`
	OptionContractMultiplier string `json:"optionContractMultiplier,omitempty"`
	OptionExercisePrice string `json:"optionExercisePrice,omitempty"`
	OrderId int32 `json:"orderId,omitempty"`
	OrderType OrderType `json:"orderType,omitempty"`
	Relation string `json:"relation,omitempty"`
	RemainQuantity string `json:"remainQuantity,omitempty"`
	StatusCode string `json:"statusCode,omitempty"`
	StatusStr string `json:"statusStr,omitempty"`
	Symbol string `json:"symbol,omitempty"`
	Ticker GetOrdersItemTicker `json:"ticker,omitempty"`
	TickerId int32 `json:"tickerId,omitempty"`
	TickerPriceDefineList []GetOrdersItemTickerPriceDefineList `json:"tickerPriceDefineList,omitempty"`
	TimeInForce TimeInForce `json:"timeInForce,omitempty"`
	TotalQuantity string `json:"totalQuantity,omitempty"`
	UpdateTime string `json:"updateTime,omitempty"`
	UpdateTime0 float32 `json:"updateTime0,omitempty"`
}
