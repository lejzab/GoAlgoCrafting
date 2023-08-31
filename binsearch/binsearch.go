package binsearch

import "log"

// BinSearch performs binary search on a sorted integer slice 't' to find the index
// of the target value 'search'.
// t: The sorted integer slice.
// search: The target value to search for.
// Returns the index of the target value if found, otherwise returns -1.
func BinSearch(t []int, search int) int {
	lo := 0      // Initialize the lower index boundary.
	hi := len(t) // Initialize the upper index boundary.

	for lo < hi {
		m := (lo + hi) / 2 // Calculate the middle index of the current search range.

		log.Printf("%v, lo=%d, hi=%d\n", t[lo:hi], lo, hi)
		if search == t[m] {
			return m // If the middle element is equal to the target value, return its index.
		}
		if search > t[m] {
			lo = m + 1 // If the target value is greater than the middle element,
			// narrow the search to the right half of the array.
		} else {
			hi = m // If the target value is smaller than or equal to the middle element,
			// narrow the search to the left half of the array.
		}
	}

	return -1 // If the target value is not found in the array, return -1.
}
