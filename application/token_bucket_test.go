package application

import (
	"context"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/leandrojacome/concurrent-rate-limiter-go/domain"
)

type fakeClock struct{ now time.Time }

func (f *fakeClock) Now() time.Time { return f.now }

func TestEnforcesCapacityUnderConcurrency(t *testing.T) {
	clock := &fakeClock{now: time.Unix(0, 0)}
	limiter := NewTokenBucket(10, 1, clock)
	var allowed atomic.Int32
	var wg sync.WaitGroup
	for range 100 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if limiter.Allow(context.Background(), domain.ClientKey("client")).Allowed {
				allowed.Add(1)
			}
		}()
	}
	wg.Wait()
	if allowed.Load() != 10 {
		t.Fatalf("allowed=%d, want 10", allowed.Load())
	}
}

func TestRefillsOverTime(t *testing.T) {
	clock := &fakeClock{now: time.Unix(0, 0)}
	limiter := NewTokenBucket(1, 1, clock)
	if !limiter.Allow(context.Background(), domain.ClientKey("client")).Allowed {
		t.Fatal("first request should pass")
	}
	if limiter.Allow(context.Background(), domain.ClientKey("client")).Allowed {
		t.Fatal("second request should be denied")
	}
	clock.now = clock.now.Add(time.Second)
	if !limiter.Allow(context.Background(), domain.ClientKey("client")).Allowed {
		t.Fatal("request should pass after refill")
	}
}
