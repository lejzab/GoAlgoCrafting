package binsearch

func BinSearch(t []int, lo, hi int, search int) int {
	m := (lo + hi) / 2
	if lo >= hi {
		return -1
	}
	if search == t[m] {
		return m
	}
	if search < t[m] {
		return BinSearch(t, lo, m, search)
	} else {
		return BinSearch(t, m+1, hi, search)
	}
}
