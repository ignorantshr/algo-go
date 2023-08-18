/*
- @lc app=leetcode.cn id=18 lang=golang
*/
package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

// @lc code=start
func fourSumReview(nums []int, target int) [][]int {
	sort.Ints(nums)
	size := len(nums)
	res := make([][]int, 0)
	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < size; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			ntarget := target - nums[i] - nums[j]
			left, right := j+1, size-1
			for left < right {
				if nums[left]+nums[right] < ntarget {
					left++
				} else if nums[left]+nums[right] > ntarget {
					right--
				} else {
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
					for left = left + 1; left < right && nums[left] == nums[left-1]; left++ {
					}
				}
			}
		}
	}
	return res
}

// @lc code=end

func Test_fourSumReview(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"1", args{
			[]int{1, 0, -1, 0, -2, 2},
			0,
		},
			[][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}},
		},
		{"4", args{
			[]int{1, -2, -5, -4, -3, 3, 3, 5},
			-11,
		},
			[][]int{{-5, -4, -3, 1}},
		},
		{"2", args{
			[]int{-2, -1, -1, 1, 1, 2, 2},
			0,
		},
			[][]int{{-2, -1, 1, 2}, {-1, -1, 1, 1}},
		},
		{"3", args{
			[]int{2, 2, 2, 2},
			8,
		},
			[][]int{{2, 2, 2, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fourSumReview(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fourSumReview() = %v, want %v", got, tt.want)
			}
		})
	}
}
