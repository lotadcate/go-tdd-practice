package arrayandslices
import "testing"

func TestSum(t *testing.T) {
  assert := func(t testing.TB, got, want int, arr [5]int) {
    t.Helper()
    if got != want {
      t.Errorf("got %d want %d given, %v", got, want, arr)
    }
  }

  t.Run("collection of 5 numbers", func(t *testing.T) {
    arr := [5]int{1, 2, 3, 4, 5}
    got := Sum(arr)
    want := 15
    assert(t, got, want, arr)
  })

  t.Run("collection of any size", func(t *testing.T){
    arr := []int{1, 2, 3}
    got := Sum(arr)
    want := 6
    assert(t, got, want, arr)
  })
}
