package server

import (
	"io"
	"strings"
	"testing"
)

type FileSystemStore struct {
	database io.Reader
}

func (f *FileSystemStore) GetPlayerScore(name string) int {
	//TODO implement me
	panic("implement me")
}

func (f *FileSystemStore) RecordWin(name string) {
	//TODO implement me
	panic("implement me")
}

func (f *FileSystemStore) GetLeague() []Player {
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
}
