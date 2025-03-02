package twopointers_test

import (
	"reflect"
	"testing"

	"github.com/davidpolme/gocodelab/solutions/arrays_and_strings/two_pointers"
)

func TestReverseString(t *testing.T) {
	tests := []struct {
		input    []byte
		expected []byte
	}{
		{[]byte("hello"), []byte("olleh")},
		{[]byte("Go"), []byte("oG")},
		{[]byte("a"), []byte("a")},
		{[]byte(""), []byte("")},
	}

	for _, test := range tests {
		inputCopy := make([]byte, len(test.input))
		copy(inputCopy, test.input)

		two_pointers.ReverseString(inputCopy)

		if !reflect.DeepEqual(inputCopy, test.expected) {
			t.Errorf("ReverseString(%q) = %q; want %q", test.input, inputCopy, test.expected)
		}
	}
}
