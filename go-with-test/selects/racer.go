package selects

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(a string, b string) (winner string) {
	startA := time.Now()
	_, _ = http.Get(a)
	aDuration := time.Since(startA)

	startB := time.Now()
	_, _ = http.Get(b)
	bDuration := time.Since(startB)

	fmt.Println(aDuration, bDuration, aDuration > bDuration)
	if aDuration < bDuration {
		return a
	}

	return b
}
