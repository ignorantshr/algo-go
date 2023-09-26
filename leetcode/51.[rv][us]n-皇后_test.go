/*
 * @lc app=leetcode.cn id=51 lang=golang
 *
 * [51] N 皇后
 *
 * https://leetcode.cn/problems/n-queens/description/
 *
 * algorithms
 * Hard (74.06%)
 * Likes:    1894
 * Dislikes: 0
 * Total Accepted:    332.9K
 * Total Submissions: 449.9K
 * Testcase Example:  '4'
 *
 * 按照国际象棋的规则，皇后可以攻击与之处在同一行或同一列或同一斜线上的棋子。
 *
 * n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
 *
 * 给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。
 *
 * 每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
 *
 * 示例 1：
 * 输入：n = 4
 * 输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
 * 解释：如上图所示，4 皇后问题存在两个不同的解法。
 *
 * 示例 2：
 * 输入：n = 1
 * 输出：[["Q"]]
 *
 * 提示：
 * 1 <= n <= 9
 */
package leetcode

import (
	"reflect"
	"testing"
)

// @lc code=start
func solveNQueens(n int) [][]string {
	res := make([][]string, 0)
	chosed := make([][]bool, n)
	for i := 0; i < n; i++ {
		chosed[i] = make([]bool, n)
	}
	rowLock := make([]bool, n)
	columnLock := make([]bool, n)
	ldiagonalLock := make([]bool, 2*n) // 左上到右下，差不变
	rdiagonalLock := make([]bool, 2*n) // 右上到左下，和不变

	var put func(row int)
	put = func(row int) {
		if row == n {
			tmp := make([]string, n)
			for i := 0; i < n; i++ {
				rowstr := ""
				for j := 0; j < n; j++ {
					if chosed[i][j] {
						rowstr += "Q"
					} else {
						rowstr += "."
					}
				}
				tmp[i] = rowstr
			}
			res = append(res, tmp)
			return
		}

		// for i := row; i < row+1; i++ { // 为了防止以前遍历过的行之后又被重新遍历，可以选择从下一行开始遍历，直接跳过它们
		// 	if rowLock[i] {
		// 		continue
		// 	}
		// 对于 i:=0;i<n; 遍历的补充条件，一整行都没有能放置的位置，往下寻找答案时此行不能再被遍历，
		// 因为后面的行可能存在可放置位置，然后寻找下一个位置时重新遍历此行，造成结果重复
		// if count < i {
		// 	return
		// }
		i := row
		for j := 0; j < n; j++ {
			if columnLock[j] || ldiagonalLock[i-j+n-1] || rdiagonalLock[i+j] {
				continue
			}
			chosed[i][j] = true
			rowLock[i] = true
			columnLock[j] = true
			ldiagonalLock[i-j+n-1] = true
			rdiagonalLock[i+j] = true
			put(row + 1)
			chosed[i][j] = false
			rowLock[i] = false
			columnLock[j] = false
			ldiagonalLock[i-j+n-1] = false
			rdiagonalLock[i+j] = false
		}
	}
	put(0)
	return res
}

// @lc code=end

func Test_solveNQueens(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want [][]string
	}{
		// {"1", 1, [][]string{{"Q"}}},
		{"4", 4, [][]string{{".Q..", "...Q", "Q...", "..Q."}, {"..Q.", "Q...", "...Q", ".Q.."}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveNQueens_RV(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solveNQueens() = %v, want %v", got, tt.want)
			}
		})
	}
}

func solveNQueens_RV(n int) [][]string {
	res := make([][]string, 0)
	choosed := make([][]byte, n)
	for i := range choosed {
		choosed[i] = make([]byte, n)
	}

	valid := func(row, col int) bool {
		// 算法是每行取一个，就不用校验行有效性了
		for i, j1, j2 := row-1, col-1, col+1; i >= 0; {
			if j1 >= 0 {
				if choosed[i][j1] == 'Q' { // 主对角线
					return false
				}
			}
			if j2 < n {
				if choosed[i][j2] == 'Q' { // 副对角线
					return false
				}
			}
			i--
			j1--
			j2++
		}
		for i := row - 1; i >= 0; i-- {
			if choosed[i][col] == 'Q' { // 同列
				return false
			}
		}
		return true
	}

	var backtrace func(row int)
	backtrace = func(row int) {
		if row == n {
			tmp := make([]string, n)
			for i, row := range choosed {
				for _, v := range row {
					if v != 'Q' {
						tmp[i] += "."
					} else {
						tmp[i] += "Q"
					}
				}
			}
			res = append(res, tmp)
			return
		}

		// 遍历这行的每一列，找到一个符合的位置
		for j := 0; j < n; j++ {
			if !valid(row, j) {
				choosed[row][j] = '.'
				continue
			}

			choosed[row][j] = 'Q'
			backtrace(row + 1)
			choosed[row][j] = '.'
		}
	}
	backtrace(0)
	return res
}
