package slidingwindow_test

import (
	"testing"

	"github.com/davidpolme/gocodelab/solutions/arrays_and_strings/sliding_window"
)

func TestFindMaxAverage(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected float64
	}{
		{[]int{1, 12, -5, -6, 50, 3}, 4, 12.75},
		{[]int{5}, 1, 5.0},
	}

	for _, test := range tests {
		result := sliding_window.FindMaxAverage(test.nums, test.k)
		if result != test.expected {
			t.Errorf("FindMaxAverage(%v, %d) = %f; want %f", test.nums, test.k, result, test.expected)
		}
	}
}
