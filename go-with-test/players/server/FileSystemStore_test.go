package server

import (
	"io"
	"strings"
	"testing"
)

type FileSystemStore struct {
	//database io.Reader
	database io.ReadSeeker
}

func (f *FileSystemStore) GetPlayerScore(name string) int {
	for _, player := range f.GetLeague() {
		if player.Name == name {
			return player.Wins
		}
	}
	return 0
}

func (f *FileSystemStore) RecordWin(name string) {
	//TODO implement me
	panic("implement me")
}

func (f *FileSystemStore) GetLeague() []Player {
	f.database.Seek(0, 0)
	league, _ := NewLeague(f.database)
	return league
}

func TestFileSystemStore(t *testing.T) {
	t.Run("/league from a reader", func(t *testing.T) {
		database := strings.NewReader(`[
			{"name": "Cleo", "Wins": 10},
			{"name": "Chris", "Wins": 33}]`)

		store := FileSystemStore{database}

		got := store.GetLeague()
		want := []Player{
			{"Cleo", 10},
			{"Chris", 33},
		}

		// read again
		got = store.GetLeague()
		assertLeague(t, want, got)
	})

	t.Run("get player score", func(t *testing.T) {
		database := strings.NewReader(`[
			{"name": "Cleo", "Wins": 10},
			{"name": "Chris", "Wins": 33}]`)

		store := FileSystemStore{database}

		got := store.GetPlayerScore("Chris")
		want := 33
		assertScoreEquals(t, want, got)
	})
}

func assertScoreEquals(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
