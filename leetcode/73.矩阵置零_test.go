/*
 * @lc app=leetcode.cn id=73 lang=golang
 *
 * [73] 矩阵置零
 *
 * https://leetcode.cn/problems/set-matrix-zeroes/description/
 *
 * algorithms
 * Medium (64.25%)
 * Likes:    1005
 * Dislikes: 0
 * Total Accepted:    310K
 * Total Submissions: 476.4K
 * Testcase Example:  '[[1,1,1],[1,0,1],[1,1,1]]'
 *
 * 给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 原地 算法。
 *
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：matrix = [[1,1,1],[1,0,1],[1,1,1]]
 * 输出：[[1,0,1],[0,0,0],[1,0,1]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
 * 输出：[[0,0,0,0],[0,4,5,0],[0,3,1,0]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == matrix.length
 * n == matrix[0].length
 * 1 <= m, n <= 200
 * -2^31 <= matrix[i][j] <= 2^31 - 1
 *
 *
 *
 *
 * 进阶：
 *
 *
 * 一个直观的解决方案是使用  O(mn) 的额外空间，但这并不是一个好的解决方案。
 * 一个简单的改进方案是使用 O(m + n) 的额外空间，但这仍然不是最好的解决方案。
 * 你能想出一个仅使用常量空间的解决方案吗？
 *
 *
 */
package leetcode

import (
	"testing"
)

// @lc code=start
func setZeroes(matrix [][]int) {
	setZeroesFirst(matrix)
	// setZeroesArray(matrix)
}

func setZeroesFirst(matrix [][]int) {
	// 使用第一行和第一列代替标记数组
	firstRow0 := false // 第一行是否包含 0
	firstCol0 := false

	m := len(matrix)
	n := len(matrix[0])

	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			firstCol0 = true
			break
		}
	}

	for i := 0; i < n; i++ {
		if matrix[0][i] == 0 {
			firstRow0 = true
			break
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if firstRow0 {
		for i := 0; i < n; i++ {
			matrix[0][i] = 0
		}
	}

	if firstCol0 {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}

func setZeroesArray(matrix [][]int) {
	rowflag := make([]bool, len(matrix))
	colflag := make([]bool, len(matrix[0]))

	for i, r := range matrix {
		for j, v := range r {
			if v == 0 {
				rowflag[i] = true
				colflag[j] = true
			}
		}
	}

	for i, r := range matrix {
		for j := range r {
			if rowflag[i] || colflag[j] {
				r[j] = 0
			}
		}
	}

}

// 输入会超过数字上限
func setZeroes1(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])

	rowflag := 0
	colflag := 0

	for i, row := range matrix {
		for j, v := range row {
			if v == 0 {
				rowflag |= 1 << i
				colflag |= 1 << j
			}
		}
	}

	for i := 0; i < m; i++ {
		if rowflag>>i&1 == 1 {
			for j := 0; j < n; j++ {
				matrix[i][j] = 0
			}
		}
	}

	for j := 0; j < n; j++ {
		if colflag>>j&1 == 1 {
			for i := 0; i < m; i++ {
				matrix[i][j] = 0
			}
		}
	}
}

// @lc code=end

func Test_setZeroes(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   [][]int
	}{
		{"1", [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}, [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}},
		{"2", [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}, [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}}},
		// {"", [][]int{}, [][]int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setZeroes(tt.matrix)
			if !equalSliceMatrix[int](tt.matrix, tt.want) {
				t.Errorf("setZeroes() = %v, want %v", tt.matrix, tt.want)
			}
		})
	}
}
