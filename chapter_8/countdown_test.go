package main

import (
  "bytes"
  "reflect"
  "testing"
)

type CountdownOperationsSpy struct {
  Calls []string
}

func (s *CountdownOperationsSpy) Sleep() {
  s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationsSpy) Write([]byte) (n int, err error) {
  s.Calls = append(s.Calls, write)
  return
}

const sleep = "sleep"
const write = "write"

func TestCountdown(t *testing.T) {
  t.Run("output correct", func (t *testing.T) {

    buffer := &bytes.Buffer{}

    Countdown(buffer, &CountdownOperationsSpy{})

    got := buffer.String()
    want := `3
2
1
Go!`

    if got != want {
      t.Errorf("got '%s' want '%s'", got, want)
    }
  })

  t.Run("sleep after every print", func (t *testing.T) {
    spy := &CountdownOperationsSpy{}
    Countdown(spy, spy)

    want := []string{
      sleep,
      write,
      sleep,
      write,
      sleep,
      write,
      sleep,
      write,
    }

    if !reflect.DeepEqual(want, spy.Calls) {
      t.Errorf("wanted calls %v got %v", want, spy.Calls)
    }
  })
}
