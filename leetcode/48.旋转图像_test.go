/*
 * @lc app=leetcode.cn id=48 lang=golang
 *
 * [48] 旋转图像
 *
 * https://leetcode.cn/problems/rotate-image/description/
 *
 * algorithms
 * Medium (74.75%)
 * Likes:    1686
 * Dislikes: 0
 * Total Accepted:    469K
 * Total Submissions: 627K
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]'
 *
 * 给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。
 *
 * 你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
 * 输出：[[7,4,1],[8,5,2],[9,6,3]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
 * 输出：[[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * n == matrix.length == matrix[i].length
 * 1 <= n <= 20
 * -1000 <= matrix[i][j] <= 1000
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
func rotate(matrix [][]int) {
	// 思路：把行变成列,列变成行

	n := len(matrix)
	// 对角线翻转
	for i := range matrix {
		for j := i + 1; j < n; j++ {
			diagonalFlip(matrix, i, j)
		}
	}
	// 行翻转
	for i := range matrix {
		for j := 0; j < n/2; j++ {
			exchange(matrix, i, j, i, n-1-j)
		}
	}
}

// 沿 斜率为-1的对角线 翻转
func diagonalFlip(matrix [][]int, x, y int) {
	exchange(matrix, x, y, y, x)
}

func exchange(matrix [][]int, x1, y1, x2, y2 int) {
	matrix[x1][y1], matrix[x2][y2] = matrix[x2][y2], matrix[x1][y1]
}

// @lc code=end

func Test_rotate(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   [][]int
	}{
		{"0",
			[][]int{},
			[][]int{},
		},
		{"1",
			[][]int{{1}},
			[][]int{{1}},
		},
		{"1",
			[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			[][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
		},
		{"1",
			[][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}},
			[][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.matrix)
			if !reflect.DeepEqual(tt.matrix, tt.want) {
				t.Fatalf("rotate = %v, want = %v", tt.matrix, tt.want)
			}
		})
	}
}
