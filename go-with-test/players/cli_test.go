package players

import (
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	t.Run("record chris win from user input", func(t *testing.T) {
		in := strings.NewReader("Chris win\n")
		playerStore := &StubPlayerStore{}

		//cli := &CLI{playerStore, in}
		cli := NewCLI(playerStore, in)
		cli.PlayPoker()

		AssertPlayerWin(t, playerStore, "Chris")
	})

	t.Run("record cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo win\n")
		playerStore := &StubPlayerStore{}

		//cli := &CLI{playerStore, in}
		cli := NewCLI(playerStore, in)
		cli.PlayPoker()

		AssertPlayerWin(t, playerStore, "Cleo")
	})
}
