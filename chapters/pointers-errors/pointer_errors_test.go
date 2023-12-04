package pointerserrors_test

import (
	pointerserrors "learn-go-with-tests/chapters/pointers-errors"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := pointerserrors.Wallet{}

		wallet.Deposit(pointerserrors.Bitcoin(10))
		assertBalance(t, wallet, pointerserrors.Bitcoin(10))
	})
	t.Run("withdraw", func(t *testing.T) {
		wallet := pointerserrors.Wallet{}

		wallet.Deposit(pointerserrors.Bitcoin(20))
		wallet.Withdraw(pointerserrors.Bitcoin(10))
		assertBalance(t, wallet, pointerserrors.Bitcoin(10))
	})

}

func assertBalance(t testing.TB, wallet pointerserrors.Wallet, want pointerserrors.Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s but want %s", got, want)
	}
}
