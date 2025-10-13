package birdeye

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestBasicRateLimiter(t *testing.T) {
	t.Run("Basic blocking behavior", func(t *testing.T) {
		limiter, err := NewRateLimiter(2, time.Second, RateLimitBlock)
		if err != nil {
			t.Fatal(err)
		}

		start := time.Now()
		ctx := context.Background()

		// First 2 calls should be immediate
		for range 2 {
			acquired, err := limiter.Acquire(ctx, 1, nil)
			if err != nil {
				t.Fatal(err)
			}
			if !acquired {
				t.Fatal("Expected to acquire token")
			}
		}

		// 3rd call should block
		acquired, err := limiter.Acquire(ctx, 1, nil)
		if err != nil {
			t.Fatal(err)
		}
		if !acquired {
			t.Fatal("Expected to acquire token after waiting")
		}

		elapsed := time.Since(start)
		// Should have waited approximately 0.5 seconds
		if elapsed < 400*time.Millisecond || elapsed > 700*time.Millisecond {
			t.Logf("Warning: Expected ~500ms wait, got %v", elapsed)
		}
	})

	t.Run("Skip behavior", func(t *testing.T) {
		limiter, err := NewRateLimiter(2, time.Second, RateLimitSkip)
		if err != nil {
			t.Fatal(err)
		}

		ctx := context.Background()
		skip := RateLimitSkip

		// First 2 should succeed
		for i := range 2 {
			acquired, err := limiter.Acquire(ctx, 1, &skip)
			if err != nil {
				t.Fatal(err)
			}
			if !acquired {
				t.Fatalf("Call %d should have been acquired", i+1)
			}
		}

		// 3rd should be skipped
		acquired, err := limiter.Acquire(ctx, 1, &skip)
		if err != nil {
			t.Fatal(err)
		}
		if acquired {
			t.Fatal("Call 3 should have been skipped")
		}
	})

	t.Run("Raise behavior", func(t *testing.T) {
		limiter, err := NewRateLimiter(1, time.Second, RateLimitRaise)
		if err != nil {
			t.Fatal(err)
		}

		ctx := context.Background()
		raise := RateLimitRaise

		// First call succeeds
		_, err = limiter.Acquire(ctx, 1, &raise)
		if err != nil {
			t.Fatal(err)
		}

		// Second call should raise error
		_, err = limiter.Acquire(ctx, 1, &raise)
		if err != ErrRateLimitExceeded {
			t.Fatalf("Expected ErrRateLimitExceeded, got %v", err)
		}
	})
}

func TestRateLimiterReset(t *testing.T) {
	limiter, err := NewRateLimiter(2, time.Second, RateLimitSkip)
	if err != nil {
		t.Fatal(err)
	}

	// Use up tokens
	limiter.TryAcquire(2)

	// Check no tokens available
	if limiter.TryAcquire(1) {
		t.Fatal("Should not have tokens available")
	}

	// Reset
	limiter.Reset()

	// Should have tokens again
	if !limiter.TryAcquire(1) {
		t.Fatal("Should have tokens after reset")
	}
}

func TestRateLimiterTokenRefill(t *testing.T) {
	limiter, err := NewRateLimiter(10, time.Second, RateLimitSkip)
	if err != nil {
		t.Fatal(err)
	}

	// Use up all tokens
	for i := range 10 {
		if !limiter.TryAcquire(1) {
			t.Fatalf("Failed to acquire token %d", i+1)
		}
	}

	// Should not have tokens
	if limiter.TryAcquire(1) {
		t.Fatal("Should not have tokens")
	}

	// Wait for refill
	time.Sleep(100 * time.Millisecond)

	// Should have at least 1 token (10 tokens per second = 1 per 100ms)
	if !limiter.TryAcquire(1) {
		t.Fatal("Should have refilled at least 1 token")
	}
}

func TestSharedRateLimiter(t *testing.T) {
	shared, err := NewSharedRateLimiter(5, time.Second, RateLimitSkip)
	if err != nil {
		t.Fatal(err)
	}

	// Simulate two functions sharing the same limiter
	successCount := 0

	// Try 10 calls total
	for range 10 {
		if shared.TryAcquire(1) {
			successCount++
		}
	}

	// Should have succeeded 5 times (the limit)
	if successCount != 5 {
		t.Fatalf("Expected 5 successful calls, got %d", successCount)
	}
}

func TestMultiRateLimiter(t *testing.T) {
	t.Run("Multi-tier enforcement", func(t *testing.T) {
		// 3 per second AND 5 per 2 seconds
		limiter, err := NewMultiRateLimiter(
			[]RateLimit{
				{Limit: 3, Period: time.Second},
				{Limit: 5, Period: 2 * time.Second},
			},
			RateLimitSkip,
		)
		if err != nil {
			t.Fatal(err)
		}

		// First 3 should succeed (within 1 second limit)
		for i := range 3 {
			if !limiter.TryAcquire(1) {
				t.Fatalf("Call %d should have succeeded", i+1)
			}
		}

		// 4th should fail (exceeds 1 second limit)
		if limiter.TryAcquire(1) {
			t.Fatal("Call 4 should have failed (exceeds 1s limit)")
		}

		// Check we have 2 tokens left in the 2-second limiter
		status := limiter.GetStatus()
		if len(status) != 2 {
			t.Fatalf("Expected 2 limiters, got %d", len(status))
		}

		// Verify the 2-second limiter has ~2 tokens left (5 - 3 used)
		if status[1].AvailableTokens < 1.5 || status[1].AvailableTokens > 2.5 {
			t.Logf("2-second limiter has %.2f tokens (expected ~2)", status[1].AvailableTokens)
		}
	})

	t.Run("Status reporting", func(t *testing.T) {
		limiter, err := NewMultiRateLimiter(
			[]RateLimit{
				{Limit: 10, Period: time.Second},
				{Limit: 50, Period: time.Minute},
			},
			RateLimitSkip,
		)
		if err != nil {
			t.Fatal(err)
		}

		// Make some calls
		for range 5 {
			limiter.TryAcquire(1)
		}

		// Check status
		status := limiter.GetStatus()
		if len(status) != 2 {
			t.Fatalf("Expected 2 limiters, got %d", len(status))
		}

		// First limiter should have ~5 tokens left
		if status[0].AvailableTokens < 4 || status[0].AvailableTokens > 6 {
			t.Logf("Warning: Expected ~5 tokens, got %.2f", status[0].AvailableTokens)
		}
	})
}

func TestConcurrentAccess(t *testing.T) {
	limiter, err := NewRateLimiter(100, time.Second, RateLimitSkip)
	if err != nil {
		t.Fatal(err)
	}

	var wg sync.WaitGroup
	successCount := int64(0)
	var mu sync.Mutex

	// Launch 200 goroutines trying to acquire
	for range 200 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if limiter.TryAcquire(1) {
				mu.Lock()
				successCount++
				mu.Unlock()
			}
		}()
	}

	wg.Wait()

	// Should have succeeded exactly 100 times
	if successCount != 100 {
		t.Fatalf("Expected 100 successful acquisitions, got %d", successCount)
	}
}

func TestContextCancellation(t *testing.T) {
	limiter, err := NewRateLimiter(1, time.Second, RateLimitBlock)
	if err != nil {
		t.Fatal(err)
	}

	// Use up the token
	limiter.TryAcquire(1)

	// Create a context that will be cancelled
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	start := time.Now()

	// This should block and then be cancelled
	_, err = limiter.Acquire(ctx, 1, nil)
	elapsed := time.Since(start)

	if err != context.DeadlineExceeded {
		t.Fatalf("Expected context.DeadlineExceeded, got %v", err)
	}

	// Should have waited approximately 100ms
	if elapsed < 80*time.Millisecond || elapsed > 150*time.Millisecond {
		t.Logf("Warning: Expected ~100ms wait, got %v", elapsed)
	}
}

// Example test to demonstrate usage
func ExampleRateLimiter_basic() {
	// Create a rate limiter: 5 calls per second
	limiter, _ := NewRateLimiter(5, time.Second, RateLimitBlock)

	ctx := context.Background()

	// Make calls
	for i := range 10 {
		acquired, err := limiter.Acquire(ctx, 1, nil)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		if acquired {
			fmt.Printf("Call %d: Success\n", i+1)
		}
	}
}

func ExampleMultiRateLimiter() {
	// Create a multi-tier rate limiter
	// 10 per second AND 100 per minute
	limiter, _ := NewMultiRateLimiter(
		[]RateLimit{
			{Limit: 10, Period: time.Second},
			{Limit: 100, Period: time.Minute},
		},
		RateLimitBlock,
	)

	ctx := context.Background()

	// Make API calls
	for i := range 15 {
		acquired, err := limiter.Acquire(ctx, 1, nil)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		if acquired {
			fmt.Printf("Call %d: Success\n", i+1)
		}
	}

	// Check status
	status := limiter.GetStatus()
	for i, s := range status {
		fmt.Printf("Limiter %d: %d/%v, %.1f tokens available\n",
			i+1, s.Limit, s.Period, s.AvailableTokens)
	}
}

func BenchmarkRateLimiter(b *testing.B) {
	limiter, _ := NewRateLimiter(1000000, time.Second, RateLimitSkip)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		limiter.TryAcquire(1)
	}
}

func BenchmarkRateLimiterParallel(b *testing.B) {
	limiter, _ := NewRateLimiter(1000000, time.Second, RateLimitSkip)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			limiter.TryAcquire(1)
		}
	})
}
