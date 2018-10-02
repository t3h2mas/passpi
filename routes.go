package main

import "time"

// register router handlers
func (s *server) routes(delay time.Duration) {
	s.router.HandleFunc("/hash", s.statsMiddleware(s.delayMiddleware(delay)(s.handleHash())))
	s.router.HandleFunc("/shutdown", s.handleShutdown())
	s.router.HandleFunc("/stats", s.handleStats())
}
