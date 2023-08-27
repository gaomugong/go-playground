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

//func TestHello(t *testing.T) {
//
//	assertCorrectMessage := func(t *testing.T, got, want string) {
//		if got != want {
//			t.Errorf("got %q, want %q", got, want)
//		}
//	}
//	t.Run("say hello to people", func(t *testing.T) {
//		got := Hello("pitou")
//		want := englishPrefix + "pitou"
//		assertCorrectMessage(t, got, want)
//	})
//
//	t.Run("say hello world when name is empty", func(t *testing.T) {
//		got := Hello("")
//		want := englishPrefix + "world"
//		assertCorrectMessage(t, got, want)
//	})
//}

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
}
