/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子数组和
 *
 * https://leetcode.cn/problems/maximum-subarray/description/
 *
 * algorithms
 * Medium (54.85%)
 * Likes:    6370
 * Dislikes: 0
 * Total Accepted:    1.5M
 * Total Submissions: 2.8M
 * Testcase Example:  '[-2,1,-3,4,-1,2,1,-5,4]'
 *
 * 给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
 *
 * 子数组 是数组中的一个连续部分。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
 * 输出：6
 * 解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1]
 * 输出：1
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [5,4,-1,7,8]
 * 输出：23
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 10^5
 * -10^4 <= nums[i] <= 10^4
 *
 *
 *
 *
 * 进阶：如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的 分治法 求解。
 *
 */
package leetcode

import "testing"

// @lc code=start
func maxSubArray(nums []int) int {
	return maxSubArrayBinary(nums)
	return maxSubArrayPrefixSumArray(nums)
	return maxSubArrayDpTableCompactSpace(nums)
	return maxSubArrayDpTable(nums)
	return maxSubArrayDpDfs(nums)
	return maxSubArrayWindow(nums)
}

type status53 struct {
	lSum int // [l,r] 内以 l 为左端点的最大子段和
	rSum int // [l,r] 内以 r 为左端点的最大子段和
	mSum int // [l,r] 内的最大子段和
	iSum int // [l,r] 区间和
}

// 分治法，线段树
func maxSubArrayBinary(nums []int) int {
	var getStatus func(l, r int) status53
	var pushUp func(l, r status53) status53

	getStatus = func(l, r int) status53 {
		if l == r {
			return status53{nums[l], nums[l], nums[l], nums[l]}
		}

		mid := (l + r) >> 1
		leftSub := getStatus(l, mid)
		rightSub := getStatus(mid+1, r)
		return pushUp(leftSub, rightSub)
	}

	pushUp = func(l, r status53) status53 {
		s := status53{}
		s.iSum = l.iSum + r.iSum
		s.lSum = max(l.lSum, l.iSum+r.lSum)
		s.rSum = max(r.rSum, r.iSum+l.rSum)
		// 1. 在左区间 2. 在右区间 3. 跨区间
		s.mSum = max(max(l.mSum, r.mSum), l.rSum+r.lSum)
		return s
	}

	return getStatus(0, len(nums)-1).mSum
}

// 前缀和数组
func maxSubArrayPrefixSumArray(nums []int) int {
	pres := make([]int, len(nums)+1)
	pres[0] = 0

	for i := 1; i <= len(nums); i++ {
		pres[i] = pres[i-1] + nums[i-1]
	}

	minv := pres[0]
	maxv := pres[0]
	for i := 0; i < len(nums); i++ {
		minv = min(minv, pres[i])        // 维护 pres[0:i] 的最小值
		maxv = max(maxv, pres[i+1]-minv) // 以 nums[i] 结尾的最大子数组和就是 preSum[i+1] - min(preSum[0..i])
	}

	return maxv
}

// 动态规划，dp 数组，压缩空间
func maxSubArrayDpTableCompactSpace(nums []int) int {
	next := nums[0] // dp[i] 是以 nums[i] 结尾的子数组的最大和
	maxv := nums[0]

	for i := 1; i < len(nums); i++ {
		next = max(nums[i], nums[i]+next)
		maxv = max(maxv, next)
	}

	return maxv
}

// 动态规划，dp 数组
func maxSubArrayDpTable(nums []int) int {
	dp := make([]int, len(nums)) // dp[i] 是以 nums[i] 结尾的子数组的最大和
	dp[0] = nums[0]
	maxv := nums[0]

	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		maxv = max(maxv, dp[i])
	}

	return maxv
}

// 动态规划，递归
func maxSubArrayDpDfs(nums []int) int {
	maxv := nums[0]

	var dp func(i int) int
	dp = func(i int) int {
		if i == len(nums) {
			return 0
		}

		child := dp(i + 1)
		result := max(nums[i]+child, nums[i])
		if maxv < result {
			maxv = result
		}
		return result
	}

	dp(0)
	return maxv
}

// 滑动窗口
func maxSubArrayWindow(nums []int) int {
	maxv := nums[0]
	sum := maxv
	left, right := 0, 1

	for right < len(nums) {
		sum += nums[right]
		right++

		if maxv < sum {
			maxv = sum
		}

		// nums 中全是负数，可得正解
		// nums 中有正有负，这种情况下元素和最大的那个子数组一定是以正数开头的，
		// 	此时我们需要穷举所有以正数开头的子数组，计算他们的元素和，找到元素和最大的那个子数组。
		// 	相当于清空滑动窗口，从新的位置继续寻找
		for sum < 0 {
			sum -= nums[left]
			left++
		}
	}
	return maxv
}

// @lc code=end

func Test_maxSubArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"1", []int{5, 4, -1, 7, 8}, 23},
		{"1", []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{"1", []int{-5, -4, -1, -7, -8}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray(tt.nums); got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
