package sort

// MergeSort sorts a slice of integers using the MergeSort algorithm.
func MergeSort(t []int) []int {
	//log.Printf("input array. %v", t)
	if len(t) <= 1 {
		return t
	}
	m := len(t) / 2
	left := t[:m]
	right := t[m:]

	left = MergeSort(left)
	right = MergeSort(right)

	return merge(left, right)
}

func merge(left, right []int) []int {
	totalLength := len(left) + len(right)
	result := make([]int, 0, totalLength)
	indexLeft := 0
	indexRight := 0

	for indexLeft < len(left) && indexRight < len(right) {
		if left[indexLeft] <= right[indexRight] {
			result = append(result, left[indexLeft])
			indexLeft++
		} else {
			result = append(result, right[indexRight])
			indexRight++
		}
	}
	result = append(result, left[indexLeft:]...)
	result = append(result, right[indexRight:]...)

	return result

}
