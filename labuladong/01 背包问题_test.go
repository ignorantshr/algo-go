package labuladong

import "testing"

/*
给你一个可装载重量为 W 的背包和 N 个物品，每个物品有重量和价值两个属性。
其中第 i 个物品的重量为 wt[i]，价值为 val[i]。
现在让你用这个背包装物品，最多能装的价值是多少？
*/

func knapsack01(W int, wt, val []int) int {
	size := len(wt)
	dp := make([][]int, W+1) // 重量为 w 且前 n 个物品的最大价值
	for i := 0; i <= W; i++ {
		dp[i] = make([]int, size+1)
		// dp[i][0] = 0
	}
	// dp[0][i] = 0

	for w := 1; w <= W; w++ {
		for n := 1; n <= size; n++ {
			if w-wt[n-1] >= 0 {
				dp[w][n] = max(dp[w][n-1], dp[w-wt[n-1]][n-1]+val[n-1]) // 选
			} else {
				dp[w][n] = dp[w][n-1]
			}
		}
	}

	return dp[W][size]
}

func Test_knapsack01(t *testing.T) {
	type args struct {
		W   int
		wt  []int
		val []int
	}
	tests := []struct {
		// name string
		args args
		want int
	}{
		{args{0, []int{}, []int{}}, 0},
		{args{5, []int{1, 2, 3}, []int{1, 2, 3}}, 5},
		{args{5, []int{3, 2, 1}, []int{1, 2, 3}}, 5},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := knapsack01(tt.args.W, tt.args.wt, tt.args.val); got != tt.want {
				t.Errorf("knapsack01() = %v, want %v", got, tt.want)
			}
		})
	}
}
