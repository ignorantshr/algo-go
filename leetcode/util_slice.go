package leetcode

type comareEle interface {
	int | string
}

// 集合相等
func equalSet[T comareEle](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	am := make(map[T]int)
	bm := make(map[T]int)
	for i := range a {
		am[a[i]]++
		bm[b[i]]++
	}
	if len(am) != len(bm) {
		return false
	}

	for k, v := range am {
		if v != bm[k] {
			return false
		}
	}
	return true
}

func equalSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
