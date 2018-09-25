package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
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

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Bad request")
			return
		}

		// body must not be empty
		if len(body) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Bad request")
			return
		}

		pieces := strings.Split(string(body), "=")
		if len(pieces) != 2 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Bad request")
			return
		}

		if pieces[0] != "password" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Bad request")
			return
		}

		if len(pieces[1]) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Bad request")
			return
		}

		// return hash
		fmt.Fprintf(w, "%s", s.hash.Calculate(pieces[1]))
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
