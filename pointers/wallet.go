package main

import (
	"errors"
	"fmt"
)

// Bitcoin represents a number of Bitcoins
type Bitcoin int

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	// ^ if the * is excluded the balance of a copy of w changes, not w itself
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	// ^ method receiver type is kept the same for consistency
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}
