/*
 * @lc app=leetcode.cn id=714 lang=golang
 *
 * [714] 买卖股票的最佳时机含手续费
 *
 * https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/description/
 *
 * algorithms
 * Medium (75.33%)
 * Likes:    1008
 * Dislikes: 0
 * Total Accepted:    247.9K
 * Total Submissions: 325.9K
 * Testcase Example:  '[1,3,2,8,4,9]\n2'
 *
 * 给定一个整数数组 prices，其中 prices[i]表示第 i 天的股票价格 ；整数 fee 代表了交易股票的手续费用。
 *
 * 你可以无限次地完成交易，但是你每笔交易都需要付手续费。如果你已经购买了一个股票，在卖出它之前你就不能再继续购买股票了。
 *
 * 返回获得利润的最大值。
 *
 * 注意：这里的一笔交易指买入持有并卖出股票的整个过程，每笔交易你只需要为支付一次手续费。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：prices = [1, 3, 2, 8, 4, 9], fee = 2
 * 输出：8
 * 解释：能够达到的最大利润:
 * 在此处买入 prices[0] = 1
 * 在此处卖出 prices[3] = 8
 * 在此处买入 prices[4] = 4
 * 在此处卖出 prices[5] = 9
 * 总利润: ((8 - 1) - 2) + ((9 - 4) - 2) = 8
 *
 * 示例 2：
 *
 *
 * 输入：prices = [1,3,7,5,10,3], fee = 3
 * 输出：6
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= prices.length <= 5 * 10^4
 * 1 <= prices[i] < 5 * 10^4
 * 0 <= fee < 5 * 10^4
 *
 *
 */
package leetcode

import (
	"math"
	"testing"
)

// @lc code=start
func maxProfit714(prices []int, fee int) int {
	return maxProfit714Optimum(prices, fee)
	return maxProfit714Base(prices, fee)
}

func maxProfit714Optimum(prices []int, fee int) int {
	size := len(prices)
	dp_i_0 := 0
	dp_i_1 := math.MinInt

	for i := 0; i < size; i++ {
		tmp := dp_i_0
		dp_i_0 = max(dp_i_0, dp_i_1+prices[i])
		dp_i_1 = max(dp_i_1, tmp-prices[i]-fee)
	}

	return dp_i_0
}

func maxProfit714Base(prices []int, fee int) int {
	size := len(prices)
	dp := make([][2]int, size)
	for i := 0; i < size; i++ {
		dp[i] = [2]int{}
	}

	for i := 0; i < size; i++ {
		if i-1 == -1 {
			dp[i][0] = 0
			dp[i][1] = -prices[i] - fee
			continue
		}

		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i]-fee)
	}

	return dp[size-1][0]
}

// @lc code=end

func Test_maxProfit714(t *testing.T) {
	type args struct {
		prices []int
		fee    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{[]int{1, 3, 2, 8, 4, 9}, 2}, 8},
		{"", args{[]int{1, 3, 7, 5, 10, 3}, 3}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit714(tt.args.prices, tt.args.fee); got != tt.want {
				t.Errorf("maxProfit714() = %v, want %v", got, tt.want)
			}
		})
	}
}
