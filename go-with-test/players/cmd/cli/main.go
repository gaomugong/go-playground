package main

import (
	"fmt"
	"log"
	"os"
	. "players"
)

const dbFileName = "game.db.json"

func main() {
	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} win to record a win}")

	//dbFile, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	//if err != nil {
	//	log.Fatalf("failed to open database file %s: %v", dbFileName, err)
	//}

	//store, err := NewFileSystemStore(dbFile)
	store, err := FileSystemStoreFromFile(dbFileName)
	if err != nil {
		//log.Fatalf("failed to create store: %v", err)
		log.Fatal(err)
	}

	game := NewCLI(store, os.Stdin)
	game.PlayPoker()

	league := store.GetLeague()
	fmt.Println(league)
}
