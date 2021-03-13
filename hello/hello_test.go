package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		name := "たけし"
		got := Hello(name, "")
		want := "Hello, " + name
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'World'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Japanese", func(t *testing.T) {
		got := Hello("ゆい", "Japanese")
		want := "今日はゆい"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Amelia", "French")
		want := "Bonjour, Amelia"
		assertCorrectMessage(t, got, want)
	})
}
