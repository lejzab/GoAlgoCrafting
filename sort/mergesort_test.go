package sort

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{name: "Simple sort",
			input:    []int{7, 2, 1, 6, 8, 5, 3, 4},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			name:     "Double elements",
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
			result := MergeSort(test.input)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("got %v, expected %v", test.input, test.expected)
			}
		})
	}
}
