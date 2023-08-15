/*
 * @lc app=leetcode.cn id=59 lang=golang
 *
 * [59] 螺旋矩阵 II
 *
 * https://leetcode.cn/problems/spiral-matrix-ii/description/
 *
 * algorithms
 * Medium (72.73%)
 * Likes:    1107
 * Dislikes: 0
 * Total Accepted:    337K
 * Total Submissions: 464.7K
 * Testcase Example:  '3'
 *
 * 给你一个正整数 n ，生成一个包含 1 到 n^2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 3
 * 输出：[[1,2,3],[8,9,4],[7,6,5]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 1
 * 输出：[[1]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= 20
 *
 *
 */
package leetcode

import (
	"reflect"
	"testing"
)

// @lc code=start
func generateMatrix(n int) [][]int {
	martrix := make([][]int, n)
	for i := range martrix {
		martrix[i] = make([]int, n)
	}

	leftBound, rightBound := 0, n-1
	upperBound, bottomBound := 0, n-1

	val := 1
	for val <= n*n {
		if upperBound <= bottomBound {
			for i := leftBound; i <= rightBound; i++ {
				martrix[upperBound][i] = val
				val++
			}
			upperBound++
		}
		if leftBound <= rightBound {
			for i := upperBound; i <= bottomBound; i++ {
				martrix[i][rightBound] = val
				val++
			}
			rightBound--
		}
		if upperBound <= bottomBound {
			for i := rightBound; i >= leftBound; i-- {
				martrix[bottomBound][i] = val
				val++
			}
			bottomBound--
		}
		if leftBound <= rightBound {
			for i := bottomBound; i >= upperBound; i-- {
				martrix[i][leftBound] = val
				val++
			}
			leftBound++
		}
	}
	return martrix
}

// @lc code=end

func Test_generateMatrix(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want [][]int
	}{
		{"1", 1, [][]int{{1}}},
		{"1", 2, [][]int{{1, 2}, {4, 3}}},
		{"1", 3, [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateMatrix(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
