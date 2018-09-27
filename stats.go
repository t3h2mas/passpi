package main

import (
	"encoding/json"
	"fmt"
	"sync/atomic"
	"time"
)

// Stats holds endpoint statistics
type Stats struct {
	requestCount int64
	totalTime    int64
}

type statsResult struct {
	Total   int64 `json:"total"`
	Average int64 `json:"average"`
}

func (s *Stats) Json() ([]byte, error) {
	var result *statsResult
	if s.requestCount == 0 {
		result = &statsResult{0, 0}
	} else {
		result = &statsResult{
			Total:   s.requestCount,
			Average: (s.totalTime / s.requestCount) / int64(time.Microsecond),
		}
	}
	fmt.Printf("total: %d, count: %d, avg?: %d\n", s.totalTime, s.requestCount, s.totalTime/s.requestCount)
	return json.Marshal(result)
}

func (s *Stats) AddPoint(t time.Duration) {
	atomic.AddInt64(&s.requestCount, 1)
	atomic.AddInt64(&s.totalTime, int64(t))
}
