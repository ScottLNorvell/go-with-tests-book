package main

import "testing"

func TestHello(t *testing.T) {
  assertCorrectMessage := func (t *testing.T, got, want string) {
    t.Helper()
    if got != want {
      t.Errorf("got '%s' want '%s'", got, want)
    }
  }

  t.Run("saying specific people are great", func (t *testing.T) {
    got := Hello("Henry", "")
    want := "Henry is Great!"

    assertCorrectMessage(t, got, want)
  })

  t.Run("default person is great", func (t *testing.T) {
    got := Hello("", "")
    want := "Scott is Great!"

    assertCorrectMessage(t, got, want)
  })

  t.Run("in Spanish", func (t *testing.T) {
    got := Hello("Cesar", "Spanish")
    want := "Cesar es fantástica!"

    assertCorrectMessage(t, got, want)
  })

  t.Run("in French", func (t *testing.T) {
    got := Hello("Louis", "French")
    want := "Louis est Remarquable!"

    assertCorrectMessage(t, got, want)
  })

  t.Run("in Pig Latin", func (t *testing.T) {
    got := Hello("Ottscay", "Pig Latin")
    want := "Ottscay isway ate-gray!"

    assertCorrectMessage(t, got, want)
  })
}
