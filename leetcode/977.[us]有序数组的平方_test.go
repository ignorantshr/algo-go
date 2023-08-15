/*
  - @lc app=leetcode.cn id=977 lang=golang

给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。

示例 1：

输入：nums = [-4,-1,0,3,10]
输出：[0,1,9,16,100]
解释：平方后，数组变为 [16,1,0,9,100]，排序后，数组变为 [0,1,9,16,100]
示例 2：

输入：nums = [-7,-3,2,3,11]
输出：[4,9,9,49,121]
*/
package leetcode

import (
	"reflect"
	"testing"
)

// @lc code=start
func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	left, right := 0, len(nums)-1

	for n := right; n >= 0; n-- {
		powRight := nums[right] * nums[right]
		powLeft := nums[left] * nums[left]
		if powLeft <= powRight {
			res[n] = powRight
			right--
		} else {
			res[n] = powLeft
			left++
		}
	}
	return res
}

// @lc code=end

func Test_sortedSquares(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"0", []int{}, []int{}},
		{"1", []int{1, 2, 3}, []int{1, 4, 9}},
		{"1", []int{-1, 2, 3}, []int{1, 4, 9}},
		{"1", []int{-2, 1, 2, 3}, []int{1, 4, 4, 9}},
		{"1", []int{-7, -3, 2, 3, 11}, []int{4, 9, 9, 49, 121}},
		{"1", []int{-4, -1, 0, 3, 10}, []int{0, 1, 9, 16, 100}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedSquares(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
