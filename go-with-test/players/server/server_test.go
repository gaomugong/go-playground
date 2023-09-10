package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func assertResponse(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func newGetScoreRequest(player string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", player), nil)
	return request
}

type StubPlayerStore struct {
	scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(player string) int {
	return s.scores[player]
}

func TestGetPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Petter": 20,
			"Floyd":  10,
		},
	}
	server := &PlayerServer{&store}

	t.Run("return Peter's store", func(t *testing.T) {
		request := newGetScoreRequest("Petter")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponse(t, response.Body.String(), "20")

	})

	t.Run("return Floyd's store", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponse(t, response.Body.String(), "10")
	})
}
