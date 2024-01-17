/*
 * @lc app=leetcode.cn id=67 lang=golang
 *
 * [67] 二进制求和
 *
 * https://leetcode.cn/problems/add-binary/description/
 *
 * algorithms
 * Easy (52.89%)
 * Likes:    1155
 * Dislikes: 0
 * Total Accepted:    367.4K
 * Total Submissions: 693.9K
 * Testcase Example:  '"11"\n"1"'
 *
 * 给你两个二进制字符串 a 和 b ，以二进制字符串的形式返回它们的和。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入:a = "11", b = "1"
 * 输出："100"
 *
 * 示例 2：
 *
 *
 * 输入：a = "1010", b = "1011"
 * 输出："10101"
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= a.length, b.length <= 10^4
 * a 和 b 仅由字符 '0' 或 '1' 组成
 * 字符串如果不是 "0" ，就不含前导零
 *
 *
 */
package leetcode

import (
	"strconv"
	"testing"
)

// @lc code=start
func addBinary(a string, b string) string {
	// 位操作
	// 统一先计算^这样没有进位，然后统一把进位求出来，加进去。 这样可能还会有进位，所以就继续循环，直到没有进位。
	x64, _ := strconv.ParseInt(a, 2, 64)
	y64, _ := strconv.ParseInt(b, 2, 64)
	x := int(x64)
	y := int(y64)
	answer := 0
	carry := 0
	for y > 0 {
		answer = x ^ y
		carry = (x & y) << 1
		x = answer
		y = carry
	}

	return strconv.FormatInt(int64(x), 2)
}

func addBinaryCarry(a string, b string) string {
	al := len(a)
	bl := len(b)
	l := max(al, bl)

	add := 0
	res := make([]byte, 0, l)
	i := 0
	for ; i < l; i++ {
		if i < al {
			add += int(a[al-i-1] - '0')
		}
		if i < bl {
			add += int(b[bl-i-1] - '0')
		}
		res = append(res, byte(add%2)+'0')
		add /= 2
	}

	if add > 0 {
		res = append(res, byte(add%2)+'0')
	}

	l = len(res)
	for i := 0; i < l/2; i++ {
		res[i], res[l-i-1] = res[l-i-1], res[i]
	}
	return string(res)
}

// @lc code=end

func Test_addBinary(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"1", "1"}, "10"},
		{"2", args{"1", "0"}, "1"},
		{"3", args{"0", "0"}, "0"},
		{"4", args{"11", "11"}, "110"},
		{"5", args{"101", "10"}, "111"},
		{"6", args{"1111", "11"}, "10010"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addBinary(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("addBinary() = %v, want %v", got, tt.want)
			}
			if got := addBinaryCarry(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("addBinaryCarry() = %v, want %v", got, tt.want)
			}
		})
	}
}
