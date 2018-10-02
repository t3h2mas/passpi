package main

import (
	"encoding/json"
	"sync/atomic"
	"time"
)

// Stats holds endpoint statistics
type Stats struct {
	requestCount int64
	totalTime    int64
}

// private, used for marshalling
type statsResult struct {
	Total   int64 `json:"total"`
	Average int64 `json:"average"`
}

// JSON marshals the Stats struct for web responses
func (s *Stats) JSON() ([]byte, error) {
	var result *statsResult
	if s.requestCount == 0 {
		// default, avoid division by zero
		result = &statsResult{0, 0}
	} else {
		// get average in microseconds
		avg := (s.totalTime / s.requestCount) / int64(time.Microsecond)

		result = &statsResult{
			Total:   s.requestCount,
			Average: avg,
		}
	}
	return json.Marshal(result)
}

// AddPoint increments the running metrics
func (s *Stats) AddPoint(t time.Duration) {
	// use atomic pkg to prevent racing
	atomic.AddInt64(&s.requestCount, 1)
	atomic.AddInt64(&s.totalTime, int64(t))
}
