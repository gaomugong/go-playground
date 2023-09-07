package v1

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		slowServer := makeDelayedServer(time.Second * 12)
		fastServer := makeDelayedServer(time.Second * 11)

		// 把它放在你创建服务器语句附近，以便函数内后面的代码仍可以使用这个服务器
		defer slowServer.Close()
		defer fastServer.Close()

		_, err := Racer(slowServer.URL, fastServer.URL)
		if err == nil {
			t.Errorf("expected an error but didn't get one")
		}
	})

}

// makeDelayedServer 将一些不感兴趣的代码移出测试并减少了重复代码
func makeDelayedServer(duration time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(duration)
		w.WriteHeader(http.StatusOK)
	}))
}
