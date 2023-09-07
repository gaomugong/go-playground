package concurrency

import (
	"reflect"
	"testing"
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
