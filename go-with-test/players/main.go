package main

import (
	"log"
	"net/http"

	. "server"
)

func main() {
	server := &PlayerServer{}
	store := NewInMemoryPlayerScore()
	server.SetStore(store)
	if err := http.ListenAndServe(":8080", server); err != nil {
		log.Fatalf("error listening on port 8080: %v", err)
	}
}
