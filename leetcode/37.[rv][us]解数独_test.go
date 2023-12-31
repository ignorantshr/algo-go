/*
 * @lc app=leetcode.cn id=37 lang=golang
 *
 * [37] 解数独
 *
 * https://leetcode.cn/problems/sudoku-solver/description/
 *
 * algorithms
 * Hard (67.56%)
 * Likes:    1719
 * Dislikes: 0
 * Total Accepted:    220.5K
 * Total Submissions: 326.3K
 * Testcase Example:  '[["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]'
 *
 * 编写一个程序，通过填充空格来解决数独问题。
 *
 * 数独的解法需 遵循如下规则：
 *
 *
 * 数字 1-9 在每一行只能出现一次。
 * 数字 1-9 在每一列只能出现一次。
 * 数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。（请参考示例图）
 *
 *
 * 数独部分空格内已填入了数字，空白格用 '.' 表示。
 *
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：board =
 * [["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]
 *
 * 输出：[["5","3","4","6","7","8","9","1","2"],["6","7","2","1","9","5","3","4","8"],["1","9","8","3","4","2","5","6","7"],["8","5","9","7","6","1","4","2","3"],["4","2","6","8","5","3","7","9","1"],["7","1","3","9","2","4","8","5","6"],["9","6","1","5","3","7","2","8","4"],["2","8","7","4","1","9","6","3","5"],["3","4","5","2","8","6","1","7","9"]]
 * 解释：输入的数独如上图所示，唯一有效的解决方案如下所示：
 *
 *
 *
 *
 *
 *
 * 提示：
 *
 *
 * board.length == 9
 * board[i].length == 9
 * board[i][j] 是一位数字或者 '.'
 * 题目数据 保证 输入数独仅有一个解
 *
 *
 *
 *
 *
 */
package leetcode

import (
	"reflect"
	"testing"
)

// @lc code=start
func solveSudoku(board [][]byte) {
	next := func(i, j int) (int, int) {
		if j < 8 {
			j++
			return i, j
		}
		i++
		j = 0
		return i, j
	}

	valid := func(i, j int, b byte) bool {
		for m := 0; m < 9; m++ {
			if b == board[i][m] {
				return false
			}
			if b == board[m][j] {
				return false
			}
			if board[i/3*3+m/3][j/3*3+m%3] == b {
				return false
			}
		}
		return true
	}

	var sovle func(i, j int) bool
	sovle = func(i, j int) bool {
		for i != 9 && board[i][j] != '.' {
			i, j = next(i, j)
		}
		if i == 9 {
			return true
		}
		for n := '1'; n <= '9'; n++ {
			if valid(i, j, byte(n)) {
				board[i][j] = byte(n)
				if sovle(next(i, j)) {
					return true
				}
				board[i][j] = '.'
			}
		}
		return false
	}
	sovle(0, 0)
}

// @lc code=end

func Test_solveSudoku(t *testing.T) {
	tests := []struct {
		name  string
		board [][]byte
		want  [][]byte
	}{
		{"1", [][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '.', '.', '1', '9', '5', '.', '.', '.'}, {'.', '9', '8', '.', '.', '.', '.', '6', '.'}, {'8', '.', '.', '.', '6', '.', '.', '.', '3'}, {'4', '.', '.', '8', '.', '3', '.', '.', '1'}, {'7', '.', '.', '.', '2', '.', '.', '.', '6'}, {'.', '6', '.', '.', '.', '.', '2', '8', '.'}, {'.', '.', '.', '4', '1', '9', '.', '.', '5'}, {'.', '.', '.', '.', '8', '.', '.', '7', '9'}},
			[][]byte{{'5', '3', '4', '6', '7', '8', '9', '1', '2'}, {'6', '7', '2', '1', '9', '5', '3', '4', '8'}, {'1', '9', '8', '3', '4', '2', '5', '6', '7'}, {'8', '5', '9', '7', '6', '1', '4', '2', '3'}, {'4', '2', '6', '8', '5', '3', '7', '9', '1'}, {'7', '1', '3', '9', '2', '4', '8', '5', '6'}, {'9', '6', '1', '5', '3', '7', '2', '8', '4'}, {'2', '8', '7', '4', '1', '9', '6', '3', '5'}, {'3', '4', '5', '2', '8', '6', '1', '7', '9'}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if solveSudoku_RV(tt.board); !reflect.DeepEqual(tt.board, tt.want) {
				t.Errorf("solveSudoku() = %v, want %v", tt.board, tt.want)
			}
		})
	}
}

func solveSudoku_RV(board [][]byte) {
	dimension := len(board)
	finalPos := dimension * dimension

	valid := func(row, col int, b byte) bool {
		for i := 0; i < dimension; i++ {
			if board[row][i] == b {
				return false
			}
			if board[i][col] == b {
				return false
			}
			if board[row/3*3+i/3][col/3*3+i%3] == b {
				return false
			}
		}
		return true
	}

	var backtrace func(pos int) bool
	backtrace = func(pos int) bool {
		for pos < finalPos && board[pos/dimension][pos%dimension] != '.' {
			pos++
		}
		if pos == finalPos {
			return true
		}

		row := pos / dimension
		col := pos % dimension
		for n := '1'; n <= '9'; n++ {
			if !valid(row, col, byte(n)) {
				continue
			}

			board[row][col] = byte(n)
			if backtrace(pos + 1) {
				return true
			}
			board[row][col] = '.'
		}
		return false
	}
	backtrace(0)
}
