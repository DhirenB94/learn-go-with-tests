package pointerserrors_test

import (
	pointerserrors "learn-go-with-tests/chapters/pointers-errors"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := pointerserrors.Wallet{}

	wallet.Deposit(10)

	got := wallet.Balance()
	want := 10

	if got != want {
		t.Errorf("got %d but want %d", got, want)
	}
}
