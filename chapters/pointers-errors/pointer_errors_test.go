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
	t.Run("withdraw with sufficient funds", func(t *testing.T) {
		wallet := pointerserrors.Wallet{}

		wallet.Deposit(pointerserrors.Bitcoin(20))
		err := wallet.Withdraw(pointerserrors.Bitcoin(10))

		assertNoError(t, err)
		assertBalance(t, wallet, pointerserrors.Bitcoin(10))
	})
	t.Run("withdraw insufficient funds", func(t *testing.T) {
		wallet := pointerserrors.Wallet{}

		wallet.Deposit(pointerserrors.Bitcoin(20))

		err := wallet.Withdraw(pointerserrors.Bitcoin(100))

		assertError(t, err, pointerserrors.ErrInsufficientFunds)
		assertBalance(t, wallet, pointerserrors.Bitcoin(20))

	})

}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("wanted an error but did not get one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertBalance(t testing.TB, wallet pointerserrors.Wallet, want pointerserrors.Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s but want %s", got, want)
	}
}
