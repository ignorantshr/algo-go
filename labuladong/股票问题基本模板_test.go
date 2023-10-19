package labuladong

import (
	"math"
	"testing"
)

/* 股票问题基本模板 */
func maxProfit_Base(maxK int, prices []int) int {
	// 每天的盈利情况可以由三个状态组成：第几天(从 0 开始算作第一天)；次数上限 k；是否持有(0：未持有；1：持有)
	size := len(prices)
	dp := make([][][2]int, size)
	for i := 0; i < size; i++ {
		dp[i] = make([][2]int, maxK+1)
		for j := range dp[i] {
			dp[i][j] = [2]int{}
		}
		if maxK > 0 {
			dp[i][0][0] = 0
			dp[i][0][1] = math.MinInt
		}
	}

	// 一次交易由买入和卖出构成，至少需要两天。
	if maxK > size/2 {
		// 相当于 k 无上限的情况
		return 0 // todo
	}

	// base case
	// dp[-1][k][0] = 0 交易还未开始，利润当然是 0
	// dp[-1][k][1] = -infinite 还没开始的时候，是不可能持有股票的。因为我们的算法要求一个最大值，所以初始值设为一个最小值，方便取最大值。
	// dp[i][0][0] = 0 没有交易的机会，利润也是 0
	// dp[i][0][1] = -infinite 不允许交易的情况下，是不可能持有股票的。因为我们的算法要求一个最大值，所以初始值设为一个最小值，方便取最大值。

	// 变化
	for i := 0; i < size; i++ {
		for k := maxK; k > 0; k-- { // j == 0 时不会再改变结果了，无需遍历
			if i-1 == -1 { // 对 i==-1 的情况特殊处理
				dp[0][k][0] = 0          // max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
				dp[0][k][1] = -prices[i] // max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
				continue
			}

			// 第 i 天未持有 = max(第 i-1 天未持有；第 i-1 天买了，第 i 天卖了)
			dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
			// 第 i 天持有 = max(第 i-1 天已持有；第 i-1 天未持有，第 i 天买了)
			// 如果今天买了，今天就用了一次购买的机会，所以要保证昨天的机会上限比今天少一次，以给今天有机会购买
			dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
		}
	}

	return dp[size-1][maxK][0] // 最后一天如果持有股票，肯定是卖了之后的利润更高
}

func Test_maxProfit_Base(t *testing.T) {
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
			if got := maxProfit_Base(tt.args.k, tt.args.prices); got != tt.want {
				t.Errorf("maxProfit_Base() = %v, want %v", got, tt.want)
			}
		})
	}
}
