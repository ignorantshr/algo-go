/*
 * @lc app=leetcode.cn id=416 lang=golang
 *
 * [416] 分割等和子集
 *
 * https://leetcode.cn/problems/partition-equal-subset-sum/description/
 *
 * algorithms
 * Medium (52.18%)
 * Likes:    1904
 * Dislikes: 0
 * Total Accepted:    464.9K
 * Total Submissions: 890.2K
 * Testcase Example:  '[1,5,11,5]'
 *
 * 给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,5,11,5]
 * 输出：true
 * 解释：数组可以分割成 [1, 5, 5] 和 [11] 。
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1,2,3,5]
 * 输出：false
 * 解释：数组不能分割成两个元素和相等的子集。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 1
 *
 *
 */
package leetcode

import "testing"

// @lc code=start
func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum&1 == 1 {
		return false
	}

	mid := sum / 2
	return canPartitionDpOptimum(mid, nums)
	return canPartitionDp(mid, nums)
	return canPartitionDfs(mid, nums)
}

func canPartitionDpOptimum(mid int, nums []int) bool {
	size := len(nums)
	dp := make([]bool, mid+1)
	dp[0] = true

	for i := 0; i < size; i++ {
		for sum := mid; sum >= 0; sum-- { // 反向
			if sum-nums[i] >= 0 {
				// 不加上这个数字  or 加上这个数字
				dp[sum] = dp[sum] || dp[sum-nums[i]]
			}
		}
	}

	return dp[mid]
}

func canPartitionDp(mid int, nums []int) bool {
	size := len(nums)
	dp := make([][]bool, size+1) // dp[i][j] 是否前 i 个数有一种数字组合，它们的和为 j
	for i := 0; i < size+1; i++ {
		dp[i] = make([]bool, mid+1)
		dp[i][0] = true
	}
	// dp[0][j] = false

	for i := 1; i <= size; i++ {
		for sum := 1; sum <= mid; sum++ {
			if sum-nums[i-1] < 0 {
				dp[i][sum] = dp[i-1][sum]
			} else {
				// 不加上这个数字  or 加上这个数字
				dp[i][sum] = dp[i-1][sum] || dp[i-1][sum-nums[i-1]]
			}
		}
	}

	return dp[size][mid]
}

func canPartitionDfs(mid int, nums []int) bool {
	var dfs func(i, sum int) bool
	dfs = func(i, sum int) bool {
		if sum == 0 {
			return true
		}
		if sum < 0 || i < 0 {
			return false
		}

		return dfs(i-1, sum) || dfs(i-1, sum-nums[i])
	}

	return dfs(len(nums), mid)
}

// @lc code=end

func Test_canPartition(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{"", []int{1, 5, 11, 5}, true},
		{"", []int{1, 2, 3, 5}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPartition(tt.nums); got != tt.want {
				t.Errorf("canPartition() = %v, want %v", got, tt.want)
			}
		})
	}
}
