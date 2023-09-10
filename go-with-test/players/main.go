package main

import (
	"log"
	"net/http"
	. "server"
)

func main() {
	handler := http.HandlerFunc(PlayerServer)
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatalf("error listening on port 8080: %v", err)
	}
}
