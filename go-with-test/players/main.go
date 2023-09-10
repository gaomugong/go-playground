package main

import (
	"log"
	"net/http"
	. "server"
)

type InMemoryPlayerScore struct {
	scores map[string]int
}

func (i *InMemoryPlayerScore) GetPlayerScore(player string) int {
	return 123
}

func main() {
	server := &PlayerServer{}
	server.SetStore(&InMemoryPlayerScore{})
	if err := http.ListenAndServe(":8080", server); err != nil {
		log.Fatalf("error listening on port 8080: %v", err)
	}
}
