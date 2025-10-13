# birdeye-go

A comprehensive Go client library for the [Birdeye API](https://birdeye.so/), providing real-time and historical cryptocurrency data across multiple blockchains.

[![Go Version](https://img.shields.io/badge/go-%3E%3D1.19-blue)](https://go.dev/)
[![License](https://img.shields.io/badge/license-MIT-green)](LICENSE)
[![Test Coverage](https://img.shields.io/badge/coverage-75%25-brightgreen)](https://github.com/dwdwow/birdeye-go)

## Features

### HTTP Client

- ✅ **Comprehensive API Coverage** - All Birdeye REST API endpoints
- ✅ **Automatic Rate Limiting** - Built-in rate limiter with configurable limits
- ✅ **Multi-Chain Support** - Solana, Ethereum, Arbitrum, Avalanche, BSC, Optimism, Polygon, Base, Zksync, Sui
- ✅ **Type Safety** - Strongly typed request/response structures
- ✅ **Context Support** - Full context cancellation and timeout support
- ✅ **Error Handling** - Comprehensive error types and handling
- ✅ **Detailed Documentation** - GoDoc style documentation for all methods

### WebSocket Client

- ✅ **Real-Time Data Streams** - Price updates, transactions, new listings, large trades
- ✅ **Synchronous Design** - Simple Read/Write methods for full control
- ✅ **Type-Safe Messages** - Strongly typed message and subscription structures
- ✅ **Multi-Chain Support** - Subscribe to data across different blockchains
- ✅ **Flexible Architecture** - Implement your own reconnection and error handling logic

## Installation

```bash
go get github.com/dwdwow/birdeye-go
```

## Quick Start

### HTTP Client Example

```go
package main

import (
    "context"
    "fmt"
    "log"
    "time"

    birdeye "github.com/dwdwow/birdeye-go"
)

func main() {
    // Create HTTP client
    client := birdeye.NewHTTPClient(birdeye.HTTPClientConfig{
        APIKey:  "your-api-key",
        Timeout: 30 * time.Second,
    })

    ctx := context.Background()

    // Get token price
    price, err := client.GetTokenPrice(ctx, "So11111111111111111111111111111111111111112", &birdeye.TokenPriceOptions{
        Chains: []birdeye.Chain{birdeye.ChainSolana},
    })
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("SOL Price: $%.2f\n", price.Value)

    // Get multiple token prices
    addresses := []string{
        "So11111111111111111111111111111111111111112", // SOL
        "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v", // USDC
    }
    prices, err := client.GetMultiTokenPrice(ctx, addresses, &birdeye.MultiTokenPriceOptions{
        Chains: []birdeye.Chain{birdeye.ChainSolana},
    })
    if err != nil {
        log.Fatal(err)
    }
    for addr, price := range prices {
        fmt.Printf("Token %s: $%.2f\n", addr[:8], price.Value)
    }
}
```

### WebSocket Client Example

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "log"

    birdeye "github.com/dwdwow/birdeye-go"
)

func main() {
    // Create WebSocket client
    client := birdeye.NewWSClient(birdeye.WSClientConfig{
        APIKey: "your-api-key",
        Chain:  birdeye.ChainSolana,
    })

    // Connect to WebSocket
    ctx := context.Background()
    if err := client.Connect(ctx); err != nil {
        log.Fatal(err)
    }
    defer client.Close()

    // Subscribe to SOL price updates
    sub := birdeye.SubDataPrice{
        Address:   "So11111111111111111111111111111111111111112",
        ChartType: birdeye.WsInterval1m,
        Currency:  birdeye.CurrencyUSD,
        QueryType: birdeye.QueryTypeSimple,
    }
    payload, _ := sub.Payload()
    if err := client.Subscribe(payload); err != nil {
        log.Fatal(err)
    }

    // Read messages in a loop
    for {
        msgType, data, err := client.Read()
        if err != nil {
            log.Printf("Read error: %v", err)
            break
        }

        // Handle different message types
        switch msgType {
        case birdeye.WsDataWelcome:
            fmt.Println("Connected to Birdeye WebSocket!")

        case birdeye.WsDataPriceData:
            var priceData birdeye.WsDataPrice
            if err := json.Unmarshal(data, &priceData); err != nil {
                log.Printf("Unmarshal error: %v", err)
                continue
            }
            fmt.Printf("SOL Price: $%.2f at %d\n", priceData.C, priceData.UnixTime)

        case birdeye.WsDataTxsData:
            var txData birdeye.WsDataTxs
            if err := json.Unmarshal(data, &txData); err != nil {
                log.Printf("Unmarshal error: %v", err)
                continue
            }
            fmt.Printf("Transaction: %s %s $%.2f\n", 
                txData.Side, txData.TxHash[:8], txData.VolumeUSD)

        case birdeye.WsDataError:
            fmt.Printf("Error from server: %s\n", string(data))
        }
    }
}
```

## API Overview

### Token APIs

```go
// Price data
price, err := client.GetTokenPrice(ctx, tokenAddress, opts)
prices, err := client.GetMultiTokenPrice(ctx, addresses, opts)

// Market data
marketData, err := client.GetTokenMarketData(ctx, tokenAddress, opts)
multiMarketData, err := client.GetMultiTokenMarketData(ctx, addresses, opts)

// Token metadata
metadata, err := client.GetTokenMetadata(ctx, tokenAddress, opts)
multiMetadata, err := client.GetMultiTokenMetadata(ctx, addresses, opts)

// Trade data
tradeData, err := client.GetTokenTradeData(ctx, tokenAddress, opts)
multiTradeData, err := client.GetMultiTokenTradeData(ctx, addresses, opts)

// OHLCV data
ohlcv, err := client.GetTokenOHLCV(ctx, tokenAddress, "1D", timeFrom, timeTo, opts)
ohlcvV3, err := client.GetTokenOHLCVV3(ctx, tokenAddress, "1D", timeFrom, timeTo, opts)

// Transactions
txs, err := client.GetTokenTxs(ctx, tokenAddress, opts)
txsV3, err := client.GetTokenTxsV3(ctx, tokenAddress, opts)
txsByTime, err := client.GetTokenTxsByTime(ctx, tokenAddress, opts)

// Token lists
tokenList, err := client.GetTokenListV3(ctx, opts)
tokenListV1, err := client.GetTokenListV1(ctx, opts)

// Token overview
overview, err := client.GetTokenOverview(ctx, tokenAddress, opts)

// Security
security, err := client.GetTokenSecurity(ctx, tokenAddress, opts)

// Holders
holders, err := client.GetTokenHolders(ctx, tokenAddress, opts)
```

### Wallet APIs

```go
// Portfolio
portfolio, err := client.GetWalletPortfolio(ctx, walletAddress, opts)

// Transactions
txs, err := client.GetWalletTxs(ctx, walletAddress, opts)
trades, err := client.GetWalletTrades(ctx, walletAddress, opts)

// Balance
balance, err := client.GetWalletTokenBalance(ctx, wallet, tokenAddress, opts)
balances, err := client.GetWalletTokensBalance(ctx, wallet, tokenAddresses, opts)
balanceChanges, err := client.GetWalletBalanceChanges(ctx, wallet, tokenAddress, opts)

// Net worth
netWorth, err := client.GetWalletNetWorth(ctx, walletAddress, opts)
netWorthHistory, err := client.GetWalletNetWorthHistories(ctx, walletAddress, opts)
netWorthDetails, err := client.GetWalletNetWorthDetails(ctx, walletAddress, opts)

// PnL
pnl, err := client.GetWalletTokensPnL(ctx, wallet, tokenAddresses, opts)
walletsPnL, err := client.GetWalletsPnLByToken(ctx, tokenAddress, wallets, opts)
```

### Pair/DEX APIs

```go
// Pair data
pairOverview, err := client.GetPairOverview(ctx, pairAddress, opts)
pairsOverview, err := client.GetPairsOverview(ctx, addresses, opts)

// Pair OHLCV
ohlcv, err := client.GetPairOHLCV(ctx, pairAddress, "1D", timeFrom, timeTo, opts)
ohlcvV3, err := client.GetPairOHLCVV3(ctx, pairAddress, "1D", timeFrom, timeTo, opts)

// Base/Quote OHLCV
ohlcv, err := client.GetOHLCVBaseQuote(ctx, baseAddr, quoteAddr, "1D", timeFrom, timeTo, opts)

// Pair transactions
txs, err := client.GetPairTxs(ctx, pairAddress, opts)
txsByTime, err := client.GetPairTxsByTime(ctx, pairAddress, opts)
```

### Market Analysis APIs

```go
// Trending tokens
trending, err := client.GetTokenTrendingList(ctx, opts)

// New listings
newListings, err := client.GetNewListing(ctx, opts)

// Gainers & losers
gainersLosers, err := client.GetGainersLosers(ctx, opts)

// Top traders
topTraders, err := client.GetTokenTopTraders(ctx, tokenAddress, opts)

// All markets
markets, err := client.GetTokenAllMarketList(ctx, tokenAddress, opts)

// Search
results, err := client.Search(ctx, keyword, opts)
```

### WebSocket Subscriptions

```go
// Price subscriptions
sub := birdeye.SubDataPrice{
    Address:   tokenAddress,
    ChartType: birdeye.WsInterval1m,
    Currency:  birdeye.CurrencyUSD,
    QueryType: birdeye.QueryTypeSimple,
}
payload, _ := sub.Payload()
client.Subscribe(payload)

// Transaction subscriptions
txSub := birdeye.SubDataTxs{
    Address:   &tokenAddress,
    QueryType: birdeye.QueryTypeSimple,
}
txPayload, _ := txSub.Payload()
client.Subscribe(txPayload)

// Complex subscriptions (multiple tokens)
prices := []birdeye.SubDataPrice{
    {Address: "addr1", ChartType: birdeye.WsInterval1m, Currency: birdeye.CurrencyUSD, QueryType: birdeye.QueryTypeComplex},
    {Address: "addr2", ChartType: birdeye.WsInterval5m, Currency: birdeye.CurrencyUSD, QueryType: birdeye.QueryTypeComplex},
}
complexPayload, _ := birdeye.PricesComplexPayload(prices)
client.Subscribe(complexPayload)

// New listing subscriptions
newListingSub := birdeye.SubDataTokenNewListing{
    MinLiquidity: &minLiq,
    MaxLiquidity: &maxLiq,
}
newListingPayload, _ := newListingSub.Payload()
client.Subscribe(newListingPayload)

// Large trade subscriptions
largeTradeSub := birdeye.SubDataLargeTradeTxs{
    MinVolume: 10000.0,
    MaxVolume: &maxVol,
}
largeTradePayload, _ := largeTradeSub.Payload()
client.Subscribe(largeTradePayload)

// Wallet transaction subscriptions
walletSub := birdeye.SubDataWalletTxs{
    Address: walletAddress,
}
walletPayload, _ := walletSub.Payload()
client.Subscribe(walletPayload)

// Token stats subscriptions
statsSub := birdeye.SubDataTokenStats{
    Address: tokenAddress,
    Select:  birdeye.NewTokenStatsSelect(),
}
statsPayload, _ := statsSub.Payload()
client.Subscribe(statsPayload)

// Unsubscribe
client.Unsubscribe(birdeye.UnsubscribePrice, sub)
```

## Configuration

### HTTP Client Configuration

```go
config := birdeye.HTTPClientConfig{
    APIKey:  "your-api-key",
    Timeout: 30 * time.Second,
    
    // Optional: Custom rate limiter
    RateLimiter: birdeye.NewRateLimiter(60, time.Minute), // 60 requests per minute
    
    // Optional: HTTP client
    HTTPClient: &http.Client{
        Timeout: 30 * time.Second,
        Transport: &http.Transport{
            MaxIdleConns:        100,
            MaxIdleConnsPerHost: 10,
        },
    },
}

client := birdeye.NewHTTPClient(config)
```

### WebSocket Client Configuration

```go
// Simple configuration
config := birdeye.WSClientConfig{
    APIKey: "your-api-key",
    Chain:  birdeye.ChainSolana,
}

client := birdeye.NewWSClient(config)

// Connect to WebSocket
ctx := context.Background()
if err := client.Connect(ctx); err != nil {
    log.Fatal(err)
}
defer client.Close()

// Subscribe to data streams
sub := birdeye.SubDataPrice{
    Address:   "token-address",
    ChartType: birdeye.WsInterval1m,
    Currency:  birdeye.CurrencyUSD,
    QueryType: birdeye.QueryTypeSimple,
}
payload, _ := sub.Payload()
client.Subscribe(payload)

// Read messages in your own loop
for {
    msgType, data, err := client.Read()
    if err != nil {
        log.Printf("Error: %v", err)
        break
    }
    // Handle msgType and data
}
```

## Supported Chains

```go
birdeye.ChainSolana      // Solana
birdeye.ChainEthereum    // Ethereum
birdeye.ChainArbitrum    // Arbitrum
birdeye.ChainAvalanche   // Avalanche
birdeye.ChainBsc         // Binance Smart Chain
birdeye.ChainOptimism    // Optimism
birdeye.ChainPolygon     // Polygon
birdeye.ChainBase        // Base
birdeye.ChainZksync      // zkSync
birdeye.ChainSui         // Sui
```

## Rate Limiting

The library includes automatic rate limiting to prevent API quota exhaustion:

```go
// Default rate limiter: 60 requests per minute
client := birdeye.NewHTTPClient(birdeye.HTTPClientConfig{
    APIKey: "your-api-key",
})

// Custom rate limiter
customLimiter := birdeye.NewRateLimiter(100, time.Minute) // 100 req/min
client := birdeye.NewHTTPClient(birdeye.HTTPClientConfig{
    APIKey:      "your-api-key",
    RateLimiter: customLimiter,
})

// Shared rate limiter across multiple clients
sharedLimiter := birdeye.NewRateLimiter(100, time.Minute)
client1 := birdeye.NewHTTPClient(birdeye.HTTPClientConfig{
    APIKey:      "api-key-1",
    RateLimiter: sharedLimiter,
})
client2 := birdeye.NewHTTPClient(birdeye.HTTPClientConfig{
    APIKey:      "api-key-2",
    RateLimiter: sharedLimiter,
})

// Handle rate limit exceeded
opts := &birdeye.TokenPriceOptions{
    OnLimitExceeded: func() error {
        return fmt.Errorf("custom rate limit error")
    },
}
```

## Error Handling

```go
price, err := client.GetTokenPrice(ctx, tokenAddress, opts)
if err != nil {
    // Check for specific error types
    switch {
    case errors.Is(err, context.Canceled):
        log.Println("Request was cancelled")
    case errors.Is(err, context.DeadlineExceeded):
        log.Println("Request timeout")
    default:
        log.Printf("API error: %v", err)
    }
    return
}
```

## Context and Timeout

All API methods support context for cancellation and timeout:

```go
// With timeout
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()

price, err := client.GetTokenPrice(ctx, tokenAddress, opts)

// With cancellation
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

go func() {
    time.Sleep(5 * time.Second)
    cancel() // Cancel after 5 seconds
}()

price, err := client.GetTokenPrice(ctx, tokenAddress, opts)
```

## Advanced Examples

### Batch Token Price Monitoring

```go
func monitorTokenPrices(ctx context.Context, client *birdeye.HTTPClient, tokens []string) {
    ticker := time.NewTicker(10 * time.Second)
    defer ticker.Stop()

    for {
        select {
        case <-ctx.Done():
            return
        case <-ticker.C:
            prices, err := client.GetMultiTokenPrice(ctx, tokens, &birdeye.MultiTokenPriceOptions{
                Chains: []birdeye.Chain{birdeye.ChainSolana},
            })
            if err != nil {
                log.Printf("Error fetching prices: %v", err)
                continue
            }

            for addr, price := range prices {
                log.Printf("Token %s: $%.2f (24h change: %.2f%%)", 
                    addr[:8], price.Value, price.PriceChange24h)
            }
        }
    }
}
```

### WebSocket with Reconnection Logic

```go
func subscribeWithReconnect(apiKey, tokenAddress string) {
    for {
        client := birdeye.NewWSClient(birdeye.WSClientConfig{
            APIKey: apiKey,
            Chain:  birdeye.ChainSolana,
        })
        
        ctx := context.Background()
        if err := client.Connect(ctx); err != nil {
            log.Printf("Connection failed: %v, retrying in 5s...", err)
            time.Sleep(5 * time.Second)
            continue
        }
        
        log.Println("Connected to Birdeye WebSocket!")
        
        // Subscribe to price updates
        sub := birdeye.SubDataPrice{
            Address:   tokenAddress,
            ChartType: birdeye.WsInterval1m,
            Currency:  birdeye.CurrencyUSD,
            QueryType: birdeye.QueryTypeSimple,
        }
        payload, _ := sub.Payload()
        if err := client.Subscribe(payload); err != nil {
            log.Printf("Subscribe failed: %v", err)
            client.Close()
            time.Sleep(5 * time.Second)
            continue
        }
        
        // Read messages
        for {
            msgType, data, err := client.Read()
            if err != nil {
                log.Printf("Read error: %v, reconnecting...", err)
                client.Close()
                time.Sleep(5 * time.Second)
                break // Break to outer loop to reconnect
            }
            
            // Handle messages
            switch msgType {
            case birdeye.WsDataPriceData:
                var priceData birdeye.WsDataPrice
                json.Unmarshal(data, &priceData)
                log.Printf("Price: $%.2f", priceData.C)
            case birdeye.WsDataError:
                log.Printf("Server error: %s", string(data))
            }
        }
    }
}
```

### Portfolio Tracking

```go
func trackWalletPortfolio(ctx context.Context, client *birdeye.HTTPClient, wallet string) {
    // Get current portfolio
    portfolio, err := client.GetWalletPortfolio(ctx, wallet, &birdeye.WalletPortfolioOptions{
        Chains: []birdeye.Chain{birdeye.ChainSolana},
    })
    if err != nil {
        log.Fatal(err)
    }

    var totalValue float64
    for _, item := range portfolio.Items {
        value := item.PriceUSD * item.UIAmount
        totalValue += value
        fmt.Printf("Token: %s, Balance: %.2f, Value: $%.2f\n", 
            item.Symbol, item.UIAmount, value)
    }
    fmt.Printf("Total Portfolio Value: $%.2f\n", totalValue)

    // Get net worth history
    history, err := client.GetWalletNetWorthHistories(ctx, wallet, &birdeye.WalletNetWorthHistoriesOptions{
        Count: 30,
        Type:  "1D",
        Chains: []birdeye.Chain{birdeye.ChainSolana},
    })
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("30-day net worth trend:\n")
    for _, h := range history.Items {
        fmt.Printf("  %s: $%.2f\n", h.Time, h.TotalValue)
    }
}
```

## Testing

Run all tests:

```bash
go test -v ./...
```

Run tests with coverage:

```bash
go test -cover -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

Run specific tests:

```bash
go test -v -run TestHTTPClient
go test -v -run TestWSClient
go test -v -run TestRateLimiter
```

## Documentation

Full API documentation is available on [GoDoc](https://pkg.go.dev/github.com/dwdwow/birdeye-go).

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Resources

- [Birdeye Official Website](https://birdeye.so/)
- [Birdeye API Documentation](https://docs.birdeye.so/)
- [GoDoc Documentation](https://pkg.go.dev/github.com/dwdwow/birdeye-go)

## Support

For bugs, questions, and discussions please use the [GitHub Issues](https://github.com/dwdwow/birdeye-go/issues).

## Acknowledgments

- Thanks to [Birdeye](https://birdeye.so/) for providing the API
- Built with [gorilla/websocket](https://github.com/gorilla/websocket) for WebSocket support
