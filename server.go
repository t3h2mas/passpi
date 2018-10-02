package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/t3h2mas/passpi/hash"
	"github.com/t3h2mas/passpi/stats"
)

type server struct {
	hash   hash.HashService
	router *http.ServeMux
	stop   chan bool
	stats  stats.Service
}

func httpErr(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

/*
 * endpoint handlers
 */

func (s *server) handleHash() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			httpErr(w, http.StatusMethodNotAllowed)
			return
		}

		err := r.ParseForm()
		if err != nil {
			log.Printf("Failed to parse body, error: %s\n", err.Error())
			httpErr(w, http.StatusBadRequest)
			return
		}

		password := r.FormValue("password")
		if len(password) == 0 {
			log.Println("Password field empty")
			httpErr(w, http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "%s", s.hash.Calculate(password))
		return
	}
}

func (s *server) handleShutdown() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Shutdown triggered")
		s.stop <- true
	}
}

func (s *server) handleStats() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		data, err := s.stats.JSON()
		if err != nil {
			fmt.Fprintf(w, "{\"error\": \"%s\"}", err.Error())
			return
		}
		w.Write(data)
	}
}

/*
 * middleware
 */

// adds metrics for wrapped handlers
func (s *server) statsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		defer func(start time.Time) {
			s.stats.AddPoint(time.Since(start))
		}(now)
		next.ServeHTTP(w, r)
	}
}

// returns middleware to add delay 'duration'
func (s *server) delayMiddleware(duration time.Duration) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(duration)
			next.ServeHTTP(w, r)
		}
	}
}

// decorates router
func (s *server) logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}
