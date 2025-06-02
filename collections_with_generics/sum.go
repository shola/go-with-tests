package collections

func Reduce[T any](list []T, fn func(accum, item T) T, initialValue T) T {
	res := initialValue

	for _, v := range list {
		res = fn(res, v)
	}
	return res
}

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
