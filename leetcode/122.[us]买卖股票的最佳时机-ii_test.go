/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] 买卖股票的最佳时机 II
 *
 * https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/description/
 *
 * algorithms
 * Medium (72.10%)
 * Likes:    2329
 * Dislikes: 0
 * Total Accepted:    966.2K
 * Total Submissions: 1.3M
 * Testcase Example:  '[7,1,5,3,6,4]'
 *
 * 给你一个整数数组 prices ，其中 prices[i] 表示某支股票第 i 天的价格。
 *
 * 在每一天，你可以决定是否购买和/或出售股票。你在任何时候 最多 只能持有 一股 股票。你也可以先购买，然后在 同一天 出售。
 *
 * 返回 你能获得的 最大 利润 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：prices = [7,1,5,3,6,4]
 * 输出：7
 * 解释：在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5 - 1 = 4
 * 。
 * 随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6 - 3 = 3 。
 * ⁠    总利润为 4 + 3 = 7 。
 *
 * 示例 2：
 *
 *
 * 输入：prices = [1,2,3,4,5]
 * 输出：4
 * 解释：在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5 - 1 = 4
 * 。
 * 总利润为 4 。
 *
 * 示例 3：
 *
 *
 * 输入：prices = [7,6,4,3,1]
 * 输出：0
 * 解释：在这种情况下, 交易无法获得正利润，所以不参与交易可以获得最大利润，最大利润为 0 。
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= prices.length <= 3 * 10^4
 * 0 <= prices[i] <= 10^4
 *
 *
 */
package leetcode

import (
	"math"
	"testing"
)

// @lc code=start
func maxProfit122(prices []int) int {
	return maxProfit122Optimum(prices)
	return maxProfit122Base(prices)
}

func maxProfit122Optimum(prices []int) int {
	size := len(prices)

	dp_i_0 := 0
	dp_i_1 := math.MinInt
	for i := 0; i < size; i++ {
		tmp := dp_i_0 // 之前的 dp_i_0 会被覆盖，暂存一份
		dp_i_0 = max(dp_i_0, dp_i_1+prices[i])
		dp_i_1 = max(dp_i_1, tmp-prices[i])
	}

	return dp_i_0
}

func maxProfit122Base(prices []int) int {
	// 没有购买次数限制，省略掉 k
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

// @lc code=end

func Test_maxProfit122(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{"", []int{7, 6, 4, 3, 1}, 0},
		{"1", []int{1, 2, 3, 4, 5}, 4},
		{"1", []int{7, 1, 5, 3, 6, 4}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit122(tt.prices); got != tt.want {
				t.Errorf("maxProfit122() = %v, want %v", got, tt.want)
			}
		})
	}
}
