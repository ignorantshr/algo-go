/*
 * @lc app=leetcode.cn id=494 lang=golang
 *
 * [494] 目标和
 *
 * https://leetcode.cn/problems/target-sum/description/
 *
 * algorithms
 * Medium (48.58%)
 * Likes:    1772
 * Dislikes: 0
 * Total Accepted:    388.9K
 * Total Submissions: 803K
 * Testcase Example:  '[1,1,1,1,1]\n3'
 *
 * 给你一个非负整数数组 nums 和一个整数 target 。
 *
 * 向数组中的每个整数前添加 '+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：
 *
 *
 * 例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表达式 "+2-1" 。
 *
 *
 * 返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,1,1,1,1], target = 3
 * 输出：5
 * 解释：一共有 5 种方法让最终目标和为 3 。
 * -1 + 1 + 1 + 1 + 1 = 3
 * +1 - 1 + 1 + 1 + 1 = 3
 * +1 + 1 - 1 + 1 + 1 = 3
 * +1 + 1 + 1 - 1 + 1 = 3
 * +1 + 1 + 1 + 1 - 1 = 3
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1], target = 1
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 20
 * 0 <= nums[i] <= 1000
 * 0 <= sum(nums[i]) <= 1000
 * -1000 <= target <= 1000
 *
 *
 */
package leetcode

import "testing"

// 题解可参考：https://leetcode.cn/problems/target-sum/solutions/2285605/yi-wen-jiang-tou-you-yi-dao-nan-yi-bu-bu-fjsz/

// @lc code=start
func findTargetSumWays(nums []int, target int) int {
	sum := target
	for _, v := range nums {
		sum += v
	}
	if (sum & 1) == 1 {
		return 0
	}

	return findTargetSumWaysDp(nums, sum)
	return findTargetSumWaysDfs2(nums, sum)
	return findTargetSumWaysDfs1(nums, target)
}

func findTargetSumWaysDp(nums []int, sum int) int {
	sum = sum / 2
	size := len(nums)
	dp := make([][]int, size+1) // 前 i 个数的 j 组合数
	for i := 0; i < size+1; i++ {
		dp[i] = make([]int, sum+1)
	}
	dp[0][0] = 1

	for i := 1; i <= size; i++ {
		for j := 0; j <= sum; j++ {
			dp[i][j] = dp[i-1][j] // 不选
			if j-nums[i-1] >= 0 {
				dp[i][j] += dp[i-1][j-nums[i-1]] // 选
			}
		}
	}
	return dp[size][sum]
}

/*
选/不选思想
@Return 返回值表示从当前层开始考虑，目标和为target的数字组合的个数
dfs(n,target) = dfs(n-1,target) + dfs(n-1,target - nums[n])
递推公式理解：表示当前解由不选当前数字和选当前数字的解的数目构成
*/
func findTargetSumWaysDfs2(nums []int, sum int) int {
	sum = sum / 2
	var dfs func(idx, res int) int
	dfs = func(idx, res int) int {
		if idx == len(nums) {
			if res == sum {
				return 1
			}
			return 0
		}

		return dfs(idx+1, res+nums[idx]) + // choose
			dfs(idx+1, res) // not choose
	}

	return dfs(0, 0)
}

func findTargetSumWaysDfs1(nums []int, target int) int {
	size := len(nums)
	var dfs func(idx, res int) int
	dfs = func(idx, res int) int {
		if idx == size {
			if res == target {
				return 1
			}
			return 0
		}

		return dfs(idx+1, res+nums[idx]) + // positive
			dfs(idx+1, res-nums[idx]) // negative
	}

	return dfs(0, 0)
}

// @lc code=end

func Test_findTargetSumWays(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{[]int{1}, 1}, 1},
		{"", args{[]int{1, 1}, 0}, 2},
		{"", args{[]int{1, 1, 1}, 1}, 3},
		{"", args{[]int{1, 1, 1, 1, 1}, 3}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTargetSumWays(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("findTargetSumWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
