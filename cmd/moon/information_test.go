package moon

import (
  "testing"
)

func TestInformation(t *testing.T) {
  result := information(8)

  if result != "Full moon" {
    t.Errorf("Result was incorrect, got: %s, want: %s.", result, "Full moon")
  }
}
