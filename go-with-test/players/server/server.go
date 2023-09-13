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
	store  PlayerStore
	router *http.ServeMux
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := &PlayerServer{store: store, router: http.NewServeMux()}
	p.router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	p.router.Handle("/players", http.HandlerFunc(p.playerHandler))

	return p
}

func (p *PlayerServer) SetStore(store PlayerStore) {
	p.store = store
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (p *PlayerServer) playerHandler(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]

	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
		return
	case http.MethodGet:
		p.showScore(w, player)
	}
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
	p.router.ServeHTTP(w, r)
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
