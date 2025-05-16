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
