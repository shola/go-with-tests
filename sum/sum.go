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

// goos: linux
// goarch: amd64
// pkg: notmain
// cpu: Intel(R) Core(TM) i7-10700 CPU @ 2.90GHz
// BenchmarkSumAllTails1-16         7053066               187.8 ns/op           144 B/op          6 allocs/op
// BenchmarkSumAllTails2-16        28372729                42.70 ns/op           24 B/op          2 allocs/op
// PASS
// ok      notmain 2.755s

func SumAllTails1(numbersToSum ...[]int) []int {
	var tailSlices [][]int

	for _, numbers := range numbersToSum {
		tailSlices = append(tailSlices, numbers[1:])
	}

	return SumAll(tailSlices...)
}

func SumAllTails2(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		tail := numbers[1:]
		sums = append(sums, Sum(tail))
	}

	return sums
}
