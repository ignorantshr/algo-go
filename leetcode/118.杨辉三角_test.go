/*
 * @lc app=leetcode.cn id=118 lang=golang
 *
 * [118] 杨辉三角
 *
 * https://leetcode.cn/problems/pascals-triangle/description/
 *
 * algorithms
 * Easy (75.50%)
 * Likes:    1120
 * Dislikes: 0
 * Total Accepted:    479.2K
 * Total Submissions: 633.5K
 * Testcase Example:  '5'
 *
 * 给定一个非负整数 numRows，生成「杨辉三角」的前 numRows 行。
 *
 * 在「杨辉三角」中，每个数是它左上方和右上方的数的和。
 *
 *
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: numRows = 5
 * 输出: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
 *
 *
 * 示例 2:
 *
 *
 * 输入: numRows = 1
 * 输出: [[1]]
 *
 */
package leetcode

import (
	"reflect"
	"testing"
)

// @lc code=start
func generate(numRows int) [][]int {
	res := make([][]int, 0)

	for i := 0; i < numRows; i++ {
		res = append(res, make([]int, i+1))
		res[i][0] = 1
		for j := 1; j < i; j++ {
			res[i][j] = res[i-1][j-1] + res[i-1][j]
		}
		res[i][i] = 1
	}

	return res
}

// @lc code=end

func Test_generate(t *testing.T) {
	tests := []struct {
		name    string
		numRows int
		want    [][]int
	}{
		{"1", 1, [][]int{{1}}},
		{"2", 2, [][]int{{1}, {1, 1}}},
		{"3", 3, [][]int{{1}, {1, 1}, {1, 2, 1}}},
		{"4", 4, [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}}},
		{"5", 5, [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generate(tt.numRows); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generate() = %v, want %v", got, tt.want)
			}
		})
	}
}
