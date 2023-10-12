/*
 * @lc app=leetcode.cn id=322 lang=c
 *
 * [322] 零钱兑换
 *
 * https://leetcode.cn/problems/coin-change/description/
 *
 * algorithms
 * Medium (46.53%)
 * Likes:    2597
 * Dislikes: 0
 * Total Accepted:    691.2K
 * Total Submissions: 1.5M
 * Testcase Example:  '[1,2,5]\n11'
 *
 * 给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
 *
 * 计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
 *
 * 你可以认为每种硬币的数量是无限的。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：coins = [1, 2, 5], amount = 11
 * 输出：3
 * 解释：11 = 5 + 5 + 1
 *
 * 示例 2：
 *
 *
 * 输入：coins = [2], amount = 3
 * 输出：-1
 *
 * 示例 3：
 *
 *
 * 输入：coins = [1], amount = 0
 * 输出：0
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= coins.length <= 12
 * 1 <= coins[i] <= 2^31 - 1
 * 0 <= amount <= 10^4
 *
 *
 */
package leetcode

import (
	"math"
	"testing"
)

// @lc code=start

/*
1、确定 base case。这个很简单，显然目标金额 amount 为 0 时算法返回 0，因为不需要任何硬币就已经凑出目标金额了。

2、确定「状态」，也就是原问题和子问题中会变化的变量。由于硬币数量无限，硬币的面额也是题目给定的，只有目标金额会不断地向 base case 靠近，所以唯一的「状态」就是目标金额 amount。

3、确定「选择」，也就是导致「状态」产生变化的行为。目标金额为什么变化呢，因为你在选择硬币，你每选择一枚硬币，就相当于减少了目标金额。所以说所有硬币的面值，就是你的「选择」。

4、明确 dp 函数/数组的定义。我们这里讲的是自顶向下的解法，所以会有一个递归的 dp 函数，一般来说函数的参数就是状态转移中会变化的量，也就是上面说到的「状态」；函数的返回值就是题目要求我们计算的量。
就本题来说，状态只有一个，即「目标金额」，题目要求我们计算凑出目标金额所需的最少硬币数量。

所以我们可以这样定义 dp 函数：dp(n) 表示，输入一个目标金额 n，返回凑出目标金额 n 所需的最少硬币数量。
*/

func coinChange(coins []int, amount int) int {
	return coinChange3(coins, amount)
}

// 自底向上
func coinChange3(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1 // 初始化为最大值，方便后续比较最小值
	}
	dp[0] = 0 // base case

	for i := 1; i <= amount; i++ { // 「状态」
		for _, v := range coins { // 「选择」
			if i-v < 0 {
				continue
			}
			dp[i] = min(dp[i], dp[i-v]+1)
		}
	}

	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

// 自顶向下，备忘录优化剪枝
func coinChange2(coins []int, amount int) int {
	memo := make(map[int]int)

	var dp func(coins []int, amount int) int
	dp = func(coins []int, amount int) int {
		if v, has := memo[amount]; has {
			return v
		}

		res := math.MaxInt
		if amount == 0 {
			return 0
		}
		if amount < 0 {
			return -1
		}

		for _, coin := range coins {
			subproblem := dp(coins, amount-coin)
			if subproblem != -1 {
				res = min(res, subproblem+1)
			}
		}
		if res == math.MaxInt {
			memo[amount] = -1
		} else {
			memo[amount] = res
		}

		return memo[amount]
	}

	return dp(coins, amount)
}

// 自顶向下
func coinChange1(coins []int, amount int) int {
	var dp func(coins []int, amount int) int
	dp = func(coins []int, amount int) int {
		res := math.MaxInt
		if amount == 0 {
			return 0
		}
		if amount < 0 {
			return -1
		}

		for _, coin := range coins {
			subproblem := dp(coins, amount-coin)
			if subproblem != -1 {
				res = min(res, subproblem+1)
			}
		}
		if res == math.MaxInt {
			return -1
		}

		return res
	}
	return dp(coins, amount)
}

// @lc code=end

func Test_coinChange(t *testing.T) {
	type args struct {
		coins  []int
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0", args{[]int{}, 0}, 0},
		{"0", args{[]int{1}, 0}, 0},
		{"0", args{[]int{1, 2, 5}, 11}, 3},
		{"0", args{[]int{1, 2}, 1}, 1},
		{"0", args{[]int{2}, 1}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coinChange(tt.args.coins, tt.args.amount); got != tt.want {
				t.Errorf("coinChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
