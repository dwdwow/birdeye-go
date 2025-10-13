package birdeye

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

// ============================================================================
// Subscription Payload Tests
// ============================================================================

func TestSubDataPrice_Query(t *testing.T) {
	sub := SubDataPrice{
		Address:   "So11111111111111111111111111111111111111112",
		ChartType: WsInterval1m,
		Currency:  CurrencyUSD,
		QueryType: QueryTypeSimple,
	}

	expected := "(address=So11111111111111111111111111111111111111112 AND chartType=1m AND currency=usd AND queryType=simple)"
	if got := sub.Query(); got != expected {
		t.Errorf("SubDataPrice.Query() = %v, want %v", got, expected)
	}
}

func TestSubDataPrice_Payload(t *testing.T) {
	sub := SubDataPrice{
		Address:   "So11111111111111111111111111111111111111112",
		ChartType: WsInterval1m,
		Currency:  CurrencyUSD,
		QueryType: QueryTypeSimple,
	}

	payload, err := sub.Payload()
	if err != nil {
		t.Fatalf("SubDataPrice.Payload() error = %v", err)
	}

	var result map[string]any
	if err := json.Unmarshal(payload, &result); err != nil {
		t.Fatalf("Failed to unmarshal payload: %v", err)
	}

	if result["type"] != string(SubscribePrice) {
		t.Errorf("Expected type %s, got %v", SubscribePrice, result["type"])
	}

	data, ok := result["data"].(map[string]any)
	if !ok {
		t.Fatal("Expected data to be a map")
	}

	if data["address"] != sub.Address {
		t.Errorf("Expected address %s, got %v", sub.Address, data["address"])
	}
}

func TestPricesComplexPayload(t *testing.T) {
	prices := []SubDataPrice{
		{
			Address:   "addr1",
			ChartType: WsInterval1m,
			Currency:  CurrencyUSD,
			QueryType: QueryTypeComplex,
		},
		{
			Address:   "addr2",
			ChartType: WsInterval5m,
			Currency:  CurrencyPair,
			QueryType: QueryTypeComplex,
		},
	}

	payload, err := PricesComplexPayload(prices)
	if err != nil {
		t.Fatalf("PricesComplexPayload() error = %v", err)
	}

	var result map[string]any
	if err := json.Unmarshal(payload, &result); err != nil {
		t.Fatalf("Failed to unmarshal payload: %v", err)
	}

	if result["type"] != string(SubscribePrice) {
		t.Errorf("Expected type %s, got %v", SubscribePrice, result["type"])
	}

	data, ok := result["data"].(map[string]any)
	if !ok {
		t.Fatal("Expected data to be a map")
	}

	query, ok := data["query"].(string)
	if !ok {
		t.Fatal("Expected query to be a string")
	}

	if !strings.Contains(query, "addr1") || !strings.Contains(query, "addr2") {
		t.Errorf("Expected query to contain both addresses, got %s", query)
	}

	if !strings.Contains(query, " OR ") {
		t.Errorf("Expected query to contain OR operator, got %s", query)
	}
}

func TestSubDataTxs_Query(t *testing.T) {
	tests := []struct {
		name     string
		sub      SubDataTxs
		expected string
	}{
		{
			name: "with address",
			sub: SubDataTxs{
				Address:   stringPtr("addr1"),
				QueryType: QueryTypeSimple,
			},
			expected: "address=addr1",
		},
		{
			name: "with pair address",
			sub: SubDataTxs{
				PairAddress: stringPtr("pair1"),
				QueryType:   QueryTypeSimple,
			},
			expected: "pairAddress=pair1",
		},
		{
			name: "empty",
			sub: SubDataTxs{
				QueryType: QueryTypeSimple,
			},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.sub.Query(); got != tt.expected {
				t.Errorf("SubDataTxs.Query() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestSubDataTxs_Payload(t *testing.T) {
	sub := SubDataTxs{
		Address:   stringPtr("So11111111111111111111111111111111111111112"),
		QueryType: QueryTypeSimple,
	}

	payload, err := sub.Payload()
	if err != nil {
		t.Fatalf("SubDataTxs.Payload() error = %v", err)
	}

	var result map[string]any
	if err := json.Unmarshal(payload, &result); err != nil {
		t.Fatalf("Failed to unmarshal payload: %v", err)
	}

	if result["type"] != string(SubscribeTxs) {
		t.Errorf("Expected type %s, got %v", SubscribeTxs, result["type"])
	}
}

func TestTxsComplexPayload(t *testing.T) {
	txs := []SubDataTxs{
		{Address: stringPtr("addr1"), QueryType: QueryTypeComplex},
		{PairAddress: stringPtr("pair1"), QueryType: QueryTypeComplex},
	}

	payload, err := TxsComplexPayload(txs)
	if err != nil {
		t.Fatalf("TxsComplexPayload() error = %v", err)
	}

	var result map[string]any
	if err := json.Unmarshal(payload, &result); err != nil {
		t.Fatalf("Failed to unmarshal payload: %v", err)
	}

	data, ok := result["data"].(map[string]any)
	if !ok {
		t.Fatal("Expected data to be a map")
	}

	query, ok := data["query"].(string)
	if !ok {
		t.Fatal("Expected query to be a string")
	}

	if !strings.Contains(query, "addr1") || !strings.Contains(query, "pair1") {
		t.Errorf("Expected query to contain both address and pairAddress, got %s", query)
	}
}

func TestSubDataBaseQuotePrice_Payload(t *testing.T) {
	sub := SubDataBaseQuotePrice{
		BaseAddress:  "base_addr",
		QuoteAddress: "quote_addr",
		ChartType:    WsInterval1m,
	}

	payload, err := sub.Payload()
	if err != nil {
		t.Fatalf("SubDataBaseQuotePrice.Payload() error = %v", err)
	}

	var result map[string]any
	if err := json.Unmarshal(payload, &result); err != nil {
		t.Fatalf("Failed to unmarshal payload: %v", err)
	}

	if result["type"] != string(SubscribeBaseQuotePrice) {
		t.Errorf("Expected type %s, got %v", SubscribeBaseQuotePrice, result["type"])
	}
}

func TestSubDataTokenNewListing_Payload(t *testing.T) {
	minLiq := 1000.0
	maxLiq := 10000.0
	memePlatform := true

	sub := SubDataTokenNewListing{
		MemePlatformEnabled: &memePlatform,
		MinLiquidity:        &minLiq,
		MaxLiquidity:        &maxLiq,
	}

	payload, err := sub.Payload()
	if err != nil {
		t.Fatalf("SubDataTokenNewListing.Payload() error = %v", err)
	}

	var result map[string]any
	if err := json.Unmarshal(payload, &result); err != nil {
		t.Fatalf("Failed to unmarshal payload: %v", err)
	}

	if result["type"] != string(SubscribeTokenNewListing) {
		t.Errorf("Expected type %s, got %v", SubscribeTokenNewListing, result["type"])
	}
}

func TestSubDataNewPair_Payload(t *testing.T) {
	minLiq := 1000.0
	sub := SubDataNewPair{
		MinLiquidity: &minLiq,
	}

	payload, err := sub.Payload()
	if err != nil {
		t.Fatalf("SubDataNewPair.Payload() error = %v", err)
	}

	var result map[string]any
	if err := json.Unmarshal(payload, &result); err != nil {
		t.Fatalf("Failed to unmarshal payload: %v", err)
	}

	if result["type"] != string(SubscribeNewPair) {
		t.Errorf("Expected type %s, got %v", SubscribeNewPair, result["type"])
	}
}

func TestSubDataLargeTradeTxs_Payload(t *testing.T) {
	maxVol := 100000.0
	sub := SubDataLargeTradeTxs{
		MinVolume: 1000.0,
		MaxVolume: &maxVol,
	}

	payload, err := sub.Payload()
	if err != nil {
		t.Fatalf("SubDataLargeTradeTxs.Payload() error = %v", err)
	}

	var result map[string]any
	if err := json.Unmarshal(payload, &result); err != nil {
		t.Fatalf("Failed to unmarshal payload: %v", err)
	}

	if result["type"] != string(SubscribeLargeTradeTxs) {
		t.Errorf("Expected type %s, got %v", SubscribeLargeTradeTxs, result["type"])
	}
}

func TestSubDataWalletTxs_Payload(t *testing.T) {
	sub := SubDataWalletTxs{
		Address: "wallet_addr",
	}

	payload, err := sub.Payload()
	if err != nil {
		t.Fatalf("SubDataWalletTxs.Payload() error = %v", err)
	}

	var result map[string]any
	if err := json.Unmarshal(payload, &result); err != nil {
		t.Fatalf("Failed to unmarshal payload: %v", err)
	}

	if result["type"] != string(SubscribeWalletTxs) {
		t.Errorf("Expected type %s, got %v", SubscribeWalletTxs, result["type"])
	}
}

func TestSubDataTokenStats_Payload(t *testing.T) {
	sub := SubDataTokenStats{
		Address: "token_addr",
		Select:  NewTokenStatsSelect(),
	}

	payload, err := sub.Payload()
	if err != nil {
		t.Fatalf("SubDataTokenStats.Payload() error = %v", err)
	}

	var result map[string]any
	if err := json.Unmarshal(payload, &result); err != nil {
		t.Fatalf("Failed to unmarshal payload: %v", err)
	}

	if result["type"] != string(SubscribeTokenStats) {
		t.Errorf("Expected type %s, got %v", SubscribeTokenStats, result["type"])
	}
}

func TestNewTokenStatsSelect(t *testing.T) {
	stats := NewTokenStatsSelect()

	if !stats.Price {
		t.Error("Expected Price to be true")
	}
	if !stats.FDV {
		t.Error("Expected FDV to be true")
	}
	if !stats.MarketCap {
		t.Error("Expected MarketCap to be true")
	}
	if !stats.Supply {
		t.Error("Expected Supply to be true")
	}
	if !stats.LastTrade {
		t.Error("Expected LastTrade to be true")
	}
	if !stats.Liquidity {
		t.Error("Expected Liquidity to be true")
	}
}

func TestNewTokenStatsSelectTradeData(t *testing.T) {
	tradeData := NewTokenStatsSelectTradeData()

	if !tradeData.Volume {
		t.Error("Expected Volume to be true")
	}
	if !tradeData.Trade {
		t.Error("Expected Trade to be true")
	}
	if !tradeData.PriceHistory {
		t.Error("Expected PriceHistory to be true")
	}
	if tradeData.UniqueWalletChange {
		t.Error("Expected UniqueWalletChange to be false")
	}
	if len(tradeData.Intervals) == 0 {
		t.Error("Expected Intervals to have default values")
	}
}

// ============================================================================
// WebSocket Client Tests
// ============================================================================

func TestNewWSClient(t *testing.T) {
	config := WSClientConfig{
		APIKey: "test-api-key",
		Chain:  ChainSolana,
	}

	client := NewWSClient(config)

	if client.APIKey != "test-api-key" {
		t.Errorf("Expected APIKey to be 'test-api-key', got %s", client.APIKey)
	}
	if client.Chain != ChainSolana {
		t.Errorf("Expected Chain to be %s, got %s", ChainSolana, client.Chain)
	}
	if client.conn != nil {
		t.Error("Expected conn to be nil before Connect()")
	}
}

func TestWSClient_Close_WithoutConnect(t *testing.T) {
	config := WSClientConfig{
		APIKey: "test-api-key",
		Chain:  ChainSolana,
	}

	client := NewWSClient(config)
	err := client.Close()
	if err != nil {
		t.Errorf("Close() error = %v, expected nil when not connected", err)
	}
}

// Mock WebSocket server for testing
func newMockWSServer(t *testing.T, handler func(*websocket.Conn)) *httptest.Server {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			t.Logf("Failed to upgrade connection: %v", err)
			return
		}
		defer conn.Close()
		handler(conn)
	}))
}

func TestWSClient_ConnectAndClose(t *testing.T) {
	server := newMockWSServer(t, func(conn *websocket.Conn) {
		// Just keep connection open
		time.Sleep(100 * time.Millisecond)
	})
	defer server.Close()

	// Create client with mock server URL
	client := NewWSClient(WSClientConfig{
		APIKey: "test-key",
		Chain:  ChainSolana,
	})

	// Manually connect to mock server
	wsURL := "ws" + strings.TrimPrefix(server.URL, "http")
	ctx := context.Background()

	dialer := websocket.Dialer{}
	conn, _, err := dialer.DialContext(ctx, wsURL, nil)
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}

	client.mu.Lock()
	client.conn = conn
	client.mu.Unlock()

	// Test Close
	err = client.Close()
	if err != nil {
		t.Errorf("Close() error = %v", err)
	}

	// Verify connection is closed
	client.mu.RLock()
	if client.conn != nil {
		// Connection object still exists but should be closed
		// Try to write to verify it's closed
		err := client.conn.WriteMessage(websocket.TextMessage, []byte("test"))
		if err == nil {
			t.Error("Expected error writing to closed connection")
		}
	}
	client.mu.RUnlock()
}

func TestWSClient_Send(t *testing.T) {
	receivedMsg := make(chan string, 1)
	server := newMockWSServer(t, func(conn *websocket.Conn) {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}
		receivedMsg <- string(msg)
	})
	defer server.Close()

	wsURL := "ws" + strings.TrimPrefix(server.URL, "http")
	ctx := context.Background()

	dialer := websocket.Dialer{}
	conn, _, err := dialer.DialContext(ctx, wsURL, nil)
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := NewWSClient(WSClientConfig{
		APIKey: "test-key",
		Chain:  ChainSolana,
	})
	client.mu.Lock()
	client.conn = conn
	client.mu.Unlock()

	// Send message
	testMsg := []byte(`{"type":"SUBSCRIBE_PRICE","data":{"address":"test"}}`)
	err = client.Send(testMsg)
	if err != nil {
		t.Fatalf("Send() error = %v", err)
	}

	// Verify server received the message
	select {
	case msg := <-receivedMsg:
		if msg != string(testMsg) {
			t.Errorf("Expected message %s, got %s", testMsg, msg)
		}
	case <-time.After(1 * time.Second):
		t.Error("Timeout waiting for message")
	}
}

func TestWSClient_Send_NilConnection(t *testing.T) {
	client := NewWSClient(WSClientConfig{
		APIKey: "test-key",
		Chain:  ChainSolana,
	})

	err := client.Send([]byte("test"))
	if err == nil {
		t.Error("Expected error when sending with nil connection")
	}
	if !strings.Contains(err.Error(), "conn is nil") {
		t.Errorf("Expected 'conn is nil' error, got: %v", err)
	}
}

func TestWSClient_Read(t *testing.T) {
	server := newMockWSServer(t, func(conn *websocket.Conn) {
		// Send welcome message
		welcome := WsData{
			Type: WsDataWelcome,
			Data: json.RawMessage(`{"message":"Connected"}`),
		}
		welcomeJSON, _ := json.Marshal(welcome)
		conn.WriteMessage(websocket.TextMessage, welcomeJSON)

		// Send price data
		priceData := WsData{
			Type: WsDataPriceData,
			Data: json.RawMessage(`{"c":100.5,"address":"test_addr"}`),
		}
		priceJSON, _ := json.Marshal(priceData)
		conn.WriteMessage(websocket.TextMessage, priceJSON)

		// Wait a bit before closing
		time.Sleep(100 * time.Millisecond)
	})
	defer server.Close()

	wsURL := "ws" + strings.TrimPrefix(server.URL, "http")
	ctx := context.Background()

	dialer := websocket.Dialer{}
	conn, _, err := dialer.DialContext(ctx, wsURL, nil)
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := NewWSClient(WSClientConfig{
		APIKey: "test-key",
		Chain:  ChainSolana,
	})
	client.mu.Lock()
	client.conn = conn
	client.mu.Unlock()

	// Read welcome message
	msgType, data, err := client.Read()
	if err != nil {
		t.Fatalf("Read() error = %v", err)
	}
	if msgType != WsDataWelcome {
		t.Errorf("Expected WELCOME message, got %s", msgType)
	}

	var welcomeData map[string]string
	if err := json.Unmarshal(data, &welcomeData); err != nil {
		t.Fatalf("Failed to unmarshal welcome data: %v", err)
	}
	if welcomeData["message"] != "Connected" {
		t.Errorf("Expected message 'Connected', got %s", welcomeData["message"])
	}

	// Read price data
	msgType, data, err = client.Read()
	if err != nil {
		t.Fatalf("Read() error = %v", err)
	}
	if msgType != WsDataPriceData {
		t.Errorf("Expected PRICE_DATA message, got %s", msgType)
	}

	var priceData WsDataPrice
	if err := json.Unmarshal(data, &priceData); err != nil {
		t.Fatalf("Failed to unmarshal price data: %v", err)
	}
	if priceData.C != 100.5 {
		t.Errorf("Expected close price 100.5, got %.2f", priceData.C)
	}
	if priceData.Address != "test_addr" {
		t.Errorf("Expected address 'test_addr', got %s", priceData.Address)
	}
}

func TestWSClient_Read_NilConnection(t *testing.T) {
	client := NewWSClient(WSClientConfig{
		APIKey: "test-key",
		Chain:  ChainSolana,
	})

	msgType, data, err := client.Read()
	if err != nil {
		t.Errorf("Read() with nil connection should not error, got: %v", err)
	}
	if msgType != "" {
		t.Errorf("Expected empty msgType, got %s", msgType)
	}
	if data != nil {
		t.Errorf("Expected nil data, got %v", data)
	}
}

func TestWSClient_Subscribe(t *testing.T) {
	receivedMsg := make(chan string, 1)
	server := newMockWSServer(t, func(conn *websocket.Conn) {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}
		receivedMsg <- string(msg)
	})
	defer server.Close()

	wsURL := "ws" + strings.TrimPrefix(server.URL, "http")
	ctx := context.Background()

	dialer := websocket.Dialer{}
	conn, _, err := dialer.DialContext(ctx, wsURL, nil)
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := NewWSClient(WSClientConfig{
		APIKey: "test-key",
		Chain:  ChainSolana,
	})
	client.mu.Lock()
	client.conn = conn
	client.mu.Unlock()

	// Subscribe to price
	sub := SubDataPrice{
		Address:   "test_addr",
		ChartType: WsInterval1m,
		Currency:  CurrencyUSD,
		QueryType: QueryTypeSimple,
	}
	payload, _ := sub.Payload()
	err = client.Subscribe(payload)
	if err != nil {
		t.Fatalf("Subscribe() error = %v", err)
	}

	// Verify server received the subscription
	select {
	case msg := <-receivedMsg:
		var result map[string]any
		json.Unmarshal([]byte(msg), &result)
		if result["type"] != string(SubscribePrice) {
			t.Errorf("Expected subscribe type, got %v", result["type"])
		}
	case <-time.After(1 * time.Second):
		t.Error("Timeout waiting for subscription message")
	}
}

func TestWSClient_Unsubscribe(t *testing.T) {
	receivedMsg := make(chan string, 1)
	server := newMockWSServer(t, func(conn *websocket.Conn) {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}
		receivedMsg <- string(msg)
	})
	defer server.Close()

	wsURL := "ws" + strings.TrimPrefix(server.URL, "http")
	ctx := context.Background()

	dialer := websocket.Dialer{}
	conn, _, err := dialer.DialContext(ctx, wsURL, nil)
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := NewWSClient(WSClientConfig{
		APIKey: "test-key",
		Chain:  ChainSolana,
	})
	client.mu.Lock()
	client.conn = conn
	client.mu.Unlock()

	// Unsubscribe from price
	sub := SubDataPrice{
		Address:   "test_addr",
		ChartType: WsInterval1m,
		Currency:  CurrencyUSD,
		QueryType: QueryTypeSimple,
	}
	err = client.Unsubscribe(UnsubscribePrice, sub)
	if err != nil {
		t.Fatalf("Unsubscribe() error = %v", err)
	}

	// Verify server received the unsubscribe message
	select {
	case msg := <-receivedMsg:
		var result map[string]any
		json.Unmarshal([]byte(msg), &result)
		if result["type"] != string(UnsubscribePrice) {
			t.Errorf("Expected unsubscribe type, got %v", result["type"])
		}
	case <-time.After(1 * time.Second):
		t.Error("Timeout waiting for unsubscribe message")
	}
}

func TestWSClient_ConcurrentReadWrite(t *testing.T) {
	messageCount := 10
	receivedMessages := make(chan string, messageCount)
	sentMessages := make(chan struct{}, messageCount)

	server := newMockWSServer(t, func(conn *websocket.Conn) {
		// Echo server - read and send back
		for i := 0; i < messageCount; i++ {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}
			// Echo back with a counter
			response := WsData{
				Type: WsDataPriceData,
				Data: json.RawMessage(fmt.Sprintf(`{"counter":%d,"original":%s}`, i, msg)),
			}
			respJSON, _ := json.Marshal(response)
			conn.WriteMessage(websocket.TextMessage, respJSON)
		}
	})
	defer server.Close()

	wsURL := "ws" + strings.TrimPrefix(server.URL, "http")
	ctx := context.Background()

	dialer := websocket.Dialer{}
	conn, _, err := dialer.DialContext(ctx, wsURL, nil)
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := NewWSClient(WSClientConfig{
		APIKey: "test-key",
		Chain:  ChainSolana,
	})
	client.mu.Lock()
	client.conn = conn
	client.mu.Unlock()

	var wg sync.WaitGroup

	// Writer goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < messageCount; i++ {
			msg := fmt.Sprintf(`{"test":"%d"}`, i)
			if err := client.Send([]byte(msg)); err != nil {
				t.Logf("Send error: %v", err)
				return
			}
			sentMessages <- struct{}{}
			time.Sleep(10 * time.Millisecond)
		}
	}()

	// Reader goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < messageCount; i++ {
			_, data, err := client.Read()
			if err != nil {
				t.Logf("Read error: %v", err)
				return
			}
			receivedMessages <- string(data)
		}
	}()

	wg.Wait()
	close(receivedMessages)
	close(sentMessages)

	// Verify we sent all messages
	sentCount := len(sentMessages)
	if sentCount != messageCount {
		t.Errorf("Expected to send %d messages, sent %d", messageCount, sentCount)
	}

	// Verify we received all responses
	receivedCount := len(receivedMessages)
	if receivedCount != messageCount {
		t.Errorf("Expected to receive %d messages, received %d", messageCount, receivedCount)
	}
}

// ============================================================================
// Integration-like Tests
// ============================================================================

func TestWSClient_SubscribeAndReceivePriceData(t *testing.T) {
	// Mock server that sends price updates after subscription
	server := newMockWSServer(t, func(conn *websocket.Conn) {
		// Read subscription message
		_, _, err := conn.ReadMessage()
		if err != nil {
			return
		}

		// Send price updates
		for i := 0; i < 3; i++ {
			priceData := WsData{
				Type: WsDataPriceData,
				Data: json.RawMessage(fmt.Sprintf(`{"c":%.2f,"address":"SOL","unixTime":%d}`, 100.0+float64(i), time.Now().Unix())),
			}
			priceJSON, _ := json.Marshal(priceData)
			conn.WriteMessage(websocket.TextMessage, priceJSON)
			time.Sleep(50 * time.Millisecond)
		}
	})
	defer server.Close()

	wsURL := "ws" + strings.TrimPrefix(server.URL, "http")
	ctx := context.Background()

	dialer := websocket.Dialer{}
	conn, _, err := dialer.DialContext(ctx, wsURL, nil)
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := NewWSClient(WSClientConfig{
		APIKey: "test-key",
		Chain:  ChainSolana,
	})
	client.mu.Lock()
	client.conn = conn
	client.mu.Unlock()

	// Subscribe
	sub := SubDataPrice{
		Address:   "SOL",
		ChartType: WsInterval1m,
		Currency:  CurrencyUSD,
		QueryType: QueryTypeSimple,
	}
	payload, _ := sub.Payload()
	if err := client.Subscribe(payload); err != nil {
		t.Fatalf("Subscribe error: %v", err)
	}

	// Read price updates
	prices := make([]float64, 0, 3)
	for i := 0; i < 3; i++ {
		msgType, data, err := client.Read()
		if err != nil {
			t.Fatalf("Read error: %v", err)
		}
		if msgType != WsDataPriceData {
			t.Errorf("Expected PRICE_DATA, got %s", msgType)
			continue
		}

		var priceData WsDataPrice
		if err := json.Unmarshal(data, &priceData); err != nil {
			t.Fatalf("Unmarshal error: %v", err)
		}
		prices = append(prices, priceData.C)
	}

	// Verify we got 3 price updates
	if len(prices) != 3 {
		t.Errorf("Expected 3 price updates, got %d", len(prices))
	}

	// Verify prices are incrementing
	for i := 0; i < len(prices); i++ {
		expected := 100.0 + float64(i)
		if prices[i] != expected {
			t.Errorf("Expected price %.2f, got %.2f", expected, prices[i])
		}
	}
}

// ============================================================================
// Helper Functions
// ============================================================================

func stringPtr(s string) *string {
	return &s
}

func getTestWSClient() *WSClient {
	apiKey := os.Getenv("BIRDEYE_API_KEY")
	if apiKey == "" {
		panic("BIRDEYE_API_KEY environment variable not set")
	}
	return NewWSClient(WSClientConfig{
		APIKey: apiKey,
		Chain:  ChainSolana,
	})
}

func TestRealWsClient(t *testing.T) {
	client := getTestWSClient()
	ctx := context.Background()
	if err := client.Connect(ctx); err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer client.Close()

	msgType, _, err := client.Read()
	if err != nil {
		t.Fatalf("Read error: %v", err)
	}
	if msgType != WsDataWelcome {
		t.Fatalf("Expected WELCOME message, got %s", msgType)
	}
	sub := SubDataPrice{
		Address:   "So11111111111111111111111111111111111111112",
		ChartType: WsInterval1m,
		Currency:  CurrencyUSD,
		QueryType: QueryTypeSimple,
	}
	payload, _ := sub.Payload()
	if err := client.Subscribe(payload); err != nil {
		t.Fatalf("Subscribe error: %v", err)
	}

	msgType, data, err := client.Read()
	if err != nil {
		t.Fatalf("Read error: %v", err)
	}
	if msgType == WsDataPriceData {
		var priceData WsDataPrice
		if err := json.Unmarshal(data, &priceData); err != nil {
			t.Fatalf("Unmarshal error: %v", err)
		}
		if priceData.C <= 0 {
			t.Fatalf("Expected price > 0, got %.2f", priceData.C)
		}
		if priceData.UnixTime <= 0 {
			t.Fatalf("Expected UnixTime > 0, got %d", priceData.UnixTime)
		}
	} else {
		t.Fatalf("Expected PRICE_DATA message, got %s", msgType)
	}

}
