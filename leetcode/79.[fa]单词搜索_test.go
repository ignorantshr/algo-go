/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] 单词搜索
 *
 * https://leetcode.cn/problems/word-search/description/
 *
 * algorithms
 * Medium (46.40%)
 * Likes:    1751
 * Dislikes: 0
 * Total Accepted:    481.8K
 * Total Submissions: 1M
 * Testcase Example:  '[["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]\n"ABCCED"'
 *
 * 给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false
 * 。
 *
 * 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word =
 * "ABCCED"
 * 输出：true
 *
 *
 * 示例 2：
 *
 *
 * 输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word =
 * "SEE"
 * 输出：true
 *
 *
 * 示例 3：
 *
 *
 * 输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word =
 * "ABCB"
 * 输出：false
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == board.length
 * n = board[i].length
 * 1
 * 1
 * board 和 word 仅由大小写英文字母组成
 *
 *
 *
 *
 * 进阶：你可以使用搜索剪枝的技术来优化解决方案，使其在 board 更大的情况下可以更快解决问题？
 *
 */
package leetcode

import "testing"

// @lc code=start
func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	directions := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	var backtrack func(row, col, idx int) bool
	backtrack = func(row, col, idx int) bool {
		if board[row][col] != word[idx] {
			return false
		}
		if idx+1 == len(word) {
			return true
		}

		visited[row][col] = true
		for _, d := range directions {
			nrow, ncol := row+d[0], col+d[1]
			if 0 <= nrow && nrow < m && 0 <= ncol && ncol < n && !visited[nrow][ncol] {
				visited[nrow][ncol] = true
				if backtrack(nrow, ncol, idx+1) {
					return true
				}
				visited[nrow][ncol] = false
			}
		}
		visited[row][col] = false
		return false
	}

	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if backtrack(row, col, 0) {
				return true
			}
		}
	}

	return false
}

// @lc code=end

func Test_exist(t *testing.T) {
	type args struct {
		board [][]byte
		word  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"x.4", args{[][]byte{
			{'a', 'a'}, {'a', 'a'},
		}, "aaaaa"}, false},
		{"x.3", args{[][]byte{
			{'a', 'a', 'a', 'a'}, {'a', 'a', 'a', 'a'}, {'a', 'a', 'a', 'a'},
		}, "aaaaaaaaaaaaa"}, false},
		{"x.2", args{[][]byte{
			{'A'},
		}, "A"}, true},
		{"x.1", args{[][]byte{
			{'A', 'B'},
			{'C', 'D'},
		}, "CDBA"}, true},
		{"0", args{[][]byte{
			{'A', 'B'},
			{'S', 'F'},
		}, "AB"}, true},
		{"1", args{[][]byte{
			{'A', 'B', 'C', 'E'},
			{'S', 'F', 'C', 'S'},
			{'A', 'D', 'E', 'E'},
		}, "ABCCED"}, true},
		{"1.1", args{[][]byte{
			{'A', 'B', 'C', 'E'},
			{'S', 'F', 'C', 'S'},
			{'A', 'D', 'E', 'E'},
		}, "SEE"}, true},
		{"1.2", args{[][]byte{
			{'A', 'B', 'C', 'E'},
			{'S', 'F', 'C', 'S'},
			{'A', 'D', 'E', 'E'},
		}, "ABCB"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exist(tt.args.board, tt.args.word); got != tt.want {
				t.Errorf("exist() = %v, want %v", got, tt.want)
			}
		})
	}
}
