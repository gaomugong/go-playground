package players

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPlayer(t *testing.T) {
	t.Run("players", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Petter", nil)
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		got := response.Body.String()
		want := "20"

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}

	})
}
