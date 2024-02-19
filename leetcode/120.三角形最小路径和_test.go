/*
 * @lc app=leetcode.cn id=120 lang=golang
 *
 * [120] 三角形最小路径和
 *
 * https://leetcode.cn/problems/triangle/description/
 *
 * algorithms
 * Medium (68.62%)
 * Likes:    1310
 * Dislikes: 0
 * Total Accepted:    330K
 * Total Submissions: 480.6K
 * Testcase Example:  '[[2],[3,4],[6,5,7],[4,1,8,3]]'
 *
 * 给定一个三角形 triangle ，找出自顶向下的最小路径和。
 *
 * 每一步只能移动到下一行中相邻的结点上。相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1
 * 的两个结点。也就是说，如果正位于当前行的下标 i ，那么下一步可以移动到下一行的下标 i 或 i + 1 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
 * 输出：11
 * 解释：如下面简图所示：
 * ⁠  2
 * ⁠ 3 4
 * ⁠6 5 7
 * 4 1 8 3
 * 自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。
 *
 *
 * 示例 2：
 *
 *
 * 输入：triangle = [[-10]]
 * 输出：-10
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= triangle.length <= 200
 * triangle[0].length == 1
 * triangle[i].length == triangle[i - 1].length + 1
 * -104 <= triangle[i][j] <= 104
 *
 *
 * 进阶：
 *
 *
 * 你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题吗？
 *
 *
 */
package leetcode

import (
	"testing"
)

// @lc code=start
func minimumTotal(triangle [][]int) int {
	return minimumTotalDpImprove(triangle)
	return minimumTotalDp(triangle)
	return minimumTotalBottom(triangle)
}

func minimumTotalDpImprove(triangle [][]int) int {
	// 空间优化
	size := len(triangle)
	dp := make([]int, size)

	for i := 0; i < size; i++ {
		for j := i; j >= 0; j-- {
			if j == 0 {
				dp[j] += triangle[i][0]
			} else if j < i {
				dp[j] = min(dp[j], dp[j-1]) + triangle[i][j]
			} else {
				dp[j] = dp[j-1] + triangle[i][j]
			}
		}
	}

	res := dp[0]
	for i := 1; i < size; i++ {
		res = min(res, dp[i])
	}

	return res
}

func minimumTotalDp(triangle [][]int) int {
	size := len(triangle)
	dp := make([][]int, size)

	for i := 0; i < size; i++ {
		dp[i] = make([]int, i+1)
		if i != 0 {
			dp[i][0] = dp[i-1][0] + triangle[i][0]
		} else {
			dp[0][0] = triangle[0][0]
		}
	}

	for i := 1; i < size; i++ {
		for j := 1; j <= i; j++ {
			if j < i {
				dp[i][j] = min(dp[i-1][j], dp[i-1][j-1]) + triangle[i][j]
			} else {
				dp[i][j] = dp[i-1][j-1] + triangle[i][j]
			}
		}
	}

	res := dp[size-1][0]
	for i := 1; i < size; i++ {
		res = min(res, dp[size-1][i])
	}

	return res
}

func minimumTotalBottom(triangle [][]int) int {
	size := len(triangle)
	tmp := make([]int, size)

	for i := size - 1; i >= 0; i-- {
		tmp[i] = triangle[size-1][i]
	}

	for i := size - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			tmp[j] = min(tmp[j], tmp[j+1]) + triangle[i][j]
		}
	}

	return tmp[0]
}

// @lc code=end

func Test_minimumTotal(t *testing.T) {
	tests := []struct {
		name     string
		triangle [][]int
		want     int
	}{
		{"0", [][]int{{0}}, 0},
		{"1", [][]int{{1}}, 1},
		{"2.1", [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}, 11},
		{"2.2", [][]int{{2}, {3, 4}, {6, 5, 1}, {4, 1, 8, 1}}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumTotal(tt.triangle); got != tt.want {
				t.Errorf("minimumTotal() = %v, want %v", got, tt.want)
			}
		})
	}
}
