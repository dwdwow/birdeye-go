package birdeye

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// Birdeye API Client with automatic rate limiting.
//
// This package provides a comprehensive, rate-limited client for the Birdeye API.
// It includes automatic rate limiting, error handling, and type safety for all API endpoints.
//
// Key Features:
// - Automatic rate limiting based on Birdeye's official limits
// - Type-safe API responses with struct definitions
// - Support for multiple blockchain networks
// - Comprehensive error handling and retry logic
//
// Rate Limiting:
// - 300 RPS: Price and market data endpoints
// - 150 RPS: Token list and security endpoints
// - 100 RPS: Historical data and transaction endpoints
// - 30 RPS / 150 RPM: Wallet endpoints (multi-tier limiting)
// - 2 RPS: Scroll endpoints (very strict)

// ============================================================================
// API Endpoints Constants
// ============================================================================

const (
	PublicAPIBaseURL = "https://public-api.birdeye.so"

	EndpointDefiPrice                        = "/defi/price"
	EndpointDefiMultiPrice                   = "/defi/multi_price"
	EndpointDefiOHLCVBaseQuote               = "/defi/ohlcv/base_quote"
	EndpointDefiPriceVolumeSingle            = "/defi/price_volume/single"
	EndpointDefiPriceVolumeMulti             = "/defi/price_volume/multi"
	EndpointDefiTokenOverview                = "/defi/token_overview"
	EndpointDefiV2TokensTopTraders           = "/defi/v2/tokens/top_traders"
	EndpointDefiV3TokenMetadataSingle        = "/defi/v3/token/meta-data/single"
	EndpointDefiV3TokenMetadataMultiple      = "/defi/v3/token/meta-data/multiple"
	EndpointDefiV3TokenMarketData            = "/defi/v3/token/market-data"
	EndpointDefiV3TokenMarketDataMultiple    = "/defi/v3/token/market-data/multiple"
	EndpointDefiV3TokenTradeDataSingle       = "/defi/v3/token/trade-data/single"
	EndpointDefiV3TokenTradeDataMultiple     = "/defi/v3/token/trade-data/multiple"
	EndpointDefiV3AllTimeTradesSingle        = "/defi/v3/all-time/trades/single"
	EndpointDefiV3AllTimeTradesMultiple      = "/defi/v3/all-time/trades/multiple"
	EndpointDefiTokenList                    = "/defi/tokenlist"
	EndpointDefiTokenSecurity                = "/defi/token_security"
	EndpointDefiHistoryPrice                 = "/defi/history_price"
	EndpointDefiHistoricalPriceUnix          = "/defi/historical_price_unix"
	EndpointDefiTxsToken                     = "/defi/txs/token"
	EndpointDefiTxsPair                      = "/defi/txs/pair"
	EndpointDefiTxsTokenSeekByTime           = "/defi/txs/token/seek_by_time"
	EndpointDefiTxsPairSeekByTime            = "/defi/txs/pair/seek_by_time"
	EndpointDefiV3TokenTxs                   = "/defi/v3/token/txs"
	EndpointDefiOHLCV                        = "/defi/ohlcv"
	EndpointDefiOHLCVPair                    = "/defi/ohlcv/pair"
	EndpointDefiV3PairOverviewSingle         = "/defi/v3/pair/overview/single"
	EndpointDefiV3PairOverviewMultiple       = "/defi/v3/pair/overview/multiple"
	EndpointDefiV3Txs                        = "/defi/v3/txs"
	EndpointDefiTokenCreationInfo            = "/defi/token_creation_info"
	EndpointDefiTokenTrending                = "/defi/token_trending"
	EndpointDefiV2TokensNewListing           = "/defi/v2/tokens/new_listing"
	EndpointDefiV2Markets                    = "/defi/v2/markets"
	EndpointDefiV3TokenHolder                = "/defi/v3/token/holder"
	EndpointTokenV1HolderBatch               = "/token/v1/holder/batch"
	EndpointDefiV3TokenMintBurnTxs           = "/defi/v3/token/mint-burn-txs"
	EndpointDefiV3TokenList                  = "/defi/v3/token/list"
	EndpointTraderGainersLosers              = "/trader/gainers-losers"
	EndpointTraderTxsSeekByTime              = "/trader/txs/seek_by_time"
	EndpointDefiNetworks                     = "/defi/networks"
	EndpointDefiV3OHLCV                      = "/defi/v3/ohlcv"
	EndpointDefiV3OHLCVPair                  = "/defi/v3/ohlcv/pair"
	EndpointDefiV3PriceStatsSingle           = "/defi/v3/price/stats/single"
	EndpointDefiV3PriceStatsMultiple         = "/defi/v3/price/stats/multiple"
	EndpointDefiV3TokenExitLiquidity         = "/defi/v3/token/exit-liquidity"
	EndpointDefiV3TokenExitLiquidityMultiple = "/defi/v3/token/exit-liquidity/multiple"
	EndpointDefiV3TokenMemeList              = "/defi/v3/token/meme/list"
	EndpointDefiV3TokenMemeDetailSingle      = "/defi/v3/token/meme/detail/single"
	EndpointV1WalletTokenList                = "/v1/wallet/token_list"
	EndpointV1WalletTxList                   = "/v1/wallet/tx_list"
	EndpointV1WalletTokenBalance             = "/v1/wallet/token_balance"
	EndpointV2WalletCurrentNetWorth          = "/wallet/v2/current-net-worth"
	EndpointV2WalletNetWorth                 = "/wallet/v2/net-worth"
	EndpointV2WalletNetWorthDetails          = "/wallet/v2/net-worth-details"
	EndpointV2WalletPnl                      = "/wallet/v2/pnl"
	EndpointV2WalletPnlMultiple              = "/wallet/v2/pnl/multiple"
	EndpointV2WalletTokenBalance             = "/wallet/v2/token-balance"
	EndpointV2WalletTxFirstFunded            = "/wallet/v2/tx/first-funded"
	EndpointDefiV3Search                     = "/defi/v3/search"
	EndpointDefiV3TxsLatestBlock             = "/defi/v3/txs/latest-block"
	EndpointV1WalletListSupportedChain       = "/v1/wallet/list_supported_chain"
	EndpointDefiV3TokenListScroll            = "/defi/v3/token/list/scroll"
)

// ============================================================================
// Error Types
// ============================================================================

// BirdeyeAPIError is the base error type for all Birdeye API errors
type BirdeyeAPIError struct {
	Message    string
	StatusCode int
	Response   map[string]any
}

func (e *BirdeyeAPIError) Error() string {
	if e.StatusCode > 0 {
		return fmt.Sprintf("[%d] %s", e.StatusCode, e.Message)
	}
	return e.Message
}

// Specific error types
var (
	ErrBadRequest          = errors.New("invalid request parameters or payload")
	ErrUnauthorized        = errors.New("unauthorized: authentication required")
	ErrForbidden           = errors.New("access denied: insufficient permissions")
	ErrNotFound            = errors.New("resource not found")
	ErrUnprocessableEntity = errors.New("invalid data format")
	ErrTooManyRequests     = errors.New("rate limit exceeded")
	ErrInternalServer      = errors.New("internal server error")
	ErrNetwork             = errors.New("network error occurred")
	ErrTimeout             = errors.New("request timeout")
)

// ============================================================================
// HTTPClient Structure
// ============================================================================

// HTTPClient is the Birdeye API client with automatic rate limiting and type safety.
//
// This client provides a comprehensive interface to the Birdeye API with built-in
// rate limiting, error handling, and type safety. All API endpoints are automatically
// rate-limited according to Birdeye's official limits.
//
// Features:
//   - Automatic rate limiting for all endpoints
//   - Type-safe responses with struct definitions
//   - Support for multiple blockchain networks
//   - Comprehensive error handling and retry logic
//   - Context support for cancellation and timeout
//
// Rate Limiting:
//   - 300 RPS: Price and market data endpoints
//   - 150 RPS: Token list and security endpoints
//   - 100 RPS: Historical data and transaction endpoints
//   - 30 RPS / 150 RPM: Wallet endpoints (multi-tier limiting)
//   - 2 RPS: Scroll endpoints (very strict)
//
// Example:
//
//	client, err := birdeye.NewHTTPClient(birdeye.HTTPClientConfig{
//	    APIKey: "your-api-key",
//	    Chains: []birdeye.Chain{birdeye.ChainSolana},
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	// Get token price
//	price, err := client.GetTokenPrice(ctx, "So11111111111111111111111111111111111111112", TokenPriceOptions{})
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("SOL Price: $%.2f\n", price.Value)
type HTTPClient struct {
	apiKey           string
	baseURL          string
	chains           []Chain
	httpClient       *http.Client
	limiter300RPS    *RateLimiter
	limiter150RPS    *RateLimiter
	limiter100RPS    *RateLimiter
	limiterWallet    *MultiRateLimiter
	limiter2RPS      *RateLimiter
	endpointLimiters map[string]any
	onLimitExceeded  RateLimitBehavior
}

// HTTPClientConfig holds configuration for creating a new HTTPClient.
type HTTPClientConfig struct {
	// APIKey is your Birdeye API key. Get one from https://birdeye.so/
	// Required field.
	APIKey string

	// Chains is the list of blockchain networks to query.
	// If nil, queries all supported networks.
	// Optional, default: nil (all networks)
	Chains []Chain

	// BaseURL is the base URL for the Birdeye API.
	// Optional, default: "https://public-api.birdeye.so"
	BaseURL string

	// HTTPClient is the HTTP client to use for requests.
	// If nil, a default client with 30 second timeout will be created.
	// Optional, default: &http.Client{Timeout: 30 * time.Second}
	HTTPClient *http.Client

	// OnLimitExceeded defines behavior when rate limit is exceeded.
	// Options: RateLimitBlock (wait), RateLimitRaise (return error).
	// Optional, default: RateLimitBlock
	OnLimitExceeded RateLimitBehavior
}

// NewHTTPClient creates a new Birdeye API client with automatic rate limiting.
//
// This function initializes the client with the provided configuration and sets up
// automatic rate limiting for all API endpoints according to Birdeye's official limits.
//
// Parameters:
//   - config: Configuration for the HTTP client including API key and options
//
// Returns:
//   - *HTTPClient: Initialized HTTP client ready to make API requests
//   - error: Error if configuration is invalid (e.g., missing API key)
//
// Rate Limiters:
// The client automatically creates and manages rate limiters for different endpoint categories:
//   - 300 RPS: Price and market data endpoints
//   - 150 RPS: Token list and security endpoints
//   - 100 RPS: Historical data and transaction endpoints
//   - 30 RPS / 150 RPM: Wallet endpoints (multi-tier limiting)
//   - 2 RPS: Scroll endpoints (very strict)
//
// Example:
//
//	// Basic initialization
//	client, err := birdeye.NewHTTPClient(birdeye.HTTPClientConfig{
//	    APIKey: "your-api-key",
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	// With specific chains
//	client, err := birdeye.NewHTTPClient(birdeye.HTTPClientConfig{
//	    APIKey: "your-api-key",
//	    Chains: []birdeye.Chain{birdeye.ChainSolana, birdeye.ChainEthereum},
//	})
//
//	// With custom rate limit behavior
//	client, err := birdeye.NewHTTPClient(birdeye.HTTPClientConfig{
//	    APIKey: "your-api-key",
//	    OnLimitExceeded: birdeye.RateLimitRaise,
//	})
func NewHTTPClient(config HTTPClientConfig) *HTTPClient {
	if config.BaseURL == "" {
		config.BaseURL = PublicAPIBaseURL
	}

	if config.HTTPClient == nil {
		config.HTTPClient = &http.Client{
			Timeout: 30 * time.Second,
		}
	}

	if config.OnLimitExceeded == "" {
		config.OnLimitExceeded = RateLimitBlock
	}

	// Create rate limiters
	limiter300, _ := NewRateLimiter(300, time.Second, config.OnLimitExceeded)
	limiter150, _ := NewRateLimiter(150, time.Second, config.OnLimitExceeded)
	limiter100, _ := NewRateLimiter(100, time.Second, config.OnLimitExceeded)
	limiter2, _ := NewRateLimiter(2, time.Second, config.OnLimitExceeded)

	limiterWallet, _ := NewMultiRateLimiter(
		[]RateLimit{
			{Limit: 30, Period: time.Second},
			{Limit: 150, Period: time.Minute},
		},
		config.OnLimitExceeded,
	)

	client := &HTTPClient{
		apiKey:           config.APIKey,
		baseURL:          strings.TrimRight(config.BaseURL, "/"),
		chains:           config.Chains,
		httpClient:       config.HTTPClient,
		limiter300RPS:    limiter300,
		limiter150RPS:    limiter150,
		limiter100RPS:    limiter100,
		limiterWallet:    limiterWallet,
		limiter2RPS:      limiter2,
		onLimitExceeded:  config.OnLimitExceeded,
		endpointLimiters: make(map[string]any),
	}

	// Map endpoints to their rate limiters
	client.initEndpointLimiters()

	return client
}

// initEndpointLimiters initializes the endpoint to limiter mapping
func (c *HTTPClient) initEndpointLimiters() {
	// 300 RPS endpoints
	endpoints300 := []string{
		EndpointDefiPrice, EndpointDefiMultiPrice, EndpointDefiOHLCVBaseQuote,
		EndpointDefiPriceVolumeSingle, EndpointDefiPriceVolumeMulti,
		EndpointDefiTokenOverview, EndpointDefiV2TokensTopTraders,
		EndpointDefiV3TokenMetadataSingle, EndpointDefiV3TokenMetadataMultiple,
		EndpointDefiV3TokenMarketData, EndpointDefiV3TokenMarketDataMultiple,
		EndpointDefiV3TokenTradeDataSingle, EndpointDefiV3TokenTradeDataMultiple,
		EndpointDefiV3AllTimeTradesSingle, EndpointDefiV3AllTimeTradesMultiple,
	}
	for _, ep := range endpoints300 {
		c.endpointLimiters[ep] = c.limiter300RPS
	}

	// 150 RPS endpoints
	endpoints150 := []string{
		EndpointDefiTokenList, EndpointDefiTokenSecurity,
	}
	for _, ep := range endpoints150 {
		c.endpointLimiters[ep] = c.limiter150RPS
	}

	// 100 RPS endpoints
	endpoints100 := []string{
		EndpointDefiHistoryPrice, EndpointDefiHistoricalPriceUnix,
		EndpointDefiTxsToken, EndpointDefiTxsPair,
		EndpointDefiTxsTokenSeekByTime, EndpointDefiTxsPairSeekByTime,
		EndpointDefiV3TokenTxs, EndpointDefiOHLCV, EndpointDefiOHLCVPair,
		EndpointDefiV3PairOverviewSingle, EndpointDefiV3PairOverviewMultiple,
		EndpointDefiV3Txs, EndpointDefiTokenCreationInfo,
		EndpointDefiTokenTrending, EndpointDefiV2TokensNewListing,
		EndpointDefiV2Markets, EndpointDefiV3TokenHolder,
		EndpointDefiV3TokenMintBurnTxs, EndpointDefiV3TokenList,
		EndpointTraderGainersLosers, EndpointTraderTxsSeekByTime,
		EndpointDefiV3Search, EndpointDefiNetworks,
		EndpointDefiV3TxsLatestBlock, EndpointDefiV3OHLCV,
		EndpointDefiV3OHLCVPair, EndpointDefiV3PriceStatsSingle,
		EndpointDefiV3PriceStatsMultiple, EndpointDefiV3TokenExitLiquidity,
		EndpointDefiV3TokenExitLiquidityMultiple, EndpointDefiV3TokenMemeList,
		EndpointDefiV3TokenMemeDetailSingle,
	}
	for _, ep := range endpoints100 {
		c.endpointLimiters[ep] = c.limiter100RPS
	}

	// Wallet endpoints (multi-tier: 30 RPS / 150 RPM)
	endpointsWallet := []string{
		EndpointV1WalletTokenList, EndpointV1WalletTxList,
		EndpointV1WalletTokenBalance, EndpointV1WalletListSupportedChain,
		EndpointV2WalletCurrentNetWorth, EndpointV2WalletNetWorth,
		EndpointV2WalletNetWorthDetails, EndpointV2WalletPnl,
		EndpointV2WalletPnlMultiple, EndpointV2WalletTokenBalance,
		EndpointV2WalletTxFirstFunded,
	}
	for _, ep := range endpointsWallet {
		c.endpointLimiters[ep] = c.limiterWallet
	}

	// 2 RPS endpoint (scroll)
	c.endpointLimiters[EndpointDefiV3TokenListScroll] = c.limiter2RPS
}

// getLimiter gets the appropriate rate limiter for an endpoint
func (c *HTTPClient) getLimiter(endpoint string) any {
	if limiter, ok := c.endpointLimiters[endpoint]; ok {
		return limiter
	}
	return c.limiter100RPS
}

// getHeaders builds the HTTP headers for API requests
func (c *HTTPClient) getHeaders(chains []Chain) http.Header {
	headers := http.Header{
		"Accept":    []string{"application/json"},
		"X-API-KEY": []string{c.apiKey},
	}

	chainsToUse := chains
	if chainsToUse == nil && c.chains != nil {
		chainsToUse = c.chains
	}

	if len(chainsToUse) > 0 {
		chainStrs := make([]string, len(chainsToUse))
		for i, chain := range chainsToUse {
			chainStrs[i] = string(chain)
		}
		headers["X-Chain"] = []string{strings.Join(chainStrs, ",")}
	}

	return headers
}

// requestOptions holds options for making API requests
type requestOptions struct {
	method          string `default:"GET"`
	chains          []Chain
	onLimitExceeded RateLimitBehavior
	paramsUseArray  bool `default:"false"`
	respJustItems   bool `default:"false"`
	paramsOrBody    map[string]any
}

// request makes a rate-limited request to the Birdeye API
func (c *HTTPClient) request(ctx context.Context, endpoint string, opts requestOptions) (map[string]any, error) {
	// Get rate limiter
	limiter := c.getLimiter(endpoint)

	// Acquire rate limit token
	behavior := c.onLimitExceeded
	if opts.onLimitExceeded != "" {
		behavior = opts.onLimitExceeded
	}

	var acquired bool
	var err error

	switch v := limiter.(type) {
	case *RateLimiter:
		acquired, err = v.Acquire(ctx, 1, &opts.onLimitExceeded)
	case *MultiRateLimiter:
		acquired, err = v.Acquire(ctx, 1, &opts.onLimitExceeded)
	default:
		return nil, errors.New("invalid limiter type")
	}

	if err != nil {
		return nil, err
	}

	if !acquired {
		if behavior == RateLimitSkip {
			return nil, nil
		}
		return nil, ErrRateLimitExceeded
	}

	// Build URL
	reqURL := c.baseURL + endpoint

	// Process parameters
	if opts.paramsOrBody != nil && !opts.paramsUseArray {
		for k, v := range opts.paramsOrBody {
			if arr, ok := v.([]string); ok {
				opts.paramsOrBody[k] = strings.Join(arr, ",")
			}
		}
	}

	// Retry logic for network errors
	var resp *http.Response
	maxRetries := 3

	for attempt := range maxRetries {
		var req *http.Request

		if opts.method == "POST" {
			// POST request with JSON body
			bodyData, _ := json.Marshal(opts.paramsOrBody)
			req, err = http.NewRequestWithContext(ctx, "POST", reqURL, bytes.NewReader(bodyData))
			if err != nil {
				return nil, err
			}
		} else {
			// GET request
			req, err = http.NewRequestWithContext(ctx, "GET", reqURL, nil)
			if err != nil {
				return nil, err
			}

			// Add query parameters
			if opts.paramsOrBody != nil {
				q := req.URL.Query()
				for k, v := range opts.paramsOrBody {
					q.Add(k, fmt.Sprintf("%v", v))
				}
				req.URL.RawQuery = q.Encode()
			}
		}

		req.Header = c.getHeaders(opts.chains)
		if opts.method == "POST" {
			req.Header.Set("Content-Type", "application/json")
		}

		// Make request
		resp, err = c.httpClient.Do(req)
		if err != nil {
			if attempt < maxRetries-1 {
				time.Sleep(time.Duration(attempt) * time.Second)
				continue
			}
			return nil, fmt.Errorf("%w: %v", ErrNetwork, err)
		}
		break
	}

	defer resp.Body.Close()

	// Read response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Parse JSON response
	var result map[string]any
	if err := json.Unmarshal(bodyBytes, &result); err != nil {
		// If not valid JSON, return error
		return nil, fmt.Errorf("invalid JSON response: %w", err)
	}

	// Handle non-200 status codes
	if resp.StatusCode != 200 {
		message, _ := result["message"].(string)
		if message == "" {
			message = string(bodyBytes)
		}
		return nil, &BirdeyeAPIError{
			Message:    message,
			StatusCode: resp.StatusCode,
			Response:   result,
		}
	}

	// Extract data field
	data, hasData := result["data"]
	if !hasData {
		return result, nil
	}

	// Handle pagination
	if pagination, ok := result["pagination"].(map[string]any); ok {
		if dataMap, ok := data.(map[string]any); ok {
			dataMap["pagination"] = pagination
			data = dataMap
		} else if dataArr, ok := data.([]any); ok {
			data = map[string]any{
				"items":      dataArr,
				"pagination": pagination,
			}
		}
	}

	// Return just items if requested
	if opts.respJustItems {
		if dataMap, ok := data.(map[string]any); ok {
			if items, ok := dataMap["items"]; ok {
				return map[string]any{"items": items}, nil
			}
		}
	}

	// Return data as map
	if dataMap, ok := data.(map[string]any); ok {
		return dataMap, nil
	}

	// If data is not a map, wrap it
	return map[string]any{"data": data}, nil
}

// ============================================================================
// Helper Functions for Building Parameters
// ============================================================================

// ============================================================================
// API Methods - Network Support
// ============================================================================

// GetSupportedNetworks retrieves the list of supported blockchain networks.
//
// This method retrieves all blockchain networks that are currently supported by the Birdeye API
// for DeFi and token data queries. Use this to discover which networks you can query.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//
// Returns:
//   - []Chain: List of supported blockchain network identifiers (e.g., "solana", "ethereum", "arbitrum")
//   - error: Error if the API request fails or if the response cannot be parsed
//
// Raises:
//   - Context errors: If the context is cancelled or times out
//   - API errors: If the Birdeye API returns an error response
//   - Network errors: If there are connectivity issues
//
// Example:
//
//	networks, err := client.GetSupportedNetworks(ctx)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Supported networks: %d\n", len(networks))
//	for _, network := range networks {
//	    fmt.Println(network)
//	}
//	// Output: solana, ethereum, arbitrum, avalanche, bsc, optimism, polygon, base, zksync, sui
func (c *HTTPClient) GetSupportedNetworks(ctx context.Context) ([]Chain, error) {
	result, err := c.request(ctx, EndpointDefiNetworks, requestOptions{
		method: "GET",
	})
	if err != nil {
		return nil, err
	}

	// Parse chains from result
	var chains []Chain
	if data, ok := result["data"].([]any); ok {
		for _, item := range data {
			if chain, ok := item.(string); ok {
				chains = append(chains, Chain(chain))
			}
		}
	}

	return chains, nil
}

// GetWalletSupportedNetworks retrieves the list of blockchain networks supported for wallet operations.
//
// This method retrieves blockchain networks that support wallet-related operations such as
// portfolio queries, transaction history, balance checks, and net worth calculations.
// Note that wallet endpoints typically support fewer networks than general token/price endpoints.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//
// Returns:
//   - []Chain: List of networks supporting wallet operations (e.g., ["solana"])
//   - error: Error if the API request fails or if the response cannot be parsed
//
// Raises:
//   - Context errors: If the context is cancelled or times out
//   - API errors: If the Birdeye API returns an error response
//   - Network errors: If there are connectivity issues
//
// Example:
//
//	walletNetworks, err := client.GetWalletSupportedNetworks(ctx)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Wallet networks: %v\n", walletNetworks)
//	// Output: [solana]
func (c *HTTPClient) GetWalletSupportedNetworks(ctx context.Context) ([]Chain, error) {
	result, err := c.request(ctx, EndpointV1WalletListSupportedChain, requestOptions{
		method: "GET",
	})
	if err != nil {
		return nil, err
	}

	// Parse chains from result
	var chains []Chain
	if data, ok := result["data"].([]any); ok {
		for _, item := range data {
			if chain, ok := item.(string); ok {
				chains = append(chains, Chain(chain))
			}
		}
	}

	return chains, nil
}

// ============================================================================
// API Methods - Token Price
// ============================================================================

// TokenPriceOptions holds options for GetTokenPrice.
type TokenPriceOptions struct {
	// CheckLiquidity is the minimum liquidity threshold in USD.
	// Tokens below this threshold may not be included in results.
	// Optional, default: 100
	CheckLiquidity int64 `default:"100"`

	// IncludeLiquidity determines whether to include liquidity information in the response.
	// Optional, default: true
	IncludeLiquidity bool `default:"true"`

	// UIAmountMode specifies the token amount display mode.
	// Options:
	//   - "raw": Raw amounts (e.g., 1000000000 for 1 token with 9 decimals)
	//   - "scaled": Human-readable amounts (e.g., 1.0)
	//   - "both": Include both raw and scaled amounts
	// Optional, default: "raw"
	UIAmountMode string `default:"raw"`

	// Chains is the list of blockchain networks to query.
	// If nil, queries all supported networks.
	// Optional, default: nil (all networks)
	Chains []Chain

	// OnLimitExceeded overrides the default rate limit behavior for this request.
	// If nil, uses the client's default behavior.
	// Optional, default: nil (use client default)
	OnLimitExceeded string `default:""`
}

// GetTokenPrice retrieves the current price of a token with optional liquidity filtering.
//
// This method retrieves the current market price for a token, including 24-hour price change
// and liquidity information. The response includes both raw and scaled token amounts based on
// the specified mode.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - address: Token contract address to query
//   - opts: Configuration options (see TokenPriceOptions for details):
//   - CheckLiquidity: Minimum liquidity threshold in USD. Tokens below this threshold
//     may not be included. Optional, default: 100
//   - IncludeLiquidity: Whether to include liquidity information in response.
//     Optional, default: true
//   - UIAmountMode: Token amount display mode.
//     Options: "raw" (e.g., 1000000000), "scaled" (e.g., 1.0), "both"
//     Optional, default: "raw"
//   - Chains: List of blockchain networks to query. If nil, queries all supported networks.
//     Optional, default: nil
//   - OnLimitExceeded: Override rate limit behavior. If nil, uses client default.
//     Optional, default: nil
//
// Returns:
//   - *RespTokenPrice: Price response containing:
//   - Value: Current token price in USD
//   - PriceChange24h: 24-hour price change percentage
//   - Liquidity: Current liquidity in USD
//   - UpdateUnixTime: Last update timestamp (Unix seconds)
//   - UpdateHumanTime: Human-readable update time
//   - PriceInNative: Price in native token (e.g., SOL, ETH)
//   - error: Error if request fails, token address is invalid, or rate limit exceeded
//
// Raises:
//   - Context errors: If ctx is cancelled or times out
//   - API errors: If the Birdeye API returns an error
//   - Validation errors: If address is invalid
//   - Rate limit errors: If rate limit exceeded and OnLimitExceeded is RateLimitRaise
//
// Example:
//
//	// Get SOL price
//	price, err := client.GetTokenPrice(ctx, "So11111111111111111111111111111111111111112", TokenPriceOptions{})
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("SOL Price: $%.2f\n", price.Value)
//	fmt.Printf("24h Change: %.2f%%\n", price.PriceChange24h)
//
//	// Get price with liquidity filtering
//	checkLiq := int64(1000)
//	price, err = client.GetTokenPrice(ctx, "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v", TokenPriceOptions{
//	    CheckLiquidity: &checkLiq,  // Only include if liquidity > $1000
//	    UIAmountMode:   "scaled",
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("USDC Liquidity: $%.2f\n", price.Liquidity)
func (c *HTTPClient) GetTokenPrice(ctx context.Context, address string, opts *TokenPriceOptions) (*RespTokenPrice, error) {
	if opts == nil {
		opts = &TokenPriceOptions{}
	}

	params, err := ApplyDefaultsAndBuildParams(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to apply defaults: %w", err)
	}

	// Add required parameters
	params["address"] = address

	result, err := c.request(ctx, EndpointDefiPrice, requestOptions{
		method:          "GET",
		chains:          opts.Chains,
		onLimitExceeded: RateLimitBehavior(opts.OnLimitExceeded),
		paramsOrBody:    params,
	})
	if err != nil {
		return nil, err
	}

	// Parse result into RespTokenPrice
	var price RespTokenPrice
	data, _ := json.Marshal(result)
	if err := json.Unmarshal(data, &price); err != nil {
		return nil, err
	}

	return &price, nil
}

// MultiTokenPriceOptions holds options for GetMultiTokenPrice.
type MultiTokenPriceOptions struct {
	// CheckLiquidity is the minimum liquidity threshold in USD.
	// Tokens below this threshold may not be included in results.
	// Optional, default: 100
	CheckLiquidity int64 `default:"100"`

	// IncludeLiquidity determines whether to include liquidity information in the response.
	// Optional, default: true
	IncludeLiquidity bool `default:"true"`

	// UIAmountMode specifies the token amount display mode.
	// Options:
	//   - "raw": Raw amounts (e.g., 1000000000 for 1 token with 9 decimals)
	//   - "scaled": Human-readable amounts (e.g., 1.0)
	//   - "both": Include both raw and scaled amounts
	// Optional, default: "raw"
	UIAmountMode string `default:"raw"`

	// Chains is the list of blockchain networks to query.
	// If nil, queries all supported networks.
	// Optional, default: nil (all networks)
	Chains []Chain

	// OnLimitExceeded overrides the default rate limit behavior for this request.
	// If nil, uses the client's default behavior.
	// Optional, default: nil (use client default)
	OnLimitExceeded string `default:""`
}

// GetMultiTokenPrice retrieves the current price of multiple tokens in a single request.
//
// This method efficiently retrieves prices for multiple tokens in one API call, which is more
// efficient than making individual requests for each token. All tokens are filtered by the same
// liquidity threshold and return the same data structure as GetTokenPrice.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - addresses: List of token contract addresses to query (max 100 addresses recommended)
//   - opts: Configuration options (see MultiTokenPriceOptions for details):
//   - CheckLiquidity: Minimum liquidity threshold in USD. Tokens below this threshold
//     may not be included. Optional, default: 100
//   - IncludeLiquidity: Whether to include liquidity information in response.
//     Optional, default: true
//   - UIAmountMode: Token amount display mode.
//     Options: "raw" (e.g., 1000000000), "scaled" (e.g., 1.0), "both"
//     Optional, default: "raw"
//   - Chains: List of blockchain networks to query. If nil, queries all supported networks.
//     Optional, default: nil
//   - OnLimitExceeded: Override rate limit behavior. If nil, uses client default.
//     Optional, default: nil
//
// Returns:
//   - map[string]RespTokenPrice: Dictionary mapping token addresses to RespTokenPrice objects.
//     Each RespTokenPrice contains:
//   - Value: Current token price in USD
//   - PriceChange24h: 24-hour price change percentage
//   - Liquidity: Current liquidity in USD
//   - UpdateUnixTime: Last update timestamp (Unix seconds)
//   - UpdateHumanTime: Human-readable update time
//   - PriceInNative: Price in native token (e.g., SOL, ETH)
//   - error: Error if request fails, any address is invalid, or rate limit exceeded
//
// Raises:
//   - Context errors: If ctx is cancelled or times out
//   - API errors: If the Birdeye API returns an error
//   - Validation errors: If any address is invalid
//   - Rate limit errors: If rate limit exceeded and OnLimitExceeded is RateLimitRaise
//
// Example:
//
//	// Get prices for multiple tokens
//	addresses := []string{
//	    "So11111111111111111111111111111111111111112",  // SOL
//	    "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v", // USDC
//	    "Es9vMFrzaCERmJfrF4H2FYD4KCoNkY11McCe8BenwNYB",  // USDT
//	}
//	prices, err := client.GetMultiTokenPrice(ctx, addresses, MultiTokenPriceOptions{})
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	for address, priceData := range prices {
//	    fmt.Printf("%s: $%.2f\n", address[:8], priceData.Value)
//	}
//
//	// Get prices with liquidity filtering
//	checkLiq := int64(1000)
//	prices, err = client.GetMultiTokenPrice(ctx, addresses, MultiTokenPriceOptions{
//	    CheckLiquidity: &checkLiq,  // Only include tokens with > $1000 liquidity
//	    IncludeLiquidity: &[]bool{true}[0],
//	})
func (c *HTTPClient) GetMultiTokenPrice(ctx context.Context, addresses []string, opts *MultiTokenPriceOptions) (map[string]RespTokenPrice, error) {
	if opts == nil {
		opts = &MultiTokenPriceOptions{}
	}

	params, err := ApplyDefaultsAndBuildParams(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to apply defaults: %w", err)
	}

	// Add required parameters
	params["list_address"] = addresses

	// Handle special cases
	if opts.IncludeLiquidity {
		params["include_liquidity"] = "true"
	}

	result, err := c.request(ctx, EndpointDefiMultiPrice, requestOptions{
		method:          "GET",
		chains:          opts.Chains,
		onLimitExceeded: RateLimitBehavior(opts.OnLimitExceeded),
		paramsOrBody:    params,
	})
	if err != nil {
		return nil, err
	}

	// Parse result into map[string]RespTokenPrice
	var prices map[string]RespTokenPrice
	data, _ := json.Marshal(result)
	if err := json.Unmarshal(data, &prices); err != nil {
		return nil, err
	}

	return prices, nil
}

// ============================================================================
// API Methods - Token Transactions
// ============================================================================

// TokenTxsOptions holds options for GetTokenTxs.
type TokenTxsOptions struct {
	// Offset is the number of transactions to skip for pagination.
	// Optional, default: 0
	Offset int64 `default:"0"`

	// Limit is the maximum number of transactions to return (1-50).
	// Optional, default: 50, max: 50
	Limit int64 `default:"50"`

	// TxType specifies the type of transactions to retrieve.
	// Options:
	//   - "swap": Only swap transactions
	//   - "add": Only liquidity add transactions
	//   - "remove": Only liquidity remove transactions
	//   - "all": All transaction types
	// Optional, default: "swap"
	TxType string `default:"swap"`

	// SortType specifies the sort order by timestamp.
	// Options:
	//   - "desc": Newest first
	//   - "asc": Oldest first
	// Optional, default: "desc"
	SortType string `default:"desc"`

	// UIAmountMode specifies the token amount display mode.
	// Options:
	//   - "raw": Raw amounts (e.g., 1000000000 for 1 token with 9 decimals)
	//   - "scaled": Human-readable amounts (e.g., 1.0)
	//   - "both": Include both raw and scaled amounts
	// Optional, default: "raw"
	UIAmountMode string `default:"raw"`

	// Chains is the list of blockchain networks to query.
	// If nil, queries all supported networks.
	// Optional, default: nil (all networks)
	Chains []Chain

	// OnLimitExceeded overrides the default rate limit behavior for this request.
	// If nil, uses the client's default behavior.
	// Optional, default: nil (use client default)
	OnLimitExceeded string `default:""`
}

// GetTokenTxs retrieves transaction history for a specific token.
//
// This method retrieves transaction history for a token, including swaps, liquidity adds, and removes.
// Results can be filtered by transaction type and sorted by timestamp.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - address: Token contract address to query
//   - opts: Configuration options (see TokenTxsOptions for details):
//   - Offset: Number of transactions to skip for pagination. Optional, default: 0
//   - Limit: Maximum number of transactions to return (1-50). Optional, default: 50, max: 50
//   - TxType: Type of transactions to retrieve.
//     Options: "swap" (swaps only), "add" (liquidity adds), "remove" (liquidity removes), "all" (all types)
//     Optional, default: "swap"
//   - SortType: Sort order by timestamp.
//     Options: "desc" (newest first), "asc" (oldest first)
//     Optional, default: "desc"
//   - UIAmountMode: Token amount display mode.
//     Options: "raw" (e.g., 1000000000), "scaled" (e.g., 1.0), "both"
//     Optional, default: "raw"
//   - Chains: List of blockchain networks to query. Optional, default: nil
//   - OnLimitExceeded: Override rate limit behavior. Optional, default: nil
//
// Returns:
//   - *RespTokenTxs: Transaction response containing:
//   - Items: List of transaction records, each containing txHash, blockTime, type, amounts, etc.
//   - HasNext: Boolean indicating whether more transactions are available for pagination
//   - error: Error if request fails, address is invalid, or rate limit exceeded
//
// Raises:
//   - Context errors: If ctx is cancelled or times out
//   - API errors: If the Birdeye API returns an error
//   - Validation errors: If address is invalid
//   - Rate limit errors: If rate limit exceeded and OnLimitExceeded is RateLimitRaise
//
// Example:
//
//	// Get recent swap transactions for SOL
//	txs, err := client.GetTokenTxs(ctx, "So11111111111111111111111111111111111111112", TokenTxsOptions{
//	    Limit:  20,
//	    TxType: "swap",
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, tx := range txs.Items {
//	    fmt.Printf("Tx: %s, Amount: %.2f\n", tx.TxHash, tx.Quote.UIAmount)
//	}
//
//	// Get all transaction types
//	allTxs, err := client.GetTokenTxs(ctx, tokenAddress, TokenTxsOptions{
//	    TxType:   "all",
//	    SortType: "desc",
//	    Limit:    100,
//	})
func (c *HTTPClient) GetTokenTxs(ctx context.Context, address string, opts *TokenTxsOptions) (*RespTokenTxs, error) {
	if opts == nil {
		opts = &TokenTxsOptions{}
	}

	params, err := ApplyDefaultsAndBuildParams(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to apply defaults: %w", err)
	}

	// Add required parameters
	params["address"] = address

	result, err := c.request(ctx, EndpointDefiTxsToken, requestOptions{
		method:          "GET",
		chains:          opts.Chains,
		onLimitExceeded: RateLimitBehavior(opts.OnLimitExceeded),
		paramsOrBody:    params,
	})
	if err != nil {
		return nil, err
	}

	var txs RespTokenTxs
	data, _ := json.Marshal(result)
	if err := json.Unmarshal(data, &txs); err != nil {
		return nil, err
	}

	return &txs, nil
}

// ============================================================================
// API Methods - OHLCV Data
// ============================================================================

// TokenOHLCVOptions holds options for GetTokenOHLCV.
type TokenOHLCVOptions struct {
	// Currency specifies the currency for OHLCV data.
	// Options:
	//   - "usd": USD denomination
	//   - "native": Native token denomination (e.g., SOL, ETH)
	// Optional, default: "usd"
	Currency string `default:"usd"`

	// UIAmountMode specifies the token amount display mode.
	// Options:
	//   - "raw": Raw amounts (e.g., 1000000000 for 1 token with 9 decimals)
	//   - "scaled": Human-readable amounts (e.g., 1.0)
	//   - "both": Include both raw and scaled amounts
	// Optional, default: "raw"
	UIAmountMode string `default:"raw"`

	// Chains is the list of blockchain networks to query.
	// If nil, queries all supported networks.
	// Optional, default: nil (all networks)
	Chains []Chain

	// OnLimitExceeded overrides the default rate limit behavior for this request.
	// If nil, uses the client's default behavior.
	// Optional, default: nil (use client default)
	OnLimitExceeded string `default:""`
}

// GetTokenOHLCV retrieves OHLCV (Open, High, Low, Close, Volume) data for a token.
//
// This method retrieves historical OHLCV data for a token over a specified time range with
// configurable time intervals. Useful for charting, technical analysis, and price history visualization.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - address: Token contract address to query
//   - intervalType: Time interval for OHLCV data.
//     Options: "1m", "3m", "5m", "15m", "30m", "1H", "2H", "4H", "6H", "8H", "12H", "1D", "3D", "1W", "1M"
//   - timeFrom: Start time in Unix timestamp (seconds). Must be between 0 and 10000000000
//   - timeTo: End time in Unix timestamp (seconds). Must be between 0 and 10000000000
//   - opts: Configuration options (see TokenOHLCVOptions for details)
//
// Returns:
//   - *RespTokenOHLCVs: OHLCV response containing:
//   - Items: List of OHLCV data points
//   - Each item contains: open, high, low, close, volume, unixTime
//   - error: Error if request fails, address is invalid, or time range is invalid
//
// Raises:
//   - Context errors: If ctx is cancelled or times out
//   - API errors: If the Birdeye API returns an error
//   - Validation errors: If address is invalid or time range is invalid
//   - Rate limit errors: If rate limit exceeded and OnLimitExceeded is RateLimitRaise
//
// Example:
//
//	// Get 1-hour OHLCV data for the last 24 hours
//	endTime := time.Now().Unix()
//	startTime := endTime - 86400  // 24 hours ago
//
//	ohlcv, err := client.GetTokenOHLCV(ctx,
//	    "So11111111111111111111111111111111111111112",  // SOL
//	    "1H", startTime, endTime, TokenOHLCVOptions{})
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, candle := range ohlcv.Items {
//	    fmt.Printf("Time: %d, Close: $%.2f\n", candle.UnixTime, candle.C)
//	}
//
//	// Get daily data for a month
//	monthlyOHLCV, err := client.GetTokenOHLCV(ctx, tokenAddress, "1D", startTime, endTime, TokenOHLCVOptions{})
func (c *HTTPClient) GetTokenOHLCV(ctx context.Context, address string, intervalType string, timeFrom, timeTo int64, opts *TokenOHLCVOptions) (*RespTokenOHLCVs, error) {
	if opts == nil {
		opts = &TokenOHLCVOptions{}
	}
	params, err := ApplyDefaultsAndBuildParams(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to apply defaults: %w", err)
	}

	if timeFrom < 0 || timeFrom > 10000000000 {
		return nil, errors.New("time_from must be between 0 and 10000000000")
	}
	if timeTo < 0 || timeTo > 10000000000 {
		return nil, errors.New("time_to must be between 0 and 10000000000")
	}

	// Add required parameters
	params["address"] = address
	params["type"] = intervalType
	params["time_from"] = timeFrom
	params["time_to"] = timeTo

	result, err := c.request(ctx, EndpointDefiOHLCV, requestOptions{
		method:          "GET",
		chains:          opts.Chains,
		onLimitExceeded: RateLimitBehavior(opts.OnLimitExceeded),
		paramsOrBody:    params,
	})
	if err != nil {
		return nil, err
	}

	var ohlcv RespTokenOHLCVs
	data, _ := json.Marshal(result)
	if err := json.Unmarshal(data, &ohlcv); err != nil {
		return nil, err
	}

	return &ohlcv, nil
}

// ============================================================================
// API Methods - Token Metadata
// ============================================================================

// TokenMetadataOptions holds options for GetTokenMetadata and GetMultiTokenMetadata.
type TokenMetadataOptions struct {
	// Chains is the list of blockchain networks to query.
	// If nil, queries all supported networks.
	// Optional, default: nil (all networks)
	Chains []Chain

	// OnLimitExceeded overrides the default rate limit behavior for this request.
	// If nil, uses the client's default behavior.
	// Optional, default: nil (use client default)
	OnLimitExceeded string `default:""`
}

// GetTokenMetadata retrieves detailed metadata information for a token.
//
// This method retrieves comprehensive metadata including name, symbol, decimals,
// logo URI, and social media links (website, Twitter, Discord, etc.).
//
// Parameters:
//   - ctx: Context for cancellation and timeout
//   - address: Token contract address
//   - opts: Optional parameters (chains, ui_amount_mode)
//
// Returns:
//   - *RespTokenMetadata: Token metadata
//   - error: Any error encountered
//
// Example:
//
//	metadata, err := client.GetTokenMetadata(ctx, tokenAddress, TokenMetadataOptions{})
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("%s (%s) - %d decimals\n", metadata.Name, metadata.Symbol, metadata.Decimals)
func (c *HTTPClient) GetTokenMetadata(ctx context.Context, address string, opts *TokenMetadataOptions) (*RespTokenMetadata, error) {
	if opts == nil {
		opts = &TokenMetadataOptions{}
	}
	params := map[string]any{
		"address": address,
	}

	result, err := c.request(ctx, EndpointDefiV3TokenMetadataSingle, requestOptions{
		method:          "GET",
		chains:          opts.Chains,
		onLimitExceeded: RateLimitBehavior(opts.OnLimitExceeded),
		paramsOrBody:    params,
	})
	if err != nil {
		return nil, err
	}

	var metadata RespTokenMetadata
	data, _ := json.Marshal(result)
	if err := json.Unmarshal(data, &metadata); err != nil {
		return nil, err
	}

	return &metadata, nil
}

// GetMultiTokenMetadata retrieves detailed metadata for multiple tokens in a single request.
//
// This method retrieves comprehensive metadata information for multiple tokens in one API call,
// improving efficiency for batch operations. Returns metadata including name, symbol, decimals,
// logo URI, and social media links.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - addresses: List of token contract addresses to query
//   - opts: Configuration options (see TokenMetadataOptions for details)
//
// Returns:
//   - []RespTokenMetadata: List of token metadata, each containing:
//   - Address, Symbol, Name, Decimals, LogoURI, Website, Twitter, Discord, etc.
//   - error: Error if request fails or any address is invalid
//
// Raises:
//   - Context errors: If ctx is cancelled or times out
//   - API errors: If the Birdeye API returns an error
//   - Validation errors: If any address is invalid
//
// Example:
//
//	addresses := []string{
//	    "So11111111111111111111111111111111111111112",  // SOL
//	    "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v", // USDC
//	}
//	metadata, err := client.GetMultiTokenMetadata(ctx, addresses, TokenMetadataOptions{})
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, meta := range metadata {
//	    fmt.Printf("%s (%s)\n", meta.Name, meta.Symbol)
//	}
func (c *HTTPClient) GetMultiTokenMetadata(ctx context.Context, addresses []string, opts *TokenMetadataOptions) (RespMultiTokenMetadata, error) {
	if opts == nil {
		opts = &TokenMetadataOptions{}
	}
	params := map[string]any{
		"list_address": addresses,
	}

	result, err := c.request(ctx, EndpointDefiV3TokenMetadataMultiple, requestOptions{
		method:          "GET",
		chains:          opts.Chains,
		onLimitExceeded: RateLimitBehavior(opts.OnLimitExceeded),
		paramsOrBody:    params,
	})
	if err != nil {
		return nil, err
	}

	var metadata RespMultiTokenMetadata
	data, _ := json.Marshal(result)
	if err := json.Unmarshal(data, &metadata); err != nil {
		return nil, err
	}

	return metadata, nil
}

// ============================================================================
// API Methods - Token Market Data
// ============================================================================

// TokenMarketDataOptions holds options for GetTokenMarketData and GetMultiTokenMarketData.
type TokenMarketDataOptions struct {
	// UIAmountMode specifies the token amount display mode.
	// Options:
	//   - "raw": Raw amounts (e.g., 1000000000 for 1 token with 9 decimals)
	//   - "scaled": Human-readable amounts (e.g., 1.0)
	//   - "both": Include both raw and scaled amounts
	// Optional, default: "raw"
	UIAmountMode string `default:"raw"`

	// Chains is the list of blockchain networks to query.
	// If nil, queries all supported networks.
	// Optional, default: nil (all networks)
	Chains []Chain

	// OnLimitExceeded overrides the default rate limit behavior for this request.
	// If nil, uses the client's default behavior.
	// Optional, default: nil (use client default)
	OnLimitExceeded string `default:""`
}

// GetTokenMarketData retrieves comprehensive market data for a token.
//
// This method retrieves detailed market data including price, volume, liquidity, and other key
// market indicators for a token.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - address: Token contract address to query
//   - opts: Configuration options (see TokenMarketDataOptions for details)
//
// Returns:
//   - *RespTokenMarketData: Token market data response containing:
//   - Price: Current token price
//   - Volume24h: 24-hour trading volume
//   - Liquidity: Total liquidity
//   - MarketCap: Market capitalization
//   - PriceChange24h: 24-hour price change percentage
//   - PriceChange7d: 7-day price change percentage
//   - error: Error if request fails or address is invalid
//
// Raises:
//   - Context errors: If ctx is cancelled or times out
//   - API errors: If the Birdeye API returns an error
//   - Validation errors: If address is invalid
//
// Example:
//
//	marketData, err := client.GetTokenMarketData(ctx, tokenAddress, TokenMarketDataOptions{})
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Market Cap: $%.2f, Liquidity: $%.2f\n", marketData.MarketCap, marketData.Liquidity)
func (c *HTTPClient) GetTokenMarketData(ctx context.Context, address string, opts *TokenMarketDataOptions) (*RespTokenMarketData, error) {
	if opts == nil {
		opts = &TokenMarketDataOptions{}
	}
	params, err := ApplyDefaultsAndBuildParams(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to apply defaults: %w", err)
	}

	// Add required parameters
	params["address"] = address

	result, err := c.request(ctx, EndpointDefiV3TokenMarketData, requestOptions{
		method:          "GET",
		chains:          opts.Chains,
		onLimitExceeded: RateLimitBehavior(opts.OnLimitExceeded),
		paramsOrBody:    params,
	})
	if err != nil {
		return nil, err
	}

	var marketData RespTokenMarketData
	data, _ := json.Marshal(result)
	if err := json.Unmarshal(data, &marketData); err != nil {
		return nil, err
	}

	return &marketData, nil
}

// GetMultiTokenMarketData retrieves comprehensive market data for multiple tokens.
//
// This method retrieves detailed market data for multiple tokens in a single request,
// improving efficiency for batch operations.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - addresses: List of token contract addresses to query
//   - opts: Configuration options (see TokenMarketDataOptions for details)
//
// Returns:
//   - map[string]RespTokenMarketData: Dictionary mapping addresses to market data
//     Each item contains: price, volume24h, liquidity, marketCap, priceChange24h
//   - error: Error if request fails or any address is invalid
//
// Raises:
//   - Context errors: If ctx is cancelled or times out
//   - API errors: If the Birdeye API returns an error
//
// Example:
//
//	addresses := []string{"token1", "token2"}
//	data, err := client.GetMultiTokenMarketData(ctx, addresses, TokenMarketDataOptions{})
//	for addr, mktData := range data {
//	    fmt.Printf("%s: MC $%.2f\n", addr[:8], mktData.MarketCap)
//	}
func (c *HTTPClient) GetMultiTokenMarketData(ctx context.Context, addresses []string, opts *TokenMarketDataOptions) (map[string]RespTokenMarketData, error) {
	if opts == nil {
		opts = &TokenMarketDataOptions{}
	}
	params, err := ApplyDefaultsAndBuildParams(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to apply defaults: %w", err)
	}

	// Add required parameters
	params["list_address"] = addresses

	result, err := c.request(ctx, EndpointDefiV3TokenMarketDataMultiple, requestOptions{
		method:          "GET",
		chains:          opts.Chains,
		onLimitExceeded: RateLimitBehavior(opts.OnLimitExceeded),
		paramsOrBody:    params,
	})
	if err != nil {
		return nil, err
	}

	var marketData map[string]RespTokenMarketData
	data, _ := json.Marshal(result)
	if err := json.Unmarshal(data, &marketData); err != nil {
		return nil, err
	}

	return marketData, nil
}

// ============================================================================
// API Methods - Token Trade Data
// ============================================================================

// TokenTradeDataOptions holds options for GetTokenTradeData and GetMultiTokenTradeData.
type TokenTradeDataOptions struct {
	// Frames specifies the time periods for trade data aggregation.
	// Options: "1m", "5m", "30m", "1h", "2h", "4h", "8h", "24h"
	// If nil or empty, returns data for all available time frames.
	// Optional, default: nil (all frames)
	Frames []string

	// UIAmountMode specifies the token amount display mode.
	// Options:
	//   - "raw": Raw amounts (e.g., 1000000000 for 1 token with 9 decimals)
	//   - "scaled": Human-readable amounts (e.g., 1.0)
	//   - "both": Include both raw and scaled amounts
	// Optional, default: "raw"
	UIAmountMode string `default:"raw"`

	// Chains is the list of blockchain networks to query.
	// If nil, queries all supported networks.
	// Optional, default: nil (all networks)
	Chains []Chain

	// OnLimitExceeded overrides the default rate limit behavior for this request.
	// If nil, uses the client's default behavior.
	// Optional, default: nil (use client default)
	OnLimitExceeded string `default:""`
}

// GetTokenTradeData retrieves comprehensive trading statistics for a token.
//
// This method provides detailed trading metrics including volume, trade counts,
// unique wallets, buy/sell ratios across multiple time periods (1m, 5m, 30m, 1h, 2h, 4h, 8h, 24h).
//
// Parameters:
//   - ctx: Context for cancellation and timeout
//   - address: Token contract address
//   - opts: Optional parameters (ui_amount_mode, chains)
//
// Returns:
//   - *RespTokenTradeData: Comprehensive trading statistics
//   - error: Any error encountered
//
// Example:
//
//	tradeData, err := client.GetTokenTradeData(ctx, tokenAddress, TokenTradeDataOptions{})
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("24h Volume: $%.2f, Holders: %d\n", tradeData.Volume24hUSD, tradeData.Holder)
func (c *HTTPClient) GetTokenTradeData(ctx context.Context, address string, opts *TokenTradeDataOptions) (*RespTokenTradeData, error) {
	if opts == nil {
		opts = &TokenTradeDataOptions{}
	}
	params, err := ApplyDefaultsAndBuildParams(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to apply defaults: %w", err)
	}

	// Add required parameters
	params["address"] = address

	if len(opts.Frames) > 0 {
		params["frames"] = opts.Frames
	}

	result, err := c.request(ctx, EndpointDefiV3TokenTradeDataSingle, requestOptions{
		method:          "GET",
		chains:          opts.Chains,
		onLimitExceeded: RateLimitBehavior(opts.OnLimitExceeded),
		paramsOrBody:    params,
	})
	if err != nil {
		return nil, err
	}

	var tradeData RespTokenTradeData
	data, _ := json.Marshal(result)
	if err := json.Unmarshal(data, &tradeData); err != nil {
		return nil, err
	}

	return &tradeData, nil
}

// GetMultiTokenTradeData retrieves trading data for multiple tokens.
//
// This method retrieves trading data for multiple tokens in a single request,
// improving efficiency for batch operations.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - addresses: List of token contract addresses to query
//   - opts: Configuration options (see TokenTradeDataOptions for details)
//
// Returns:
//   - map[string]RespTokenTradeData: Dictionary mapping addresses to trade data
//     Each item contains: recentTrades, volume24h, tradeCount, avgTradeSize
//   - error: Error if request fails or any address is invalid
//
// Example:
//
//	addresses := []string{"token1", "token2"}
//	tradeData, err := client.GetMultiTokenTradeData(ctx, addresses, TokenTradeDataOptions{})
//	for addr, data := range tradeData {
//	    fmt.Printf("%s: Volume $%.2f\n", addr[:8], data.Volume24hUSD)
//	}
func (c *HTTPClient) GetMultiTokenTradeData(ctx context.Context, addresses []string, opts *TokenTradeDataOptions) (map[string]RespTokenTradeData, error) {
	if opts == nil {
		opts = &TokenTradeDataOptions{}
	}
	params, err := ApplyDefaultsAndBuildParams(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to apply defaults: %w", err)
	}

	// Add required parameters
	params["list_address"] = addresses

	if len(opts.Frames) > 0 {
		params["frames"] = opts.Frames
	}

	result, err := c.request(ctx, EndpointDefiV3TokenTradeDataMultiple, requestOptions{
		method:          "GET",
		chains:          opts.Chains,
		onLimitExceeded: RateLimitBehavior(opts.OnLimitExceeded),
		paramsOrBody:    params,
	})
	if err != nil {
		return nil, err
	}

	var tradeData map[string]RespTokenTradeData
	data, _ := json.Marshal(result)
	if err := json.Unmarshal(data, &tradeData); err != nil {
		return nil, err
	}

	return tradeData, nil
}

// ============================================================================
// API Methods - Token Security
// ============================================================================

// TokenSecurityOptions holds options for GetTokenSecurity.
//
// Note: Security endpoint is only available for Solana chain.
type TokenSecurityOptions struct {
	// Chains is the list of blockchain networks to query.
	// If nil, queries all supported networks.
	// Note: Currently only Solana is supported for security data.
	// Optional, default: nil (all networks)
	Chains []Chain

	// OnLimitExceeded overrides the default rate limit behavior for this request.
	// If nil, uses the client's default behavior.
	// Optional, default: nil (use client default)
	OnLimitExceeded string `default:""`
}

// GetTokenSecurity retrieves security information for a token.
//
// This method retrieves security-related information about a token including risk assessment,
// security scores, and safety indicators. Essential for due diligence before trading or investing.
//
// Note: This endpoint is only available for Solana chain.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - address: Token contract address to query
//   - opts: Configuration options (see TokenSecurityOptions for details)
//
// Returns:
//   - *RespTokenSecurity: Security response containing:
//   - RiskScore: Overall risk score (0-100)
//   - SecurityScore: Security score (0-100)
//   - IsHoneypot: Whether token is a honeypot
//   - IsRugpull: Whether token is a rugpull
//   - IsVerified: Whether token is verified
//   - Warnings: List of security warnings
//   - error: Error if request fails or address is invalid
//
// Raises:
//   - Context errors: If ctx is cancelled or times out
//   - API errors: If the Birdeye API returns an error
//   - Validation errors: If address is invalid
//
// Example:
//
//	security, err := client.GetTokenSecurity(ctx, tokenAddress, TokenSecurityOptions{})
//	if err != nil {
//	    log.Fatal(err)
//	}
//	if security.MutableMetadata != nil && *security.MutableMetadata {
//	    fmt.Println("Warning: Token has mutable metadata")
//	}
//	fmt.Printf("Top 10 holders own: %.2f%%\n", *security.Top10HolderPercent)
func (c *HTTPClient) GetTokenSecurity(ctx context.Context, address string, opts *TokenSecurityOptions) (*RespTokenSecurity, error) {
	if opts == nil {
		opts = &TokenSecurityOptions{}
	}
	params := map[string]any{
		"address": address,
	}

	result, err := c.request(ctx, EndpointDefiTokenSecurity, requestOptions{
		method:          "GET",
		chains:          opts.Chains,
		onLimitExceeded: RateLimitBehavior(opts.OnLimitExceeded),
		paramsOrBody:    params,
	})
	if err != nil {
		return nil, err
	}

	var security RespTokenSecurity
	data, _ := json.Marshal(result)
	if err := json.Unmarshal(data, &security); err != nil {
		return nil, err
	}

	return &security, nil
}

// ============================================================================
// API Methods - Token Holders
// ============================================================================

// TokenHoldersOptions holds options for GetTokenHolders.
type TokenHoldersOptions struct {
	// Offset is the number of holders to skip for pagination.
	// Optional, default: 0
	Offset int64 `default:"0"`

	// Limit is the maximum number of holders to return (1-100).
	// Optional, default: 100, max: 100
	Limit int64 `default:"100"`

	// UIAmountMode specifies the token amount display mode.
	// Options:
	//   - "raw": Raw amounts (e.g., 1000000000 for 1 token with 9 decimals)
	//   - "scaled": Human-readable amounts (e.g., 1.0)
	//   - "both": Include both raw and scaled amounts
	// Optional, default: "scaled"
	UIAmountMode string `default:"scaled"`

	// Chains is the list of blockchain networks to query.
	// If nil, queries all supported networks.
	// Optional, default: nil (all networks)
	Chains []Chain

	// OnLimitExceeded overrides the default rate limit behavior for this request.
	// If nil, uses the client's default behavior.
	// Optional, default: nil (use client default)
	OnLimitExceeded string `default:""`
}

// GetTokenHolders retrieves token holders information.
//
// This method retrieves information about token holders including holder addresses, balances,
// and percentage of total supply.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - address: Token contract address to query
//   - opts: Configuration options (see TokenHoldersOptions for details)
//
// Returns:
//   - []RespTokenHoldersItem: Token holders response containing:
//   - TotalHolders: Total number of holders
//   - Holders: List of holder information
//   - Each item contains: address, balance, percentage
//   - error: Error if request fails or address is invalid
//
// Raises:
//   - Context errors: If ctx is cancelled or times out
//   - API errors: If the Birdeye API returns an error
//   - Validation errors: If address is invalid or offset+limit > 10000 or limit > 100
//   - Rate limit errors: If rate limit exceeded and OnLimitExceeded is RateLimitRaise
//
// Example:
//
//	holders, err := client.GetTokenHolders(ctx, tokenAddress, TokenHoldersOptions{
//	    Limit: 100,
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, holder := range holders {
//	    fmt.Printf("Owner: %s, Amount: %.2f\n", holder.Owner, holder.UIAmount)
//	}
func (c *HTTPClient) GetTokenHolders(ctx context.Context, address string, opts *TokenHoldersOptions) (RespMultiTokenHolders, error) {
	if opts == nil {
		opts = &TokenHoldersOptions{}
	}
	params, err := ApplyDefaultsAndBuildParams(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to apply defaults: %w", err)
	}

	if opts.Offset < 0 || opts.Offset > 10000 {
		return nil, errors.New("offset must be between 0 and 10000")
	}
	if opts.Limit < 1 || opts.Limit > 100 {
		return nil, errors.New("limit must be between 1 and 100")
	}
	if opts.Offset+opts.Limit > 10000 {
		return nil, errors.New("offset + limit must not exceed 10000")
	}

	// Add required parameters
	params["address"] = address

	result, err := c.request(ctx, EndpointDefiV3TokenHolder, requestOptions{
		method:          "GET",
		chains:          opts.Chains,
		onLimitExceeded: RateLimitBehavior(opts.OnLimitExceeded),
		paramsOrBody:    params,
	})
	if err != nil {
		return nil, err
	}

	var holders map[string]RespMultiTokenHolders
	data, _ := json.Marshal(result)
	if err := json.Unmarshal(data, &holders); err != nil {
		return nil, err
	}

	return holders["items"], nil
}

// ============================================================================
// API Methods - Wallet Portfolio
// ============================================================================

// WalletPortfolioOptions holds options for GetWalletPortfolio.
//
// Note: This endpoint is only available for Solana chain.
type WalletPortfolioOptions struct {
	// UIAmountMode specifies the token amount display mode.
	// Options:
	//   - "raw": Raw amounts (e.g., 1000000000 for 1 token with 9 decimals)
	//   - "scaled": Human-readable amounts (e.g., 1.0)
	//   - "both": Include both raw and scaled amounts
	// Optional, default: "raw"
	UIAmountMode string `default:"raw"`

	// Chains is the list of blockchain networks to query.
	// If nil, queries all supported networks.
	// Note: Currently only Solana is supported.
	// Optional, default: nil (all networks)
	Chains []Chain

	// OnLimitExceeded overrides the default rate limit behavior for this request.
	// If nil, uses the client's default behavior.
	// Optional, default: nil (use client default)
	OnLimitExceeded string `default:""`
}

// GetWalletPortfolio retrieves the complete portfolio overview for a wallet.
//
// This method retrieves comprehensive portfolio information including total value, token holdings,
// and performance metrics for a wallet across multiple blockchain networks.
//
// Note: This endpoint is only available for Solana chain.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - wallet: Wallet address to query
//   - opts: Configuration options (see WalletPortfolioOptions for details)
//
// Returns:
//   - []RespWalletPortfolioItem: Portfolio response containing:
//   - total_value: Total portfolio value in USD
//   - tokens: List of token holdings
//   - performance: Performance metrics and statistics
//   - error: Error if request fails or wallet address is invalid
//
// Raises:
//   - Context errors: If ctx is cancelled or times out
//   - API errors: If the Birdeye API returns an error
//   - Validation errors: If wallet address is invalid
//   - Rate limit errors: If rate limit exceeded and OnLimitExceeded is RateLimitRaise
//
// Example:
//
//	// Get wallet portfolio
//	portfolio, err := client.GetWalletPortfolio(ctx, "wallet-address", WalletPortfolioOptions{})
//	if err != nil {
//	    log.Fatal(err)
//	}
//	var totalValue float64
//	for _, token := range portfolio {
//	    totalValue += token.ValueUSD
//	    fmt.Printf("%s: $%.2f\n", token.Symbol, token.ValueUSD)
//	}
//	fmt.Printf("Total Value: $%.2f\n", totalValue)
//
//	// Get portfolio for specific chains
//	portfolio, err = client.GetWalletPortfolio(ctx, "wallet-address", WalletPortfolioOptions{
//	    Chains: []Chain{ChainSolana, ChainEthereum},
//	})
func (c *HTTPClient) GetWalletPortfolio(ctx context.Context, wallet string, opts *WalletPortfolioOptions) (RespWalletPortfolio, error) {
	if opts == nil {
		opts = &WalletPortfolioOptions{}
	}
	params, err := ApplyDefaultsAndBuildParams(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to apply defaults: %w", err)
	}

	// Add required parameters
	params["wallet"] = wallet

	result, err := c.request(ctx, EndpointV1WalletTokenList, requestOptions{
		method:          "GET",
		chains:          opts.Chains,
		onLimitExceeded: RateLimitBehavior(opts.OnLimitExceeded),
		paramsOrBody:    params,
		respJustItems:   true,
	})
	if err != nil {
		return nil, err
	}

	var portfolio map[string]RespWalletPortfolio
	data, _ := json.Marshal(result)
	if err := json.Unmarshal(data, &portfolio); err != nil {
		return nil, err
	}

	return portfolio["items"], nil
}

// ============================================================================
// API Methods - Wallet Transactions
// ============================================================================

// WalletTxsOptions holds options for GetWalletTxs.
//
// Note: This endpoint is only available for Solana chain.
type WalletTxsOptions struct {
	// Limit is the maximum number of transactions to return (1-100).
	// Optional, default: 50, max: 100
	Limit int64 `default:"50"`

	// Before is a cursor for pagination, typically a transaction hash.
	// Get transactions before this cursor.
	// Optional, default: nil (start from latest)
	Before string `default:""`

	// UIAmountMode specifies the token amount display mode.
	// Options:
	//   - "raw": Raw amounts (e.g., 1000000000 for 1 token with 9 decimals)
	//   - "scaled": Human-readable amounts (e.g., 1.0)
	//   - "both": Include both raw and scaled amounts
	// Optional, default: "raw"
	UIAmountMode string `default:"raw"`

	// Chains is the list of blockchain networks to query.
	// If nil, queries all supported networks.
	// Note: Currently only Solana is supported.
	// Optional, default: nil (all networks)
	Chains []Chain

	// OnLimitExceeded overrides the default rate limit behavior for this request.
	// If nil, uses the client's default behavior.
	// Optional, default: nil (use client default)
	OnLimitExceeded string `default:""`
}

// GetWalletTxs retrieves transaction history for a wallet.
//
// This method retrieves the transaction history for a specific wallet address including
// all transactions and their details.
//
// Note: This endpoint is only available for Solana chain.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - wallet: Wallet address to query
//   - opts: Configuration options (see WalletTxsOptions for details)
//
// Returns:
//   - map[Chain][]RespWalletTx: Wallet transactions response containing:
//   - totalTxs: Total number of transactions
//   - txs: List of transactions
//   - Each item contains: txHash, timestamp, type, amount, token
//   - error: Error if request fails or wallet address is invalid
//
// Raises:
//   - Context errors: If ctx is cancelled or times out
//   - API errors: If the Birdeye API returns an error
//   - Validation errors: If wallet address is invalid or limit out of range (1-100)
//   - Rate limit errors: If rate limit exceeded and OnLimitExceeded is RateLimitRaise
//
// Example:
//
//	txs, err := client.GetWalletTxs(ctx, "wallet-address", WalletTxsOptions{
//	    Limit: 50,
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for chain, chainTxs := range txs {
//	    fmt.Printf("Chain: %s, Txs: %d\n", chain, len(chainTxs))
//	    for _, tx := range chainTxs {
//	        fmt.Printf("  TX: %s\n", tx.TxHash)
//	    }
//	}
func (c *HTTPClient) GetWalletTxs(ctx context.Context, wallet string, opts *WalletTxsOptions) (map[Chain][]RespWalletTx, error) {
	if opts == nil {
		opts = &WalletTxsOptions{}
	}
	params, err := ApplyDefaultsAndBuildParams(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to apply defaults: %w", err)
	}

	if opts.Limit < 1 || opts.Limit > 100 {
		return nil, errors.New("limit must be between 1 and 100")
	}

	// Add required parameters
	params["wallet"] = wallet

	result, err := c.request(ctx, EndpointV1WalletTxList, requestOptions{
		method:          "GET",
		chains:          opts.Chains,
		onLimitExceeded: RateLimitBehavior(opts.OnLimitExceeded),
		paramsOrBody:    params,
	})
	if err != nil {
		return nil, err
	}

	var txs map[Chain][]RespWalletTx
	data, _ := json.Marshal(result)
	if err := json.Unmarshal(data, &txs); err != nil {
		return nil, err
	}

	return txs, nil
}

// ============================================================================
// API Methods - Wallet Net Worth
// ============================================================================

// WalletNetWorthOptions holds options for GetWalletNetWorth.
//
// Note: This endpoint is only available for Solana chain.
type WalletNetWorthOptions struct {
	// FilterValue filters tokens by minimum value in USD. Only tokens with value >= this will be returned. Default: 0 (no filter)
	FilterValue float64 `default:"0"`
	// SortBy specifies the field to sort by. Options: "value", "amount". Default: "value"
	SortBy string `default:"value"`
	// SortType specifies the sort order. Options: "desc", "asc". Default: "desc"
	SortType string `default:"desc"`
	// Limit is the maximum number of tokens to return. Default: 100
	Limit int64 `default:"100"`
	// Offset is the number of tokens to skip for pagination. Default: 0
	Offset int64 `default:"0"`
	// Chains is the list of blockchain networks to query. Default: nil
	Chains []Chain
	// OnLimitExceeded overrides the default rate limit behavior. Default: "" (use client default)
	OnLimitExceeded string `default:""`
}

// GetWalletNetWorth retrieves the current net worth for a wallet.
//
// This method retrieves the current net worth for a specific wallet address including
// total value and asset breakdown across all holdings.
//
// Note: Solana only
//
// Parameters:
//   - ctx: Context for cancellation and timeout
//   - wallet: Wallet address to query
//   - opts: Optional parameters including:
//   - Limit: Maximum number of tokens to return (default: 100)
//   - Offset: Number of tokens to skip for pagination (default: 0)
//   - Chains: List of blockchain networks to query
//   - OnLimitExceeded: Rate limit behavior override
//
// Returns:
//   - *RespWalletNetWorth: Net worth data containing:
//   - WalletAddress: The queried wallet address
//   - TotalValue: Total wallet value in USD
//   - Items: List of token holdings with values
//   - Pagination: Pagination information
//   - error: Any error encountered during the request
//
// Example:
//
//	// Get wallet net worth
//	netWorth, err := client.GetWalletNetWorth(ctx, walletAddress, WalletNetWorthOptions{})
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Total Net Worth: $%s\n", netWorth.TotalValue)
//	fmt.Printf("Holdings: %d tokens\n", len(netWorth.Items))
func (c *HTTPClient) GetWalletNetWorth(ctx context.Context, wallet string, opts *WalletNetWorthOptions) (*RespWalletNetWorth, error) {
	if opts == nil {
		opts = &WalletNetWorthOptions{}
	}
	params, err := ApplyDefaultsAndBuildParams(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to apply defaults: %w", err)
	}

	if opts.Limit < 1 || opts.Limit > 100 {
		return nil, errors.New("limit must be between 1 and 100")
	}
	if opts.Offset < 0 || opts.Offset > 10000 {
		return nil, errors.New("offset must be between 0 and 10000")
	}

	// Add required parameters
	params["wallet"] = wallet

	if opts.FilterValue > 0 {
		params["filter_value"] = opts.FilterValue
	}

	result, err := c.request(ctx, EndpointV2WalletCurrentNetWorth, requestOptions{
		method:          "GET",
		chains:          opts.Chains,
		onLimitExceeded: RateLimitBehavior(opts.OnLimitExceeded),
		paramsOrBody:    params,
	})
	if err != nil {
		return nil, err
	}

	var netWorth RespWalletNetWorth
	data, _ := json.Marshal(result)
	if err := json.Unmarshal(data, &netWorth); err != nil {
		return nil, err
	}

	return &netWorth, nil
}

// ============================================================================
// API Methods - Search
// ============================================================================

// SearchOptions holds options for Search.
type SearchOptions struct {
	// Chain specifies a single blockchain network to search. If nil, searches all networks. Default: nil
	Chain *Chain
	// Target specifies what to search for. Options: "all", "token", "pair". Default: "all"
	Target string `default:"all"`
	// SearchMode specifies the search matching mode. Options: "exact", "prefix", "fuzzy". Default: "exact"
	SearchMode string `default:"exact"`
	// SearchBy specifies which field to search by. Options: "address", "symbol", "name". Default: "symbol"
	SearchBy string `default:"symbol"`
	// SortBy specifies the field to sort results by. Options: "liquidity", "volume", "market_cap". Default: "liquidity"
	SortBy string `default:"liquidity"`
	// SortType specifies the sort order. Options: "desc", "asc". Default: "desc"
	SortType string `default:"desc"`
	// VerifyToken filters to only show verified tokens if true. Default: false (show all)
	VerifyToken bool `default:"false"`
	// Markets filters by specific markets/exchanges. Comma-separated string. Default: "" (no filter)
	Markets string `default:""`
	// Offset is the number of results to skip for pagination. Default: 0
	Offset int64 `default:"0"`
	// Limit is the maximum number of results to return. Default: 10
	Limit int64 `default:"10"`
	// UIAmountMode specifies the token amount display mode. Options: "raw", "scaled", "both". Default: "raw"
	UIAmountMode string `default:"raw"`
	// OnLimitExceeded overrides the default rate limit behavior. Default: "" (use client default)
	OnLimitExceeded string `default:""`
}

// Search searches for tokens, markets, and other entities across blockchain networks.
//
// This method provides comprehensive search functionality across multiple blockchain networks,
// allowing users to find tokens, markets, and other entities by name, symbol, or address.
//
// Parameters:
//   - ctx: Context for cancellation and timeout
//   - keyword: Search keyword (token name, symbol, or address)
//   - opts: Optional parameters including:
//   - Offset: Number of results to skip for pagination (default: 0)
//   - Limit: Maximum number of results to return, 1-100 (default: 10)
//   - VerifyToken: Whether to verify token authenticity
//   - Target: Search target type - "token" or "market"
//   - Chains: List of blockchain networks to search
//   - OnLimitExceeded: Rate limit behavior override
//
// Returns:
//   - []RespSearchItem: List of search results with token/market information
//   - error: Any error encountered during the request
//
// Example:
//
//	// Search for tokens
//	results, err := client.Search(ctx, "SOL", SearchOptions{
//	    Target: "token",
//	    Limit:  10,
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, item := range results {
//	    for _, token := range item.Result {
//	        fmt.Printf("%s (%s): $%.2f\n", token.Name, token.Symbol, token.Price)
//	    }
//	}
func (c *HTTPClient) Search(ctx context.Context, keyword string, opts *SearchOptions) (RespSearchItems, error) {
	if opts == nil {
		opts = &SearchOptions{}
	}
	params, err := ApplyDefaultsAndBuildParams(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to apply defaults: %w", err)
	}

	// Add required parameters
	params["keyword"] = keyword

	if opts.VerifyToken {
		params["verify_token"] = "true"
	}
	if opts.Markets != "" {
		params["markets"] = opts.Markets
	}

	var chains []Chain
	if opts.Chain != nil {
		chains = []Chain{*opts.Chain}
	}

	result, err := c.request(ctx, EndpointDefiV3Search, requestOptions{
		method:          "GET",
		chains:          chains,
		onLimitExceeded: RateLimitBehavior(opts.OnLimitExceeded),
		paramsOrBody:    params,
		respJustItems:   true,
	})
	if err != nil {
		return nil, err
	}

	var searchResults map[string]RespSearchItems
	data, _ := json.Marshal(result)
	if err := json.Unmarshal(data, &searchResults); err != nil {
		return nil, err
	}

	return searchResults["items"], nil
}

// ============================================================================
// API Methods - Pair Transactions
// ============================================================================

// PairTxsOptions holds options for GetPairTxs.
type PairTxsOptions struct {
	// Offset is the number of transactions to skip for pagination.
	// Optional, default: 0
	Offset int64 `default:"0"`

	// Limit is the maximum number of transactions to return (1-50).
	// Optional, default: 50, max: 50
	Limit int64 `default:"50"`

	// TxType specifies the type of transactions to retrieve.
	// Options:
	//   - "swap": Only swap transactions
	//   - "add": Only liquidity add transactions
	//   - "remove": Only liquidity remove transactions
	//   - "all": All transaction types
	// Optional, default: "swap"
	TxType string `default:"swap"`

	// SortType specifies the sort order by timestamp.
	// Options:
	//   - "desc": Newest first
	//   - "asc": Oldest first
	// Optional, default: "desc"
	SortType string `default:"desc"`

	// UIAmountMode specifies the token amount display mode.
	// Options:
	//   - "raw": Raw amounts (e.g., 1000000000 for 1 token with 9 decimals)
	//   - "scaled": Human-readable amounts (e.g., 1.0)
	//   - "both": Include both raw and scaled amounts
	// Optional, default: "raw"
	UIAmountMode string `default:"raw"`

	// Chains is the list of blockchain networks to query.
	// If nil, queries all supported networks.
	// Optional, default: nil (all networks)
	Chains []Chain

	// OnLimitExceeded overrides the default rate limit behavior for this request.
	// If nil, uses the client's default behavior.
	// Optional, default: nil (use client default)
	OnLimitExceeded string `default:""`
}

// GetPairTxs retrieves transaction history for a trading pair.
//
// This method retrieves detailed transaction information for a specific trading pair,
// including swaps, liquidity adds, and removes.
//
// Parameters:
//   - ctx: Context for cancellation and timeout
//   - address: Trading pair address
//   - opts: Optional parameters
//   - Offset: Number of transactions to skip for pagination (default: 0)
//   - Limit: Maximum number of transactions, 1-1000 (default: 50)
//   - TxType: Type - "swap", "add", "remove", or "all" (default: "swap")
//   - SortType: Sort order - "desc" or "asc" (default: "desc")
//   - Chains: List of blockchain networks
//
// Returns:
//   - *RespPairTxs: Pair transaction data with pagination
//   - error: Any error encountered
//
// Example:
//
//	pairTxs, err := client.GetPairTxs(ctx, pairAddress, PairTxsOptions{Limit: 20, TxType: "swap"})
func (c *HTTPClient) GetPairTxs(ctx context.Context, address string, opts *PairTxsOptions) (*RespPairTxs, error) {
	if opts == nil {
		opts = &PairTxsOptions{}
	}
	params, err := ApplyDefaultsAndBuildParams(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to apply defaults: %w", err)
	}

	// Add required parameters
	params["address"] = address

	result, err := c.request(ctx, EndpointDefiTxsPair, requestOptions{
		method:          "GET",
		chains:          opts.Chains,
		onLimitExceeded: RateLimitBehavior(opts.OnLimitExceeded),
		paramsOrBody:    params,
	})
	if err != nil {
		return nil, err
	}

	var txs RespPairTxs
	data, _ := json.Marshal(result)
	if err := json.Unmarshal(data, &txs); err != nil {
		return nil, err
	}

	return &txs, nil
}

// ============================================================================
// API Methods - Transactions By Time
// ============================================================================

// TokenTxsByTimeOptions holds options for GetTokenTxsByTime.
type TokenTxsByTimeOptions struct {
	// AfterTime specifies the start time filter in Unix timestamp (seconds).
	// Get transactions after this time.
	// Optional, default: nil (no lower time limit)
	AfterTime int64 `default:"0"`

	// BeforeTime specifies the end time filter in Unix timestamp (seconds).
	// Get transactions before this time.
	// Optional, default: nil (no upper time limit)
	BeforeTime int64 `default:"0"`

	// Offset is the number of transactions to skip for pagination.
	// Optional, default: 0
	Offset int64 `default:"0"`

	// Limit is the maximum number of transactions to return (1-100).
	// Optional, default: 100, max: 100
	Limit int64 `default:"100"`

	// TxType specifies the type of transactions to retrieve.
	// Options:
	//   - "swap": Only swap transactions
	//   - "add": Only liquidity add transactions
	//   - "remove": Only liquidity remove transactions
	//   - "all": All transaction types
	// Optional, default: "swap"
	TxType string `default:"swap"`

	// SortType specifies the sort order by timestamp.
	// Options:
	//   - "desc": Newest first
	//   - "asc": Oldest first
	// Optional, default: "desc"
	SortType string `default:"desc"`

	// UIAmountMode specifies the token amount display mode.
	// Options:
	//   - "raw": Raw amounts (e.g., 1000000000 for 1 token with 9 decimals)
	//   - "scaled": Human-readable amounts (e.g., 1.0)
	//   - "both": Include both raw and scaled amounts
	// Optional, default: "raw"
	UIAmountMode string `default:"raw"`

	// Chains is the list of blockchain networks to query.
	// If nil, queries all supported networks.
	// Optional, default: nil (all networks)
	Chains []Chain

	// OnLimitExceeded overrides the default rate limit behavior for this request.
	// If nil, uses the client's default behavior.
	// Optional, default: nil (use client default)
	OnLimitExceeded string `default:""`
}

// GetTokenTxsByTime retrieves token transactions within a specific time range.
//
// This method retrieves transaction history for a token within a specified time window,
// allowing for precise time-based filtering of transactions.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - address: Token contract address to query
//   - opts: Configuration options (see TokenTxsByTimeOptions for details)
//
// Returns:
//   - *RespTokenTxsByTime: Transaction response containing:
//   - items: List of transaction records
//   - hasNext: Whether more transactions are available
//   - error: Error if request fails, address is invalid, or time range is invalid
//
// Raises:
//   - Context errors: If ctx is cancelled or times out
//   - API errors: If the Birdeye API returns an error
//   - Validation errors: If address is invalid or time range is invalid
//   - Rate limit errors: If rate limit exceeded and OnLimitExceeded is RateLimitRaise
//
// Example:
//
//	// Get transactions from the last hour
//	currentTime := time.Now().Unix()
//	oneHourAgo := currentTime - 3600
//
//	txs, err := client.GetTokenTxsByTime(ctx,
//	    "So11111111111111111111111111111111111111112",  // SOL
//	    TokenTxsByTimeOptions{
//	        AfterTime:  &oneHourAgo,
//	        BeforeTime: &currentTime,
//	        TxType:     "swap",
//	        Limit:      100,
//	    })
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, tx := range txs.Items {
//	    fmt.Printf("Tx: %s, Time: %d\n", tx.TxHash, tx.BlockUnixTime)
//	}
//
//	// Get all transaction types in a specific time window
//	allTxs, err := client.GetTokenTxsByTime(ctx, tokenAddress, TokenTxsByTimeOptions{
//	    AfterTime:  &[]int64{1640995200}[0],  // Jan 1, 2022
//	    BeforeTime: &[]int64{1672531200}[0],  // Jan 1, 2023
//	    TxType:     "all",
//	})
func (c *HTTPClient) GetTokenTxsByTime(ctx context.Context, address string, opts *TokenTxsByTimeOptions) (*RespTokenTxsByTime, error) {
	if opts == nil {
		opts = &TokenTxsByTimeOptions{}
	}
	params, err := ApplyDefaultsAndBuildParams(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to apply defaults: %w", err)
	}

	// Add required parameters
	params["address"] = address

	if opts.AfterTime > 0 {
		params["after_time"] = opts.AfterTime
	}
	if opts.BeforeTime > 0 {
		params["before_time"] = opts.BeforeTime
	}

	result, err := c.request(ctx, EndpointDefiTxsTokenSeekByTime, requestOptions{
		method:          "GET",
		chains:          opts.Chains,
		onLimitExceeded: RateLimitBehavior(opts.OnLimitExceeded),
		paramsOrBody:    params,
	})
	if err != nil {
		return nil, err
	}

	var txs RespTokenTxsByTime
	data, _ := json.Marshal(result)
	if err := json.Unmarshal(data, &txs); err != nil {
		return nil, err
	}

	return &txs, nil
}

// PairTxsByTimeOptions holds options for GetPairTxsByTime.
type PairTxsByTimeOptions struct {
	// AfterTime filters transactions after this Unix timestamp (seconds). Default: 0 (no filter)
	AfterTime int64 `default:"0"`
	// BeforeTime filters transactions before this Unix timestamp (seconds). Default: 0 (no filter)
	BeforeTime int64 `default:"0"`
	// Offset is the number of transactions to skip for pagination. Default: 0
	Offset int64 `default:"0"`
	// Limit is the maximum number of transactions to return (1-100). Default: 100, max: 100
	Limit int64 `default:"100"`
	// TxType specifies the transaction type. Options: "swap", "add", "remove", "all". Default: "swap"
	TxType string `default:"swap"`
	// SortType specifies the sort order. Options: "desc", "asc". Default: "desc"
	SortType string `default:"desc"`
	// UIAmountMode specifies the token amount display mode. Options: "raw", "scaled", "both". Default: "raw"
	UIAmountMode string `default:"raw"`
	// Chains is the list of blockchain networks to query. Default: nil
	Chains []Chain
	// OnLimitExceeded overrides the default rate limit behavior. Default: "" (use client default)
	OnLimitExceeded string `default:""`
}

// GetPairTxsByTime retrieves trading pair transactions within a specific time range.
//
// This method retrieves transaction history for a trading pair within a specified time window,
// allowing for precise time-based filtering of pair transactions.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - address: Trading pair address to query
//   - opts: Configuration options (see PairTxsByTimeOptions for details)
//
// Returns:
//   - *RespPairTxsByTime: Transaction response containing:
//   - items: List of transaction records
//   - hasNext: Whether more transactions are available
//   - error: Error if request fails, address is invalid, or time range is invalid
//
// Raises:
//   - Context errors: If ctx is cancelled or times out
//   - API errors: If the Birdeye API returns an error
//   - Validation errors: If address is invalid or time range is invalid
//   - Rate limit errors: If rate limit exceeded and OnLimitExceeded is RateLimitRaise
//
// Example:
//
//	// Get pair transactions from the last 24 hours
//	currentTime := time.Now().Unix()
//	oneDayAgo := currentTime - 86400
//
//	pairTxs, err := client.GetPairTxsByTime(ctx, "pair-address", PairTxsByTimeOptions{
//	    AfterTime:  &oneDayAgo,
//	    BeforeTime: &currentTime,
//	    TxType:     "swap",
//	    Limit:      100,
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, tx := range pairTxs.Items {
//	    fmt.Printf("Tx: %s, Type: %s\n", tx.TxHash, tx.TxType)
//	}
//
//	// Get liquidity operations in a specific time window
//	liquidityTxs, err := client.GetPairTxsByTime(ctx, "pair-address", PairTxsByTimeOptions{
//	    AfterTime:  &[]int64{1640995200}[0],  // Jan 1, 2022
//	    BeforeTime: &[]int64{1672531200}[0],  // Jan 1, 2023
//	    TxType:     "all",
//	})
func (c *HTTPClient) GetPairTxsByTime(ctx context.Context, address string, opts *PairTxsByTimeOptions) (*RespPairTxsByTime, error) {
	if opts == nil {
		opts = &PairTxsByTimeOptions{}
	}
	params, err := ApplyDefaultsAndBuildParams(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to apply defaults: %w", err)
	}

	// Add required parameters
	params["address"] = address

	if opts.AfterTime > 0 {
		params["after_time"] = opts.AfterTime
	}
	if opts.BeforeTime > 0 {
		params["before_time"] = opts.BeforeTime
	}

	result, err := c.request(ctx, EndpointDefiTxsPairSeekByTime, requestOptions{
		method:          "GET",
		chains:          opts.Chains,
		onLimitExceeded: RateLimitBehavior(opts.OnLimitExceeded),
		paramsOrBody:    params,
	})
	if err != nil {
		return nil, err
	}

	var txs RespPairTxsByTime
	data, _ := json.Marshal(result)
	if err := json.Unmarshal(data, &txs); err != nil {
		return nil, err
	}

	return &txs, nil
}

// ============================================================================
// API Methods - Token Transactions V3
// ============================================================================

// TokenTxsV3Options holds options for GetTokenTxsV3.
type TokenTxsV3Options struct {
	// Offset is the number of transactions to skip for pagination. Default: 0
	Offset int64 `default:"0"`
	// Limit is the maximum number of transactions to return (1-100). Default: 100, max: 100
	Limit int64 `default:"100"`
	// SortBy specifies the field to sort by. Options: "block_unix_time", "block_number". Default: "block_unix_time"
	SortBy string `default:"block_unix_time"`
	// SortType specifies the sort order. Options: "desc", "asc". Default: "desc"
	SortType string `default:"desc"`
	// TxType specifies the transaction type. Options: "swap", "add", "remove", "all". Default: "swap"
	TxType string `default:"swap"`
	// Source filters by DEX source (e.g., "raydium", "orca"). Default: "" (no filter)
	Source string `default:""`
	// Owner filters by owner/wallet address. Default: "" (no filter)
	Owner string `default:""`
	// PoolID filters by pool/pair ID. Default: "" (no filter)
	PoolID string `default:""`
	// BeforeTime filters transactions before this Unix timestamp (seconds). Default: 0 (no filter)
	BeforeTime int64 `default:"0"`
	// AfterTime filters transactions after this Unix timestamp (seconds). Default: 0 (no filter)
	AfterTime int64 `default:"0"`
	// BeforeBlockNumber filters transactions before this block number. Default: 0 (no filter)
	BeforeBlockNumber int64 `default:"0"`
	// AfterBlockNumber filters transactions after this block number. Default: 0 (no filter)
	AfterBlockNumber int64 `default:"0"`
	// UIAmountMode specifies the token amount display mode. Options: "raw", "scaled", "both". Default: "raw"
	UIAmountMode string `default:"raw"`
	// Chains is the list of blockchain networks to query. Default: nil
	Chains []Chain
	// OnLimitExceeded overrides the default rate limit behavior. Default: "" (use client default)
	OnLimitExceeded string `default:""`
}

// GetTokenTxsV3 retrieves token transactions using the V3 API with enhanced filtering.
//
// This method retrieves transaction history for a token using the V3 API, which provides
// enhanced filtering options and improved data structure.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - address: Token contract address to query
//   - opts: Configuration options (see TokenTxsV3Options for details)
//
// Returns:
//   - *RespTokenTxsV3: Transaction response containing:
//   - items: List of transaction records
//   - hasNext: Whether more transactions are available
//   - error: Error if request fails, address is invalid, or time range is invalid
//
// Raises:
//   - Context errors: If ctx is cancelled or times out
//   - API errors: If the Birdeye API returns an error
//   - Validation errors: If address is invalid, time range is invalid, or offset+limit > 10000 or limit > 100
//   - Rate limit errors: If rate limit exceeded and OnLimitExceeded is RateLimitRaise
//
// Example:
//
//	// Get recent swap transactions for SOL
//	currentTime := time.Now().Unix()
//	oneHourAgo := currentTime - 3600
//
//	txs, err := client.GetTokenTxsV3(ctx,
//	    "So11111111111111111111111111111111111111112",  // SOL
//	    TokenTxsV3Options{
//	        TxType:     "swap",
//	        AfterTime:  &oneHourAgo,
//	        Limit:      100,
//	    })
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, tx := range txs.Items {
//	    fmt.Printf("Tx: %s, Volume: %.2f\n", tx.TxHash, tx.Volume)
//	}
//
//	// Get liquidity operations
//	liquidityTxs, err := client.GetTokenTxsV3(ctx, tokenAddress, TokenTxsV3Options{
//	    TxType:   "add_liquidity",
//	    SortType: "desc",
//	    Limit:    50,
//	})
func (c *HTTPClient) GetTokenTxsV3(ctx context.Context, address string, opts *TokenTxsV3Options) (*RespTokenTxsV3, error) {
	if opts == nil {
		opts = &TokenTxsV3Options{}
	}
	params, err := ApplyDefaultsAndBuildParams(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to apply defaults: %w", err)
	}

	if opts.Offset+opts.Limit > 10000 {
		return nil, errors.New("offset + limit must be <= 10000")
	}
	if opts.Limit > 100 {
		return nil, errors.New("limit must be <= 100")
	}

	// Add required parameters
	params["address"] = address

	if opts.Source != "" {
		params["source"] = opts.Source
	}
	if opts.Owner != "" {
		params["owner"] = opts.Owner
	}
	if opts.PoolID != "" {
		params["pool_id"] = opts.PoolID
	}
	if opts.BeforeTime > 0 {
		params["before_time"] = opts.BeforeTime
	}
	if opts.AfterTime > 0 {
		params["after_time"] = opts.AfterTime
	}
	if opts.BeforeBlockNumber > 0 {
		params["before_block_number"] = opts.BeforeBlockNumber
	}
	if opts.AfterBlockNumber > 0 {
		params["after_block_number"] = opts.AfterBlockNumber
	}

	result, err := c.request(ctx, EndpointDefiV3TokenTxs, requestOptions{
		method:          "GET",
		chains:          opts.Chains,
		onLimitExceeded: RateLimitBehavior(opts.OnLimitExceeded),
		paramsOrBody:    params,
	})
	if err != nil {
		return nil, err
	}

	var txs RespTokenTxsV3
	data, _ := json.Marshal(result)
	if err := json.Unmarshal(data, &txs); err != nil {
		return nil, err
	}

	return &txs, nil
}

// ============================================================================
// API Methods - Pair OHLCV
// ============================================================================

// PairOHLCVOptions holds options for GetPairOHLCV.
type PairOHLCVOptions struct {
	// UIAmountMode specifies the token amount display mode.
	// Options:
	//   - "raw": Raw amounts (e.g., 1000000000 for 1 token with 9 decimals)
	//   - "scaled": Human-readable amounts (e.g., 1.0)
	//   - "both": Include both raw and scaled amounts
	// Optional, default: "raw"
	UIAmountMode string `default:"raw"`

	// Chains is the list of blockchain networks to query.
	// If nil, queries all supported networks.
	// Optional, default: nil (all networks)
	Chains []Chain

	// OnLimitExceeded overrides the default rate limit behavior for this request.
	// If nil, uses the client's default behavior.
	// Optional, default: nil (use client default)
	OnLimitExceeded string `default:""`
}

// GetPairOHLCV retrieves OHLCV (Open, High, Low, Close, Volume) data for a trading pair.
//
// This method retrieves historical OHLCV data for a trading pair over a specified time range
// with configurable time intervals.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - address: Trading pair address to query
//   - intervalType: Time interval for OHLCV data.
//     Options: "1m", "3m", "5m", "15m", "30m", "1H", "2H", "4H", "6H", "8H", "12H", "1D", "3D", "1W", "1M"
//   - timeFrom: Start time in Unix timestamp (seconds). Must be between 0 and 10000000000
//   - timeTo: End time in Unix timestamp (seconds). Must be between 0 and 10000000000
//   - opts: Configuration options (see PairOHLCVOptions for details)
//
// Returns:
//   - []RespPairOHLCVItem: OHLCV response containing list of OHLCV data points
//     Each item contains: open, high, low, close, volume, timestamp
//   - error: Error if request fails, address is invalid, or time range is invalid
//
// Raises:
//   - Context errors: If ctx is cancelled or times out
//   - API errors: If the Birdeye API returns an error
//   - Validation errors: If address is invalid or time range is invalid
//   - Rate limit errors: If rate limit exceeded and OnLimitExceeded is RateLimitRaise
//
// Example:
//
//	// Get 1-hour OHLCV data for a trading pair over the last 24 hours
//	endTime := time.Now().Unix()
//	startTime := endTime - 86400  // 24 hours ago
//
//	pairOHLCV, err := client.GetPairOHLCV(ctx, "pair-address", "1H", startTime, endTime, PairOHLCVOptions{})
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, candle := range pairOHLCV {
//	    fmt.Printf("Time: %d, Close: $%.2f\n", candle.UnixTime, candle.C)
//	}
//
//	// Get daily data for a month
//	monthlyOHLCV, err := client.GetPairOHLCV(ctx, "pair-address", "1D", startTime, endTime, PairOHLCVOptions{})
func (c *HTTPClient) GetPairOHLCV(ctx context.Context, address, intervalType string, timeFrom, timeTo int64, opts *PairOHLCVOptions) ([]RespPairOHLCVItem, error) {
	if opts == nil {
		opts = &PairOHLCVOptions{}
	}
	params, err := ApplyDefaultsAndBuildParams(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to apply defaults: %w", err)
	}

	if timeFrom < 0 || timeFrom > 10000000000 {
		return nil, errors.New("time_from must be between 0 and 10000000000")
	}
	if timeTo < 0 || timeTo > 10000000000 {
		return nil, errors.New("time_to must be between 0 and 10000000000")
	}

	// Add required parameters
	params["address"] = address
	params["type"] = intervalType
	params["time_from"] = timeFrom
	params["time_to"] = timeTo

	result, err := c.request(ctx, EndpointDefiOHLCVPair, requestOptions{
		method:          "GET",
		chains:          opts.Chains,
		onLimitExceeded: RateLimitBehavior(opts.OnLimitExceeded),
		paramsOrBody:    params,
		respJustItems:   true,
	})
	if err != nil {
		return nil, err
	}

	var ohlcv map[string][]RespPairOHLCVItem
	data, _ := json.Marshal(result)
	if err := json.Unmarshal(data, &ohlcv); err != nil {
		return nil, err
	}

	return ohlcv["items"], nil
}

// ============================================================================
// API Methods - Pair Overview
// ============================================================================

// PairOverviewOptions holds options for GetPairOverview and GetMultiPairOverview.
//
// Note: This endpoint is only available for Solana chain.
type PairOverviewOptions struct {
	// UIAmountMode specifies the token amount display mode.
	// Options:
	//   - "raw": Raw amounts (e.g., 1000000000 for 1 token with 9 decimals)
	//   - "scaled": Human-readable amounts (e.g., 1.0)
	//   - "both": Include both raw and scaled amounts
	// Optional, default: "raw"
	UIAmountMode string `default:"raw"`

	// Chains is the list of blockchain networks to query.
	// If nil, queries all supported networks.
	// Note: Currently only Solana is supported.
	// Optional, default: nil (all networks)
	Chains []Chain

	// OnLimitExceeded overrides the default rate limit behavior for this request.
	// If nil, uses the client's default behavior.
	// Optional, default: nil (use client default)
	OnLimitExceeded string `default:""`
}

// GetPairOverview retrieves comprehensive overview data for a trading pair.
//
// This method retrieves detailed information about a trading pair including liquidity, volume,
// price changes, and market statistics.
//
// Note: This endpoint is only available for Solana chain.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - address: Trading pair address to query
//   - opts: Configuration options (see PairOverviewOptions for details)
//
// Returns:
//   - *RespPairOverview: Pair overview response containing:
//   - baseToken: Base token information
//   - quoteToken: Quote token information
//   - price: Current pair price
//   - liquidity: Total liquidity
//   - volume24h: 24-hour trading volume
//   - priceChange24h: 24-hour price change percentage
//   - error: Error if request fails or address is invalid
//
// Raises:
//   - Context errors: If ctx is cancelled or times out
//   - API errors: If the Birdeye API returns an error
//   - Validation errors: If address is invalid
//   - Rate limit errors: If rate limit exceeded and OnLimitExceeded is RateLimitRaise
//
// Example:
//
//	// Get SOL/USDC pair overview
//	pairOverview, err := client.GetPairOverview(ctx, "pair-address", PairOverviewOptions{})
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Price: $%.2f\n", pairOverview.Price)
//	fmt.Printf("Liquidity: $%.2f\n", pairOverview.Liquidity)
//	fmt.Printf("24h Volume: $%.2f\n", pairOverview.Volume24h)
func (c *HTTPClient) GetPairOverview(ctx context.Context, address string, opts *PairOverviewOptions) (*RespPairOverview, error) {
	if opts == nil {
		opts = &PairOverviewOptions{}
	}
	params, err := ApplyDefaultsAndBuildParams(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to apply defaults: %w", err)
	}

	// Add required parameters
	params["address"] = address

	result, err := c.request(ctx, EndpointDefiV3PairOverviewSingle, requestOptions{
		method:          "GET",
		chains:          opts.Chains,
		onLimitExceeded: RateLimitBehavior(opts.OnLimitExceeded),
		paramsOrBody:    params,
	})
	if err != nil {
		return nil, err
	}

	var overview RespPairOverview
	data, _ := json.Marshal(result)
	if err := json.Unmarshal(data, &overview); err != nil {
		return nil, err
	}

	return &overview, nil
}

// GetPairsOverview retrieves comprehensive overview data for multiple trading pairs.
//
// This method retrieves detailed information about multiple trading pairs in a single request,
// improving efficiency for batch operations.
//
// Note: This endpoint is only available for Solana chain.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - addresses: List of trading pair addresses to query
//   - opts: Configuration options (see PairOverviewOptions for details)
//
// Returns:
//   - map[string]RespPairOverview: Pairs overview response containing:
//   - items: List of pair overview data
//   - Each item contains: baseToken, quoteToken, price, liquidity, volume24h
//   - error: Error if request fails or any address is invalid
//
// Raises:
//   - Context errors: If ctx is cancelled or times out
//   - API errors: If the Birdeye API returns an error
//   - Validation errors: If any address is invalid
//   - Rate limit errors: If rate limit exceeded and OnLimitExceeded is RateLimitRaise
//
// Example:
//
//	// Get overview for multiple pairs
//	pairsOverview, err := client.GetPairsOverview(ctx, []string{
//	    "pair-address-1",
//	    "pair-address-2",
//	    "pair-address-3",
//	}, PairOverviewOptions{})
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for addr, pair := range pairsOverview {
//	    fmt.Printf("Pair %s: $%.2f\n", addr[:8], pair.Price)
//	}
func (c *HTTPClient) GetPairsOverview(ctx context.Context, addresses []string, opts *PairOverviewOptions) (map[string]RespPairOverview, error) {
	if opts == nil {
		opts = &PairOverviewOptions{}
	}
	params, err := ApplyDefaultsAndBuildParams(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to apply defaults: %w", err)
	}

	// Add required parameters
	params["list_address"] = addresses

	result, err := c.request(ctx, EndpointDefiV3PairOverviewMultiple, requestOptions{
		method:          "GET",
		chains:          opts.Chains,
		onLimitExceeded: RateLimitBehavior(opts.OnLimitExceeded),
		paramsOrBody:    params,
	})
	if err != nil {
		return nil, err
	}

	var overviews map[string]RespPairOverview
	data, _ := json.Marshal(result)
	if err := json.Unmarshal(data, &overviews); err != nil {
		return nil, err
	}

	return overviews, nil
}

// ============================================================================
// API Methods - Token List
// ============================================================================

// TokenListV3Options holds options for GetTokenListV3.
type TokenListV3Options struct {
	// SortBy specifies the field to sort by. Options: "liquidity", "market_cap", "fdv", "volume_24h". Default: "liquidity"
	SortBy string `default:"liquidity"`
	// SortType specifies the sort order. Options: "desc", "asc". Default: "desc"
	SortType string `default:"desc"`
	// MinLiquidity filters tokens with liquidity >= this value in USD. Default: 0 (no filter)
	MinLiquidity float64 `default:"0"`
	// MaxLiquidity filters tokens with liquidity <= this value in USD. Default: 0 (no filter)
	MaxLiquidity float64 `default:"0"`
	// MinMarketCap filters tokens with market cap >= this value in USD. Default: 0 (no filter)
	MinMarketCap float64 `default:"0"`
	// MaxMarketCap filters tokens with market cap <= this value in USD. Default: 0 (no filter)
	MaxMarketCap float64 `default:"0"`
	// MinFDV filters tokens with fully diluted valuation >= this value in USD. Default: 0 (no filter)
	MinFDV float64 `default:"0"`
	// MaxFDV filters tokens with fully diluted valuation <= this value in USD. Default: 0 (no filter)
	MaxFDV float64 `default:"0"`
	// Offset is the number of tokens to skip for pagination. Default: 0
	Offset int64 `default:"0"`
	// Limit is the maximum number of tokens to return (1-100). Default: 100, max: 100
	Limit int64 `default:"100"`
	// UIAmountMode specifies the token amount display mode. Options: "raw", "scaled", "both". Default: "raw"
	UIAmountMode string `default:"raw"`
	// Chains is the list of blockchain networks to query. Default: nil
	Chains []Chain
	// OnLimitExceeded overrides the default rate limit behavior. Default: "" (use client default)
	OnLimitExceeded string `default:""`
}

// GetTokenListV3 retrieves token list using V3 API with advanced filtering.
//
// This method retrieves a list of tokens with advanced filtering options using the V3 API endpoint,
// providing more comprehensive data.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - opts: Configuration options (see TokenListV3Options for details)
//
// Returns:
//   - *RespTokenListV3: Token list response containing:
//   - items: List of token information
//   - hasNext: Whether more tokens are available
//   - error: Error if request fails or validation fails
//
// Raises:
//   - Context errors: If ctx is cancelled or times out
//   - API errors: If the Birdeye API returns an error
//   - Validation errors: If limit out of range (1-100) or offset+limit > 10000
//   - Rate limit errors: If rate limit exceeded and OnLimitExceeded is RateLimitRaise
//
// Example:
//
//	tokenList, err := client.GetTokenListV3(ctx, TokenListV3Options{
//	    SortBy:   "liquidity",
//	    SortType: "desc",
//	    Limit:    50,
//	    MinLiquidity: &[]float64{10000}[0],  // Min $10k liquidity
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, token := range tokenList.Items {
//	    fmt.Printf("%s: $%.2f\n", token.Symbol, token.Price)
//	}
func (c *HTTPClient) GetTokenListV3(ctx context.Context, opts *TokenListV3Options) (*RespTokenListV3, error) {
	if opts == nil {
		opts = &TokenListV3Options{}
	}
	params, err := ApplyDefaultsAndBuildParams(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to apply defaults: %w", err)
	}

	if opts.Limit < 1 || opts.Limit > 100 {
		return nil, errors.New("limit must be between 1 and 100")
	}
	if opts.Offset+opts.Limit > 10000 {
		return nil, errors.New("offset + limit must not exceed 10000")
	}

	// No additional parameters needed - all handled by ApplyDefaultsAndBuildParams
	if opts.MaxLiquidity > 0 {
		params["max_liquidity"] = opts.MaxLiquidity
	}
	if opts.MinMarketCap > 0 {
		params["min_market_cap"] = opts.MinMarketCap
	}
	if opts.MaxMarketCap > 0 {
		params["max_market_cap"] = opts.MaxMarketCap
	}
	if opts.MinFDV > 0 {
		params["min_fdv"] = opts.MinFDV
	}
	if opts.MaxFDV > 0 {
		params["max_fdv"] = opts.MaxFDV
	}

	result, err := c.request(ctx, EndpointDefiV3TokenList, requestOptions{
		method:          "GET",
		chains:          opts.Chains,
		onLimitExceeded: RateLimitBehavior(opts.OnLimitExceeded),
		paramsOrBody:    params,
	})
	if err != nil {
		return nil, err
	}

	var tokenList RespTokenListV3
	data, _ := json.Marshal(result)
	if err := json.Unmarshal(data, &tokenList); err != nil {
		return nil, err
	}

	return &tokenList, nil
}

// ============================================================================
// API Methods - Token Overview
// ============================================================================

// TokenOverviewOptions holds options for GetTokenOverview.
type TokenOverviewOptions struct {
	// Frames specifies the time periods for statistics. Options: "1m", "5m", "30m", "1h", "2h", "4h", "8h", "24h". Default: nil (all)
	Frames []string
	// UIAmountMode specifies the token amount display mode. Options: "raw", "scaled", "both". Default: "scaled"
	UIAmountMode string `default:"raw"`
	// Chains is the list of blockchain networks to query. Default: nil
	Chains []Chain
	// OnLimitExceeded overrides the default rate limit behavior. Default: "" (use client default)
	OnLimitExceeded string `default:""`
}

// GetTokenOverview retrieves comprehensive overview information for a token.
//
// This method retrieves detailed overview information about a token including basic metadata,
// market data, and key statistics.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - address: Token contract address to query
//   - opts: Configuration options (see TokenOverviewOptions for details)
//
// Returns:
//   - *RespTokenOverview: Token overview response containing:
//   - address: Token contract address
//   - symbol: Token symbol
//   - name: Token name
//   - decimals: Token decimals
//   - price: Current token price
//   - marketCap: Market capitalization
//   - volume24h: 24-hour trading volume
//   - priceChange24h: 24-hour price change percentage
//   - error: Error if request fails or address is invalid
//
// Raises:
//   - Context errors: If ctx is cancelled or times out
//   - API errors: If the Birdeye API returns an error
//   - Validation errors: If address is invalid
//   - Rate limit errors: If rate limit exceeded
//
// Example:
//
//	// Get SOL token overview
//	solOverview, err := client.GetTokenOverview(ctx,
//	    "So11111111111111111111111111111111111111112",  // SOL
//	    TokenOverviewOptions{})
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Symbol: %s\n", solOverview.Symbol)
//	fmt.Printf("Name: %s\n", solOverview.Name)
//	fmt.Printf("Price: $%.2f\n", solOverview.Price)
//	fmt.Printf("Market Cap: $%.2f\n", solOverview.MarketCap)
func (c *HTTPClient) GetTokenOverview(ctx context.Context, address string, opts *TokenOverviewOptions) (*RespTokenOverview, error) {
	if opts == nil {
		opts = &TokenOverviewOptions{}
	}
	params, err := ApplyDefaultsAndBuildParams(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to apply defaults: %w", err)
	}

	// Add required parameters
	params["address"] = address

	if len(opts.Frames) > 0 {
		params["frames"] = opts.Frames
	}

	result, err := c.request(ctx, EndpointDefiTokenOverview, requestOptions{
		method:          "GET",
		chains:          opts.Chains,
		onLimitExceeded: RateLimitBehavior(opts.OnLimitExceeded),
		paramsOrBody:    params,
	})
	if err != nil {
		return nil, err
	}

	var overview RespTokenOverview
	data, _ := json.Marshal(result)
	if err := json.Unmarshal(data, &overview); err != nil {
		return nil, err
	}

	return &overview, nil
}

// ============================================================================
// API Methods - Token Creation and Trending
// ============================================================================

// TokenCreationInfoOptions holds options for GetTokenCreationInfo.
//
// Note: This endpoint is only available for Solana chain.
type TokenCreationInfoOptions struct {
	// Chains is the list of blockchain networks to query.
	// Note: Currently only Solana is supported.
	// Optional, default: nil (all networks)
	Chains []Chain

	// OnLimitExceeded overrides the default rate limit behavior for this request.
	// Optional, default: nil (use client default)
	OnLimitExceeded string `default:""`
}

// GetTokenCreationInfo retrieves token creation information.
//
// This method retrieves information about when and how a token was created, including creation
// details and initial parameters.
//
// Note: This endpoint is only available for Solana chain.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - address: Token contract address to query
//   - opts: Configuration options (see TokenCreationInfoOptions for details)
//
// Returns:
//   - *RespTokenCreationInfo: Creation info response containing:
//   - createdAt: Creation timestamp
//   - creator: Creator wallet address
//   - initialSupply: Initial token supply
//   - decimals: Token decimals
//   - symbol: Token symbol
//   - name: Token name
//   - error: Error if request fails or address is invalid
//
// Raises:
//   - Context errors: If ctx is cancelled or times out
//   - API errors: If the Birdeye API returns an error
//   - Validation errors: If address is invalid
//   - Rate limit errors: If rate limit exceeded
//
// Example:
//
//	creationInfo, err := client.GetTokenCreationInfo(ctx, tokenAddress, TokenCreationInfoOptions{})
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Created at: %d\n", creationInfo.CreatedAt)
//	fmt.Printf("Creator: %s\n", creationInfo.Creator)
func (c *HTTPClient) GetTokenCreationInfo(ctx context.Context, address string, opts *TokenCreationInfoOptions) (*RespTokenCreationInfo, error) {
	if opts == nil {
		opts = &TokenCreationInfoOptions{}
	}
	params := map[string]any{
		"address": address,
	}

	result, err := c.request(ctx, EndpointDefiTokenCreationInfo, requestOptions{
		method:          "GET",
		chains:          opts.Chains,
		onLimitExceeded: RateLimitBehavior(opts.OnLimitExceeded),
		paramsOrBody:    params,
	})
	if err != nil {
		return nil, err
	}

	var info RespTokenCreationInfo
	data, _ := json.Marshal(result)
	if err := json.Unmarshal(data, &info); err != nil {
		return nil, err
	}

	return &info, nil
}

// TrendingListOptions holds options for GetTokenTrendingList.
type TrendingListOptions struct {
	// SortBy specifies the field to sort by.
	// Options: "rank", "volume", "volume_change_percent", "trade", "trade_change_percent",
	// "unique_wallet_24h", "unique_wallet_24h_change_percent"
	// Optional, default: "rank"
	SortBy string `default:"liquidity"`

	// SortType specifies the sort order.
	// Options: "asc", "desc"
	// Optional, default: "asc"
	SortType string `default:"desc"`

	// Offset is the number of items to skip for pagination.
	// Optional, default: 0
	Offset int64 `default:"0"`

	// Limit is the maximum number of tokens to return (1-20).
	// Optional, default: 20, max: 20
	Limit int64 `default:"20"`

	// UIAmountMode specifies the token amount display mode.
	// Options: "raw", "scaled", "both"
	// Optional, default: "raw"
	UIAmountMode string `default:"raw"`

	// Chains is the list of blockchain networks to query.
	// Optional, default: nil (all networks)
	Chains []Chain

	// OnLimitExceeded overrides the default rate limit behavior.
	// Optional, default: nil (use client default)
	OnLimitExceeded string `default:""`
}

// GetTokenTrendingList retrieves trending tokens list.
//
// This method retrieves a list of trending tokens based on various metrics like volume, trades,
// and unique wallet activity.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - opts: Configuration options (see TrendingListOptions for details)
//
// Returns:
//   - *RespTokenTrendingList: Trending list response containing:
//   - items: List of trending tokens
//   - hasNext: Whether more tokens are available
//   - error: Error if request fails
//
// Raises:
//   - Context errors: If ctx is cancelled or times out
//   - API errors: If the Birdeye API returns an error
//   - Validation errors: If limit out of range (1-20)
//   - Rate limit errors: If rate limit exceeded
//
// Example:
//
//	trending, err := client.GetTokenTrendingList(ctx, TrendingListOptions{
//	    SortBy:   "unique_wallet_24h",
//	    SortType: "desc",
//	    Limit:    10,
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, token := range trending.Items {
//	    fmt.Printf("%s: Rank %d\n", token.Symbol, token.Rank)
//	}
func (c *HTTPClient) GetTokenTrendingList(ctx context.Context, opts *TrendingListOptions) (*RespTokenTrendingList, error) {
	if opts == nil {
		opts = &TrendingListOptions{}
	}
	params, err := ApplyDefaultsAndBuildParams(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to apply defaults: %w", err)
	}

	if opts.Limit < 1 || opts.Limit > 20 {
		return nil, errors.New("limit must be between 1 and 20")
	}

	// No additional parameters needed - all handled by ApplyDefaultsAndBuildParams

	result, err := c.request(ctx, EndpointDefiTokenTrending, requestOptions{
		method:          "GET",
		chains:          opts.Chains,
		onLimitExceeded: RateLimitBehavior(opts.OnLimitExceeded),
		paramsOrBody:    params,
	})
	if err != nil {
		return nil, err
	}

	var trending RespTokenTrendingList
	data, _ := json.Marshal(result)
	if err := json.Unmarshal(data, &trending); err != nil {
		return nil, err
	}

	return &trending, nil
}

// NewListingOptions holds options for GetNewListing.
type NewListingOptions struct {
	// TimeTo filters listings before this Unix timestamp (seconds). Default: 0 (current time)
	TimeTo int64 `default:"0"`
	// Limit is the maximum number of listings to return (1-20). Default: 20, max: 20
	Limit int64 `default:"20"`
	// MemePlatformEnabled includes meme platform tokens if true. Default: false
	MemePlatformEnabled bool `default:"false"`
	// Chains is the list of blockchain networks to query. Default: nil
	Chains []Chain
	// OnLimitExceeded overrides the default rate limit behavior. Default: "" (use client default)
	OnLimitExceeded string `default:""`
}

// GetNewListing retrieves newly listed tokens.
//
// This method retrieves a list of newly listed tokens across supported networks,
// useful for discovering new opportunities.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - opts: Configuration options (see NewListingOptions for details)
//
// Returns:
//   - []RespNewTokenListingItem: New listing response containing:
//   - items: List of newly listed tokens
//   - hasNext: Whether more tokens are available
//   - error: Error if request fails
//
// Raises:
//   - Context errors: If ctx is cancelled or times out
//   - API errors: If the Birdeye API returns an error
//   - Validation errors: If limit out of range (1-20)
//   - Rate limit errors: If rate limit exceeded
//
// Example:
//
//	newListings, err := client.GetNewListing(ctx, NewListingOptions{
//	    Limit: 20,
//	    MemePlatformEnabled: false,
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, token := range newListings {
//	    fmt.Printf("New: %s (%s)\n", token.Symbol, token.Name)
//	}
func (c *HTTPClient) GetNewListing(ctx context.Context, opts *NewListingOptions) (RespTokenNewListing, error) {
	if opts == nil {
		opts = &NewListingOptions{}
	}
	params, err := ApplyDefaultsAndBuildParams(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to apply defaults: %w", err)
	}

	if opts.Limit < 1 || opts.Limit > 20 {
		return nil, errors.New("limit must be between 1 and 20")
	}

	// No additional parameters needed - all handled by ApplyDefaultsAndBuildParams

	result, err := c.request(ctx, EndpointDefiV2TokensNewListing, requestOptions{
		method:          "GET",
		chains:          opts.Chains,
		onLimitExceeded: RateLimitBehavior(opts.OnLimitExceeded),
		paramsOrBody:    params,
		respJustItems:   true,
	})
	if err != nil {
		return nil, err
	}

	var newListings map[string]RespTokenNewListing
	data, _ := json.Marshal(result)
	if err := json.Unmarshal(data, &newListings); err != nil {
		return nil, err
	}

	return newListings["items"], nil
}

// ============================================================================
// API Methods - Wallet Trades and Balance
// ============================================================================

// WalletTradesOptions holds options for GetWalletTrades.
//
// Note: This endpoint is only available for Solana chain.
type WalletTradesOptions struct {
	// Offset is the number of trades to skip for pagination. Optional, default: 0
	Offset int64 `default:"0"`

	// Limit is the maximum number of trades to return (1-100). Optional, default: 100, max: 100
	Limit int64 `default:"100"`

	// BeforeTime filters trades before this Unix timestamp. Optional, default: nil
	BeforeTime int64 `default:"0"`

	// AfterTime filters trades after this Unix timestamp. Optional, default: nil
	AfterTime int64 `default:"0"`

	// UIAmountMode specifies the token amount display mode. Options: "raw", "scaled", "both". Optional, default: "raw"
	UIAmountMode string `default:"raw"`

	// Chains is the list of blockchain networks to query. Optional, default: nil
	Chains []Chain

	// OnLimitExceeded overrides the default rate limit behavior. Optional, default: nil
	OnLimitExceeded string `default:""`
}

// GetWalletTrades retrieves trading history for a wallet.
//
// This method retrieves the trading history for a specific wallet including all trades,
// volumes, and profit/loss information.
//
// Note: This endpoint is only available for Solana chain.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - walletAddress: Wallet address to query
//   - opts: Configuration options (see WalletTradesOptions for details)
//
// Returns:
//   - *RespWalletTrades: Wallet trades response containing:
//   - totalTrades: Total number of trades
//   - totalVolume: Total trading volume
//   - totalProfit: Total profit/loss
//   - trades: List of individual trades
//   - Each trade contains: token, amount, price, timestamp, type
//   - error: Error if request fails or wallet address is invalid
//
// Raises:
//   - Context errors: If ctx is cancelled or times out
//   - API errors: If the Birdeye API returns an error
//   - Validation errors: If wallet address is invalid
//   - Rate limit errors: If rate limit exceeded
//
// Example:
//
//	trades, err := client.GetWalletTrades(ctx, "wallet-address", WalletTradesOptions{
//	    Limit: 50,
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Total trades: %d\n", len(trades.Items))
func (c *HTTPClient) GetWalletTrades(ctx context.Context, walletAddress string, opts *WalletTradesOptions) (*RespWalletTrades, error) {
	if opts == nil {
		opts = &WalletTradesOptions{}
	}
	params, err := ApplyDefaultsAndBuildParams(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to apply defaults: %w", err)
	}

	// Add required parameters
	params["address"] = walletAddress

	if opts.BeforeTime > 0 {
		params["before_time"] = opts.BeforeTime
	}
	if opts.AfterTime > 0 {
		params["after_time"] = opts.AfterTime
	}

	result, err := c.request(ctx, EndpointTraderTxsSeekByTime, requestOptions{
		method:          "GET",
		chains:          opts.Chains,
		onLimitExceeded: RateLimitBehavior(opts.OnLimitExceeded),
		paramsOrBody:    params,
	})
	if err != nil {
		return nil, err
	}

	var trades RespWalletTrades
	data, _ := json.Marshal(result)
	if err := json.Unmarshal(data, &trades); err != nil {
		return nil, err
	}

	return &trades, nil
}

// WalletTokenBalanceOptions holds options for GetWalletTokenBalance.
type WalletTokenBalanceOptions struct {
	// UIAmountMode specifies the token amount display mode. Options: "raw", "scaled", "both". Optional, default: "raw"
	UIAmountMode string `default:"raw"`

	// Chains is the list of blockchain networks to query. Optional, default: nil
	Chains []Chain

	// OnLimitExceeded overrides the default rate limit behavior. Optional, default: nil
	OnLimitExceeded string `default:""`
}

// GetWalletTokenBalance retrieves token balance for a wallet.
//
// This method retrieves the current token balance for a specific wallet address
// across supported networks.
//
// Note: This endpoint is only available for Solana chain.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - wallet: Wallet address to query
//   - tokenAddress: Token contract address to query
//   - opts: Configuration options (see WalletTokenBalanceOptions for details)
//
// Returns:
//   - *RespWalletTokenBalance: Wallet token balance response containing:
//   - totalBalance: Total wallet balance in USD
//   - tokens: List of token balances
//   - Each item contains: token, balance, value, price
//   - error: Error if request fails or wallet address is invalid
//
// Raises:
//   - Context errors: If ctx is cancelled or times out
//   - API errors: If the Birdeye API returns an error
//   - Validation errors: If wallet address is invalid
//   - Rate limit errors: If rate limit exceeded
//
// Example:
//
//	balance, err := client.GetWalletTokenBalance(ctx, "wallet-address", "token-address",
//	    WalletTokenBalanceOptions{})
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Balance: %.2f\n", balance.UIAmount)
func (c *HTTPClient) GetWalletTokenBalance(ctx context.Context, wallet, tokenAddress string, opts *WalletTokenBalanceOptions) (*RespWalletTokenBalance, error) {
	if opts == nil {
		opts = &WalletTokenBalanceOptions{}
	}
	params, err := ApplyDefaultsAndBuildParams(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to apply defaults: %w", err)
	}

	// Add required parameters
	params["wallet"] = wallet
	params["token_address"] = tokenAddress

	result, err := c.request(ctx, EndpointV1WalletTokenBalance, requestOptions{
		method:          "GET",
		chains:          opts.Chains,
		onLimitExceeded: RateLimitBehavior(opts.OnLimitExceeded),
		paramsOrBody:    params,
	})
	if err != nil {
		return nil, err
	}

	var balance RespWalletTokenBalance
	data, _ := json.Marshal(result)
	if err := json.Unmarshal(data, &balance); err != nil {
		return nil, err
	}

	return &balance, nil
}

// WalletNetWorthHistoriesOptions holds options for GetWalletNetWorthHistories.
//
// Note: This endpoint is only available for Solana chain.
type WalletNetWorthHistoriesOptions struct {
	// Count is the number of data points to return (1-30). Optional, default: 7, max: 30
	Count int64 `default:"7"`

	// Direction specifies the direction to query. Options: "back", "forward". Optional, default: "back"
	Direction string `default:"back"`

	// Time is the reference time for the query. Optional, default: nil (current time)
	Time string `default:""`

	// Type specifies the time interval. Options: "1d", "1w", "1m". Optional, default: "1d"
	Type string `default:"1d"`

	// SortType specifies the sort order. Options: "desc", "asc". Optional, default: "desc"
	SortType string `default:"desc"`

	// Chains is the list of blockchain networks to query. Optional, default: nil
	Chains []Chain

	// OnLimitExceeded overrides the default rate limit behavior. Optional, default: nil
	OnLimitExceeded string `default:""`
}

// GetWalletNetWorthHistories retrieves net worth history for a wallet.
//
// This method retrieves the historical net worth data for a specific wallet address over time,
// useful for tracking portfolio performance.
//
// Note: This endpoint is only available for Solana chain.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - wallet: Wallet address to query
//   - opts: Configuration options (see WalletNetWorthHistoriesOptions for details)
//
// Returns:
//   - *RespWalletNetWorthHistories: Wallet net worth history response containing:
//   - histories: List of historical data
//   - Each item contains: timestamp, totalValue, changePercent
//   - error: Error if request fails or wallet address is invalid
//
// Raises:
//   - Context errors: If ctx is cancelled or times out
//   - API errors: If the Birdeye API returns an error
//   - Validation errors: If wallet address is invalid or count out of range (1-30)
//   - Rate limit errors: If rate limit exceeded
//
// Example:
//
//	histories, err := client.GetWalletNetWorthHistories(ctx, "wallet-address",
//	    WalletNetWorthHistoriesOptions{
//	        Count: 7,
//	        Type:  "1d",
//	    })
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, h := range histories.Items {
//	    fmt.Printf("Date: %s, Value: $%.2f\n", h.Time, h.TotalValue)
//	}
func (c *HTTPClient) GetWalletNetWorthHistories(ctx context.Context, wallet string, opts *WalletNetWorthHistoriesOptions) (*RespWalletNetWorthHistories, error) {
	if opts == nil {
		opts = &WalletNetWorthHistoriesOptions{}
	}
	params, err := ApplyDefaultsAndBuildParams(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to apply defaults: %w", err)
	}

	if opts.Count < 1 || opts.Count > 30 {
		return nil, errors.New("count must be between 1 and 30")
	}

	// Add required parameters
	params["wallet"] = wallet

	if opts.Time != "" {
		params["time"] = opts.Time
	}

	result, err := c.request(ctx, EndpointV2WalletNetWorth, requestOptions{
		method:          "GET",
		chains:          opts.Chains,
		onLimitExceeded: RateLimitBehavior(opts.OnLimitExceeded),
		paramsOrBody:    params,
	})
	if err != nil {
		return nil, err
	}

	var histories RespWalletNetWorthHistories
	data, _ := json.Marshal(result)
	if err := json.Unmarshal(data, &histories); err != nil {
		return nil, err
	}

	return &histories, nil
}

// ============================================================================
// API Methods - Latest Block Number
// ============================================================================

// GetLatestBlockNumber retrieves the latest block number for supported blockchain networks.
//
// This method retrieves the current block number for the specified blockchain networks,
// which is useful for tracking blockchain state and ensuring data freshness.
//
// Parameters:
//   - ctx: Context for cancellation and timeout
//   - chains: List of blockchain networks to query
//
// Returns:
//   - int64: Latest block number
//   - error: Any error encountered during the request
//
// Example:
//
//	blockNum, err := client.GetLatestBlockNumber(ctx, []Chain{ChainSolana})
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Latest block: %d\n", blockNum)
func (c *HTTPClient) GetLatestBlockNumber(ctx context.Context, chains []Chain) (int64, error) {
	result, err := c.request(ctx, EndpointDefiV3TxsLatestBlock, requestOptions{
		method: "GET",
		chains: chains,
	})
	if err != nil {
		return 0, err
	}

	if blockNum, ok := result["block_number"].(float64); ok {
		return int64(blockNum), nil
	}

	return 0, errors.New("block_number not found in response")
}

// ============================================================================
// API Methods - Token Top Traders
// ============================================================================

// TokenTopTradersOptions holds options for GetTokenTopTraders.
type TokenTopTradersOptions struct {
	// TimeFrame specifies the time period. Options: "30m", "1h", "2h", "4h", "6h", "8h", "12h", "24h". Default: "24h"
	TimeFrame string `default:"24h"`

	// SortType specifies the sort order. Options: "desc", "asc". Default: "desc"
	SortType string `default:"desc"`

	// SortBy specifies the field to sort by. Options: "volume", "trade". Default: "volume"
	SortBy string `default:"volume"`

	// Offset is the number of traders to skip for pagination. Default: 0
	Offset int64 `default:"0"`

	// Limit is the maximum number of traders to return (1-10). Default: 10, max: 10
	Limit int64 `default:"10"`

	// UIAmountMode specifies the token amount display mode. Options: "raw", "scaled", "both". Default: "raw"
	UIAmountMode string `default:"raw"`

	// Chains is the list of blockchain networks to query. Default: nil
	Chains []Chain

	// OnLimitExceeded overrides the default rate limit behavior. Default: "" (use client default)
	OnLimitExceeded string `default:""`
}

// GetTokenTopTraders retrieves top traders for a token.
//
// This method retrieves information about the top traders for a specific token, including
// their trading activity and statistics.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - address: Token contract address to query
//   - opts: Configuration options (see TokenTopTradersOptions for details)
//
// Returns:
//   - []RespTokenTopTraderItem: Top traders response containing list of top traders
//     Each item contains: wallet address, trade count, volume, profit/loss
//   - error: Error if request fails or address is invalid
//
// Raises:
//   - Context errors: If ctx is cancelled or times out
//   - API errors: If the Birdeye API returns an error
//   - Validation errors: If address is invalid or offset/limit out of range
//   - Rate limit errors: If rate limit exceeded
//
// Example:
//
//	topTraders, err := client.GetTokenTopTraders(ctx, tokenAddress, TokenTopTradersOptions{
//	    Limit: 10,
//	    SortBy: "volume",
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, trader := range topTraders {
//	    fmt.Printf("Trader: %s, Volume: $%.2f\n", trader.Address, trader.Volume)
//	}
func (c *HTTPClient) GetTokenTopTraders(ctx context.Context, address string, opts *TokenTopTradersOptions) (RespTokenTopTraders, error) {
	if opts == nil {
		opts = &TokenTopTradersOptions{}
	}
	params, err := ApplyDefaultsAndBuildParams(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to apply defaults: %w", err)
	}

	if opts.Offset < 0 || opts.Offset > 10000 {
		return nil, errors.New("offset must be between 0 and 10000")
	}
	if opts.Limit < 1 || opts.Limit > 10 {
		return nil, errors.New("limit must be between 1 and 10")
	}
	if opts.Offset+opts.Limit > 10000 {
		return nil, errors.New("offset + limit must not exceed 10000")
	}

	// Add required parameters
	params["address"] = address

	result, err := c.request(ctx, EndpointDefiV2TokensTopTraders, requestOptions{
		method:          "GET",
		chains:          opts.Chains,
		onLimitExceeded: RateLimitBehavior(opts.OnLimitExceeded),
		paramsOrBody:    params,
		respJustItems:   true,
	})
	if err != nil {
		return nil, err
	}

	var traders map[string]RespTokenTopTraders
	data, _ := json.Marshal(result)
	if err := json.Unmarshal(data, &traders); err != nil {
		return nil, err
	}

	return traders["items"], nil
}

// ============================================================================
// API Methods - Token All Market List
// ============================================================================

// TokenAllMarketListOptions holds options for GetTokenAllMarketList.
type TokenAllMarketListOptions struct {
	// TimeFrame specifies the time period. Options: "30m", "1h", "2h", "4h", "6h", "8h", "12h", "24h". Default: "24h"
	TimeFrame string `default:"24h"`
	// SortType specifies the sort order. Options: "desc", "asc". Default: "desc"
	SortType string `default:"desc"`
	// SortBy specifies the field to sort by. Options: "liquidity", "volume24h". Default: "liquidity"
	SortBy string `default:"liquidity"`
	// Offset is the number of markets to skip for pagination. Default: 0
	Offset int64 `default:"0"`
	// Limit is the maximum number of markets to return (1-20). Default: 20, max: 20
	Limit int64 `default:"20"`
	// Chains is the list of blockchain networks to query. Default: nil
	Chains []Chain
	// OnLimitExceeded overrides the default rate limit behavior. Default: "" (use client default)
	OnLimitExceeded string `default:""`
}

// GetTokenAllMarketList retrieves all market information for a token.
//
// This method retrieves comprehensive market information including all trading pairs, exchanges,
// and market data for a token.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - address: Token contract address to query
//   - opts: Configuration options (see TokenAllMarketListOptions for details)
//
// Returns:
//   - *RespTokenAllMarketList: All market list response containing:
//   - items: List of market information
//   - Each item contains: exchange, pair, price, volume, liquidity
//   - error: Error if request fails or address is invalid
//
// Raises:
//   - Context errors: If ctx is cancelled or times out
//   - API errors: If the Birdeye API returns an error
//   - Validation errors: If address is invalid or limit out of range (1-20)
//   - Rate limit errors: If rate limit exceeded
//
// Example:
//
//	markets, err := client.GetTokenAllMarketList(ctx, tokenAddress, TokenAllMarketListOptions{
//	    SortBy: "liquidity",
//	    Limit:  20,
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, market := range markets.Items {
//	    fmt.Printf("Market: %s, Liquidity: $%.2f\n", market.Name, market.Liquidity)
//	}
func (c *HTTPClient) GetTokenAllMarketList(ctx context.Context, address string, opts *TokenAllMarketListOptions) (*RespTokenAllMarketList, error) {
	if opts == nil {
		opts = &TokenAllMarketListOptions{}
	}
	params, err := ApplyDefaultsAndBuildParams(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to apply defaults: %w", err)
	}

	if opts.Limit < 1 || opts.Limit > 20 {
		return nil, errors.New("limit must be between 1 and 20")
	}

	// Add required parameters
	params["address"] = address

	result, err := c.request(ctx, EndpointDefiV2Markets, requestOptions{
		method:          "GET",
		chains:          opts.Chains,
		onLimitExceeded: RateLimitBehavior(opts.OnLimitExceeded),
		paramsOrBody:    params,
	})
	if err != nil {
		return nil, err
	}

	var marketList RespTokenAllMarketList
	data, _ := json.Marshal(result)
	if err := json.Unmarshal(data, &marketList); err != nil {
		return nil, err
	}

	return &marketList, nil
}

// ============================================================================
// API Methods - Gainers and Losers
// ============================================================================

// GainersLosersOptions holds options for GetGainersLosers.
//
// Note: This endpoint is only available for Solana chain.
type GainersLosersOptions struct {
	// Type specifies the time period. Options: "yesterday", "today", "1W". Default: "1W"
	Type string `default:"1W"`
	// SortBy specifies the field to sort by. Options: "PnL". Default: "PnL"
	SortBy string `default:"PnL"`
	// SortType specifies the sort order. Options: "desc", "asc". Default: "desc"
	SortType string `default:"desc"`
	// Offset is the number of traders to skip for pagination. Default: 0
	Offset int64 `default:"0"`
	// Limit is the maximum number of traders to return (1-10). Default: 10, max: 10
	Limit int64 `default:"10"`
	// Chains is the list of blockchain networks to query. Default: nil
	Chains []Chain
	// OnLimitExceeded overrides the default rate limit behavior. Default: "" (use client default)
	OnLimitExceeded string `default:""`
}

// GetGainersLosers retrieves top gainers and losers tokens.
//
// This method retrieves a list of tokens with the highest gains and losses over a specified
// time period.
//
// Note: This endpoint is only available for Solana chain.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - opts: Configuration options (see GainersLosersOptions for details)
//
// Returns:
//   - []RespGainerLoserItem: Gainers/losers response containing:
//   - gainers: List of top gaining tokens
//   - losers: List of top losing tokens
//   - Each item contains: address, symbol, price, changePercent
//   - error: Error if request fails
//
// Raises:
//   - Context errors: If ctx is cancelled or times out
//   - API errors: If the Birdeye API returns an error
//   - Rate limit errors: If rate limit exceeded
//
// Example:
//
//	gainersLosers, err := client.GetGainersLosers(ctx, GainersLosersOptions{
//	    Type:  "1W",
//	    Limit: 10,
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, item := range gainersLosers {
//	    fmt.Printf("%s: %.2f%%\n", item.Symbol, item.PriceChangePercent)
//	}
func (c *HTTPClient) GetGainersLosers(ctx context.Context, opts *GainersLosersOptions) (RespGainerLosers, error) {
	if opts == nil {
		opts = &GainersLosersOptions{}
	}
	params, err := ApplyDefaultsAndBuildParams(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to apply defaults: %w", err)
	}

	// No additional parameters needed - all handled by ApplyDefaultsAndBuildParams

	result, err := c.request(ctx, EndpointTraderGainersLosers, requestOptions{
		method:          "GET",
		chains:          opts.Chains,
		onLimitExceeded: RateLimitBehavior(opts.OnLimitExceeded),
		paramsOrBody:    params,
	})
	if err != nil {
		return nil, err
	}

	var gainersLosers map[string]RespGainerLosers
	data, _ := json.Marshal(result)
	if err := json.Unmarshal(data, &gainersLosers); err != nil {
		return nil, err
	}

	return gainersLosers["items"], nil
}

// ============================================================================
// API Methods - Token All Time Trades
// ============================================================================

// TokenAllTimeTradesOptions holds options for GetTokenAllTimeTrades and GetMultiTokenAllTimeTrades.
type TokenAllTimeTradesOptions struct {
	// TimeFrame specifies the time period. Options: "1m", "5m", "30m", "1h", "2h", "4h", "8h", "24h", "3d", "7d", "14d", "30d", "90d", "180d", "1y", "alltime". Default: "24h"
	TimeFrame string `default:"24h"`
	// UIAmountMode specifies the token amount display mode. Options: "raw", "scaled", "both". Default: "raw"
	UIAmountMode string `default:"raw"`
	// Chains is the list of blockchain networks to query. Default: nil
	Chains []Chain
	// OnLimitExceeded overrides the default rate limit behavior. Default: "" (use client default)
	OnLimitExceeded string `default:""`
}

// GetTokenAllTimeTrades retrieves all-time trading data for a token.
//
// This method retrieves comprehensive trading data for a token including all-time statistics
// and trading history.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - address: Token contract address to query
//   - opts: Configuration options (see TokenAllTimeTradesOptions for details)
//
// Returns:
//   - *RespTokenAllTimeTrades: All-time trades response containing:
//   - totalTrades: Total number of trades
//   - totalVolume: Total trading volume
//   - avgTradeSize: Average trade size
//   - firstTrade: First trade timestamp
//   - lastTrade: Last trade timestamp
//   - error: Error if request fails or address is invalid
//
// Raises:
//   - Context errors: If ctx is cancelled or times out
//   - API errors: If the Birdeye API returns an error
//   - Validation errors: If address is invalid
//   - Rate limit errors: If rate limit exceeded
//
// Example:
//
//	allTimeTrades, err := client.GetTokenAllTimeTrades(ctx, tokenAddress, TokenAllTimeTradesOptions{
//	    TimeFrame: "alltime",
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Total Trades: %d\n", allTimeTrades.TotalTrades)
//	fmt.Printf("Total Volume: $%.2f\n", allTimeTrades.TotalVolume)
func (c *HTTPClient) GetTokenAllTimeTrades(ctx context.Context, address string, opts *TokenAllTimeTradesOptions) (*RespTokenAllTimeTrades, error) {
	if opts == nil {
		opts = &TokenAllTimeTradesOptions{}
	}
	params, err := ApplyDefaultsAndBuildParams(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to apply defaults: %w", err)
	}

	// Add required parameters
	params["address"] = address

	result, err := c.request(ctx, EndpointDefiV3AllTimeTradesSingle, requestOptions{
		method:          "GET",
		chains:          opts.Chains,
		onLimitExceeded: RateLimitBehavior(opts.OnLimitExceeded),
		paramsOrBody:    params,
	})
	if err != nil {
		return nil, err
	}

	// Result may be a list, get first element
	data, _ := json.Marshal(result)

	// Try to unmarshal as array first
	// var tradesArray []RespTokenAllTimeTrades
	// if err := json.Unmarshal(data, &tradesArray); err == nil && len(tradesArray) > 0 {
	// 	return &tradesArray[0], nil
	// }

	// Otherwise unmarshal as single object
	var trades map[string][]RespTokenAllTimeTrades
	if err := json.Unmarshal(data, &trades); err != nil {
		return nil, err
	}

	l := trades["data"]

	if len(l) > 0 {
		return &l[0], nil
	}

	return nil, nil
}

// GetMultiTokenAllTimeTrades retrieves all-time trading data for multiple tokens.
//
// This method retrieves comprehensive trading data for multiple tokens in a single request,
// improving efficiency for batch operations.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - addresses: List of token contract addresses to query
//   - opts: Configuration options (see TokenAllTimeTradesOptions for details)
//
// Returns:
//   - []RespTokenAllTimeTrades: All-time trades response containing list of token all-time trades
//     Each item contains: totalTrades, totalVolume, avgTradeSize
//   - error: Error if request fails or any address is invalid
//
// Raises:
//   - Context errors: If ctx is cancelled or times out
//   - API errors: If the Birdeye API returns an error
//   - Validation errors: If any address is invalid
//   - Rate limit errors: If rate limit exceeded
//
// Example:
//
//	addresses := []string{"token1", "token2"}
//	allTimeTrades, err := client.GetMultiTokenAllTimeTrades(ctx, addresses, TokenAllTimeTradesOptions{
//	    TimeFrame: "24h",
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, trades := range allTimeTrades {
//	    fmt.Printf("Total Volume: $%.2f\n", trades.TotalVolume)
//	}
func (c *HTTPClient) GetMultiTokenAllTimeTrades(ctx context.Context, addresses []string, opts *TokenAllTimeTradesOptions) (RespMultiTokenAllTimeTrades, error) {
	if opts == nil {
		opts = &TokenAllTimeTradesOptions{}
	}
	params, err := ApplyDefaultsAndBuildParams(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to apply defaults: %w", err)
	}

	// Add required parameters
	params["list_address"] = addresses

	result, err := c.request(ctx, EndpointDefiV3AllTimeTradesMultiple, requestOptions{
		method:          "POST",
		chains:          opts.Chains,
		onLimitExceeded: RateLimitBehavior(opts.OnLimitExceeded),
		paramsOrBody:    params,
	})
	if err != nil {
		return nil, err
	}

	var trades map[string]RespMultiTokenAllTimeTrades
	data, _ := json.Marshal(result)
	if err := json.Unmarshal(data, &trades); err != nil {
		return nil, err
	}

	return trades["data"], nil
}

// ============================================================================
// API Methods - Token Price Volume
// ============================================================================

// TokenPriceVolumeOptions holds options for GetTokenPriceVolume and GetMultiTokenPriceVolume.
type TokenPriceVolumeOptions struct {
	// Type specifies the time period for volume calculation. Options: "1h", "2h", "4h", "8h", "24h". Default: "24h"
	Type string `default:"24h"`
	// UIAmountMode specifies the token amount display mode. Options: "raw", "scaled", "both". Default: "raw"
	UIAmountMode string `default:"raw"`
	// Chains is the list of blockchain networks to query. Default: nil
	Chains []Chain
	// OnLimitExceeded overrides the default rate limit behavior. Default: "" (use client default)
	OnLimitExceeded string `default:""`
}

// GetTokenPriceVolume retrieves token price and trading volume data.
//
// This method retrieves current price and trading volume information for a specific token
// across supported networks.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - address: Token contract address to query
//   - opts: Configuration options (see TokenPriceVolumeOptions for details)
//
// Returns:
//   - *RespTokenPriceVolume: Price and volume response containing:
//   - price: Current token price
//   - volume24h: 24-hour trading volume
//   - priceChange24h: 24-hour price change percentage
//   - error: Error if request fails or address is invalid
//
// Raises:
//   - Context errors: If ctx is cancelled or times out
//   - API errors: If the Birdeye API returns an error
//   - Validation errors: If address is invalid
//   - Rate limit errors: If rate limit exceeded
//
// Example:
//
//	// Get SOL price and volume
//	solData, err := client.GetTokenPriceVolume(ctx,
//	    "So11111111111111111111111111111111111111112",  // SOL
//	    TokenPriceVolumeOptions{})
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("SOL Price: $%.2f\n", solData.Price)
//	fmt.Printf("24h Volume: $%.2f\n", solData.Volume24h)
func (c *HTTPClient) GetTokenPriceVolume(ctx context.Context, address string, opts *TokenPriceVolumeOptions) (*RespTokenPriceVolume, error) {
	if opts == nil {
		opts = &TokenPriceVolumeOptions{}
	}
	params, err := ApplyDefaultsAndBuildParams(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to apply defaults: %w", err)
	}

	// Add required parameters
	params["address"] = address

	result, err := c.request(ctx, EndpointDefiPriceVolumeSingle, requestOptions{
		method:          "GET",
		chains:          opts.Chains,
		onLimitExceeded: RateLimitBehavior(opts.OnLimitExceeded),
		paramsOrBody:    params,
	})
	if err != nil {
		return nil, err
	}

	var priceVolume RespTokenPriceVolume
	data, _ := json.Marshal(result)
	if err := json.Unmarshal(data, &priceVolume); err != nil {
		return nil, err
	}

	return &priceVolume, nil
}

// GetMultiTokenPriceVolume retrieves price and trading volume data for multiple tokens.
//
// This method retrieves current price and trading volume information for multiple tokens in
// a single request, improving efficiency.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - addresses: List of token contract addresses to query
//   - opts: Configuration options (see TokenPriceVolumeOptions for details)
//
// Returns:
//   - map[string]RespTokenPriceVolume: Price and volume response containing:
//   - items: List of token price/volume data
//   - Each item contains: price, volume24h, priceChange24h
//   - error: Error if request fails or any address is invalid
//
// Raises:
//   - Context errors: If ctx is cancelled or times out
//   - API errors: If the Birdeye API returns an error
//   - Validation errors: If any address is invalid
//   - Rate limit errors: If rate limit exceeded
//
// Example:
//
//	tokensData, err := client.GetMultiTokenPriceVolume(ctx, []string{
//	    "So11111111111111111111111111111111111111112",  // SOL
//	    "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v", // USDC
//	    "DezXAZ8z7PnrnRJjz3wXBoRgixCa6xjnB7YaB1pPB263",  // Bonk
//	}, TokenPriceVolumeOptions{})
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for addr, tokenData := range tokensData {
//	    fmt.Printf("%s: Price $%.2f\n", addr[:8], tokenData.Price)
//	}
func (c *HTTPClient) GetMultiTokenPriceVolume(ctx context.Context, addresses []string, opts *TokenPriceVolumeOptions) (map[string]RespTokenPriceVolume, error) {
	if opts == nil {
		opts = &TokenPriceVolumeOptions{}
	}
	params, err := ApplyDefaultsAndBuildParams(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to apply defaults: %w", err)
	}

	// Add required parameters
	params["list_address"] = addresses

	result, err := c.request(ctx, EndpointDefiPriceVolumeMulti, requestOptions{
		method:          "POST",
		chains:          opts.Chains,
		onLimitExceeded: RateLimitBehavior(opts.OnLimitExceeded),
		paramsOrBody:    params,
	})
	if err != nil {
		return nil, err
	}

	var priceVolumes map[string]RespTokenPriceVolume
	data, _ := json.Marshal(result)
	if err := json.Unmarshal(data, &priceVolumes); err != nil {
		return nil, err
	}

	return priceVolumes, nil
}

// ============================================================================
// API Methods - Token Price Histories
// ============================================================================

// TokenPriceHistoriesOptions holds options for GetTokenPriceHistories.
type TokenPriceHistoriesOptions struct {
	// UIAmountMode specifies the token amount display mode. Options: "raw", "scaled", "both". Default: "raw"
	UIAmountMode string `default:"raw"`
	// Chains is the list of blockchain networks to query. Default: nil
	Chains []Chain
	// OnLimitExceeded overrides the default rate limit behavior. Default: "" (use client default)
	OnLimitExceeded string `default:""`
}

// GetTokenPriceHistories retrieves historical price data for a token or trading pair.
//
// This method retrieves historical price data over a specified time range with configurable
// time intervals for tokens or trading pairs.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - address: Token or pair contract address to query
//   - addressType: Type of address. Options: "token", "pair"
//   - intervalType: Time interval. Options: "1m", "3m", "5m", "15m", "30m", "1H", "2H", "4H", "6H", "8H", "12H", "1D", "3D", "1W", "1M"
//   - timeFrom: Start time in Unix timestamp (seconds)
//   - timeTo: End time in Unix timestamp (seconds)
//   - opts: Configuration options (see TokenPriceHistoriesOptions for details)
//
// Returns:
//   - *RespTokenPriceHistories: Price history response containing:
//   - items: List of price data points
//   - Each item contains: price, timestamp, and other metrics
//   - error: Error if request fails, address is invalid, or time range is invalid
//
// Raises:
//   - Context errors: If ctx is cancelled or times out
//   - API errors: If the Birdeye API returns an error
//   - Validation errors: If address is invalid or time range is invalid
//   - Rate limit errors: If rate limit exceeded
//
// Example:
//
//	// Get daily price history for SOL over the last month
//	endTime := time.Now().Unix()
//	startTime := endTime - 2592000  // 30 days ago
//
//	priceHistory, err := client.GetTokenPriceHistories(ctx,
//	    "So11111111111111111111111111111111111111112",  // SOL
//	    "token", "1D", startTime, endTime, TokenPriceHistoriesOptions{})
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, pricePoint := range priceHistory.Items {
//	    fmt.Printf("Date: %d, Price: $%.2f\n", pricePoint.UnixTime, pricePoint.Value)
//	}
func (c *HTTPClient) GetTokenPriceHistories(ctx context.Context, address, addressType, intervalType string, timeFrom, timeTo int64, opts *TokenPriceHistoriesOptions) (*RespTokenPriceHistories, error) {
	if opts == nil {
		opts = &TokenPriceHistoriesOptions{}
	}
	params, err := ApplyDefaultsAndBuildParams(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to apply defaults: %w", err)
	}

	if timeFrom < 0 || timeFrom > 10000000000 {
		return nil, errors.New("time_from must be between 0 and 10000000000")
	}
	if timeTo < 0 || timeTo > 10000000000 {
		return nil, errors.New("time_to must be between 0 and 10000000000")
	}

	// Add required parameters
	params["address"] = address
	params["address_type"] = addressType
	params["type"] = intervalType
	params["time_from"] = timeFrom
	params["time_to"] = timeTo

	result, err := c.request(ctx, EndpointDefiHistoryPrice, requestOptions{
		method:          "GET",
		chains:          opts.Chains,
		onLimitExceeded: RateLimitBehavior(opts.OnLimitExceeded),
		paramsOrBody:    params,
	})
	if err != nil {
		return nil, err
	}

	var histories RespTokenPriceHistories
	data, _ := json.Marshal(result)
	if err := json.Unmarshal(data, &histories); err != nil {
		return nil, err
	}

	return &histories, nil
}

// GetTokenPriceHistoryByTime retrieves token price at a specific time point.
//
// This method retrieves the token price at a specific Unix timestamp, useful for historical
// price analysis and backtesting.
//
// Note: This endpoint is only available for Solana chain.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - address: Token contract address to query
//   - unixTime: Unix timestamp in seconds for the price query (0-10000000000)
//   - opts: Configuration options (see TokenPriceHistoriesOptions for details)
//
// Returns:
//   - *RespTokenPriceHistoryByTime: Price history response containing:
//   - value: Token price at the specified time
//   - unixtime: The queried timestamp
//   - priceChange24h: 24-hour price change at that time
//   - error: Error if request fails, address is invalid, or timestamp is invalid
//
// Raises:
//   - Context errors: If ctx is cancelled or times out
//   - API errors: If the Birdeye API returns an error
//   - Validation errors: If address or timestamp is invalid
//
// Example:
//
//	// Get SOL price on January 1, 2024
//	solPrice, err := client.GetTokenPriceHistoryByTime(ctx,
//	    "So11111111111111111111111111111111111111112",  // SOL
//	    1704067200,  // Jan 1, 2024 00:00:00 UTC
//	    TokenPriceHistoriesOptions{})
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("SOL price on Jan 1, 2024: $%.2f\n", solPrice.Value)
func (c *HTTPClient) GetTokenPriceHistoryByTime(ctx context.Context, address string, unixTime int64, opts *TokenPriceHistoriesOptions) (*RespTokenPriceHistoryByTime, error) {
	if opts == nil {
		opts = &TokenPriceHistoriesOptions{}
	}
	params, err := ApplyDefaultsAndBuildParams(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to apply defaults: %w", err)
	}

	if unixTime < 0 || unixTime > 10000000000 {
		return nil, errors.New("unixtime must be between 0 and 10000000000")
	}

	// Add required parameters
	params["address"] = address
	params["unixtime"] = unixTime

	result, err := c.request(ctx, EndpointDefiHistoricalPriceUnix, requestOptions{
		method:          "GET",
		chains:          opts.Chains,
		onLimitExceeded: RateLimitBehavior(opts.OnLimitExceeded),
		paramsOrBody:    params,
	})
	if err != nil {
		return nil, err
	}

	var history RespTokenPriceHistoryByTime
	data, _ := json.Marshal(result)
	if err := json.Unmarshal(data, &history); err != nil {
		return nil, err
	}

	return &history, nil
}

// ============================================================================
// API Methods - OHLCV V3
// ============================================================================

// TokenOHLCVV3Options holds options for GetTokenOHLCVV3 and GetPairOHLCVV3.
type TokenOHLCVV3Options struct {
	// Currency specifies the price currency. Options: "usd", "native". Default: "usd"
	Currency string `default:"usd"`
	// Mode specifies the query mode. Options: "range" (by time range), "count" (by count). Default: "range"
	Mode string `default:"range"`
	// CountLimit is the maximum number of OHLCV data points when mode is "count". Default: 5000
	CountLimit int64 `default:"5000"`
	// Padding adds empty candles for missing time periods if true. Default: false
	Padding bool `default:"false"`
	// Outlier includes outlier detection data if true. Default: true
	Outlier bool `default:"false"`
	// UIAmountMode specifies the token amount display mode. Options: "raw", "scaled", "both". Default: "raw"
	UIAmountMode string `default:"raw"`
	// Chains is the list of blockchain networks to query. Default: nil
	Chains []Chain
	// OnLimitExceeded overrides the default rate limit behavior. Default: "" (use client default)
	OnLimitExceeded string `default:""`
}

// GetTokenOHLCVV3 retrieves OHLCV data for a token using the V3 API.
//
// This method retrieves historical OHLCV data for a token using the V3 API, which provides
// enhanced data structure and improved performance.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - address: Token contract address to query
//   - intervalType: Time interval. Options: "1s", "15s", "1m", "3m", "5m", "15m", "30m", "1H", "2H", "4H", "6H", "8H", "12H", "1D", "3D", "1W", "1M"
//   - timeFrom: Start time in Unix timestamp (seconds)
//   - timeTo: End time in Unix timestamp (seconds)
//   - opts: Configuration options (see TokenOHLCVV3Options for details)
//
// Returns:
//   - *RespTokenOHLCVsV3: OHLCV response containing list of OHLCV data points
//     Each item contains: open, high, low, close, volume, timestamp
//   - error: Error if request fails, address is invalid, or time range/count_limit is invalid
//
// Raises:
//   - Context errors: If ctx is cancelled or times out
//   - API errors: If the Birdeye API returns an error
//   - Validation errors: If address, time range, or count_limit is invalid
//   - Rate limit errors: If rate limit exceeded
func (c *HTTPClient) GetTokenOHLCVV3(ctx context.Context, address, intervalType string, timeFrom, timeTo int64, opts *TokenOHLCVV3Options) (*RespTokenOHLCVsV3, error) {
	if opts == nil {
		opts = &TokenOHLCVV3Options{}
	}
	params, err := ApplyDefaultsAndBuildParams(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to apply defaults: %w", err)
	}

	if timeFrom < 0 || timeFrom > 10000000000 {
		return nil, errors.New("time_from must be between 0 and 10000000000")
	}
	if timeTo < 0 || timeTo > 10000000000 {
		return nil, errors.New("time_to must be between 0 and 10000000000")
	}
	if opts.CountLimit < 0 || opts.CountLimit > 5000 {
		return nil, errors.New("count_limit must be between 0 and 5000")
	}

	// Add required parameters
	params["address"] = address
	params["type"] = intervalType
	params["time_from"] = timeFrom
	params["time_to"] = timeTo
	params["currency"] = opts.Currency
	params["mode"] = opts.Mode
	params["count_limit"] = opts.CountLimit
	params["padding"] = fmt.Sprintf("%t", opts.Padding)
	params["outlier"] = fmt.Sprintf("%t", opts.Outlier)

	result, err := c.request(ctx, EndpointDefiV3OHLCV, requestOptions{
		method:          "GET",
		chains:          opts.Chains,
		onLimitExceeded: RateLimitBehavior(opts.OnLimitExceeded),
		paramsOrBody:    params,
	})
	if err != nil {
		return nil, err
	}

	var ohlcv RespTokenOHLCVsV3
	data, _ := json.Marshal(result)
	if err := json.Unmarshal(data, &ohlcv); err != nil {
		return nil, err
	}

	return &ohlcv, nil
}

// GetPairOHLCVV3 retrieves OHLCV data for a trading pair using the V3 API.
//
// This method retrieves historical OHLCV data for a trading pair using the V3 API, which provides
// enhanced data structure and improved performance.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - address: Trading pair address to query
//   - intervalType: Time interval. Options: "1s", "15s", "1m", "3m", "5m", "15m", "30m", "1H", "2H", "4H", "6H", "8H", "12H", "1D", "3D", "1W", "1M"
//   - timeFrom: Start time in Unix timestamp (seconds)
//   - timeTo: End time in Unix timestamp (seconds)
//   - opts: Configuration options (see TokenOHLCVV3Options for details)
//
// Returns:
//   - []RespPairOHLCVItemV3: OHLCV response containing list of OHLCV data points
//     Each item contains: open, high, low, close, volume, timestamp
//   - error: Error if request fails, address is invalid, or time range/count_limit is invalid
//
// Raises:
//   - Context errors: If ctx is cancelled or times out
//   - API errors: If the Birdeye API returns an error
//   - Validation errors: If address, time range, or count_limit is invalid
func (c *HTTPClient) GetPairOHLCVV3(ctx context.Context, address, intervalType string, timeFrom, timeTo int64, opts *TokenOHLCVV3Options) ([]RespPairOHLCVItemV3, error) {
	if opts == nil {
		opts = &TokenOHLCVV3Options{}
	}
	params, err := ApplyDefaultsAndBuildParams(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to apply defaults: %w", err)
	}

	if timeFrom < 0 || timeFrom > 10000000000 {
		return nil, errors.New("time_from must be between 0 and 10000000000")
	}
	if timeTo < 0 || timeTo > 10000000000 {
		return nil, errors.New("time_to must be between 0 and 10000000000")
	}
	if opts.CountLimit < 0 || opts.CountLimit > 5000 {
		return nil, errors.New("count_limit must be between 0 and 5000")
	}

	// Add required parameters
	params["address"] = address
	params["type"] = intervalType
	params["time_from"] = timeFrom
	params["time_to"] = timeTo
	params["currency"] = opts.Currency
	params["mode"] = opts.Mode
	params["count_limit"] = opts.CountLimit
	params["padding"] = fmt.Sprintf("%t", opts.Padding)
	params["outlier"] = fmt.Sprintf("%t", opts.Outlier)

	result, err := c.request(ctx, EndpointDefiV3OHLCVPair, requestOptions{
		method:          "GET",
		chains:          opts.Chains,
		onLimitExceeded: RateLimitBehavior(opts.OnLimitExceeded),
		paramsOrBody:    params,
		respJustItems:   true,
	})
	if err != nil {
		return nil, err
	}

	var ohlcv map[string][]RespPairOHLCVItemV3
	data, _ := json.Marshal(result)
	if err := json.Unmarshal(data, &ohlcv); err != nil {
		return nil, err
	}

	return ohlcv["items"], nil
}

// ============================================================================
// API Methods - Token Price Stats
// ============================================================================

// TokenPriceStatsOptions holds options for GetTokenPriceStats and GetMultiTokenPriceStats.
type TokenPriceStatsOptions struct {
	// UIAmountMode specifies the token amount display mode. Options: "raw", "scaled", "both". Default: "raw"
	UIAmountMode string `default:"raw"`
	// Chains is the list of blockchain networks to query. Default: nil
	Chains []Chain
	// OnLimitExceeded overrides the default rate limit behavior. Default: "" (use client default)
	OnLimitExceeded string `default:""`
}

// GetTokenPriceStats retrieves comprehensive price statistics for a token.
//
// This method retrieves detailed price statistics including historical performance, volatility
// metrics, and market indicators.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - address: Token contract address to query
//   - timeframes: List of timeframes, max 3. Options: "1m", "5m", "30m", "1h", "2h", "4h", "8h", "24h", "2d", "3d", "7d"
//   - opts: Configuration options (see TokenPriceStatsOptions for details)
//
// Returns:
//   - *RespTokenPriceStats: Price statistics response containing:
//   - currentPrice: Current token price
//   - priceChange24h: 24-hour price change percentage
//   - priceChange7d: 7-day price change percentage
//   - priceChange30d: 30-day price change percentage
//   - volume24h: 24-hour trading volume
//   - marketCap: Current market capitalization
//   - error: Error if request fails or address is invalid
//
// Raises:
//   - Context errors: If ctx is cancelled or times out
//   - API errors: If the Birdeye API returns an error
//   - Validation errors: If address is invalid
func (c *HTTPClient) GetTokenPriceStats(ctx context.Context, address string, timeframes []string, opts *TokenPriceStatsOptions) (*RespTokenPriceStats, error) {
	if opts == nil {
		opts = &TokenPriceStatsOptions{}
	}
	params, err := ApplyDefaultsAndBuildParams(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to apply defaults: %w", err)
	}

	// Add required parameters
	params["address"] = address
	params["list_timeframe"] = timeframes

	result, err := c.request(ctx, EndpointDefiV3PriceStatsSingle, requestOptions{
		method:          "GET",
		chains:          opts.Chains,
		onLimitExceeded: RateLimitBehavior(opts.OnLimitExceeded),
		paramsOrBody:    params,
	})
	if err != nil {
		return nil, err
	}

	// Result may be a list, get first element
	var stats RespTokenPriceStats
	data, _ := json.Marshal(result)

	// Try to unmarshal as array first
	var statsArray []RespTokenPriceStats
	if err := json.Unmarshal(data, &statsArray); err == nil && len(statsArray) > 0 {
		return &statsArray[0], nil
	}

	// Otherwise unmarshal as single object
	if err := json.Unmarshal(data, &stats); err != nil {
		return nil, err
	}

	return &stats, nil
}

// GetMultiTokenPriceStats retrieves comprehensive price statistics for multiple tokens.
//
// This method retrieves detailed price statistics for multiple tokens in a single request,
// improving efficiency for batch operations.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - addresses: List of token contract addresses to query
//   - timeframes: List of timeframes. Options: "1m", "5m", "30m", "1h", "2h", "4h", "8h", "24h", "2d", "3d", "7d"
//   - opts: Configuration options (see TokenPriceStatsOptions for details)
//
// Returns:
//   - []RespTokenPriceStats: Price statistics response containing list of token price statistics
//     Each item contains: currentPrice, priceChange24h, volume24h, marketCap
//   - error: Error if request fails or any address is invalid
//
// Raises:
//   - Context errors: If ctx is cancelled or times out
//   - API errors: If the Birdeye API returns an error
//   - Validation errors: If any address is invalid
func (c *HTTPClient) GetMultiTokenPriceStats(ctx context.Context, addresses, timeframes []string, opts *TokenPriceStatsOptions) (RespMultiTokenPriceStats, error) {
	if opts == nil {
		opts = &TokenPriceStatsOptions{}
	}
	params, err := ApplyDefaultsAndBuildParams(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to apply defaults: %w", err)
	}

	// Add required parameters - handle different formats for different parameters
	// params["list_address"] = strings.Join(addresses, ",")

	path := EndpointDefiV3PriceStatsMultiple + "?" + "&list_timeframe=" + strings.Join(timeframes, ",") + "&ui_amount_mode=" + params["ui_amount_mode"].(string)

	result, err := c.request(ctx, path, requestOptions{
		method:          "POST",
		chains:          opts.Chains,
		onLimitExceeded: RateLimitBehavior(opts.OnLimitExceeded),
		paramsOrBody:    map[string]any{"list_address": strings.Join(addresses, ",")},
	})
	if err != nil {
		return nil, err
	}

	var stats map[string]RespMultiTokenPriceStats
	data, _ := json.Marshal(result)
	if err := json.Unmarshal(data, &stats); err != nil {
		return nil, err
	}

	return stats["data"], nil
}

// ============================================================================
// API Methods - Token Mint Burn Transactions
// ============================================================================

// TokenMintBurnTxsOptions holds options for GetTokenMintBurnTxs. Note: Solana only.
type TokenMintBurnTxsOptions struct {
	// SortBy specifies the field to sort by. Options: "block_time". Default: "block_time"
	SortBy string `default:"block_time"`
	// SortType specifies the sort order. Options: "desc", "asc". Default: "desc"
	SortType string `default:"desc"`
	// Type specifies the transaction type filter. Options: "mint", "burn", "all". Default: "all"
	Type string `default:"all"`
	// AfterTime filters transactions after this Unix timestamp. Default: 0 (no filter)
	AfterTime int64 `default:"0"`
	// BeforeTime filters transactions before this Unix timestamp. Default: 0 (no filter)
	BeforeTime int64 `default:"0"`
	// Offset is the number of transactions to skip for pagination. Default: 0
	Offset int64 `default:"0"`
	// Limit is the maximum number of transactions to return. Default: 100
	Limit int64 `default:"100"`
	// Chains is the list of blockchain networks to query. Default: nil
	Chains []Chain
	// OnLimitExceeded overrides the default rate limit behavior. Default: "" (use client default)
	OnLimitExceeded string `default:""`
}

// GetTokenMintBurnTxs retrieves mint and burn transactions for a token.
//
// This method retrieves all mint and burn transactions for a token, useful for tracking
// token supply changes.
//
// Note: This endpoint is only available for Solana chain.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - address: Token contract address to query
//   - opts: Configuration options (see TokenMintBurnTxsOptions for details)
//
// Returns:
//   - []RespTokenMintBurnTxItem: Mint/burn transactions response
//     Each item contains: txHash, type, amount, timestamp, from, to
//   - error: Error if request fails or address is invalid
//
// Raises:
//   - Context errors: If ctx is cancelled or times out
//   - API errors: If the Birdeye API returns an error
//   - Validation errors: If address is invalid
func (c *HTTPClient) GetTokenMintBurnTxs(ctx context.Context, address string, opts *TokenMintBurnTxsOptions) (RespTokenMintBurnTxs, error) {
	if opts == nil {
		opts = &TokenMintBurnTxsOptions{}
	}
	params, err := ApplyDefaultsAndBuildParams(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to apply defaults: %w", err)
	}

	// Add required parameters
	params["address"] = address

	if opts.AfterTime > 0 {
		params["after_time"] = opts.AfterTime
	}
	if opts.BeforeTime > 0 {
		params["before_time"] = opts.BeforeTime
	}

	result, err := c.request(ctx, EndpointDefiV3TokenMintBurnTxs, requestOptions{
		method:          "GET",
		chains:          opts.Chains,
		onLimitExceeded: RateLimitBehavior(opts.OnLimitExceeded),
		paramsOrBody:    params,
		respJustItems:   true,
	})
	if err != nil {
		return nil, err
	}

	var txs map[string]RespTokenMintBurnTxs
	data, _ := json.Marshal(result)
	if err := json.Unmarshal(data, &txs); err != nil {
		return nil, err
	}

	return txs["items"], nil
}

// ============================================================================
// API Methods - Token Exit Liquidity
// ============================================================================

// TokenExitLiquidityOptions holds options for GetTokenExitLiquidity and GetMultiTokenExitLiquidity. Note: Base chain only.
type TokenExitLiquidityOptions struct {
	// Chains is the list of blockchain networks to query. Default: nil
	Chains []Chain
	// OnLimitExceeded overrides the default rate limit behavior. Default: "" (use client default)
	OnLimitExceeded string `default:""`
}

// GetTokenExitLiquidity retrieves exit liquidity information for a token.
//
// This method retrieves information about exit liquidity for a token, including liquidity removal
// events and related data.
//
// Note: This endpoint is only available for Base chain.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - address: Token contract address to query
//   - opts: Configuration options (see TokenExitLiquidityOptions for details)
//
// Returns:
//   - *RespTokenExitLiquidity: Exit liquidity response
//   - error: Error if request fails or address is invalid
//
// Raises:
//   - Context errors: If ctx is cancelled or times out
//   - API errors: If the Birdeye API returns an error
func (c *HTTPClient) GetTokenExitLiquidity(ctx context.Context, address string, opts *TokenExitLiquidityOptions) (*RespTokenExitLiquidity, error) {
	if opts == nil {
		opts = &TokenExitLiquidityOptions{}
	}
	params := map[string]any{
		"address": address,
	}

	result, err := c.request(ctx, EndpointDefiV3TokenExitLiquidity, requestOptions{
		method:          "GET",
		chains:          opts.Chains,
		onLimitExceeded: RateLimitBehavior(opts.OnLimitExceeded),
		paramsOrBody:    params,
	})
	if err != nil {
		return nil, err
	}

	var exitLiquidity RespTokenExitLiquidity
	data, _ := json.Marshal(result)
	if err := json.Unmarshal(data, &exitLiquidity); err != nil {
		return nil, err
	}

	return &exitLiquidity, nil
}

// GetMultiTokenExitLiquidity retrieves exit liquidity information for multiple tokens.
//
// This method retrieves exit liquidity information for multiple tokens in a single request,
// improving efficiency for batch operations.
//
// Note: This endpoint is only available for Base chain.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - addresses: List of token contract addresses to query
//   - opts: Configuration options (see TokenExitLiquidityOptions for details)
//
// Returns:
//   - []RespTokenExitLiquidity: Exit liquidity response
//   - error: Error if request fails or any address is invalid
//
// Raises:
//   - Context errors: If ctx is cancelled or times out
//   - API errors: If the Birdeye API returns an error
func (c *HTTPClient) GetMultiTokenExitLiquidity(ctx context.Context, addresses []string, opts *TokenExitLiquidityOptions) ([]RespTokenExitLiquidity, error) {
	if opts == nil {
		opts = &TokenExitLiquidityOptions{}
	}
	params := map[string]any{
		"list_address": addresses,
	}

	result, err := c.request(ctx, EndpointDefiV3TokenExitLiquidityMultiple, requestOptions{
		method:          "GET",
		chains:          opts.Chains,
		onLimitExceeded: RateLimitBehavior(opts.OnLimitExceeded),
		paramsOrBody:    params,
	})
	if err != nil {
		return nil, err
	}

	var exitLiquidity []RespTokenExitLiquidity
	data, _ := json.Marshal(result)
	if err := json.Unmarshal(data, &exitLiquidity); err != nil {
		return nil, err
	}

	return exitLiquidity, nil
}

// ============================================================================
// API Methods - Meme Tokens
// ============================================================================

// MemeListOptions holds options for GetMemeList. Note: Solana only.
type MemeListOptions struct {
	// SortBy specifies the field to sort by. Default: "progress_percent"
	SortBy string `default:"liquidity"`
	// SortType specifies the sort order. Options: "desc", "asc". Default: "desc"
	SortType string `default:"desc"`
	// Source specifies the platform source. Options: "all", "pump_dot_fun". Default: "all"
	Source string `default:"all"`
	// Creator filters by creator address. Default: "" (no filter)
	Creator string `default:""`
	// PlatformID filters by platform ID. Default: "" (no filter)
	PlatformID string `default:""`
	// Graduated filters by graduation status. Default: false (no filter)
	Graduated bool `default:"false"`
	// Offset is the number of items to skip for pagination. Default: 0
	Offset int64 `default:"0"`
	// Limit is the maximum number of items to return. Default: 100
	Limit int64 `default:"100"`
	// Chains is the list of blockchain networks to query. Default: nil
	Chains []Chain
	// OnLimitExceeded overrides the default rate limit behavior. Default: "" (use client default)
	OnLimitExceeded string `default:""`
}

// GetMemeList retrieves list of meme tokens.
//
// This method retrieves a list of meme tokens with their basic information and market data.
//
// Note: This endpoint is only available for Solana chain.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - opts: Configuration options (see MemeListOptions for details)
//
// Returns:
//   - *RespMemeList: Meme list response containing list of meme tokens
//   - error: Error if request fails
//
// Raises:
//   - Context errors: If ctx is cancelled or times out
//   - API errors: If the Birdeye API returns an error
func (c *HTTPClient) GetMemeList(ctx context.Context, opts *MemeListOptions) (*RespMemeList, error) {
	if opts == nil {
		opts = &MemeListOptions{}
	}
	params, err := ApplyDefaultsAndBuildParams(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to apply defaults: %w", err)
	}

	// No additional parameters needed - all handled by ApplyDefaultsAndBuildParams

	if opts.Creator != "" {
		params["creator"] = opts.Creator
	}
	if opts.PlatformID != "" {
		params["platform_id"] = opts.PlatformID
	}
	if opts.Graduated {
		params["graduated"] = "true"
	}

	result, err := c.request(ctx, EndpointDefiV3TokenMemeList, requestOptions{
		method:          "GET",
		chains:          opts.Chains,
		onLimitExceeded: RateLimitBehavior(opts.OnLimitExceeded),
		paramsOrBody:    params,
	})
	if err != nil {
		return nil, err
	}

	var memeList RespMemeList
	data, _ := json.Marshal(result)
	if err := json.Unmarshal(data, &memeList); err != nil {
		return nil, err
	}

	return &memeList, nil
}

// MemeDetailOptions holds options for GetMemeDetail. Note: Solana only.
type MemeDetailOptions struct {
	// Chains is the list of blockchain networks to query. Default: nil
	Chains []Chain
	// OnLimitExceeded overrides the default rate limit behavior. Default: "" (use client default)
	OnLimitExceeded string `default:""`
}

// GetMemeDetail retrieves detailed information for a meme token.
//
// This method retrieves comprehensive information about a specific meme token including metadata,
// market data, and social metrics.
//
// Note: This endpoint is only available for Solana chain.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - address: Meme token contract address to query
//   - opts: Configuration options (see MemeDetailOptions for details)
//
// Returns:
//   - *RespMemeDetail: Meme detail response
//   - error: Error if request fails or address is invalid
//
// Raises:
//   - Context errors: If ctx is cancelled or times out
//   - API errors: If the Birdeye API returns an error
func (c *HTTPClient) GetMemeDetail(ctx context.Context, address string, opts *MemeDetailOptions) (*RespMemeDetail, error) {
	if opts == nil {
		opts = &MemeDetailOptions{}
	}
	params := map[string]any{
		"address": address,
	}

	result, err := c.request(ctx, EndpointDefiV3TokenMemeDetailSingle, requestOptions{
		method:          "GET",
		chains:          opts.Chains,
		onLimitExceeded: RateLimitBehavior(opts.OnLimitExceeded),
		paramsOrBody:    params,
	})
	if err != nil {
		return nil, err
	}

	var detail RespMemeDetail
	data, _ := json.Marshal(result)
	if err := json.Unmarshal(data, &detail); err != nil {
		return nil, err
	}

	return &detail, nil
}

// ============================================================================
// API Methods - Wallet PnL
// ============================================================================

// WalletTokensPnLOptions holds options for GetWalletTokensPnL and GetWalletsPnLByToken. Note: Solana only.
type WalletTokensPnLOptions struct {
	// Chains is the list of blockchain networks to query. Default: nil
	Chains []Chain
	// OnLimitExceeded overrides the default rate limit behavior. Default: "" (use client default)
	OnLimitExceeded string `default:""`
}

// GetWalletTokensPnL retrieves profit and loss for wallet tokens.
//
// This method retrieves profit and loss information for all tokens in a specific wallet address.
//
// Note: This endpoint is only available for Solana chain.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - wallet: Wallet address to query
//   - tokenAddresses: List of token addresses
//   - opts: Configuration options (see WalletTokensPnLOptions for details)
//
// Returns:
//   - *RespWalletTokensPnL: Wallet tokens PnL response
//   - error: Error if request fails or wallet address is invalid
//
// Raises:
//   - Context errors: If ctx is cancelled or times out
//   - API errors: If the Birdeye API returns an error
func (c *HTTPClient) GetWalletTokensPnL(ctx context.Context, wallet string, tokenAddresses []string, opts *WalletTokensPnLOptions) (*RespWalletTokensPnL, error) {
	if opts == nil {
		opts = &WalletTokensPnLOptions{}
	}
	params := map[string]any{
		"wallet":          wallet,
		"token_addresses": tokenAddresses,
	}

	result, err := c.request(ctx, EndpointV2WalletPnl, requestOptions{
		method:          "GET",
		chains:          opts.Chains,
		onLimitExceeded: RateLimitBehavior(opts.OnLimitExceeded),
		paramsOrBody:    params,
	})
	if err != nil {
		return nil, err
	}

	var pnl RespWalletTokensPnL
	data, _ := json.Marshal(result)
	if err := json.Unmarshal(data, &pnl); err != nil {
		return nil, err
	}

	return &pnl, nil
}

// GetWalletsPnLByToken retrieves profit and loss for wallets by token.
//
// This method retrieves profit and loss information for all wallets that hold a specific token.
//
// Note: This endpoint is only available for Solana chain.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - tokenAddress: Token contract address to query
//   - wallets: List of wallet addresses
//   - opts: Configuration options (see WalletTokensPnLOptions for details)
//
// Returns:
//   - *RespWalletsPnLByToken: Wallets PnL by token response
//   - error: Error if request fails or token address is invalid
//
// Raises:
//   - Context errors: If ctx is cancelled or times out
//   - API errors: If the Birdeye API returns an error
func (c *HTTPClient) GetWalletsPnLByToken(ctx context.Context, tokenAddress string, wallets []string, opts *WalletTokensPnLOptions) (*RespWalletsPnLByToken, error) {
	if opts == nil {
		opts = &WalletTokensPnLOptions{}
	}
	params := map[string]any{
		"token_address": tokenAddress,
		"wallets":       wallets,
	}

	result, err := c.request(ctx, EndpointV2WalletPnlMultiple, requestOptions{
		method:          "GET",
		chains:          opts.Chains,
		onLimitExceeded: RateLimitBehavior(opts.OnLimitExceeded),
		paramsOrBody:    params,
	})
	if err != nil {
		return nil, err
	}

	var pnl RespWalletsPnLByToken
	data, _ := json.Marshal(result)
	if err := json.Unmarshal(data, &pnl); err != nil {
		return nil, err
	}

	return &pnl, nil
}

// ============================================================================
// API Methods - Wallet Token Balance
// ============================================================================

// WalletTokensBalanceOptions holds options for GetWalletTokensBalance and GetWalletTokenFirstTx.
type WalletTokensBalanceOptions struct {
	// Chains is the list of blockchain networks to query. Default: nil
	Chains []Chain
	// OnLimitExceeded overrides the default rate limit behavior. Default: "" (use client default)
	OnLimitExceeded string `default:""`
}

// GetWalletTokensBalance retrieves token balances for a wallet.
//
// This method retrieves the current token balances for a specific wallet address across
// all supported networks.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - wallet: Wallet address to query
//   - tokenAddresses: List of token addresses to query balances for
//   - opts: Configuration options (see WalletTokensBalanceOptions for details)
//
// Returns:
//   - []RespWalletTokensBalanceItem: Wallet tokens balances response
//   - error: Error if request fails or wallet address is invalid
//
// Raises:
//   - Context errors: If ctx is cancelled or times out
//   - API errors: If the Birdeye API returns an error
func (c *HTTPClient) GetWalletTokensBalance(ctx context.Context, wallet string, tokenAddresses []string, opts *WalletTokensBalanceOptions) (RespWalletTokensBalances, error) {
	if opts == nil {
		opts = &WalletTokensBalanceOptions{}
	}
	result, err := c.request(ctx, EndpointV2WalletTokenBalance, requestOptions{
		method:          "POST",
		chains:          opts.Chains,
		onLimitExceeded: RateLimitBehavior(opts.OnLimitExceeded),
		paramsOrBody: map[string]any{
			"wallet":          wallet,
			"token_addresses": tokenAddresses,
		},
		paramsUseArray: true,
	})
	if err != nil {
		return nil, err
	}

	var balances map[string]RespWalletTokensBalances
	data, _ := json.Marshal(result)
	if err := json.Unmarshal(data, &balances); err != nil {
		return nil, err
	}

	return balances["items"], nil
}

// GetWalletTokenFirstTx retrieves first transaction for a wallet and token.
//
// This method retrieves the first transaction between a specific wallet and token, useful for
// tracking initial interactions.
//
// Note: This endpoint is only available for Solana chain.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - wallets: List of wallet addresses to query
//   - tokenAddress: Token contract address to query
//   - opts: Configuration options (see WalletTokensBalanceOptions for details)
//
// Returns:
//   - map[string]RespWalletTokenFirstTx: First transaction response
//   - error: Error if request fails or wallet/token address is invalid
//
// Raises:
//   - Context errors: If ctx is cancelled or times out
//   - API errors: If the Birdeye API returns an error
func (c *HTTPClient) GetWalletTokenFirstTx(ctx context.Context, wallets []string, tokenAddress string, opts *WalletTokensBalanceOptions) (map[string]RespWalletTokenFirstTx, error) {
	if opts == nil {
		opts = &WalletTokensBalanceOptions{}
	}
	result, err := c.request(ctx, EndpointV2WalletTxFirstFunded, requestOptions{
		method:          "POST",
		chains:          opts.Chains,
		onLimitExceeded: RateLimitBehavior(opts.OnLimitExceeded),
		paramsOrBody: map[string]any{
			"wallets":       wallets,
			"token_address": tokenAddress,
		},
		paramsUseArray: true,
	})
	if err != nil {
		return nil, err
	}

	var firstTx map[string]RespWalletTokenFirstTx
	data, _ := json.Marshal(result)
	if err := json.Unmarshal(data, &firstTx); err != nil {
		return nil, err
	}

	return firstTx, nil
}

// ============================================================================
// API Methods - Wallet Net Worth Details
// ============================================================================

// WalletNetWorthDetailsOptions holds options for GetWalletNetWorthDetails. Note: Solana only.
type WalletNetWorthDetailsOptions struct {
	// Time is the reference time for the query. Default: "" (current time)
	Time string `default:""`
	// Type specifies the time interval. Options: "1d", "1w", "1m". Default: "1d"
	Type string `default:"1d"`
	// SortType specifies the sort order. Options: "desc", "asc". Default: "desc"
	SortType string `default:"desc"`
	// Limit is the maximum number of items to return (1-100). Default: 20, max: 100
	Limit int64 `default:"20"`
	// Offset is the number of items to skip for pagination (0-10000). Default: 0, max: 10000
	Offset int64 `default:"0"`
	// Chains is the list of blockchain networks to query. Default: nil
	Chains []Chain
	// OnLimitExceeded overrides the default rate limit behavior. Default: "" (use client default)
	OnLimitExceeded string `default:""`
}

// GetWalletNetWorthDetails retrieves detailed net worth information for a wallet.
//
// This method retrieves comprehensive net worth details for a specific wallet address including
// asset breakdown and performance metrics.
//
// Note: This endpoint is only available for Solana chain.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - wallet: Wallet address to query
//   - opts: Configuration options (see WalletNetWorthDetailsOptions for details)
//
// Returns:
//   - *RespWalletNetWorthDetails: Wallet net worth details response
//   - error: Error if request fails or wallet address is invalid
//
// Raises:
//   - Context errors: If ctx is cancelled or times out
//   - API errors: If the Birdeye API returns an error
//   - Validation errors: If wallet address is invalid or limit/offset out of range
func (c *HTTPClient) GetWalletNetWorthDetails(ctx context.Context, wallet string, opts *WalletNetWorthDetailsOptions) (*RespWalletNetWorthDetails, error) {
	if opts == nil {
		opts = &WalletNetWorthDetailsOptions{}
	}
	params, err := ApplyDefaultsAndBuildParams(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to apply defaults: %w", err)
	}

	if opts.Limit < 1 || opts.Limit > 100 {
		return nil, errors.New("limit must be between 1 and 100")
	}
	if opts.Offset < 0 || opts.Offset > 10000 {
		return nil, errors.New("offset must be between 0 and 10000")
	}

	// Add required parameters
	params["wallet"] = wallet

	if opts.Time != "" {
		params["time"] = opts.Time
	}

	result, err := c.request(ctx, EndpointV2WalletNetWorthDetails, requestOptions{
		method:          "GET",
		chains:          opts.Chains,
		onLimitExceeded: RateLimitBehavior(opts.OnLimitExceeded),
		paramsOrBody:    params,
	})
	if err != nil {
		return nil, err
	}

	var details RespWalletNetWorthDetails
	data, _ := json.Marshal(result)
	if err := json.Unmarshal(data, &details); err != nil {
		return nil, err
	}

	return &details, nil
}

// ============================================================================
// API Methods - Token Holder Batch
// ============================================================================

// TokenHolderBatchOptions holds options for GetTokenHolderBatch. Note: Solana only.
type TokenHolderBatchOptions struct {
	// UIAmountMode specifies the token amount display mode. Options: "raw", "scaled", "both". Default: "scaled"
	UIAmountMode string `default:"raw"`
	// Chains is the list of blockchain networks to query. Default: nil
	Chains []Chain
	// OnLimitExceeded overrides the default rate limit behavior. Default: "" (use client default)
	OnLimitExceeded string `default:""`
}

// GetTokenHolderBatch retrieves token holder information for multiple wallets.
//
// This method retrieves holder information for multiple wallets in a single request,
// improving efficiency for batch operations.
//
// Note: This endpoint is only available for Solana chain.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - tokenAddress: Token contract address
//   - wallets: List of wallet addresses to check
//   - opts: Configuration options (see TokenHolderBatchOptions for details)
//
// Returns:
//   - []RespTokenHolderBatchItem: Token holder batch response
//   - error: Error if request fails or addresses are invalid
//
// Raises:
//   - Context errors: If ctx is cancelled or times out
//   - API errors: If the Birdeye API returns an error
func (c *HTTPClient) GetTokenHolderBatch(ctx context.Context, tokenAddress string, wallets []string, opts *TokenHolderBatchOptions) (RespTokenHolderBatch, error) {
	if opts == nil {
		opts = &TokenHolderBatchOptions{}
	}
	params, err := ApplyDefaultsAndBuildParams(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to apply defaults: %w", err)
	}

	// Add required parameters
	params["token_address"] = tokenAddress
	params["wallets"] = wallets

	result, err := c.request(ctx, EndpointTokenV1HolderBatch, requestOptions{
		method:          "POST",
		chains:          opts.Chains,
		onLimitExceeded: RateLimitBehavior(opts.OnLimitExceeded),
		paramsOrBody:    params,
		paramsUseArray:  true,
	})
	if err != nil {
		return nil, err
	}

	var holders map[string]RespTokenHolderBatch
	data, _ := json.Marshal(result)
	if err := json.Unmarshal(data, &holders); err != nil {
		return nil, err
	}

	return holders["items"], nil
}

// ============================================================================
// API Methods - Token List V1
// ============================================================================

// TokenListV1Options holds options for GetTokenListV1.
type TokenListV1Options struct {
	// SortBy specifies the field to sort by. Default: "liquidity"
	SortBy string `default:"liquidity"`
	// SortType specifies the sort order. Options: "desc", "asc". Default: "desc"
	SortType string `default:"desc"`
	// Offset is the number of tokens to skip for pagination. Default: 0
	Offset int64 `default:"0"`
	// Limit is the maximum number of tokens to return (1-50). Default: 50, max: 50
	Limit int64 `default:"50"`
	// MinLiquidity is the minimum liquidity filter in USD. Default: 100
	MinLiquidity float64
	// MaxLiquidity is the maximum liquidity filter in USD. Default: 0 (no filter)
	MaxLiquidity float64 `default:"0"`
	// UIAmountMode specifies the token amount display mode. Options: "raw", "scaled", "both". Default: "raw"
	UIAmountMode string `default:"raw"`
	// Chains is the list of blockchain networks to query. Default: nil
	Chains []Chain
	// OnLimitExceeded overrides the default rate limit behavior. Default: "" (use client default)
	OnLimitExceeded string `default:""`
}

// GetTokenListV1 retrieves token list using V1 API with basic filtering.
//
// This method retrieves a list of tokens with basic information using the V1 API endpoint.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - opts: Configuration options (see TokenListV1Options for details)
//
// Returns:
//   - *RespTokenListV1: Token list response containing list of tokens and pagination info
//   - error: Error if request fails or validation fails
//
// Raises:
//   - Context errors: If ctx is cancelled or times out
//   - API errors: If the Birdeye API returns an error
//   - Validation errors: If limit out of range (1-50)
func (c *HTTPClient) GetTokenListV1(ctx context.Context, opts *TokenListV1Options) (*RespTokenListV1, error) {
	if opts == nil {
		opts = &TokenListV1Options{}
	}
	params, err := ApplyDefaultsAndBuildParams(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to apply defaults: %w", err)
	}

	if opts.Limit < 1 || opts.Limit > 50 {
		return nil, errors.New("limit must be between 1 and 50")
	}

	// No additional parameters needed - all handled by ApplyDefaultsAndBuildParams

	if opts.MaxLiquidity > 0 {
		params["max_liquidity"] = opts.MaxLiquidity
	}

	result, err := c.request(ctx, EndpointDefiTokenList, requestOptions{
		method:          "GET",
		chains:          opts.Chains,
		onLimitExceeded: RateLimitBehavior(opts.OnLimitExceeded),
		paramsOrBody:    params,
	})
	if err != nil {
		return nil, err
	}

	var tokenList RespTokenListV1
	data, _ := json.Marshal(result)
	if err := json.Unmarshal(data, &tokenList); err != nil {
		return nil, err
	}

	return &tokenList, nil
}

// ============================================================================
// API Methods - All Transactions V3
// ============================================================================

// AllTxsV3Options holds options for GetAllTxs.
type AllTxsV3Options struct {
	// Offset is the number of transactions to skip for pagination. Default: 0
	Offset int64 `default:"0"`
	// Limit is the maximum number of transactions to return (1-100). Default: 100, max: 100
	Limit int64 `default:"100"`
	// SortBy specifies the field to sort by. Options: "block_unix_time", "block_number". Default: "block_unix_time"
	SortBy string `default:"block_unix_time"`
	// SortType specifies the sort order. Options: "desc", "asc". Default: "desc"
	SortType string `default:"desc"`
	// TxType specifies the transaction type filter. Options: "swap", "add", "remove", "all". Default: "swap"
	TxType string `default:"swap"`
	// Source filters by DEX source. Default: "" (no filter)
	Source string `default:""`
	// Owner filters by owner address. Default: "" (no filter)
	Owner string `default:""`
	// PoolID filters by pool ID. Default: "" (no filter)
	PoolID string `default:""`
	// BeforeTime filters transactions before this Unix timestamp. Default: 0 (no filter)
	BeforeTime int64 `default:"0"`
	// AfterTime filters transactions after this Unix timestamp. Default: 0 (no filter)
	AfterTime int64 `default:"0"`
	// BeforeBlockNumber filters transactions before this block number. Default: 0 (no filter)
	BeforeBlockNumber int64 `default:"0"`
	// AfterBlockNumber filters transactions after this block number. Default: 0 (no filter)
	AfterBlockNumber int64 `default:"0"`
	// UIAmountMode specifies the token amount display mode. Options: "raw", "scaled", "both". Default: "scaled"
	UIAmountMode string `default:"scaled"`
	// Chains is the list of blockchain networks to query. Default: nil
	Chains []Chain
	// OnLimitExceeded overrides the default rate limit behavior. Default: "" (use client default)
	OnLimitExceeded string `default:""`
}

// GetAllTxs retrieves all transactions across the platform with advanced filtering.
//
// This method retrieves transactions from across all supported networks with comprehensive
// filtering options for pairs, tokens, wallets, and time ranges.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - opts: Configuration options (see AllTxsV3Options for details)
//
// Returns:
//   - *RespAllTxsV3: Transaction response containing items and hasNext flag
//   - error: Error if request fails or validation fails
//
// Raises:
//   - Context errors: If ctx is cancelled or times out
//   - API errors: If the Birdeye API returns an error
//   - Validation errors: If offset+limit > 10000 or limit > 100
func (c *HTTPClient) GetAllTxs(ctx context.Context, opts *AllTxsV3Options) (*RespAllTxsV3, error) {
	if opts == nil {
		opts = &AllTxsV3Options{}
	}
	params, err := ApplyDefaultsAndBuildParams(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to apply defaults: %w", err)
	}

	if opts.Offset+opts.Limit > 10000 {
		return nil, errors.New("offset + limit must be <= 10000")
	}
	if opts.Limit > 100 {
		return nil, errors.New("limit must be <= 100")
	}

	// No additional parameters needed - all handled by ApplyDefaultsAndBuildParams

	if opts.Source != "" {
		params["source"] = opts.Source
	}
	if opts.Owner != "" {
		params["owner"] = opts.Owner
	}
	if opts.PoolID != "" {
		params["pool_id"] = opts.PoolID
	}
	if opts.BeforeTime > 0 {
		params["before_time"] = opts.BeforeTime
	}
	if opts.AfterTime > 0 {
		params["after_time"] = opts.AfterTime
	}
	if opts.BeforeBlockNumber > 0 {
		params["before_block_number"] = opts.BeforeBlockNumber
	}
	if opts.AfterBlockNumber > 0 {
		params["after_block_number"] = opts.AfterBlockNumber
	}

	result, err := c.request(ctx, EndpointDefiV3Txs, requestOptions{
		method:          "GET",
		chains:          opts.Chains,
		onLimitExceeded: RateLimitBehavior(opts.OnLimitExceeded),
		paramsOrBody:    params,
	})
	if err != nil {
		return nil, err
	}

	var txs RespAllTxsV3
	data, _ := json.Marshal(result)
	if err := json.Unmarshal(data, &txs); err != nil {
		return nil, err
	}

	return &txs, nil
}

// RecentTxsV3Options holds options for GetRecentTxs.
type RecentTxsV3Options struct {
	// Offset is the number of transactions to skip for pagination (0-9999). Default: 0, max: 9999
	Offset int64 `default:"0"`
	// Limit is the maximum number of transactions to return (1-100). Default: 100, max: 100
	Limit int64 `default:"100"`
	// TxType specifies the transaction type filter. Options: "swap", "add", "remove", "all". Default: "swap"
	TxType string `default:"swap"`
	// Owner filters by owner address. Default: "" (no filter)
	Owner string `default:""`
	// BeforeTime filters transactions before this Unix timestamp. Default: 0 (no filter)
	BeforeTime int64 `default:"0"`
	// AfterTime filters transactions after this Unix timestamp. Default: 0 (no filter)
	AfterTime int64 `default:"0"`
	// BeforeBlockNumber filters transactions before this block number. Default: 0 (no filter)
	BeforeBlockNumber int64 `default:"0"`
	// AfterBlockNumber filters transactions after this block number. Default: 0 (no filter)
	AfterBlockNumber int64 `default:"0"`
	// UIAmountMode specifies the token amount display mode. Options: "raw", "scaled", "both". Default: "raw"
	UIAmountMode string `default:"raw"`
	// Chains is the list of blockchain networks to query. Default: nil
	Chains []Chain
	// OnLimitExceeded overrides the default rate limit behavior. Default: "" (use client default)
	OnLimitExceeded string `default:""`
}

// GetRecentTxs retrieves recent transactions across the platform with advanced filtering.
//
// This method retrieves the most recent transactions across all supported networks with
// comprehensive filtering options and sorting capabilities.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - opts: Configuration options (see RecentTxsV3Options for details)
//
// Returns:
//   - *RespRecentTxsV3: Transaction response containing recent transaction records
//   - error: Error if request fails or validation fails
//
// Raises:
//   - Context errors: If ctx is cancelled or times out
//   - API errors: If the Birdeye API returns an error
//   - Validation errors: If offset/limit out of range or offset+limit > 10000
func (c *HTTPClient) GetRecentTxs(ctx context.Context, opts *RecentTxsV3Options) (*RespRecentTxsV3, error) {
	if opts == nil {
		opts = &RecentTxsV3Options{}
	}
	params, err := ApplyDefaultsAndBuildParams(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to apply defaults: %w", err)
	}

	if opts.Offset < 0 || opts.Offset > 9999 {
		return nil, errors.New("offset must be between 0 and 9999")
	}
	if opts.Limit < 1 || opts.Limit > 100 {
		return nil, errors.New("limit must be between 1 and 100")
	}
	if opts.Offset+opts.Limit > 10000 {
		return nil, errors.New("offset + limit cannot exceed 10000")
	}

	// No additional parameters needed - all handled by ApplyDefaultsAndBuildParams

	if opts.Owner != "" {
		params["owner"] = opts.Owner
	}
	if opts.BeforeTime > 0 {
		params["before_time"] = opts.BeforeTime
	}
	if opts.AfterTime > 0 {
		params["after_time"] = opts.AfterTime
	}
	if opts.BeforeBlockNumber > 0 {
		params["before_block_number"] = opts.BeforeBlockNumber
	}
	if opts.AfterBlockNumber > 0 {
		params["after_block_number"] = opts.AfterBlockNumber
	}

	result, err := c.request(ctx, EndpointDefiV3Txs, requestOptions{
		method:          "GET",
		chains:          opts.Chains,
		onLimitExceeded: RateLimitBehavior(opts.OnLimitExceeded),
		paramsOrBody:    params,
	})
	if err != nil {
		return nil, err
	}

	var txs RespRecentTxsV3
	data, _ := json.Marshal(result)
	if err := json.Unmarshal(data, &txs); err != nil {
		return nil, err
	}

	return &txs, nil
}

// ============================================================================
// API Methods - OHLCV Base Quote
// ============================================================================

// OHLCVBaseQuoteOptions holds options for GetOHLCVBaseQuote.
type OHLCVBaseQuoteOptions struct {
	// UIAmountMode specifies the token amount display mode. Options: "raw", "scaled", "both". Default: "raw"
	UIAmountMode string `default:"raw"`
	// Chains is the list of blockchain networks to query. Default: nil
	Chains []Chain
	// OnLimitExceeded overrides the default rate limit behavior. Default: "" (use client default)
	OnLimitExceeded string `default:""`
}

// GetOHLCVBaseQuote retrieves OHLCV data for a trading pair by base and quote token addresses.
//
// This method retrieves OHLCV data for a trading pair by specifying the base and quote token
// addresses, which is useful when you don't have the pair address.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - baseAddress: Base token contract address
//   - quoteAddress: Quote token contract address
//   - intervalType: Time interval. Options: "1m", "3m", "5m", "15m", "30m", "1H", "2H", "4H", "6H", "8H", "12H", "1D", "3D", "1W", "1M"
//   - timeFrom: Start time in Unix timestamp (seconds)
//   - timeTo: End time in Unix timestamp (seconds)
//   - opts: Configuration options (see OHLCVBaseQuoteOptions for details)
//
// Returns:
//   - *RespOHLCVBaseQuote: OHLCV response containing list of OHLCV data points
//   - error: Error if request fails, addresses are invalid, or time range is invalid
//
// Raises:
//   - Context errors: If ctx is cancelled or times out
//   - API errors: If the Birdeye API returns an error
//   - Validation errors: If addresses are invalid or time range is invalid
func (c *HTTPClient) GetOHLCVBaseQuote(ctx context.Context, baseAddress, quoteAddress, intervalType string, timeFrom, timeTo int64, opts *OHLCVBaseQuoteOptions) (*RespOHLCVBaseQuote, error) {
	if opts == nil {
		opts = &OHLCVBaseQuoteOptions{}
	}
	params, err := ApplyDefaultsAndBuildParams(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to apply defaults: %w", err)
	}

	if timeFrom < 0 || timeFrom > 10000000000 {
		return nil, errors.New("time_from must be between 0 and 10000000000")
	}
	if timeTo < 0 || timeTo > 10000000000 {
		return nil, errors.New("time_to must be between 0 and 10000000000")
	}

	// Add required parameters
	params["base_address"] = baseAddress
	params["quote_address"] = quoteAddress
	params["type"] = intervalType
	params["time_from"] = timeFrom
	params["time_to"] = timeTo

	result, err := c.request(ctx, EndpointDefiOHLCVBaseQuote, requestOptions{
		method:          "GET",
		chains:          opts.Chains,
		onLimitExceeded: RateLimitBehavior(opts.OnLimitExceeded),
		paramsOrBody:    params,
	})
	if err != nil {
		return nil, err
	}

	var ohlcv RespOHLCVBaseQuote
	data, _ := json.Marshal(result)
	if err := json.Unmarshal(data, &ohlcv); err != nil {
		return nil, err
	}

	return &ohlcv, nil
}

// ============================================================================
// API Methods - Token List V3 Scroll
// ============================================================================

// TokenListV3ScrollOptions holds options for GetTokenListV3Scroll.
//
// This endpoint uses scroll pagination for efficient retrieval of large token lists.
// Note: This endpoint is rate limited to 2 RPS (very strict).
type TokenListV3ScrollOptions struct {
	// SortBy specifies the field to sort by.
	// Options: "liquidity", "market_cap", "volume_24h", etc.
	// Optional, default: "liquidity"
	SortBy string `default:"liquidity"`

	// SortType specifies the sort order.
	// Options:
	//   - "desc": Highest first
	//   - "asc": Lowest first
	// Optional, default: "desc"
	SortType string `default:"desc"`

	// Limit is the maximum number of tokens to return per page (1-5000).
	// Higher limits allow fetching more data per request but use more quota.
	// Optional, default: 5000, max: 5000
	Limit int64 `default:"5000"`

	// ScrollID is the pagination cursor from the previous response.
	// Use the scroll_id from the previous response to get the next page.
	// For the first request, leave this nil to start from the beginning.
	// Optional, default: nil (start from beginning)
	ScrollID string `default:""`

	// MinLiquidity is the minimum liquidity filter in USD.
	// Only tokens with liquidity >= this value will be returned.
	// Optional, default: nil (no minimum)
	MinLiquidity float64 `default:"0"`

	// MaxLiquidity is the maximum liquidity filter in USD.
	// Only tokens with liquidity <= this value will be returned.
	// Optional, default: nil (no maximum)
	MaxLiquidity float64 `default:"0"`

	// UIAmountMode specifies the token amount display mode.
	// Options:
	//   - "raw": Raw amounts (e.g., 1000000000 for 1 token with 9 decimals)
	//   - "scaled": Human-readable amounts (e.g., 1.0)
	//   - "both": Include both raw and scaled amounts
	// Optional, default: "raw"
	UIAmountMode string `default:"raw"`

	// Chains is the list of blockchain networks to query.
	// If nil, queries all supported networks.
	// Optional, default: nil (all networks)
	Chains []Chain

	// OnLimitExceeded overrides the default rate limit behavior for this request.
	// If nil, uses the client's default behavior.
	// Note: This endpoint has a very strict 2 RPS limit.
	// Optional, default: nil (use client default)
	OnLimitExceeded string `default:""`
}

// GetTokenListV3Scroll retrieves token list using V3 API with Scroll network support.
//
// This method retrieves a list of tokens with advanced filtering options specifically optimized
// for the Scroll network. Uses scroll pagination for efficient retrieval of large token lists.
//
// Note: This endpoint is rate limited to 2 RPS (very strict).
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - opts: Configuration options (see TokenListV3ScrollOptions for details)
//
// Returns:
//   - *RespTokenListV3Scroll: Token list response containing:
//   - items: List of token information
//   - hasNext: Whether more tokens are available
//   - scrollID: Pagination cursor for next request
//   - error: Error if request fails or validation fails
//
// Raises:
//   - Context errors: If ctx is cancelled or times out
//   - API errors: If the Birdeye API returns an error
//   - Validation errors: If limit out of range (1-5000)
//   - Rate limit errors: If rate limit exceeded (strict 2 RPS limit)
func (c *HTTPClient) GetTokenListV3Scroll(ctx context.Context, opts *TokenListV3ScrollOptions) (*RespTokenListV3Scroll, error) {
	if opts == nil {
		opts = &TokenListV3ScrollOptions{}
	}
	params, err := ApplyDefaultsAndBuildParams(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to apply defaults: %w", err)
	}

	if opts.Limit < 1 || opts.Limit > 5000 {
		return nil, errors.New("limit must be between 1 and 5000")
	}

	// No additional parameters needed - all handled by ApplyDefaultsAndBuildParams

	if opts.ScrollID != "" {
		params["scroll_id"] = opts.ScrollID
	}
	if opts.MinLiquidity > 0 {
		params["min_liquidity"] = opts.MinLiquidity
	}
	if opts.MaxLiquidity > 0 {
		params["max_liquidity"] = opts.MaxLiquidity
	}

	result, err := c.request(ctx, EndpointDefiV3TokenListScroll, requestOptions{
		method:          "GET",
		chains:          opts.Chains,
		onLimitExceeded: RateLimitBehavior(opts.OnLimitExceeded),
		paramsOrBody:    params,
	})
	if err != nil {
		return nil, err
	}

	var tokenList RespTokenListV3Scroll
	data, _ := json.Marshal(result)
	if err := json.Unmarshal(data, &tokenList); err != nil {
		return nil, err
	}

	return &tokenList, nil
}

// ============================================================================
// End of API Methods
// ============================================================================

// ============================================================================
// API Methods - Wallet Balance Changes
// ============================================================================

// WalletBalanceChangesOptions holds options for GetWalletBalanceChanges.
//
// This endpoint retrieves the history of token balance changes for a specific wallet and token.
// Note: This endpoint is only available for Solana chain.
type WalletBalanceChangesOptions struct {
	// TimeFrom is the start time filter in Unix timestamp (seconds).
	// Get balance changes after this time.
	// Optional, default: nil (no lower time limit)
	TimeFrom int64 `default:"0"`

	// TimeTo is the end time filter in Unix timestamp (seconds).
	// Get balance changes before this time.
	// Optional, default: nil (no upper time limit)
	TimeTo int64 `default:"0"`

	// Type specifies the token type.
	// Options:
	//   - "SPL": SPL tokens (Solana Program Library tokens)
	//   - "NFT": NFT tokens
	// Optional, default: "SPL"
	Type string `default:"1d"`

	// ChangeType filters by the type of balance change.
	// Options:
	//   - "increase": Only show increases in balance
	//   - "decrease": Only show decreases in balance
	// Optional, default: nil (show all changes)
	ChangeType string `default:""`

	// Offset is the number of balance changes to skip for pagination.
	// Optional, default: 0
	Offset int64 `default:"0"`

	// Limit is the maximum number of balance changes to return (1-100).
	// Optional, default: 100, max: 100
	Limit int64 `default:"100"`

	// UIAmountMode specifies the token amount display mode.
	// Options:
	//   - "raw": Raw amounts (e.g., 1000000000 for 1 token with 9 decimals)
	//   - "scaled": Human-readable amounts (e.g., 1.0)
	//   - "both": Include both raw and scaled amounts
	// Optional, default: "raw"
	UIAmountMode string `default:"raw"`

	// Chains is the list of blockchain networks to query.
	// Note: Currently only Solana is supported for this endpoint.
	// Optional, default: nil (all networks)
	Chains []Chain

	// OnLimitExceeded overrides the default rate limit behavior for this request.
	// If nil, uses the client's default behavior.
	// Optional, default: nil (use client default)
	OnLimitExceeded string `default:""`
}

// GetWalletBalanceChanges retrieves balance changes for a wallet.
//
// This method retrieves balance changes for a specific wallet including token balance updates
// and transaction history.
//
// Note: This endpoint is only available for Solana chain.
//
// Parameters:
//   - ctx: Context for cancellation and timeout control
//   - wallet: Wallet address to query
//   - tokenAddress: Token contract address to query
//   - opts: Configuration options (see WalletBalanceChangesOptions for details)
//
// Returns:
//   - []RespWalletBalanceChangesItem: Wallet balance changes response containing:
//   - totalChanges: Total number of balance changes
//   - changes: List of balance changes
//   - Each change contains: token, oldBalance, newBalance, timestamp
//   - error: Error if request fails or wallet address is invalid
//
// Raises:
//   - Context errors: If ctx is cancelled or times out
//   - API errors: If the Birdeye API returns an error
//   - Validation errors: If wallet address is invalid
func (c *HTTPClient) GetWalletBalanceChanges(ctx context.Context, wallet, tokenAddress string, opts *WalletBalanceChangesOptions) (RespWalletBalanceChanges, error) {
	if opts == nil {
		opts = &WalletBalanceChangesOptions{}
	}
	params, err := ApplyDefaultsAndBuildParams(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to apply defaults: %w", err)
	}

	// Add required parameters
	params["wallet"] = wallet
	params["token_address"] = tokenAddress

	if opts.TimeFrom > 0 {
		params["time_from"] = opts.TimeFrom
	}
	if opts.TimeTo > 0 {
		params["time_to"] = opts.TimeTo
	}
	if opts.Type != "" {
		params["type"] = opts.Type
	}
	if opts.ChangeType != "" {
		params["change_type"] = opts.ChangeType
	}

	result, err := c.request(ctx, EndpointV1WalletTokenBalance, requestOptions{
		method:          "GET",
		chains:          opts.Chains,
		onLimitExceeded: RateLimitBehavior(opts.OnLimitExceeded),
		paramsOrBody:    params,
		respJustItems:   true,
	})
	if err != nil {
		return nil, err
	}

	var changes map[string]RespWalletBalanceChanges
	data, _ := json.Marshal(result)
	if err := json.Unmarshal(data, &changes); err != nil {
		return nil, err
	}

	return changes["items"], nil
}
