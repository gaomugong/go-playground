package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/hello", hello)
	fmt.Println("Listening on port 8080...")
	_ = http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {

	// A `context.Context` is created for each request by
	// the `net/http` machinery, and is available with
	// the `Context()` method.
	ctx := r.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	// simulate some work and keep an eye on the ctx's `Done()` channel
	// for a signal that we should cancel the work and return
	select {
	case <-time.After(10 * time.Second):
		_, _ = fmt.Fprint(w, "hello\n")
	case <-ctx.Done():
		// `Err()` returns an error that explains why `Done()` channel is closed
		err := ctx.Err()
		fmt.Println("server: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// open http://localhost:8080/hello
// server: hello handler started
// server: hello handler ended

// open http://localhost:8080/hello
// server: hello handler started
// ctrl+w -> close http://localhost:8080/hello
// server:  context canceled
