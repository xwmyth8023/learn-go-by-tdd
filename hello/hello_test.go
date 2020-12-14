package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("sandy", "english")
		want := "Hello, sandy"
		assertCorrectMessage(t, got, want)
		// if got != want {
		// 	t.Errorf("got '%q' want '%q'", got, want)
		// }
	})

	t.Run("saying hello world when input empty name", func(t *testing.T) {
		got := Hello("", "english")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
		// if got != want {
		// 	t.Errorf("got '%q' want '%q'", got, want)
		// }
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("World", "French")
		want := "Bonjour, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("when language is empty", func(t *testing.T) {
		got := Hello("World", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}
