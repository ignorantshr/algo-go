package leetcode

import "sort"

type compareEle interface {
	int | string
}

// 集合相等
func equalSet[T compareEle](a, b []T) bool {
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

func equalSlice[T compareEle](a, b []T) bool {
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

func equalSliceMatrix[T compareEle](a, b [][]T) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if !equalSlice(a[i], b[i]) {
			return false
		}
	}
	return true
}

func equalSetMatrix[T compareEle](a, b [][]T) bool {
	if len(a) != len(b) {
		return false
	}

	matrixSort(a)
	matrixSort(b)

	for i := range a {
		if !equalSet(a[i], b[i]) {
			return false
		}
	}
	return true
}

func matrixSort[T compareEle](a [][]T) {
	sort.Slice(a, func(i, j int) bool {
		if len(a[i]) < len(a[j]) {
			return true
		} else if len(a[i]) > len(a[j]) {
			return false
		}
		for m := range a[i] {
			if a[i][m] > a[j][m] {
				return false
			}
		}
		return true
	})
}
