package arraysandslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		expected := 6

		if got != expected {
			t.Errorf("expected %d but got %d, %v", expected, got, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 2})
	expected := []int{3, 2}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %v want %v", got, expected)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, expected []int) {
		t.Helper()

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got %v want %v", got, expected)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 2})
		expected := []int{2, 2}

		checkSums(t, got, expected)
	})

	t.Run("safety sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{1}, []int{0, 2})
		expected := []int{0, 2}

		checkSums(t, got, expected)
	})
}
