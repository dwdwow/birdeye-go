package birdeye

import (
	"context"
	"os"
	"testing"
	"time"
)

// Test constants
const (
	// Solana addresses for testing
	testTokenSOL    = "So11111111111111111111111111111111111111112"
	testTokenUSDC   = "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v"
	testTokenBONK   = "DezXAZ8z7PnrnRJjz3wXBoRgixCa6xjnB7YaB1pPB263"
	testWalletAddr  = "GBJ4MZe8fqpA6UVgjh19BwJPMb79KDfMv78XnFVxgH2Q"
	testPairAddress = "Czfq3xZZDmsdGdUyrNLtRhGc47cXcZtLG4crryfu44zE"
)

// getTestClient returns a test HTTP client
func getTestClient(t *testing.T) *HTTPClient {
	apiKey := os.Getenv("BIRDEYE_API_KEY")
	if apiKey == "" {
		t.Skip("BIRDEYE_API_KEY environment variable not set")
	}

	client := NewHTTPClient(HTTPClientConfig{
		APIKey: apiKey,
		Chains: []Chain{ChainSolana},
	})
	return client
}

// Test Network Support APIs
func TestGetSupportedNetworks(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	networks, err := client.GetSupportedNetworks(ctx)
	if err != nil {
		t.Fatalf("GetSupportedNetworks failed: %v", err)
	}

	if len(networks) == 0 {
		t.Error("Expected at least one supported network")
	}

	t.Logf("Supported networks: %d", len(networks))
}

func TestGetWalletSupportedNetworks(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	networks, err := client.GetWalletSupportedNetworks(ctx)
	if err != nil {
		t.Fatalf("GetWalletSupportedNetworks failed: %v", err)
	}

	if len(networks) == 0 {
		t.Error("Expected at least one wallet supported network")
	}

	t.Logf("Wallet supported networks: %d", len(networks))
}

// Test Token Price APIs
func TestGetTokenPrice(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	price, err := client.GetTokenPrice(ctx, testTokenSOL, nil)
	if err != nil {
		t.Fatalf("GetTokenPrice failed: %v", err)
	}

	if price.Value <= 0 {
		t.Errorf("Expected positive price, got %f", price.Value)
	}

	t.Logf("SOL Price: $%.2f", price.Value)
}

func TestGetMultiTokenPrice(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	addresses := []string{testTokenSOL, testTokenUSDC}
	prices, err := client.GetMultiTokenPrice(ctx, addresses, nil)
	if err != nil {
		t.Fatalf("GetMultiTokenPrice failed: %v", err)
	}

	if len(prices) == 0 {
		t.Error("Expected at least one price")
	}

	for addr, price := range prices {
		t.Logf("Token %s: $%.4f", addr[:8], price.Value)
	}
}

func TestGetTokenPriceVolume(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	data, err := client.GetTokenPriceVolume(ctx, testTokenSOL, nil)
	if err != nil {
		t.Fatalf("GetTokenPriceVolume failed: %v", err)
	}

	if data.Price <= 0 {
		t.Errorf("Expected positive price, got %f", data.Price)
	}

	t.Logf("SOL Price: $%.2f, Volume: $%.2f", data.Price, data.VolumeUSD)
}

func TestGetMultiTokenPriceVolume(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	addresses := []string{testTokenSOL, testTokenUSDC}
	volumes, err := client.GetMultiTokenPriceVolume(ctx, addresses, nil)
	if err != nil {
		t.Fatalf("GetMultiTokenPriceVolume failed: %v", err)
	}

	if len(volumes) == 0 {
		t.Error("Expected at least one volume data")
	}

	t.Logf("Retrieved volume data for %d tokens", len(volumes))
}

func TestGetTokenPriceHistories(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	now := time.Now().Unix()
	from := now - 3600 // 1 hour ago
	to := now

	histories, err := client.GetTokenPriceHistories(ctx, testTokenSOL, "token", "15m", from, to, nil)
	if err != nil {
		t.Fatalf("GetTokenPriceHistories failed: %v", err)
	}

	if len(histories.Items) == 0 {
		t.Error("Expected at least one price history item")
	}

	t.Logf("Retrieved %d price history items", len(histories.Items))
}

func TestGetTokenPriceHistoryByTime(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	unixTime := time.Now().Unix() - 3600 // 1 hour ago

	history, err := client.GetTokenPriceHistoryByTime(ctx, testTokenSOL, unixTime, nil)
	if err != nil {
		t.Fatalf("GetTokenPriceHistoryByTime failed: %v", err)
	}

	if history.Value <= 0 {
		t.Errorf("Expected positive price, got %f", history.Value)
	}

	t.Logf("SOL Price at %d: $%.2f", unixTime, history.Value)
}

// Test Token Metadata APIs
func TestGetTokenMetadata(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	metadata, err := client.GetTokenMetadata(ctx, testTokenSOL, nil)
	if err != nil {
		t.Fatalf("GetTokenMetadata failed: %v", err)
	}

	if metadata.Symbol == "" {
		t.Error("Expected non-empty symbol")
	}

	t.Logf("Token: %s (%s)", metadata.Name, metadata.Symbol)
}

func TestGetMultiTokenMetadata(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	addresses := []string{testTokenSOL, testTokenUSDC}
	metadatas, err := client.GetMultiTokenMetadata(ctx, addresses, nil)
	if err != nil {
		t.Fatalf("GetMultiTokenMetadata failed: %v", err)
	}

	if len(metadatas) == 0 {
		t.Error("Expected at least one metadata")
	}

	t.Logf("Retrieved metadata for %d tokens", len(metadatas))
}

// Test Token Market Data APIs
func TestGetTokenMarketData(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	marketData, err := client.GetTokenMarketData(ctx, testTokenSOL, nil)
	if err != nil {
		t.Fatalf("GetTokenMarketData failed: %v", err)
	}

	if marketData.Price <= 0 {
		t.Errorf("Expected positive price, got %f", marketData.Price)
	}

	t.Logf("Market Cap: $%.2f, Liquidity: $%.2f", marketData.MarketCap, marketData.Liquidity)
}

func TestGetMultiTokenMarketData(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	addresses := []string{testTokenSOL, testTokenUSDC}
	marketDatas, err := client.GetMultiTokenMarketData(ctx, addresses, nil)
	if err != nil {
		t.Fatalf("GetMultiTokenMarketData failed: %v", err)
	}

	if len(marketDatas) == 0 {
		t.Error("Expected at least one market data")
	}

	t.Logf("Retrieved market data for %d tokens", len(marketDatas))
}

// Test Token Trade Data APIs
func TestGetTokenTradeData(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	tradeData, err := client.GetTokenTradeData(ctx, testTokenSOL, nil)
	if err != nil {
		t.Fatalf("GetTokenTradeData failed: %v", err)
	}

	if tradeData.Price <= 0 {
		t.Errorf("Expected positive price, got %f", tradeData.Price)
	}

	t.Logf("Trade data - Price: $%.2f, Volume 24h: $%.2f", tradeData.Price, tradeData.Volume24hUSD)
}

func TestGetMultiTokenTradeData(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	addresses := []string{testTokenSOL, testTokenUSDC}
	tradeDatas, err := client.GetMultiTokenTradeData(ctx, addresses, nil)
	if err != nil {
		t.Fatalf("GetMultiTokenTradeData failed: %v", err)
	}

	if len(tradeDatas) == 0 {
		t.Error("Expected at least one trade data")
	}

	t.Logf("Retrieved trade data for %d tokens", len(tradeDatas))
}

// Test Transaction APIs
func TestGetTokenTxs(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	txs, err := client.GetTokenTxs(ctx, testTokenSOL, nil)
	if err != nil {
		t.Fatalf("GetTokenTxs failed: %v", err)
	}

	t.Logf("Retrieved %d transactions", len(txs.Items))
}

func TestGetTokenTxsByTime(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	now := time.Now().Unix()
	// from := now - 3600 // 1 hour ago

	txs, err := client.GetTokenTxsByTime(ctx, testTokenSOL, &TokenTxsByTimeOptions{
		// AfterTime:  from,
		BeforeTime: now,
	})
	if err != nil {
		t.Fatalf("GetTokenTxsByTime failed: %v", err)
	}

	t.Logf("Retrieved %d transactions", len(txs.Items))
}

func TestGetTokenTxsV3(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	txs, err := client.GetTokenTxsV3(ctx, testTokenSOL, &TokenTxsV3Options{
		Limit: 10,
	})
	if err != nil {
		t.Fatalf("GetTokenTxsV3 failed: %v", err)
	}

	t.Logf("Retrieved %d transactions (V3)", len(txs.Items))
}

func TestGetPairTxs(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	txs, err := client.GetPairTxs(ctx, testPairAddress, &PairTxsOptions{
		Limit: 10,
	})
	if err != nil {
		t.Fatalf("GetPairTxs failed: %v", err)
	}

	t.Logf("Retrieved %d pair transactions", len(txs.Items))
}

func TestGetPairTxsByTime(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	now := time.Now().Unix()
	from := now - 3600

	txs, err := client.GetPairTxsByTime(ctx, testPairAddress, &PairTxsByTimeOptions{
		AfterTime: from,
		// BeforeTime: now,
	})
	if err != nil {
		t.Fatalf("GetPairTxsByTime failed: %v", err)
	}

	t.Logf("Retrieved %d pair transactions", len(txs.Items))
}

// Test OHLCV APIs
func TestGetTokenOHLCV(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	now := time.Now().Unix()
	from := now - 3600

	ohlcv, err := client.GetTokenOHLCV(ctx, testTokenSOL, "15m", from, now, nil)
	if err != nil {
		t.Fatalf("GetTokenOHLCV failed: %v", err)
	}

	if len(ohlcv.Items) == 0 {
		t.Error("Expected at least one OHLCV item")
	}

	t.Logf("Retrieved %d OHLCV items", len(ohlcv.Items))
}

func TestGetTokenOHLCVV3(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	now := time.Now().Unix()
	from := now - 3600

	ohlcv, err := client.GetTokenOHLCVV3(ctx, testTokenSOL, "15m", from, now, nil)
	if err != nil {
		t.Fatalf("GetTokenOHLCVV3 failed: %v", err)
	}

	if len(ohlcv.Items) == 0 {
		t.Error("Expected at least one OHLCV item")
	}

	t.Logf("Retrieved %d OHLCV V3 items", len(ohlcv.Items))
}

func TestGetPairOHLCV(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	now := time.Now().Unix()
	from := now - 3600

	ohlcv, err := client.GetPairOHLCV(ctx, testPairAddress, "15m", from, now, nil)
	if err != nil {
		t.Fatalf("GetPairOHLCV failed: %v", err)
	}

	t.Logf("Retrieved %d pair OHLCV items", len(ohlcv))
}

func TestGetPairOHLCVV3(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	now := time.Now().Unix()
	from := now - 3600

	ohlcv, err := client.GetPairOHLCVV3(ctx, testPairAddress, "15m", from, now, nil)
	if err != nil {
		t.Fatalf("GetPairOHLCVV3 failed: %v", err)
	}

	t.Logf("Retrieved %d pair OHLCV V3 items", len(ohlcv))
}

func TestGetOHLCVBaseQuote(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	now := time.Now().Unix()
	from := now - 3600

	ohlcv, err := client.GetOHLCVBaseQuote(ctx, testTokenSOL, testTokenUSDC, "15m", from, now, nil)
	if err != nil {
		t.Fatalf("GetOHLCVBaseQuote failed: %v", err)
	}

	if len(ohlcv.Items) == 0 {
		t.Error("Expected at least one OHLCV item")
	}

	t.Logf("Retrieved %d base/quote OHLCV items", len(ohlcv.Items))
}

// Test Pair Overview APIs
func TestGetPairOverview(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	overview, err := client.GetPairOverview(ctx, testPairAddress, nil)
	if err != nil {
		t.Fatalf("GetPairOverview failed: %v", err)
	}

	if overview.Address == "" {
		t.Error("Expected non-empty address")
	}

	t.Logf("Pair: %s, Liquidity: $%.2f", overview.Name, overview.Liquidity)
}

func TestGetPairsOverview(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	addresses := []string{testPairAddress}
	overviews, err := client.GetPairsOverview(ctx, addresses, nil)
	if err != nil {
		t.Fatalf("GetPairsOverview failed: %v", err)
	}

	if len(overviews) == 0 {
		t.Error("Expected at least one pair overview")
	}

	t.Logf("Retrieved %d pair overviews", len(overviews))
}

// Test Token List APIs
func TestGetTokenListV1(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	tokenList, err := client.GetTokenListV1(ctx, &TokenListV1Options{
		Limit: 10,
	})
	if err != nil {
		t.Fatalf("GetTokenListV1 failed: %v", err)
	}

	if len(tokenList.Tokens) == 0 {
		t.Error("Expected at least one token")
	}

	t.Logf("Retrieved %d tokens, Total: %d", len(tokenList.Tokens), tokenList.Total)
}

func TestGetTokenListV3(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	tokenList, err := client.GetTokenListV3(ctx, &TokenListV3Options{
		Limit: 10,
	})
	if err != nil {
		t.Fatalf("GetTokenListV3 failed: %v", err)
	}

	if len(tokenList.Items) == 0 {
		t.Error("Expected at least one token")
	}

	t.Logf("Retrieved %d tokens (V3)", len(tokenList.Items))
}

func TestGetTokenListV3Scroll(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	tokenList, err := client.GetTokenListV3Scroll(ctx, &TokenListV3ScrollOptions{
		Limit: 100,
	})
	if err != nil {
		t.Fatalf("GetTokenListV3Scroll failed: %v", err)
	}

	if len(tokenList.Items) == 0 {
		t.Error("Expected at least one token")
	}

	t.Logf("Retrieved %d tokens (V3 Scroll), HasNext: %v", len(tokenList.Items), tokenList.HasNext)
}

// Test Token Overview and Stats APIs
func TestGetTokenOverview(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	overview, err := client.GetTokenOverview(ctx, testTokenSOL, nil)
	if err != nil {
		t.Fatalf("GetTokenOverview failed: %v", err)
	}

	if overview.Symbol == "" {
		t.Error("Expected non-empty symbol")
	}

	t.Logf("Token: %s, Price: $%.2f, Market Cap: $%.2f", overview.Symbol, overview.Price, overview.MarketCap)
}

func TestGetTokenPriceStats(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	timeframes := []string{"1m", "5m", "1h"}
	stats, err := client.GetTokenPriceStats(ctx, testTokenSOL, timeframes, nil)
	if err != nil {
		t.Fatalf("GetTokenPriceStats failed: %v", err)
	}

	if len(stats.Data) == 0 {
		t.Error("Expected at least one stats data")
	}

	t.Logf("Retrieved price stats for %d timeframes", len(stats.Data))
}

func TestGetMultiTokenPriceStats(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	addresses := []string{testTokenSOL}
	timeframes := []string{"1h", "24h"}
	stats, err := client.GetMultiTokenPriceStats(ctx, addresses, timeframes, nil)
	if err != nil {
		t.Fatalf("GetMultiTokenPriceStats failed: %v", err)
	}

	if len(stats) == 0 {
		t.Error("Expected at least one stats")
	}

	t.Logf("Retrieved price stats for %d tokens", len(stats))
}

func TestGetTokenTrendingList(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	trending, err := client.GetTokenTrendingList(ctx, &TrendingListOptions{
		Limit: 10,
	})
	if err != nil {
		t.Fatalf("GetTokenTrendingList failed: %v", err)
	}

	if len(trending.Tokens) == 0 {
		t.Error("Expected at least one trending token")
	}

	t.Logf("Retrieved %d trending tokens, Total: %d", len(trending.Tokens), trending.Total)
}

// Test Token Security and Creation APIs
func TestGetTokenSecurity(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	security, err := client.GetTokenSecurity(ctx, testTokenSOL, nil)
	if err != nil {
		t.Fatalf("GetTokenSecurity failed: %v", err)
	}

	t.Logf("Token security - Is Token 2022: %v", security.IsToken2022)
}

func TestGetTokenCreationInfo(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	creationInfo, err := client.GetTokenCreationInfo(ctx, testTokenBONK, nil)
	if err != nil {
		t.Fatalf("GetTokenCreationInfo failed: %v", err)
	}

	if creationInfo.TxHash == "" {
		t.Error("Expected non-empty txHash")
	}

	t.Logf("Token created at slot: %d", creationInfo.Slot)
}

// Test Token Holders APIs
func TestGetTokenHolders(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	holders, err := client.GetTokenHolders(ctx, testTokenSOL, &TokenHoldersOptions{
		Limit: 10,
	})
	if err != nil {
		t.Fatalf("GetTokenHolders failed: %v", err)
	}

	if len(holders) == 0 {
		t.Error("Expected at least one holder")
	}

	t.Logf("Retrieved %d token holders", len(holders))
}

func TestGetTokenHolderBatch(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	wallets := []string{testWalletAddr}
	holders, err := client.GetTokenHolderBatch(ctx, testTokenSOL, wallets, nil)
	if err != nil {
		t.Fatalf("GetTokenHolderBatch failed: %v", err)
	}

	t.Logf("Retrieved holder info for %d wallets", len(holders))
}

// Test Other Token APIs
func TestGetTokenTopTraders(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	traders, err := client.GetTokenTopTraders(ctx, testTokenSOL, &TokenTopTradersOptions{
		Limit: 5,
	})
	if err != nil {
		t.Fatalf("GetTokenTopTraders failed: %v", err)
	}

	t.Logf("Retrieved %d top traders", len(traders))
}

func TestGetTokenAllMarketList(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	markets, err := client.GetTokenAllMarketList(ctx, testTokenSOL, &TokenAllMarketListOptions{
		Limit: 10,
	})
	if err != nil {
		t.Fatalf("GetTokenAllMarketList failed: %v", err)
	}

	if len(markets.Items) == 0 {
		t.Error("Expected at least one market")
	}

	t.Logf("Retrieved %d markets, Total: %d", len(markets.Items), markets.Total)
}

func TestGetTokenAllTimeTrades(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	trades, err := client.GetTokenAllTimeTrades(ctx, testTokenSOL, nil)
	if err != nil {
		t.Fatalf("GetTokenAllTimeTrades failed: %v", err)
	}

	if trades.TotalTrade == 0 {
		t.Error("Expected non-zero total trades")
	}

	t.Logf("All-time trades: %d, Total volume: $%.2f", trades.TotalTrade, trades.TotalVolumeUSD)
}

func TestGetMultiTokenAllTimeTrades(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	addresses := []string{testTokenSOL, testTokenUSDC}
	trades, err := client.GetMultiTokenAllTimeTrades(ctx, addresses, nil)
	if err != nil {
		t.Fatalf("GetMultiTokenAllTimeTrades failed: %v", err)
	}

	if len(trades) == 0 {
		t.Error("Expected at least one trade data")
	}

	t.Logf("Retrieved all-time trades for %d tokens", len(trades))
}

func TestGetTokenMintBurnTxs(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	txs, err := client.GetTokenMintBurnTxs(ctx, testTokenSOL, &TokenMintBurnTxsOptions{
		Limit: 10,
	})
	if err != nil {
		t.Fatalf("GetTokenMintBurnTxs failed: %v", err)
	}

	t.Logf("Retrieved %d mint/burn transactions", len(txs))
}

func TestGetNewListing(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	listings, err := client.GetNewListing(ctx, &NewListingOptions{
		Limit: 10,
	})
	if err != nil {
		t.Fatalf("GetNewListing failed: %v", err)
	}

	t.Logf("Retrieved %d new listings", len(listings))
}

// Test Token Exit Liquidity APIs
func TestGetTokenExitLiquidity(t *testing.T) {
	t.Skip("Exit liquidity is Base chain only - skipping for Solana")
}

func TestGetMultiTokenExitLiquidity(t *testing.T) {
	t.Skip("Exit liquidity is Base chain only - skipping for Solana")
}

// Test Wallet APIs
func TestGetWalletPortfolio(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	portfolio, err := client.GetWalletPortfolio(ctx, testWalletAddr, nil)
	if err != nil {
		t.Fatalf("GetWalletPortfolio failed: %v", err)
	}

	t.Logf("Wallet has %d tokens in portfolio", len(portfolio))
}

func TestGetWalletTxs(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	txs, err := client.GetWalletTxs(ctx, testWalletAddr, &WalletTxsOptions{
		Limit: 10,
	})
	if err != nil {
		t.Fatalf("GetWalletTxs failed: %v", err)
	}

	for chain, txList := range txs {
		t.Logf("Chain %s: %d transactions", chain, len(txList))
	}
}

func TestGetWalletTrades(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	trades, err := client.GetWalletTrades(ctx, testWalletAddr, &WalletTradesOptions{
		Limit: 10,
	})
	if err != nil {
		t.Fatalf("GetWalletTrades failed: %v", err)
	}

	t.Logf("Retrieved %d wallet trades, HasNext: %v", len(trades.Items), trades.HasNext)
}

func TestGetWalletBalanceChanges(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	changes, err := client.GetWalletBalanceChanges(ctx, testWalletAddr, testTokenSOL, &WalletBalanceChangesOptions{
		Limit: 10,
	})
	if err != nil {
		t.Fatalf("GetWalletBalanceChanges failed: %v", err)
	}

	t.Logf("Retrieved %d balance changes", len(changes))
}

func TestGetWalletTokenBalance(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	balance, err := client.GetWalletTokenBalance(ctx, testWalletAddr, "GTArCNbysqJYqN1rwZnyPwSvzNy7sM3fD6MH2qyerNqc", nil)
	if err != nil {
		t.Fatalf("GetWalletTokenBalance failed: %v", err)
	}

	if balance.Address == "" {
		t.Error("Expected non-empty address")
	}

	t.Logf("Token balance: %.4f %s", balance.UIAmount, balance.Symbol)
}

func TestGetWalletTokensBalance(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	addresses := []string{testTokenSOL, testTokenUSDC}
	balances, err := client.GetWalletTokensBalance(ctx, testWalletAddr, addresses, nil)
	if err != nil {
		t.Fatalf("GetWalletTokensBalance failed: %v", err)
	}

	t.Logf("Retrieved balances for %d tokens", len(balances))
}

func TestGetWalletNetWorth(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	netWorth, err := client.GetWalletNetWorth(ctx, testWalletAddr, nil)
	if err != nil {
		t.Fatalf("GetWalletNetWorth failed: %v", err)
	}

	t.Logf("Wallet net worth: $%s, %d items", netWorth.TotalValue, len(netWorth.Items))
}

func TestGetWalletNetWorthHistories(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	histories, err := client.GetWalletNetWorthHistories(ctx, testWalletAddr, nil)
	if err != nil {
		t.Fatalf("GetWalletNetWorthHistories failed: %v", err)
	}

	t.Logf("Retrieved %d net worth history items", len(histories.History))
}

func TestGetWalletNetWorthDetails(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	details, err := client.GetWalletNetWorthDetails(ctx, testWalletAddr, nil)
	if err != nil {
		t.Fatalf("GetWalletNetWorthDetails failed: %v", err)
	}

	t.Logf("Net worth: $%.2f, Assets: %d", details.NetWorth, len(details.NetAssets))
}

func TestGetWalletTokensPnL(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	tokens := []string{testTokenSOL}
	pnl, err := client.GetWalletTokensPnL(ctx, testWalletAddr, tokens, nil)
	if err != nil {
		t.Fatalf("GetWalletTokensPnL failed: %v", err)
	}

	if pnl.Meta.Address == "" {
		t.Error("Expected non-empty address")
	}

	t.Logf("Retrieved PnL for %d tokens", len(pnl.Tokens))
}

func TestGetWalletsPnLByToken(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	wallets := []string{testWalletAddr}
	pnl, err := client.GetWalletsPnLByToken(ctx, testTokenSOL, wallets, nil)
	if err != nil {
		t.Fatalf("GetWalletsPnLByToken failed: %v", err)
	}

	t.Logf("Retrieved PnL for %d wallets", len(pnl.Data))
}

func TestGetWalletTokenFirstTx(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	wallets := []string{testWalletAddr}
	firstTxs, err := client.GetWalletTokenFirstTx(ctx, wallets, testTokenSOL, nil)
	if err != nil {
		t.Fatalf("GetWalletTokenFirstTx failed: %v", err)
	}

	t.Logf("Retrieved first tx info for %d wallets", len(firstTxs))
}

// Test Meme APIs
func TestGetMemeList(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	memes, err := client.GetMemeList(ctx, &MemeListOptions{
		Limit: 10,
	})
	if err != nil {
		t.Fatalf("GetMemeList failed: %v", err)
	}

	t.Logf("Retrieved %d meme tokens, HasNext: %v", len(memes.Items), memes.HasNext)
}

func TestGetMemeDetail(t *testing.T) {
	t.Skip("Skipping GetMemeDetail - requires specific meme token address")
}

// Test Other APIs
func TestGetGainersLosers(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	gainersLosers, err := client.GetGainersLosers(ctx, &GainersLosersOptions{
		Limit: 10,
	})
	if err != nil {
		t.Fatalf("GetGainersLosers failed: %v", err)
	}

	t.Logf("Retrieved %d gainers/losers", len(gainersLosers))
}

func TestSearch(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	results, err := client.Search(ctx, "SOL", &SearchOptions{
		Limit: 10,
	})
	if err != nil {
		t.Fatalf("Search failed: %v", err)
	}

	if len(results) == 0 {
		t.Error("Expected at least one search result")
	}

	t.Logf("Search returned %d results", len(results))
}

func TestGetLatestBlockNumber(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	blockNumber, err := client.GetLatestBlockNumber(ctx, []Chain{ChainSolana})
	if err != nil {
		t.Fatalf("GetLatestBlockNumber failed: %v", err)
	}

	if blockNumber == 0 {
		t.Error("Expected non-zero block number")
	}

	t.Logf("Latest block number: %d", blockNumber)
}

// Test All Transactions V3
func TestGetAllTxs(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	txs, err := client.GetAllTxs(ctx, &AllTxsV3Options{
		Limit: 10,
	})
	if err != nil {
		t.Fatalf("GetAllTxs failed: %v", err)
	}

	t.Logf("Retrieved %d transactions, HasNext: %v", len(txs.Items), txs.HasNext)
}

func TestGetRecentTxs(t *testing.T) {
	client := getTestClient(t)
	ctx := context.Background()

	txs, err := client.GetRecentTxs(ctx, &RecentTxsV3Options{
		Limit: 10,
	})
	if err != nil {
		t.Fatalf("GetRecentTxs failed: %v", err)
	}

	t.Logf("Retrieved %d recent transactions, HasNext: %v", len(txs.Items), txs.HasNext)
}
