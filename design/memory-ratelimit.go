package design

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

func New() *RateLimiter {
	return &RateLimiter{
		limit:    3,
		window:   1 * time.Minute,
		requests: make(map[string][]time.Time),
		mu:       sync.Mutex{},
	}
}

func (r *RateLimiter) Allow(userID string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	now := time.Now()

	// Clean up old requests outside the window
	for i := range r.requests[userID] {
		if now.Sub(r.requests[userID][i]) > r.window {
			r.requests[userID] = r.requests[userID][i+1:]
			break
		}
	}

	// Check if under limit
	if len(r.requests[userID]) < r.limit {
		r.requests[userID] = append(r.requests[userID], now)
		return true
	}

	return false
}
