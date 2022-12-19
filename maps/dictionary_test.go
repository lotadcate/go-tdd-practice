package dictionary
import (
  "testing"
)

func TestSearch(t *testing.T) {
  dictionary := Dictionary {
    "test": "this is just a test",
  }

  // expect to get "string"
  t.Run("known world", func(t *testing.T) {
    got, _  := dictionary.Search("test")
    want := "this is just a test"
    assertStrings(t, got, want)
  })

  // expect to get "error"
  t.Run("unknown world", func(t *testing.T) {
    _, got := dictionary.Search("unknown")
    assertError(t, got, ErrNotFound)
  })
}

func assertStrings(t *testing.T, got, want string) {
  t.Helper()
  if got != want {
    t.Errorf("got %q want %q", got, want)
  }
}

func assertError(t *testing.T, got, want error) {
    t.Helper()
    if got != want {
      t.Errorf("got %q want %q", got, want)
    }
}
