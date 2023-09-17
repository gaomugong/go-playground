package main

import (
	"log"
	"net/http"
	. "server"
)

const dbFileName = "game.db.json"

func main() {
	//db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	//if err != nil {
	//	log.Fatalf("failed to open database %s, error: %v", dbFileName, err)
	//}

	server := &PlayerServer{}
	//store := NewInMemoryPlayerScore()
	//store, err := NewFileSystemStore(db)
	store, err := FileSystemStoreFromFile(dbFileName)
	if err != nil {
		//log.Fatalf("create player store failed: %v", err)
		log.Fatal(err)
	}

	server.SetStore(store)
	if err := http.ListenAndServe(":8080", server); err != nil {
		log.Fatalf("error listening on port 8080: %v", err)
	}
}
