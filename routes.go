package main

import "time"

func (s *server) routes() {
	s.router.HandleFunc("/hash", s.handleHash())
	s.router.HandleFunc("/delay", s.delayMiddleware(5*time.Second)(s.handleHash()))
}
