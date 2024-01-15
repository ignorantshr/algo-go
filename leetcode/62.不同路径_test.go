/*
 * @lc app=leetcode.cn id=62 lang=golang
 *
 * [62] 不同路径
 *
 * https://leetcode.cn/problems/unique-paths/description/
 *
 * algorithms
 * Medium (67.88%)
 * Likes:    1973
 * Dislikes: 0
 * Total Accepted:    716.7K
 * Total Submissions: 1.1M
 * Testcase Example:  '3\n7'
 *
 * 一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
 *
 * 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
 *
 * 问总共有多少条不同的路径？
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：m = 3, n = 7
 * 输出：28
 *
 * 示例 2：
 *
 *
 * 输入：m = 3, n = 2
 * 输出：3
 * 解释：
 * 从左上角开始，总共有 3 条路径可以到达右下角。
 * 1. 向右 -> 向下 -> 向下
 * 2. 向下 -> 向下 -> 向右
 * 3. 向下 -> 向右 -> 向下
 *
 *
 * 示例 3：
 *
 *
 * 输入：m = 7, n = 3
 * 输出：28
 *
 *
 * 示例 4：
 *
 *
 * 输入：m = 3, n = 3
 * 输出：6
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= m, n <= 100
 * 题目数据保证答案小于等于 2 * 10^9
 *
 *
 */
package leetcode

import "testing"

// @lc code=start
func uniquePaths(m int, n int) int {
	// dp improve space
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] = dp[j] + dp[j-1]
		}
	}

	return dp[n-1]
}

func uniquePathsInitSingle(m int, n int) int {
	// dp improve space
	dp := make([]int, n)
	// for i := 0; i < n; i++ {
	// 	dp[i] = 1
	// }

	for i := 0; i < m; i++ {
		dp[0] = 1
		for j := 1; j < n; j++ {
			dp[j] = dp[j] + dp[j-1]
		}
	}

	return dp[n-1]
}

func uniquePathsDp(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}

func uniquePathsMineDp(m int, n int) int {
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		dp[i][1] = 1
	}
	for i := 1; i <= n; i++ {
		dp[1][i] = 1
	}

	// . .
	// x x /
	sumrow := func(r, c int) int {
		sum := 0
		for i := 1; i <= c; i++ {
			sum += dp[r][i]
		}
		return sum
	}

	// . x
	// . x
	//   /
	sumcol := func(r, c int) int {
		sum := 0
		for i := 1; i <= r; i++ {
			sum += dp[i][c]
		}
		return sum
	}

	for i := 2; i <= m; i++ {
		for j := 2; j <= n; j++ {
			dp[i][j] = sumrow(i-1, j-1) + sumcol(i-1, j-1)
		}
	}

	return dp[m][n]
}

func uniquePathsTimeout(m int, n int) int {
	if m == 1 && n == 1 {
		return 1
	}
	if m == 0 || n == 0 {
		return 0
	}
	return uniquePathsTimeout(m-1, n) + uniquePathsTimeout(m, n-1)
}

// @lc code=end

func Test_uniquePaths(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{1, 1}, 1},
		{"1.1", args{2, 1}, 1},
		{"1.2", args{1, 2}, 1},
		{"2", args{3, 3}, 6},
		{"2.1", args{3, 2}, 3},
		{"2.2", args{2, 3}, 3},
		{"3", args{3, 7}, 28},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePaths(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("uniquePaths() = %v, want %v", got, tt.want)
			}
			if got := uniquePathsInitSingle(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("uniquePathsInitSingle() = %v, want %v", got, tt.want)
			}
			if got := uniquePathsDp(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("uniquePathsDp() = %v, want %v", got, tt.want)
			}
			if got := uniquePathsMineDp(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("uniquePathsMineDp() = %v, want %v", got, tt.want)
			}
			if got := uniquePathsTimeout(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("uniquePathsTimeout() = %v, want %v", got, tt.want)
			}
		})
	}
}
