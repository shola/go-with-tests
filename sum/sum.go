package main

func Sum(numbers []int) int {
	sum := 0

	for _, v := range numbers {
		sum += v
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	return nil
}
