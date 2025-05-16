package main

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	assertNumbersEqual := func(t testing.TB, got, want int, numbers []int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	}
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6
		assertNumbersEqual(t, got, want, numbers)
	})
	t.Run("empty collection", func(t *testing.T) {
		numbers := []int{}
		got := Sum(numbers)
		want := 0
		assertNumbersEqual(t, got, want, numbers)
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}
	assertSlicesEqual(t, got, want)
}

func TestSumAllTails(t *testing.T) {
	t.Run("all slices are non-empty", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{0, 2, 9})
		want := []int{5, 11}
		assertSlicesEqual(t, got, want)
	})

	t.Run("empty slice", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{})
		want := []int{5, 0}
		assertSlicesEqual(t, got, want)
	})
}

func assertSlicesEqual(t testing.TB, got, want []int) {
	t.Helper()
	// NOTE: for 2D slices, use reflect.DeepEqual to compare
	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
func BenchmarkSumAllTails(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAllTails([]int{1, 2, 3}, []int{0, 2, 9})
	}
}
