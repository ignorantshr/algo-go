/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 *
 * https://leetcode.cn/problems/combination-sum-ii/description/
 *
 * algorithms
 * Medium (59.63%)
 * Likes:    1444
 * Dislikes: 0
 * Total Accepted:    463K
 * Total Submissions: 776.8K
 * Testcase Example:  '[10,1,2,7,6,1,5]\n8'
 *
 * 给定一个候选人编号的集合 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
 *
 * candidates 中的每个数字在每个组合中只能使用 一次 。
 *
 * 注意：解集不能包含重复的组合。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: candidates = [10,1,2,7,6,1,5], target = 8,
 * 输出:
 * [
 * [1,1,6],
 * [1,2,5],
 * [1,7],
 * [2,6]
 * ]
 *
 * 示例 2:
 *
 *
 * 输入: candidates = [2,5,2,1,2], target = 5,
 * 输出:
 * [
 * [1,2,2],
 * [5]
 * ]
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= candidates.length <= 100
 * 1 <= candidates[i] <= 50
 * 1 <= target <= 30
 *
 *
 */
package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

// @lc code=start
func combinationSum2(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	sum := 0
	sort.Ints(candidates)

	var backtrack func(startIdx int)
	backtrack = func(startIdx int) {
		if sum == target {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		if sum > target {
			return
		}

		for i := startIdx; i < len(candidates); i++ {
			if i != startIdx && candidates[i] == candidates[i-1] {
				continue
			}
			path = append(path, candidates[i])
			sum += candidates[i]
			backtrack(i + 1)
			path = path[:len(path)-1]
			sum -= candidates[i]
		}
	}

	backtrack(0)
	return res
}

// @lc code=end

func Test_combinationSum2(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"1", args{[]int{2, 5, 2, 1, 2}, 5}, [][]int{{1, 2, 2}, {5}}},
		{"1", args{[]int{10, 1, 2, 7, 6, 1, 5}, 8}, [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSum2_RV(tt.args.candidates, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combinationSum2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func combinationSum2_RV(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	sum := 0
	sort.Ints(candidates)

	var backtrack func(idx int)
	backtrack = func(idx int) {
		if sum == target {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for i := idx; i < len(candidates); i++ {
			if i > idx && candidates[i] == candidates[i-1] {
				continue
			}
			if candidates[i]+sum > target {
				break
			}
			path = append(path, candidates[i])
			sum += candidates[i]
			backtrack(i + 1)
			path = path[:len(path)-1]
			sum -= candidates[i]
		}
	}
	backtrack(0)
	return res
}
