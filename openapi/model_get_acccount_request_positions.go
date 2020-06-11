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
// GetAcccountRequestPositions struct for GetAcccountRequestPositions
type GetAcccountRequestPositions struct {
	AssetType string `json:"assetType,omitempty"`
	BrokerId float32 `json:"brokerId,omitempty"`
	Cost string `json:"cost,omitempty"`
	CostPrice string `json:"costPrice,omitempty"`
	Currency string `json:"currency,omitempty"`
	ExchangeRate string `json:"exchangeRate,omitempty"`
	Id float32 `json:"id,omitempty"`
	LastOpenTime string `json:"lastOpenTime,omitempty"`
	LastPrice string `json:"lastPrice,omitempty"`
	Lock bool `json:"lock,omitempty"`
	MarketValue string `json:"marketValue,omitempty"`
	Position string `json:"position,omitempty"`
	PositionProportion string `json:"positionProportion,omitempty"`
	Ticker GetAcccountRequestTicker `json:"ticker,omitempty"`
	UnrealizedProfitLoss string `json:"unrealizedProfitLoss,omitempty"`
	UnrealizedProfitLossRate string `json:"unrealizedProfitLossRate,omitempty"`
	UpdatePositionTimeStamp float32 `json:"updatePositionTimeStamp,omitempty"`
}
