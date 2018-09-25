package main

import (
	"fmt"
	"net/http"

	"github.com/t3h2mas/passpi/hash"
)

type server struct {
	hash   hash.HashService
	router *http.ServeMux
}

func (s *server) handleHash() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	}
}
