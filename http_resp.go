package birdeye

// Response type definitions for Birdeye API.
// This file contains all struct definitions for API response types.

// ============================================================================
// Chain Type and Constants
// ============================================================================

type Chain string

const (
	ChainSolana    Chain = "solana"
	ChainEthereum  Chain = "ethereum"
	ChainArbitrum  Chain = "arbitrum"
	ChainAvalanche Chain = "avalanche"
	ChainBSC       Chain = "bsc"
	ChainOptimism  Chain = "optimism"
	ChainPolygon   Chain = "polygon"
	ChainBase      Chain = "base"
	ChainZksync    Chain = "zksync"
	ChainSui       Chain = "sui"
)

// ============================================================================
// Transaction Type Constants
// ============================================================================

type TxType string

const (
	TxTypeSwap   TxType = "swap"
	TxTypeAdd    TxType = "add"
	TxTypeRemove TxType = "remove"
	TxTypeBuy    TxType = "buy"
	TxTypeSell   TxType = "sell"
)

type TradeSide string

const (
	TradeSideBuy  TradeSide = "buy"
	TradeSideSell TradeSide = "sell"
)

type TypeSwap string

const (
	TypeSwapFrom TypeSwap = "from"
	TypeSwapTo   TypeSwap = "to"
)

// ============================================================================
// Time Interval Constants
// ============================================================================

type TimeInterval string

const (
	Interval1s  TimeInterval = "1s"
	Interval15s TimeInterval = "15s"
	Interval1m  TimeInterval = "1m"
	Interval3m  TimeInterval = "3m"
	Interval5m  TimeInterval = "5m"
	Interval15m TimeInterval = "15m"
	Interval30m TimeInterval = "30m"
	Interval1H  TimeInterval = "1H"
	Interval2H  TimeInterval = "2H"
	Interval4H  TimeInterval = "4H"
	Interval6H  TimeInterval = "6H"
	Interval8H  TimeInterval = "8H"
	Interval12H TimeInterval = "12H"
	Interval1D  TimeInterval = "1D"
	Interval3D  TimeInterval = "3D"
	Interval1W  TimeInterval = "1W"
	Interval1M  TimeInterval = "1M"
)

type TimeFrame string

const (
	TimeFrame1m  TimeFrame = "1m"
	TimeFrame5m  TimeFrame = "5m"
	TimeFrame30m TimeFrame = "30m"
	TimeFrame1h  TimeFrame = "1h"
	TimeFrame2h  TimeFrame = "2h"
	TimeFrame4h  TimeFrame = "4h"
	TimeFrame8h  TimeFrame = "8h"
	TimeFrame24h TimeFrame = "24h"
	TimeFrame2d  TimeFrame = "2d"
	TimeFrame3d  TimeFrame = "3d"
	TimeFrame7d  TimeFrame = "7d"
)

type MintBurnType string

const (
	MintBurnTypeMint MintBurnType = "mint"
	MintBurnTypeBurn MintBurnType = "burn"
)

type BalanceChangeType string

const (
	BalanceChangeTypeSOL BalanceChangeType = "SOL"
	BalanceChangeTypeSPL BalanceChangeType = "SPL"
)

type BalanceChangeDirection string

const (
	BalanceChangeDirectionINCR BalanceChangeDirection = "INCR"
	BalanceChangeDirectionDECR BalanceChangeDirection = "DECR"
)

// ============================================================================
// Basic Response Types
// ============================================================================

// RespTokenPrice represents the response type for token price endpoints.
//
// Example:
//
//	{
//	    "isScaledUiToken": false,
//	    "value": 0.38622452197470425,
//	    "updateUnixTime": 1745058945,
//	    "updateHumanTime": "2025-04-19T10:35:45",
//	    "priceChange24h": 1.933391934259418,
//	    "priceInNative": 0.0027761598581298626,
//	    "liquidity": 10854103.37938592
//	}
type RespTokenPrice struct {
	IsScaledUiToken bool    `json:"isScaledUiToken" bson:"isScaledUiToken"`
	Value           float64 `json:"value" bson:"value"`
	UpdateUnixTime  int64   `json:"updateUnixTime" bson:"updateUnixTime"`
	UpdateHumanTime string  `json:"updateHumanTime" bson:"updateHumanTime"`
	PriceChange24h  float64 `json:"priceChange24h" bson:"priceChange24h"`
	PriceInNative   float64 `json:"priceInNative" bson:"priceInNative"`
	Liquidity       float64 `json:"liquidity" bson:"liquidity"`
}

// ============================================================================
// Token Transaction Types
// ============================================================================

// RespTokenTradeToken represents token details in a trade response
type RespTokenTradeToken struct {
	Symbol          string   `json:"symbol" bson:"symbol"`
	Decimals        int64    `json:"decimals" bson:"decimals"`
	Address         string   `json:"address" bson:"address"`
	Amount          any      `json:"amount" bson:"amount"`
	UIAmount        float64  `json:"uiAmount" bson:"uiAmount"`
	Price           float64  `json:"price" bson:"price"`
	NearestPrice    float64  `json:"nearestPrice" bson:"nearestPrice"`
	ChangeAmount    int64    `json:"changeAmount" bson:"changeAmount"`
	UIChangeAmount  float64  `json:"uiChangeAmount" bson:"uiChangeAmount"`
	IsScaledUIToken bool     `json:"isScaledUiToken" bson:"isScaledUiToken"`
	Multiplier      *float64 `json:"multiplier" bson:"multiplier"`
}

// RespTokenTxsItem represents a single token transaction item
type RespTokenTxsItem struct {
	Quote         RespTokenTradeToken `json:"quote" bson:"quote"`
	Base          RespTokenTradeToken `json:"base" bson:"base"`
	BasePrice     float64             `json:"basePrice" bson:"basePrice"`
	QuotePrice    float64             `json:"quotePrice" bson:"quotePrice"`
	TxHash        string              `json:"txHash" bson:"txHash"`
	Source        string              `json:"source" bson:"source"`
	BlockUnixTime int64               `json:"blockUnixTime" bson:"blockUnixTime"`
	TxType        TxType              `json:"txType" bson:"txType"`
	Owner         string              `json:"owner" bson:"owner"`
	Side          TradeSide           `json:"side" bson:"side"`
	Alias         *string             `json:"alias" bson:"alias"`
	PricePair     float64             `json:"pricePair" bson:"pricePair"`
	From          RespTokenTradeToken `json:"from" bson:"from"`
	To            RespTokenTradeToken `json:"to" bson:"to"`
	TokenPrice    float64             `json:"tokenPrice" bson:"tokenPrice"`
	PoolID        string              `json:"poolId" bson:"poolId"`
}

// RespTokenTxs represents the response for token transactions
type RespTokenTxs struct {
	Items   []RespTokenTxsItem `json:"items" bson:"items"`
	HasNext bool               `json:"hasNext" bson:"hasNext"`
}

// ============================================================================
// Pair Transaction Types
// ============================================================================

// TokenTradeToken represents token trade information in a transaction
type TokenTradeToken struct {
	Symbol   string `json:"symbol" bson:"symbol"`
	Decimals int64  `json:"decimals" bson:"decimals"`
	Address  string `json:"address" bson:"address"`
	// Amount can be either int64 or string depending on the API response format.
	// Some endpoints return numeric values, others return string representations.
	// Use type assertion to convert to the desired type:
	//   - For int64: amount, ok := token.Amount.(int64)
	//   - For string: amount, ok := token.Amount.(string)
	//   - For string to int64 conversion: str, ok := token.Amount.(string); if ok { intVal, _ := strconv.ParseInt(str, 10, 64) }
	Amount          any      `json:"amount" bson:"amount"`
	Type            string   `json:"type" bson:"type"`
	TypeSwap        string   `json:"typeSwap" bson:"typeSwap"`
	UIAmount        float64  `json:"uiAmount" bson:"uiAmount"`
	Price           float64  `json:"price" bson:"price"`
	NearestPrice    float64  `json:"nearestPrice" bson:"nearestPrice"`
	ChangeAmount    int64    `json:"changeAmount" bson:"changeAmount"`
	UIChangeAmount  float64  `json:"uiChangeAmount" bson:"uiChangeAmount"`
	IsScaledUIToken bool     `json:"isScaledUiToken" bson:"isScaledUiToken"`
	Multiplier      *float64 `json:"multiplier" bson:"multiplier"`
}

// RespPairTxsItem represents a single pair transaction item
type RespPairTxsItem struct {
	TxHash        string          `json:"txHash" bson:"txHash"`
	Source        string          `json:"source" bson:"source"`
	BlockUnixTime int64           `json:"blockUnixTime" bson:"blockUnixTime"`
	TxType        TxType          `json:"txType" bson:"txType"`
	Address       string          `json:"address" bson:"address"`
	Owner         string          `json:"owner" bson:"owner"`
	From          TokenTradeToken `json:"from" bson:"from"`
	To            TokenTradeToken `json:"to" bson:"to"`
}

// RespPairTxs represents the response for pair transactions
type RespPairTxs struct {
	Items   []RespPairTxsItem `json:"items" bson:"items"`
	HasNext bool              `json:"hasNext" bson:"hasNext"`
}

// RespTokenTxsByTime represents token transactions by time
type RespTokenTxsByTime struct {
	Items   []RespTokenTxsItem `json:"items" bson:"items"`
	HasNext bool               `json:"hasNext" bson:"hasNext"`
}

// RespPairTxsByTime represents pair transactions by time
type RespPairTxsByTime struct {
	Items   []RespPairTxsItem `json:"items" bson:"items"`
	HasNext bool              `json:"hasNext" bson:"hasNext"`
}

// ============================================================================
// V3 Transaction Types
// ============================================================================

// RespAllTxsTokenV3 represents token details in an all transactions v3 response
type RespAllTxsTokenV3 struct {
	Symbol          string   `json:"symbol" bson:"symbol"`
	Address         string   `json:"address" bson:"address"`
	Decimals        int64    `json:"decimals" bson:"decimals"`
	Price           float64  `json:"price" bson:"price"`
	Amount          string   `json:"amount" bson:"amount"`
	UIAmount        float64  `json:"ui_amount" bson:"ui_amount"`
	UIChangeAmount  float64  `json:"ui_change_amount" bson:"ui_change_amount"`
	TypeSwap        TypeSwap `json:"type_swap" bson:"type_swap"`
	IsScaledUIToken bool     `json:"is_scaled_ui_token" bson:"is_scaled_ui_token"`
	Multiplier      *float64 `json:"multiplier" bson:"multiplier"`
}

// RespAllTxsItemV3 represents a single transaction in all transactions v3
type RespAllTxsItemV3 struct {
	Base                RespAllTxsTokenV3 `json:"base" bson:"base"`
	Quote               RespAllTxsTokenV3 `json:"quote" bson:"quote"`
	TxType              TxType            `json:"tx_type" bson:"tx_type"`
	TxHash              string            `json:"tx_hash" bson:"tx_hash"`
	InsIndex            int64             `json:"ins_index" bson:"ins_index"`
	InnerInsIndex       int64             `json:"inner_ins_index" bson:"inner_ins_index"`
	BlockUnixTime       int64             `json:"block_unix_time" bson:"block_unix_time"`
	BlockNumber         int64             `json:"block_number" bson:"block_number"`
	VolumeUSD           float64           `json:"volume_usd" bson:"volume_usd"`
	Volume              float64           `json:"volume" bson:"volume"`
	Owner               string            `json:"owner" bson:"owner"`
	Signers             []string          `json:"signers" bson:"signers"`
	Source              string            `json:"source" bson:"source"`
	InteractedProgramID string            `json:"interacted_program_id" bson:"interacted_program_id"`
	PoolID              string            `json:"pool_id" bson:"pool_id"`
}

// RespAllTxsV3 represents the response for all transactions v3
type RespAllTxsV3 struct {
	Items   []RespAllTxsItemV3 `json:"items" bson:"items"`
	HasNext bool               `json:"hasNext" bson:"hasNext"`
}

// TokenInfo represents token information in a V3 transaction
type TokenInfo struct {
	Symbol         string  `json:"symbol" bson:"symbol"`
	Address        string  `json:"address" bson:"address"`
	Decimals       int64   `json:"decimals" bson:"decimals"`
	Price          float64 `json:"price" bson:"price"`
	Amount         string  `json:"amount" bson:"amount"`
	UIAmount       float64 `json:"ui_amount" bson:"ui_amount"`
	UIChangeAmount float64 `json:"ui_change_amount" bson:"ui_change_amount"`
}

type LiquidityTokenInfo struct {
	Symbol   string  `json:"symbol" bson:"symbol"`
	Address  string  `json:"address" bson:"address"`
	Decimals int64   `json:"decimals" bson:"decimals"`
	Amount   string  `json:"amount" bson:"amount"`
	UIAmount float64 `json:"ui_amount" bson:"ui_amount"`
}

// RespTokenTxsItemV3 represents a single token transaction in V3 API
type RespTokenTxsItemV3 struct {
	TxType              TxType    `json:"tx_type" bson:"tx_type"`
	TxHash              string    `json:"tx_hash" bson:"tx_hash"`
	InsIndex            int64     `json:"ins_index" bson:"ins_index"`
	InnerInsIndex       int64     `json:"inner_ins_index" bson:"inner_ins_index"`
	BlockUnixTime       int64     `json:"block_unix_time" bson:"block_unix_time"`
	BlockNumber         int64     `json:"block_number" bson:"block_number"`
	VolumeUSD           float64   `json:"volume_usd" bson:"volume_usd"`
	Volume              float64   `json:"volume" bson:"volume"`
	Owner               string    `json:"owner" bson:"owner"`
	Signers             []string  `json:"signers" bson:"signers"`
	Source              string    `json:"source" bson:"source"`
	Side                TradeSide `json:"side" bson:"side"`
	InteractedProgramID string    `json:"interacted_program_id" bson:"interacted_program_id"`
	Alias               *string   `json:"alias" bson:"alias"`
	PricePair           float64   `json:"price_pair" bson:"price_pair"`
	// From just for buy and sell tx type
	From TokenInfo `json:"from" bson:"from"`
	// To just for buy and sell tx type
	To TokenInfo `json:"to" bson:"to"`
	// Tokens just for add and remove tx type
	Tokens []LiquidityTokenInfo `json:"tokens" bson:"tokens"`
	PoolID string               `json:"pool_id" bson:"pool_id"`
}

// RespTokenTxsV3 represents the response for token transactions V3
type RespTokenTxsV3 struct {
	Items   []RespTokenTxsItemV3 `json:"items" bson:"items"`
	HasNext bool                 `json:"has_next" bson:"has_next"`
}

// RespRecentTxsTokenV3 represents token details in a recent transactions v3 response
type RespRecentTxsTokenV3 struct {
	Symbol          string   `json:"symbol" bson:"symbol"`
	Address         string   `json:"address" bson:"address"`
	Decimals        int64    `json:"decimals" bson:"decimals"`
	Price           float64  `json:"price" bson:"price"`
	Amount          string   `json:"amount" bson:"amount"`
	UIAmount        float64  `json:"ui_amount" bson:"ui_amount"`
	UIChangeAmount  float64  `json:"ui_change_amount" bson:"ui_change_amount"`
	TypeSwap        TypeSwap `json:"type_swap" bson:"type_swap"`
	IsScaledUIToken bool     `json:"is_scaled_ui_token" bson:"is_scaled_ui_token"`
	Multiplier      *float64 `json:"multiplier" bson:"multiplier"`
}

// RespRecentTxsItemV3 represents a single recent transaction in V3
type RespRecentTxsItemV3 struct {
	Base                RespRecentTxsTokenV3 `json:"base" bson:"base"`
	Quote               RespRecentTxsTokenV3 `json:"quote" bson:"quote"`
	TxType              TxType               `json:"tx_type" bson:"tx_type"`
	TxHash              string               `json:"tx_hash" bson:"tx_hash"`
	InsIndex            int64                `json:"ins_index" bson:"ins_index"`
	InnerInsIndex       int64                `json:"inner_ins_index" bson:"inner_ins_index"`
	BlockUnixTime       int64                `json:"block_unix_time" bson:"block_unix_time"`
	BlockNumber         int64                `json:"block_number" bson:"block_number"`
	VolumeUSD           float64              `json:"volume_usd" bson:"volume_usd"`
	Volume              float64              `json:"volume" bson:"volume"`
	Owner               string               `json:"owner" bson:"owner"`
	Signers             []string             `json:"signers" bson:"signers"`
	Source              string               `json:"source" bson:"source"`
	InteractedProgramID string               `json:"interacted_program_id" bson:"interacted_program_id"`
	PoolID              string               `json:"pool_id" bson:"pool_id"`
}

// RespRecentTxsV3 represents the response for recent transactions V3
type RespRecentTxsV3 struct {
	Items   []RespRecentTxsItemV3 `json:"items" bson:"items"`
	HasNext bool                  `json:"hasNext" bson:"hasNext"`
}

// ============================================================================
// OHLCV Types
// ============================================================================

// RespTokenOHLCVItem represents an OHLCV data point for a token
type RespTokenOHLCVItem struct {
	O        float64      `json:"o" bson:"o"`
	H        float64      `json:"h" bson:"h"`
	L        float64      `json:"l" bson:"l"`
	C        float64      `json:"c" bson:"c"`
	V        float64      `json:"v" bson:"v"`
	UnixTime int64        `json:"unixTime" bson:"unixTime"`
	Address  string       `json:"address" bson:"address"`
	Type     TimeInterval `json:"type" bson:"type"`
	Currency string       `json:"currency" bson:"currency"`
}

// RespTokenOHLCVs represents OHLCV response for a token
type RespTokenOHLCVs struct {
	IsScaledUIToken bool                 `json:"isScaledUiToken" bson:"isScaledUiToken"`
	Items           []RespTokenOHLCVItem `json:"items" bson:"items"`
}

// RespPairOHLCVItem represents an OHLCV data point for a trading pair
type RespPairOHLCVItem struct {
	Address  string       `json:"address" bson:"address"`
	C        float64      `json:"c" bson:"c"`
	H        float64      `json:"h" bson:"h"`
	L        float64      `json:"l" bson:"l"`
	O        float64      `json:"o" bson:"o"`
	Type     TimeInterval `json:"type" bson:"type"`
	UnixTime int64        `json:"unixTime" bson:"unixTime"`
	V        float64      `json:"v" bson:"v"`
}

// RespOHLCVBaseQuoteItem represents an OHLCV data point for base/quote token pair
type RespOHLCVBaseQuoteItem struct {
	O        float64 `json:"o" bson:"o"`
	C        float64 `json:"c" bson:"c"`
	H        float64 `json:"h" bson:"h"`
	L        float64 `json:"l" bson:"l"`
	VBase    float64 `json:"vBase" bson:"vBase"`
	VQuote   float64 `json:"vQuote" bson:"vQuote"`
	UnixTime int64   `json:"unixTime" bson:"unixTime"`
}

// RespOHLCVBaseQuote represents OHLCV response for a base/quote token pair
type RespOHLCVBaseQuote struct {
	Items                []RespOHLCVBaseQuoteItem `json:"items" bson:"items"`
	BaseAddress          string                   `json:"baseAddress" bson:"baseAddress"`
	QuoteAddress         string                   `json:"quoteAddress" bson:"quoteAddress"`
	IsScaledUITokenBase  bool                     `json:"isScaledUiTokenBase" bson:"isScaledUiTokenBase"`
	IsScaledUITokenQuote bool                     `json:"isScaledUiTokenQuote" bson:"isScaledUiTokenQuote"`
	Type                 TimeInterval             `json:"type" bson:"type"`
}

// RespTokenOHLCVItemV3 represents an OHLCV V3 data point for a token
type RespTokenOHLCVItemV3 struct {
	O        float64      `json:"o" bson:"o"`
	H        float64      `json:"h" bson:"h"`
	L        float64      `json:"l" bson:"l"`
	C        float64      `json:"c" bson:"c"`
	V        float64      `json:"v" bson:"v"`
	VUSD     float64      `json:"v_usd" bson:"v_usd"`
	UnixTime int64        `json:"unix_time" bson:"unix_time"`
	Address  string       `json:"address" bson:"address"`
	Type     TimeInterval `json:"type" bson:"type"`
	Currency string       `json:"currency" bson:"currency"`
}

// RespTokenOHLCVsV3 represents OHLCV V3 response for a token
type RespTokenOHLCVsV3 struct {
	IsScaledUIToken bool                   `json:"is_scaled_ui_token" bson:"is_scaled_ui_token"`
	Items           []RespTokenOHLCVItemV3 `json:"items" bson:"items"`
}

// RespPairOHLCVItemV3 represents an OHLCV V3 data point for a trading pair
type RespPairOHLCVItemV3 struct {
	Address  string       `json:"address" bson:"address"`
	H        float64      `json:"h" bson:"h"`
	O        float64      `json:"o" bson:"o"`
	L        float64      `json:"l" bson:"l"`
	C        float64      `json:"c" bson:"c"`
	Type     TimeInterval `json:"type" bson:"type"`
	V        float64      `json:"v" bson:"v"`
	UnixTime int64        `json:"unix_time" bson:"unix_time"`
	VUSD     float64      `json:"v_usd" bson:"v_usd"`
}

// ============================================================================
// Price History and Statistics Types
// ============================================================================

// RespPriceHistoryItem represents an individual price history data point
type RespPriceHistoryItem struct {
	UnixTime int64   `json:"unixTime" bson:"unixTime"`
	Value    float64 `json:"value" bson:"value"`
}

// RespTokenPriceHistories represents the response for token price history
type RespTokenPriceHistories struct {
	IsScaledUIToken bool                   `json:"isScaledUiToken" bson:"isScaledUiToken"`
	Items           []RespPriceHistoryItem `json:"items" bson:"items"`
}

// RespTokenPriceHistoryByTime represents price history for a token at a specific time
type RespTokenPriceHistoryByTime struct {
	IsScaledUIToken bool    `json:"isScaledUiToken" bson:"isScaledUiToken"`
	Value           float64 `json:"value" bson:"value"`
	UpdateUnixTime  int64   `json:"updateUnixTime" bson:"updateUnixTime"`
	PriceChange24h  float64 `json:"priceChange24h" bson:"priceChange24h"`
}

// RespTokenPriceVolume represents token price and volume data
type RespTokenPriceVolume struct {
	IsScaledUIToken     bool    `json:"isScaledUiToken" bson:"isScaledUiToken"`
	Price               float64 `json:"price" bson:"price"`
	UpdateUnixTime      int64   `json:"updateUnixTime" bson:"updateUnixTime"`
	UpdateHumanTime     string  `json:"updateHumanTime" bson:"updateHumanTime"`
	VolumeUSD           float64 `json:"volumeUSD" bson:"volumeUSD"`
	VolumeChangePercent float64 `json:"volumeChangePercent" bson:"volumeChangePercent"`
	PriceChangePercent  float64 `json:"priceChangePercent" bson:"priceChangePercent"`
}

// PriceStatsData represents price statistics data point
type PriceStatsData struct {
	UnixTimeUpdatePrice int64     `json:"unix_time_update_price" bson:"unix_time_update_price"`
	TimeFrame           TimeFrame `json:"time_frame" bson:"time_frame"`
	Price               float64   `json:"price" bson:"price"`
	PriceChangePercent  float64   `json:"price_change_percent" bson:"price_change_percent"`
	High                float64   `json:"high" bson:"high"`
	Low                 float64   `json:"low" bson:"low"`
}

// RespTokenPriceStats represents price statistics response for a token
type RespTokenPriceStats struct {
	Address         string           `json:"address" bson:"address"`
	IsScaledUIToken bool             `json:"is_scaled_ui_token" bson:"is_scaled_ui_token"`
	Data            []PriceStatsData `json:"data" bson:"data"`
}

type RespMultiTokenPriceStats = []RespTokenPriceStats

// ============================================================================
// Pair Overview Types
// ============================================================================

// TokenInfoInPair represents token information in pair overview
type TokenInfoInPair struct {
	Address         string  `json:"address" bson:"address"`
	Decimals        int64   `json:"decimals" bson:"decimals"`
	Icon            string  `json:"icon" bson:"icon"`
	Symbol          string  `json:"symbol" bson:"symbol"`
	IsScaledUIToken bool    `json:"is_scaled_ui_token" bson:"is_scaled_ui_token"`
	Multiplier      float64 `json:"multiplier" bson:"multiplier"`
}

// RespPairOverview represents overview data for a trading pair
type RespPairOverview struct {
	Address                      string          `json:"address" bson:"address"`
	Base                         TokenInfoInPair `json:"base" bson:"base"`
	Quote                        TokenInfoInPair `json:"quote" bson:"quote"`
	Name                         string          `json:"name" bson:"name"`
	Source                       string          `json:"source" bson:"source"`
	CreatedAt                    string          `json:"created_at" bson:"created_at"`
	Liquidity                    float64         `json:"liquidity" bson:"liquidity"`
	LiquidityChangePercentage24h *float64        `json:"liquidity_change_percentage_24h" bson:"liquidity_change_percentage_24h"`
	Price                        float64         `json:"price" bson:"price"`
	Trade24h                     int64           `json:"trade_24h" bson:"trade_24h"`
	Trade12h                     int64           `json:"trade_12h" bson:"trade_12h"`
	Trade8h                      int64           `json:"trade_8h" bson:"trade_8h"`
	Trade4h                      int64           `json:"trade_4h" bson:"trade_4h"`
	Trade2h                      int64           `json:"trade_2h" bson:"trade_2h"`
	Trade1h                      int64           `json:"trade_1h" bson:"trade_1h"`
	Trade30m                     int64           `json:"trade_30m" bson:"trade_30m"`
	Trade24hChangePercent        float64         `json:"trade_24h_change_percent" bson:"trade_24h_change_percent"`
	Trade12hChangePercent        float64         `json:"trade_12h_change_percent" bson:"trade_12h_change_percent"`
	Trade8hChangePercent         float64         `json:"trade_8h_change_percent" bson:"trade_8h_change_percent"`
	Trade4hChangePercent         float64         `json:"trade_4h_change_percent" bson:"trade_4h_change_percent"`
	Trade2hChangePercent         float64         `json:"trade_2h_change_percent" bson:"trade_2h_change_percent"`
	Trade1hChangePercent         float64         `json:"trade_1h_change_percent" bson:"trade_1h_change_percent"`
	Trade30mChangePercent        float64         `json:"trade_30m_change_percent" bson:"trade_30m_change_percent"`
	TradeHistory24h              int64           `json:"trade_history_24h" bson:"trade_history_24h"`
	TradeHistory12h              int64           `json:"trade_history_12h" bson:"trade_history_12h"`
	TradeHistory8h               int64           `json:"trade_history_8h" bson:"trade_history_8h"`
	TradeHistory4h               int64           `json:"trade_history_4h" bson:"trade_history_4h"`
	TradeHistory2h               int64           `json:"trade_history_2h" bson:"trade_history_2h"`
	TradeHistory1h               int64           `json:"trade_history_1h" bson:"trade_history_1h"`
	TradeHistory30m              int64           `json:"trade_history_30m" bson:"trade_history_30m"`
	UniqueWallet24h              int64           `json:"unique_wallet_24h" bson:"unique_wallet_24h"`
	UniqueWallet12h              int64           `json:"unique_wallet_12h" bson:"unique_wallet_12h"`
	UniqueWallet8h               int64           `json:"unique_wallet_8h" bson:"unique_wallet_8h"`
	UniqueWallet4h               int64           `json:"unique_wallet_4h" bson:"unique_wallet_4h"`
	UniqueWallet2h               int64           `json:"unique_wallet_2h" bson:"unique_wallet_2h"`
	UniqueWallet1h               int64           `json:"unique_wallet_1h" bson:"unique_wallet_1h"`
	UniqueWallet30m              int64           `json:"unique_wallet_30m" bson:"unique_wallet_30m"`
	UniqueWallet24hChangePercent float64         `json:"unique_wallet_24h_change_percent" bson:"unique_wallet_24h_change_percent"`
	UniqueWallet12hChangePercent float64         `json:"unique_wallet_12h_change_percent" bson:"unique_wallet_12h_change_percent"`
	UniqueWallet8hChangePercent  float64         `json:"unique_wallet_8h_change_percent" bson:"unique_wallet_8h_change_percent"`
	UniqueWallet4hChangePercent  float64         `json:"unique_wallet_4h_change_percent" bson:"unique_wallet_4h_change_percent"`
	UniqueWallet2hChangePercent  float64         `json:"unique_wallet_2h_change_percent" bson:"unique_wallet_2h_change_percent"`
	UniqueWallet1hChangePercent  float64         `json:"unique_wallet_1h_change_percent" bson:"unique_wallet_1h_change_percent"`
	UniqueWallet30mChangePercent float64         `json:"unique_wallet_30m_change_percent" bson:"unique_wallet_30m_change_percent"`
	Volume24h                    float64         `json:"volume_24h" bson:"volume_24h"`
	Volume12h                    float64         `json:"volume_12h" bson:"volume_12h"`
	Volume8h                     float64         `json:"volume_8h" bson:"volume_8h"`
	Volume4h                     float64         `json:"volume_4h" bson:"volume_4h"`
	Volume2h                     float64         `json:"volume_2h" bson:"volume_2h"`
	Volume1h                     float64         `json:"volume_1h" bson:"volume_1h"`
	Volume30m                    float64         `json:"volume_30m" bson:"volume_30m"`
	Volume24hBase                float64         `json:"volume_24h_base" bson:"volume_24h_base"`
	Volume12hBase                float64         `json:"volume_12h_base" bson:"volume_12h_base"`
	Volume8hBase                 float64         `json:"volume_8h_base" bson:"volume_8h_base"`
	Volume4hBase                 float64         `json:"volume_4h_base" bson:"volume_4h_base"`
	Volume2hBase                 float64         `json:"volume_2h_base" bson:"volume_2h_base"`
	Volume1hBase                 float64         `json:"volume_1h_base" bson:"volume_1h_base"`
	Volume30mBase                float64         `json:"volume_30m_base" bson:"volume_30m_base"`
	Volume24hQuote               float64         `json:"volume_24h_quote" bson:"volume_24h_quote"`
	Volume12hQuote               float64         `json:"volume_12h_quote" bson:"volume_12h_quote"`
	Volume8hQuote                float64         `json:"volume_8h_quote" bson:"volume_8h_quote"`
	Volume4hQuote                float64         `json:"volume_4h_quote" bson:"volume_4h_quote"`
	Volume2hQuote                float64         `json:"volume_2h_quote" bson:"volume_2h_quote"`
	Volume1hQuote                float64         `json:"volume_1h_quote" bson:"volume_1h_quote"`
	Volume30mQuote               float64         `json:"volume_30m_quote" bson:"volume_30m_quote"`
	Volume24hChangePercentage24h *float64        `json:"volume_24h_change_percentage_24h" bson:"volume_24h_change_percentage_24h"`
}

// ============================================================================
// Token List Types
// ============================================================================

// RespTokenListV1Token represents token details in a token list v1 response
type RespTokenListV1Token struct {
	IsScaledUIToken   bool     `json:"isScaledUiToken" bson:"isScaledUiToken"`
	Multiplier        *float64 `json:"multiplier" bson:"multiplier"`
	Address           string   `json:"address" bson:"address"`
	Decimals          int64    `json:"decimals" bson:"decimals"`
	Price             float64  `json:"price" bson:"price"`
	LastTradeUnixTime int64    `json:"lastTradeUnixTime" bson:"lastTradeUnixTime"`
	Liquidity         float64  `json:"liquidity" bson:"liquidity"`
	LogoURI           string   `json:"logoURI" bson:"logoURI"`
	MC                float64  `json:"mc" bson:"mc"`
	Name              string   `json:"name" bson:"name"`
	Symbol            string   `json:"symbol" bson:"symbol"`
	V24hChangePercent float64  `json:"v24hChangePercent" bson:"v24hChangePercent"`
	V24hUSD           float64  `json:"v24hUSD" bson:"v24hUSD"`
}

// RespTokenListV1 represents token list v1 response
type RespTokenListV1 struct {
	UpdateUnixTime int64                  `json:"updateUnixTime" bson:"updateUnixTime"`
	UpdateTime     string                 `json:"updateTime" bson:"updateTime"`
	Tokens         []RespTokenListV1Token `json:"tokens" bson:"tokens"`
	Total          int64                  `json:"total" bson:"total"`
}

// TokenExtensions represents token extension metadata
type TokenExtensions struct {
	CoingeckoID *string `json:"coingecko_id,omitempty" bson:"coingecko_id,omitempty"`
	SerumV3USDC *string `json:"serum_v3_usdc,omitempty" bson:"serum_v3_usdc,omitempty"`
	SerumV3USDT *string `json:"serum_v3_usdt,omitempty" bson:"serum_v3_usdt,omitempty"`
	Website     *string `json:"website,omitempty" bson:"website,omitempty"`
	Telegram    *string `json:"telegram,omitempty" bson:"telegram,omitempty"`
	Twitter     *string `json:"twitter,omitempty" bson:"twitter,omitempty"`
	Description *string `json:"description,omitempty" bson:"description,omitempty"`
	Discord     *string `json:"discord,omitempty" bson:"discord,omitempty"`
	Medium      *string `json:"medium,omitempty" bson:"medium,omitempty"`
}

// RespTokenListV3TokenItem represents individual token details in V3
type RespTokenListV3TokenItem struct {
	Address                      string          `json:"address" bson:"address"`
	LogoURI                      string          `json:"logo_uri" bson:"logo_uri"`
	Name                         string          `json:"name" bson:"name"`
	Symbol                       string          `json:"symbol" bson:"symbol"`
	Decimals                     int64           `json:"decimals" bson:"decimals"`
	Extensions                   TokenExtensions `json:"extensions" bson:"extensions"`
	MarketCap                    float64         `json:"market_cap" bson:"market_cap"`
	FDV                          float64         `json:"fdv" bson:"fdv"`
	TotalSupply                  float64         `json:"total_supply" bson:"total_supply"`
	CirculatingSupply            float64         `json:"circulating_supply" bson:"circulating_supply"`
	Liquidity                    float64         `json:"liquidity" bson:"liquidity"`
	LastTradeUnixTime            int64           `json:"last_trade_unix_time" bson:"last_trade_unix_time"`
	Volume1hUSD                  float64         `json:"volume_1h_usd" bson:"volume_1h_usd"`
	Volume1hChangePercent        float64         `json:"volume_1h_change_percent" bson:"volume_1h_change_percent"`
	Volume2hUSD                  float64         `json:"volume_2h_usd" bson:"volume_2h_usd"`
	Volume2hChangePercent        float64         `json:"volume_2h_change_percent" bson:"volume_2h_change_percent"`
	Volume4hUSD                  float64         `json:"volume_4h_usd" bson:"volume_4h_usd"`
	Volume4hChangePercent        float64         `json:"volume_4h_change_percent" bson:"volume_4h_change_percent"`
	Volume8hUSD                  float64         `json:"volume_8h_usd" bson:"volume_8h_usd"`
	Volume8hChangePercent        float64         `json:"volume_8h_change_percent" bson:"volume_8h_change_percent"`
	Volume24hUSD                 float64         `json:"volume_24h_usd" bson:"volume_24h_usd"`
	Volume24hChangePercent       float64         `json:"volume_24h_change_percent" bson:"volume_24h_change_percent"`
	Trade1hCount                 int64           `json:"trade_1h_count" bson:"trade_1h_count"`
	Trade2hCount                 int64           `json:"trade_2h_count" bson:"trade_2h_count"`
	Trade4hCount                 int64           `json:"trade_4h_count" bson:"trade_4h_count"`
	Trade8hCount                 int64           `json:"trade_8h_count" bson:"trade_8h_count"`
	Trade24hCount                int64           `json:"trade_24h_count" bson:"trade_24h_count"`
	Buy24h                       int64           `json:"buy_24h" bson:"buy_24h"`
	Buy24hChangePercent          float64         `json:"buy_24h_change_percent" bson:"buy_24h_change_percent"`
	VolumeBuy24hUSD              float64         `json:"volume_buy_24h_usd" bson:"volume_buy_24h_usd"`
	VolumeBuy24hChangePercent    float64         `json:"volume_buy_24h_change_percent" bson:"volume_buy_24h_change_percent"`
	Sell24h                      int64           `json:"sell_24h" bson:"sell_24h"`
	Sell24hChangePercent         float64         `json:"sell_24h_change_percent" bson:"sell_24h_change_percent"`
	VolumeSell24hUSD             float64         `json:"volume_sell_24h_usd" bson:"volume_sell_24h_usd"`
	VolumeSell24hChangePercent   float64         `json:"volume_sell_24h_change_percent" bson:"volume_sell_24h_change_percent"`
	UniqueWallet24h              int64           `json:"unique_wallet_24h" bson:"unique_wallet_24h"`
	UniqueWallet24hChangePercent float64         `json:"unique_wallet_24h_change_percent" bson:"unique_wallet_24h_change_percent"`
	Price                        float64         `json:"price" bson:"price"`
	PriceChange1hPercent         float64         `json:"price_change_1h_percent" bson:"price_change_1h_percent"`
	PriceChange2hPercent         float64         `json:"price_change_2h_percent" bson:"price_change_2h_percent"`
	PriceChange4hPercent         float64         `json:"price_change_4h_percent" bson:"price_change_4h_percent"`
	PriceChange8hPercent         float64         `json:"price_change_8h_percent" bson:"price_change_8h_percent"`
	PriceChange24hPercent        float64         `json:"price_change_24h_percent" bson:"price_change_24h_percent"`
	Holder                       int64           `json:"holder" bson:"holder"`
	RecentListingTime            *int64          `json:"recent_listing_time" bson:"recent_listing_time"`
	IsScaledUIToken              bool            `json:"is_scaled_ui_token" bson:"is_scaled_ui_token"`
	Multiplier                   *float64        `json:"multiplier" bson:"multiplier"`
}

// RespTokenListV3 represents token list V3 response
type RespTokenListV3 struct {
	Items   []RespTokenListV3TokenItem `json:"items" bson:"items"`
	HasNext bool                       `json:"hasNext" bson:"hasNext"`
}

// RespTokenListV3Scroll represents token list v3 scroll response
type RespTokenListV3Scroll struct {
	Items   []RespTokenListV3TokenItem `json:"items" bson:"items"`
	HasNext bool                       `json:"hasNext" bson:"hasNext"`
}

// ============================================================================
// Token Overview Type
// ============================================================================

// RespTokenOverview represents token overview response including price, volume and trading statistics
type RespTokenOverview struct {
	Address            string          `json:"address" bson:"address"`
	Decimals           int64           `json:"decimals" bson:"decimals"`
	Symbol             string          `json:"symbol" bson:"symbol"`
	Name               string          `json:"name" bson:"name"`
	MarketCap          float64         `json:"marketCap" bson:"marketCap"`
	FDV                float64         `json:"fdv" bson:"fdv"`
	Extensions         TokenExtensions `json:"extensions" bson:"extensions"`
	LogoURI            string          `json:"logoURI" bson:"logoURI"`
	Liquidity          float64         `json:"liquidity" bson:"liquidity"`
	LastTradeUnixTime  int64           `json:"lastTradeUnixTime" bson:"lastTradeUnixTime"`
	LastTradeHumanTime string          `json:"lastTradeHumanTime" bson:"lastTradeHumanTime"`
	Price              float64         `json:"price" bson:"price"`
	History1mPrice     float64         `json:"history1mPrice" bson:"history1mPrice"`
	History5mPrice     float64         `json:"history5mPrice" bson:"history5mPrice"`
	History30mPrice    float64         `json:"history30mPrice" bson:"history30mPrice"`
	History1hPrice     float64         `json:"history1hPrice" bson:"history1hPrice"`
	History2hPrice     float64         `json:"history2hPrice" bson:"history2hPrice"`
	History4hPrice     float64         `json:"history4hPrice" bson:"history4hPrice"`
	History6hPrice     float64         `json:"history6hPrice" bson:"history6hPrice"`
	History8hPrice     float64         `json:"history8hPrice" bson:"history8hPrice"`
	History12hPrice    float64         `json:"history12hPrice" bson:"history12hPrice"`
	History24hPrice    float64         `json:"history24hPrice" bson:"history24hPrice"`

	PriceChange1mPercent  float64 `json:"priceChange1mPercent" bson:"priceChange1mPercent"`
	PriceChange5mPercent  float64 `json:"priceChange5mPercent" bson:"priceChange5mPercent"`
	PriceChange30mPercent float64 `json:"priceChange30mPercent" bson:"priceChange30mPercent"`
	PriceChange1hPercent  float64 `json:"priceChange1hPercent" bson:"priceChange1hPercent"`
	PriceChange2hPercent  float64 `json:"priceChange2hPercent" bson:"priceChange2hPercent"`
	PriceChange4hPercent  float64 `json:"priceChange4hPercent" bson:"priceChange4hPercent"`
	PriceChange6hPercent  float64 `json:"priceChange6hPercent" bson:"priceChange6hPercent"`
	PriceChange8hPercent  float64 `json:"priceChange8hPercent" bson:"priceChange8hPercent"`
	PriceChange12hPercent float64 `json:"priceChange12hPercent" bson:"priceChange12hPercent"`
	PriceChange24hPercent float64 `json:"priceChange24hPercent" bson:"priceChange24hPercent"`

	UniqueWallet1m               int64   `json:"uniqueWallet1m" bson:"uniqueWallet1m"`
	UniqueWalletHistory1m        int64   `json:"uniqueWalletHistory1m" bson:"uniqueWalletHistory1m"`
	UniqueWallet1mChangePercent  float64 `json:"uniqueWallet1mChangePercent" bson:"uniqueWallet1mChangePercent"`
	UniqueWallet5m               int64   `json:"uniqueWallet5m" bson:"uniqueWallet5m"`
	UniqueWalletHistory5m        int64   `json:"uniqueWalletHistory5m" bson:"uniqueWalletHistory5m"`
	UniqueWallet5mChangePercent  float64 `json:"uniqueWallet5mChangePercent" bson:"uniqueWallet5mChangePercent"`
	UniqueWallet30m              int64   `json:"uniqueWallet30m" bson:"uniqueWallet30m"`
	UniqueWalletHistory30m       int64   `json:"uniqueWalletHistory30m" bson:"uniqueWalletHistory30m"`
	UniqueWallet30mChangePercent float64 `json:"uniqueWallet30mChangePercent" bson:"uniqueWallet30mChangePercent"`
	UniqueWallet1h               int64   `json:"uniqueWallet1h" bson:"uniqueWallet1h"`
	UniqueWalletHistory1h        int64   `json:"uniqueWalletHistory1h" bson:"uniqueWalletHistory1h"`
	UniqueWallet1hChangePercent  float64 `json:"uniqueWallet1hChangePercent" bson:"uniqueWallet1hChangePercent"`
	UniqueWallet2h               int64   `json:"uniqueWallet2h" bson:"uniqueWallet2h"`
	UniqueWalletHistory2h        int64   `json:"uniqueWalletHistory2h" bson:"uniqueWalletHistory2h"`
	UniqueWallet2hChangePercent  float64 `json:"uniqueWallet2hChangePercent" bson:"uniqueWallet2hChangePercent"`
	UniqueWallet4h               int64   `json:"uniqueWallet4h" bson:"uniqueWallet4h"`
	UniqueWalletHistory4h        int64   `json:"uniqueWalletHistory4h" bson:"uniqueWalletHistory4h"`
	UniqueWallet4hChangePercent  float64 `json:"uniqueWallet4hChangePercent" bson:"uniqueWallet4hChangePercent"`
	UniqueWallet8h               int64   `json:"uniqueWallet8h" bson:"uniqueWallet8h"`
	UniqueWalletHistory8h        int64   `json:"uniqueWalletHistory8h" bson:"uniqueWalletHistory8h"`
	UniqueWallet8hChangePercent  float64 `json:"uniqueWallet8hChangePercent" bson:"uniqueWallet8hChangePercent"`
	UniqueWallet24h              int64   `json:"uniqueWallet24h" bson:"uniqueWallet24h"`
	UniqueWalletHistory24h       int64   `json:"uniqueWalletHistory24h" bson:"uniqueWalletHistory24h"`
	UniqueWallet24hChangePercent float64 `json:"uniqueWallet24hChangePercent" bson:"uniqueWallet24hChangePercent"`
	TotalSupply                  float64 `json:"totalSupply" bson:"totalSupply"`
	CirculatingSupply            float64 `json:"circulatingSupply" bson:"circulatingSupply"`
	Holder                       int64   `json:"holder" bson:"holder"`

	Trade1m              int64   `json:"trade1m" bson:"trade1m"`
	TradeHistory1m       int64   `json:"tradeHistory1m" bson:"tradeHistory1m"`
	Trade1mChangePercent float64 `json:"trade1mChangePercent" bson:"trade1mChangePercent"`
	Sell1m               int64   `json:"sell1m" bson:"sell1m"`
	SellHistory1m        int64   `json:"sellHistory1m" bson:"sellHistory1m"`
	Sell1mChangePercent  float64 `json:"sell1mChangePercent" bson:"sell1mChangePercent"`
	Buy1m                int64   `json:"buy1m" bson:"buy1m"`
	BuyHistory1m         int64   `json:"buyHistory1m" bson:"buyHistory1m"`
	Buy1mChangePercent   float64 `json:"buy1mChangePercent" bson:"buy1mChangePercent"`
	V1m                  float64 `json:"v1m" bson:"v1m"`
	V1mUSD               float64 `json:"v1mUSD" bson:"v1mUSD"`
	VHistory1m           float64 `json:"vHistory1m" bson:"vHistory1m"`
	VHistory1mUSD        float64 `json:"vHistory1mUSD" bson:"vHistory1mUSD"`
	V1mChangePercent     float64 `json:"v1mChangePercent" bson:"v1mChangePercent"`
	VBuy1m               float64 `json:"vBuy1m" bson:"vBuy1m"`
	VBuy1mUSD            float64 `json:"vBuy1mUSD" bson:"vBuy1mUSD"`
	VBuyHistory1m        float64 `json:"vBuyHistory1m" bson:"vBuyHistory1m"`
	VBuyHistory1mUSD     float64 `json:"vBuyHistory1mUSD" bson:"vBuyHistory1mUSD"`
	VBuy1mChangePercent  float64 `json:"vBuy1mChangePercent" bson:"vBuy1mChangePercent"`
	VSell1m              float64 `json:"vSell1m" bson:"vSell1m"`
	VSell1mUSD           float64 `json:"vSell1mUSD" bson:"vSell1mUSD"`
	VSellHistory1m       float64 `json:"vSellHistory1m" bson:"vSellHistory1m"`
	VSellHistory1mUSD    float64 `json:"vSellHistory1mUSD" bson:"vSellHistory1mUSD"`
	VSell1mChangePercent float64 `json:"vSell1mChangePercent" bson:"vSell1mChangePercent"`

	Trade5m              int64   `json:"trade5m" bson:"trade5m"`
	TradeHistory5m       int64   `json:"tradeHistory5m" bson:"tradeHistory5m"`
	Trade5mChangePercent float64 `json:"trade5mChangePercent" bson:"trade5mChangePercent"`
	Sell5m               int64   `json:"sell5m" bson:"sell5m"`
	SellHistory5m        int64   `json:"sellHistory5m" bson:"sellHistory5m"`
	Sell5mChangePercent  float64 `json:"sell5mChangePercent" bson:"sell5mChangePercent"`
	Buy5m                int64   `json:"buy5m" bson:"buy5m"`
	BuyHistory5m         int64   `json:"buyHistory5m" bson:"buyHistory5m"`
	Buy5mChangePercent   float64 `json:"buy5mChangePercent" bson:"buy5mChangePercent"`
	V5m                  float64 `json:"v5m" bson:"v5m"`
	V5mUSD               float64 `json:"v5mUSD" bson:"v5mUSD"`
	VHistory5m           float64 `json:"vHistory5m" bson:"vHistory5m"`
	VHistory5mUSD        float64 `json:"vHistory5mUSD" bson:"vHistory5mUSD"`
	V5mChangePercent     float64 `json:"v5mChangePercent" bson:"v5mChangePercent"`
	VBuy5m               float64 `json:"vBuy5m" bson:"vBuy5m"`
	VBuy5mUSD            float64 `json:"vBuy5mUSD" bson:"vBuy5mUSD"`
	VBuyHistory5m        float64 `json:"vBuyHistory5m" bson:"vBuyHistory5m"`
	VBuyHistory5mUSD     float64 `json:"vBuyHistory5mUSD" bson:"vBuyHistory5mUSD"`
	VBuy5mChangePercent  float64 `json:"vBuy5mChangePercent" bson:"vBuy5mChangePercent"`
	VSell5m              float64 `json:"vSell5m" bson:"vSell5m"`
	VSell5mUSD           float64 `json:"vSell5mUSD" bson:"vSell5mUSD"`
	VSellHistory5m       float64 `json:"vSellHistory5m" bson:"vSellHistory5m"`
	VSellHistory5mUSD    float64 `json:"vSellHistory5mUSD" bson:"vSellHistory5mUSD"`
	VSell5mChangePercent float64 `json:"vSell5mChangePercent" bson:"vSell5mChangePercent"`

	Trade30m              int64   `json:"trade30m" bson:"trade30m"`
	TradeHistory30m       int64   `json:"tradeHistory30m" bson:"tradeHistory30m"`
	Trade30mChangePercent float64 `json:"trade30mChangePercent" bson:"trade30mChangePercent"`
	Sell30m               int64   `json:"sell30m" bson:"sell30m"`
	SellHistory30m        int64   `json:"sellHistory30m" bson:"sellHistory30m"`
	Sell30mChangePercent  float64 `json:"sell30mChangePercent" bson:"sell30mChangePercent"`
	Buy30m                int64   `json:"buy30m" bson:"buy30m"`
	BuyHistory30m         int64   `json:"buyHistory30m" bson:"buyHistory30m"`
	Buy30mChangePercent   float64 `json:"buy30mChangePercent" bson:"buy30mChangePercent"`
	V30m                  float64 `json:"v30m" bson:"v30m"`
	V30mUSD               float64 `json:"v30mUSD" bson:"v30mUSD"`
	VHistory30m           float64 `json:"vHistory30m" bson:"vHistory30m"`
	VHistory30mUSD        float64 `json:"vHistory30mUSD" bson:"vHistory30mUSD"`
	V30mChangePercent     float64 `json:"v30mChangePercent" bson:"v30mChangePercent"`
	VBuy30m               float64 `json:"vBuy30m" bson:"vBuy30m"`
	VBuy30mUSD            float64 `json:"vBuy30mUSD" bson:"vBuy30mUSD"`
	VBuyHistory30m        float64 `json:"vBuyHistory30m" bson:"vBuyHistory30m"`
	VBuyHistory30mUSD     float64 `json:"vBuyHistory30mUSD" bson:"vBuyHistory30mUSD"`
	VBuy30mChangePercent  float64 `json:"vBuy30mChangePercent" bson:"vBuy30mChangePercent"`
	VSell30m              float64 `json:"vSell30m" bson:"vSell30m"`
	VSell30mUSD           float64 `json:"vSell30mUSD" bson:"vSell30mUSD"`
	VSellHistory30m       float64 `json:"vSellHistory30m" bson:"vSellHistory30m"`
	VSellHistory30mUSD    float64 `json:"vSellHistory30mUSD" bson:"vSellHistory30mUSD"`
	VSell30mChangePercent float64 `json:"vSell30mChangePercent" bson:"vSell30mChangePercent"`

	Trade1h              int64   `json:"trade1h" bson:"trade1h"`
	TradeHistory1h       int64   `json:"tradeHistory1h" bson:"tradeHistory1h"`
	Trade1hChangePercent float64 `json:"trade1hChangePercent" bson:"trade1hChangePercent"`
	Sell1h               int64   `json:"sell1h" bson:"sell1h"`
	SellHistory1h        int64   `json:"sellHistory1h" bson:"sellHistory1h"`
	Sell1hChangePercent  float64 `json:"sell1hChangePercent" bson:"sell1hChangePercent"`
	Buy1h                int64   `json:"buy1h" bson:"buy1h"`
	BuyHistory1h         int64   `json:"buyHistory1h" bson:"buyHistory1h"`
	Buy1hChangePercent   float64 `json:"buy1hChangePercent" bson:"buy1hChangePercent"`
	V1h                  float64 `json:"v1h" bson:"v1h"`
	V1hUSD               float64 `json:"v1hUSD" bson:"v1hUSD"`
	VHistory1h           float64 `json:"vHistory1h" bson:"vHistory1h"`
	VHistory1hUSD        float64 `json:"vHistory1hUSD" bson:"vHistory1hUSD"`
	V1hChangePercent     float64 `json:"v1hChangePercent" bson:"v1hChangePercent"`
	VBuy1h               float64 `json:"vBuy1h" bson:"vBuy1h"`
	VBuy1hUSD            float64 `json:"vBuy1hUSD" bson:"vBuy1hUSD"`
	VBuyHistory1h        float64 `json:"vBuyHistory1h" bson:"vBuyHistory1h"`
	VBuyHistory1hUSD     float64 `json:"vBuyHistory1hUSD" bson:"vBuyHistory1hUSD"`
	VBuy1hChangePercent  float64 `json:"vBuy1hChangePercent" bson:"vBuy1hChangePercent"`
	VSell1h              float64 `json:"vSell1h" bson:"vSell1h"`
	VSell1hUSD           float64 `json:"vSell1hUSD" bson:"vSell1hUSD"`
	VSellHistory1h       float64 `json:"vSellHistory1h" bson:"vSellHistory1h"`
	VSellHistory1hUSD    float64 `json:"vSellHistory1hUSD" bson:"vSellHistory1hUSD"`
	VSell1hChangePercent float64 `json:"vSell1hChangePercent" bson:"vSell1hChangePercent"`

	Trade2h              int64   `json:"trade2h" bson:"trade2h"`
	TradeHistory2h       int64   `json:"tradeHistory2h" bson:"tradeHistory2h"`
	Trade2hChangePercent float64 `json:"trade2hChangePercent" bson:"trade2hChangePercent"`
	Sell2h               int64   `json:"sell2h" bson:"sell2h"`
	SellHistory2h        int64   `json:"sellHistory2h" bson:"sellHistory2h"`
	Sell2hChangePercent  float64 `json:"sell2hChangePercent" bson:"sell2hChangePercent"`
	Buy2h                int64   `json:"buy2h" bson:"buy2h"`
	BuyHistory2h         int64   `json:"buyHistory2h" bson:"buyHistory2h"`
	Buy2hChangePercent   float64 `json:"buy2hChangePercent" bson:"buy2hChangePercent"`
	V2h                  float64 `json:"v2h" bson:"v2h"`
	V2hUSD               float64 `json:"v2hUSD" bson:"v2hUSD"`
	VHistory2h           float64 `json:"vHistory2h" bson:"vHistory2h"`
	VHistory2hUSD        float64 `json:"vHistory2hUSD" bson:"vHistory2hUSD"`
	V2hChangePercent     float64 `json:"v2hChangePercent" bson:"v2hChangePercent"`
	VBuy2h               float64 `json:"vBuy2h" bson:"vBuy2h"`
	VBuy2hUSD            float64 `json:"vBuy2hUSD" bson:"vBuy2hUSD"`
	VBuyHistory2h        float64 `json:"vBuyHistory2h" bson:"vBuyHistory2h"`
	VBuyHistory2hUSD     float64 `json:"vBuyHistory2hUSD" bson:"vBuyHistory2hUSD"`
	VBuy2hChangePercent  float64 `json:"vBuy2hChangePercent" bson:"vBuy2hChangePercent"`
	VSell2h              float64 `json:"vSell2h" bson:"vSell2h"`
	VSell2hUSD           float64 `json:"vSell2hUSD" bson:"vSell2hUSD"`
	VSellHistory2h       float64 `json:"vSellHistory2h" bson:"vSellHistory2h"`
	VSellHistory2hUSD    float64 `json:"vSellHistory2hUSD" bson:"vSellHistory2hUSD"`
	VSell2hChangePercent float64 `json:"vSell2hChangePercent" bson:"vSell2hChangePercent"`

	Trade4h              int64   `json:"trade4h" bson:"trade4h"`
	TradeHistory4h       int64   `json:"tradeHistory4h" bson:"tradeHistory4h"`
	Trade4hChangePercent float64 `json:"trade4hChangePercent" bson:"trade4hChangePercent"`
	Sell4h               int64   `json:"sell4h" bson:"sell4h"`
	SellHistory4h        int64   `json:"sellHistory4h" bson:"sellHistory4h"`
	Sell4hChangePercent  float64 `json:"sell4hChangePercent" bson:"sell4hChangePercent"`
	Buy4h                int64   `json:"buy4h" bson:"buy4h"`
	BuyHistory4h         int64   `json:"buyHistory4h" bson:"buyHistory4h"`
	Buy4hChangePercent   float64 `json:"buy4hChangePercent" bson:"buy4hChangePercent"`
	V4h                  float64 `json:"v4h" bson:"v4h"`
	V4hUSD               float64 `json:"v4hUSD" bson:"v4hUSD"`
	VHistory4h           float64 `json:"vHistory4h" bson:"vHistory4h"`
	VHistory4hUSD        float64 `json:"vHistory4hUSD" bson:"vHistory4hUSD"`
	V4hChangePercent     float64 `json:"v4hChangePercent" bson:"v4hChangePercent"`
	VBuy4h               float64 `json:"vBuy4h" bson:"vBuy4h"`
	VBuy4hUSD            float64 `json:"vBuy4hUSD" bson:"vBuy4hUSD"`
	VBuyHistory4h        float64 `json:"vBuyHistory4h" bson:"vBuyHistory4h"`
	VBuyHistory4hUSD     float64 `json:"vBuyHistory4hUSD" bson:"vBuyHistory4hUSD"`
	VBuy4hChangePercent  float64 `json:"vBuy4hChangePercent" bson:"vBuy4hChangePercent"`
	VSell4h              float64 `json:"vSell4h" bson:"vSell4h"`
	VSell4hUSD           float64 `json:"vSell4hUSD" bson:"vSell4hUSD"`
	VSellHistory4h       float64 `json:"vSellHistory4h" bson:"vSellHistory4h"`
	VSellHistory4hUSD    float64 `json:"vSellHistory4hUSD" bson:"vSellHistory4hUSD"`
	VSell4hChangePercent float64 `json:"vSell4hChangePercent" bson:"vSell4hChangePercent"`

	Trade8h              int64   `json:"trade8h" bson:"trade8h"`
	TradeHistory8h       int64   `json:"tradeHistory8h" bson:"tradeHistory8h"`
	Trade8hChangePercent float64 `json:"trade8hChangePercent" bson:"trade8hChangePercent"`
	Sell8h               int64   `json:"sell8h" bson:"sell8h"`
	SellHistory8h        int64   `json:"sellHistory8h" bson:"sellHistory8h"`
	Sell8hChangePercent  float64 `json:"sell8hChangePercent" bson:"sell8hChangePercent"`
	Buy8h                int64   `json:"buy8h" bson:"buy8h"`
	BuyHistory8h         int64   `json:"buyHistory8h" bson:"buyHistory8h"`
	Buy8hChangePercent   float64 `json:"buy8hChangePercent" bson:"buy8hChangePercent"`
	V8h                  float64 `json:"v8h" bson:"v8h"`
	V8hUSD               float64 `json:"v8hUSD" bson:"v8hUSD"`
	VHistory8h           float64 `json:"vHistory8h" bson:"vHistory8h"`
	VHistory8hUSD        float64 `json:"vHistory8hUSD" bson:"vHistory8hUSD"`
	V8hChangePercent     float64 `json:"v8hChangePercent" bson:"v8hChangePercent"`
	VBuy8h               float64 `json:"vBuy8h" bson:"vBuy8h"`
	VBuy8hUSD            float64 `json:"vBuy8hUSD" bson:"vBuy8hUSD"`
	VBuyHistory8h        float64 `json:"vBuyHistory8h" bson:"vBuyHistory8h"`
	VBuyHistory8hUSD     float64 `json:"vBuyHistory8hUSD" bson:"vBuyHistory8hUSD"`
	VBuy8hChangePercent  float64 `json:"vBuy8hChangePercent" bson:"vBuy8hChangePercent"`
	VSell8h              float64 `json:"vSell8h" bson:"vSell8h"`
	VSell8hUSD           float64 `json:"vSell8hUSD" bson:"vSell8hUSD"`
	VSellHistory8h       float64 `json:"vSellHistory8h" bson:"vSellHistory8h"`
	VSellHistory8hUSD    float64 `json:"vSellHistory8hUSD" bson:"vSellHistory8hUSD"`
	VSell8hChangePercent float64 `json:"vSell8hChangePercent" bson:"vSell8hChangePercent"`

	Trade24h              int64   `json:"trade24h" bson:"trade24h"`
	TradeHistory24h       int64   `json:"tradeHistory24h" bson:"tradeHistory24h"`
	Trade24hChangePercent float64 `json:"trade24hChangePercent" bson:"trade24hChangePercent"`
	Sell24h               int64   `json:"sell24h" bson:"sell24h"`
	SellHistory24h        int64   `json:"sellHistory24h" bson:"sellHistory24h"`
	Sell24hChangePercent  float64 `json:"sell24hChangePercent" bson:"sell24hChangePercent"`
	Buy24h                int64   `json:"buy24h" bson:"buy24h"`
	BuyHistory24h         int64   `json:"buyHistory24h" bson:"buyHistory24h"`
	Buy24hChangePercent   float64 `json:"buy24hChangePercent" bson:"buy24hChangePercent"`
	V24h                  float64 `json:"v24h" bson:"v24h"`
	V24hUSD               float64 `json:"v24hUSD" bson:"v24hUSD"`
	VHistory24h           float64 `json:"vHistory24h" bson:"vHistory24h"`
	VHistory24hUSD        float64 `json:"vHistory24hUSD" bson:"vHistory24hUSD"`
	V24hChangePercent     float64 `json:"v24hChangePercent" bson:"v24hChangePercent"`
	VBuy24h               float64 `json:"vBuy24h" bson:"vBuy24h"`
	VBuy24hUSD            float64 `json:"vBuy24hUSD" bson:"vBuy24hUSD"`
	VBuyHistory24h        float64 `json:"vBuyHistory24h" bson:"vBuyHistory24h"`
	VBuyHistory24hUSD     float64 `json:"vBuyHistory24hUSD" bson:"vBuyHistory24hUSD"`
	VBuy24hChangePercent  float64 `json:"vBuy24hChangePercent" bson:"vBuy24hChangePercent"`
	VSell24h              float64 `json:"vSell24h" bson:"vSell24h"`
	VSell24hUSD           float64 `json:"vSell24hUSD" bson:"vSell24hUSD"`
	VSellHistory24h       float64 `json:"vSellHistory24h" bson:"vSellHistory24h"`
	VSellHistory24hUSD    float64 `json:"vSellHistory24hUSD" bson:"vSellHistory24hUSD"`
	VSell24hChangePercent float64 `json:"vSell24hChangePercent" bson:"vSell24hChangePercent"`

	NumberMarkets   int64    `json:"numberMarkets" bson:"numberMarkets"`
	IsScaledUIToken bool     `json:"isScaledUiToken" bson:"isScaledUiToken"`
	Multiplier      *float64 `json:"multiplier" bson:"multiplier"`
}

// ============================================================================
// Token Metadata, Market Data, and Trade Data Types
// ============================================================================

// RespTokenMetadata represents token metadata response
type RespTokenMetadata struct {
	Address    string          `json:"address" bson:"address"`
	Symbol     string          `json:"symbol" bson:"symbol"`
	Name       string          `json:"name" bson:"name"`
	Decimals   int64           `json:"decimals" bson:"decimals"`
	Extensions TokenExtensions `json:"extensions" bson:"extensions"`
	LogoURI    string          `json:"logo_uri" bson:"logo_uri"`
}

type RespMultiTokenMetadata = map[string]RespTokenMetadata

// RespTokenMarketData represents token market data response
type RespTokenMarketData struct {
	Address           string   `json:"address" bson:"address"`
	Price             float64  `json:"price" bson:"price"`
	Liquidity         float64  `json:"liquidity" bson:"liquidity"`
	TotalSupply       float64  `json:"total_supply" bson:"total_supply"`
	CirculatingSupply float64  `json:"circulating_supply" bson:"circulating_supply"`
	FDV               float64  `json:"fdv" bson:"fdv"`
	MarketCap         float64  `json:"market_cap" bson:"market_cap"`
	IsScaledUIToken   bool     `json:"is_scaled_ui_token" bson:"is_scaled_ui_token"`
	Multiplier        *float64 `json:"multiplier" bson:"multiplier"`
}

// RespTokenTradeData represents token trade data response (this is a large struct, split for readability)
type RespTokenTradeData struct {
	Address                      string  `json:"address" bson:"address"`
	Holder                       int64   `json:"holder" bson:"holder"`
	Market                       int64   `json:"market" bson:"market"`
	LastTradeUnixTime            int64   `json:"last_trade_unix_time" bson:"last_trade_unix_time"`
	LastTradeHumanTime           string  `json:"last_trade_human_time" bson:"last_trade_human_time"`
	Price                        float64 `json:"price" bson:"price"`
	History1mPrice               float64 `json:"history_1m_price" bson:"history_1m_price"`
	PriceChange1mPercent         float64 `json:"price_change_1m_percent" bson:"price_change_1m_percent"`
	History5mPrice               float64 `json:"history_5m_price" bson:"history_5m_price"`
	PriceChange5mPercent         float64 `json:"price_change_5m_percent" bson:"price_change_5m_percent"`
	History30mPrice              float64 `json:"history_30m_price" bson:"history_30m_price"`
	PriceChange30mPercent        float64 `json:"price_change_30m_percent" bson:"price_change_30m_percent"`
	History1hPrice               float64 `json:"history_1h_price" bson:"history_1h_price"`
	PriceChange1hPercent         float64 `json:"price_change_1h_percent" bson:"price_change_1h_percent"`
	History2hPrice               float64 `json:"history_2h_price" bson:"history_2h_price"`
	PriceChange2hPercent         float64 `json:"price_change_2h_percent" bson:"price_change_2h_percent"`
	History4hPrice               float64 `json:"history_4h_price" bson:"history_4h_price"`
	PriceChange4hPercent         float64 `json:"price_change_4h_percent" bson:"price_change_4h_percent"`
	History6hPrice               float64 `json:"history_6h_price" bson:"history_6h_price"`
	PriceChange6hPercent         float64 `json:"price_change_6h_percent" bson:"price_change_6h_percent"`
	History8hPrice               float64 `json:"history_8h_price" bson:"history_8h_price"`
	PriceChange8hPercent         float64 `json:"price_change_8h_percent" bson:"price_change_8h_percent"`
	History12hPrice              float64 `json:"history_12h_price" bson:"history_12h_price"`
	PriceChange12hPercent        float64 `json:"price_change_12h_percent" bson:"price_change_12h_percent"`
	History24hPrice              float64 `json:"history_24h_price" bson:"history_24h_price"`
	PriceChange24hPercent        float64 `json:"price_change_24h_percent" bson:"price_change_24h_percent"`
	UniqueWallet1m               int64   `json:"unique_wallet_1m" bson:"unique_wallet_1m"`
	UniqueWalletHistory1m        int64   `json:"unique_wallet_history_1m" bson:"unique_wallet_history_1m"`
	UniqueWallet1mChangePercent  float64 `json:"unique_wallet_1m_change_percent" bson:"unique_wallet_1m_change_percent"`
	UniqueWallet5m               int64   `json:"unique_wallet_5m" bson:"unique_wallet_5m"`
	UniqueWalletHistory5m        int64   `json:"unique_wallet_history_5m" bson:"unique_wallet_history_5m"`
	UniqueWallet5mChangePercent  float64 `json:"unique_wallet_5m_change_percent" bson:"unique_wallet_5m_change_percent"`
	UniqueWallet30m              int64   `json:"unique_wallet_30m" bson:"unique_wallet_30m"`
	UniqueWalletHistory30m       int64   `json:"unique_wallet_history_30m" bson:"unique_wallet_history_30m"`
	UniqueWallet30mChangePercent float64 `json:"unique_wallet_30m_change_percent" bson:"unique_wallet_30m_change_percent"`
	UniqueWallet1h               int64   `json:"unique_wallet_1h" bson:"unique_wallet_1h"`
	UniqueWalletHistory1h        int64   `json:"unique_wallet_history_1h" bson:"unique_wallet_history_1h"`
	UniqueWallet1hChangePercent  float64 `json:"unique_wallet_1h_change_percent" bson:"unique_wallet_1h_change_percent"`
	UniqueWallet2h               int64   `json:"unique_wallet_2h" bson:"unique_wallet_2h"`
	UniqueWalletHistory2h        int64   `json:"unique_wallet_history_2h" bson:"unique_wallet_history_2h"`
	UniqueWallet2hChangePercent  float64 `json:"unique_wallet_2h_change_percent" bson:"unique_wallet_2h_change_percent"`
	UniqueWallet4h               int64   `json:"unique_wallet_4h" bson:"unique_wallet_4h"`
	UniqueWalletHistory4h        int64   `json:"unique_wallet_history_4h" bson:"unique_wallet_history_4h"`
	UniqueWallet4hChangePercent  float64 `json:"unique_wallet_4h_change_percent" bson:"unique_wallet_4h_change_percent"`
	UniqueWallet8h               int64   `json:"unique_wallet_8h" bson:"unique_wallet_8h"`
	UniqueWalletHistory8h        int64   `json:"unique_wallet_history_8h" bson:"unique_wallet_history_8h"`
	UniqueWallet8hChangePercent  float64 `json:"unique_wallet_8h_change_percent" bson:"unique_wallet_8h_change_percent"`
	UniqueWallet24h              int64   `json:"unique_wallet_24h" bson:"unique_wallet_24h"`
	UniqueWalletHistory24h       int64   `json:"unique_wallet_history_24h" bson:"unique_wallet_history_24h"`
	UniqueWallet24hChangePercent float64 `json:"unique_wallet_24h_change_percent" bson:"unique_wallet_24h_change_percent"`
	Trade1m                      int64   `json:"trade_1m" bson:"trade_1m"`
	TradeHistory1m               int64   `json:"trade_history_1m" bson:"trade_history_1m"`
	Trade1mChangePercent         float64 `json:"trade_1m_change_percent" bson:"trade_1m_change_percent"`
	Sell1m                       int64   `json:"sell_1m" bson:"sell_1m"`
	SellHistory1m                int64   `json:"sell_history_1m" bson:"sell_history_1m"`
	Sell1mChangePercent          float64 `json:"sell_1m_change_percent" bson:"sell_1m_change_percent"`
	Buy1m                        int64   `json:"buy_1m" bson:"buy_1m"`
	BuyHistory1m                 int64   `json:"buy_history_1m" bson:"buy_history_1m"`
	Buy1mChangePercent           float64 `json:"buy_1m_change_percent" bson:"buy_1m_change_percent"`
	Volume1m                     float64 `json:"volume_1m" bson:"volume_1m"`
	Volume1mUSD                  float64 `json:"volume_1m_usd" bson:"volume_1m_usd"`
	VolumeHistory1m              float64 `json:"volume_history_1m" bson:"volume_history_1m"`
	VolumeHistory1mUSD           float64 `json:"volume_history_1m_usd" bson:"volume_history_1m_usd"`
	Volume1mChangePercent        float64 `json:"volume_1m_change_percent" bson:"volume_1m_change_percent"`
	VolumeBuy1m                  float64 `json:"volume_buy_1m" bson:"volume_buy_1m"`
	VolumeBuy1mUSD               float64 `json:"volume_buy_1m_usd" bson:"volume_buy_1m_usd"`
	VolumeBuyHistory1m           float64 `json:"volume_buy_history_1m" bson:"volume_buy_history_1m"`
	VolumeBuyHistory1mUSD        float64 `json:"volume_buy_history_1m_usd" bson:"volume_buy_history_1m_usd"`
	VolumeBuy1mChangePercent     float64 `json:"volume_buy_1m_change_percent" bson:"volume_buy_1m_change_percent"`
	VolumeSell1m                 float64 `json:"volume_sell_1m" bson:"volume_sell_1m"`
	VolumeSell1mUSD              float64 `json:"volume_sell_1m_usd" bson:"volume_sell_1m_usd"`
	VolumeSellHistory1m          float64 `json:"volume_sell_history_1m" bson:"volume_sell_history_1m"`
	VolumeSellHistory1mUSD       float64 `json:"volume_sell_history_1m_usd" bson:"volume_sell_history_1m_usd"`
	VolumeSell1mChangePercent    float64 `json:"volume_sell_1m_change_percent" bson:"volume_sell_1m_change_percent"`

	// 5m time period statistics
	Trade5m                   int64   `json:"trade_5m" bson:"trade_5m"`
	TradeHistory5m            int64   `json:"trade_history_5m" bson:"trade_history_5m"`
	Trade5mChangePercent      float64 `json:"trade_5m_change_percent" bson:"trade_5m_change_percent"`
	Sell5m                    int64   `json:"sell_5m" bson:"sell_5m"`
	SellHistory5m             int64   `json:"sell_history_5m" bson:"sell_history_5m"`
	Sell5mChangePercent       float64 `json:"sell_5m_change_percent" bson:"sell_5m_change_percent"`
	Buy5m                     int64   `json:"buy_5m" bson:"buy_5m"`
	BuyHistory5m              int64   `json:"buy_history_5m" bson:"buy_history_5m"`
	Buy5mChangePercent        float64 `json:"buy_5m_change_percent" bson:"buy_5m_change_percent"`
	Volume5m                  float64 `json:"volume_5m" bson:"volume_5m"`
	Volume5mUSD               float64 `json:"volume_5m_usd" bson:"volume_5m_usd"`
	VolumeHistory5m           float64 `json:"volume_history_5m" bson:"volume_history_5m"`
	VolumeHistory5mUSD        float64 `json:"volume_history_5m_usd" bson:"volume_history_5m_usd"`
	Volume5mChangePercent     float64 `json:"volume_5m_change_percent" bson:"volume_5m_change_percent"`
	VolumeBuy5m               float64 `json:"volume_buy_5m" bson:"volume_buy_5m"`
	VolumeBuy5mUSD            float64 `json:"volume_buy_5m_usd" bson:"volume_buy_5m_usd"`
	VolumeBuyHistory5m        float64 `json:"volume_buy_history_5m" bson:"volume_buy_history_5m"`
	VolumeBuyHistory5mUSD     float64 `json:"volume_buy_history_5m_usd" bson:"volume_buy_history_5m_usd"`
	VolumeBuy5mChangePercent  float64 `json:"volume_buy_5m_change_percent" bson:"volume_buy_5m_change_percent"`
	VolumeSell5m              float64 `json:"volume_sell_5m" bson:"volume_sell_5m"`
	VolumeSell5mUSD           float64 `json:"volume_sell_5m_usd" bson:"volume_sell_5m_usd"`
	VolumeSellHistory5m       float64 `json:"volume_sell_history_5m" bson:"volume_sell_history_5m"`
	VolumeSellHistory5mUSD    float64 `json:"volume_sell_history_5m_usd" bson:"volume_sell_history_5m_usd"`
	VolumeSell5mChangePercent float64 `json:"volume_sell_5m_change_percent" bson:"volume_sell_5m_change_percent"`

	// 30m time period statistics
	Trade30m                   int64   `json:"trade_30m" bson:"trade_30m"`
	TradeHistory30m            int64   `json:"trade_history_30m" bson:"trade_history_30m"`
	Trade30mChangePercent      float64 `json:"trade_30m_change_percent" bson:"trade_30m_change_percent"`
	Sell30m                    int64   `json:"sell_30m" bson:"sell_30m"`
	SellHistory30m             int64   `json:"sell_history_30m" bson:"sell_history_30m"`
	Sell30mChangePercent       float64 `json:"sell_30m_change_percent" bson:"sell_30m_change_percent"`
	Buy30m                     int64   `json:"buy_30m" bson:"buy_30m"`
	BuyHistory30m              int64   `json:"buy_history_30m" bson:"buy_history_30m"`
	Buy30mChangePercent        float64 `json:"buy_30m_change_percent" bson:"buy_30m_change_percent"`
	Volume30m                  float64 `json:"volume_30m" bson:"volume_30m"`
	Volume30mUSD               float64 `json:"volume_30m_usd" bson:"volume_30m_usd"`
	VolumeHistory30m           float64 `json:"volume_history_30m" bson:"volume_history_30m"`
	VolumeHistory30mUSD        float64 `json:"volume_history_30m_usd" bson:"volume_history_30m_usd"`
	Volume30mChangePercent     float64 `json:"volume_30m_change_percent" bson:"volume_30m_change_percent"`
	VolumeBuy30m               float64 `json:"volume_buy_30m" bson:"volume_buy_30m"`
	VolumeBuy30mUSD            float64 `json:"volume_buy_30m_usd" bson:"volume_buy_30m_usd"`
	VolumeBuyHistory30m        float64 `json:"volume_buy_history_30m" bson:"volume_buy_history_30m"`
	VolumeBuyHistory30mUSD     float64 `json:"volume_buy_history_30m_usd" bson:"volume_buy_history_30m_usd"`
	VolumeBuy30mChangePercent  float64 `json:"volume_buy_30m_change_percent" bson:"volume_buy_30m_change_percent"`
	VolumeSell30m              float64 `json:"volume_sell_30m" bson:"volume_sell_30m"`
	VolumeSell30mUSD           float64 `json:"volume_sell_30m_usd" bson:"volume_sell_30m_usd"`
	VolumeSellHistory30m       float64 `json:"volume_sell_history_30m" bson:"volume_sell_history_30m"`
	VolumeSellHistory30mUSD    float64 `json:"volume_sell_history_30m_usd" bson:"volume_sell_history_30m_usd"`
	VolumeSell30mChangePercent float64 `json:"volume_sell_30m_change_percent" bson:"volume_sell_30m_change_percent"`

	// 1h time period statistics
	Trade1h                   int64   `json:"trade_1h" bson:"trade_1h"`
	TradeHistory1h            int64   `json:"trade_history_1h" bson:"trade_history_1h"`
	Trade1hChangePercent      float64 `json:"trade_1h_change_percent" bson:"trade_1h_change_percent"`
	Sell1h                    int64   `json:"sell_1h" bson:"sell_1h"`
	SellHistory1h             int64   `json:"sell_history_1h" bson:"sell_history_1h"`
	Sell1hChangePercent       float64 `json:"sell_1h_change_percent" bson:"sell_1h_change_percent"`
	Buy1h                     int64   `json:"buy_1h" bson:"buy_1h"`
	BuyHistory1h              int64   `json:"buy_history_1h" bson:"buy_history_1h"`
	Buy1hChangePercent        float64 `json:"buy_1h_change_percent" bson:"buy_1h_change_percent"`
	Volume1h                  float64 `json:"volume_1h" bson:"volume_1h"`
	Volume1hUSD               float64 `json:"volume_1h_usd" bson:"volume_1h_usd"`
	VolumeHistory1h           float64 `json:"volume_history_1h" bson:"volume_history_1h"`
	VolumeHistory1hUSD        float64 `json:"volume_history_1h_usd" bson:"volume_history_1h_usd"`
	Volume1hChangePercent     float64 `json:"volume_1h_change_percent" bson:"volume_1h_change_percent"`
	VolumeBuy1h               float64 `json:"volume_buy_1h" bson:"volume_buy_1h"`
	VolumeBuy1hUSD            float64 `json:"volume_buy_1h_usd" bson:"volume_buy_1h_usd"`
	VolumeBuyHistory1h        float64 `json:"volume_buy_history_1h" bson:"volume_buy_history_1h"`
	VolumeBuyHistory1hUSD     float64 `json:"volume_buy_history_1h_usd" bson:"volume_buy_history_1h_usd"`
	VolumeBuy1hChangePercent  float64 `json:"volume_buy_1h_change_percent" bson:"volume_buy_1h_change_percent"`
	VolumeSell1h              float64 `json:"volume_sell_1h" bson:"volume_sell_1h"`
	VolumeSell1hUSD           float64 `json:"volume_sell_1h_usd" bson:"volume_sell_1h_usd"`
	VolumeSellHistory1h       float64 `json:"volume_sell_history_1h" bson:"volume_sell_history_1h"`
	VolumeSellHistory1hUSD    float64 `json:"volume_sell_history_1h_usd" bson:"volume_sell_history_1h_usd"`
	VolumeSell1hChangePercent float64 `json:"volume_sell_1h_change_percent" bson:"volume_sell_1h_change_percent"`

	// 2h time period statistics
	Trade2h                   int64   `json:"trade_2h" bson:"trade_2h"`
	TradeHistory2h            int64   `json:"trade_history_2h" bson:"trade_history_2h"`
	Trade2hChangePercent      float64 `json:"trade_2h_change_percent" bson:"trade_2h_change_percent"`
	Sell2h                    int64   `json:"sell_2h" bson:"sell_2h"`
	SellHistory2h             int64   `json:"sell_history_2h" bson:"sell_history_2h"`
	Sell2hChangePercent       float64 `json:"sell_2h_change_percent" bson:"sell_2h_change_percent"`
	Buy2h                     int64   `json:"buy_2h" bson:"buy_2h"`
	BuyHistory2h              int64   `json:"buy_history_2h" bson:"buy_history_2h"`
	Buy2hChangePercent        float64 `json:"buy_2h_change_percent" bson:"buy_2h_change_percent"`
	Volume2h                  float64 `json:"volume_2h" bson:"volume_2h"`
	Volume2hUSD               float64 `json:"volume_2h_usd" bson:"volume_2h_usd"`
	VolumeHistory2h           float64 `json:"volume_history_2h" bson:"volume_history_2h"`
	VolumeHistory2hUSD        float64 `json:"volume_history_2h_usd" bson:"volume_history_2h_usd"`
	Volume2hChangePercent     float64 `json:"volume_2h_change_percent" bson:"volume_2h_change_percent"`
	VolumeBuy2h               float64 `json:"volume_buy_2h" bson:"volume_buy_2h"`
	VolumeBuy2hUSD            float64 `json:"volume_buy_2h_usd" bson:"volume_buy_2h_usd"`
	VolumeBuyHistory2h        float64 `json:"volume_buy_history_2h" bson:"volume_buy_history_2h"`
	VolumeBuyHistory2hUSD     float64 `json:"volume_buy_history_2h_usd" bson:"volume_buy_history_2h_usd"`
	VolumeBuy2hChangePercent  float64 `json:"volume_buy_2h_change_percent" bson:"volume_buy_2h_change_percent"`
	VolumeSell2h              float64 `json:"volume_sell_2h" bson:"volume_sell_2h"`
	VolumeSell2hUSD           float64 `json:"volume_sell_2h_usd" bson:"volume_sell_2h_usd"`
	VolumeSellHistory2h       float64 `json:"volume_sell_history_2h" bson:"volume_sell_history_2h"`
	VolumeSellHistory2hUSD    float64 `json:"volume_sell_history_2h_usd" bson:"volume_sell_history_2h_usd"`
	VolumeSell2hChangePercent float64 `json:"volume_sell_2h_change_percent" bson:"volume_sell_2h_change_percent"`

	// 4h time period statistics
	Trade4h                   int64   `json:"trade_4h" bson:"trade_4h"`
	TradeHistory4h            int64   `json:"trade_history_4h" bson:"trade_history_4h"`
	Trade4hChangePercent      float64 `json:"trade_4h_change_percent" bson:"trade_4h_change_percent"`
	Sell4h                    int64   `json:"sell_4h" bson:"sell_4h"`
	SellHistory4h             int64   `json:"sell_history_4h" bson:"sell_history_4h"`
	Sell4hChangePercent       float64 `json:"sell_4h_change_percent" bson:"sell_4h_change_percent"`
	Buy4h                     int64   `json:"buy_4h" bson:"buy_4h"`
	BuyHistory4h              int64   `json:"buy_history_4h" bson:"buy_history_4h"`
	Buy4hChangePercent        float64 `json:"buy_4h_change_percent" bson:"buy_4h_change_percent"`
	Volume4h                  float64 `json:"volume_4h" bson:"volume_4h"`
	Volume4hUSD               float64 `json:"volume_4h_usd" bson:"volume_4h_usd"`
	VolumeHistory4h           float64 `json:"volume_history_4h" bson:"volume_history_4h"`
	VolumeHistory4hUSD        float64 `json:"volume_history_4h_usd" bson:"volume_history_4h_usd"`
	Volume4hChangePercent     float64 `json:"volume_4h_change_percent" bson:"volume_4h_change_percent"`
	VolumeBuy4h               float64 `json:"volume_buy_4h" bson:"volume_buy_4h"`
	VolumeBuy4hUSD            float64 `json:"volume_buy_4h_usd" bson:"volume_buy_4h_usd"`
	VolumeBuyHistory4h        float64 `json:"volume_buy_history_4h" bson:"volume_buy_history_4h"`
	VolumeBuyHistory4hUSD     float64 `json:"volume_buy_history_4h_usd" bson:"volume_buy_history_4h_usd"`
	VolumeBuy4hChangePercent  float64 `json:"volume_buy_4h_change_percent" bson:"volume_buy_4h_change_percent"`
	VolumeSell4h              float64 `json:"volume_sell_4h" bson:"volume_sell_4h"`
	VolumeSell4hUSD           float64 `json:"volume_sell_4h_usd" bson:"volume_sell_4h_usd"`
	VolumeSellHistory4h       float64 `json:"volume_sell_history_4h" bson:"volume_sell_history_4h"`
	VolumeSellHistory4hUSD    float64 `json:"volume_sell_history_4h_usd" bson:"volume_sell_history_4h_usd"`
	VolumeSell4hChangePercent float64 `json:"volume_sell_4h_change_percent" bson:"volume_sell_4h_change_percent"`

	// 8h time period statistics
	Trade8h                   int64   `json:"trade_8h" bson:"trade_8h"`
	TradeHistory8h            int64   `json:"trade_history_8h" bson:"trade_history_8h"`
	Trade8hChangePercent      float64 `json:"trade_8h_change_percent" bson:"trade_8h_change_percent"`
	Sell8h                    int64   `json:"sell_8h" bson:"sell_8h"`
	SellHistory8h             int64   `json:"sell_history_8h" bson:"sell_history_8h"`
	Sell8hChangePercent       float64 `json:"sell_8h_change_percent" bson:"sell_8h_change_percent"`
	Buy8h                     int64   `json:"buy_8h" bson:"buy_8h"`
	BuyHistory8h              int64   `json:"buy_history_8h" bson:"buy_history_8h"`
	Buy8hChangePercent        float64 `json:"buy_8h_change_percent" bson:"buy_8h_change_percent"`
	Volume8h                  float64 `json:"volume_8h" bson:"volume_8h"`
	Volume8hUSD               float64 `json:"volume_8h_usd" bson:"volume_8h_usd"`
	VolumeHistory8h           float64 `json:"volume_history_8h" bson:"volume_history_8h"`
	VolumeHistory8hUSD        float64 `json:"volume_history_8h_usd" bson:"volume_history_8h_usd"`
	Volume8hChangePercent     float64 `json:"volume_8h_change_percent" bson:"volume_8h_change_percent"`
	VolumeBuy8h               float64 `json:"volume_buy_8h" bson:"volume_buy_8h"`
	VolumeBuy8hUSD            float64 `json:"volume_buy_8h_usd" bson:"volume_buy_8h_usd"`
	VolumeBuyHistory8h        float64 `json:"volume_buy_history_8h" bson:"volume_buy_history_8h"`
	VolumeBuyHistory8hUSD     float64 `json:"volume_buy_history_8h_usd" bson:"volume_buy_history_8h_usd"`
	VolumeBuy8hChangePercent  float64 `json:"volume_buy_8h_change_percent" bson:"volume_buy_8h_change_percent"`
	VolumeSell8h              float64 `json:"volume_sell_8h" bson:"volume_sell_8h"`
	VolumeSell8hUSD           float64 `json:"volume_sell_8h_usd" bson:"volume_sell_8h_usd"`
	VolumeSellHistory8h       float64 `json:"volume_sell_history_8h" bson:"volume_sell_history_8h"`
	VolumeSellHistory8hUSD    float64 `json:"volume_sell_history_8h_usd" bson:"volume_sell_history_8h_usd"`
	VolumeSell8hChangePercent float64 `json:"volume_sell_8h_change_percent" bson:"volume_sell_8h_change_percent"`

	// 24h time period statistics
	Trade24h                   int64   `json:"trade_24h" bson:"trade_24h"`
	TradeHistory24h            int64   `json:"trade_history_24h" bson:"trade_history_24h"`
	Trade24hChangePercent      float64 `json:"trade_24h_change_percent" bson:"trade_24h_change_percent"`
	Sell24h                    int64   `json:"sell_24h" bson:"sell_24h"`
	SellHistory24h             int64   `json:"sell_history_24h" bson:"sell_history_24h"`
	Sell24hChangePercent       float64 `json:"sell_24h_change_percent" bson:"sell_24h_change_percent"`
	Buy24h                     int64   `json:"buy_24h" bson:"buy_24h"`
	BuyHistory24h              int64   `json:"buy_history_24h" bson:"buy_history_24h"`
	Buy24hChangePercent        float64 `json:"buy_24h_change_percent" bson:"buy_24h_change_percent"`
	Volume24h                  float64 `json:"volume_24h" bson:"volume_24h"`
	Volume24hUSD               float64 `json:"volume_24h_usd" bson:"volume_24h_usd"`
	VolumeHistory24h           float64 `json:"volume_history_24h" bson:"volume_history_24h"`
	VolumeHistory24hUSD        float64 `json:"volume_history_24h_usd" bson:"volume_history_24h_usd"`
	Volume24hChangePercent     float64 `json:"volume_24h_change_percent" bson:"volume_24h_change_percent"`
	VolumeBuy24h               float64 `json:"volume_buy_24h" bson:"volume_buy_24h"`
	VolumeBuy24hUSD            float64 `json:"volume_buy_24h_usd" bson:"volume_buy_24h_usd"`
	VolumeBuyHistory24h        float64 `json:"volume_buy_history_24h" bson:"volume_buy_history_24h"`
	VolumeBuyHistory24hUSD     float64 `json:"volume_buy_history_24h_usd" bson:"volume_buy_history_24h_usd"`
	VolumeBuy24hChangePercent  float64 `json:"volume_buy_24h_change_percent" bson:"volume_buy_24h_change_percent"`
	VolumeSell24h              float64 `json:"volume_sell_24h" bson:"volume_sell_24h"`
	VolumeSell24hUSD           float64 `json:"volume_sell_24h_usd" bson:"volume_sell_24h_usd"`
	VolumeSellHistory24h       float64 `json:"volume_sell_history_24h" bson:"volume_sell_history_24h"`
	VolumeSellHistory24hUSD    float64 `json:"volume_sell_history_24h_usd" bson:"volume_sell_history_24h_usd"`
	VolumeSell24hChangePercent float64 `json:"volume_sell_24h_change_percent" bson:"volume_sell_24h_change_percent"`

	IsScaledUIToken bool     `json:"is_scaled_ui_token" bson:"is_scaled_ui_token"`
	Multiplier      *float64 `json:"multiplier" bson:"multiplier"`
}

// ============================================================================
// Token Security and Additional Types
// ============================================================================

// RespTokenSecurity represents token security information response
type RespTokenSecurity struct {
	CreatorAddress                 *string  `json:"creatorAddress" bson:"creatorAddress"`
	CreatorOwnerAddress            *string  `json:"creatorOwnerAddress" bson:"creatorOwnerAddress"`
	OwnerAddress                   *string  `json:"ownerAddress" bson:"ownerAddress"`
	OwnerOfOwnerAddress            *string  `json:"ownerOfOwnerAddress" bson:"ownerOfOwnerAddress"`
	CreationTx                     *string  `json:"creationTx" bson:"creationTx"`
	CreationTime                   *int64   `json:"creationTime" bson:"creationTime"`
	CreationSlot                   *int64   `json:"creationSlot" bson:"creationSlot"`
	MintTx                         *string  `json:"mintTx" bson:"mintTx"`
	MintTime                       *int64   `json:"mintTime" bson:"mintTime"`
	MintSlot                       *int64   `json:"mintSlot" bson:"mintSlot"`
	CreatorBalance                 *float64 `json:"creatorBalance" bson:"creatorBalance"`
	OwnerBalance                   *float64 `json:"ownerBalance" bson:"ownerBalance"`
	OwnerPercentage                *float64 `json:"ownerPercentage" bson:"ownerPercentage"`
	CreatorPercentage              *float64 `json:"creatorPercentage" bson:"creatorPercentage"`
	MetaplexUpdateAuthority        *string  `json:"metaplexUpdateAuthority" bson:"metaplexUpdateAuthority"`
	MetaplexOwnerUpdateAuthority   *string  `json:"metaplexOwnerUpdateAuthority" bson:"metaplexOwnerUpdateAuthority"`
	MetaplexUpdateAuthorityBalance *float64 `json:"metaplexUpdateAuthorityBalance" bson:"metaplexUpdateAuthorityBalance"`
	MetaplexUpdateAuthorityPercent *float64 `json:"metaplexUpdateAuthorityPercent" bson:"metaplexUpdateAuthorityPercent"`
	MutableMetadata                *bool    `json:"mutableMetadata" bson:"mutableMetadata"`
	Top10HolderBalance             *float64 `json:"top10HolderBalance" bson:"top10HolderBalance"`
	Top10HolderPercent             *float64 `json:"top10HolderPercent" bson:"top10HolderPercent"`
	Top10UserBalance               *float64 `json:"top10UserBalance" bson:"top10UserBalance"`
	Top10UserPercent               *float64 `json:"top10UserPercent" bson:"top10UserPercent"`
	IsTrueToken                    *bool    `json:"isTrueToken" bson:"isTrueToken"`
	FakeToken                      *bool    `json:"fakeToken" bson:"fakeToken"`
	TotalSupply                    *float64 `json:"totalSupply" bson:"totalSupply"`
	PreMarketHolder                []any    `json:"preMarketHolder" bson:"preMarketHolder"`
	LockInfo                       any      `json:"lockInfo" bson:"lockInfo"`
	Freezeable                     *bool    `json:"freezeable" bson:"freezeable"`
	FreezeAuthority                *string  `json:"freezeAuthority" bson:"freezeAuthority"`
	TransferFeeEnable              *bool    `json:"transferFeeEnable" bson:"transferFeeEnable"`
	TransferFeeData                any      `json:"transferFeeData" bson:"transferFeeData"`
	IsToken2022                    bool     `json:"isToken2022" bson:"isToken2022"`
	NonTransferable                *bool    `json:"nonTransferable" bson:"nonTransferable"`
	JupStrictList                  bool     `json:"jupStrictList" bson:"jupStrictList"`
}

// RespTokenCreationInfo represents token creation information response
type RespTokenCreationInfo struct {
	TxHash         string `json:"txHash" bson:"txHash"`
	Slot           int64  `json:"slot" bson:"slot"`
	TokenAddress   string `json:"tokenAddress" bson:"tokenAddress"`
	Decimals       int64  `json:"decimals" bson:"decimals"`
	Owner          string `json:"owner" bson:"owner"`
	BlockUnixTime  int64  `json:"blockUnixTime" bson:"blockUnixTime"`
	BlockHumanTime string `json:"blockHumanTime" bson:"blockHumanTime"`
}

// RespTokenMintBurnTxItem represents token mint/burn transaction details
type RespTokenMintBurnTxItem struct {
	Amount         string       `json:"amount" bson:"amount"`
	BlockHumanTime string       `json:"block_human_time" bson:"block_human_time"`
	BlockTime      int64        `json:"block_time" bson:"block_time"`
	CommonType     MintBurnType `json:"common_type" bson:"common_type"`
	Decimals       int64        `json:"decimals" bson:"decimals"`
	Mint           string       `json:"mint" bson:"mint"`
	ProgramID      string       `json:"program_id" bson:"program_id"`
	Slot           int64        `json:"slot" bson:"slot"`
	TxHash         string       `json:"tx_hash" bson:"tx_hash"`
	UIAmount       float64      `json:"ui_amount" bson:"ui_amount"`
	UIAmountString string       `json:"ui_amount_string" bson:"ui_amount_string"`
}

// RespTokenAllTimeTrades represents token all-time trade statistics response
type RespTokenAllTimeTrades struct {
	Address        string  `json:"address" bson:"address"`
	TotalVolume    float64 `json:"total_volume" bson:"total_volume"`
	TotalVolumeUSD float64 `json:"total_volume_usd" bson:"total_volume_usd"`
	VolumeBuyUSD   float64 `json:"volume_buy_usd" bson:"volume_buy_usd"`
	VolumeSellUSD  float64 `json:"volume_sell_usd" bson:"volume_sell_usd"`
	VolumeBuy      float64 `json:"volume_buy" bson:"volume_buy"`
	VolumeSell     float64 `json:"volume_sell" bson:"volume_sell"`
	TotalTrade     int64   `json:"total_trade" bson:"total_trade"`
	Buy            int64   `json:"buy" bson:"buy"`
	Sell           int64   `json:"sell" bson:"sell"`
}

type RespMultiTokenAllTimeTrades = []RespTokenAllTimeTrades

// RespTokenHoldersItem represents token holder information
type RespTokenHoldersItem struct {
	Amount          string   `json:"amount" bson:"amount"`
	Decimals        int64    `json:"decimals" bson:"decimals"`
	Mint            string   `json:"mint" bson:"mint"`
	Owner           string   `json:"owner" bson:"owner"`
	TokenAccount    string   `json:"token_account" bson:"token_account"`
	UIAmount        float64  `json:"ui_amount" bson:"ui_amount"`
	IsScaledUIToken bool     `json:"is_scaled_ui_token" bson:"is_scaled_ui_token"`
	Multiplier      *float64 `json:"multiplier" bson:"multiplier"`
}

type RespMultiTokenHolders = []RespTokenHoldersItem

// RespNewTokenListingItem represents newly listed token information
type RespNewTokenListingItem struct {
	Address          string  `json:"address" bson:"address"`
	Symbol           string  `json:"symbol" bson:"symbol"`
	Name             string  `json:"name" bson:"name"`
	Decimals         int64   `json:"decimals" bson:"decimals"`
	Source           string  `json:"source" bson:"source"`
	LiquidityAddedAt string  `json:"liquidityAddedAt" bson:"liquidityAddedAt"`
	LogoURI          *string `json:"logoURI" bson:"logoURI"`
	Liquidity        float64 `json:"liquidity" bson:"liquidity"`
}

// RespGainerLoserItem represents gainer/loser item
type RespGainerLoserItem struct {
	Network    string  `json:"network" bson:"network"`
	Address    string  `json:"address" bson:"address"`
	PNL        float64 `json:"pnl" bson:"pnl"`
	TradeCount int64   `json:"trade_count" bson:"trade_count"`
	Volume     float64 `json:"volume" bson:"volume"`
}

// ============================================================================
// Wallet Related Types
// ============================================================================

// RespWalletPortfolioItem represents an individual token in wallet portfolio
type RespWalletPortfolioItem struct {
	Address         string   `json:"address" bson:"address"`
	Decimals        int64    `json:"decimals" bson:"decimals"`
	Balance         int64    `json:"balance" bson:"balance"`
	UIAmount        float64  `json:"uiAmount" bson:"uiAmount"`
	ChainID         string   `json:"chainId" bson:"chainId"`
	Name            string   `json:"name" bson:"name"`
	Symbol          string   `json:"symbol" bson:"symbol"`
	Icon            string   `json:"icon" bson:"icon"`
	LogoURI         string   `json:"logoURI" bson:"logoURI"`
	PriceUSD        float64  `json:"priceUsd" bson:"priceUsd"`
	ValueUSD        float64  `json:"valueUsd" bson:"valueUsd"`
	IsScaledUIToken bool     `json:"isScaledUiToken" bson:"isScaledUiToken"`
	Multiplier      *float64 `json:"multiplier" bson:"multiplier"`
}

// RespWalletTokenBalance represents wallet token balance
type RespWalletTokenBalance struct {
	Address         string   `json:"address" bson:"address"`
	Decimals        int64    `json:"decimals" bson:"decimals"`
	Balance         int64    `json:"balance" bson:"balance"`
	UIAmount        float64  `json:"uiAmount" bson:"uiAmount"`
	ChainID         string   `json:"chainId" bson:"chainId"`
	LogoURI         string   `json:"logoURI" bson:"logoURI"`
	Name            string   `json:"name" bson:"name"`
	Symbol          string   `json:"symbol" bson:"symbol"`
	PriceUSD        float64  `json:"priceUsd" bson:"priceUsd"`
	ValueUSD        float64  `json:"valueUsd" bson:"valueUsd"`
	IsScaledUIToken bool     `json:"isScaledUiToken" bson:"isScaledUiToken"`
	Multiplier      *float64 `json:"multiplier" bson:"multiplier"`
}

// RespWalletTokenFirstTx represents first token transaction in a wallet
type RespWalletTokenFirstTx struct {
	TxHash        string `json:"tx_hash" bson:"tx_hash"`
	BlockUnixTime int64  `json:"block_unix_time" bson:"block_unix_time"`
	BlockNumber   int64  `json:"block_number" bson:"block_number"`
	BalanceChange string `json:"balance_change" bson:"balance_change"`
	TokenAddress  string `json:"token_address" bson:"token_address"`
	TokenDecimals int64  `json:"token_decimals" bson:"token_decimals"`
}

// RespWalletTokensBalanceItem represents wallet token balance item
type RespWalletTokensBalanceItem struct {
	Address  string  `json:"address" bson:"address"`
	Decimals int64   `json:"decimals" bson:"decimals"`
	Price    float64 `json:"price" bson:"price"`
	Balance  string  `json:"balance" bson:"balance"`
	Amount   float64 `json:"amount" bson:"amount"`
	Network  string  `json:"network" bson:"network"`
	Name     string  `json:"name" bson:"name"`
	Symbol   string  `json:"symbol" bson:"symbol"`
	LogoURI  string  `json:"logo_uri" bson:"logo_uri"`
	Value    string  `json:"value" bson:"value"`
}

// RespWalletBalanceChangesTokenInfo represents token info in balance change response
type RespWalletBalanceChangesTokenInfo struct {
	Address         string   `json:"address" bson:"address"`
	Decimals        int64    `json:"decimals" bson:"decimals"`
	Symbol          string   `json:"symbol" bson:"symbol"`
	Name            string   `json:"name" bson:"name"`
	LogoURI         string   `json:"logo_uri" bson:"logo_uri"`
	IsScaledUIToken bool     `json:"is_scaled_ui_token" bson:"is_scaled_ui_token"`
	Multiplier      *float64 `json:"multiplier" bson:"multiplier"`
}

// RespWalletBalanceChangesItem represents wallet balance change item
type RespWalletBalanceChangesItem struct {
	Time           string                            `json:"time" bson:"time"`
	BlockNumber    int64                             `json:"block_number" bson:"block_number"`
	BlockUnixTime  int64                             `json:"block_unix_time" bson:"block_unix_time"`
	Address        string                            `json:"address" bson:"address"`
	TokenAccount   string                            `json:"token_account" bson:"token_account"`
	TxHash         string                            `json:"tx_hash" bson:"tx_hash"`
	PreBalance     string                            `json:"pre_balance" bson:"pre_balance"`
	PostBalance    string                            `json:"post_balance" bson:"post_balance"`
	Amount         string                            `json:"amount" bson:"amount"`
	TokenInfo      RespWalletBalanceChangesTokenInfo `json:"token_info" bson:"token_info"`
	Type           int64                             `json:"type" bson:"type"`
	TypeText       BalanceChangeType                 `json:"type_text" bson:"type_text"`
	ChangeType     int64                             `json:"change_type" bson:"change_type"`
	ChangeTypeText BalanceChangeDirection            `json:"change_type_text" bson:"change_type_text"`
}

// RespWalletTradesToken represents token in a wallet trade
type RespWalletTradesToken struct {
	Symbol          string   `json:"symbol" bson:"symbol"`
	Decimals        int64    `json:"decimals" bson:"decimals"`
	Address         string   `json:"address" bson:"address"`
	Amount          int64    `json:"amount" bson:"amount"`
	Type            string   `json:"type" bson:"type"`
	TypeSwap        string   `json:"type_swap" bson:"type_swap"`
	UIAmount        float64  `json:"ui_amount" bson:"ui_amount"`
	Price           float64  `json:"price" bson:"price"`
	NearestPrice    float64  `json:"nearest_price" bson:"nearest_price"`
	ChangeAmount    int64    `json:"change_amount" bson:"change_amount"`
	UIChangeAmount  float64  `json:"ui_change_amount" bson:"ui_change_amount"`
	IsScaledUIToken bool     `json:"is_scaled_ui_token" bson:"is_scaled_ui_token"`
	Multiplier      *float64 `json:"multiplier" bson:"multiplier"`
}

// RespWalletTradesItem represents a single wallet trade
type RespWalletTradesItem struct {
	Quote               RespWalletTradesToken `json:"quote" bson:"quote"`
	Base                RespWalletTradesToken `json:"base" bson:"base"`
	BasePrice           float64               `json:"base_price" bson:"base_price"`
	QuotePrice          float64               `json:"quote_price" bson:"quote_price"`
	TxHash              string                `json:"tx_hash" bson:"tx_hash"`
	Source              string                `json:"source" bson:"source"`
	BlockUnixTime       int64                 `json:"block_unix_time" bson:"block_unix_time"`
	TxType              string                `json:"tx_type" bson:"tx_type"`
	Address             string                `json:"address" bson:"address"`
	Owner               string                `json:"owner" bson:"owner"`
	BlockNumber         int64                 `json:"block_number" bson:"block_number"`
	VolumeUSD           float64               `json:"volume_usd" bson:"volume_usd"`
	Volume              float64               `json:"volume" bson:"volume"`
	InsIndex            int64                 `json:"ins_index" bson:"ins_index"`
	InnerInsIndex       int64                 `json:"inner_ins_index" bson:"inner_ins_index"`
	Signers             []string              `json:"signers" bson:"signers"`
	InteractedProgramID string                `json:"interacted_program_id" bson:"interacted_program_id"`
}

// RespWalletTrades represents wallet trades response
type RespWalletTrades struct {
	Items   []RespWalletTradesItem `json:"items" bson:"items"`
	HasNext bool                   `json:"hasNext" bson:"hasNext"`
}

// ============================================================================
// Type Aliases for Response Types
// ============================================================================

type RespSupportedNetworks = []Chain
type RespWalletSupportedNetworks = []Chain
type RespTokensPrice = map[string]RespTokenPrice
type RespPairOHLCVs = []RespPairOHLCVItem
type RespPairOHLCVsV3 = []RespPairOHLCVItemV3
type RespTokensPriceVolume = map[string]RespTokenPriceVolume
type RespPairsOverview = map[string]RespPairOverview
type RespTokensPriceStats = []RespTokenPriceStats
type RespTokensMetadata = []RespTokenMetadata
type RespTokensMarketData = map[string]RespTokenMarketData
type RespTokensTradeData = map[string]RespTokenTradeData
type RespTokenNewListing = []RespNewTokenListingItem
type RespTokenMintBurnTxs = []RespTokenMintBurnTxItem
type RespTokensAllTimeTrades = []RespTokenAllTimeTrades
type RespTokenHolders = []RespTokenHoldersItem
type RespGainerLosers = []RespGainerLoserItem
type RespWalletBalanceChanges = []RespWalletBalanceChangesItem
type RespWalletPortfolio = []RespWalletPortfolioItem
type RespWalletTokensBalances = []RespWalletTokensBalanceItem
type RespWalletsTokenFirstTx = map[string]RespWalletTokenFirstTx
type RespLatestBlockNumber = int64

// ============================================================================
// Meme Token Types
// ============================================================================

// MemeCreatedAt represents creation transaction details for meme tokens
type MemeCreatedAt struct {
	TxHash    string `json:"tx_hash" bson:"tx_hash"`
	Slot      int64  `json:"slot" bson:"slot"`
	BlockTime int64  `json:"block_time" bson:"block_time"`
}

// MemePool represents pool configuration for meme tokens
type MemePool struct {
	Address               string  `json:"address" bson:"address"`
	CurveAmount           *string `json:"curve_amount,omitempty" bson:"curve_amount,omitempty"`
	TotalSupply           *string `json:"total_supply,omitempty" bson:"total_supply,omitempty"`
	MarketcapThreshold    *string `json:"marketcap_threshold,omitempty" bson:"marketcap_threshold,omitempty"`
	CoefB                 *string `json:"coef_b,omitempty" bson:"coef_b,omitempty"`
	Bump                  *string `json:"bump,omitempty" bson:"bump,omitempty"`
	VirtualBase           *string `json:"virtual_base,omitempty" bson:"virtual_base,omitempty"`
	Creator               *string `json:"creator,omitempty" bson:"creator,omitempty"`
	BaseDecimals          *int64  `json:"base_decimals,omitempty" bson:"base_decimals,omitempty"`
	QuoteMint             *string `json:"quote_mint,omitempty" bson:"quote_mint,omitempty"`
	AuthBump              *int64  `json:"auth_bump,omitempty" bson:"auth_bump,omitempty"`
	TotalQuoteFundRaising *string `json:"total_quote_fund_raising,omitempty" bson:"total_quote_fund_raising,omitempty"`
	Supply                *string `json:"supply,omitempty" bson:"supply,omitempty"`
	PlatformFee           *string `json:"platform_fee,omitempty" bson:"platform_fee,omitempty"`
	QuoteProtocolFee      *string `json:"quote_protocol_fee,omitempty" bson:"quote_protocol_fee,omitempty"`
	TotalBaseSell         *string `json:"total_base_sell,omitempty" bson:"total_base_sell,omitempty"`
	VirtualQuote          *string `json:"virtual_quote,omitempty" bson:"virtual_quote,omitempty"`
	BaseMint              *string `json:"base_mint,omitempty" bson:"base_mint,omitempty"`
	BaseVault             *string `json:"base_vault,omitempty" bson:"base_vault,omitempty"`
	PlatformConfig        *string `json:"platform_config,omitempty" bson:"platform_config,omitempty"`
	QuoteDecimals         *int64  `json:"quote_decimals,omitempty" bson:"quote_decimals,omitempty"`
	RealQuote             *string `json:"real_quote,omitempty" bson:"real_quote,omitempty"`
	QuoteVault            *string `json:"quote_vault,omitempty" bson:"quote_vault,omitempty"`
	RealBase              *string `json:"real_base,omitempty" bson:"real_base,omitempty"`
	Status                *int64  `json:"status,omitempty" bson:"status,omitempty"`
	RealSolReserves       *string `json:"real_sol_reserves,omitempty" bson:"real_sol_reserves,omitempty"`
	RealTokenReserves     *string `json:"real_token_reserves,omitempty" bson:"real_token_reserves,omitempty"`
	TokenTotalSupply      *string `json:"token_total_supply,omitempty" bson:"token_total_supply,omitempty"`
	VirtualTokenReserves  *string `json:"virtual_token_reserves,omitempty" bson:"virtual_token_reserves,omitempty"`
}

// MemeExtensions represents token extension information for meme tokens
type MemeExtensions struct {
	Twitter     string `json:"twitter" bson:"twitter"`
	Website     string `json:"website" bson:"website"`
	Description string `json:"description" bson:"description"`
}

// MemeInfo represents meme token specific information
type MemeInfo struct {
	Source          string         `json:"source" bson:"source"`
	PlatformID      string         `json:"platform_id" bson:"platform_id"`
	Address         string         `json:"address" bson:"address"`
	CreatedAt       MemeCreatedAt  `json:"created_at" bson:"created_at"`
	CreationTime    int64          `json:"creation_time" bson:"creation_time"`
	Creator         string         `json:"creator" bson:"creator"`
	UpdatedAt       *MemeCreatedAt `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	GraduatedAt     *MemeCreatedAt `json:"graduated_at,omitempty" bson:"graduated_at,omitempty"`
	Graduated       bool           `json:"graduated" bson:"graduated"`
	GraduatedTime   *int64         `json:"graduated_time" bson:"graduated_time"`
	Pool            MemePool       `json:"pool" bson:"pool"`
	ProgressPercent float64        `json:"progress_percent" bson:"progress_percent"`
}

// RespMemeListItem represents individual meme token details
type RespMemeListItem struct {
	Address                string          `json:"address" bson:"address"`
	LogoURI                string          `json:"logo_uri" bson:"logo_uri"`
	Name                   string          `json:"name" bson:"name"`
	Symbol                 string          `json:"symbol" bson:"symbol"`
	Decimals               int64           `json:"decimals" bson:"decimals"`
	Extensions             *MemeExtensions `json:"extensions" bson:"extensions"`
	MarketCap              float64         `json:"market_cap" bson:"market_cap"`
	FDV                    float64         `json:"fdv" bson:"fdv"`
	Liquidity              float64         `json:"liquidity" bson:"liquidity"`
	LastTradeUnixTime      int64           `json:"last_trade_unix_time" bson:"last_trade_unix_time"`
	Volume1hUSD            float64         `json:"volume_1h_usd" bson:"volume_1h_usd"`
	Volume1hChangePercent  float64         `json:"volume_1h_change_percent" bson:"volume_1h_change_percent"`
	Volume2hUSD            float64         `json:"volume_2h_usd" bson:"volume_2h_usd"`
	Volume2hChangePercent  float64         `json:"volume_2h_change_percent" bson:"volume_2h_change_percent"`
	Volume4hUSD            float64         `json:"volume_4h_usd" bson:"volume_4h_usd"`
	Volume4hChangePercent  float64         `json:"volume_4h_change_percent" bson:"volume_4h_change_percent"`
	Volume8hUSD            float64         `json:"volume_8h_usd" bson:"volume_8h_usd"`
	Volume8hChangePercent  float64         `json:"volume_8h_change_percent" bson:"volume_8h_change_percent"`
	Volume24hUSD           float64         `json:"volume_24h_usd" bson:"volume_24h_usd"`
	Volume24hChangePercent *float64        `json:"volume_24h_change_percent" bson:"volume_24h_change_percent"`
	Trade1hCount           int64           `json:"trade_1h_count" bson:"trade_1h_count"`
	Trade2hCount           int64           `json:"trade_2h_count" bson:"trade_2h_count"`
	Trade4hCount           int64           `json:"trade_4h_count" bson:"trade_4h_count"`
	Trade8hCount           int64           `json:"trade_8h_count" bson:"trade_8h_count"`
	Trade24hCount          int64           `json:"trade_24h_count" bson:"trade_24h_count"`
	Price                  float64         `json:"price" bson:"price"`
	PriceChange1hPercent   float64         `json:"price_change_1h_percent" bson:"price_change_1h_percent"`
	PriceChange2hPercent   float64         `json:"price_change_2h_percent" bson:"price_change_2h_percent"`
	PriceChange4hPercent   float64         `json:"price_change_4h_percent" bson:"price_change_4h_percent"`
	PriceChange8hPercent   float64         `json:"price_change_8h_percent" bson:"price_change_8h_percent"`
	PriceChange24hPercent  float64         `json:"price_change_24h_percent" bson:"price_change_24h_percent"`
	Holder                 int64           `json:"holder" bson:"holder"`
	RecentListingTime      int64           `json:"recent_listing_time" bson:"recent_listing_time"`
	MemeInfo               MemeInfo        `json:"meme_info" bson:"meme_info"`
}

// RespMemeList represents response type for meme token list
type RespMemeList struct {
	Items   []RespMemeListItem `json:"items" bson:"items"`
	HasNext bool               `json:"has_next" bson:"has_next"`
}

// RespMemeDetail represents response type for meme token detail
type RespMemeDetail struct {
	Address           string         `json:"address" bson:"address"`
	Name              string         `json:"name" bson:"name"`
	Symbol            string         `json:"symbol" bson:"symbol"`
	Decimals          int64          `json:"decimals" bson:"decimals"`
	Extensions        MemeExtensions `json:"extensions" bson:"extensions"`
	LogoURI           string         `json:"logo_uri" bson:"logo_uri"`
	Price             float64        `json:"price" bson:"price"`
	Liquidity         float64        `json:"liquidity" bson:"liquidity"`
	CirculatingSupply int64          `json:"circulating_supply" bson:"circulating_supply"`
	MarketCap         int64          `json:"market_cap" bson:"market_cap"`
	TotalSupply       int64          `json:"total_supply" bson:"total_supply"`
	FDV               int64          `json:"fdv" bson:"fdv"`
	MemeInfo          MemeInfo       `json:"meme_info" bson:"meme_info"`
}

// ============================================================================
// Search Types
// ============================================================================

// RespSearchTokenResult represents token search result information
type RespSearchTokenResult struct {
	Name                         string   `json:"name" bson:"name"`
	Symbol                       string   `json:"symbol" bson:"symbol"`
	Address                      string   `json:"address" bson:"address"`
	Network                      string   `json:"network" bson:"network"`
	Decimals                     int64    `json:"decimals" bson:"decimals"`
	Verified                     bool     `json:"verified" bson:"verified"`
	FDV                          float64  `json:"fdv" bson:"fdv"`
	MarketCap                    float64  `json:"market_cap" bson:"market_cap"`
	Liquidity                    float64  `json:"liquidity" bson:"liquidity"`
	Price                        float64  `json:"price" bson:"price"`
	PriceChange24hPercent        float64  `json:"price_change_24h_percent" bson:"price_change_24h_percent"`
	Sell24h                      int64    `json:"sell_24h" bson:"sell_24h"`
	Sell24hChangePercent         *float64 `json:"sell_24h_change_percent" bson:"sell_24h_change_percent"`
	Buy24h                       int64    `json:"buy_24h" bson:"buy_24h"`
	Buy24hChangePercent          *float64 `json:"buy_24h_change_percent" bson:"buy_24h_change_percent"`
	UniqueWallet24h              int64    `json:"unique_wallet_24h" bson:"unique_wallet_24h"`
	UniqueWallet24hChangePercent *float64 `json:"unique_wallet_24h_change_percent" bson:"unique_wallet_24h_change_percent"`
	Trade24h                     int64    `json:"trade_24h" bson:"trade_24h"`
	Trade24hChangePercent        *float64 `json:"trade_24h_change_percent" bson:"trade_24h_change_percent"`
	Volume24hChangePercent       *float64 `json:"volume_24h_change_percent" bson:"volume_24h_change_percent"`
	Volume24hUSD                 float64  `json:"volume_24h_usd" bson:"volume_24h_usd"`
	LastTradeUnixTime            int64    `json:"last_trade_unix_time" bson:"last_trade_unix_time"`
	LastTradeHumanTime           string   `json:"last_trade_human_time" bson:"last_trade_human_time"`
	UpdatedTime                  int64    `json:"updated_time" bson:"updated_time"`
	CreationTime                 string   `json:"creation_time" bson:"creation_time"`
	IsScaledUIToken              bool     `json:"is_scaled_ui_token" bson:"is_scaled_ui_token"`
	Multiplier                   *float64 `json:"multiplier" bson:"multiplier"`
}

// RespSearchItem represents search result containing token and market results
type RespSearchItem struct {
	Type   string                  `json:"type" bson:"type"`
	Result []RespSearchTokenResult `json:"result" bson:"result"`
}

type RespSearchItems = []RespSearchItem

// ============================================================================
// Wallet Net Worth Types
// ============================================================================

// RespWalletNetWorthItem represents token details in wallet net worth response
type RespWalletNetWorthItem struct {
	Address  string  `json:"address" bson:"address"`
	Decimals int64   `json:"decimals" bson:"decimals"`
	Price    float64 `json:"price" bson:"price"`
	Balance  string  `json:"balance" bson:"balance"`
	Amount   float64 `json:"amount" bson:"amount"`
	Network  string  `json:"network" bson:"network"`
	Name     string  `json:"name" bson:"name"`
	Symbol   string  `json:"symbol" bson:"symbol"`
	LogoURI  string  `json:"logo_uri" bson:"logo_uri"`
	Value    string  `json:"value" bson:"value"`
}

// RespWalletNetWorthPagination represents pagination details
type RespWalletNetWorthPagination struct {
	Limit  int64 `json:"limit" bson:"limit"`
	Offset int64 `json:"offset" bson:"offset"`
	Total  int64 `json:"total" bson:"total"`
}

// RespWalletNetWorth represents wallet net worth endpoint response
type RespWalletNetWorth struct {
	WalletAddress    string                       `json:"wallet_address" bson:"wallet_address"`
	Currency         string                       `json:"currency" bson:"currency"`
	TotalValue       string                       `json:"total_value" bson:"total_value"`
	CurrentTimestamp string                       `json:"current_timestamp" bson:"current_timestamp"`
	Items            []RespWalletNetWorthItem     `json:"items" bson:"items"`
	Pagination       RespWalletNetWorthPagination `json:"pagination" bson:"pagination"`
}

// RespWalletNetWorthHistoryItem represents individual net worth history data point
type RespWalletNetWorthHistoryItem struct {
	Timestamp             string  `json:"timestamp" bson:"timestamp"`
	NetWorth              float64 `json:"net_worth" bson:"net_worth"`
	NetWorthChange        float64 `json:"net_worth_change" bson:"net_worth_change"`
	NetWorthChangePercent float64 `json:"net_worth_change_percent" bson:"net_worth_change_percent"`
}

// RespWalletNetWorthHistories represents wallet net worth history response
type RespWalletNetWorthHistories struct {
	WalletAddress    string                          `json:"wallet_address" bson:"wallet_address"`
	Currency         string                          `json:"currency" bson:"currency"`
	CurrentTimestamp string                          `json:"current_timestamp" bson:"current_timestamp"`
	PastTimestamp    string                          `json:"past_timestamp" bson:"past_timestamp"`
	History          []RespWalletNetWorthHistoryItem `json:"history" bson:"history"`
}

// RespWalletNetWorthDetailsNetAsset represents details of an individual asset in the wallet
type RespWalletNetWorthDetailsNetAsset struct {
	Symbol       string  `json:"symbol" bson:"symbol"`
	TokenAddress string  `json:"token_address" bson:"token_address"`
	Decimal      int64   `json:"decimal" bson:"decimal"`
	Balance      string  `json:"balance" bson:"balance"`
	Price        float64 `json:"price" bson:"price"`
	Value        float64 `json:"value" bson:"value"`
}

// RespWalletNetWorthDetails represents wallet net worth details response
type RespWalletNetWorthDetails struct {
	WalletAddress      string                              `json:"wallet_address" bson:"wallet_address"`
	Currency           string                              `json:"currency" bson:"currency"`
	NetWorth           float64                             `json:"net_worth" bson:"net_worth"`
	RequestedTimestamp string                              `json:"requested_timestamp" bson:"requested_timestamp"`
	ResolvedTimestamp  string                              `json:"resolved_timestamp" bson:"resolved_timestamp"`
	NetAssets          []RespWalletNetWorthDetailsNetAsset `json:"net_assets" bson:"net_assets"`
}

// ============================================================================
// Wallet PnL Types
// ============================================================================

// WalletPnLTokenCounts represents trade count statistics
type WalletPnLTokenCounts struct {
	TotalBuy   int64 `json:"total_buy" bson:"total_buy"`
	TotalSell  int64 `json:"total_sell" bson:"total_sell"`
	TotalTrade int64 `json:"total_trade" bson:"total_trade"`
}

// WalletPnLTokenQuantity represents token quantity statistics
type WalletPnLTokenQuantity struct {
	TotalBoughtAmount float64 `json:"total_bought_amount" bson:"total_bought_amount"`
	TotalSoldAmount   float64 `json:"total_sold_amount" bson:"total_sold_amount"`
	Holding           float64 `json:"holding" bson:"holding"`
}

// WalletPnLTokenCashflow represents USD cashflow statistics
type WalletPnLTokenCashflow struct {
	CostOfQuantitySold float64 `json:"cost_of_quantity_sold" bson:"cost_of_quantity_sold"`
	TotalInvested      float64 `json:"total_invested" bson:"total_invested"`
	TotalSold          float64 `json:"total_sold" bson:"total_sold"`
	CurrentValue       float64 `json:"current_value" bson:"current_value"`
}

// WalletPnLTokenPnL represents profit/loss metrics
type WalletPnLTokenPnL struct {
	RealizedProfitUSD     float64 `json:"realized_profit_usd" bson:"realized_profit_usd"`
	RealizedProfitPercent float64 `json:"realized_profit_percent" bson:"realized_profit_percent"`
	UnrealizedUSD         float64 `json:"unrealized_usd" bson:"unrealized_usd"`
	UnrealizedPercent     float64 `json:"unrealized_percent" bson:"unrealized_percent"`
	TotalUSD              float64 `json:"total_usd" bson:"total_usd"`
	TotalPercent          float64 `json:"total_percent" bson:"total_percent"`
	AvgProfitPerTradeUSD  float64 `json:"avg_profit_per_trade_usd" bson:"avg_profit_per_trade_usd"`
}

// WalletPnLTokenPricing represents token pricing data
type WalletPnLTokenPricing struct {
	CurrentPrice float64  `json:"current_price" bson:"current_price"`
	AvgBuyCost   float64  `json:"avg_buy_cost" bson:"avg_buy_cost"`
	AvgSellCost  *float64 `json:"avg_sell_cost" bson:"avg_sell_cost"`
}

// WalletPnLTokenStats represents per-token statistics
type WalletPnLTokenStats struct {
	Symbol      string                 `json:"symbol" bson:"symbol"`
	Decimals    int64                  `json:"decimals" bson:"decimals"`
	Counts      WalletPnLTokenCounts   `json:"counts" bson:"counts"`
	Quantity    WalletPnLTokenQuantity `json:"quantity" bson:"quantity"`
	CashflowUSD WalletPnLTokenCashflow `json:"cashflow_usd" bson:"cashflow_usd"`
	PnL         WalletPnLTokenPnL      `json:"pnl" bson:"pnl"`
	Pricing     WalletPnLTokenPricing  `json:"pricing" bson:"pricing"`
}

// WalletPnLMeta represents metadata about the PnL request
type WalletPnLMeta struct {
	Address      string `json:"address" bson:"address"`
	Currency     string `json:"currency" bson:"currency"`
	HoldingCheck bool   `json:"holding_check" bson:"holding_check"`
	Time         string `json:"time" bson:"time"`
}

// RespWalletTokensPnL represents wallet tokens PnL endpoint response
type RespWalletTokensPnL struct {
	Meta   WalletPnLMeta                  `json:"meta" bson:"meta"`
	Tokens map[string]WalletPnLTokenStats `json:"tokens" bson:"tokens"`
}

// WalletsPnLByTokenMetadata represents token metadata information
type WalletsPnLByTokenMetadata struct {
	Symbol   string `json:"symbol" bson:"symbol"`
	Decimals int64  `json:"decimals" bson:"decimals"`
}

// RespWalletsPnLByToken represents wallet PnL by token endpoint response
type RespWalletsPnLByToken struct {
	TokenMetadata WalletsPnLByTokenMetadata      `json:"token_metadata" bson:"token_metadata"`
	Data          map[string]WalletPnLTokenStats `json:"data" bson:"data"`
}

// ============================================================================
// Wallet Transaction Types
// ============================================================================

// WalletTxBalanceChange represents token balance change details
type WalletTxBalanceChange struct {
	Amount          int64    `json:"amount" bson:"amount"`
	Symbol          string   `json:"symbol" bson:"symbol"`
	Name            string   `json:"name" bson:"name"`
	Decimals        int64    `json:"decimals" bson:"decimals"`
	Address         string   `json:"address" bson:"address"`
	LogoURI         string   `json:"logoURI" bson:"logoURI"`
	IsScaledUIToken bool     `json:"isScaledUiToken" bson:"isScaledUiToken"`
	Multiplier      *float64 `json:"multiplier" bson:"multiplier"`
}

// WalletTxContractLabelMetadata represents contract metadata
type WalletTxContractLabelMetadata struct {
	Icon string `json:"icon" bson:"icon"`
}

// WalletTxContractLabel represents contract label information
type WalletTxContractLabel struct {
	Address  string                        `json:"address" bson:"address"`
	Name     string                        `json:"name" bson:"name"`
	Metadata WalletTxContractLabelMetadata `json:"metadata" bson:"metadata"`
}

// WalletTxTokenTransfer represents token transfer details
type WalletTxTokenTransfer struct {
	FromTokenAccount string   `json:"fromTokenAccount" bson:"fromTokenAccount"`
	ToTokenAccount   string   `json:"toTokenAccount" bson:"toTokenAccount"`
	FromUserAccount  string   `json:"fromUserAccount" bson:"fromUserAccount"`
	ToUserAccount    string   `json:"toUserAccount" bson:"toUserAccount"`
	TokenAmount      float64  `json:"tokenAmount" bson:"tokenAmount"`
	Mint             string   `json:"mint" bson:"mint"`
	TransferNative   bool     `json:"transferNative" bson:"transferNative"`
	IsScaledUIToken  bool     `json:"isScaledUiToken" bson:"isScaledUiToken"`
	Multiplier       *float64 `json:"multiplier" bson:"multiplier"`
}

// RespWalletTx represents wallet transaction details response
type RespWalletTx struct {
	TxHash         string                  `json:"txHash" bson:"txHash"`
	BlockNumber    int64                   `json:"blockNumber" bson:"blockNumber"`
	BlockTime      string                  `json:"blockTime" bson:"blockTime"`
	Status         bool                    `json:"status" bson:"status"`
	From           string                  `json:"from" bson:"from"`
	To             string                  `json:"to" bson:"to"`
	Fee            int64                   `json:"fee" bson:"fee"`
	MainAction     string                  `json:"mainAction" bson:"mainAction"`
	BalanceChange  []WalletTxBalanceChange `json:"balanceChange" bson:"balanceChange"`
	ContractLabel  WalletTxContractLabel   `json:"contractLabel" bson:"contractLabel"`
	TokenTransfers []WalletTxTokenTransfer `json:"tokenTransfers" bson:"tokenTransfers"`
}

// RespWalletTxs represents wallet transactions by chain
type RespWalletTxs = map[Chain][]RespWalletTx

// ============================================================================
// Token Additional Types
// ============================================================================

// RespTokenTopTraderItem represents top trader details for a token
type RespTokenTopTraderItem struct {
	TokenAddress    string   `json:"tokenAddress" bson:"tokenAddress"`
	Owner           string   `json:"owner" bson:"owner"`
	Tags            []string `json:"tags" bson:"tags"`
	Type            string   `json:"type" bson:"type"`
	Volume          float64  `json:"volume" bson:"volume"`
	Trade           int64    `json:"trade" bson:"trade"`
	TradeBuy        int64    `json:"tradeBuy" bson:"tradeBuy"`
	TradeSell       int64    `json:"tradeSell" bson:"tradeSell"`
	VolumeBuy       float64  `json:"volumeBuy" bson:"volumeBuy"`
	VolumeSell      float64  `json:"volumeSell" bson:"volumeSell"`
	IsScaledUIToken bool     `json:"isScaledUiToken" bson:"isScaledUiToken"`
	Multiplier      *float64 `json:"multiplier" bson:"multiplier"`
}

// RespTokenTopTraders is a list of top traders
type RespTokenTopTraders = []RespTokenTopTraderItem

// RespTokenAllMarketListTokenInfo represents token information in market list
type RespTokenAllMarketListTokenInfo struct {
	Address  string `json:"address" bson:"address"`
	Decimals int64  `json:"decimals" bson:"decimals"`
	Symbol   string `json:"symbol" bson:"symbol"`
	Icon     string `json:"icon" bson:"icon"`
}

// RespTokenAllMarketListItem represents individual market item
type RespTokenAllMarketListItem struct {
	Address                      string                          `json:"address" bson:"address"`
	Base                         RespTokenAllMarketListTokenInfo `json:"base" bson:"base"`
	CreatedAt                    string                          `json:"createdAt" bson:"createdAt"`
	Liquidity                    float64                         `json:"liquidity" bson:"liquidity"`
	Name                         string                          `json:"name" bson:"name"`
	Price                        *float64                        `json:"price" bson:"price"`
	Quote                        RespTokenAllMarketListTokenInfo `json:"quote" bson:"quote"`
	Source                       string                          `json:"source" bson:"source"`
	Trade24h                     int64                           `json:"trade24h" bson:"trade24h"`
	Trade24hChangePercent        float64                         `json:"trade24hChangePercent" bson:"trade24hChangePercent"`
	UniqueWallet24h              int64                           `json:"uniqueWallet24h" bson:"uniqueWallet24h"`
	UniqueWallet24hChangePercent float64                         `json:"uniqueWallet24hChangePercent" bson:"uniqueWallet24hChangePercent"`
	Volume24h                    float64                         `json:"volume24h" bson:"volume24h"`
}

// RespTokenAllMarketList represents token all market list response
type RespTokenAllMarketList struct {
	Items []RespTokenAllMarketListItem `json:"items" bson:"items"`
	Total int64                        `json:"total" bson:"total"`
}

// RespTrendingToken represents individual trending token details
type RespTrendingToken struct {
	Address                string  `json:"address" bson:"address"`
	Decimals               int64   `json:"decimals" bson:"decimals"`
	Liquidity              float64 `json:"liquidity" bson:"liquidity"`
	LogoURI                string  `json:"logoURI" bson:"logoURI"`
	Name                   string  `json:"name" bson:"name"`
	Symbol                 string  `json:"symbol" bson:"symbol"`
	Volume24hUSD           float64 `json:"volume24hUSD" bson:"volume24hUSD"`
	Volume24hChangePercent float64 `json:"volume24hChangePercent" bson:"volume24hChangePercent"`
	Rank                   int64   `json:"rank" bson:"rank"`
	Price                  float64 `json:"price" bson:"price"`
	Price24hChangePercent  float64 `json:"price24hChangePercent" bson:"price24hChangePercent"`
	FDV                    float64 `json:"fdv" bson:"fdv"`
	MarketCap              float64 `json:"marketcap" bson:"marketcap"`
}

// RespTokenTrendingList represents trending token list response
type RespTokenTrendingList struct {
	UpdateUnixTime int64               `json:"updateUnixTime" bson:"updateUnixTime"`
	UpdateTime     string              `json:"updateTime" bson:"updateTime"`
	Tokens         []RespTrendingToken `json:"tokens" bson:"tokens"`
	Total          int64               `json:"total" bson:"total"`
}

// RespTokenHolderBatchItem represents token holder batch response item
type RespTokenHolderBatchItem struct {
	Balance  string  `json:"balance" bson:"balance"`
	Decimals int64   `json:"decimals" bson:"decimals"`
	Mint     string  `json:"mint" bson:"mint"`
	Owner    string  `json:"owner" bson:"owner"`
	Amount   float64 `json:"amount" bson:"amount"`
}

// RespTokenHolderBatch is a list of token holder batch items
type RespTokenHolderBatch = []RespTokenHolderBatchItem

// RespTokenExitLiquidityPrice represents price information in exit liquidity response
type RespTokenExitLiquidityPrice struct {
	Value           float64 `json:"value" bson:"value"`
	UpdateUnixTime  int64   `json:"update_unix_time" bson:"update_unix_time"`
	UpdateHumanTime string  `json:"update_human_time" bson:"update_human_time"`
	UpdateInSlot    int64   `json:"update_in_slot" bson:"update_in_slot"`
}

// RespTokenExitLiquidity represents token exit liquidity information
type RespTokenExitLiquidity struct {
	Token         string                      `json:"token" bson:"token"`
	ExitLiquidity float64                     `json:"exit_liquidity" bson:"exit_liquidity"`
	Liquidity     float64                     `json:"liquidity" bson:"liquidity"`
	Price         RespTokenExitLiquidityPrice `json:"price" bson:"price"`
	Currency      string                      `json:"currency" bson:"currency"`
	Address       string                      `json:"address" bson:"address"`
	Name          string                      `json:"name" bson:"name"`
	Symbol        string                      `json:"symbol" bson:"symbol"`
	Decimals      int64                       `json:"decimals" bson:"decimals"`
	Extensions    TokenExtensions             `json:"extensions" bson:"extensions"`
	LogoURI       string                      `json:"logo_uri" bson:"logo_uri"`
}

// RespTokensExitLiquidity is a list of exit liquidity items
type RespTokensExitLiquidity = []RespTokenExitLiquidity
