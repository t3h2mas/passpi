package main

import "time"

func (s *server) routes(delay time.Duration) {
	s.router.HandleFunc("/hashQuick", s.handleHash())
	s.router.HandleFunc("/hash", s.withStats(s.delayMiddleware(delay)(s.handleHash())))
	s.router.HandleFunc("/shutdown", s.handleShutdown())
	s.router.HandleFunc("/stats", s.handleStats())
}
