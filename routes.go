package main

func (s *server) routes() {
	s.router.HandleFunc("/hash", s.handleHash())
}
