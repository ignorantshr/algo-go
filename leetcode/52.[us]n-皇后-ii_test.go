/*
 * @lc app=leetcode.cn id=52 lang=golang
 *
 * [52] N 皇后 II
 *
 * https://leetcode.cn/problems/n-queens-ii/description/
 *
 * algorithms
 * Hard (82.35%)
 * Likes:    467
 * Dislikes: 0
 * Total Accepted:    128K
 * Total Submissions: 155.4K
 * Testcase Example:  '4'
 *
 * n 皇后问题 研究的是如何将 n 个皇后放置在 n × n 的棋盘上，并且使皇后彼此之间不能相互攻击。
 *
 * 给你一个整数 n ，返回 n 皇后问题 不同的解决方案的数量。
 *
 * 示例 1：
 * 输入：n = 4
 * 输出：2
 * 解释：如上图所示，4 皇后问题存在两个不同的解法。
 *
 * 示例 2：
 * 输入：n = 1
 * 输出：1
 *
 * 提示：
 *
 * 1 <= n <= 9
 */
package leetcode

import "testing"

// @lc code=start
func totalNQueens(n int) int {
	chosed := make([][]bool, n)
	for i := range chosed {
		chosed[i] = make([]bool, n)
	}
	rlock := make([]bool, n)
	clock := make([]bool, n)
	ldlock := make([]bool, 2*n-1)
	rdlock := make([]bool, 2*n-1)
	res := 0

	var walk func(row int)
	walk = func(row int) {
		if row == n {
			res++
			return
		}

		for j := 0; j < n; j++ {
			if clock[j] || ldlock[row-j+n-1] || rdlock[row+j] {
				continue
			}
			chosed[row][j] = true
			rlock[row] = true
			clock[j] = true
			ldlock[row-j+n-1] = true
			rdlock[row+j] = true
			walk(row + 1)
			chosed[row][j] = false
			rlock[row] = false
			clock[j] = false
			ldlock[row-j+n-1] = false
			rdlock[row+j] = false
		}
	}
	walk(0)
	return res
}

// @lc code=end

func Test_totalNQueens(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"1", 1, 1},
		{"2", 2, 0},
		{"3", 3, 0},
		{"4", 4, 2},
		{"5", 5, 10},
		{"6", 6, 4},
		{"7", 7, 40},
		{"8", 8, 92},
		{"9", 9, 352},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := totalNQueens(tt.n); got != tt.want {
				t.Errorf("totalNQueens() = %v, want %v", got, tt.want)
			}
		})
	}
}
