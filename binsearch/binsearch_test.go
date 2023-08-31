package binsearch

import "testing"

func TestBinSearch(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		search   int
		expected []int
	}{
		{
			name:     "Element found at the beginning",
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			search:   1,
			expected: []int{0},
		},
		{
			name:     "Element found in the middle",
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			search:   5,
			expected: []int{4},
		},
		{
			name:     "Element found at the end",
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			search:   9,
			expected: []int{8},
		},
		{
			name:     "Element not found",
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			search:   10,
			expected: []int{-1},
		},
		{
			name:     "Element somewhere in array",
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			search:   4,
			expected: []int{3},
		},
		{
			name:     "Empty array",
			input:    []int{},
			search:   4,
			expected: []int{-1},
		},
		{
			name:     "Doubled elements",
			input:    []int{1, 2, 3, 3, 4, 5, 6, 6, 7, 8, 9},
			search:   6,
			expected: []int{6, 7},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := BinSearch(testCase.input, testCase.search)

			// Check if the result is in the expected array.
			found := false
			for _, idx := range testCase.expected {
				if result == idx {
					found = true
					break
				}
			}

			if !found {
				t.Errorf("Expected index in %v, but got %d", testCase.expected, result)
			}
		})
	}
}
