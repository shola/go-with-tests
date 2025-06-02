package collections

type Transaction struct {
	From string
	To   string
	Sum  float64
}

func BalanceFor(transactions []Transaction, name string) float64 {

	adjustBalance := func(accum float64, item Transaction) float64 {
		if item.From == name {
			return accum - item.Sum
		}
		if item.To == name {
			return accum + item.Sum
		}
		return accum
	}
	return Reduce(transactions, adjustBalance, 0.0)
}
