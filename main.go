package main

import (
	"log"
	"net/http"

	"github.com/t3h2mas/passpi/hash"
)

func main() {
	hash := &hash.HashSha512{}

	server := &server{
		hash:   hash,
		router: http.NewServeMux(),
	}

	server.routes()

	err := http.ListenAndServe(":8080", server.logMiddleware(server.router))
	if err != nil {
		log.Fatal(err)
	}
}
