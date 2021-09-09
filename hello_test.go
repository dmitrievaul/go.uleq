package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := ""
		want := ""
		assertCorrectMessage(t, got, want)
	})
	t.Run("empty string defaults to 'World'", func(t *testing.T) {
		got := ""
		want := ""
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Bonjour", "French")
		want := "Bonjour, Bonjour"
		assertCorrectMessage(t, got, want)
	})
	t.Run("russiancold", func(t *testing.T) {
		got := Hello("Привет", "Russian")
		want := "Привет, Привет"
		assertCorrectMessage(t, got, want)
	})
}
