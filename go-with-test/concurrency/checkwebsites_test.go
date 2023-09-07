package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	actualResults := CheckWebsites(mockWebSiteChecker, websites)

	want := len(websites)
	got := len(actualResults)

	if want != got {
		t.Fatalf("Wanted %v, got %v", want, got)
	}

	expectedResults := map[string]bool{
		"http://google.com":          false,
		"http://blog.gypsydave5.com": false,
		"waat://furhurterwe.geds":    true,
	}

	if !reflect.DeepEqual(actualResults, expectedResults) {
		t.Fatalf("Want %v, got %v", expectedResults, actualResults)
	}
}

func mockWebSiteChecker(url string) bool {
	return url == "waat://furhurterwe.geds"
}

func slowStubWebsiteChecker(url string) bool {
	time.Sleep(time.Millisecond * 20)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)

	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
