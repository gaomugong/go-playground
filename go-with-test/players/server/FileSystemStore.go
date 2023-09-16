package server

import (
	"encoding/json"
	"fmt"
	"io"
)

type FileSystemStore struct {
	//database io.Reader
	//database io.ReadSeeker
	database io.ReadWriteSeeker
}

func NewFileSystemStore(database io.ReadWriteSeeker) *FileSystemStore {
	return &FileSystemStore{database: database}
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
	league := f.GetLeague()
	fmt.Printf("league: %#v\n", league)
	//for i, player := range league {
	//	if player.Name == name {
	//		//player.Wins++
	//		league[i].Wins++
	//		break
	//	}
	//}
	player := league.Find(name)
	if player != nil {
		player.Wins++
	} else {
		league = append(league, Player{Name: name, Wins: 1})
	}
	f.database.Seek(0, 0)
	json.NewEncoder(f.database).Encode(league)
}

// GetLeague func (f *FileSystemStore) GetLeague() []Player {
func (f *FileSystemStore) GetLeague() League {
	f.database.Seek(0, 0)
	league, _ := NewLeague(f.database)
	return league
}
