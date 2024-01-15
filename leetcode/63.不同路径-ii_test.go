/*
 * @lc app=leetcode.cn id=63 lang=golang
 *
 * [63] 不同路径 II
 *
 * https://leetcode.cn/problems/unique-paths-ii/description/
 *
 * algorithms
 * Medium (41.17%)
 * Likes:    1199
 * Dislikes: 0
 * Total Accepted:    440.5K
 * Total Submissions: 1.1M
 * Testcase Example:  '[[0,0,0],[0,1,0],[0,0,0]]'
 *
 * 一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
 *
 * 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish”）。
 *
 * 现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？
 *
 * 网格中的障碍物和空位置分别用 1 和 0 来表示。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：obstacleGrid = [[0,0,0],[0,1,0],[0,0,0]]
 * 输出：2
 * 解释：3x3 网格的正中间有一个障碍物。
 * 从左上角到右下角一共有 2 条不同的路径：
 * 1. 向右 -> 向右 -> 向下 -> 向下
 * 2. 向下 -> 向下 -> 向右 -> 向右
 *
 *
 * 示例 2：
 *
 *
 * 输入：obstacleGrid = [[0,1],[0,0]]
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == obstacleGrid.length
 * n == obstacleGrid[i].length
 * 1 <= m, n <= 100
 * obstacleGrid[i][j] 为 0 或 1
 *
 *
 */
package leetcode

import "testing"

// @lc code=start
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	dp := make([]int, n)
	if obstacleGrid[0][0] == 0 {
		dp[0] = 1
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[j] = 0
				continue
			}

			if j > 0 && obstacleGrid[i][j-1] == 0 {
				dp[j] = dp[j] + dp[j-1]
			}
		}
	}

	return dp[n-1]
}

func uniquePathsWithObstaclesDpMine(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		if i > 0 {
			dp[i] = min(dp[i-1], 1-obstacleGrid[0][i])
		} else {
			dp[i] = 1 - obstacleGrid[0][i]
		}
	}

	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[j] = 0
				continue
			}

			if j > 0 {
				dp[j] = dp[j] + dp[j-1]
			}
		}
	}

	return dp[n-1]
}

func uniquePathsWithObstaclesDp(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		if i > 0 {
			dp[i][0] = min(dp[i-1][0], 1-obstacleGrid[i][0])
		} else {
			dp[i][0] = 1 - obstacleGrid[i][0]
		}
	}
	for i := 0; i < n; i++ {
		if i > 0 {
			dp[0][i] = min(dp[0][i-1], 1-obstacleGrid[0][i])
		} else {
			dp[0][i] = 1 - obstacleGrid[0][i]
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	return dp[m-1][n-1]
}

// @lc code=end

func Test_uniquePathsWithObstacles(t *testing.T) {
	type args struct {
		obstacleGrid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"x.3", args{[][]int{{1, 0}}}, 0},
		{"x", args{[][]int{{0}}}, 1},
		{"x.4", args{[][]int{{0}, {1}}}, 0},
		{"x.1", args{[][]int{{1, 1}}}, 0},
		{"x.2", args{[][]int{{0, 0, 0, 0}, {0, 1, 0, 0}, {0, 0, 0, 0}, {0, 0, 1, 0}, {0, 0, 0, 0}}}, 7},
		{"1", args{[][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}}, 2},
		{"1.1", args{[][]int{{0, 1}, {0, 0}}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePathsWithObstacles(tt.args.obstacleGrid); got != tt.want {
				t.Errorf("uniquePathsWithObstacles() = %v, want %v", got, tt.want)
			}
			if got := uniquePathsWithObstaclesDp(tt.args.obstacleGrid); got != tt.want {
				t.Errorf("uniquePathsWithObstaclesDp() = %v, want %v", got, tt.want)
			}
			if got := uniquePathsWithObstaclesDpMine(tt.args.obstacleGrid); got != tt.want {
				t.Errorf("uniquePathsWithObstaclesDpMine() = %v, want %v", got, tt.want)
			}
		})
	}
}
