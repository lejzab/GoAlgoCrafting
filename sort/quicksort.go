package sort

// QuickSort sorts a slice of integers using the QuickSort algorithm.
func QuickSort(t []int, low, high int) {
	if low >= high {
		return
	}
	pivotIndex := partition(t, low, high)
	QuickSort(t, low, pivotIndex-1)
	QuickSort(t, pivotIndex+1, high)
}

// partition rearranges the elements in the slice such that elements smaller
// than the pivot are on the left, and elements greater than the pivot are on the right.
func partition(t []int, low, high int) int {
	pivot := t[high]
	i := low - 1 // index of the smallest element

	for j := low; j < high; j++ {
		if t[j] <= pivot {
			i++
			t[i], t[j] = t[j], t[i]
		}
	}
	i++ // move index one position to the right, after the smallest then pivot
	t[i], t[high] = t[high], t[i]
	return i
}
