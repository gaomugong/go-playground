package main

import (
	. "server"
	"strings"
	"testing"
)

func assertPlayerWin(t *testing.T, store *StubPlayerStore, winner string) {
	t.Helper()

	if len(store.WinCalls()) != 1 {
		t.Fatalf("got %d calls to RecordWin, want %d", len(store.WinCalls()), 1)
	}

	got := store.WinCalls()[0]
	if got != winner {
		t.Fatalf("did not store correct winner got %#v, want %#v", got, winner)
	}

}

func TestCLI(t *testing.T) {
	t.Run("record chris win from user input", func(t *testing.T) {
		in := strings.NewReader("Chris win\n")
		playerStore := &StubPlayerStore{}

		//cli := &CLI{playerStore, in}
		cli := NewCLI(playerStore, in)
		cli.PlayPoker()

		assertPlayerWin(t, playerStore, "Chris")
	})

	t.Run("record cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo win\n")
		playerStore := &StubPlayerStore{}

		//cli := &CLI{playerStore, in}
		cli := NewCLI(playerStore, in)
		cli.PlayPoker()

		assertPlayerWin(t, playerStore, "Cleo")
	})
}
