package structsmethodsinterfaces
import "testing"

func TestPerimeter(t *testing.T) {
  rectangle := Rectangle{10.0, 10.0}
  got := Perimeter(rectangle)
  want := 40.0
  if got != want {
    t.Errorf("got %.2f want %.2f", got, want)
  }
}

func TestArea(t *testing.T) {
  assert := func(t testing.TB, got, want float64) {
    t.Helper()
    if got != want {
      t.Errorf("got %g want %g", got, want)
    }
  }

  t.Run("rectangles", func(t *testing.T) {
    rectangle := Rectangle{12.0, 6.0}
    got := Area(rectangle)
    want := 72.0
    assert(t, got, want)
  })

  t.Run("circles", func(t *testing.T) {
    circle := Circle{10}
    got := Area(circle)
    want := 314.1592653589793
    assert(t, got, want)
  })
}
