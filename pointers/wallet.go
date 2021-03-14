package main

type Bitcoin int

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
