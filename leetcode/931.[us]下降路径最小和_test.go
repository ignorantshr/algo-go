/*
 * @lc app=leetcode.cn id=931 lang=golang
 *
 * [931] 下降路径最小和
 *
 * https://leetcode.cn/problems/minimum-falling-path-sum/description/
 *
 * algorithms
 * Medium (67.96%)
 * Likes:    316
 * Dislikes: 0
 * Total Accepted:    88.7K
 * Total Submissions: 130.8K
 * Testcase Example:  '[[2,1,3],[6,5,4],[7,8,9]]'
 *
 * 给你一个 n x n 的 方形 整数数组 matrix ，请你找出并返回通过 matrix 的下降路径 的 最小和 。
 *
 * 下降路径
 * 可以从第一行中的任何元素开始，并从每一行中选择一个元素。在下一行选择的元素和当前行所选元素最多相隔一列（即位于正下方或者沿对角线向左或者向右的第一个元素）。具体来说，位置
 * (row, col) 的下一个元素应当是 (row + 1, col - 1)、(row + 1, col) 或者 (row + 1, col + 1)
 * 。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：matrix = [[2,1,3],[6,5,4],[7,8,9]]
 * 输出：13
 * 解释：如图所示，为和最小的两条下降路径
 *
 *
 * 示例 2：
 *
 *
 *
 *
 * 输入：matrix = [[-19,57],[-40,-5]]
 * 输出：-59
 * 解释：如图所示，为和最小的下降路径
 *
 *
 *
 *
 * 提示：
 *
 *
 * n == matrix.length == matrix[i].length
 * 1 <= n <= 100
 * -100 <= matrix[i][j] <= 100
 *
 *
 */
package leetcode

import "testing"

// @lc code=start
func minFallingPathSum(matrix [][]int) int {
	size := len(matrix)
	// dp := make([][]int, size) // 压缩到 2xN 维数组
	dp := [2][]int{}

	for i := 0; i < 2; i++ {
		dp[i] = make([]int, size)
		if i == 0 {
			copy(dp[i], matrix[i])
		}
	}

	for i := 1; i < size; i++ {
		for j := 0; j < size; j++ {
			minp := dp[0][j]
			if j != size-1 {
				minp = min(minp, dp[0][j+1])
			}
			if j != 0 {
				minp = min(minp, dp[0][j-1])
			}
			dp[1][j] = matrix[i][j] + minp
		}
		copy(dp[0], dp[1])
	}

	res := dp[0][0] // n==1 时，只有 dp[0] 存储结果
	for i := size - 1; i > 0; i-- {
		res = min(res, dp[0][i])
	}
	return res
}

// @lc code=end

func Test_minFallingPathSum(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   int
	}{
		{"1", [][]int{{57}}, 57},
		{"1", [][]int{{-19, 57}, {-40, -5}}, -59},
		{"1", [][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}}, 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minFallingPathSum(tt.matrix); got != tt.want {
				t.Errorf("minFallingPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
