/*
- @lc app=leetcode.cn id=1 lang=golang
*/
package leetcode

import (
	"reflect"
	"testing"
)

// @lc code=start

func twoSumReview(nums []int, target int) []int {
	records := make(map[int]int)
	for i, v := range nums {
		if idx, ok := records[target-v]; ok {
			return []int{idx, i}
		} else {
			records[v] = i
		}
	}
	return []int{}
}

// @lc code=end

func Test_twoSumReview(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{
			[]int{},
			0,
		}, []int{}},
		{"1", args{
			[]int{1, 43, 2, 6, 8},
			0,
		}, []int{}},
		{"1", args{
			[]int{1, 43, 2, 6, 8},
			7,
		}, []int{0, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSumReview(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSumReview() = %v, want %v", got, tt.want)
			}
		})
	}
}
