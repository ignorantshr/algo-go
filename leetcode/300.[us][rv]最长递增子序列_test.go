/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长递增子序列
 *
 * https://leetcode.cn/problems/longest-increasing-subsequence/description/
 *
 * algorithms
 * Medium (54.88%)
 * Likes:    3423
 * Dislikes: 0
 * Total Accepted:    807K
 * Total Submissions: 1.5M
 * Testcase Example:  '[10,9,2,5,3,7,101,18]'
 *
 * 给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
 *
 * 子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7]
 * 的子序列。
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [10,9,2,5,3,7,101,18]
 * 输出：4
 * 解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [0,1,0,3,2,3]
 * 输出：4
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [7,7,7,7,7,7,7]
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 2500
 * -10^4 <= nums[i] <= 10^4
 *
 *
 *
 *
 * 进阶：
 *
 *
 * 你能将算法的时间复杂度降低到 O(n log(n)) 吗?
 *
 *
 */
package leetcode

import "testing"

// @lc code=start

func lengthOfLIS(nums []int) int {
	return lengthOfLISBinarySearch(nums)
	// return lengthOfLISDp(nums)
}

func lengthOfLISDp(nums []int) int {
	dp := make([]int, len(nums)) // 前 i 个数字最长递增子序列长度
	for i := range nums {
		dp[i] = 1 // base case
	}

	for i := 0; i < len(dp); i++ { // 状态
		for j := i - 1; j >= 0; j-- { // 选择
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	res := 1
	for _, v := range dp {
		res = max(res, v)
	}
	return res
}

func lengthOfLISBinarySearch(nums []int) int {
	// 基于特定的规则，就能算出结果
	size := len(nums)
	top := make([]int, size)
	piles := 0

	for _, card := range nums {
		left, right := 0, piles
		for left < right {
			mid := left + (right-left)/2
			if top[mid] < card {
				left = mid + 1
			} else if top[mid] > card {
				right = mid
			} else {
				left = mid
				break
			}
		}

		if left == piles {
			piles++
		}
		top[left] = card
	}
	return piles
}

// @lc code=end

func Test_lengthOfLIS(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"1", []int{1}, 1},
		{"2", []int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		{"3", []int{0, 1, 0, 3, 2, 3}, 4},
		{"4", []int{7, 7, 7, 7, 7, 7, 7}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLISDp_RV(tt.nums); got != tt.want {
				t.Errorf("lengthOfLIS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func lengthOfLISDp_RV(nums []int) int {
	dp := make([]int, len(nums)) // dp[i] 代表了以 nums[i] 结尾的最长递增子序列的长度
	dp[0] = 1

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			// 找到所有的比 nums[i] 小的数结尾的 dp,
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	m := 0
	for _, v := range dp {
		m = max(m, v)
	}
	return m
}
