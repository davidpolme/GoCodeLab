package twopointers_test

import (
	"reflect"
	"testing"

	"github.com/davidpolme/gocodelab/solutions/arrays_and_strings/two_pointers"
)

func TestSortedSquares(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{-4, -1, 0, 3, 10}, []int{0, 1, 9, 16, 100}},
		{[]int{-7, -3, 2, 3, 11}, []int{4, 9, 9, 49, 121}},
		{[]int{1}, []int{1}},
		{[]int{-2, -1}, []int{1, 4}},
	}

	for _, test := range tests {
		result := two_pointers.SortedSquares(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("SortedSquares(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}
