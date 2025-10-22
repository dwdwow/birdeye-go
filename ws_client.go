package birdeye

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/gorilla/websocket"
)

// ============================================================================
// WebSocket Constants
// ============================================================================

// QueryType represents the type of query for WebSocket subscriptions
type QueryType string

const (
	QueryTypeSimple  QueryType = "simple"
	QueryTypeComplex QueryType = "complex"
)

// WsIntervalType represents time intervals for price data
type WsIntervalType string

const (
	// 1s, 15s, 30s only for Solana
	WsInterval1s  WsIntervalType = "1s"
	WsInterval15s WsIntervalType = "15s"
	WsInterval30s WsIntervalType = "30s"

	WsInterval1m  WsIntervalType = "1m"
	WsInterval3m  WsIntervalType = "3m"
	WsInterval5m  WsIntervalType = "5m"
	WsInterval15m WsIntervalType = "15m"
	WsInterval30m WsIntervalType = "30m"

	WsInterval1H  WsIntervalType = "1H"
	WsInterval2H  WsIntervalType = "2H"
	WsInterval4H  WsIntervalType = "4H"
	WsInterval6H  WsIntervalType = "6H"
	WsInterval8H  WsIntervalType = "8H"
	WsInterval12H WsIntervalType = "12H"

	WsInterval1D WsIntervalType = "1D"
	WsInterval3D WsIntervalType = "3D"
	WsInterval1W WsIntervalType = "1W"
	WsInterval1M WsIntervalType = "1M"
)

// SubUnsubType represents subscription/unsubscription action types
type SubUnsubType string

const (
	SubscribePrice             SubUnsubType = "SUBSCRIBE_PRICE"
	SubscribeBaseQuotePrice    SubUnsubType = "SUBSCRIBE_BASE_QUOTE_PRICE"
	SubscribeTxs               SubUnsubType = "SUBSCRIBE_TXS"
	SubscribeTokenNewListing   SubUnsubType = "SUBSCRIBE_TOKEN_NEW_LISTING"
	SubscribeNewPair           SubUnsubType = "SUBSCRIBE_NEW_PAIR"
	SubscribeLargeTradeTxs     SubUnsubType = "SUBSCRIBE_LARGE_TRADE_TXS"
	SubscribeWalletTxs         SubUnsubType = "SUBSCRIBE_WALLET_TXS"
	SubscribeTokenStats        SubUnsubType = "SUBSCRIBE_TOKEN_STATS"
	UnsubscribePrice           SubUnsubType = "UNSUBSCRIBE_PRICE"
	UnsubscribeBaseQuotePrice  SubUnsubType = "UNSUBSCRIBE_BASE_QUOTE_PRICE"
	UnsubscribeTxs             SubUnsubType = "UNSUBSCRIBE_TXS"
	UnsubscribeTokenNewListing SubUnsubType = "UNSUBSCRIBE_TOKEN_NEW_LISTING"
	UnsubscribeNewPair         SubUnsubType = "UNSUBSCRIBE_NEW_PAIR"
	UnsubscribeLargeTradeTxs   SubUnsubType = "UNSUBSCRIBE_LARGE_TRADE_TXS"
	UnsubscribeWalletTxs       SubUnsubType = "UNSUBSCRIBE_WALLET_TXS"
	UnsubscribeTokenStats      SubUnsubType = "UNSUBSCRIBE_TOKEN_STATS"
)

// WsDataType represents the type of data received from WebSocket
type WsDataType string

const (
	WsDataWelcome             WsDataType = "WELCOME"
	WsDataError               WsDataType = "ERROR"
	WsDataPriceData           WsDataType = "PRICE_DATA"
	WsDataTxsData             WsDataType = "TXS_DATA"
	WsDataBaseQuotePriceData  WsDataType = "BASE_QUOTE_PRICE_DATA"
	WsDataTokenNewListingData WsDataType = "TOKEN_NEW_LISTING_DATA"
	WsDataNewPairData         WsDataType = "NEW_PAIR_DATA"
	WsDataTxsLargeTradeData   WsDataType = "TXS_LARGE_TRADE_DATA"
	WsDataWalletTxsData       WsDataType = "WALLET_TXS_DATA"
	WsDataTokenStatsData      WsDataType = "TOKEN_STATS_DATA"
)

// CurrencyType represents the currency type for price data
type CurrencyType string

const (
	CurrencyUSD  CurrencyType = "usd"
	CurrencyPair CurrencyType = "pair"
)

// ============================================================================
// WebSocket Subscription Data Structures
// ============================================================================

// SubDataPrice represents price subscription data
type SubDataPrice struct {
	Address   string         `json:"address"`
	ChartType WsIntervalType `json:"chartType"`
	Currency  CurrencyType   `json:"currency"`
	QueryType QueryType      `json:"queryType"`
}

// Query returns the query string for complex subscriptions
func (s *SubDataPrice) Query() string {
	return fmt.Sprintf("(address=%s AND chartType=%s AND currency=%s AND queryType=%s)",
		s.Address, s.ChartType, s.Currency, s.QueryType)
}

// Payload returns the JSON payload for subscription
func (s *SubDataPrice) Payload() ([]byte, error) {
	msg := map[string]any{
		"type": SubscribePrice,
		"data": s,
	}
	data, err := json.Marshal(msg)
	return data, err
}

// PricesComplexPayload creates a complex query payload for multiple price subscriptions
func PricesComplexPayload(prices []SubDataPrice) ([]byte, error) {
	queries := make([]string, len(prices))
	for i, price := range prices {
		queries[i] = price.Query()
	}

	msg := map[string]any{
		"type": SubscribePrice,
		"data": map[string]any{
			"queryType": QueryTypeComplex,
			"query":     strings.Join(queries, " OR "),
		},
	}
	data, err := json.Marshal(msg)
	return data, err
}

// SubDataTxs represents transaction subscription data
type SubDataTxs struct {
	Address     *string   `json:"address,omitempty"`
	PairAddress *string   `json:"pairAddress,omitempty"`
	QueryType   QueryType `json:"queryType"`
}

// Query returns the query string for complex subscriptions
func (s *SubDataTxs) Query() string {
	if s.Address != nil {
		return fmt.Sprintf("address=%s", *s.Address)
	}
	if s.PairAddress != nil {
		return fmt.Sprintf("pairAddress=%s", *s.PairAddress)
	}
	return ""
}

// Payload returns the JSON payload for subscription
func (s *SubDataTxs) Payload() ([]byte, error) {
	if s.Address == nil && s.PairAddress == nil {
		return nil, errors.New("birdeye: either address or pairAddress must be provided")
	}
	msg := map[string]any{
		"type": SubscribeTxs,
		"data": s,
	}
	data, err := json.Marshal(msg)
	return data, err
}

// TxsComplexPayload creates a complex query payload for multiple transaction subscriptions
func TxsComplexPayload(txs []SubDataTxs) ([]byte, error) {
	queries := make([]string, 0, len(txs))
	for _, tx := range txs {
		if q := tx.Query(); q != "" {
			queries = append(queries, q)
		}
	}

	msg := map[string]any{
		"type": SubscribeTxs,
		"data": map[string]any{
			"queryType": QueryTypeComplex,
			"query":     strings.Join(queries, " OR "),
		},
	}
	data, err := json.Marshal(msg)
	return data, err
}

// SubDataBaseQuotePrice represents base/quote price subscription data
type SubDataBaseQuotePrice struct {
	BaseAddress  string         `json:"baseAddress"`
	QuoteAddress string         `json:"quoteAddress"`
	ChartType    WsIntervalType `json:"chartType"`
}

// Payload returns the JSON payload for subscription
func (s *SubDataBaseQuotePrice) Payload() ([]byte, error) {
	msg := map[string]any{
		"type": SubscribeBaseQuotePrice,
		"data": s,
	}
	data, err := json.Marshal(msg)
	return data, err
}

// SubDataTokenNewListing represents new token listing subscription data
type SubDataTokenNewListing struct {
	MemePlatformEnabled *bool    `json:"meme_plateform_enabled,omitempty"`
	MinLiquidity        *float64 `json:"min_liquidity,omitempty"`
	MaxLiquidity        *float64 `json:"max_liquidity,omitempty"`
}

// Payload returns the JSON payload for subscription
func (s *SubDataTokenNewListing) Payload() ([]byte, error) {
	msg := map[string]any{
		"type": SubscribeTokenNewListing,
		"data": s,
	}
	data, err := json.Marshal(msg)
	return data, err
}

// SubDataNewPair represents new pair subscription data
type SubDataNewPair struct {
	MinLiquidity *float64 `json:"min_liquidity,omitempty"`
	MaxLiquidity *float64 `json:"max_liquidity,omitempty"`
}

// Payload returns the JSON payload for subscription
func (s *SubDataNewPair) Payload() ([]byte, error) {
	msg := map[string]any{
		"type": SubscribeNewPair,
		"data": s,
	}
	data, err := json.Marshal(msg)
	return data, err
}

// SubDataLargeTradeTxs represents large trade subscription data
type SubDataLargeTradeTxs struct {
	MinVolume float64  `json:"min_volume"`
	MaxVolume *float64 `json:"max_volume,omitempty"`
}

// Payload returns the JSON payload for subscription
func (s *SubDataLargeTradeTxs) Payload() ([]byte, error) {
	msg := map[string]any{
		"type":       SubscribeLargeTradeTxs,
		"min_volume": s.MinVolume,
	}
	if s.MaxVolume != nil {
		msg["max_volume"] = *s.MaxVolume
	}
	data, err := json.Marshal(msg)
	return data, err
}

// SubDataWalletTxs represents wallet transaction subscription data
type SubDataWalletTxs struct {
	Address string `json:"address"`
}

// Payload returns the JSON payload for subscription
func (s *SubDataWalletTxs) Payload() ([]byte, error) {
	msg := map[string]any{
		"type": SubscribeWalletTxs,
		"data": s,
	}
	data, err := json.Marshal(msg)
	return data, err
}

// TokenStatsSelectTradeData represents trade data selection for token stats
type TokenStatsSelectTradeData struct {
	Volume             bool             `json:"volume"`
	Trade              bool             `json:"trade"`
	PriceHistory       bool             `json:"price_history"`
	VolumeHistory      bool             `json:"volume_history"`
	PriceChange        bool             `json:"price_change"`
	TradeHistory       bool             `json:"trade_history"`
	TradeChange        bool             `json:"trade_change"`
	VolumeChange       bool             `json:"volume_change"`
	UniqueWallet       bool             `json:"unique_wallet"`
	UniqueWalletChange bool             `json:"unique_wallet_change"`
	Intervals          []WsIntervalType `json:"intervals"`
}

// NewTokenStatsSelectTradeData creates a new instance with default values
func NewTokenStatsSelectTradeData() TokenStatsSelectTradeData {
	return TokenStatsSelectTradeData{
		Volume:             true,
		Trade:              true,
		PriceHistory:       true,
		VolumeHistory:      true,
		PriceChange:        true,
		TradeHistory:       true,
		TradeChange:        true,
		VolumeChange:       true,
		UniqueWallet:       true,
		UniqueWalletChange: false,
		Intervals:          []WsIntervalType{WsInterval30m, WsInterval1H, WsInterval2H, WsInterval4H, WsInterval8H, "24h"},
	}
}

// TokenStatsSelect represents token stats selection
type TokenStatsSelect struct {
	Price     bool                      `json:"price"`
	TradeData TokenStatsSelectTradeData `json:"trade_data"`
	FDV       bool                      `json:"fdv"`
	MarketCap bool                      `json:"marketcap"`
	Supply    bool                      `json:"supply"`
	LastTrade bool                      `json:"last_trade"`
	Liquidity bool                      `json:"liquidity"`
}

// NewTokenStatsSelect creates a new instance with default values
func NewTokenStatsSelect() TokenStatsSelect {
	return TokenStatsSelect{
		Price:     true,
		TradeData: NewTokenStatsSelectTradeData(),
		FDV:       true,
		MarketCap: true,
		Supply:    true,
		LastTrade: true,
		Liquidity: true,
	}
}

// SubDataTokenStats represents token stats subscription data
type SubDataTokenStats struct {
	Address string           `json:"address"`
	Select  TokenStatsSelect `json:"select"`
}

// Payload returns the JSON payload for subscription
func (s *SubDataTokenStats) Payload() ([]byte, error) {
	msg := map[string]any{
		"type": SubscribeTokenStats,
		"data": s,
	}
	data, err := json.Marshal(msg)
	return data, err
}

// ============================================================================
// WebSocket Response Data Structures
// ============================================================================

// WsData represents a generic WebSocket message
type WsData struct {
	Type WsDataType      `json:"type"`
	Data json.RawMessage `json:"data"`
}

// WsDataPrice represents price data from WebSocket
type WsDataPrice struct {
	O         float64        `json:"o"`         // Open price
	H         float64        `json:"h"`         // High price
	L         float64        `json:"l"`         // Low price
	C         float64        `json:"c"`         // Close price
	EventType string         `json:"eventType"` // Should be "ohlcv"
	Type      WsIntervalType `json:"type"`      // Interval type
	UnixTime  int64          `json:"unixTime"`
	V         float64        `json:"v"` // Volume
	Symbol    string         `json:"symbol"`
	Address   string         `json:"address"`
}

// WsDataTxsTokenInfo represents token info in transaction data
type WsDataTxsTokenInfo struct {
	Address        string   `json:"address"`
	Amount         int64    `json:"amount"`
	ChangeAmount   int64    `json:"changeAmount"`
	Decimals       int64    `json:"decimals"`
	NearestPrice   float64  `json:"nearestPrice"`
	Price          *float64 `json:"price"`
	Symbol         string   `json:"symbol"`
	Type           string   `json:"type"`
	TypeSwap       string   `json:"typeSwap"`
	UIAmount       float64  `json:"uiAmount"`
	UIChangeAmount float64  `json:"uiChangeAmount"`
	FeeInfo        any      `json:"feeInfo,omitempty"`
}

// WsDataTxs represents transaction data from WebSocket
type WsDataTxs struct {
	BlockUnixTime int64              `json:"blockUnixTime"`
	Owner         string             `json:"owner"`
	Source        string             `json:"source"`
	TxHash        string             `json:"txHash"`
	Side          string             `json:"side"` // "buy" or "sell"
	TokenAddress  string             `json:"tokenAddress"`
	Alias         *string            `json:"alias"`
	IsTradeOnBe   bool               `json:"isTradeOnBe"`
	Platform      string             `json:"platform"`
	PricePair     float64            `json:"pricePair"`
	VolumeUSD     float64            `json:"volumeUSD"`
	From          WsDataTxsTokenInfo `json:"from"`
	To            WsDataTxsTokenInfo `json:"to"`
	PriceMark     bool               `json:"priceMark"`
	TokenPrice    float64            `json:"tokenPrice"`
	Network       string             `json:"network"`
	PoolID        string             `json:"poolId"`
}

// WsDataBaseQuotePrice represents base/quote price data from WebSocket
type WsDataBaseQuotePrice struct {
	O            float64        `json:"o"`            // Open price
	H            float64        `json:"h"`            // High price
	L            float64        `json:"l"`            // Low price
	C            float64        `json:"c"`            // Close price
	EventType    string         `json:"eventType"`    // Should be "ohlcv"
	Type         WsIntervalType `json:"type"`         // Time interval
	UnixTime     int64          `json:"unixTime"`     // Unix timestamp
	V            float64        `json:"v"`            // Volume
	BaseAddress  string         `json:"baseAddress"`  // Base token address
	QuoteAddress string         `json:"quoteAddress"` // Quote token address
}

// WsDataTokenNewListing represents new token listing data
type WsDataTokenNewListing struct {
	Address          string `json:"address"`
	Decimals         int64  `json:"decimals"`
	Name             string `json:"name"`
	Symbol           string `json:"symbol"`
	Liquidity        string `json:"liquidity"`
	LiquidityAddedAt int64  `json:"liquidityAddedAt"`
}

// WsDataNewPairTokenInfo represents token info in new pair data
type WsDataNewPairTokenInfo struct {
	Address  string `json:"address"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Decimals int64  `json:"decimals"`
}

// WsDataNewPair represents new pair data from WebSocket
type WsDataNewPair struct {
	Address   string                 `json:"address"`
	Name      string                 `json:"name"`
	Source    string                 `json:"source"`
	Base      WsDataNewPairTokenInfo `json:"base"`
	Quote     WsDataNewPairTokenInfo `json:"quote"`
	TxHash    string                 `json:"txHash"`
	BlockTime int64                  `json:"blockTime"`
}

// WsDataLargeTradeTxsTokenInfo represents token info in large trade data
type WsDataLargeTradeTxsTokenInfo struct {
	Symbol         string   `json:"symbol"`
	Decimals       int64    `json:"decimals"`
	Address        string   `json:"address"`
	UIAmount       float64  `json:"uiAmount"`
	Price          *float64 `json:"price"`
	NearestPrice   float64  `json:"nearestPrice"`
	UIChangeAmount float64  `json:"uiChangeAmount"`
}

// WsDataLargeTradeTxs represents large trade transaction data
type WsDataLargeTradeTxs struct {
	BlockUnixTime  int64                        `json:"blockUnixTime"`
	BlockHumanTime string                       `json:"blockHumanTime"`
	Owner          string                       `json:"owner"`
	Source         string                       `json:"source"`
	PoolAddress    string                       `json:"poolAddress"`
	TxHash         string                       `json:"txHash"`
	VolumeUSD      float64                      `json:"volumeUSD"`
	Network        string                       `json:"network"`
	From           WsDataLargeTradeTxsTokenInfo `json:"from"`
	To             WsDataLargeTradeTxsTokenInfo `json:"to"`
}

// WsDataWalletMintAddLiquidityTxTokenInfo represents token info in wallet mint/add liquidity tx
type WsDataWalletMintAddLiquidityTxTokenInfo struct {
	Symbol   string  `json:"symbol"`
	Decimals int64   `json:"decimals"`
	Address  string  `json:"address"`
	UIAmount float64 `json:"uiAmount"`
}

// WsDataWalletMintAddLiquidityTx represents wallet mint/add liquidity transaction
type WsDataWalletMintAddLiquidityTx struct {
	Type           string                                  `json:"type"`
	BlockUnixTime  int64                                   `json:"blockUnixTime"`
	BlockHumanTime string                                  `json:"blockHumanTime"`
	Owner          string                                  `json:"owner"`
	Source         string                                  `json:"source"`
	TxHash         string                                  `json:"txHash"`
	VolumeUSD      float64                                 `json:"volumeUSD"`
	Network        string                                  `json:"network"`
	Base           WsDataWalletMintAddLiquidityTxTokenInfo `json:"base"`
	Quote          WsDataWalletMintAddLiquidityTxTokenInfo `json:"quote"`
}

// WsDataWalletSwapTxTokenInfo represents token info in wallet swap tx
type WsDataWalletSwapTxTokenInfo struct {
	Symbol         string  `json:"symbol"`
	Decimals       int64   `json:"decimals"`
	Address        string  `json:"address"`
	UIAmount       float64 `json:"uiAmount"`
	Amount         int64   `json:"amount"`
	Price          float64 `json:"price"`
	NearestPrice   float64 `json:"nearestPrice"`
	UIChangeAmount float64 `json:"uiChangeAmount"`
}

// WsDataWalletSwapTx represents wallet swap transaction
type WsDataWalletSwapTx struct {
	Type                string                      `json:"type"`
	BlockUnixTime       int64                       `json:"blockUnixTime"`
	BlockHumanTime      string                      `json:"blockHumanTime"`
	Owner               string                      `json:"owner"`
	Source              string                      `json:"source"`
	PoolAddress         string                      `json:"poolAddress"`
	TxHash              string                      `json:"txHash"`
	VolumeUSD           float64                     `json:"volumeUSD"`
	Network             string                      `json:"network"`
	From                WsDataWalletSwapTxTokenInfo `json:"from"`
	To                  WsDataWalletSwapTxTokenInfo `json:"to"`
	InteractedProgramID string                      `json:"interactedProgramId"`
	LogIndex            int64                       `json:"logIndex"`
	InsIndex            int64                       `json:"insIndex"`
	BlockNumber         int64                       `json:"blockNumber"`
}

// WsDataTokenStats represents token statistics data
type WsDataTokenStats struct {
	Price                      float64 `json:"price"`
	LastTradeHumanTime         string  `json:"last_trade_human_time"`
	LastTradeUnixTime          int64   `json:"last_trade_unix_time"`
	CirculatingSupply          float64 `json:"circulating_supply"`
	TotalSupply                float64 `json:"total_supply"`
	FDV                        float64 `json:"fdv"`
	MarketCap                  float64 `json:"marketcap"`
	Liquidity                  float64 `json:"liquidity"`
	Volume30mUSD               float64 `json:"volume_30m_usd"`
	Volume30m                  float64 `json:"volume_30m"`
	VolumeBuy30m               float64 `json:"volume_buy_30m"`
	VolumeBuy30mUSD            float64 `json:"volume_buy_30m_usd"`
	VolumeSell30m              float64 `json:"volume_sell_30m"`
	VolumeSell30mUSD           float64 `json:"volume_sell_30m_usd"`
	Trade30m                   int64   `json:"trade_30m"`
	Buy30m                     int64   `json:"buy_30m"`
	Sell30m                    int64   `json:"sell_30m"`
	VolumeHistory30m           float64 `json:"volume_history_30m"`
	VolumeHistory30mUSD        float64 `json:"volume_history_30m_usd"`
	VolumeSellHistory30mUSD    float64 `json:"volume_sell_history_30m_usd"`
	VolumeBuyHistory30mUSD     float64 `json:"volume_buy_history_30m_usd"`
	PriceChange30mPercent      float64 `json:"price_change_30m_percent"`
	TradeHistory30m            int64   `json:"trade_history_30m"`
	BuyHistory30m              int64   `json:"buy_history_30m"`
	SellHistory30m             int64   `json:"sell_history_30m"`
	Trade30mChangePercent      float64 `json:"trade_30m_change_percent"`
	Buy30mChangePercent        float64 `json:"buy_30m_change_percent"`
	Sell30mChangePercent       float64 `json:"sell_30m_change_percent"`
	Volume30mChangePercent     float64 `json:"volume_30m_change_percent"`
	VolumeBuy30mChangePercent  float64 `json:"volume_buy_30m_change_percent"`
	VolumeSell30mChangePercent float64 `json:"volume_sell_30m_change_percent"`
	UniqueWallet30m            int64   `json:"unique_wallet_30m"`
}

// ============================================================================
// WebSocket Client
// ============================================================================

// WSClient represents a WebSocket client for Birdeye API
type WSClient struct {
	APIKey string
	Chain  Chain
	conn   *websocket.Conn
	mu     sync.RWMutex
}

// WSClientConfig holds configuration for WebSocket client
type WSClientConfig struct {
	APIKey string
	Chain  Chain
}

// NewWSClient creates a new WebSocket client
func NewWSClient(config WSClientConfig) *WSClient {
	return &WSClient{
		APIKey: config.APIKey,
		Chain:  config.Chain,
	}
}

// Connect establishes WebSocket connection
func (c *WSClient) Connect(ctx context.Context) error {
	uri := fmt.Sprintf("wss://public-api.birdeye.so/socket/%s?x-api-key=%s", c.Chain, c.APIKey)

	header := http.Header{}
	header.Set("Origin", "ws://public-api.birdeye.so")

	dialer := websocket.Dialer{
		Subprotocols: []string{"echo-protocol"},
	}

	conn, _, err := dialer.DialContext(ctx, uri, header)
	if err != nil {
		return fmt.Errorf("birdeye: failed to connect to websocket: %w", err)
	}

	c.mu.Lock()
	c.conn = conn
	c.mu.Unlock()

	return nil
}

// readMessages reads messages from WebSocket
func (c *WSClient) Read() (dataType WsDataType, data json.RawMessage, err error) {
	c.mu.RLock()
	conn := c.conn
	c.mu.RUnlock()

	if conn == nil {
		return
	}

	_, message, err := conn.ReadMessage()
	if err != nil {
		err = fmt.Errorf("birdeye: failed to read ws message: %w", err)
		return
	}

	var wsData WsData
	if err = json.Unmarshal(message, &wsData); err != nil {
		err = fmt.Errorf("birdeye: failed to unmarshal ws message: %w", err)
		return
	}

	return wsData.Type, wsData.Data, nil
}

// Send sends a message through WebSocket
func (c *WSClient) Send(payload []byte) error {
	c.mu.RLock()
	defer c.mu.RUnlock()
	conn := c.conn

	if conn == nil {
		return fmt.Errorf("birdeye: conn is nil")
	}

	return conn.WriteMessage(websocket.TextMessage, payload)
}

// Subscribe subscribes to a data stream
func (c *WSClient) Subscribe(payload []byte) error {
	return c.Send(payload)
}

// Unsubscribe unsubscribes from a data stream
func (c *WSClient) Unsubscribe(subType SubUnsubType, data any) error {
	msg := map[string]any{
		"type": subType,
	}
	if data != nil {
		msg["data"] = data
	}

	jsonData, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	return c.Send(jsonData)
}

// Close closes the WebSocket connection
func (c *WSClient) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.conn != nil {
		return c.conn.Close()
	}

	return nil
}
