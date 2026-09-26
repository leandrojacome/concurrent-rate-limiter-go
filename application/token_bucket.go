package application

import (
	"context"
	"sync"
	"time"

	"github.com/leandrojacome/concurrent-rate-limiter-go/domain"
)

type bucket struct {
	tokens  float64
	updated time.Time
}

type TokenBucket struct {
	mu              sync.Mutex
	capacity        int
	refillPerSecond float64
	clock           Clock
	buckets         map[string]bucket
}

func NewTokenBucket(capacity int, refillPerSecond float64, clock Clock) *TokenBucket {
	if capacity <= 0 || refillPerSecond <= 0 {
		panic("capacity and refill rate must be positive")
	}
	return &TokenBucket{capacity: capacity, refillPerSecond: refillPerSecond, clock: clock, buckets: make(map[string]bucket)}
}

func (l *TokenBucket) Allow(_ context.Context, key string) domain.Decision {
	l.mu.Lock()
	defer l.mu.Unlock()
	now := l.clock.Now()
	b, found := l.buckets[key]
	if !found {
		b = bucket{tokens: float64(l.capacity), updated: now}
	}
	elapsed := now.Sub(b.updated).Seconds()
	b.tokens = min(float64(l.capacity), b.tokens+elapsed*l.refillPerSecond)
	b.updated = now
	if b.tokens >= 1 {
		b.tokens--
		l.buckets[key] = b
		return domain.Decision{Allowed: true, Remaining: int(b.tokens)}
	}
	l.buckets[key] = b
	wait := time.Duration((1 - b.tokens) / l.refillPerSecond * float64(time.Second))
	return domain.Decision{Allowed: false, Remaining: 0, RetryAt: now.Add(wait)}
}
