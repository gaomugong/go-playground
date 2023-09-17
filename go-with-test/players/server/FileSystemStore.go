package server

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type FileSystemStore struct {
	//database io.Reader
	//database io.ReadSeeker
	//database io.ReadWriteSeeker
	//database io.Writer
	database *json.Encoder
	league   League
}

func NewFileSystemStore(database *os.File) *FileSystemStore {
	database.Seek(0, io.SeekStart)
	league, _ := NewLeague(database)
	//return &FileSystemStore{database: database, league: league}
	//return &FileSystemStore{database: &tape{database}, league: league}
	return &FileSystemStore{database: json.NewEncoder(&tape{database}), league: league}
}

type League []Player

func (players League) Find(name string) *Player {
	for i, player := range players {
		if player.Name == name {
			return &players[i]
		}
	}

	return nil
}

func (f *FileSystemStore) GetPlayerScore(name string) int {
	//for _, player := range f.GetLeague() {
	//	if player.Name == name {
	//		return player.Wins
	//	}
	//}
	//return 0
	player := f.GetLeague().Find(name)
	if player != nil {
		return player.Wins
	}
	return 0
}

func (f *FileSystemStore) RecordWin(name string) {
	//league := f.GetLeague()
	fmt.Printf("league: %#v\n", f.league)
	//for i, player := range league {
	//	if player.Name == name {
	//		//player.Wins++
	//		league[i].Wins++
	//		break
	//	}
	//}
	player := f.league.Find(name)
	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{Name: name, Wins: 1})
	}
	//f.database.Seek(0, 0)
	//json.NewEncoder(f.database).Encode(f.league)
	f.database.Encode(f.league)
}

// GetLeague func (f *FileSystemStore) GetLeague() []Player {
func (f *FileSystemStore) GetLeague() League {
	//f.database.Seek(0, 0)
	//league, _ := NewLeague(f.database)
	//return league
	return f.league
}
