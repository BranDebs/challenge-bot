package slices

func Uint64Exists(val uint64, vals ...uint64) bool {
	for _, v := range vals {
		if val == v {
			return true
		}
	}
	return false
}
