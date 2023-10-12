/*
 * @lc app=leetcode.cn id=509 lang=c
 *
 * [509] 斐波那契数
 *
 * https://leetcode.cn/problems/fibonacci-number/description/
 *
 * algorithms
 * Easy (66.13%)
 * Likes:    707
 * Dislikes: 0
 * Total Accepted:    607.5K
 * Total Submissions: 920.2K
 * Testcase Example:  '2'
 *
 * 斐波那契数 （通常用 F(n) 表示）形成的序列称为 斐波那契数列 。该数列由 0 和 1
 * 开始，后面的每一项数字都是前面两项数字的和。也就是：
 *
 *
 * F(0) = 0，F(1) = 1
 * F(n) = F(n - 1) + F(n - 2)，其中 n > 1
 *
 *
 * 给定 n ，请计算 F(n) 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 2
 * 输出：1
 * 解释：F(2) = F(1) + F(0) = 1 + 0 = 1
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 3
 * 输出：2
 * 解释：F(3) = F(2) + F(1) = 1 + 1 = 2
 *
 *
 * 示例 3：
 *
 *
 * 输入：n = 4
 * 输出：3
 * 解释：F(4) = F(3) + F(2) = 2 + 1 = 3
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= n <= 30
 *
 *
 */
package leetcode

import "testing"

// @lc code=start
func fib(n int) int {
	return fib4(n)
}

// 初级版本
func fib1(n int) int {
	if n <= 1 {
		return n
	}
	return fib1(n-1) + fib1(n-2)
}

// 备忘录版本
func fib2(n int) int {
	mem := make([]int, n+1)

	return fib2Body(n, mem)
}

func fib2Body(n int, mem []int) int {
	if n <= 1 {
		return n
	}

	if mem[n] != 0 {
		return mem[n]
	}
	mem[n] = fib2Body(n-1, mem) + fib2Body(n-2, mem)
	return mem[n]
}

// 从底向上递推版本
func fib3(n int) int {
	mem := make([]int, n+1)

	mem[0], mem[1] = 0, 1

	for i := 2; i <= n; i++ {
		mem[i] = mem[i-1] + mem[i-2]
	}

	return mem[n]
}

// 从底向上递推优化版本
func fib4(n int) int {
	if n <= 1 {
		return n
	}

	a, b, res := 0, 1, 0

	for i := 2; i <= n; i++ {
		res = a + b
		a = b
		b = res
	}

	return res
}

// @lc code=end

func Test_fib(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"0", 0, 0},
		{"1", 1, 1},
		{"1", 2, 1},
		{"1", 3, 2},
		{"1", 4, 3},
		{"1", 20, 6765},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib(tt.n); got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}
