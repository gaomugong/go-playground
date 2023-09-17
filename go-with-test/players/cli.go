package players

import (
	"bufio"
	"io"
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
