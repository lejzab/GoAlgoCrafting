package binsearch

import "testing"

func TestBinSearch(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		lo       int
		hi       int
		search   int
		expected int
	}{
		{
			name:     "Element found at the beginning",
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			lo:       0,
			hi:       9,
			search:   1,
			expected: 0,
		},
		{
			name:     "Element found in the middle",
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			lo:       0,
			hi:       9,
			search:   5,
			expected: 4,
		},
		{
			name:     "Element found at the end",
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			lo:       0,
			hi:       9,
			search:   9,
			expected: 8,
		},
		{
			name:     "Element not found",
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			lo:       0,
			hi:       9,
			search:   10,
			expected: -1,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := BinSearch(testCase.input, testCase.lo, testCase.hi, testCase.search)
			if result != testCase.expected {
				t.Errorf("Expected index %d, but got %d", testCase.expected, result)
			}
		})
	}
}
