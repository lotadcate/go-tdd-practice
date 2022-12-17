package iteration
import "testing"

func TestIteration(t *testing.T) {
  repeated := Repeat("a", 6)
  expected := "aaaaaa"
  if repeated != expected {
    t.Errorf("expected %q but got %q", expected, repeated)
  }
}

// go test -bench=.
func BenchmarkRepeat(b *testing.B) {
  for i:=0; i<b.N; i++ {
    Repeat("a", 6)
  }
}
