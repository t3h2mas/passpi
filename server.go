package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/t3h2mas/passpi/hash"
)

type server struct {
	hash   hash.HashService
	router *http.ServeMux
}

func (s *server) handleHash() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Not allowed")
			return
		}
		fmt.Fprintf(w, "Hello, World!")
	}
}

func (s *server) delayMiddleware(duration time.Duration) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(duration)
			next.ServeHTTP(w, r)
		}
	}
}
