package server

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
)

type FileSystemStore struct {
	//database io.Reader
	//database io.ReadSeeker
	//database io.ReadWriteSeeker
	//database io.Writer
	database *json.Encoder
	league   League
}

func initialPlayerDBFile(file *os.File) error {
	info, err := file.Stat()
	if err != nil {
		return fmt.Errorf("get info from file %s failed: %v", file.Name(), err)
	}

	if info.Size() == 0 {
		file.Write([]byte("[]"))
		file.Seek(0, io.SeekStart)
	}
	return nil
}

func NewFileSystemStore(file *os.File) (*FileSystemStore, error) {
	file.Seek(0, io.SeekStart)

	//info, err := file.Stat()
	//if err != nil {
	//	return nil, fmt.Errorf("get info from file %s failed: %v", file.Name(), err)
	//}
	err := initialPlayerDBFile(file)
	if err != nil {
		return nil, fmt.Errorf("init db file %s failed: %v", file.Name(), err)
	}
	//
	//if info.Size() == 0 {
	//	file.Write([]byte("[]"))
	//	file.Seek(0, io.SeekStart)
	//}

	league, err := NewLeague(file)
	//return &FileSystemStore{file: file, league: league}
	//return &FileSystemStore{file: &tape{file}, league: league}
	if err != nil {
		return nil, fmt.Errorf("problem loading player store from file %s: %s", file.Name(), err)
	}
	return &FileSystemStore{
		database: json.NewEncoder(&tape{file}),
		league:   league,
	}, nil
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
	sort.Slice(f.league, func(i, j int) bool {
		return f.league[i].Wins > f.league[j].Wins
	})
	return f.league
}
