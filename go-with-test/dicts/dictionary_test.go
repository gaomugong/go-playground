package dicts

import "testing"

const testStr = "this is just a test"

func TestSearch(t *testing.T) {
	dict := Dict{
		"test": testStr,
	}

	got := dict.Search("test")
	want := testStr

	assertString(t, got, want)
}

func assertString(t *testing.T, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %#v, want %#v, given %#v", got, want, "test")
	}
}
