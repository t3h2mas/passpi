package stats

import (
	"encoding/json"
	"sync/atomic"
	"time"
)

// Service provides an interface for http statistics
type Service interface {
	AddPoint(time.Duration)
	JSON() ([]byte, error)
}

// Memory holds endpoint statistics in memory
type Memory struct {
	RequestCount int64
	TotalTime    int64
}

// private, used for marshalling
type statsResult struct {
	Total   int64 `json:"total"`
	Average int64 `json:"average"`
}

// JSON marshals the Stats struct for web responses
func (s *Memory) JSON() ([]byte, error) {
	var result *statsResult
	if s.RequestCount == 0 {
		// default, avoid division by zero
		result = &statsResult{0, 0}
	} else {
		// get average in microseconds
		avg := (s.TotalTime / s.RequestCount) / int64(time.Microsecond)

		result = &statsResult{
			Total:   s.RequestCount,
			Average: avg,
		}
	}
	return json.Marshal(result)
}

// AddPoint increments the running metrics
func (s *Memory) AddPoint(t time.Duration) {
	// use atomic pkg to prevent racing
	atomic.AddInt64(&s.RequestCount, 1)
	atomic.AddInt64(&s.TotalTime, int64(t))
}
