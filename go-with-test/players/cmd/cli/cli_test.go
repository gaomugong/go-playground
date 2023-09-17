package main

import (
	. "server"
	"testing"
)

type CLI struct {
	store PlayerStore
}

func (c *CLI) PlayPoker() {
	c.store.RecordWin("Cleo")
}

func TestCLI(t *testing.T) {
	playerStore := &StubPlayerStore{}
	cli := &CLI{playerStore}
	cli.PlayPoker()

	if len(playerStore.WinCalls()) != 1 {
		t.Fatalf("expected a win call but didn't get any")
	}
}
