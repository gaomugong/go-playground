package server

import "testing"

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	//league   []Player
	league League
}

func (s *StubPlayerStore) WinCalls() []string {
	return s.winCalls
}

func (s *StubPlayerStore) SetWinCalls(winCalls []string) {
	s.winCalls = winCalls
}

// func (s *StubPlayerStore) GetLeague() []Player {
func (s *StubPlayerStore) GetLeague() League {
	return s.league
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) GetPlayerScore(player string) int {
	return s.scores[player]
}

func AssertPlayerWin(t *testing.T, store *StubPlayerStore, winner string) {
	t.Helper()

	if len(store.WinCalls()) != 1 {
		t.Fatalf("got %d calls to RecordWin, want %d", len(store.WinCalls()), 1)
	}

	got := store.WinCalls()[0]
	if got != winner {
		t.Fatalf("did not store correct winner got %#v, want %#v", got, winner)
	}

}
