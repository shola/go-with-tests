package main

func Sum(numbers []int) int {
	sum := 0

	for _, v := range numbers {
		sum += v
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	results := make([]int, lengthOfNumbers)

	for i, v := range numbersToSum {
		results[i] = Sum(v)
	}

	return results
}
