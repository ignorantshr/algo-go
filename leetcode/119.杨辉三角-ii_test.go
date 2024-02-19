/*
 * @lc app=leetcode.cn id=119 lang=golang
 *
 * [119] 杨辉三角 II
 *
 * https://leetcode.cn/problems/pascals-triangle-ii/description/
 *
 * algorithms
 * Easy (68.93%)
 * Likes:    529
 * Dislikes: 0
 * Total Accepted:    298.1K
 * Total Submissions: 432.3K
 * Testcase Example:  '3'
 *
 * 给定一个非负索引 rowIndex，返回「杨辉三角」的第 rowIndex 行。
 *
 * 在「杨辉三角」中，每个数是它左上方和右上方的数的和。
 *
 * 示例 1:
 *
 *
 * 输入: rowIndex = 3
 * 输出: [1,3,3,1]
 *
 *
 * 示例 2:
 *
 *
 * 输入: rowIndex = 0
 * 输出: [1]
 *
 *
 * 示例 3:
 *
 *
 * 输入: rowIndex = 1
 * 输出: [1,1]
 *
 * 进阶：
 *
 * 你可以优化你的算法到 O(rowIndex) 空间复杂度吗？
 *
 */
package leetcode

import (
	"reflect"
	"testing"
)

// @lc code=start
func getRow(rowIndex int) []int {
	// 倒着递推，就不用担心覆盖的问题了
	res := make([]int, rowIndex+1)

	res[0] = 1
	for i := 1; i <= rowIndex; i++ {
		for j := i; j > 0; j-- {
			res[j] = res[j-1] + res[j]
		}
	}

	return res
}

// 递推
func getRowDp(rowIndex int) []int {
	res := make([]int, rowIndex+1)

	for i := 0; i <= rowIndex; i++ {
		res[0] = 1
		res[i] = 1
		add := 1
		for j := 1; j < i; j++ {
			tmp := res[j]
			res[j] = add + res[j]
			add = tmp
		}
	}

	return res
}

// @lc code=end

func Test_getRow(t *testing.T) {
	tests := []struct {
		name     string
		rowIndex int
		want     []int
	}{
		{"0", 0, []int{1}},
		{"1", 1, []int{1, 1}},
		{"2", 2, []int{1, 2, 1}},
		{"3", 3, []int{1, 3, 3, 1}},
		{"4", 4, []int{1, 4, 6, 4, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getRow(tt.rowIndex); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getRow() = %v, want %v", got, tt.want)
			}
		})
	}
}
