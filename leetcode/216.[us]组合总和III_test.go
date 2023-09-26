/*
- @lc app=leetcode.cn id=216 lang=golang

找出所有相加之和为 n 的 k 个数的组合。组合中只允许含有 1 - 9 的正整数，并且每种组合中不存在重复的数字。

说明：

所有数字都是正整数。
解集不能包含重复的组合。
示例 1: 输入: k = 3, n = 7 输出: [[1,2,4]]

示例 2: 输入: k = 3, n = 9 输出: [[1,2,6], [1,3,5], [2,3,4]]
*/
package leetcode

import (
	"testing"
)

// @lc code=start
func combinationSum3(k, n int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	sum := 0

	var backtrack func(idx int) bool
	backtrack = func(idx int) bool {
		if len(path) == k {
			if sum == n {
				tmp := make([]int, len(path))
				copy(tmp, path)
				res = append(res, tmp)
				return true
			}
			return false
		}

		for i := idx; i <= 9-(k-len(path))+1; i++ {
			path = append(path, i)
			sum += i
			found := backtrack(i + 1)
			path = path[:len(path)-1]
			sum -= i
			if found {
				break
			}
		}
		return false
	}
	backtrack(1)
	return res
}

// @lc code=end

func Test_combinationSum3(t *testing.T) {
	type args struct {
		k int
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"0", args{0, 1}, [][]int{}},
		{"0", args{1, 0}, [][]int{}},
		{"1", args{1, 1}, [][]int{{1}}},
		{"1", args{3, 7}, [][]int{{1, 2, 4}}},
		{"1", args{3, 9}, [][]int{{1, 2, 6}, {1, 3, 5}, {2, 3, 4}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSum3(tt.args.k, tt.args.n); !equalSliceMatrix(got, tt.want) {
				t.Errorf("combinationSum3() = %v, want %v", got, tt.want)
			}
		})
	}
}
