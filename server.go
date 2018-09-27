package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/t3h2mas/passpi/hash"
)

type server struct {
	hash   hash.HashService
	router *http.ServeMux
	stop   chan bool
}

func httpErr(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (s *server) handleHash() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			httpErr(w, http.StatusMethodNotAllowed)
			return
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("Failed to parse body, error: %s\n", err.Error())
			httpErr(w, http.StatusBadRequest)
			return
		}

		// body must not be empty
		if len(body) == 0 {
			log.Println("Body empty")
			httpErr(w, http.StatusBadRequest)
			return
		}

		pieces := strings.Split(string(body), "=")
		if len(pieces) != 2 {
			log.Printf("Bad body format:\n%s\n", string(body))
			httpErr(w, http.StatusBadRequest)
			return
		}

		if pieces[0] != "password" {
			log.Printf("Bad body format:\n%s\n", string(body))
			httpErr(w, http.StatusBadRequest)
			return
		}

		if len(pieces[1]) == 0 {
			log.Printf("Password field empty:\n%s\n", string(body))
			httpErr(w, http.StatusBadRequest)
			return
		}

		// return hash
		fmt.Fprintf(w, "%s", s.hash.Calculate(pieces[1]))
	}
}

func (s *server) handleShutdown() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Shutdown triggered")
		s.stop <- true
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

// decorates router
func (s *server) logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}
