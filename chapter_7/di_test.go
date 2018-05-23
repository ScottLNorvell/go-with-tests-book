package main

import (
  "bytes"
  "testing"
)

func TestEncourage(t *testing.T) {
  buffer := bytes.Buffer{}
  Encourage(&buffer, "Scott")

  got := buffer.String()
  want := "Scott is Great!"

  if got != want {
    t.Errorf("got '%s' want '%s'", got, want)
  }
}
