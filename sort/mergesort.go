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
	var result []int
	indexLeft := 0
	indexRight := 0

	//log.Printf("merge start. %v, %v\n", left, right)
	for indexLeft < len(left) && indexRight < len(right) {
		if left[indexLeft] <= right[indexRight] {
			result = append(result, left[indexLeft])
			indexLeft += 1
		} else {
			result = append(result, right[indexRight])
			indexRight += 1
		}
	}
	//log.Printf("merge po pętli. %v, %v, %v\n", result,
	//	left[indexLeft:],
	//	right[indexRight:])
	result = append(result, left[indexLeft:]...)
	result = append(result, right[indexRight:]...)

	//log.Printf("merged array. %v\n", result)
	return result

}
