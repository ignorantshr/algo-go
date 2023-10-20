/*
 * @lc app=leetcode.cn id=518 lang=golang
 *
 * [518] 零钱兑换 II
 *
 * https://leetcode.cn/problems/coin-change-ii/description/
 *
 * algorithms
 * Medium (70.63%)
 * Likes:    1178
 * Dislikes: 0
 * Total Accepted:    269.1K
 * Total Submissions: 380K
 * Testcase Example:  '5\n[1,2,5]'
 *
 * 给你一个整数数组 coins 表示不同面额的硬币，另给一个整数 amount 表示总金额。
 *
 * 请你计算并返回可以凑成总金额的硬币组合数。如果任何硬币组合都无法凑出总金额，返回 0 。
 *
 * 假设每一种面额的硬币有无限个。
 *
 * 题目数据保证结果符合 32 位带符号整数。
 *
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：amount = 5, coins = [1, 2, 5]
 * 输出：4
 * 解释：有四种方式可以凑成总金额：
 * 5=5
 * 5=2+2+1
 * 5=2+1+1+1
 * 5=1+1+1+1+1
 *
 *
 * 示例 2：
 *
 *
 * 输入：amount = 3, coins = [2]
 * 输出：0
 * 解释：只用面额 2 的硬币不能凑成总金额 3 。
 *
 *
 * 示例 3：
 *
 *
 * 输入：amount = 10, coins = [10]
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 1
 * coins 中的所有值 互不相同
 * 0
 *
 *
 */
package leetcode

import (
	"testing"
)

// @lc code=start
func change(amount int, coins []int) int {
	return changeDpOptimum(amount, coins)
	return changeDp(amount, coins)
	return changeDfs(amount, coins)
}

func changeDpOptimum(amount int, coins []int) int {
	size := len(coins)
	dp := make([]int, amount+1)

	for i := 1; i <= size; i++ {
		for j := 1; j <= amount; j++ {
			if j-coins[i-1] >= 0 {
				dp[j] = dp[j] + dp[j-coins[i-1]]
			}
		}
	}

	return dp[amount]
}

func changeDp(amount int, coins []int) int {
	// 状态就是选中的 coins 的和
	// 选择就是 coins
	size := len(coins)
	dp := make([][]int, size+1) // amount 为 j 时，前 i 个硬币的组合数
	for i := 0; i < size+1; i++ {
		dp[i] = make([]int, amount+1)
		dp[i][0] = 1
	}
	// dp[0][j] = 0

	for i := 1; i <= size; i++ {
		for j := 1; j <= amount; j++ {
			// sum := 0
			// // 给本次使用的金币留出 m 倍的空间来填充
			// for m := 0; m*coins[i-1] <= j; m++ {
			// 	sum += dp[i-1][j-m*coins[i-1]]
			// }
			// dp[i][j] = sum

			if j-coins[i-1] >= 0 {
				dp[i][j] = dp[i-1][j] + // 不用这枚金币
					dp[i][j-coins[i-1]] // 使用这枚金币，注意 [i]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[size][amount]
}

// 回溯
func changeDfs(amount int, coins []int) int {
	res := 0
	var backtrace func(idx, sum int)
	backtrace = func(idx, sum int) {
		if sum == amount {
			res++
			return
		}
		if sum > amount {
			return
		}

		for i := idx; i < len(coins); i++ {
			backtrace(i, sum+coins[i]) // 选择
		}
	}
	backtrace(0, 0)
	return res
}

// @lc code=end

func Test_change(t *testing.T) {
	type args struct {
		amount int
		coins  []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{2, []int{1, 2}}, 2},
		{"", args{4, []int{5}}, 0},
		{"", args{5, []int{5}}, 1},
		{"", args{5, []int{1, 2, 5}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := change(tt.args.amount, tt.args.coins); got != tt.want {
				t.Errorf("change() = %v, want %v", got, tt.want)
			}
		})
	}
}
