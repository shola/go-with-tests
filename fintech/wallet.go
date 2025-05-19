package fintech

type Wallet struct {
	balance int
}

func (w *Wallet) Deposit(amt int)  {
	// golang auto dereferences this pointer!!!
	w.balance += amt
}

func (w *Wallet) Balance() int {
	return w.balance
}