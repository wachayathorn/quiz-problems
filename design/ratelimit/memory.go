package ratelimit

import (
	"sync"
	"time"
)

type RateLimiter struct {
	limit    int
	window   time.Duration
	requests map[string][]time.Time
	mu       sync.Mutex
}

func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		limit:    limit,
		window:   window,
		requests: make(map[string][]time.Time),
	}
}

func (r *RateLimiter) Allow(userID string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	now := time.Now()

	for i := range r.requests[userID] {
		if now.Sub(r.requests[userID][i]) > r.window {
			r.requests[userID] = r.requests[userID][i+1:]
			break
		}
	}

	if len(r.requests[userID]) < r.limit {
		r.requests[userID] = append(r.requests[userID], now)
		return true
	}

	return false
}
