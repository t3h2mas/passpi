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

	hash := &hash.HashSha512{}

	stop := make(chan bool, 1)

	server := &server{
		hash:   hash,
		router: http.NewServeMux(),
		stop:   stop,
	}

	server.routes(routeDelay)

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
