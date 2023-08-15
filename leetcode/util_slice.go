package leetcode

import "sort"

// 集合相等
func equalSet(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	a1 := make([]int, len(a))
	b1 := make([]int, len(b))
	copy(a1, a)
	copy(b1, b)
	sort.Ints(a1)
	sort.Ints(b1)
	for i := 0; i < len(a1); i++ {
		if a1[i] != b1[i] {
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
