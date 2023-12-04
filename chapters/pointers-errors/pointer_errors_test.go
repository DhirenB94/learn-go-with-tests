package pointerserrors_test

import (
	pointerserrors "learn-go-with-tests/chapters/pointers-errors"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := pointerserrors.Wallet{}

		wallet.Deposit(pointerserrors.Bitcoin(10))

		got := wallet.Balance()
		want := pointerserrors.Bitcoin(10)

		if got != want {
			t.Errorf("got %s but want %s", got, want)
		}
	})
	t.Run("withdraw", func(t *testing.T) {
		wallet := pointerserrors.Wallet{}

		wallet.Deposit(pointerserrors.Bitcoin(20))
		wallet.Withdraw(pointerserrors.Bitcoin(10))

		got := wallet.Balance()
		want := pointerserrors.Bitcoin(10)

		if got != want {
			t.Errorf("got %s but want %s", got, want)
		}
	})

}
