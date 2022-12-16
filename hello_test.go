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
    got := hello("Chlis", "")
    want := "Hello, Chlis"
    assertCorrectMessage(t, got, want)
  })

  t.Run("say 'Hello, World' when no one is here", func(t *testing.T) {
    got := hello("", "");
    want := "Hello, World"
    assertCorrectMessage(t, got, want)
  })
  t.Run("in Spanish", func(t *testing.T) {
    got := hello("Elodie", "Spanish")
    want := "Hola, Elodie"
    assertCorrectMessage(t, got, want)
  })
  t.Run("in French", func(t *testing.T) {
    got := hello("Kevin", "French")
    want := "Bonjure, Kevin"
    assertCorrectMessage(t, got, want)
  })
}
