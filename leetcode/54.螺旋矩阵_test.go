/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 *
 * https://leetcode.cn/problems/spiral-matrix/description/
 *
 * algorithms
 * Medium (49.34%)
 * Likes:    1439
 * Dislikes: 0
 * Total Accepted:    387.5K
 * Total Submissions: 784.3K
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]'
 *
 * 给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
 * 输出：[1,2,3,6,9,8,7,4,5]
 *
 *
 * 示例 2：
 *
 *
 * 输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
 * 输出：[1,2,3,4,8,12,11,10,9,5,6,7]
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == matrix.length
 * n == matrix[i].length
 * 1 <= m, n <= 10
 * -100 <= matrix[i][j] <= 100
 *
 *
 */
package leetcode

import (
	"reflect"
	"testing"
)

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	upperBound, bottomBound := 0, len(matrix)-1
	leftBound, rightBound := 0, len(matrix[0])-1
	m, n := bottomBound+1, rightBound+1

	var res []int
	for len(res) < m*n {
		if upperBound <= bottomBound {
			for j := leftBound; j <= rightBound; j++ {
				// to right
				res = append(res, matrix[upperBound][j])
			}
			upperBound++
		}
		if leftBound <= rightBound {
			for i := upperBound; i <= bottomBound; i++ {
				// to bottom
				res = append(res, matrix[i][rightBound])
			}
			rightBound--
		}
		if upperBound <= bottomBound {
			for j := rightBound; j >= leftBound; j-- {
				// to left
				res = append(res, matrix[bottomBound][j])
			}
			bottomBound--
		}
		if leftBound <= rightBound {
			for i := bottomBound; i >= upperBound; i-- {
				// to upper
				res = append(res, matrix[i][leftBound])
			}
			leftBound++
		}
	}
	return res
}

// @lc code=end

func Test_spiralOrder(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   []int
	}{
		{"1", [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}, []int{1, 2, 3, 4, 8, 12, 16, 15, 14, 13, 9, 5, 6, 7, 11, 10}},
		{"1", [][]int{{1}}, []int{1}},
		{"1", [][]int{{1, 2, 3}}, []int{1, 2, 3}},
		{"1", [][]int{{1}, {2}, {3}}, []int{1, 2, 3}},
		{"1", [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
		{"1", [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}, []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := spiralOrder(tt.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("spiralOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
