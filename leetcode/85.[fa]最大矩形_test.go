/*
 * @lc app=leetcode.cn id=85 lang=golang
 *
 * [85] 最大矩形
 *
 * https://leetcode.cn/problems/maximal-rectangle/description/
 *
 * algorithms
 * Hard (54.88%)
 * Likes:    1622
 * Dislikes: 0
 * Total Accepted:    190.4K
 * Total Submissions: 346.3K
 * Testcase Example:  '[["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]'
 *
 * 给定一个仅包含 0 和 1 、大小为 rows x cols 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：matrix =
 * [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
 * 输出：6
 * 解释：最大矩形如上图所示。
 *
 *
 * 示例 2：
 *
 *
 * 输入：matrix = [["0"]]
 * 输出：0
 *
 *
 * 示例 3：
 *
 *
 * 输入：matrix = [["1"]]
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * rows == matrix.length
 * cols == matrix[0].length
 * 1 <= row, cols <= 200
 * matrix[i][j] 为 '0' 或 '1'
 *
 *
 */
package leetcode

import (
	"testing"
)

// @lc code=start
// 单调栈
// 时间复杂度：O(mn)
func maximalRectangle(matrix [][]byte) int {
	m := len(matrix)
	n := len(matrix[0])
	left := make([][]int, m) // 在第 i 行中以 j 为结尾的连续方块个数
	ans := 0

	for i := 0; i < m; i++ {
		left[i] = make([]int, n)
		left[i][0] = int(matrix[i][0] - '0')
		for j := 1; j < n; j++ {
			if matrix[i][j] == '1' {
				left[i][j] = left[i][j-1] + 1
			}
		}
	}

	// 对 每一列柱子 应用84.题的单调栈算法
	for j := 0; j < n; j++ {
		stack := make([]int, 0)
		lpillar := make([]int, m)
		rpillar := make([]int, m)
		for i := 0; i < m; i++ {
			rpillar[i] = m
		}

		for i := 0; i < m; i++ {
			for len(stack) > 0 && left[stack[len(stack)-1]][j] >= left[i][j] {
				rpillar[stack[len(stack)-1]] = i
				stack = stack[:len(stack)-1]
			}
			if len(stack) != 0 {
				lpillar[i] = stack[len(stack)-1]
			} else {
				lpillar[i] = -1
			}
			stack = append(stack, i)
		}

		area := 0
		for i := 0; i < m; i++ {
			area = max(area, left[i][j]*(rpillar[i]-lpillar[i]-1))
		}
		ans = max(ans, area)
	}

	return ans
}

// 转化为柱状图计算面积
// 时间复杂度：O(m^2n)
func maximalRectangleEnum(matrix [][]byte) int {
	m := len(matrix)
	n := len(matrix[0])
	left := make([][]int, m) // 在第 i 行中以 j 为结尾的连续方块个数
	ans := 0

	for i := 0; i < m; i++ {
		left[i] = make([]int, n)
		left[i][0] = int(matrix[i][0] - '0')
		for j := 1; j < n; j++ {
			if matrix[i][j] == '1' {
				left[i][j] = left[i][j-1] + 1
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			width := left[i][j]
			if width == 0 {
				continue
			}
			area := width

			for k := i - 1; k >= 0 && left[k][j] > 0; k-- {
				width = min(left[k][j], width)
				area = max(area, width*(i-k+1))
			}
			ans = max(ans, area)
		}
	}

	return ans
}

// @lc code=end

func Test_maximalRectangle(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]byte
		want   int
	}{
		{"x.1", [][]byte{{'0', '1'}}, 1},
		{"0", [][]byte{{'0'}}, 0},
		{"1", [][]byte{{'1'}}, 1},
		{"2", [][]byte{
			{'1', '0', '1', '0', '0'},
			{'1', '0', '1', '1', '1'},
			{'1', '1', '1', '1', '1'},
			{'1', '0', '0', '1', '0'},
		}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximalRectangle(tt.matrix); got != tt.want {
				t.Errorf("maximalRectangle() = %v, want %v", got, tt.want)
			}
		})
	}
}
