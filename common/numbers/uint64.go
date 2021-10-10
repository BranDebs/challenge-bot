package numbers

func Uint64Compare(l uint64, r uint64) int {
	if l > r {
		return 1
	}
	if l < r {
		return -1
	}
	return 0
}
