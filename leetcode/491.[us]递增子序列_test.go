/*
- @lc app=leetcode.cn id=491 lang=golang

给定一个整型数组, 你的任务是找到所有该数组的递增子序列，递增子序列的长度至少是2。

示例:

输入: [4, 6, 7, 7]
输出: [[4, 6], [4, 7], [4, 6, 7], [4, 6, 7, 7], [6, 7], [6, 7, 7], [7,7], [4,7,7]]
说明:

给定数组的长度不会超过15。
数组中的整数范围是 [-100,100]。
给定数组中可能包含重复数字，相等的数字应该被视为递增的一种情况。
*/
package leetcode

import (
	"testing"
)

// @lc code=start
func findSubsequences(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)

	var backtrack func(idx int)
	backtrack = func(idx int) {
		if len(path) > 1 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
		}

		visited := [201]int{} // 同层元素去重
		for i := idx; i < len(nums); i++ {
			if len(path) > 0 && nums[i] < path[len(path)-1] {
				continue
			}

			if visited[nums[i]+100] == 1 {
				continue
			}

			visited[nums[i]+100] = 1
			path = append(path, nums[i])
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return res
}

// @lc code=end

func Test_findSubsequences(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{"0", []int{}, [][]int{}},
		{"0", []int{1}, [][]int{}},
		{"1", []int{1, 2}, [][]int{{1, 2}}},
		{"1", []int{4, 6, 7, 7}, [][]int{{4, 6}, {4, 7}, {4, 6, 7}, {4, 6, 7, 7}, {6, 7}, {6, 7, 7}, {7, 7}, {4, 7, 7}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSubsequences(tt.nums); !equalSetMatrix(got, tt.want) {
				t.Errorf("findSubsequences() = %v, want %v", got, tt.want)
			}
		})
	}
}
