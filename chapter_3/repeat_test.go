package iteration

import (
  "fmt"
  "testing"
)

func TestRepeat(t *testing.T) {
  repeated := Repeat("a", 7)
  expected := "aaaaaaa"

  if repeated != expected {
    t.Errorf("expected '%s' but got  '%s'", expected, repeated)
  }
}

func TestRepeatAnotherWay(t *testing.T) {
  repeated := RepeatAnotherWay("a", 7)
  expected := "aaaaaaa"

  if repeated != expected {
    t.Errorf("expected '%s' but got  '%s'", expected, repeated)
  }
}

func BenchmarkRepeat(b *testing.B) {
  for i := 0; i < b.N; i++ {
    Repeat("a", 6)
  }
}

func ExampleRepeat() {
  result := Repeat("a", 5)
  fmt.Println(result)
  // Output: aaaaa
}
