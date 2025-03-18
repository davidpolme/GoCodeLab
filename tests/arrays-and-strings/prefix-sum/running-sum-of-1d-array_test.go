package prefixsum_test

import (
	"testing"
	"reflect"

	// Importa la función desde el archivo original
	"github.com/davidpolme/gocodelab/solutions/arrays_and_strings/prefix_sum"
)

func TestRunningSum_Basic(t *testing.T) {
	input := []int{1, 2, 3, 4}
	expected := []int{1, 3, 6, 10}
	result := prefix_sum.RunningSum(input) 

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func TestRunningSum_WithZeros(t *testing.T) {
	input := []int{0, 0, 0, 0}
	expected := []int{0, 0, 0, 0}
	result := prefix_sum.RunningSum(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func TestRunningSum_SingleElement(t *testing.T) {
	input := []int{5}
	expected := []int{5}
	result := prefix_sum.RunningSum(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func TestRunningSum_NegativeNumbers(t *testing.T) {
	input := []int{-1, -2, -3, -4}
	expected := []int{-1, -3, -6, -10}
	result := prefix_sum.RunningSum(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func TestRunningSum_MixedNumbers(t *testing.T) {
	input := []int{-1, 3, -2, 5}
	expected := []int{-1, 2, 0, 5}
	result := prefix_sum.RunningSum(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}
