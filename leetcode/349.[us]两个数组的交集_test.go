/*
- @lc app=leetcode.cn id=349 lang=golang
给定两个数组，编写一个函数来计算它们的交集。
说明： 输出结果中的每个元素一定是唯一的。 我们可以不考虑输出结果的顺序。
*/
package leetcode

import (
	"reflect"
	"testing"
)

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
	fre := make(map[int]bool)
	for _, v := range nums1 {
		fre[v] = true
	}

	res := make([]int, 0)
	for _, v := range nums2 {
		if fre[v] {
			res = append(res, v)
			delete(fre, v)
		}
	}
	return res
}

// @lc code=end

func Test_intersection(t *testing.T) {
	type args struct {
		a []int
		b []int
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
			[]int{1, 2, 4, 2, 4},
			[]int{2, 2, 3},
		}, []int{2}},
		{"1", args{
			[]int{4, 9, 5},
			[]int{9, 4, 9, 8, 4},
		}, []int{9, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersection(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intersection() = %v, want %v", got, tt.want)
			}
		})
	}
}
