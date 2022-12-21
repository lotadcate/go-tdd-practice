package main
import (
  "testing"
  "bytes"
)

func TestCountdown(t *testing.T) {
  buffer := &bytes.Buffer{}
  defaultSleeper := &DefaultSleeper{}

  Countdown(buffer, defaultSleeper)

  got := buffer.String()
  want := `3
2
1
Go!`

  if got != want {
    t.Errorf("got %q want %q", got, want)
  }

  if spySleeper.Calls != 4 {
    t.Errorf("not enough calls to sleeper, want 4 got %d", spySleeper.Calls)
  }
}
