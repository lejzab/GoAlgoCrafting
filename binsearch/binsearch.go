package binsearch

// BinSearch is a recursive binary search algorithm that searches for a target value
// in a sorted integer array.
// t: The sorted integer array.
// lo: The lower index boundary of the search range.
// hi: The upper index boundary of the search range.
// search: The target value to search for.
// Returns the index of the target value if found, otherwise returns -1.
func BinSearch(t []int, lo, hi int, search int) int {
	// Calculate the middle index of the current search range.
	m := (lo + hi) / 2

	// Base case: If the lower index is greater than or equal to the upper index,
	// the search range is empty, and the target value is not found.
	if lo >= hi {
		return -1
	}

	// If the middle element is equal to the target value, return its index.
	if search == t[m] {
		return m
	}

	// If the target value is smaller than the middle element,
	// recursively search the left half of the array.
	if search < t[m] {
		return BinSearch(t, lo, m, search)
	} else {
		// If the target value is larger than or equal to the middle element,
		// recursively search the right half of the array.
		return BinSearch(t, m+1, hi, search)
	}
}
