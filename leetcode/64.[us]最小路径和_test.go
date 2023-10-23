/*
 * @lc app=leetcode.cn id=64 lang=golang
 *
 * [64] 最小路径和
 *
 * https://leetcode.cn/problems/minimum-path-sum/description/
 *
 * algorithms
 * Medium (69.58%)
 * Likes:    1595
 * Dislikes: 0
 * Total Accepted:    526.7K
 * Total Submissions: 755.2K
 * Testcase Example:  '[[1,3,1],[1,5,1],[4,2,1]]'
 *
 * 给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
 *
 * 说明：每次只能向下或者向右移动一步。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
 * 输出：7
 * 解释：因为路径 1→3→1→1→1 的总和最小。
 *
 *
 * 示例 2：
 *
 *
 * 输入：grid = [[1,2,3],[4,5,6]]
 * 输出：12
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == grid.length
 * n == grid[i].length
 * 1 <= m, n <= 200
 * 0 <= grid[i][j] <= 200
 *
 *
 */
package leetcode

import "testing"

// @lc code=start
func minPathSum(grid [][]int) int {
	return minPathSumDpOptimum(grid)
	return minPathSumDp(grid)
}

func minPathSumDpOptimum(grid [][]int) int {
	rowSize := len(grid)
	colSize := len(grid[0])
	dp := make([]int, colSize)
	dp[0] = grid[0][0]
	for j := 1; j < colSize; j++ {
		dp[j] = dp[j-1] + grid[0][j] // base case 相当于dp[0][j-1]=sum(grid[0][0...j])
	}

	for i := 1; i < rowSize; i++ {
		dp[0] += grid[i][0] // base case  相当于dp[i][0]=sum(grid[0...i][0])
		for j := 1; j < colSize; j++ {
			// dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
			dp[j] = min(dp[j], dp[j-1]) + grid[i][j]
		}
	}

	return dp[colSize-1]
}

func minPathSumDp(grid [][]int) int {
	// d[i][j] 以 i,j 为终点的最小路径和
	rowSize := len(grid)
	colSize := len(grid[0])
	dp := make([][]int, rowSize)
	for i := 0; i < rowSize; i++ {
		dp[i] = make([]int, colSize)
		if i > 0 {
			dp[i][0] = dp[i-1][0] + grid[i][0]
		} else {
			dp[0][0] = grid[0][0]
		}
	}
	for j := 1; j < colSize; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}

	for i := 1; i < rowSize; i++ {
		for j := 1; j < colSize; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[rowSize-1][colSize-1]
}

// @lc code=end

func Test_minPathSum(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{"", [][]int{{1}}, 1},
		{"", [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}, 7},
		{"", [][]int{{1, 2, 3}, {4, 5, 6}}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minPathSum(tt.grid); got != tt.want {
				t.Errorf("minPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
