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

func TestPlayer(t *testing.T) {
	t.Run("return Peter's score", func(t *testing.T) {
		request := newGetScoreRequest("Petter")
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		assertResponse(t, response.Body.String(), "20")

	})

	t.Run("return Floyd's score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		assertResponse(t, response.Body.String(), "10")
	})
}
