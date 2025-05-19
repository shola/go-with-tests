package fintech

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(amt Bitcoin) {
	// golang auto dereferences this pointer!!!
	w.balance += amt
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amt Bitcoin) error {
	if amt > w.Balance() {
		return errors.New("you can't withdraw more than you have!")
	}
	w.balance -= amt

	return nil
}
