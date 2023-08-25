package leetcode

import (
	"reflect"
	"testing"
)

func nextGreaterElementReview(nums1 []int, nums2 []int) []int {
	res := make([]int, 0)
	stack := make([]int, 0)
	reflect := make(map[int]int)
	last := -1

	for i := len(nums2) - 1; i >= 0; i-- {
		for last > -1 && stack[last] <= nums2[i] {
			stack = stack[:last]
			last--
		}
		if last >= 0 {
			reflect[nums2[i]] = stack[last]
		} else {
			reflect[nums2[i]] = -1
		}
		stack = append(stack, nums2[i])
		last++
	}

	for _, v := range nums1 {
		res = append(res, reflect[v])
	}

	return res
}

func Test_nextGreaterElementReview(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"0", args{
			[]int{},
			[]int{},
		}, []int{}},
		{"1", args{
			[]int{1},
			[]int{1},
		}, []int{-1}},
		{"1", args{
			[]int{4, 1, 2},
			[]int{1, 3, 4, 2},
		}, []int{-1, 3, -1}},
		{"1", args{
			[]int{2, 4},
			[]int{1, 2, 3, 4},
		}, []int{3, -1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreaterElementReview(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextGreaterElementReview() = %v, want %v", got, tt.want)
			}
		})
	}
}
