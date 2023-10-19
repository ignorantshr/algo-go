/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 *
 * https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/description/
 *
 * algorithms
 * Easy (57.99%)
 * Likes:    3239
 * Dislikes: 0
 * Total Accepted:    1.2M
 * Total Submissions: 2.1M
 * Testcase Example:  '[7,1,5,3,6,4]'
 *
 * 给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
 *
 * 你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
 *
 * 返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：[7,1,5,3,6,4]
 * 输出：5
 * 解释：在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
 * ⁠    注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。
 *
 *
 * 示例 2：
 *
 *
 * 输入：prices = [7,6,4,3,1]
 * 输出：0
 * 解释：在这种情况下, 没有交易完成, 所以最大利润为 0。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 0
 *
 *
 */
package leetcode

import (
	"math"
	"testing"
)

// @lc code=start
func maxProfit121(prices []int) int {
	return maxProfit121Optimum(prices)
	return maxProfit121Base(prices)
}

func maxProfit121Optimum(prices []int) int {
	// 优化：
	// 	简化省去 k 的状态
	// 	数组用变量代替
	size := len(prices)

	dp_i_0 := 0
	dp_i_1 := math.MinInt
	for i := 0; i < size; i++ {
		dp_i_0 = max(dp_i_0, dp_i_1+prices[i])
		dp_i_1 = max(dp_i_1, -prices[i])
	}

	return dp_i_0
}

func maxProfit121Base(prices []int) int {
	size := len(prices)
	dp := make([][][2]int, size) // 可以简化省去 k 的状态
	for i := 0; i < size; i++ {
		dp[i] = make([][2]int, 2)
		dp[i][0][0] = 0
		dp[i][0][1] = math.MinInt
	}

	for i := 0; i < size; i++ {
		for k := 1; k > 0; k-- { // 最大限制 k 只有一个数值，可以简化省去该状态
			if i-1 == -1 {
				dp[i][k][0] = 0
				dp[i][k][1] = -prices[i]
				continue
			}
			dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i]) // 可以简化省去 k 的状态
			dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
		}
	}

	return dp[size-1][1][0]
}

// @lc code=end

func Test_maxProfit121(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{"", []int{7, 1, 5, 3, 6, 4}, 5},
		{"", []int{7, 6, 4, 3, 1}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit121(tt.prices); got != tt.want {
				t.Errorf("maxProfit121() = %v, want %v", got, tt.want)
			}
		})
	}
}
