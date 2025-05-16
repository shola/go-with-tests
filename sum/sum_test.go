package main

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	// NOTE: for 2D slices, use reflect.DeepEqual to compare
	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails1(t *testing.T) {
	got := SumAllTails1([]int{1, 2, 3}, []int{0, 2, 9})
	want := []int{5, 11}

	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
func TestSumAllTails2(t *testing.T) {
	got := SumAllTails2([]int{1, 2, 3}, []int{0, 2, 9})
	want := []int{5, 11}

	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func BenchmarkSumAllTails1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAllTails1([]int{1, 2, 3}, []int{0, 2, 9})
	}
}

func BenchmarkSumAllTails2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAllTails2([]int{1, 2, 3}, []int{0, 2, 9})
	}
}
