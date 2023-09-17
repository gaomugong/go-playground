package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	. "server"
	"strings"
)

type CLI struct {
	store PlayerStore
	//in    io.Reader
	in *bufio.Scanner
}

func NewCLI(store PlayerStore, in io.Reader) *CLI {
	return &CLI{store: store, in: bufio.NewScanner(in)}
}

func (c *CLI) PlayPoker() {
	//name := make([]byte, 100)
	//c.in.Read(name)
	//players := strings.Split(string(name), " ")
	//fmt.Printf("players=%#v\n", players)
	userInput := c.readLine()
	c.store.RecordWin(extractWinner(userInput))
}

func extractWinner(input string) string {
	return strings.Replace(input, " win", "", -1)
}

func (c *CLI) readLine() string {
	c.in.Scan()
	return c.in.Text()
}

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
