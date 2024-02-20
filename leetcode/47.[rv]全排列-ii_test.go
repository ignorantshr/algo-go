/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 *
 * https://leetcode.cn/problems/permutations-ii/description/
 *
 * algorithms
 * Medium (65.57%)
 * Likes:    1458
 * Dislikes: 0
 * Total Accepted:    489.4K
 * Total Submissions: 746.2K
 * Testcase Example:  '[1,1,2]'
 *
 * 给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,1,2]
 * 输出：
 * [[1,1,2],⁠[1,2,1],⁠[2,1,1]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1,2,3]
 * 输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 8
 * -10 <= nums[i] <= 10
 *
 *
 */
package leetcode

import (
	"sort"
	"testing"
)

// @lc code=start
func permuteUnique(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	used := make([]bool, len(nums))
	sort.Ints(nums)

	var backtrack func(idx int)
	backtrack = func(idx int) {
		if len(path) == len(nums) {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			// 只有前面相同元素用上了才会使用本元素，这样保证了相同元素之间的相对位置
			//
			// 当出现重复元素时，比如输入 nums = [1,2,2',2'']，
			// 2' 只有在 2 已经被使用的情况下才会被选择，
			// 同理，2'' 只有在 2' 已经被使用的情况下才会被选择，
			// 这就保证了相同元素在排列中的相对位置保证固定。
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}

			path = append(path, nums[i])
			used[i] = true
			backtrack(idx + 1)
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	backtrack(0)
	return res
}

func permuteUnique1(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	used := make([]bool, len(nums))
	sort.Ints(nums)

	var backtrack func(startIdx int)
	backtrack = func(startIdx int) {
		if len(path) == len(nums) {
			dest := make([]int, len(path))
			copy(dest, path)
			res = append(res, dest)
			return
		}

		preVal := 999
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			if nums[i] == preVal {
				continue
			}
			path = append(path, nums[i])
			used[i] = true
			preVal = nums[i]
			backtrack(startIdx + 1)
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	backtrack(0)
	return res
}

// @lc code=end

func Test_permuteUnique(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{"x", []int{0, 1, 0, 0, 9}, [][]int{{0, 0, 0, 1, 9}, {0, 0, 0, 9, 1}, {0, 0, 1, 0, 9}, {0, 0, 1, 9, 0}, {0, 0, 9, 0, 1}, {0, 0, 9, 1, 0}, {0, 1, 0, 0, 9}, {0, 1, 0, 9, 0}, {0, 1, 9, 0, 0}, {0, 9, 0, 0, 1}, {0, 9, 0, 1, 0}, {0, 9, 1, 0, 0}, {1, 0, 0, 0, 9}, {1, 0, 0, 9, 0}, {1, 0, 9, 0, 0}, {1, 9, 0, 0, 0}, {9, 0, 0, 0, 1}, {9, 0, 0, 1, 0}, {9, 0, 1, 0, 0}, {9, 1, 0, 0, 0}}},
		{"x.2", []int{2, 2, 1, 1}, [][]int{{1, 1, 2, 2}, {1, 2, 1, 2}, {1, 2, 2, 1}, {2, 1, 1, 2}, {2, 1, 2, 1}, {2, 2, 1, 1}}},
		{"2", []int{1, 1, 1}, [][]int{{1, 1, 1}}},
		{"2.1", []int{1, 2, 1}, [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}}},
		{"0", []int{}, [][]int{{}}},
		{"1", []int{1}, [][]int{{1}}},
		{"1", []int{0, 1}, [][]int{{0, 1}, {1, 0}}},
		{"1", []int{1, 2, 3}, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 2, 1}, {3, 1, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permuteUnique(tt.nums); !equalSetMatrix(got, tt.want) {
				t.Errorf("permuteUnique() = %v, want %v", got, tt.want)
			}
			if got := permuteUnique1(tt.nums); !equalSetMatrix(got, tt.want) {
				t.Errorf("permuteUnique1() = %v, want %v", got, tt.want)
			}
			if got := permuteUnique_RV(tt.nums); !equalSetMatrix(got, tt.want) {
				t.Errorf("permuteUnique_RV() = %v, want %v", got, tt.want)
			}
		})
	}
}

func permuteUnique_RV(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	used := make([]bool, len(nums))
	sort.Ints(nums)

	var backtrack func(idx int)
	backtrack = func(idx int) {
		if len(path) == len(nums) {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			// 只有前面相同元素用上了才会使用本元素，这样保证了相同元素之间的相对位置
			//
			// 当出现重复元素时，比如输入 nums = [1,2,2',2'']，
			// 2' 只有在 2 已经被使用的情况下才会被选择，
			// 同理，2'' 只有在 2' 已经被使用的情况下才会被选择，
			// 这就保证了相同元素在排列中的相对位置保证固定。
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}

			path = append(path, nums[i])
			used[i] = true
			backtrack(idx + 1)
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	backtrack(0)
	return res
}
