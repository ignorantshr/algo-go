/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 *
 * https://leetcode.cn/problems/sqrtx/description/
 *
 * algorithms
 * Easy (38.42%)
 * Likes:    1499
 * Dislikes: 0
 * Total Accepted:    839.6K
 * Total Submissions: 2.2M
 * Testcase Example:  '4'
 *
 * 给你一个非负整数 x ，计算并返回 x 的 算术平方根 。
 *
 * 由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。
 *
 * 注意：不允许使用任何内置指数函数和算符，例如 pow(x, 0.5) 或者 x ** 0.5 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：x = 4
 * 输出：2
 *
 *
 * 示例 2：
 *
 *
 * 输入：x = 8
 * 输出：2
 * 解释：8 的算术平方根是 2.82842..., 由于返回类型是整数，小数部分将被舍去。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= x <= 2^31 - 1
 *
 *
 */
package leetcode

import (
	"math"
	"testing"
)

// @lc code=start
func mySqrt(x int) int {
	// 牛顿迭代法 快速求解函数零点
	// 	作者：力扣官方题解
	// 链接：https://leetcode.cn/problems/sqrtx/solutions/238553/x-de-ping-fang-gen-by-leetcode-solution/
	// 来源：力扣（LeetCode）
	// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
	if x == 0 {
		return 0
	}
	C, x0 := float64(x), float64(x)
	for {
		xi := 0.5 * (x0 + C/x0)
		if math.Abs(x0-xi) < 1e-7 {
			break
		}
		x0 = xi
	}
	return int(x0)
}

func mySqrtBinary(x int) int {
	// 二分法
	l, r := 1, x
	for l <= r {
		mid := l + (r-l)/2
		// upper bound的形式，因为我们要找的ans要是最接近于x的最大的数，利用upper bound
		if mid <= x/mid {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l - 1
}

func mySqrtE(x int) int {
	// e^(1/2*lnx)
	ans := int(math.Exp(0.5 * math.Log(float64(x))))
	if (ans+1)*(ans+1) <= x { // 运算过程中会存在误差
		return ans + 1
	}
	return ans
}

func mySqrt1(x int) int {
	if x <= 1 {
		return x
	}

	// 2^31 - 1 = 2147483647
	// 2147483647的平方根大约等于46340.950001052。
	res := x / 2
	if res > 46340 {
		res = 46340
	}
	for res*res >= x {
		res /= 2
	}

	for res*res <= x {
		res++
	}
	return res - 1
}

// @lc code=end

func Test_mySqrt(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want int
	}{
		{"0", 0, 0},
		{"1", 1, 1},
		{"1.1", 2, 1},
		{"1.2", 3, 1},
		{"2", 4, 2},
		{"2.2", 7, 2},
		{"3", 2147483647, 46340},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySqrt(tt.x); got != tt.want {
				t.Errorf("mySqrt() = %v, want %v", got, tt.want)
			}
			if got := mySqrtBinary(tt.x); got != tt.want {
				t.Errorf("mySqrtBinary() = %v, want %v", got, tt.want)
			}
			if got := mySqrtE(tt.x); got != tt.want {
				t.Errorf("mySqrtE() = %v, want %v", got, tt.want)
			}
			if got := mySqrt1(tt.x); got != tt.want {
				t.Errorf("mySqrt1() = %v, want %v", got, tt.want)
			}
		})
	}
}
