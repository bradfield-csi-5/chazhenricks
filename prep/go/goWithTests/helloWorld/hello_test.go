package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello to name passed in", func(t *testing.T) {
		got := Hello("Chaz")
		want := "Hello, Chaz"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello world when no argument passed", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q - want %q", got, want)
	}
}
