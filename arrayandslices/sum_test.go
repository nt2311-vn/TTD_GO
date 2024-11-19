package arrayandslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		assertCorrectResult(t, numbers, got, want)
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 50}

		got := Sum(numbers)
		want := 56

		assertCorrectResult(t, numbers, got, want)
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func assertCorrectResult(t testing.TB, numbers []int, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got '%d' but want '%d' input numbers %v", numbers, got, want)
	}
}