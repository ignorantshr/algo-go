/*
 * @lc app=leetcode.cn id=29 lang=golang
 *
 * [29] 两数相除
 *
 * https://leetcode.cn/problems/divide-two-integers/description/
 *
 * algorithms
 * Medium (22.24%)
 * Likes:    1203
 * Dislikes: 0
 * Total Accepted:    225.4K
 * Total Submissions: 1M
 * Testcase Example:  '10\n3'
 *
 * 给你两个整数，被除数 dividend 和除数 divisor。将两数相除，要求 不使用 乘法、除法和取余运算。
 *
 * 整数除法应该向零截断，也就是截去（truncate）其小数部分。例如，8.345 将被截断为 8 ，-2.7335 将被截断至 -2 。
 *
 * 返回被除数 dividend 除以除数 divisor 得到的 商 。
 *
 * 注意：假设我们的环境只能存储 32 位 有符号整数，其数值范围是 [−2^31,  2^31 − 1] 。本题中，如果商 严格大于 2^31 − 1
 * ，则返回 2^31 − 1 ；如果商 严格小于 -2^31 ，则返回 -2^31^ 。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: dividend = 10, divisor = 3
 * 输出: 3
 * 解释: 10/3 = 3.33333.. ，向零截断后得到 3 。
 *
 * 示例 2:
 *
 *
 * 输入: dividend = 7, divisor = -3
 * 输出: -2
 * 解释: 7/-3 = -2.33333.. ，向零截断后得到 -2 。
 *
 *
 *
 * 提示：
 *
 *
 * -2^31 <= dividend, divisor <= 2^31 - 1
 * divisor != 0
 *
 *
 */
package leetcode

import (
	"testing"
)

// @lc code=start
func divide(dividend int, divisor int) int {
	negative := false
	if dividend > 0 {
		negative = !negative
		dividend = -dividend
	}
	if divisor > 0 {
		negative = !negative
		divisor = -divisor
	}
	limit := int32(-1 << 31) // 使用负数避免溢出

	// 优化
	// if divisor == -1 {
	// 	if !negative && int32(dividend) == limit {
	// 		return -int(limit + 1)
	// 	}
	// 	if !negative {
	// 		return -dividend
	// 	}
	// 	return dividend
	// }

	closest := int32(divisor)
	mod := int32(-1) // 若是正阶数也可能超过最大值

	for closest >= int32(dividend)>>1 && mod >= limit>>1 { // 限制为一半就不必关注负数溢出的情况
		closest <<= 1 // 倍乘接近被除数
		mod <<= 1     // 记录 closest 的阶数
	}

	ans := int32(0)
	remain := int32(dividend)
	for remain <= int32(divisor) && remain != 0 {
		if remain <= closest {
			ans += mod
			remain -= closest
		}
		// 每个倍数只用得到一次
		closest >>= 1
		mod >>= 1
	}

	if ans == limit && !negative {
		return -int(limit + 1)
	}
	if !negative {
		return -int(ans)
	}
	return int(ans)
}

func divide2(dividend int, divisor int) int {
	positive := true
	if dividend > 0 {
		positive = !positive
		dividend = -dividend
	}
	if divisor > 0 {
		positive = !positive
		divisor = -divisor
	}

	limit := int32(-1 << 31)
	nums := make([]int32, 1)
	i := 0
	nums[i] = int32(divisor)
	// 找到最接近的倍乘数
	for nums[i] >= int32(dividend)>>1 && nums[i] >= limit>>1 {
		nums = append(nums, nums[i]<<1)
		i++
	}

	count := int32(0)
	// 不断逼近余数
	for j, remain := i, int32(dividend); j >= 0 && remain != 0; j-- {
		if nums[j] >= remain {
			remain -= nums[j]
			count -= int32(1 << j)
		}
	}

	if count == limit && positive {
		return 1<<31 - 1
	}
	if positive {
		return -int(count)
	}
	return int(count)
}

// 超时
func divide1(dividend int, divisor int) int {
	// dividend / divisor
	negative := false
	if dividend < 0 {
		negative = !negative
		dividend = -dividend
	}
	if divisor < 0 {
		negative = !negative
		divisor = -divisor
	}

	limit := int32(-1 << 31) // 整数溢出后变成最小的负数
	count := int32(0)
	sum := 0
	// 3/2 = 1.5 c=2, sum=4
	// -3/2 = -1.5 c=2, sum=4
	for sum <= dividend && count+1 != limit {
		count++
		sum += divisor
	}

	if sum != dividend {
		// overflow
		if count+1 == limit {
			if !negative {
				return int(count) // 1<<32 -1
			}

			if sum > dividend { // sum超过 并且 正数即将溢出
				return int(-(count - 1))
				// } else if sum == dividend { // 被相等的情况处理了
				// 	return int(-count)
			}
			return -(1 << 31)
		} else {
			count--
		}
	}

	if negative {
		count = -count
	}
	return int(count)
}

// @lc code=end

func Test_divide(t *testing.T) {
	tests := []struct {
		name     string
		dividend int
		divisor  int
		want     int
	}{
		{"1", 10, 3, 3},
		{"1", 7, 2, 3},
		{"1", 7, 1, 7},
		{"1", 7, -2, -3},
		{"1", -7, -2, 3},
		{"1", 7, -3, -2},
		{"1", -7, 3, -2},
		{"2", 1<<31 - 1, 1, (1 << 31) - 1}, // 正数最大值
		{"2", -1 << 31, 1, -(1 << 31)},     // 负数最大值
		{"3", -1 << 31, -1, (1 << 31) - 1}, // 正数溢出
		// {"3", 1<<31 + 2, -1, -(1 << 31)},   // 负数溢出
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divide(tt.dividend, tt.divisor); got != tt.want {
				t.Errorf("divide() = %v, want %v", got, tt.want)
			}
		})
	}
}
