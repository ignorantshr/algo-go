package leetcode

import (
	"math"
	"testing"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	return resolution2(nums1, nums2)
}

func resolution2(nums1 []int, nums2 []int) float64 {
	l := len(nums1) + len(nums2)
	if (l)&1 == 1 {
		return float64(findKthNum2(l/2+1, nums1, nums2))
	}
	return float64(findKthNum2(l/2, nums1, nums2)+findKthNum2(l/2+1, nums1, nums2)) / 2
}

// 查找第K个数，从1开始
func findKthNum(k int, nums1, nums2 []int) int {
	offsetA, offsetB := 0, 0
	l1 := len(nums1)
	l2 := len(nums2)

	for {
		if offsetA >= l1 {
			return nums2[offsetB+k-1]
		}
		if offsetB >= l2 {
			return nums1[offsetA+k-1]
		}

		if k == 1 {
			return min(nums1[offsetA], nums2[offsetB])
		}

		h := k/2 - 1
		h1 := min(offsetA+h, l1-1)
		h2 := min(offsetB+h, l2-1)

		// a[0,...,k/2-1], b[0,...,k/2-1]. a[k/2-1]<=b[k/2-1].
		// 至多 (a 中的 k/2 + b 中的 k/2-1) 个数（求和等于 k-1）比 b[k/2-1] 小，
		// 所以 b[k/2-1] 不可能是 第 k 个小数，所以可以把 a[0,...,k/2-1] 排除
		if nums1[h1] <= nums2[h2] {
			k -= (h1 - offsetA + 1)
			offsetA = h1 + 1
		} else {
			k -= (h2 - offsetB + 1)
			offsetB = h2 + 1
		}

	}
}

func resolution1(nums1 []int, nums2 []int) float64 {
	nums := merge(nums1, nums2)
	if len(nums) == 0 {
		return 0
	}

	a, b := 0, len(nums)-1
	mid := a + (b-a)/2
	if len(nums)&1 == 1 {
		return float64(nums[mid])
	} else {
		return float64(nums[mid]+nums[mid+1]) / 2
	}
}

func merge(nums1 []int, nums2 []int) []int {
	res := make([]int, len(nums1)+len(nums2))
	for a, b := 0, 0; a < len(nums1) || b < len(nums2); {
		if a == len(nums1) {
			res[a+b] = nums2[b]
			b++
			continue
		}
		if b == len(nums2) {
			res[a+b] = nums1[a]
			a++
			continue
		}

		if nums1[a] > nums2[b] {
			res[a+b] = nums2[b]
			b++
		} else {
			res[a+b] = nums1[a]
			a++
		}
	}
	return res
}

func TestFindMedianSortedArrays(t *testing.T) {
	table := []struct {
		name string
		arr1 []int
		arr2 []int
		re   float64
	}{
		{
			"1",
			[]int{1, 4, 5},
			[]int{2, 4},
			4,
		},
		{
			"2",
			[]int{1, 4, 5},
			[]int{6, 9},
			5,
		},
		{
			"3",
			[]int{6, 9},
			[]int{1, 4, 5},
			5,
		},
		{
			"4",
			[]int{2, 6, 9},
			[]int{1, 4, 5},
			4.5,
		},
		{
			"5",
			[]int{6, 8, 9},
			[]int{1, 4, 5},
			5.5,
		},
		{
			"6",
			[]int{1, 4, 5},
			[]int{6, 8, 9},
			5.5,
		},
		{
			"7",
			[]int{1, 3},
			[]int{2},
			2,
		},
		{
			"8",
			[]int{1, 3},
			[]int{},
			2,
		},
		{
			"9",
			[]int{0, 0, 0, 0, 0},
			[]int{-1, 0, 0, 0, 0, 0, 1},
			0,
		},
		{
			"10",
			[]int{1, 2, 3, 4},
			[]int{5, 6},
			3.5,
		},
	}
	for _, v := range table {
		t.Run(v.name, func(t *testing.T) {
			re := findMedianSortedArrays(v.arr1, v.arr2)
			if v.re != re {
				t.Fatalf("\n%x\n%x\n%f", v.arr1, v.arr2, re)
			}
		})
	}
}

func findKthNum2(k int, arr1, arr2 []int) int {
	l1 := len(arr1)
	l2 := len(arr2)
	if k > l1+l2 {
		return int(math.Inf(-1))
	}

	// a1{1,...,k/2}, a2{1,...,k/2}
	// if a1[k/2] <= a2[k/2], 则比 a1[k/2] 小的数有至多这么多个：
	// (k/2-1) + (k/2-1) = k-2
	// 所以 a2[k/2] 最多也只是第 k-1 的小数，不可能是第 k 小数
	// 故可以把 a1{1,...,k/2} 去除

	offset1, offset2 := 0, 0
	for {
		if offset2 >= l2 {
			return arr1[offset1+k-1]
		}
		if offset1 >= l1 {
			return arr2[offset2+k-1]
		}
		if k == 1 {
			return min(arr1[offset1], arr2[offset2])
		}

		h := k/2 - 1
		h1 := min(offset1+h, l1-1)
		h2 := min(offset2+h, l2-1)
		if arr1[h1] > arr2[h2] {
			k -= h2 - offset2 + 1
			offset2 = h2 + 1
		} else {
			k -= h1 - offset1 + 1
			offset1 = h1 + 1
		}
	}
}
