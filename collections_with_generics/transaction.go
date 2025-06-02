package collections

type Transaction struct {
	From string
	To   string
	Sum  float64
}

func BalanceFor(transactions []Transaction, name string) float64 {
	// createLedger := func(accum float64, item Transaction) Transaction {
	// 	// do addition and subtraction for each in the pair
	// }
	// return Reduce(transactions, createLedger, 0.0)
	var balance float64
	for _, t := range transactions {
		if t.From == name {
			balance -= t.Sum
		}
		if t.To == name {
			balance += t.Sum
		}
	}
	return balance
}
