package server

type InMemoryPlayerScore struct {
	scores map[string]int
}

func (i *InMemoryPlayerScore) GetLeague() []Player {
	return nil
}

func (i *InMemoryPlayerScore) RecordWin(name string) {
	i.scores[name]++
}

func (i *InMemoryPlayerScore) GetPlayerScore(name string) int {
	return i.scores[name]
}

func NewInMemoryPlayerScore() *InMemoryPlayerScore {
	return &InMemoryPlayerScore{map[string]int{}}
}
