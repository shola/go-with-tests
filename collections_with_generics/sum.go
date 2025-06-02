package collections

func Sum(numbers []int) int {
	add := func(accum, item int) int {
		return accum + item
	}
	return Reduce(numbers, add, 0)
}

func SumAll(numbersToSum ...[]int) []int {
	addChildren := func(accum, item []int) []int {
		return append(accum, Sum(item))
	}
	return Reduce(numbersToSum, addChildren, []int{})
}

func SumAllTails(numbersToSum ...[]int) []int {
	addTails := func(accum, item []int) []int {
		var tail []int
		if len(item) == 0 {
			tail = []int{}
		} else {
			tail = item[1:]
		}
		return append(accum, Sum(tail))
	}
	return Reduce(numbersToSum, addTails, []int{})
}
