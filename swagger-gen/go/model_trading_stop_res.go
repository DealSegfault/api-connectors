/*
 * Bybit API
 *
 * ## REST API for the Bybit Exchange. 
 *
 * API version: 1.0.0
 * Contact: support@bybit.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Set Trading-Stop Condition response
type TradingStopRes struct {
	Id float32 `json:"id,omitempty"`
	UserId float32 `json:"user_id,omitempty"`
	Symbol string `json:"symbol,omitempty"`
	Side string `json:"side,omitempty"`
	Size float32 `json:"size,omitempty"`
	PositionValue float64 `json:"position_value,omitempty"`
	EntryPrice float64 `json:"entry_price,omitempty"`
	RiskId float32 `json:"risk_id,omitempty"`
	AutoAddMargin float64 `json:"auto_add_margin,omitempty"`
	Leverage float64 `json:"leverage,omitempty"`
	PositionMargin float64 `json:"position_margin,omitempty"`
	LiqPrice float64 `json:"liq_price,omitempty"`
	BustPrice float64 `json:"bust_price,omitempty"`
	OccClosingFee float64 `json:"occ_closing_fee,omitempty"`
	OccFundingFee float64 `json:"occ_funding_fee,omitempty"`
	TakeProfit float64 `json:"take_profit,omitempty"`
	StopLoss float64 `json:"stop_loss,omitempty"`
	TrailingStop float64 `json:"trailing_stop,omitempty"`
	PositionStatus string `json:"position_status,omitempty"`
	DeleverageIndicator float32 `json:"deleverage_indicator,omitempty"`
	OcCalcData string `json:"oc_calc_data,omitempty"`
	OrderMargin float64 `json:"order_margin,omitempty"`
	WalletBalance float64 `json:"wallet_balance,omitempty"`
	RealisedPnl float64 `json:"realised_pnl,omitempty"`
	CumRealisedPnl float64 `json:"cum_realised_pnl,omitempty"`
	CumCommission float32 `json:"cum_commission,omitempty"`
	CrossSeq float32 `json:"cross_seq,omitempty"`
	PositionSeq float32 `json:"position_seq,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}
