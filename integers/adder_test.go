package integers

import "testing"

func TestAdder(t *testing.T) {
  assert := func(t testing.TB, sum, expected int) {
    t.Helper()
    if sum != expected {
      t.Errorf("expected '%d' but got '%d'", expected, sum)
    }
  }

  t.Run("2+7", func(t *testing.T) {
    sum := Add(2, 7)
    expected := 9
    assert(t, sum, expected)
  })
}
