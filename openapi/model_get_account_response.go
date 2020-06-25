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
// GetAccountResponse struct for GetAccountResponse
type GetAccountResponse struct {
	SecAccountId int32 `json:"secAccountId,omitempty"`
	BrokerId int32 `json:"brokerId,omitempty"`
	AccountType string `json:"accountType,omitempty"`
	BrokerAccountId string `json:"brokerAccountId,omitempty"`
	Currency string `json:"currency,omitempty"`
	CurrencyId int32 `json:"currencyId,omitempty"`
	Pdt bool `json:"pdt,omitempty"`
	Professional bool `json:"professional,omitempty"`
	ShowUpgrade bool `json:"showUpgrade,omitempty"`
	TotalCost string `json:"totalCost,omitempty"`
	NetLiquidation string `json:"netLiquidation,omitempty"`
	UnrealizedProfitLoss string `json:"unrealizedProfitLoss,omitempty"`
	UnrealizedProfitLossRate string `json:"unrealizedProfitLossRate,omitempty"`
	UnrealizedProfitLossBase float32 `json:"unrealizedProfitLossBase,omitempty"`
	Warning bool `json:"warning,omitempty"`
	Banners []GetAccountResponseBanners `json:"banners,omitempty"`
	OpenIpoOrders []map[string]interface{} `json:"openIpoOrders,omitempty"`
	AccountMembers []GetAcccountRequestAccountMembers `json:"accountMembers,omitempty"`
	Positions []GetAccountResponsePositions `json:"positions,omitempty"`
	OpenOrders []map[string]interface{} `json:"openOrders,omitempty"`
	OpenOrderSize int32 `json:"openOrderSize,omitempty"`
	Deposits []map[string]interface{} `json:"deposits,omitempty"`
}