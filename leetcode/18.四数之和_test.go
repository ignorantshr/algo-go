/*
 * @lc app=leetcode.cn id=18 lang=golang
 *
 * [18] 四数之和
 *
 * https://leetcode.cn/problems/4sum/description/
 *
 * algorithms
 * Medium (36.78%)
 * Likes:    1688
 * Dislikes: 0
 * Total Accepted:    487K
 * Total Submissions: 1.3M
 * Testcase Example:  '[1,0,-1,0,-2,2]\n0'
 *
 * 给你一个由 n 个整数组成的数组 nums ，和一个目标值 target 。请你找出并返回满足下述全部条件且不重复的四元组 [nums[a],
 * nums[b], nums[c], nums[d]] （若两个四元组元素一一对应，则认为两个四元组重复）：
 *
 *
 * 0 <= a, b, c, d < n
 * a、b、c 和 d 互不相同
 * nums[a] + nums[b] + nums[c] + nums[d] == target
 *
 *
 * 你可以按 任意顺序 返回答案 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,0,-1,0,-2,2], target = 0
 * 输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [2,2,2,2,2], target = 8
 * 输出：[[2,2,2,2]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 200
 * -10^9 <= nums[i] <= 10^9
 * -10^9 <= target <= 10^9
 *
 *
 */
package leetcode

import (
	"sort"
	"testing"
)

// @lc code=start
func fourSum(nums []int, target int) [][]int {
	return fourSumCommon2(nums, target)
}

// n数和 通用解法
func fourSumCommon2(nums []int, target int) [][]int {
	sort.Ints(nums)
	return nSumCommon(nums, 4, target, 0, len(nums)-1)
}

func nSumCommon(nums []int, n, target, start, end int) [][]int {
	if n < 2 {
		return nil
	}

	var res [][]int
	if n == 2 {
		first := true
		for ; start < end; start++ {
			for !first && start > 0 && start < end && nums[start] == nums[start-1] {
				start++
			}

			for start < end && nums[start]+nums[end] > target {
				end--
			}

			if start >= end {
				break
			}

			if nums[start]+nums[end] == target {
				res = append(res, []int{nums[start], nums[end]})
			}
			first = false
		}
	} else {
		for i := start; i <= end; i++ {
			subsequence := nSumCommon(nums, n-1, target-nums[i], i+1, end)
			for _, list := range subsequence {
				list = append(list, nums[i])
				res = append(res, list)
			}
			for i < end && nums[i] == nums[i+1] {
				i++
			}
		}
	}

	return res
}

func fourSum1(nums []int, target int) [][]int {
	sort.Ints(nums)

	var res [][]int
	size := len(nums)

	for i := 0; i < size; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < size; j++ {
			if j > i+1 && j < size && nums[j] == nums[j-1] {
				continue
			}

			for l, r := j+1, size-1; l < r; l++ {
				for l > j+1 && l < r && nums[l] == nums[l-1] {
					l++
				}

				for l < r && nums[i]+nums[j]+nums[l]+nums[r] > target {
					r--
				}
				if l >= r {
					break
				}

				if nums[i]+nums[j]+nums[l]+nums[r] == target {
					res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})
				}
			}
		}
	}

	return res
}

// @lc code=end

func Test_fourSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"1", args{
			[]int{1, 0, -1, 0, -2, 2},
			0,
		},
			[][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}},
		},
		{"4", args{
			[]int{1, -2, -5, -4, -3, 3, 3, 5},
			-11,
		},
			[][]int{{-5, -4, -3, 1}},
		},
		{"2", args{
			[]int{-2, -1, -1, 1, 1, 2, 2},
			0,
		},
			[][]int{{-2, -1, 1, 2}, {-1, -1, 1, 1}},
		},
		{"3", args{
			[]int{2, 2, 2, 2},
			8,
		},
			[][]int{{2, 2, 2, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fourSum(tt.args.nums, tt.args.target)
			t.Logf("fourSum() = %v, want %v", got, tt.want)
		})
	}
}
