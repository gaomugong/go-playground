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

func assertError(t *testing.T, want, got error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %#v, want %#v", got, want)
	}
}

func TestFind(t *testing.T) {
	dict := Dict{"test": testStr}

	t.Run("test: found", func(t *testing.T) {
		got, _ := dict.Find("test")
		want := testStr
		assertString(t, got, want)
	})

	t.Run("unknown: not found", func(t *testing.T) {
		_, err := dict.Find("unknown")

		if err == nil {
			t.Fatal("expected to get an error")
		}
		assertError(t, err, DictKeyError)
	})

}
