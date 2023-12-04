package pointerserrors

import "fmt"

type Bitcoin int

// this uses the Stringer interface, defined in the fmt package
// allows you to define how your type is printed when using the %s in prints
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
