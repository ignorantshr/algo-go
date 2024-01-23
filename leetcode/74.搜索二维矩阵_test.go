/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] 搜索二维矩阵
 *
 * https://leetcode.cn/problems/search-a-2d-matrix/description/
 *
 * algorithms
 * Medium (48.99%)
 * Likes:    889
 * Dislikes: 0
 * Total Accepted:    360.3K
 * Total Submissions: 731.4K
 * Testcase Example:  '[[1,3,5,7],[10,11,16,20],[23,30,34,60]]\n3'
 *
 * 给你一个满足下述两条属性的 m x n 整数矩阵：
 *
 *
 * 每行中的整数从左到右按非严格递增顺序排列。
 * 每行的第一个整数大于前一行的最后一个整数。
 *
 *
 * 给你一个整数 target ，如果 target 在矩阵中，返回 true ；否则，返回 false 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
 * 输出：true
 *
 *
 * 示例 2：
 *
 *
 * 输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
 * 输出：false
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == matrix.length
 * n == matrix[i].length
 * 1 <= m, n <= 100
 * -10^4 <= matrix[i][j], target <= 10^4
 *
 *
 */
package leetcode

import "testing"

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	// 一次二分
	m := len(matrix)
	n := len(matrix[0])

	l, r := 0, m*n-1
	for l <= r {
		mid := (l + r) / 2
		if matrix[mid/n][mid%n] > target { // ⚠️映射是对列操作
			r = mid - 1
		} else if matrix[mid/n][mid%n] < target {
			l = mid + 1
		} else {
			return true
		}
	}
	return false
}

// 两次二分
func searchMatrixDoubleBinary(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])

	up, bottom := 0, m
	for up < bottom {
		mid := (up + bottom) / 2
		if matrix[mid][0] > target {
			bottom = mid
		} else if matrix[mid][0] < target {
			up = mid + 1
		} else {
			return true
		}
	}

	if bottom <= 0 {
		return false
	}
	bottom--

	l, r := 0, n-1
	for l <= r {
		mid := (l + r) / 2
		if matrix[bottom][mid] > target {
			r = mid - 1
		} else if matrix[bottom][mid] < target {
			l = mid + 1
		} else {
			return true
		}
	}
	return false
}

// @lc code=end

func Test_searchMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1.1", args{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 1}, true},
		{"1.2", args{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3}, true},
		{"1.3", args{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 60}, true},
		{"2.1", args{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, -1}, false},
		{"2.2", args{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 61}, false},
		{"2.3", args{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 21}, false},
		{"2.4", args{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchMatrix(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("searchMatrix() = %v, want %v", got, tt.want)
			}
			if got := searchMatrixDoubleBinary(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("searchMatrix1() = %v, want %v", got, tt.want)
			}
		})
	}
}
