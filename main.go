package main

import (
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

	http.ListenAndServe(":8080", server.router)
}
