/*
 * @lc app=leetcode.cn id=43 lang=golang
 *
 * [43] 字符串相乘
 *
 * https://leetcode.cn/problems/multiply-strings/description/
 *
 * algorithms
 * Medium (44.39%)
 * Likes:    1307
 * Dislikes: 0
 * Total Accepted:    323.3K
 * Total Submissions: 729.5K
 * Testcase Example:  '"2"\n"3"'
 *
 * 给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。
 *
 * 注意：不能使用任何内置的 BigInteger 库或直接将输入转换为整数。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: num1 = "2", num2 = "3"
 * 输出: "6"
 *
 * 示例 2:
 *
 *
 * 输入: num1 = "123", num2 = "456"
 * 输出: "56088"
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= num1.length, num2.length <= 200
 * num1 和 num2 只能由数字组成。
 * num1 和 num2 都不包含任何前导零，除了数字0本身。
 *
 *
 */
package leetcode

import (
	"strconv"
	"strings"
	"testing"
)

// @lc code=start
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	size1 := len(num1)
	size2 := len(num2)
	res := make([]int, size1+size2)
	promotion := 0

	for i := size2 - 1; i >= 0; i-- {
		k := size2 - 1 - i
		for j := size1 - 1; j >= 0; j-- {
			tmp := int((num1[j]-'0')*(num2[i]-'0')) + promotion + res[k]
			promotion = tmp / 10
			res[k] = tmp % 10
			k++
		}

		for promotion > 0 {
			tmp := res[k] + promotion
			promotion = tmp / 10
			res[k] = tmp % 10
			k++
		}
	}

	str := &strings.Builder{}
	k := len(res) - 1
	for i := k; i > 0 && res[i] == 0; i-- {
		k--
	}
	for i := k; i >= 0; i-- {
		str.WriteString(strconv.Itoa(res[i]))
	}
	return str.String()
}

// @lc code=end

func Test_multiply(t *testing.T) {
	type args struct {
		num1 string
		num2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"0", args{"0", "0"}, "0"},
		{"1", args{"12", "0"}, "0"},
		{"2", args{"2", "3"}, "6"},
		{"3", args{"9", "9"}, "81"},
		{"4", args{"123", "0456"}, "56088"},
		{"5", args{"10000000000", "10000000000"}, "100000000000000000000"},
		{"6", args{"100000000000000000000", "100000000000000000000"}, "10000000000000000000000000000000000000000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiply(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}
