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
	IsScaledUiToken bool    `json:"isScaledUiToken"`
	Value           float64 `json:"value"`
	UpdateUnixTime  int64   `json:"updateUnixTime"`
	UpdateHumanTime string  `json:"updateHumanTime"`
	PriceChange24h  float64 `json:"priceChange24h"`
	PriceInNative   float64 `json:"priceInNative"`
	Liquidity       float64 `json:"liquidity"`
}

// ============================================================================
// Token Transaction Types
// ============================================================================

// RespTokenTradeToken represents token details in a trade response
type RespTokenTradeToken struct {
	Symbol          string   `json:"symbol"`
	Decimals        int64    `json:"decimals"`
	Address         string   `json:"address"`
	Amount          any      `json:"amount"`
	UIAmount        float64  `json:"uiAmount"`
	Price           float64  `json:"price"`
	NearestPrice    float64  `json:"nearestPrice"`
	ChangeAmount    int64    `json:"changeAmount"`
	UIChangeAmount  float64  `json:"uiChangeAmount"`
	IsScaledUIToken bool     `json:"isScaledUiToken"`
	Multiplier      *float64 `json:"multiplier"`
}

// RespTokenTxsItem represents a single token transaction item
type RespTokenTxsItem struct {
	Quote         RespTokenTradeToken `json:"quote"`
	Base          RespTokenTradeToken `json:"base"`
	BasePrice     float64             `json:"basePrice"`
	QuotePrice    float64             `json:"quotePrice"`
	TxHash        string              `json:"txHash"`
	Source        string              `json:"source"`
	BlockUnixTime int64               `json:"blockUnixTime"`
	TxType        TxType              `json:"txType"`
	Owner         string              `json:"owner"`
	Side          TradeSide           `json:"side"`
	Alias         *string             `json:"alias"`
	PricePair     float64             `json:"pricePair"`
	From          RespTokenTradeToken `json:"from"`
	To            RespTokenTradeToken `json:"to"`
	TokenPrice    float64             `json:"tokenPrice"`
	PoolID        string              `json:"poolId"`
}

// RespTokenTxs represents the response for token transactions
type RespTokenTxs struct {
	Items   []RespTokenTxsItem `json:"items"`
	HasNext bool               `json:"hasNext"`
}

// ============================================================================
// Pair Transaction Types
// ============================================================================

// TokenTradeToken represents token trade information in a transaction
type TokenTradeToken struct {
	Symbol   string `json:"symbol"`
	Decimals int64  `json:"decimals"`
	Address  string `json:"address"`
	// Amount can be either int64 or string depending on the API response format.
	// Some endpoints return numeric values, others return string representations.
	// Use type assertion to convert to the desired type:
	//   - For int64: amount, ok := token.Amount.(int64)
	//   - For string: amount, ok := token.Amount.(string)
	//   - For string to int64 conversion: str, ok := token.Amount.(string); if ok { intVal, _ := strconv.ParseInt(str, 10, 64) }
	Amount          any      `json:"amount"`
	Type            string   `json:"type"`
	TypeSwap        string   `json:"typeSwap"`
	UIAmount        float64  `json:"uiAmount"`
	Price           float64  `json:"price"`
	NearestPrice    float64  `json:"nearestPrice"`
	ChangeAmount    int64    `json:"changeAmount"`
	UIChangeAmount  float64  `json:"uiChangeAmount"`
	IsScaledUIToken bool     `json:"isScaledUiToken"`
	Multiplier      *float64 `json:"multiplier"`
}

// RespPairTxsItem represents a single pair transaction item
type RespPairTxsItem struct {
	TxHash        string          `json:"txHash"`
	Source        string          `json:"source"`
	BlockUnixTime int64           `json:"blockUnixTime"`
	TxType        TxType          `json:"txType"`
	Address       string          `json:"address"`
	Owner         string          `json:"owner"`
	From          TokenTradeToken `json:"from"`
	To            TokenTradeToken `json:"to"`
}

// RespPairTxs represents the response for pair transactions
type RespPairTxs struct {
	Items   []RespPairTxsItem `json:"items"`
	HasNext bool              `json:"hasNext"`
}

// RespTokenTxsByTime represents token transactions by time
type RespTokenTxsByTime struct {
	Items   []RespTokenTxsItem `json:"items"`
	HasNext bool               `json:"hasNext"`
}

// RespPairTxsByTime represents pair transactions by time
type RespPairTxsByTime struct {
	Items   []RespPairTxsItem `json:"items"`
	HasNext bool              `json:"hasNext"`
}

// ============================================================================
// V3 Transaction Types
// ============================================================================

// RespAllTxsTokenV3 represents token details in an all transactions v3 response
type RespAllTxsTokenV3 struct {
	Symbol          string   `json:"symbol"`
	Address         string   `json:"address"`
	Decimals        int64    `json:"decimals"`
	Price           float64  `json:"price"`
	Amount          string   `json:"amount"`
	UIAmount        float64  `json:"ui_amount"`
	UIChangeAmount  float64  `json:"ui_change_amount"`
	TypeSwap        TypeSwap `json:"type_swap"`
	IsScaledUIToken bool     `json:"is_scaled_ui_token"`
	Multiplier      *float64 `json:"multiplier"`
}

// RespAllTxsItemV3 represents a single transaction in all transactions v3
type RespAllTxsItemV3 struct {
	Base                RespAllTxsTokenV3 `json:"base"`
	Quote               RespAllTxsTokenV3 `json:"quote"`
	TxType              TxType            `json:"tx_type"`
	TxHash              string            `json:"tx_hash"`
	InsIndex            int64             `json:"ins_index"`
	InnerInsIndex       int64             `json:"inner_ins_index"`
	BlockUnixTime       int64             `json:"block_unix_time"`
	BlockNumber         int64             `json:"block_number"`
	VolumeUSD           float64           `json:"volume_usd"`
	Volume              float64           `json:"volume"`
	Owner               string            `json:"owner"`
	Signers             []string          `json:"signers"`
	Source              string            `json:"source"`
	InteractedProgramID string            `json:"interacted_program_id"`
	PoolID              string            `json:"pool_id"`
}

// RespAllTxsV3 represents the response for all transactions v3
type RespAllTxsV3 struct {
	Items   []RespAllTxsItemV3 `json:"items"`
	HasNext bool               `json:"hasNext"`
}

// TokenInfo represents token information in a V3 transaction
type TokenInfo struct {
	Symbol         string  `json:"symbol"`
	Address        string  `json:"address"`
	Decimals       int64   `json:"decimals"`
	Price          float64 `json:"price"`
	Amount         string  `json:"amount"`
	UIAmount       float64 `json:"ui_amount"`
	UIChangeAmount float64 `json:"ui_change_amount"`
}

// RespTokenTxsItemV3 represents a single token transaction in V3 API
type RespTokenTxsItemV3 struct {
	TxType        TxType    `json:"tx_type"`
	TxHash        string    `json:"tx_hash"`
	InsIndex      int64     `json:"ins_index"`
	InnerInsIndex int64     `json:"inner_ins_index"`
	BlockUnixTime int64     `json:"block_unix_time"`
	BlockNumber   int64     `json:"block_number"`
	VolumeUSD     float64   `json:"volume_usd"`
	Volume        float64   `json:"volume"`
	Owner         string    `json:"owner"`
	Signers       []string  `json:"signers"`
	Source        string    `json:"source"`
	Side          TradeSide `json:"side"`
	Alias         *string   `json:"alias"`
	PricePair     float64   `json:"price_pair"`
	From          TokenInfo `json:"from"`
	To            TokenInfo `json:"to"`
	PoolID        string    `json:"pool_id"`
}

// RespTokenTxsV3 represents the response for token transactions V3
type RespTokenTxsV3 struct {
	Items   []RespTokenTxsItemV3 `json:"items"`
	HasNext bool                 `json:"hasNext"`
}

// RespRecentTxsTokenV3 represents token details in a recent transactions v3 response
type RespRecentTxsTokenV3 struct {
	Symbol          string   `json:"symbol"`
	Address         string   `json:"address"`
	Decimals        int64    `json:"decimals"`
	Price           float64  `json:"price"`
	Amount          string   `json:"amount"`
	UIAmount        float64  `json:"ui_amount"`
	UIChangeAmount  float64  `json:"ui_change_amount"`
	TypeSwap        TypeSwap `json:"type_swap"`
	IsScaledUIToken bool     `json:"is_scaled_ui_token"`
	Multiplier      *float64 `json:"multiplier"`
}

// RespRecentTxsItemV3 represents a single recent transaction in V3
type RespRecentTxsItemV3 struct {
	Base                RespRecentTxsTokenV3 `json:"base"`
	Quote               RespRecentTxsTokenV3 `json:"quote"`
	TxType              TxType               `json:"tx_type"`
	TxHash              string               `json:"tx_hash"`
	InsIndex            int64                `json:"ins_index"`
	InnerInsIndex       int64                `json:"inner_ins_index"`
	BlockUnixTime       int64                `json:"block_unix_time"`
	BlockNumber         int64                `json:"block_number"`
	VolumeUSD           float64              `json:"volume_usd"`
	Volume              float64              `json:"volume"`
	Owner               string               `json:"owner"`
	Signers             []string             `json:"signers"`
	Source              string               `json:"source"`
	InteractedProgramID string               `json:"interacted_program_id"`
	PoolID              string               `json:"pool_id"`
}

// RespRecentTxsV3 represents the response for recent transactions V3
type RespRecentTxsV3 struct {
	Items   []RespRecentTxsItemV3 `json:"items"`
	HasNext bool                  `json:"hasNext"`
}

// ============================================================================
// OHLCV Types
// ============================================================================

// RespTokenOHLCVItem represents an OHLCV data point for a token
type RespTokenOHLCVItem struct {
	O        float64      `json:"o"`        // Open price
	H        float64      `json:"h"`        // High price
	L        float64      `json:"l"`        // Low price
	C        float64      `json:"c"`        // Close price
	V        float64      `json:"v"`        // Volume
	UnixTime int64        `json:"unixTime"` // Unix timestamp
	Address  string       `json:"address"`
	Type     TimeInterval `json:"type"`     // Interval type
	Currency string       `json:"currency"` // Price currency
}

// RespTokenOHLCVs represents OHLCV response for a token
type RespTokenOHLCVs struct {
	IsScaledUIToken bool                 `json:"isScaledUiToken"`
	Items           []RespTokenOHLCVItem `json:"items"`
}

// RespPairOHLCVItem represents an OHLCV data point for a trading pair
type RespPairOHLCVItem struct {
	Address  string       `json:"address"`
	C        float64      `json:"c"`        // Close price
	H        float64      `json:"h"`        // High price
	L        float64      `json:"l"`        // Low price
	O        float64      `json:"o"`        // Open price
	Type     TimeInterval `json:"type"`     // Time interval
	UnixTime int64        `json:"unixTime"` // Unix timestamp
	V        float64      `json:"v"`        // Volume
}

// RespOHLCVBaseQuoteItem represents an OHLCV data point for base/quote token pair
type RespOHLCVBaseQuoteItem struct {
	O        float64 `json:"o"`        // Open price
	C        float64 `json:"c"`        // Close price
	H        float64 `json:"h"`        // High price
	L        float64 `json:"l"`        // Low price
	VBase    float64 `json:"vBase"`    // Volume in base token
	VQuote   float64 `json:"vQuote"`   // Volume in quote token
	UnixTime int64   `json:"unixTime"` // Unix timestamp in seconds
}

// RespOHLCVBaseQuote represents OHLCV response for a base/quote token pair
type RespOHLCVBaseQuote struct {
	Items                []RespOHLCVBaseQuoteItem `json:"items"`
	BaseAddress          string                   `json:"baseAddress"`
	QuoteAddress         string                   `json:"quoteAddress"`
	IsScaledUITokenBase  bool                     `json:"isScaledUiTokenBase"`
	IsScaledUITokenQuote bool                     `json:"isScaledUiTokenQuote"`
	Type                 TimeInterval             `json:"type"`
}

// RespTokenOHLCVItemV3 represents an OHLCV V3 data point for a token
type RespTokenOHLCVItemV3 struct {
	O        float64      `json:"o"`         // Open price
	H        float64      `json:"h"`         // High price
	L        float64      `json:"l"`         // Low price
	C        float64      `json:"c"`         // Close price
	V        float64      `json:"v"`         // Volume
	VUSD     float64      `json:"v_usd"`     // Volume in USD
	UnixTime int64        `json:"unix_time"` // Unix timestamp
	Address  string       `json:"address"`
	Type     TimeInterval `json:"type"`     // Interval type
	Currency string       `json:"currency"` // Price currency
}

// RespTokenOHLCVsV3 represents OHLCV V3 response for a token
type RespTokenOHLCVsV3 struct {
	IsScaledUIToken bool                   `json:"is_scaled_ui_token"`
	Items           []RespTokenOHLCVItemV3 `json:"items"`
}

// RespPairOHLCVItemV3 represents an OHLCV V3 data point for a trading pair
type RespPairOHLCVItemV3 struct {
	Address  string       `json:"address"`
	H        float64      `json:"h"`         // High price
	O        float64      `json:"o"`         // Open price
	L        float64      `json:"l"`         // Low price
	C        float64      `json:"c"`         // Close price
	Type     TimeInterval `json:"type"`      // Time interval
	V        float64      `json:"v"`         // Volume
	UnixTime int64        `json:"unix_time"` // Unix timestamp
	VUSD     float64      `json:"v_usd"`     // Volume in USD
}

// ============================================================================
// Price History and Statistics Types
// ============================================================================

// RespPriceHistoryItem represents an individual price history data point
type RespPriceHistoryItem struct {
	UnixTime int64   `json:"unixTime"`
	Value    float64 `json:"value"`
}

// RespTokenPriceHistories represents the response for token price history
type RespTokenPriceHistories struct {
	IsScaledUIToken bool                   `json:"isScaledUiToken"`
	Items           []RespPriceHistoryItem `json:"items"`
}

// RespTokenPriceHistoryByTime represents price history for a token at a specific time
type RespTokenPriceHistoryByTime struct {
	IsScaledUIToken bool    `json:"isScaledUiToken"`
	Value           float64 `json:"value"`
	UpdateUnixTime  int64   `json:"updateUnixTime"` // Unix time in seconds
	PriceChange24h  float64 `json:"priceChange24h"`
}

// RespTokenPriceVolume represents token price and volume data
type RespTokenPriceVolume struct {
	IsScaledUIToken     bool    `json:"isScaledUiToken"`
	Price               float64 `json:"price"`
	UpdateUnixTime      int64   `json:"updateUnixTime"` // Unix time in seconds
	UpdateHumanTime     string  `json:"updateHumanTime"`
	VolumeUSD           float64 `json:"volumeUSD"`
	VolumeChangePercent float64 `json:"volumeChangePercent"`
	PriceChangePercent  float64 `json:"priceChangePercent"`
}

// PriceStatsData represents price statistics data point
type PriceStatsData struct {
	UnixTimeUpdatePrice int64     `json:"unix_time_update_price"`
	TimeFrame           TimeFrame `json:"time_frame"`
	Price               float64   `json:"price"`
	PriceChangePercent  float64   `json:"price_change_percent"`
	High                float64   `json:"high"`
	Low                 float64   `json:"low"`
}

// RespTokenPriceStats represents price statistics response for a token
type RespTokenPriceStats struct {
	Address         string           `json:"address"`
	IsScaledUIToken bool             `json:"is_scaled_ui_token"`
	Data            []PriceStatsData `json:"data"`
}

type RespMultiTokenPriceStats = []RespTokenPriceStats

// ============================================================================
// Pair Overview Types
// ============================================================================

// TokenInfoInPair represents token information in pair overview
type TokenInfoInPair struct {
	Address         string  `json:"address"`
	Decimals        int64   `json:"decimals"`
	Icon            string  `json:"icon"`
	Symbol          string  `json:"symbol"`
	IsScaledUIToken bool    `json:"is_scaled_ui_token"`
	Multiplier      float64 `json:"multiplier"`
}

// RespPairOverview represents overview data for a trading pair
type RespPairOverview struct {
	Address                      string          `json:"address"`
	Base                         TokenInfoInPair `json:"base"`
	Quote                        TokenInfoInPair `json:"quote"`
	Name                         string          `json:"name"`
	Source                       string          `json:"source"`
	CreatedAt                    string          `json:"created_at"`
	Liquidity                    float64         `json:"liquidity"`
	LiquidityChangePercentage24h *float64        `json:"liquidity_change_percentage_24h"`
	Price                        float64         `json:"price"`
	Trade24h                     int64           `json:"trade_24h"`
	Trade12h                     int64           `json:"trade_12h"`
	Trade8h                      int64           `json:"trade_8h"`
	Trade4h                      int64           `json:"trade_4h"`
	Trade2h                      int64           `json:"trade_2h"`
	Trade1h                      int64           `json:"trade_1h"`
	Trade30m                     int64           `json:"trade_30m"`
	Trade24hChangePercent        float64         `json:"trade_24h_change_percent"`
	Trade12hChangePercent        float64         `json:"trade_12h_change_percent"`
	Trade8hChangePercent         float64         `json:"trade_8h_change_percent"`
	Trade4hChangePercent         float64         `json:"trade_4h_change_percent"`
	Trade2hChangePercent         float64         `json:"trade_2h_change_percent"`
	Trade1hChangePercent         float64         `json:"trade_1h_change_percent"`
	Trade30mChangePercent        float64         `json:"trade_30m_change_percent"`
	TradeHistory24h              int64           `json:"trade_history_24h"`
	TradeHistory12h              int64           `json:"trade_history_12h"`
	TradeHistory8h               int64           `json:"trade_history_8h"`
	TradeHistory4h               int64           `json:"trade_history_4h"`
	TradeHistory2h               int64           `json:"trade_history_2h"`
	TradeHistory1h               int64           `json:"trade_history_1h"`
	TradeHistory30m              int64           `json:"trade_history_30m"`
	UniqueWallet24h              int64           `json:"unique_wallet_24h"`
	UniqueWallet12h              int64           `json:"unique_wallet_12h"`
	UniqueWallet8h               int64           `json:"unique_wallet_8h"`
	UniqueWallet4h               int64           `json:"unique_wallet_4h"`
	UniqueWallet2h               int64           `json:"unique_wallet_2h"`
	UniqueWallet1h               int64           `json:"unique_wallet_1h"`
	UniqueWallet30m              int64           `json:"unique_wallet_30m"`
	UniqueWallet24hChangePercent float64         `json:"unique_wallet_24h_change_percent"`
	UniqueWallet12hChangePercent float64         `json:"unique_wallet_12h_change_percent"`
	UniqueWallet8hChangePercent  float64         `json:"unique_wallet_8h_change_percent"`
	UniqueWallet4hChangePercent  float64         `json:"unique_wallet_4h_change_percent"`
	UniqueWallet2hChangePercent  float64         `json:"unique_wallet_2h_change_percent"`
	UniqueWallet1hChangePercent  float64         `json:"unique_wallet_1h_change_percent"`
	UniqueWallet30mChangePercent float64         `json:"unique_wallet_30m_change_percent"`
	Volume24h                    float64         `json:"volume_24h"`
	Volume12h                    float64         `json:"volume_12h"`
	Volume8h                     float64         `json:"volume_8h"`
	Volume4h                     float64         `json:"volume_4h"`
	Volume2h                     float64         `json:"volume_2h"`
	Volume1h                     float64         `json:"volume_1h"`
	Volume30m                    float64         `json:"volume_30m"`
	Volume24hBase                float64         `json:"volume_24h_base"`
	Volume12hBase                float64         `json:"volume_12h_base"`
	Volume8hBase                 float64         `json:"volume_8h_base"`
	Volume4hBase                 float64         `json:"volume_4h_base"`
	Volume2hBase                 float64         `json:"volume_2h_base"`
	Volume1hBase                 float64         `json:"volume_1h_base"`
	Volume30mBase                float64         `json:"volume_30m_base"`
	Volume24hQuote               float64         `json:"volume_24h_quote"`
	Volume12hQuote               float64         `json:"volume_12h_quote"`
	Volume8hQuote                float64         `json:"volume_8h_quote"`
	Volume4hQuote                float64         `json:"volume_4h_quote"`
	Volume2hQuote                float64         `json:"volume_2h_quote"`
	Volume1hQuote                float64         `json:"volume_1h_quote"`
	Volume30mQuote               float64         `json:"volume_30m_quote"`
	Volume24hChangePercentage24h *float64        `json:"volume_24h_change_percentage_24h"`
}

// ============================================================================
// Token List Types
// ============================================================================

// RespTokenListV1Token represents token details in a token list v1 response
type RespTokenListV1Token struct {
	IsScaledUIToken   bool     `json:"isScaledUiToken"`
	Multiplier        *float64 `json:"multiplier"`
	Address           string   `json:"address"`
	Decimals          int64    `json:"decimals"`
	Price             float64  `json:"price"`
	LastTradeUnixTime int64    `json:"lastTradeUnixTime"`
	Liquidity         float64  `json:"liquidity"`
	LogoURI           string   `json:"logoURI"`
	MC                float64  `json:"mc"` // Market cap
	Name              string   `json:"name"`
	Symbol            string   `json:"symbol"`
	V24hChangePercent float64  `json:"v24hChangePercent"`
	V24hUSD           float64  `json:"v24hUSD"`
}

// RespTokenListV1 represents token list v1 response
type RespTokenListV1 struct {
	UpdateUnixTime int64                  `json:"updateUnixTime"`
	UpdateTime     string                 `json:"updateTime"`
	Tokens         []RespTokenListV1Token `json:"tokens"`
	Total          int64                  `json:"total"`
}

// TokenExtensions represents token extension metadata
type TokenExtensions struct {
	CoingeckoID *string `json:"coingecko_id,omitempty"`
	SerumV3USDC *string `json:"serum_v3_usdc,omitempty"`
	SerumV3USDT *string `json:"serum_v3_usdt,omitempty"`
	Website     *string `json:"website,omitempty"`
	Telegram    *string `json:"telegram,omitempty"`
	Twitter     *string `json:"twitter,omitempty"`
	Description *string `json:"description,omitempty"`
	Discord     *string `json:"discord,omitempty"`
	Medium      *string `json:"medium,omitempty"`
}

// RespTokenListV3TokenItem represents individual token details in V3
type RespTokenListV3TokenItem struct {
	Address                      string          `json:"address"`
	LogoURI                      string          `json:"logo_uri"`
	Name                         string          `json:"name"`
	Symbol                       string          `json:"symbol"`
	Decimals                     int64           `json:"decimals"`
	Extensions                   TokenExtensions `json:"extensions"`
	MarketCap                    float64         `json:"market_cap"`
	FDV                          float64         `json:"fdv"`
	TotalSupply                  float64         `json:"total_supply"`
	CirculatingSupply            float64         `json:"circulating_supply"`
	Liquidity                    float64         `json:"liquidity"`
	LastTradeUnixTime            int64           `json:"last_trade_unix_time"`
	Volume1hUSD                  float64         `json:"volume_1h_usd"`
	Volume1hChangePercent        float64         `json:"volume_1h_change_percent"`
	Volume2hUSD                  float64         `json:"volume_2h_usd"`
	Volume2hChangePercent        float64         `json:"volume_2h_change_percent"`
	Volume4hUSD                  float64         `json:"volume_4h_usd"`
	Volume4hChangePercent        float64         `json:"volume_4h_change_percent"`
	Volume8hUSD                  float64         `json:"volume_8h_usd"`
	Volume8hChangePercent        float64         `json:"volume_8h_change_percent"`
	Volume24hUSD                 float64         `json:"volume_24h_usd"`
	Volume24hChangePercent       float64         `json:"volume_24h_change_percent"`
	Trade1hCount                 int64           `json:"trade_1h_count"`
	Trade2hCount                 int64           `json:"trade_2h_count"`
	Trade4hCount                 int64           `json:"trade_4h_count"`
	Trade8hCount                 int64           `json:"trade_8h_count"`
	Trade24hCount                int64           `json:"trade_24h_count"`
	Buy24h                       int64           `json:"buy_24h"`
	Buy24hChangePercent          float64         `json:"buy_24h_change_percent"`
	VolumeBuy24hUSD              float64         `json:"volume_buy_24h_usd"`
	VolumeBuy24hChangePercent    float64         `json:"volume_buy_24h_change_percent"`
	Sell24h                      int64           `json:"sell_24h"`
	Sell24hChangePercent         float64         `json:"sell_24h_change_percent"`
	VolumeSell24hUSD             float64         `json:"volume_sell_24h_usd"`
	VolumeSell24hChangePercent   float64         `json:"volume_sell_24h_change_percent"`
	UniqueWallet24h              int64           `json:"unique_wallet_24h"`
	UniqueWallet24hChangePercent float64         `json:"unique_wallet_24h_change_percent"`
	Price                        float64         `json:"price"`
	PriceChange1hPercent         float64         `json:"price_change_1h_percent"`
	PriceChange2hPercent         float64         `json:"price_change_2h_percent"`
	PriceChange4hPercent         float64         `json:"price_change_4h_percent"`
	PriceChange8hPercent         float64         `json:"price_change_8h_percent"`
	PriceChange24hPercent        float64         `json:"price_change_24h_percent"`
	Holder                       int64           `json:"holder"`
	RecentListingTime            *int64          `json:"recent_listing_time"`
	IsScaledUIToken              bool            `json:"is_scaled_ui_token"`
	Multiplier                   *float64        `json:"multiplier"`
}

// RespTokenListV3 represents token list V3 response
type RespTokenListV3 struct {
	Items   []RespTokenListV3TokenItem `json:"items"`
	HasNext bool                       `json:"hasNext"`
}

// RespTokenListV3Scroll represents token list v3 scroll response
type RespTokenListV3Scroll struct {
	Items   []RespTokenListV3TokenItem `json:"items"`
	HasNext bool                       `json:"hasNext"`
}

// ============================================================================
// Token Overview Type
// ============================================================================

// RespTokenOverview represents token overview response including price, volume and trading statistics
type RespTokenOverview struct {
	Address            string          `json:"address"`
	Decimals           int64           `json:"decimals"`
	Symbol             string          `json:"symbol"`
	Name               string          `json:"name"`
	MarketCap          float64         `json:"marketCap"`
	FDV                float64         `json:"fdv"`
	Extensions         TokenExtensions `json:"extensions"`
	LogoURI            string          `json:"logoURI"`
	Liquidity          float64         `json:"liquidity"`
	LastTradeUnixTime  int64           `json:"lastTradeUnixTime"`
	LastTradeHumanTime string          `json:"lastTradeHumanTime"`
	Price              float64         `json:"price"`
	TotalSupply        float64         `json:"totalSupply"`
	CirculatingSupply  float64         `json:"circulatingSupply"`
	Holder             int64           `json:"holder"`
	NumberMarkets      int64           `json:"numberMarkets"`
	IsScaledUIToken    bool            `json:"isScaledUiToken"`
	Multiplier         *float64        `json:"multiplier"`
	// Note: Contains many more price history and stats fields - truncated for brevity
	// The full Python version has extensive price/volume/trade statistics fields
}

// ============================================================================
// Token Metadata, Market Data, and Trade Data Types
// ============================================================================

// RespTokenMetadata represents token metadata response
type RespTokenMetadata struct {
	Address    string          `json:"address"`
	Symbol     string          `json:"symbol"`
	Name       string          `json:"name"`
	Decimals   int64           `json:"decimals"`
	Extensions TokenExtensions `json:"extensions"`
	LogoURI    string          `json:"logo_uri"`
}

type RespMultiTokenMetadata = map[string]RespTokenMetadata

// RespTokenMarketData represents token market data response
type RespTokenMarketData struct {
	Address           string   `json:"address"`
	Price             float64  `json:"price"`
	Liquidity         float64  `json:"liquidity"`
	TotalSupply       float64  `json:"total_supply"`
	CirculatingSupply float64  `json:"circulating_supply"`
	FDV               float64  `json:"fdv"` // Fully diluted valuation
	MarketCap         float64  `json:"market_cap"`
	IsScaledUIToken   bool     `json:"is_scaled_ui_token"`
	Multiplier        *float64 `json:"multiplier"`
}

// RespTokenTradeData represents token trade data response (this is a large struct, split for readability)
type RespTokenTradeData struct {
	Address                      string  `json:"address"`
	Holder                       int64   `json:"holder"`
	Market                       int64   `json:"market"`
	LastTradeUnixTime            int64   `json:"last_trade_unix_time"`
	LastTradeHumanTime           string  `json:"last_trade_human_time"`
	Price                        float64 `json:"price"`
	History1mPrice               float64 `json:"history_1m_price"`
	PriceChange1mPercent         float64 `json:"price_change_1m_percent"`
	History5mPrice               float64 `json:"history_5m_price"`
	PriceChange5mPercent         float64 `json:"price_change_5m_percent"`
	History30mPrice              float64 `json:"history_30m_price"`
	PriceChange30mPercent        float64 `json:"price_change_30m_percent"`
	History1hPrice               float64 `json:"history_1h_price"`
	PriceChange1hPercent         float64 `json:"price_change_1h_percent"`
	History2hPrice               float64 `json:"history_2h_price"`
	PriceChange2hPercent         float64 `json:"price_change_2h_percent"`
	History4hPrice               float64 `json:"history_4h_price"`
	PriceChange4hPercent         float64 `json:"price_change_4h_percent"`
	History6hPrice               float64 `json:"history_6h_price"`
	PriceChange6hPercent         float64 `json:"price_change_6h_percent"`
	History8hPrice               float64 `json:"history_8h_price"`
	PriceChange8hPercent         float64 `json:"price_change_8h_percent"`
	History12hPrice              float64 `json:"history_12h_price"`
	PriceChange12hPercent        float64 `json:"price_change_12h_percent"`
	History24hPrice              float64 `json:"history_24h_price"`
	PriceChange24hPercent        float64 `json:"price_change_24h_percent"`
	UniqueWallet1m               int64   `json:"unique_wallet_1m"`
	UniqueWalletHistory1m        int64   `json:"unique_wallet_history_1m"`
	UniqueWallet1mChangePercent  float64 `json:"unique_wallet_1m_change_percent"`
	UniqueWallet5m               int64   `json:"unique_wallet_5m"`
	UniqueWalletHistory5m        int64   `json:"unique_wallet_history_5m"`
	UniqueWallet5mChangePercent  float64 `json:"unique_wallet_5m_change_percent"`
	UniqueWallet30m              int64   `json:"unique_wallet_30m"`
	UniqueWalletHistory30m       int64   `json:"unique_wallet_history_30m"`
	UniqueWallet30mChangePercent float64 `json:"unique_wallet_30m_change_percent"`
	UniqueWallet1h               int64   `json:"unique_wallet_1h"`
	UniqueWalletHistory1h        int64   `json:"unique_wallet_history_1h"`
	UniqueWallet1hChangePercent  float64 `json:"unique_wallet_1h_change_percent"`
	UniqueWallet2h               int64   `json:"unique_wallet_2h"`
	UniqueWalletHistory2h        int64   `json:"unique_wallet_history_2h"`
	UniqueWallet2hChangePercent  float64 `json:"unique_wallet_2h_change_percent"`
	UniqueWallet4h               int64   `json:"unique_wallet_4h"`
	UniqueWalletHistory4h        int64   `json:"unique_wallet_history_4h"`
	UniqueWallet4hChangePercent  float64 `json:"unique_wallet_4h_change_percent"`
	UniqueWallet8h               int64   `json:"unique_wallet_8h"`
	UniqueWalletHistory8h        int64   `json:"unique_wallet_history_8h"`
	UniqueWallet8hChangePercent  float64 `json:"unique_wallet_8h_change_percent"`
	UniqueWallet24h              int64   `json:"unique_wallet_24h"`
	UniqueWalletHistory24h       int64   `json:"unique_wallet_history_24h"`
	UniqueWallet24hChangePercent float64 `json:"unique_wallet_24h_change_percent"`
	Trade1m                      int64   `json:"trade_1m"`
	TradeHistory1m               int64   `json:"trade_history_1m"`
	Trade1mChangePercent         float64 `json:"trade_1m_change_percent"`
	Sell1m                       int64   `json:"sell_1m"`
	SellHistory1m                int64   `json:"sell_history_1m"`
	Sell1mChangePercent          float64 `json:"sell_1m_change_percent"`
	Buy1m                        int64   `json:"buy_1m"`
	BuyHistory1m                 int64   `json:"buy_history_1m"`
	Buy1mChangePercent           float64 `json:"buy_1m_change_percent"`
	Volume1m                     float64 `json:"volume_1m"`
	Volume1mUSD                  float64 `json:"volume_1m_usd"`
	VolumeHistory1m              float64 `json:"volume_history_1m"`
	VolumeHistory1mUSD           float64 `json:"volume_history_1m_usd"`
	Volume1mChangePercent        float64 `json:"volume_1m_change_percent"`
	VolumeBuy1m                  float64 `json:"volume_buy_1m"`
	VolumeBuy1mUSD               float64 `json:"volume_buy_1m_usd"`
	VolumeBuyHistory1m           float64 `json:"volume_buy_history_1m"`
	VolumeBuyHistory1mUSD        float64 `json:"volume_buy_history_1m_usd"`
	VolumeBuy1mChangePercent     float64 `json:"volume_buy_1m_change_percent"`
	VolumeSell1m                 float64 `json:"volume_sell_1m"`
	VolumeSell1mUSD              float64 `json:"volume_sell_1m_usd"`
	VolumeSellHistory1m          float64 `json:"volume_sell_history_1m"`
	VolumeSellHistory1mUSD       float64 `json:"volume_sell_history_1m_usd"`
	VolumeSell1mChangePercent    float64 `json:"volume_sell_1m_change_percent"`

	// 5m time period statistics
	Trade5m                   int64   `json:"trade_5m"`
	TradeHistory5m            int64   `json:"trade_history_5m"`
	Trade5mChangePercent      float64 `json:"trade_5m_change_percent"`
	Sell5m                    int64   `json:"sell_5m"`
	SellHistory5m             int64   `json:"sell_history_5m"`
	Sell5mChangePercent       float64 `json:"sell_5m_change_percent"`
	Buy5m                     int64   `json:"buy_5m"`
	BuyHistory5m              int64   `json:"buy_history_5m"`
	Buy5mChangePercent        float64 `json:"buy_5m_change_percent"`
	Volume5m                  float64 `json:"volume_5m"`
	Volume5mUSD               float64 `json:"volume_5m_usd"`
	VolumeHistory5m           float64 `json:"volume_history_5m"`
	VolumeHistory5mUSD        float64 `json:"volume_history_5m_usd"`
	Volume5mChangePercent     float64 `json:"volume_5m_change_percent"`
	VolumeBuy5m               float64 `json:"volume_buy_5m"`
	VolumeBuy5mUSD            float64 `json:"volume_buy_5m_usd"`
	VolumeBuyHistory5m        float64 `json:"volume_buy_history_5m"`
	VolumeBuyHistory5mUSD     float64 `json:"volume_buy_history_5m_usd"`
	VolumeBuy5mChangePercent  float64 `json:"volume_buy_5m_change_percent"`
	VolumeSell5m              float64 `json:"volume_sell_5m"`
	VolumeSell5mUSD           float64 `json:"volume_sell_5m_usd"`
	VolumeSellHistory5m       float64 `json:"volume_sell_history_5m"`
	VolumeSellHistory5mUSD    float64 `json:"volume_sell_history_5m_usd"`
	VolumeSell5mChangePercent float64 `json:"volume_sell_5m_change_percent"`

	// 30m time period statistics
	Trade30m                   int64   `json:"trade_30m"`
	TradeHistory30m            int64   `json:"trade_history_30m"`
	Trade30mChangePercent      float64 `json:"trade_30m_change_percent"`
	Sell30m                    int64   `json:"sell_30m"`
	SellHistory30m             int64   `json:"sell_history_30m"`
	Sell30mChangePercent       float64 `json:"sell_30m_change_percent"`
	Buy30m                     int64   `json:"buy_30m"`
	BuyHistory30m              int64   `json:"buy_history_30m"`
	Buy30mChangePercent        float64 `json:"buy_30m_change_percent"`
	Volume30m                  float64 `json:"volume_30m"`
	Volume30mUSD               float64 `json:"volume_30m_usd"`
	VolumeHistory30m           float64 `json:"volume_history_30m"`
	VolumeHistory30mUSD        float64 `json:"volume_history_30m_usd"`
	Volume30mChangePercent     float64 `json:"volume_30m_change_percent"`
	VolumeBuy30m               float64 `json:"volume_buy_30m"`
	VolumeBuy30mUSD            float64 `json:"volume_buy_30m_usd"`
	VolumeBuyHistory30m        float64 `json:"volume_buy_history_30m"`
	VolumeBuyHistory30mUSD     float64 `json:"volume_buy_history_30m_usd"`
	VolumeBuy30mChangePercent  float64 `json:"volume_buy_30m_change_percent"`
	VolumeSell30m              float64 `json:"volume_sell_30m"`
	VolumeSell30mUSD           float64 `json:"volume_sell_30m_usd"`
	VolumeSellHistory30m       float64 `json:"volume_sell_history_30m"`
	VolumeSellHistory30mUSD    float64 `json:"volume_sell_history_30m_usd"`
	VolumeSell30mChangePercent float64 `json:"volume_sell_30m_change_percent"`

	// 1h time period statistics
	Trade1h                   int64   `json:"trade_1h"`
	TradeHistory1h            int64   `json:"trade_history_1h"`
	Trade1hChangePercent      float64 `json:"trade_1h_change_percent"`
	Sell1h                    int64   `json:"sell_1h"`
	SellHistory1h             int64   `json:"sell_history_1h"`
	Sell1hChangePercent       float64 `json:"sell_1h_change_percent"`
	Buy1h                     int64   `json:"buy_1h"`
	BuyHistory1h              int64   `json:"buy_history_1h"`
	Buy1hChangePercent        float64 `json:"buy_1h_change_percent"`
	Volume1h                  float64 `json:"volume_1h"`
	Volume1hUSD               float64 `json:"volume_1h_usd"`
	VolumeHistory1h           float64 `json:"volume_history_1h"`
	VolumeHistory1hUSD        float64 `json:"volume_history_1h_usd"`
	Volume1hChangePercent     float64 `json:"volume_1h_change_percent"`
	VolumeBuy1h               float64 `json:"volume_buy_1h"`
	VolumeBuy1hUSD            float64 `json:"volume_buy_1h_usd"`
	VolumeBuyHistory1h        float64 `json:"volume_buy_history_1h"`
	VolumeBuyHistory1hUSD     float64 `json:"volume_buy_history_1h_usd"`
	VolumeBuy1hChangePercent  float64 `json:"volume_buy_1h_change_percent"`
	VolumeSell1h              float64 `json:"volume_sell_1h"`
	VolumeSell1hUSD           float64 `json:"volume_sell_1h_usd"`
	VolumeSellHistory1h       float64 `json:"volume_sell_history_1h"`
	VolumeSellHistory1hUSD    float64 `json:"volume_sell_history_1h_usd"`
	VolumeSell1hChangePercent float64 `json:"volume_sell_1h_change_percent"`

	// 2h time period statistics
	Trade2h                   int64   `json:"trade_2h"`
	TradeHistory2h            int64   `json:"trade_history_2h"`
	Trade2hChangePercent      float64 `json:"trade_2h_change_percent"`
	Sell2h                    int64   `json:"sell_2h"`
	SellHistory2h             int64   `json:"sell_history_2h"`
	Sell2hChangePercent       float64 `json:"sell_2h_change_percent"`
	Buy2h                     int64   `json:"buy_2h"`
	BuyHistory2h              int64   `json:"buy_history_2h"`
	Buy2hChangePercent        float64 `json:"buy_2h_change_percent"`
	Volume2h                  float64 `json:"volume_2h"`
	Volume2hUSD               float64 `json:"volume_2h_usd"`
	VolumeHistory2h           float64 `json:"volume_history_2h"`
	VolumeHistory2hUSD        float64 `json:"volume_history_2h_usd"`
	Volume2hChangePercent     float64 `json:"volume_2h_change_percent"`
	VolumeBuy2h               float64 `json:"volume_buy_2h"`
	VolumeBuy2hUSD            float64 `json:"volume_buy_2h_usd"`
	VolumeBuyHistory2h        float64 `json:"volume_buy_history_2h"`
	VolumeBuyHistory2hUSD     float64 `json:"volume_buy_history_2h_usd"`
	VolumeBuy2hChangePercent  float64 `json:"volume_buy_2h_change_percent"`
	VolumeSell2h              float64 `json:"volume_sell_2h"`
	VolumeSell2hUSD           float64 `json:"volume_sell_2h_usd"`
	VolumeSellHistory2h       float64 `json:"volume_sell_history_2h"`
	VolumeSellHistory2hUSD    float64 `json:"volume_sell_history_2h_usd"`
	VolumeSell2hChangePercent float64 `json:"volume_sell_2h_change_percent"`

	// 4h time period statistics
	Trade4h                   int64   `json:"trade_4h"`
	TradeHistory4h            int64   `json:"trade_history_4h"`
	Trade4hChangePercent      float64 `json:"trade_4h_change_percent"`
	Sell4h                    int64   `json:"sell_4h"`
	SellHistory4h             int64   `json:"sell_history_4h"`
	Sell4hChangePercent       float64 `json:"sell_4h_change_percent"`
	Buy4h                     int64   `json:"buy_4h"`
	BuyHistory4h              int64   `json:"buy_history_4h"`
	Buy4hChangePercent        float64 `json:"buy_4h_change_percent"`
	Volume4h                  float64 `json:"volume_4h"`
	Volume4hUSD               float64 `json:"volume_4h_usd"`
	VolumeHistory4h           float64 `json:"volume_history_4h"`
	VolumeHistory4hUSD        float64 `json:"volume_history_4h_usd"`
	Volume4hChangePercent     float64 `json:"volume_4h_change_percent"`
	VolumeBuy4h               float64 `json:"volume_buy_4h"`
	VolumeBuy4hUSD            float64 `json:"volume_buy_4h_usd"`
	VolumeBuyHistory4h        float64 `json:"volume_buy_history_4h"`
	VolumeBuyHistory4hUSD     float64 `json:"volume_buy_history_4h_usd"`
	VolumeBuy4hChangePercent  float64 `json:"volume_buy_4h_change_percent"`
	VolumeSell4h              float64 `json:"volume_sell_4h"`
	VolumeSell4hUSD           float64 `json:"volume_sell_4h_usd"`
	VolumeSellHistory4h       float64 `json:"volume_sell_history_4h"`
	VolumeSellHistory4hUSD    float64 `json:"volume_sell_history_4h_usd"`
	VolumeSell4hChangePercent float64 `json:"volume_sell_4h_change_percent"`

	// 8h time period statistics
	Trade8h                   int64   `json:"trade_8h"`
	TradeHistory8h            int64   `json:"trade_history_8h"`
	Trade8hChangePercent      float64 `json:"trade_8h_change_percent"`
	Sell8h                    int64   `json:"sell_8h"`
	SellHistory8h             int64   `json:"sell_history_8h"`
	Sell8hChangePercent       float64 `json:"sell_8h_change_percent"`
	Buy8h                     int64   `json:"buy_8h"`
	BuyHistory8h              int64   `json:"buy_history_8h"`
	Buy8hChangePercent        float64 `json:"buy_8h_change_percent"`
	Volume8h                  float64 `json:"volume_8h"`
	Volume8hUSD               float64 `json:"volume_8h_usd"`
	VolumeHistory8h           float64 `json:"volume_history_8h"`
	VolumeHistory8hUSD        float64 `json:"volume_history_8h_usd"`
	Volume8hChangePercent     float64 `json:"volume_8h_change_percent"`
	VolumeBuy8h               float64 `json:"volume_buy_8h"`
	VolumeBuy8hUSD            float64 `json:"volume_buy_8h_usd"`
	VolumeBuyHistory8h        float64 `json:"volume_buy_history_8h"`
	VolumeBuyHistory8hUSD     float64 `json:"volume_buy_history_8h_usd"`
	VolumeBuy8hChangePercent  float64 `json:"volume_buy_8h_change_percent"`
	VolumeSell8h              float64 `json:"volume_sell_8h"`
	VolumeSell8hUSD           float64 `json:"volume_sell_8h_usd"`
	VolumeSellHistory8h       float64 `json:"volume_sell_history_8h"`
	VolumeSellHistory8hUSD    float64 `json:"volume_sell_history_8h_usd"`
	VolumeSell8hChangePercent float64 `json:"volume_sell_8h_change_percent"`

	// 24h time period statistics
	Trade24h                   int64   `json:"trade_24h"`
	TradeHistory24h            int64   `json:"trade_history_24h"`
	Trade24hChangePercent      float64 `json:"trade_24h_change_percent"`
	Sell24h                    int64   `json:"sell_24h"`
	SellHistory24h             int64   `json:"sell_history_24h"`
	Sell24hChangePercent       float64 `json:"sell_24h_change_percent"`
	Buy24h                     int64   `json:"buy_24h"`
	BuyHistory24h              int64   `json:"buy_history_24h"`
	Buy24hChangePercent        float64 `json:"buy_24h_change_percent"`
	Volume24h                  float64 `json:"volume_24h"`
	Volume24hUSD               float64 `json:"volume_24h_usd"`
	VolumeHistory24h           float64 `json:"volume_history_24h"`
	VolumeHistory24hUSD        float64 `json:"volume_history_24h_usd"`
	Volume24hChangePercent     float64 `json:"volume_24h_change_percent"`
	VolumeBuy24h               float64 `json:"volume_buy_24h"`
	VolumeBuy24hUSD            float64 `json:"volume_buy_24h_usd"`
	VolumeBuyHistory24h        float64 `json:"volume_buy_history_24h"`
	VolumeBuyHistory24hUSD     float64 `json:"volume_buy_history_24h_usd"`
	VolumeBuy24hChangePercent  float64 `json:"volume_buy_24h_change_percent"`
	VolumeSell24h              float64 `json:"volume_sell_24h"`
	VolumeSell24hUSD           float64 `json:"volume_sell_24h_usd"`
	VolumeSellHistory24h       float64 `json:"volume_sell_history_24h"`
	VolumeSellHistory24hUSD    float64 `json:"volume_sell_history_24h_usd"`
	VolumeSell24hChangePercent float64 `json:"volume_sell_24h_change_percent"`

	IsScaledUIToken bool     `json:"is_scaled_ui_token"`
	Multiplier      *float64 `json:"multiplier"`
}

// ============================================================================
// Token Security and Additional Types
// ============================================================================

// RespTokenSecurity represents token security information response
type RespTokenSecurity struct {
	CreatorAddress                 *string  `json:"creatorAddress"`
	CreatorOwnerAddress            *string  `json:"creatorOwnerAddress"`
	OwnerAddress                   *string  `json:"ownerAddress"`
	OwnerOfOwnerAddress            *string  `json:"ownerOfOwnerAddress"`
	CreationTx                     *string  `json:"creationTx"`
	CreationTime                   *int64   `json:"creationTime"`
	CreationSlot                   *int64   `json:"creationSlot"`
	MintTx                         *string  `json:"mintTx"`
	MintTime                       *int64   `json:"mintTime"`
	MintSlot                       *int64   `json:"mintSlot"`
	CreatorBalance                 *float64 `json:"creatorBalance"`
	OwnerBalance                   *float64 `json:"ownerBalance"`
	OwnerPercentage                *float64 `json:"ownerPercentage"`
	CreatorPercentage              *float64 `json:"creatorPercentage"`
	MetaplexUpdateAuthority        *string  `json:"metaplexUpdateAuthority"`
	MetaplexOwnerUpdateAuthority   *string  `json:"metaplexOwnerUpdateAuthority"`
	MetaplexUpdateAuthorityBalance *float64 `json:"metaplexUpdateAuthorityBalance"`
	MetaplexUpdateAuthorityPercent *float64 `json:"metaplexUpdateAuthorityPercent"`
	MutableMetadata                *bool    `json:"mutableMetadata"`
	Top10HolderBalance             *float64 `json:"top10HolderBalance"`
	Top10HolderPercent             *float64 `json:"top10HolderPercent"`
	Top10UserBalance               *float64 `json:"top10UserBalance"`
	Top10UserPercent               *float64 `json:"top10UserPercent"`
	IsTrueToken                    *bool    `json:"isTrueToken"`
	FakeToken                      *bool    `json:"fakeToken"`
	TotalSupply                    *float64 `json:"totalSupply"`
	PreMarketHolder                []any    `json:"preMarketHolder"`
	LockInfo                       any      `json:"lockInfo"`
	Freezeable                     *bool    `json:"freezeable"`
	FreezeAuthority                *string  `json:"freezeAuthority"`
	TransferFeeEnable              *bool    `json:"transferFeeEnable"`
	TransferFeeData                any      `json:"transferFeeData"`
	IsToken2022                    bool     `json:"isToken2022"`
	NonTransferable                *bool    `json:"nonTransferable"`
	JupStrictList                  bool     `json:"jupStrictList"`
}

// RespTokenCreationInfo represents token creation information response
type RespTokenCreationInfo struct {
	TxHash         string `json:"txHash"`
	Slot           int64  `json:"slot"`
	TokenAddress   string `json:"tokenAddress"`
	Decimals       int64  `json:"decimals"`
	Owner          string `json:"owner"`
	BlockUnixTime  int64  `json:"blockUnixTime"`
	BlockHumanTime string `json:"blockHumanTime"`
}

// RespTokenMintBurnTxItem represents token mint/burn transaction details
type RespTokenMintBurnTxItem struct {
	Amount         string       `json:"amount"`
	BlockHumanTime string       `json:"block_human_time"`
	BlockTime      int64        `json:"block_time"`
	CommonType     MintBurnType `json:"common_type"`
	Decimals       int64        `json:"decimals"`
	Mint           string       `json:"mint"`
	ProgramID      string       `json:"program_id"`
	Slot           int64        `json:"slot"`
	TxHash         string       `json:"tx_hash"`
	UIAmount       float64      `json:"ui_amount"`
	UIAmountString string       `json:"ui_amount_string"`
}

// RespTokenAllTimeTrades represents token all-time trade statistics response
type RespTokenAllTimeTrades struct {
	Address        string  `json:"address"`
	TotalVolume    float64 `json:"total_volume"`
	TotalVolumeUSD float64 `json:"total_volume_usd"`
	VolumeBuyUSD   float64 `json:"volume_buy_usd"`
	VolumeSellUSD  float64 `json:"volume_sell_usd"`
	VolumeBuy      float64 `json:"volume_buy"`
	VolumeSell     float64 `json:"volume_sell"`
	TotalTrade     int64   `json:"total_trade"`
	Buy            int64   `json:"buy"`
	Sell           int64   `json:"sell"`
}

type RespMultiTokenAllTimeTrades = []RespTokenAllTimeTrades

// RespTokenHoldersItem represents token holder information
type RespTokenHoldersItem struct {
	Amount          string   `json:"amount"`
	Decimals        int64    `json:"decimals"`
	Mint            string   `json:"mint"`
	Owner           string   `json:"owner"`
	TokenAccount    string   `json:"token_account"`
	UIAmount        float64  `json:"ui_amount"`
	IsScaledUIToken bool     `json:"is_scaled_ui_token"`
	Multiplier      *float64 `json:"multiplier"`
}

type RespMultiTokenHolders = []RespTokenHoldersItem

// RespNewTokenListingItem represents newly listed token information
type RespNewTokenListingItem struct {
	Address          string  `json:"address"`
	Symbol           string  `json:"symbol"`
	Name             string  `json:"name"`
	Decimals         int64   `json:"decimals"`
	Source           string  `json:"source"`
	LiquidityAddedAt string  `json:"liquidityAddedAt"`
	LogoURI          *string `json:"logoURI"`
	Liquidity        float64 `json:"liquidity"`
}

// RespGainerLoserItem represents gainer/loser item
type RespGainerLoserItem struct {
	Network    string  `json:"network"`
	Address    string  `json:"address"`
	PNL        float64 `json:"pnl"`
	TradeCount int64   `json:"trade_count"`
	Volume     float64 `json:"volume"`
}

// ============================================================================
// Wallet Related Types
// ============================================================================

// RespWalletPortfolioItem represents an individual token in wallet portfolio
type RespWalletPortfolioItem struct {
	Address         string   `json:"address"`
	Decimals        int64    `json:"decimals"`
	Balance         int64    `json:"balance"`
	UIAmount        float64  `json:"uiAmount"`
	ChainID         string   `json:"chainId"`
	Name            string   `json:"name"`
	Symbol          string   `json:"symbol"`
	Icon            string   `json:"icon"`
	LogoURI         string   `json:"logoURI"`
	PriceUSD        float64  `json:"priceUsd"`
	ValueUSD        float64  `json:"valueUsd"`
	IsScaledUIToken bool     `json:"isScaledUiToken"`
	Multiplier      *float64 `json:"multiplier"`
}

// RespWalletTokenBalance represents wallet token balance
type RespWalletTokenBalance struct {
	Address         string   `json:"address"`
	Decimals        int64    `json:"decimals"`
	Balance         int64    `json:"balance"`
	UIAmount        float64  `json:"uiAmount"`
	ChainID         string   `json:"chainId"`
	LogoURI         string   `json:"logoURI"`
	Name            string   `json:"name"`
	Symbol          string   `json:"symbol"`
	PriceUSD        float64  `json:"priceUsd"`
	ValueUSD        float64  `json:"valueUsd"`
	IsScaledUIToken bool     `json:"isScaledUiToken"`
	Multiplier      *float64 `json:"multiplier"`
}

// RespWalletTokenFirstTx represents first token transaction in a wallet
type RespWalletTokenFirstTx struct {
	TxHash        string `json:"tx_hash"`
	BlockUnixTime int64  `json:"block_unix_time"`
	BlockNumber   int64  `json:"block_number"`
	BalanceChange string `json:"balance_change"`
	TokenAddress  string `json:"token_address"`
	TokenDecimals int64  `json:"token_decimals"`
}

// RespWalletTokensBalanceItem represents wallet token balance item
type RespWalletTokensBalanceItem struct {
	Address  string  `json:"address"`
	Decimals int64   `json:"decimals"`
	Price    float64 `json:"price"`
	Balance  string  `json:"balance"`
	Amount   float64 `json:"amount"`
	Network  string  `json:"network"`
	Name     string  `json:"name"`
	Symbol   string  `json:"symbol"`
	LogoURI  string  `json:"logo_uri"`
	Value    string  `json:"value"`
}

// RespWalletBalanceChangesTokenInfo represents token info in balance change response
type RespWalletBalanceChangesTokenInfo struct {
	Address         string   `json:"address"`
	Decimals        int64    `json:"decimals"`
	Symbol          string   `json:"symbol"`
	Name            string   `json:"name"`
	LogoURI         string   `json:"logo_uri"`
	IsScaledUIToken bool     `json:"is_scaled_ui_token"`
	Multiplier      *float64 `json:"multiplier"`
}

// RespWalletBalanceChangesItem represents wallet balance change item
type RespWalletBalanceChangesItem struct {
	Time           string                            `json:"time"`
	BlockNumber    int64                             `json:"block_number"`
	BlockUnixTime  int64                             `json:"block_unix_time"`
	Address        string                            `json:"address"`
	TokenAccount   string                            `json:"token_account"`
	TxHash         string                            `json:"tx_hash"`
	PreBalance     string                            `json:"pre_balance"`
	PostBalance    string                            `json:"post_balance"`
	Amount         string                            `json:"amount"`
	TokenInfo      RespWalletBalanceChangesTokenInfo `json:"token_info"`
	Type           int64                             `json:"type"`
	TypeText       BalanceChangeType                 `json:"type_text"`
	ChangeType     int64                             `json:"change_type"`
	ChangeTypeText BalanceChangeDirection            `json:"change_type_text"`
}

// RespWalletTradesToken represents token in a wallet trade
type RespWalletTradesToken struct {
	Symbol          string   `json:"symbol"`
	Decimals        int64    `json:"decimals"`
	Address         string   `json:"address"`
	Amount          int64    `json:"amount"`
	Type            string   `json:"type"`
	TypeSwap        string   `json:"type_swap"`
	UIAmount        float64  `json:"ui_amount"`
	Price           float64  `json:"price"`
	NearestPrice    float64  `json:"nearest_price"`
	ChangeAmount    int64    `json:"change_amount"`
	UIChangeAmount  float64  `json:"ui_change_amount"`
	IsScaledUIToken bool     `json:"is_scaled_ui_token"`
	Multiplier      *float64 `json:"multiplier"`
}

// RespWalletTradesItem represents a single wallet trade
type RespWalletTradesItem struct {
	Quote               RespWalletTradesToken `json:"quote"`
	Base                RespWalletTradesToken `json:"base"`
	BasePrice           float64               `json:"base_price"`
	QuotePrice          float64               `json:"quote_price"`
	TxHash              string                `json:"tx_hash"`
	Source              string                `json:"source"`
	BlockUnixTime       int64                 `json:"block_unix_time"`
	TxType              string                `json:"tx_type"`
	Address             string                `json:"address"`
	Owner               string                `json:"owner"`
	BlockNumber         int64                 `json:"block_number"`
	VolumeUSD           float64               `json:"volume_usd"`
	Volume              float64               `json:"volume"`
	InsIndex            int64                 `json:"ins_index"`
	InnerInsIndex       int64                 `json:"inner_ins_index"`
	Signers             []string              `json:"signers"`
	InteractedProgramID string                `json:"interacted_program_id"`
}

// RespWalletTrades represents wallet trades response
type RespWalletTrades struct {
	Items   []RespWalletTradesItem `json:"items"`
	HasNext bool                   `json:"hasNext"`
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
	TxHash    string `json:"tx_hash"`
	Slot      int64  `json:"slot"`
	BlockTime int64  `json:"block_time"`
}

// MemePool represents pool configuration for meme tokens
type MemePool struct {
	Address               string  `json:"address"`
	CurveAmount           *string `json:"curve_amount,omitempty"`
	TotalSupply           *string `json:"total_supply,omitempty"`
	MarketcapThreshold    *string `json:"marketcap_threshold,omitempty"`
	CoefB                 *string `json:"coef_b,omitempty"`
	Bump                  *string `json:"bump,omitempty"`
	VirtualBase           *string `json:"virtual_base,omitempty"`
	Creator               *string `json:"creator,omitempty"`
	BaseDecimals          *int64  `json:"base_decimals,omitempty"`
	QuoteMint             *string `json:"quote_mint,omitempty"`
	AuthBump              *int64  `json:"auth_bump,omitempty"`
	TotalQuoteFundRaising *string `json:"total_quote_fund_raising,omitempty"`
	Supply                *string `json:"supply,omitempty"`
	PlatformFee           *string `json:"platform_fee,omitempty"`
	QuoteProtocolFee      *string `json:"quote_protocol_fee,omitempty"`
	TotalBaseSell         *string `json:"total_base_sell,omitempty"`
	VirtualQuote          *string `json:"virtual_quote,omitempty"`
	BaseMint              *string `json:"base_mint,omitempty"`
	BaseVault             *string `json:"base_vault,omitempty"`
	PlatformConfig        *string `json:"platform_config,omitempty"`
	QuoteDecimals         *int64  `json:"quote_decimals,omitempty"`
	RealQuote             *string `json:"real_quote,omitempty"`
	QuoteVault            *string `json:"quote_vault,omitempty"`
	RealBase              *string `json:"real_base,omitempty"`
	Status                *int64  `json:"status,omitempty"`
	RealSolReserves       *string `json:"real_sol_reserves,omitempty"`
	RealTokenReserves     *string `json:"real_token_reserves,omitempty"`
	TokenTotalSupply      *string `json:"token_total_supply,omitempty"`
	VirtualTokenReserves  *string `json:"virtual_token_reserves,omitempty"`
}

// MemeExtensions represents token extension information for meme tokens
type MemeExtensions struct {
	Twitter     string `json:"twitter"`
	Website     string `json:"website"`
	Description string `json:"description"`
}

// MemeInfo represents meme token specific information
type MemeInfo struct {
	Source          string         `json:"source"`
	PlatformID      string         `json:"platform_id"`
	Address         string         `json:"address"`
	CreatedAt       MemeCreatedAt  `json:"created_at"`
	CreationTime    int64          `json:"creation_time"`
	Creator         string         `json:"creator"`
	UpdatedAt       *MemeCreatedAt `json:"updated_at,omitempty"`
	GraduatedAt     *MemeCreatedAt `json:"graduated_at,omitempty"`
	Graduated       bool           `json:"graduated"`
	GraduatedTime   *int64         `json:"graduated_time"`
	Pool            MemePool       `json:"pool"`
	ProgressPercent float64        `json:"progress_percent"`
}

// RespMemeListItem represents individual meme token details
type RespMemeListItem struct {
	Address                string          `json:"address"`
	LogoURI                string          `json:"logo_uri"`
	Name                   string          `json:"name"`
	Symbol                 string          `json:"symbol"`
	Decimals               int64           `json:"decimals"`
	Extensions             *MemeExtensions `json:"extensions"`
	MarketCap              float64         `json:"market_cap"`
	FDV                    float64         `json:"fdv"`
	Liquidity              float64         `json:"liquidity"`
	LastTradeUnixTime      int64           `json:"last_trade_unix_time"`
	Volume1hUSD            float64         `json:"volume_1h_usd"`
	Volume1hChangePercent  float64         `json:"volume_1h_change_percent"`
	Volume2hUSD            float64         `json:"volume_2h_usd"`
	Volume2hChangePercent  float64         `json:"volume_2h_change_percent"`
	Volume4hUSD            float64         `json:"volume_4h_usd"`
	Volume4hChangePercent  float64         `json:"volume_4h_change_percent"`
	Volume8hUSD            float64         `json:"volume_8h_usd"`
	Volume8hChangePercent  float64         `json:"volume_8h_change_percent"`
	Volume24hUSD           float64         `json:"volume_24h_usd"`
	Volume24hChangePercent *float64        `json:"volume_24h_change_percent"`
	Trade1hCount           int64           `json:"trade_1h_count"`
	Trade2hCount           int64           `json:"trade_2h_count"`
	Trade4hCount           int64           `json:"trade_4h_count"`
	Trade8hCount           int64           `json:"trade_8h_count"`
	Trade24hCount          int64           `json:"trade_24h_count"`
	Price                  float64         `json:"price"`
	PriceChange1hPercent   float64         `json:"price_change_1h_percent"`
	PriceChange2hPercent   float64         `json:"price_change_2h_percent"`
	PriceChange4hPercent   float64         `json:"price_change_4h_percent"`
	PriceChange8hPercent   float64         `json:"price_change_8h_percent"`
	PriceChange24hPercent  float64         `json:"price_change_24h_percent"`
	Holder                 int64           `json:"holder"`
	RecentListingTime      int64           `json:"recent_listing_time"`
	MemeInfo               MemeInfo        `json:"meme_info"`
}

// RespMemeList represents response type for meme token list
type RespMemeList struct {
	Items   []RespMemeListItem `json:"items"`
	HasNext bool               `json:"has_next"`
}

// RespMemeDetail represents response type for meme token detail
type RespMemeDetail struct {
	Address           string         `json:"address"`
	Name              string         `json:"name"`
	Symbol            string         `json:"symbol"`
	Decimals          int64          `json:"decimals"`
	Extensions        MemeExtensions `json:"extensions"`
	LogoURI           string         `json:"logo_uri"`
	Price             float64        `json:"price"`
	Liquidity         float64        `json:"liquidity"`
	CirculatingSupply int64          `json:"circulating_supply"`
	MarketCap         int64          `json:"market_cap"`
	TotalSupply       int64          `json:"total_supply"`
	FDV               int64          `json:"fdv"`
	MemeInfo          MemeInfo       `json:"meme_info"`
}

// ============================================================================
// Search Types
// ============================================================================

// RespSearchTokenResult represents token search result information
type RespSearchTokenResult struct {
	Name                         string   `json:"name"`
	Symbol                       string   `json:"symbol"`
	Address                      string   `json:"address"`
	Network                      string   `json:"network"`
	Decimals                     int64    `json:"decimals"`
	Verified                     bool     `json:"verified"`
	FDV                          float64  `json:"fdv"`
	MarketCap                    float64  `json:"market_cap"`
	Liquidity                    float64  `json:"liquidity"`
	Price                        float64  `json:"price"`
	PriceChange24hPercent        float64  `json:"price_change_24h_percent"`
	Sell24h                      int64    `json:"sell_24h"`
	Sell24hChangePercent         *float64 `json:"sell_24h_change_percent"`
	Buy24h                       int64    `json:"buy_24h"`
	Buy24hChangePercent          *float64 `json:"buy_24h_change_percent"`
	UniqueWallet24h              int64    `json:"unique_wallet_24h"`
	UniqueWallet24hChangePercent *float64 `json:"unique_wallet_24h_change_percent"`
	Trade24h                     int64    `json:"trade_24h"`
	Trade24hChangePercent        *float64 `json:"trade_24h_change_percent"`
	Volume24hChangePercent       *float64 `json:"volume_24h_change_percent"`
	Volume24hUSD                 float64  `json:"volume_24h_usd"`
	LastTradeUnixTime            int64    `json:"last_trade_unix_time"`
	LastTradeHumanTime           string   `json:"last_trade_human_time"`
	UpdatedTime                  int64    `json:"updated_time"`
	CreationTime                 string   `json:"creation_time"`
	IsScaledUIToken              bool     `json:"is_scaled_ui_token"`
	Multiplier                   *float64 `json:"multiplier"`
}

// RespSearchItem represents search result containing token and market results
type RespSearchItem struct {
	Type   string                  `json:"type"` // "token" or "market"
	Result []RespSearchTokenResult `json:"result"`
}

type RespSearchItems = []RespSearchItem

// ============================================================================
// Wallet Net Worth Types
// ============================================================================

// RespWalletNetWorthItem represents token details in wallet net worth response
type RespWalletNetWorthItem struct {
	Address  string  `json:"address"`
	Decimals int64   `json:"decimals"`
	Price    float64 `json:"price"`
	Balance  string  `json:"balance"`
	Amount   float64 `json:"amount"`
	Network  string  `json:"network"`
	Name     string  `json:"name"`
	Symbol   string  `json:"symbol"`
	LogoURI  string  `json:"logo_uri"`
	Value    string  `json:"value"`
}

// RespWalletNetWorthPagination represents pagination details
type RespWalletNetWorthPagination struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
	Total  int64 `json:"total"`
}

// RespWalletNetWorth represents wallet net worth endpoint response
type RespWalletNetWorth struct {
	WalletAddress    string                       `json:"wallet_address"`
	Currency         string                       `json:"currency"`
	TotalValue       string                       `json:"total_value"`
	CurrentTimestamp string                       `json:"current_timestamp"`
	Items            []RespWalletNetWorthItem     `json:"items"`
	Pagination       RespWalletNetWorthPagination `json:"pagination"`
}

// RespWalletNetWorthHistoryItem represents individual net worth history data point
type RespWalletNetWorthHistoryItem struct {
	Timestamp             string  `json:"timestamp"`
	NetWorth              float64 `json:"net_worth"`
	NetWorthChange        float64 `json:"net_worth_change"`
	NetWorthChangePercent float64 `json:"net_worth_change_percent"`
}

// RespWalletNetWorthHistories represents wallet net worth history response
type RespWalletNetWorthHistories struct {
	WalletAddress    string                          `json:"wallet_address"`
	Currency         string                          `json:"currency"`
	CurrentTimestamp string                          `json:"current_timestamp"`
	PastTimestamp    string                          `json:"past_timestamp"`
	History          []RespWalletNetWorthHistoryItem `json:"history"`
}

// RespWalletNetWorthDetailsNetAsset represents details of an individual asset in the wallet
type RespWalletNetWorthDetailsNetAsset struct {
	Symbol       string  `json:"symbol"`
	TokenAddress string  `json:"token_address"`
	Decimal      int64   `json:"decimal"`
	Balance      string  `json:"balance"`
	Price        float64 `json:"price"`
	Value        float64 `json:"value"`
}

// RespWalletNetWorthDetails represents wallet net worth details response
type RespWalletNetWorthDetails struct {
	WalletAddress      string                              `json:"wallet_address"`
	Currency           string                              `json:"currency"`
	NetWorth           float64                             `json:"net_worth"`
	RequestedTimestamp string                              `json:"requested_timestamp"`
	ResolvedTimestamp  string                              `json:"resolved_timestamp"`
	NetAssets          []RespWalletNetWorthDetailsNetAsset `json:"net_assets"`
}

// ============================================================================
// Wallet PnL Types
// ============================================================================

// WalletPnLTokenCounts represents trade count statistics
type WalletPnLTokenCounts struct {
	TotalBuy   int64 `json:"total_buy"`
	TotalSell  int64 `json:"total_sell"`
	TotalTrade int64 `json:"total_trade"`
}

// WalletPnLTokenQuantity represents token quantity statistics
type WalletPnLTokenQuantity struct {
	TotalBoughtAmount float64 `json:"total_bought_amount"`
	TotalSoldAmount   float64 `json:"total_sold_amount"`
	Holding           float64 `json:"holding"`
}

// WalletPnLTokenCashflow represents USD cashflow statistics
type WalletPnLTokenCashflow struct {
	CostOfQuantitySold float64 `json:"cost_of_quantity_sold"`
	TotalInvested      float64 `json:"total_invested"`
	TotalSold          float64 `json:"total_sold"`
	CurrentValue       float64 `json:"current_value"`
}

// WalletPnLTokenPnL represents profit/loss metrics
type WalletPnLTokenPnL struct {
	RealizedProfitUSD     float64 `json:"realized_profit_usd"`
	RealizedProfitPercent float64 `json:"realized_profit_percent"`
	UnrealizedUSD         float64 `json:"unrealized_usd"`
	UnrealizedPercent     float64 `json:"unrealized_percent"`
	TotalUSD              float64 `json:"total_usd"`
	TotalPercent          float64 `json:"total_percent"`
	AvgProfitPerTradeUSD  float64 `json:"avg_profit_per_trade_usd"`
}

// WalletPnLTokenPricing represents token pricing data
type WalletPnLTokenPricing struct {
	CurrentPrice float64  `json:"current_price"`
	AvgBuyCost   float64  `json:"avg_buy_cost"`
	AvgSellCost  *float64 `json:"avg_sell_cost"`
}

// WalletPnLTokenStats represents per-token statistics
type WalletPnLTokenStats struct {
	Symbol      string                 `json:"symbol"`
	Decimals    int64                  `json:"decimals"`
	Counts      WalletPnLTokenCounts   `json:"counts"`
	Quantity    WalletPnLTokenQuantity `json:"quantity"`
	CashflowUSD WalletPnLTokenCashflow `json:"cashflow_usd"`
	PnL         WalletPnLTokenPnL      `json:"pnl"`
	Pricing     WalletPnLTokenPricing  `json:"pricing"`
}

// WalletPnLMeta represents metadata about the PnL request
type WalletPnLMeta struct {
	Address      string `json:"address"`
	Currency     string `json:"currency"`
	HoldingCheck bool   `json:"holding_check"`
	Time         string `json:"time"`
}

// RespWalletTokensPnL represents wallet tokens PnL endpoint response
type RespWalletTokensPnL struct {
	Meta   WalletPnLMeta                  `json:"meta"`
	Tokens map[string]WalletPnLTokenStats `json:"tokens"`
}

// WalletsPnLByTokenMetadata represents token metadata information
type WalletsPnLByTokenMetadata struct {
	Symbol   string `json:"symbol"`
	Decimals int64  `json:"decimals"`
}

// RespWalletsPnLByToken represents wallet PnL by token endpoint response
type RespWalletsPnLByToken struct {
	TokenMetadata WalletsPnLByTokenMetadata      `json:"token_metadata"`
	Data          map[string]WalletPnLTokenStats `json:"data"`
}

// ============================================================================
// Wallet Transaction Types
// ============================================================================

// WalletTxBalanceChange represents token balance change details
type WalletTxBalanceChange struct {
	Amount          int64    `json:"amount"`
	Symbol          string   `json:"symbol"`
	Name            string   `json:"name"`
	Decimals        int64    `json:"decimals"`
	Address         string   `json:"address"`
	LogoURI         string   `json:"logoURI"`
	IsScaledUIToken bool     `json:"isScaledUiToken"`
	Multiplier      *float64 `json:"multiplier"`
}

// WalletTxContractLabelMetadata represents contract metadata
type WalletTxContractLabelMetadata struct {
	Icon string `json:"icon"`
}

// WalletTxContractLabel represents contract label information
type WalletTxContractLabel struct {
	Address  string                        `json:"address"`
	Name     string                        `json:"name"`
	Metadata WalletTxContractLabelMetadata `json:"metadata"`
}

// WalletTxTokenTransfer represents token transfer details
type WalletTxTokenTransfer struct {
	FromTokenAccount string   `json:"fromTokenAccount"`
	ToTokenAccount   string   `json:"toTokenAccount"`
	FromUserAccount  string   `json:"fromUserAccount"`
	ToUserAccount    string   `json:"toUserAccount"`
	TokenAmount      float64  `json:"tokenAmount"`
	Mint             string   `json:"mint"`
	TransferNative   bool     `json:"transferNative"`
	IsScaledUIToken  bool     `json:"isScaledUiToken"`
	Multiplier       *float64 `json:"multiplier"`
}

// RespWalletTx represents wallet transaction details response
type RespWalletTx struct {
	TxHash         string                  `json:"txHash"`
	BlockNumber    int64                   `json:"blockNumber"`
	BlockTime      string                  `json:"blockTime"`
	Status         bool                    `json:"status"`
	From           string                  `json:"from"`
	To             string                  `json:"to"`
	Fee            int64                   `json:"fee"`
	MainAction     string                  `json:"mainAction"`
	BalanceChange  []WalletTxBalanceChange `json:"balanceChange"`
	ContractLabel  WalletTxContractLabel   `json:"contractLabel"`
	TokenTransfers []WalletTxTokenTransfer `json:"tokenTransfers"`
}

// RespWalletTxs represents wallet transactions by chain
type RespWalletTxs = map[Chain][]RespWalletTx

// ============================================================================
// Token Additional Types
// ============================================================================

// RespTokenTopTraderItem represents top trader details for a token
type RespTokenTopTraderItem struct {
	TokenAddress    string   `json:"tokenAddress"`
	Owner           string   `json:"owner"`
	Tags            []string `json:"tags"`
	Type            string   `json:"type"`
	Volume          float64  `json:"volume"`
	Trade           int64    `json:"trade"`
	TradeBuy        int64    `json:"tradeBuy"`
	TradeSell       int64    `json:"tradeSell"`
	VolumeBuy       float64  `json:"volumeBuy"`
	VolumeSell      float64  `json:"volumeSell"`
	IsScaledUIToken bool     `json:"isScaledUiToken"`
	Multiplier      *float64 `json:"multiplier"`
}

// RespTokenTopTraders is a list of top traders
type RespTokenTopTraders = []RespTokenTopTraderItem

// RespTokenAllMarketListTokenInfo represents token information in market list
type RespTokenAllMarketListTokenInfo struct {
	Address  string `json:"address"`
	Decimals int64  `json:"decimals"`
	Symbol   string `json:"symbol"`
	Icon     string `json:"icon"`
}

// RespTokenAllMarketListItem represents individual market item
type RespTokenAllMarketListItem struct {
	Address                      string                          `json:"address"`
	Base                         RespTokenAllMarketListTokenInfo `json:"base"`
	CreatedAt                    string                          `json:"createdAt"`
	Liquidity                    float64                         `json:"liquidity"`
	Name                         string                          `json:"name"`
	Price                        *float64                        `json:"price"`
	Quote                        RespTokenAllMarketListTokenInfo `json:"quote"`
	Source                       string                          `json:"source"`
	Trade24h                     int64                           `json:"trade24h"`
	Trade24hChangePercent        float64                         `json:"trade24hChangePercent"`
	UniqueWallet24h              int64                           `json:"uniqueWallet24h"`
	UniqueWallet24hChangePercent float64                         `json:"uniqueWallet24hChangePercent"`
	Volume24h                    float64                         `json:"volume24h"`
}

// RespTokenAllMarketList represents token all market list response
type RespTokenAllMarketList struct {
	Items []RespTokenAllMarketListItem `json:"items"`
	Total int64                        `json:"total"`
}

// RespTrendingToken represents individual trending token details
type RespTrendingToken struct {
	Address                string  `json:"address"`
	Decimals               int64   `json:"decimals"`
	Liquidity              float64 `json:"liquidity"`
	LogoURI                string  `json:"logoURI"`
	Name                   string  `json:"name"`
	Symbol                 string  `json:"symbol"`
	Volume24hUSD           float64 `json:"volume24hUSD"`
	Volume24hChangePercent float64 `json:"volume24hChangePercent"`
	Rank                   int64   `json:"rank"`
	Price                  float64 `json:"price"`
	Price24hChangePercent  float64 `json:"price24hChangePercent"`
	FDV                    float64 `json:"fdv"`
	MarketCap              float64 `json:"marketcap"`
}

// RespTokenTrendingList represents trending token list response
type RespTokenTrendingList struct {
	UpdateUnixTime int64               `json:"updateUnixTime"`
	UpdateTime     string              `json:"updateTime"`
	Tokens         []RespTrendingToken `json:"tokens"`
	Total          int64               `json:"total"`
}

// RespTokenHolderBatchItem represents token holder batch response item
type RespTokenHolderBatchItem struct {
	Balance  string  `json:"balance"`
	Decimals int64   `json:"decimals"`
	Mint     string  `json:"mint"`
	Owner    string  `json:"owner"`
	Amount   float64 `json:"amount"`
}

// RespTokenHolderBatch is a list of token holder batch items
type RespTokenHolderBatch = []RespTokenHolderBatchItem

// RespTokenExitLiquidityPrice represents price information in exit liquidity response
type RespTokenExitLiquidityPrice struct {
	Value           float64 `json:"value"`
	UpdateUnixTime  int64   `json:"update_unix_time"`
	UpdateHumanTime string  `json:"update_human_time"`
	UpdateInSlot    int64   `json:"update_in_slot"`
}

// RespTokenExitLiquidity represents token exit liquidity information
type RespTokenExitLiquidity struct {
	Token         string                      `json:"token"`
	ExitLiquidity float64                     `json:"exit_liquidity"`
	Liquidity     float64                     `json:"liquidity"`
	Price         RespTokenExitLiquidityPrice `json:"price"`
	Currency      string                      `json:"currency"`
	Address       string                      `json:"address"`
	Name          string                      `json:"name"`
	Symbol        string                      `json:"symbol"`
	Decimals      int64                       `json:"decimals"`
	Extensions    TokenExtensions             `json:"extensions"`
	LogoURI       string                      `json:"logo_uri"`
}

// RespTokensExitLiquidity is a list of exit liquidity items
type RespTokensExitLiquidity = []RespTokenExitLiquidity
