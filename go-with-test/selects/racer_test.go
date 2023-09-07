package selects

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	// v1
	// slowURL := "http://www.facebook.com"
	// fastURL := "http://www.quii.co.uk"

	// v2
	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Millisecond * 20)
		w.WriteHeader(http.StatusOK)
	}))

	fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	// want := fastURL
	want := fastServer.URL

	// v1
	// got := Racer(slowURL, fastURL)

	// v2
	fmt.Println(slowServer.URL, fastServer.URL)
	got := Racer(slowServer.URL, fastServer.URL)

	if want != got {
		t.Errorf("Want %#v, got %#v", want, got)
	}
}
