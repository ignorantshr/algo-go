/*
 * @lc app=leetcode.cn id=188 lang=golang
 *
 * [188] 买卖股票的最佳时机 IV
 *
 * https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iv/description/
 *
 * algorithms
 * Hard (45.98%)
 * Likes:    1094
 * Dislikes: 0
 * Total Accepted:    222.7K
 * Total Submissions: 465.4K
 * Testcase Example:  '2\n[2,4,1]'
 *
 * 给你一个整数数组 prices 和一个整数 k ，其中 prices[i] 是某支给定的股票在第 i 天的价格。
 *
 * 设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。也就是说，你最多可以买 k 次，卖 k 次。
 *
 * 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：k = 2, prices = [2,4,1]
 * 输出：2
 * 解释：在第 1 天 (股票价格 = 2) 的时候买入，在第 2 天 (股票价格 = 4) 的时候卖出，这笔交易所能获得利润 = 4-2 = 2 。
 *
 * 示例 2：
 *
 *
 * 输入：k = 2, prices = [3,2,6,5,0,3]
 * 输出：7
 * 解释：在第 2 天 (股票价格 = 2) 的时候买入，在第 3 天 (股票价格 = 6) 的时候卖出, 这笔交易所能获得利润 = 6-2 = 4 。
 * ⁠    随后，在第 5 天 (股票价格 = 0) 的时候买入，在第 6 天 (股票价格 = 3) 的时候卖出, 这笔交易所能获得利润 = 3-0 =
 * 3 。
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= k <= 100
 * 1 <= prices.length <= 1000
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
func maxProfit188(k int, prices []int) int {
	return maxProfit188Optimum(k, prices)
	return maxProfit188Base(k, prices)
}

func maxProfit188Optimum(k int, prices []int) int {
	size := len(prices)
	//一次交易由买入和卖出构成，至少需要两天。
	if k > size/2 {
		return maxProfit188Unlimit2(prices)
	}

	return maxProfit188Base(k, prices)
}

func maxProfit188Unlimit2(prices []int) int {
	size := len(prices)
	dp_i_0 := 0
	dp_i_1 := -prices[0]
	for i := 0; i < size; i++ {
		tmp := dp_i_0
		dp_i_0 = max(dp_i_0, dp_i_1+prices[i])
		dp_i_1 = max(dp_i_1, tmp-prices[i])
	}
	return dp_i_0
}

func maxProfit188Unlimit1(prices []int) int {
	size := len(prices)
	dp := make([][2]int, size)
	for i := 0; i < size; i++ {
		dp[i] = [2]int{}
	}

	for i := 0; i < size; i++ {
		if i-1 == -1 {
			dp[i][0] = 0
			dp[i][1] = -prices[i]
			continue
		}
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[size-1][0]
}

func maxProfit188Base(k int, prices []int) int {
	// 三种状态：第几天(从 0 开始算作第一天)；次数上限 k；是否持有(0：未持有；1：持有)
	size := len(prices)
	dp := make([][][2]int, size)
	for i := 0; i < size; i++ {
		dp[i] = make([][2]int, k+1)
		for j := range dp[i] {
			dp[i][j] = [2]int{}
		}
		if k > 0 {
			dp[i][0][0] = 0
			dp[i][0][1] = math.MinInt
		}
	}
	// base case
	// dp[-1][j][0] = 0 交易还未开始，利润当然是 0
	// dp[-1][j][1] = -infinite 还没开始的时候，是不可能持有股票的。因为我们的算法要求一个最大值，所以初始值设为一个最小值，方便取最大值。
	// dp[i][0][0] = 0 没有交易的机会，利润也是 0
	// dp[i][0][1] = -infinite 不允许交易的情况下，是不可能持有股票的。因为我们的算法要求一个最大值，所以初始值设为一个最小值，方便取最大值。

	// 变化
	for i := 0; i < size; i++ {
		for j := k; j > 0; j-- { // j == 0 时不会再改变结果了，无需遍历
			if i-1 == -1 { // 对 i==-1 的情况特殊处理
				dp[0][j][0] = 0          // max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
				dp[0][j][1] = -prices[i] // max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
				continue
			}

			// 第 i 天未持有 = max(第 i-1 天未持有；第 i-1 天买了，第 i 天卖了)
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
			// 第 i 天持有 = max(第 i-1 天已持有；第 i-1 天未持有，第 i 天买了)
			// 如果今天买了，今天就用了一次购买的机会，所以要保证昨天的机会上限比今天少一次，以给今天有机会购买
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
		}
	}

	return dp[size-1][k][0] // 最后一天如果持有股票，肯定是卖了之后的利润更高
}

// @lc code=end

func Test_maxProfit188(t *testing.T) {
	type args struct {
		k      int
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0", args{0, []int{2, 4, 1}}, 0},
		{"1", args{1, []int{2, 4, 1}}, 2},
		{"1", args{2, []int{2, 4, 1}}, 2},
		{"2", args{1, []int{3, 2, 6, 5, 0, 3}}, 4},
		{"2", args{2, []int{3, 2, 6, 5, 0, 3}}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit188(tt.args.k, tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
