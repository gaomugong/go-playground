package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

func main() {
	client := http.Client{
		Timeout: time.Second * 5,
	}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "http://localhost:8000/", nil)
	if err != nil {
		log.Fatalf("new request error: %s", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("do request error: %s", err)
	}
	log.Printf("resp: %v", resp)
}
