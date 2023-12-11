/*
 * @lc app=leetcode.cn id=36 lang=golang
 *
 * [36] 有效的数独
 *
 * https://leetcode.cn/problems/valid-sudoku/description/
 *
 * algorithms
 * Medium (63.10%)
 * Likes:    1186
 * Dislikes: 0
 * Total Accepted:    406.3K
 * Total Submissions: 643.2K
 * Testcase Example:  '[["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]'
 *
 * 请你判断一个 9 x 9 的数独是否有效。只需要 根据以下规则 ，验证已经填入的数字是否有效即可。
 *
 *
 * 数字 1-9 在每一行只能出现一次。
 * 数字 1-9 在每一列只能出现一次。
 * 数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。（请参考示例图）
 *
 *
 *
 *
 * 注意：
 *
 *
 * 一个有效的数独（部分已被填充）不一定是可解的。
 * 只需要根据以上规则，验证已经填入的数字是否有效即可。
 * 空白格用 '.' 表示。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：board =
 * [["5","3",".",".","7",".",".",".","."]
 * ,["6",".",".","1","9","5",".",".","."]
 * ,[".","9","8",".",".",".",".","6","."]
 * ,["8",".",".",".","6",".",".",".","3"]
 * ,["4",".",".","8",".","3",".",".","1"]
 * ,["7",".",".",".","2",".",".",".","6"]
 * ,[".","6",".",".",".",".","2","8","."]
 * ,[".",".",".","4","1","9",".",".","5"]
 * ,[".",".",".",".","8",".",".","7","9"]]
 * 输出：true
 *
 *
 * 示例 2：
 *
 *
 * 输入：board =
 * [["8","3",".",".","7",".",".",".","."]
 * ,["6",".",".","1","9","5",".",".","."]
 * ,[".","9","8",".",".",".",".","6","."]
 * ,["8",".",".",".","6",".",".",".","3"]
 * ,["4",".",".","8",".","3",".",".","1"]
 * ,["7",".",".",".","2",".",".",".","6"]
 * ,[".","6",".",".",".",".","2","8","."]
 * ,[".",".",".","4","1","9",".",".","5"]
 * ,[".",".",".",".","8",".",".","7","9"]]
 * 输出：false
 * 解释：除了第一行的第一个数字从 5 改为 8 以外，空格内其他数字均与 示例1 相同。 但由于位于左上角的 3x3 宫内有两个 8 存在,
 * 因此这个数独是无效的。
 *
 *
 *
 * 提示：
 *
 *
 * board.length == 9
 * board[i].length == 9
 * board[i][j] 是一位数字（1-9）或者 '.'
 *
 *
 */
package leetcode

import (
	"fmt"
	"testing"
)

// @lc code=start
func isValidSudoku(board [][]byte) bool {
	// 统计哈希表中每个元素出现的次数
	rows, cols := [9][9]int{}, [9][9]int{}
	subs := [3][3][9]int{}

	for i, row := range board {
		for j, c := range row {
			if c == '.' {
				continue
			}

			n := c - '1'
			rows[i][n]++
			cols[j][n]++
			subs[i/3][j/3][n]++
			if rows[i][n] > 1 || cols[j][n] > 1 || subs[i/3][j/3][n] > 1 {
				return false
			}
		}
	}

	return true
}

func isValidSudokuBit(board [][]byte) bool {
	// 位运算判断存在
	rows, cols := [9]int{}, [9]int{}
	subs := [9]int{}

	for i, row := range board {
		for j, c := range row {
			if c == '.' {
				continue
			}

			idx := 1 << (c - '1')
			if rows[i]&idx > 0 || cols[j]&idx > 0 || subs[(i/3)*3+j/3]&idx > 0 {
				return false
			}
			rows[i] |= idx
			cols[j] |= idx
			subs[(i/3)*3+j/3] |= idx
		}
	}

	return true
}

func isValidSudokuMe(board [][]byte) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			if board[i][j] != '.' && !checkSudoku(board, i, j) {
				return false
			}
		}
	}

	return true
}

func checkSudoku(board [][]byte, row, col int) bool {
	for i := 0; i < len(board); i++ {
		if i != col && board[row][i] == board[row][col] {
			return false
		}

		if i != row && board[i][col] == board[row][col] {
			return false
		}

		x := 3*(row/3) + i/3
		y := 3*(col/3) + i%3
		if (x != row || y != col) && board[x][y] == board[row][col] {
			return false
		}
	}

	return true
}

// @lc code=end

// 回溯填充找解
func fillSudoku(board [][]byte, row, col int) bool {
	if board[row][col] != '.' {
		if !checkSudoku(board, row, col) {
			return false
		}

		row, col := nextCoordinate(row, col)
		if row == -1 {
			return true
		}
		return fillSudoku(board, row, col)
	}

	for i := 1; i <= len(board); i++ {
		board[row][col] = byte(i + '0')
		if fillSudoku(board, row, col) {
			return true
		}
		board[row][col] = '.' // ⚠️此处需要改回 . ，因为前面的字符回溯时还需要对这个位置进行重新判断
	}

	// boardDisplay(board)
	return false
}

func nextCoordinate(row, col int) (int, int) {
	col++
	if col == 9 {
		row++
		col = 0
	}
	if row == 9 {
		return -1, -1
	}

	return row, col
}

func boardDisplay(board [][]byte) {
	for i := 0; i < len(board); i++ {
		fmt.Println(board[i])
	}
	fmt.Println("----------------")
}

func Test_isValidSudoku(t *testing.T) {
	tests := []struct {
		name  string
		board [][]byte
		want  bool
	}{
		{"1",
			[][]byte{
				{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
			true,
		},
		{"2",
			[][]byte{
				{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
			false,
		},
		{"3",
			[][]byte{
				{'.', '8', '7', '6', '5', '4', '3', '2', '1'},
				{'2', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'3', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'4', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'5', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'6', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'7', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'8', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'9', '.', '.', '.', '.', '.', '.', '.', '.'},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidSudoku(tt.board); got != tt.want {
				t.Errorf("isValidSudoku() = %v, want %v", got, tt.want)
			}
			if got := isValidSudokuBit(tt.board); got != tt.want {
				t.Errorf("isValidSudokuBit() = %v, want %v", got, tt.want)
			}
			if got := isValidSudokuMe(tt.board); got != tt.want {
				t.Errorf("isValidSudokuMe() = %v, want %v", got, tt.want)
			}
		})
	}
}
