package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/t3h2mas/passpi/hash"
)

func main() {
	// "global" configuration
	Addr := ":8080"
	shutdownTimeout := time.Second * 10
	routeDelay := time.Second * 5

	// initialize vars
	hash := &hash.HashSha512{}

	// used for gracefull shutdown
	stop := make(chan bool, 1)

	// main server struct, holds router and dependencies
	server := &server{
		hash:   hash,
		router: http.NewServeMux(),
		stop:   stop,
		stats: &Stats{
			requestCount: 0,
			totalTime:    0,
		},
	}

	// register routes
	server.routes(routeDelay)

	// implement http.Server so we can use `s` for shutdown
	s := &http.Server{
		Addr:    Addr,
		Handler: server.logMiddleware(server.router),
	}

	go func() {
		log.Printf("Listening on http://0.0.0.0%s\n", s.Addr)
		if err := s.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	// blocks on stop channel until shutdown endpoint called
	<-stop
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	log.Printf("\nShutdown with timeout: %s\n", shutdownTimeout)
	if err := s.Shutdown(ctx); err != nil {
		log.Printf("Error during shutdown: %s\n", err.Error())
	} else {
		log.Println("Server stopped")
	}
}
