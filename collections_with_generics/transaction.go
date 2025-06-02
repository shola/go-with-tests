package collections

type Transaction struct {
	From string
	To   string
	Sum  int
}

func BalanceFor(transactions []Transaction, name string) float64 {
	return 0.0
}
