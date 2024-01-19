/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 *
 * https://leetcode.cn/problems/climbing-stairs/description/
 *
 * algorithms
 * Easy (54.15%)
 * Likes:    3419
 * Dislikes: 0
 * Total Accepted:    1.3M
 * Total Submissions: 2.5M
 * Testcase Example:  '2'
 *
 * 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
 *
 * 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 2
 * 输出：2
 * 解释：有两种方法可以爬到楼顶。
 * 1. 1 阶 + 1 阶
 * 2. 2 阶
 *
 * 示例 2：
 *
 *
 * 输入：n = 3
 * 输出：3
 * 解释：有三种方法可以爬到楼顶。
 * 1. 1 阶 + 1 阶 + 1 阶
 * 2. 1 阶 + 2 阶
 * 3. 2 阶 + 1 阶
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= 45
 *
 *
 */
package leetcode

import "testing"

// @lc code=start
func climbStairs(n int) int {
	res := pow70(matrix{{1, 1}, {1, 0}}, n)
	return res[0][0]
}

type matrix [2][2]int

func mul(a, b matrix) (c matrix) {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			c[i][j] = a[i][0]*b[0][j] + a[i][1]*b[1][j]
		}
	}
	return c
}

// 快速幂算法，参考第 50 题
func pow70(a matrix, n int) matrix {
	res := matrix{{1, 0}, {0, 1}}
	sum := a
	for n > 0 {
		if n&1 == 1 {
			res = mul(res, sum)
		}
		sum = mul(sum, sum)
		n >>= 1
	}
	return res
}

func climbStairsDPImprove(n int) int {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return b
}

func climbStairsDP(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// timeout,O(2^n)
// 可以用记忆化搜索优化成 O(n)
func climbStairsDFS(n int) int {
	if n < 2 {
		return 1
	}
	return climbStairsDFS(n-1) + climbStairsDFS(n-2)
}

// @lc code=end

func Test_climbStairs(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"1", 1, 1},
		{"2", 2, 2},
		{"3", 3, 3},
		{"4", 4, 5},
		{"44", 44, 1134903170},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs(tt.n); got != tt.want {
				t.Errorf("climbStairs() = %v, want %v", got, tt.want)
			}
			if got := climbStairsDPImprove(tt.n); got != tt.want {
				t.Errorf("climbStairsDPImprove() = %v, want %v", got, tt.want)
			}
			if got := climbStairsDP(tt.n); got != tt.want {
				t.Errorf("climbStairsDP() = %v, want %v", got, tt.want)
			}
			// if got := climbStairsDFS(tt.n); got != tt.want {
			// 	t.Errorf("climbStairsDFS() = %v, want %v", got, tt.want)
			// }
		})
	}
}
