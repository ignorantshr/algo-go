/*
 * @lc app=leetcode.cn id=309 lang=golang
 *
 * [309] 买卖股票的最佳时机含冷冻期
 *
 * https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-cooldown/description/
 *
 * algorithms
 * Medium (64.19%)
 * Likes:    1651
 * Dislikes: 0
 * Total Accepted:    291.7K
 * Total Submissions: 450.8K
 * Testcase Example:  '[1,2,3,0,2]'
 *
 * 给定一个整数数组prices，其中第  prices[i] 表示第 i 天的股票价格 。​
 *
 * 设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:
 *
 *
 * 卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。
 *
 *
 * 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: prices = [1,2,3,0,2]
 * 输出: 3
 * 解释: 对应的交易状态为: [买入, 卖出, 冷冻期, 买入, 卖出]
 *
 * 示例 2:
 *
 *
 * 输入: prices = [1]
 * 输出: 0
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= prices.length <= 5000
 * 0 <= prices[i] <= 1000
 *
 *
 */
package leetcode

import (
	"math"
	"testing"
)

// @lc code=start
func maxProfit309(prices []int) int {
	return maxProfit309Optimum(prices)
	return maxProfit309Base(prices)
}

func maxProfit309Optimum(prices []int) int {
	size := len(prices)
	dp_i_0 := 0
	dp_i_1 := math.MinInt
	dp_pre_2 := 0

	for i := 0; i < size; i++ {
		// dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		// if i-2 < 0 {
		// 	dp[i][1] = max(dp[i-1][1], -prices[i])
		// } else {
		// 	dp[i][1] = max(dp[i-1][1], dp[i-2][0]-prices[i]) //前天卖出后的状态
		// }

		tmp := dp_i_0
		dp_i_0 = max(dp_i_0, dp_i_1+prices[i])
		dp_i_1 = max(dp_i_1, dp_pre_2-prices[i])
		dp_pre_2 = tmp
	}

	return dp_i_0
}

func maxProfit309Base(prices []int) int {
	size := len(prices)
	dp := make([][2]int, size)

	for i := 0; i < size; i++ {
		if i-1 < 0 {
			dp[i][0] = 0
			dp[i][1] = -prices[i]
			continue
		}
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		if i-2 < 0 {
			dp[i][1] = max(dp[i-1][1], -prices[i])
		} else {
			dp[i][1] = max(dp[i-1][1], dp[i-2][0]-prices[i]) //前天卖出后的状态
		}
	}

	return dp[size-1][0]
}

// @lc code=end

func Test_maxProfit309(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{"0", []int{1}, 0},
		{"0", []int{2, 1}, 0},
		{"0", []int{5, 4, 2, 1}, 0},
		{"1", []int{1, 2}, 1},
		{"1", []int{1, 2, 3, 0, 2}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit309(tt.prices); got != tt.want {
				t.Errorf("maxProfit309() = %v, want %v", got, tt.want)
			}
		})
	}
}
