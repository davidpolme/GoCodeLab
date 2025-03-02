package slidingwindow_test

import (
	"testing"

	"github.com/davidpolme/gocodelab/solutions/arrays_and_strings/sliding_window"
)

func TestLongestOnes(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected int
	}{
		{[]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2, 6},
		{[]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3, 10},
	}

	for _, test := range tests {
		result := sliding_window.LongestOnes(test.nums, test.k)
		if result != test.expected {
			t.Errorf("LongestOnes(%v, %d) = %d; want %d", test.nums, test.k, result, test.expected)
		}
	}
}
