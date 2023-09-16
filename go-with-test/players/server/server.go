package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PlayerStore interface {
	GetPlayerScore(player string) int
	RecordWin(name string)
	//GetLeague() []Player
	GetLeague() League
}

type PlayerServer struct {
	store PlayerStore
	// router *http.ServeMux
	// 嵌入：PlayerServer拥有了http.Handler的所有方法，即 ServeHTTP
	// 在使用嵌入接口的方式时，需要确保实现了接口中的所有方法
	http.Handler
}

type Player struct {
	Name string
	Wins int
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := new(PlayerServer)
	p.store = store

	// p := &PlayerServer{store: store, router: http.NewServeMux()}
	// p.router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	// p.router.Handle("/players", http.HandlerFunc(p.playerHandler))

	// router为 ServeMux 类型（实现了http.Handler接口）
	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	// 不要漏了结尾的/
	router.Handle("/players/", http.HandlerFunc(p.playerHandler))
	p.Handler = router

	return p
}

func (p *PlayerServer) SetStore(store PlayerStore) {
	p.store = store
}

// func (p *PlayerServer) getLeagueTable() []Player {
// 	return []Player{
// 		{"Chris", 20},
// 	}
// }

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	// leagueTable := []Player{
	// 	{"Chris", 20},
	// }
	// leagueTable := p.getLeagueTable()
	_ = json.NewEncoder(w).Encode(p.store.GetLeague())
	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
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

// 这允许我们删除我们的 ServeHTTP 方法，因为我们已经通过嵌入类型http.Handler公开了它。
// func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	p.router.ServeHTTP(w, r)
// }

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
