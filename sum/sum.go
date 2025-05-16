package main

func Sum(numbers []int) int {
	sum := 0

	for _, v := range numbers {
		sum += v
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, v := range numbersToSum {
		sums = append(sums, Sum(v))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		var tail []int
		if len(numbers) == 0 {
			tail = []int{}
		} else {
			tail = numbers[1:]
		}
		sums = append(sums, Sum(tail))
	}

	return sums
}
