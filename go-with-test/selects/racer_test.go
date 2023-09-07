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
	slowServer := makeDelayedServer(time.Millisecond * 20)
	fastServer := makeDelayedServer(time.Millisecond * 0)

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

	slowServer.Close()
	fastServer.Close()
}

// makeDelayedServer 将一些不感兴趣的代码移出测试并减少了重复代码
func makeDelayedServer(duration time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(duration)
		w.WriteHeader(http.StatusOK)
	}))
}
