/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 *
 * https://leetcode.cn/problems/3sum/description/
 *
 * algorithms
 * Medium (37.04%)
 * Likes:    6118
 * Dislikes: 0
 * Total Accepted:    1.4M
 * Total Submissions: 3.8M
 * Testcase Example:  '[-1,0,1,2,-1,-4]'
 *
 * 给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j !=
 * k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请
 *
 * 你返回所有和为 0 且不重复的三元组。
 *
 * 注意：答案中不可以包含重复的三元组。
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [-1,0,1,2,-1,-4]
 * 输出：[[-1,-1,2],[-1,0,1]]
 * 解释：
 * nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
 * nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
 * nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
 * 不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
 * 注意，输出的顺序和三元组的顺序并不重要。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [0,1,1]
 * 输出：[]
 * 解释：唯一可能的三元组和不为 0 。
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [0,0,0]
 * 输出：[[0,0,0]]
 * 解释：唯一可能的三元组和为 0 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 3 <= nums.length <= 3000
 * -10^5 <= nums[i] <= 10^5
 *
 *
 */
package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

// @lc code=start
func threeSum(nums []int) [][]int {
	return threeSum3(nums)
}

// 双指针通用 nNums 解法
func threeSum3(nums []int) [][]int {
	sort.Ints(nums)
	return nNums(nums, 3, 0, 0)
}

func nNums(nums []int, n, target, start int) [][]int {
	l := len(nums)
	if n > l {
		return [][]int{}
	}
	res := make([][]int, 0)
	if n == 2 {
		lo, hi := start, l-1
		for lo < hi {
			sum := nums[lo] + nums[hi]
			left, right := nums[lo], nums[hi]
			if sum < target {
				for lo < hi && nums[lo] == left {
					lo++
				}
			} else if sum > target {
				for lo < hi && nums[hi] == right {
					hi--
				}
			} else {
				res = append(res, []int{nums[lo], nums[hi]})
				for lo < hi && nums[lo] == left {
					lo++
				}
				for lo < hi && nums[hi] == right {
					hi--
				}
			}
		}
	} else {
		for i := 0; i < l; i++ {
			subsequence := nNums(nums, n-1, target-nums[i], i+1)
			for _, v := range subsequence {
				v = append(v, nums[i])
				res = append(res, v)
			}
			for i < l-1 && nums[i] == nums[i+1] {
				i++ // 循环体还会递增一次
			}
		}
	}
	return res
}

// 官方题解
func threeSum2(nums []int) [][]int {
	sort.Ints(nums)
	if nums[0] > 0 {
		return [][]int{}
	}

	res := make([][]int, 0)
	l := len(nums)
	for i := 0; i < l; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		k := l - 1
		for j := i + 1; j < l; j++ {
			for j > i+1 && j < l && nums[j] == nums[j-1] {
				j++
			}

			for j < k && nums[j]+nums[k] > -nums[i] {
				k--
			}

			if j >= k {
				break
			}
			if nums[j]+nums[k] == -nums[i] {
				res = append(res, []int{nums[i], nums[j], nums[k]})
			}
		}
	}
	return res
}

// 暴力解法
func threeSum1(nums []int) [][]int {
	solutions := make(map[string][]int)
	l := len(nums)

	genK := func(slice []int) string {
		sort.Ints(slice)
		return fmt.Sprintf("%d%d%d", slice[0], slice[1], slice[2])
	}

	for i := 0; i < l-2; i++ {
		for j := i + 1; j < l-1; j++ {
			for k := j + 1; k < l; k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					solutions[genK([]int{nums[i], nums[j], nums[k]})] = []int{nums[i], nums[j], nums[k]}
				}
			}
		}
	}

	res := make([][]int, 0, len(solutions))
	for _, v := range solutions {
		res = append(res, v)
	}
	return res
}

// @lc code=end

func Test_threeSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{"1", []int{-1, 0, 1, 2, -1, -4}, [][]int{{0, 1, -1}, {-1, 2, -1}}},
		{"2", []int{0, 1, 1}, [][]int{}},
		{"2", []int{3, 0, -2, -1, 1, 2}, [][]int{{0, -2, 2}, {0, -1, 1}, {3, -2, -1}}},
		{"2", []int{0, 0, 0}, [][]int{{0, 0, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := threeSum(tt.nums)
			t.Logf("threeSum() = %v, want %v", got, tt.want)
		})
	}
}
