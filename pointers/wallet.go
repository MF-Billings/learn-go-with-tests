package main

type Wallet struct {
	balance int
}

func (w *Wallet) Deposit(amount int) {
	// ^ if the * is excluded the balance of a copy of w changes, not w itself
	w.balance += amount
}

func (w *Wallet) Balance() int {
	// ^ method receiver type is kept the same for consistency
	return w.balance
}
