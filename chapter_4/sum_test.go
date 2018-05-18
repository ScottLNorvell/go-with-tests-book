package main

import (
  "fmt"
  "testing"
)

func assertSumsMatch(t *testing.T, got, want int, numbers []int) {
  t.Helper()
  if got != want {
    t.Errorf("got '%d', want '%d', %v", got, want, numbers)
  }
}

func TestSum(t *testing.T) {
  t.Run("collection of 5 numbers", func (t *testing.T) {
    numbers := []int{1,2,3,4,5}
    got := Sum(numbers)
    want := 15

    assertSumsMatch(t, got, want, numbers)
  })

  t.Run("collection of any size", func (t *testing.T) {
    numbers := []int{1,2,3}
    got := Sum(numbers)
    want := 6

    assertSumsMatch(t, got, want, numbers)
  })
}

func ExampleSum() {
  numbers := []int{5,10,15}
  sum := Sum(numbers)
  fmt.Println(sum)
  // Output: 30
}
