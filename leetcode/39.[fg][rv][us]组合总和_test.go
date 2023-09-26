/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 *
 * https://leetcode.cn/problems/combination-sum/description/
 *
 * algorithms
 * Medium (72.33%)
 * Likes:    2620
 * Dislikes: 0
 * Total Accepted:    794.2K
 * Total Submissions: 1.1M
 * Testcase Example:  '[2,3,6,7]\n7'
 *
 * 给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，找出 candidates 中可以使数字和为目标数 target
 * 的 所有 不同组合 ，并以列表形式返回。你可以按 任意顺序 返回这些组合。
 *
 * candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。
 *
 * 对于给定的输入，保证和为 target 的不同组合数少于 150 个。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：candidates = [2,3,6,7], target = 7
 * 输出：[[2,2,3],[7]]
 * 解释：
 * 2 和 3 可以形成一组候选，2 + 2 + 3 = 7 。注意 2 可以使用多次。
 * 7 也是一个候选， 7 = 7 。
 * 仅有这两种组合。
 *
 * 示例 2：
 *
 *
 * 输入: candidates = [2,3,5], target = 8
 * 输出: [[2,2,2,2],[2,3,3],[3,5]]
 *
 * 示例 3：
 *
 *
 * 输入: candidates = [2], target = 1
 * 输出: []
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= candidates.length <= 30
 * 2 <= candidates[i] <= 40
 * candidates 的所有元素 互不相同
 * 1 <= target <= 40
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
func combinationSum(candidates []int, target int) [][]int {
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
			if candidates[i]+sum > target { // 小优化，排序之后无需判断后面的元素了
				break
			}
			path = append(path, candidates[i])
			sum += candidates[i]
			backtrack(startIdx + 1)
			path = path[:len(path)-1]
			sum -= candidates[i]
		}
	}
	backtrack(0)
	return res
}

func permuteRepeat(nums []int) [][]int {
	var res [][]int
	track := []int{}

	// 回溯核心函数
	var backtrack func()
	backtrack = func() {
		// 如果选择的路径等于原数列长度，就收录这个排列
		if len(track) == len(nums) {
			res = append(res, append([]int(nil), track...))
			return
		}

		for i := 0; i < len(nums); i++ {
			// 如果路径中已经包含了当前选择，就跳过
			// if contains(track, nums[i]) {
			//     continue
			// }

			// 做选择
			track = append(track, nums[i])
			// 进入下一层回溯树
			backtrack()
			// 取消选择
			track = track[:len(track)-1]
		}
	}

	backtrack()
	return res
}

// @lc code=end

func Test_combinationSum(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"1", args{[]int{2, 3, 6, 7}, 7}, [][]int{{2, 2, 3}, {7}}},
		{"1", args{[]int{2, 3, 5}, 8}, [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}},
		{"1", args{[]int{2, 1}, 0}, [][]int{{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSum_RV(tt.args.candidates, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combinationSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_permuteRepeat(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{"1", []int{1}, [][]int{{1}}},
		{"1", []int{1, 2}, [][]int{{1, 1}, {1, 2}, {2, 1}, {2, 2}}},
		{"1", []int{1, 2, 3}, [][]int{
			{1, 1, 1}, {1, 1, 2}, {1, 1, 3}, {1, 2, 1}, {1, 2, 2}, {1, 2, 3}, {1, 3, 1}, {1, 3, 2}, {1, 3, 3},
			{2, 1, 1}, {2, 1, 2}, {2, 1, 3}, {2, 2, 1}, {2, 2, 2}, {2, 2, 3}, {2, 3, 1}, {2, 3, 2}, {2, 3, 3},
			{3, 1, 1}, {3, 1, 2}, {3, 1, 3}, {3, 2, 1}, {3, 2, 2}, {3, 2, 3}, {3, 3, 1}, {3, 3, 2}, {3, 3, 3},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permuteRepeat(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permuteRepeat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func combinationSum_RV(candidates []int, target int) [][]int {
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
		if sum > target {
			return
		}

		for i := idx; i < len(candidates); i++ {
			if candidates[i]+sum > target {
				break
			}
			path = append(path, candidates[i])
			sum += candidates[i]
			backtrack(i)
			path = path[:len(path)-1]
			sum -= candidates[i]
		}
	}
	backtrack(0)
	return res
}
