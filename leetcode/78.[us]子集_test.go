/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 *
 * https://leetcode.cn/problems/subsets/description/
 *
 * algorithms
 * Medium (81.13%)
 * Likes:    2143
 * Dislikes: 0
 * Total Accepted:    675.9K
 * Total Submissions: 833K
 * Testcase Example:  '[1,2,3]'
 *
 * 给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
 *
 * 解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,2,3]
 * 输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [0]
 * 输出：[[],[0]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * -10
 * nums 中的所有元素 互不相同
 *
 *
 */
package leetcode

import (
	"testing"
)

// @lc code=start
func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	var backtrack func(startIdx int)
	backtrack = func(startIdx int) {
		dest := make([]int, len(path))
		copy(dest, path)
		res = append(res, dest)

		for i := startIdx; i < len(nums); i++ {
			path = append(path, nums[i])
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return res
}

// @lc code=end

func Test_subsets(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{"0", []int{}, [][]int{{}}},
		{"1", []int{1, 2, 3}, [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subsets(tt.nums); !equalSetMatrix(got, tt.want) {
				t.Errorf("subsets() = %v, want %v", got, tt.want)
			}
		})
	}
}
