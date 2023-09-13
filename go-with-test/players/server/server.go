package server

import (
	"fmt"
	"net/http"
)

type PlayerStore interface {
	GetPlayerScore(player string) int
	RecordWin(name string)
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) SetStore(store PlayerStore) {
	p.store = store
}

// func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	player := r.URL.Path[len("/players/"):]
//
//	if r.Method == http.MethodPost {
//		w.WriteHeader(http.StatusAccepted)
//		return
//	}
//
//	score := p.store.GetPlayerScore(player)
//
//	if score == 0 {
//		w.WriteHeader(http.StatusNotFound)
//		return
//	}
//
//	//w.WriteHeader(http.StatusOK)
//	fmt.Fprint(w, score)
// }

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router := http.NewServeMux()

	router.Handle("/league", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	router.Handle("/players", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		player := r.URL.Path[len("/players/"):]

		switch r.Method {
		case http.MethodPost:
			p.processWin(w, player)
			return
		case http.MethodGet:
			p.showScore(w, player)
		}
	}))

	router.ServeHTTP(w, r)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, score)
}
