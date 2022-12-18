package arrayandslices
import (
  "testing"
  "reflect"
)

func TestSumAllTails(t *testing.T) {
  assert := func(t testing.TB, got, want []int) {
    t.Helper()
    if !reflect.DeepEqual(got, want) {
      t.Errorf("got %v want %v", got, want)
    }
  }

  t.Run("make the sums of two slices which contain two elements", func(t *testing.T) {
    got := SumAllTails([]int{1, 2}, []int{0, 9})
    want := []int{2, 9}
    assert(t, got, want)
  })

  t.Run("make the sums of three slices which contain three elements", func(t *testing.T) {
    got := SumAllTails([]int{1, 2, 3}, []int{0, 9, 8}, []int{9, 7, 3})
    want := []int{5, 17, 10}
    assert(t, got, want)
  })

  t.Run("safely sum empty slices", func(t *testing.T) {
    got := SumAllTails([]int{}, []int{3, 4, 5})
    want := []int{0, 9}
    assert(t, got, want)
  })
}
