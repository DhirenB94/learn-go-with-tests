package pointerserrors_test

import (
	"fmt"
	pointerserrors "learn-go-with-tests/chapters/pointers-errors"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := pointerserrors.Wallet{}

	wallet.Deposit(10)

	got := wallet.Balance()
	fmt.Printf("address of balance in test is %p \n", &got)

	want := 10

	if got != want {
		t.Errorf("got %d but want %d", got, want)
	}
}
