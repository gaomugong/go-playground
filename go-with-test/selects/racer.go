package selects

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(a string, b string) (winner string) {
	aDuration := measureResponseTime(a)
	bDuration := measureResponseTime(b)

	fmt.Println(aDuration, bDuration, aDuration > bDuration)
	if aDuration < bDuration {
		return a
	}

	return b
}

func measureResponseTime(url string) time.Duration {
	startA := time.Now()
	_, _ = http.Get(url)
	return time.Since(startA)
}
