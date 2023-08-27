package main

import "testing"

//func TestHello(t *testing.T) {
//	got := Hello()
//	want := "Hello world"
//
//	if got != want {
//		t.Errorf("got %q, want %q", got, want)
//	}
//}

//func TestHello(t *testing.T) {
//	got := Hello("pitou")
//	want := englishPrefix + "pitou"
//
//	if got != want {
//		t.Errorf("got %q, want %q", got, want)
//	}
//}

func TestHello(t *testing.T) {
	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("pitou")
		want := englishPrefix + "pitou"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("say hello world when name is empty", func(t *testing.T) {
		got := Hello("")
		want := englishPrefix + "world"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
