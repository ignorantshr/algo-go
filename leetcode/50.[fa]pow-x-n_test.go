/*
 * @lc app=leetcode.cn id=50 lang=golang
 *
 * [50] Pow(x, n)
 *
 * https://leetcode.cn/problems/powx-n/description/
 *
 * algorithms
 * Medium (38.02%)
 * Likes:    1297
 * Dislikes: 0
 * Total Accepted:    419.1K
 * Total Submissions: 1.1M
 * Testcase Example:  '2.00000\n10'
 *
 * 实现 pow(x, n) ，即计算 x 的整数 n 次幂函数（即，x^n^ ）。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：x = 2.00000, n = 10
 * 输出：1024.00000
 *
 *
 * 示例 2：
 *
 *
 * 输入：x = 2.10000, n = 3
 * 输出：9.26100
 *
 *
 * 示例 3：
 *
 *
 * 输入：x = 2.00000, n = -2
 * 输出：0.25000
 * 解释：2^-2 = 1/2^2 = 1/4 = 0.25
 *
 *
 *
 *
 * 提示：
 *
 *
 * -100.0 < x < 100.0
 * -2^31 <= n <= 2^31-1
 * n 是一个整数
 * 要么 x 不为零，要么 n > 0 。
 * -10^4 <= x^n <= 10^4
 *
 *
 */
package leetcode

import (
	"testing"
)

// @lc code=start
func myPow(x float64, n int) float64 {
	if n < 0 {
		// return 1 / myPowDFS(x, -n)
		return 1 / myPowBinary(x, -n)
	}
	// return myPowDFS(x, n)
	return myPowBinary(x, n)
}

// n >= 0
func myPowDFS(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	res := myPowDFS(x, n/2)
	res = res * res
	if n&1 == 1 {
		res = res * x
	}
	return res
}

// 二进制
// https://leetcode.cn/problems/powx-n/solutions/241471/50-powx-n-kuai-su-mi-qing-xi-tu-jie-by-jyd/
func myPowBinary(x float64, n int) float64 {
	// 可以利用 n 的二进制进行计算 n = b_0 * 2^0 + b_1 * 2^1 + ... + b_(m-1) * 2^(m-1)
	// x^n = x^(b_0 * 2^0 + b_1 * 2^1 + ... + b_(m-1) * 2^(m-1))
	// 	   = x^(b_0 * 2^0) * ... * x^(b_(m-1) * 2^(m-1))

	res := 1.0
	sum := x

	for n > 0 {
		if n&1 == 1 {
			res *= sum
		}
		sum *= sum
		n = n >> 1
	}

	return res
}

// @lc code=end

func Test_myPow(t *testing.T) {
	type args struct {
		x float64
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"1", args{1, 0}, 1},
		{"1", args{2, 0}, 1},
		{"2", args{0, 1}, 0},
		{"22", args{0, 2}, 0},
		{"3", args{2, 1}, 2},
		{"33", args{2, 2}, 4},
		{"4", args{2, -2}, 0.25},
		{"44", args{2, -1}, 0.5},
		{"5", args{-1, 0}, 1},
		{"5.1", args{-1, -1}, -1},
		{"5.2", args{-1, -2}, 1},
		{"5.3", args{-2, 2}, 4},
		{"5.4", args{-2, 3}, -8},
		{"5.4", args{-2, -3}, -0.125},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myPow(tt.args.x, tt.args.n); got != tt.want {
				t.Errorf("myPow() = %v, want %v", got, tt.want)
			}
		})
	}
}
