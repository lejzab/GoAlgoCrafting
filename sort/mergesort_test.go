package sort

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{[]int{7, 2, 1, 6, 8}, []int{1, 2, 6, 7, 8}},
		{[]int{7, 2, 1, 6, 8, 5, 3, 4}, []int{1, 2, 3, 4, 5, 6, 7, 8}},
		{[]int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}, []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}},
		{[]int{-1}, []int{-1}},
		{[]int{}, []int{}},
		// Add more test cases here

	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			result := MergeSort(test.input)
			if !reflect.DeepEqual(result, test.want) {
				t.Errorf("got %v, want %v", test.input, test.want)
			}
		})
	}
}
