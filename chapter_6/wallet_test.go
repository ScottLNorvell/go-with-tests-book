package main

import (
  _ "fmt"
  "testing"
)

func TestWallet(t *testing.T) {
  t.Run("Deposit", func (t * testing.T) {
    w := Wallet{}
    w.Deposit(Bitcoin(10))
    assertBalance(t, w, Bitcoin(10))
  })

  t.Run("Withdraw", func (t *testing.T) {
    w := Wallet{Bitcoin(30)}
    err := w.Withdraw(Bitcoin(10))
    assertBalance(t, w, Bitcoin(20))
    assertNoError(t, err)
  })

  t.Run("Withdraw with insufficient funds", func (t *testing.T) {
    startingBalance := Bitcoin(10)
    w := Wallet{startingBalance}
    err := w.Withdraw(Bitcoin(100))
    assertBalance(t, w, startingBalance)
    assertError(t, err, ErrInsufficientFunds)
  })
}

func assertBalance(t *testing.T, w Wallet, want Bitcoin) {
  t.Helper()
  got := w.Balance()
  if got != want {
    t.Errorf("got '%s' want '%s'", got, want)
  }
}

func assertNoError(t *testing.T, got error) {
  t.Helper()
  if got != nil {
    t.Fatal("got an error but didn't want one")
  }
}


func assertError(t *testing.T, got error, want error) {
  t.Helper()
  if got == nil {
    t.Fatal("wanted an error but didn't get one")
  }

  if got != want {
    t.Errorf("got '%s' want '%s'", got, want)
  }
}
