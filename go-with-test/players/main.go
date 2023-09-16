package main

import (
	"log"
	"net/http"
	"os"

	. "server"
)

const dbFileName = "game.db.json"

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("failed to open database %s, error: %v", dbFileName, err)
	}

	server := &PlayerServer{}
	//store := NewInMemoryPlayerScore()
	store := NewFileSystemStore(db)
	server.SetStore(store)
	if err := http.ListenAndServe(":8080", server); err != nil {
		log.Fatalf("error listening on port 8080: %v", err)
	}
}
