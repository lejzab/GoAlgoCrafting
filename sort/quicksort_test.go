package sort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Simple sort",
			input:    []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5},
			expected: []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9},
		},
		{
			name:     "Single element",
			input:    []int{-1},
			expected: []int{-1},
		},
		{
			name:     "Empty array",
			input:    []int{},
			expected: []int{},
		},
		// Add more test cases here

	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			QuickSort(test.input, 0, len(test.input)-1)
			if !reflect.DeepEqual(test.input, test.expected) {
				t.Errorf("got %v, expected %v", test.input, test.expected)
			}
		})
	}
}
