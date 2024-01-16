/*
 * @lc app=leetcode.cn id=65 lang=golang
 *
 * [65] 有效数字
 *
 * https://leetcode.cn/problems/valid-number/description/
 *
 * algorithms
 * Hard (27.63%)
 * Likes:    372
 * Dislikes: 0
 * Total Accepted:    70.2K
 * Total Submissions: 253.9K
 * Testcase Example:  '"0"'
 *
 * 有效数字（按顺序）可以分成以下几个部分：
 *
 *
 * 一个 小数 或者 整数
 * （可选）一个 'e' 或 'E' ，后面跟着一个 整数
 *
 *
 * 小数（按顺序）可以分成以下几个部分：
 *
 *
 * （可选）一个符号字符（'+' 或 '-'）
 * 下述格式之一：
 *
 * 至少一位数字，后面跟着一个点 '.'
 * 至少一位数字，后面跟着一个点 '.' ，后面再跟着至少一位数字
 * 一个点 '.' ，后面跟着至少一位数字
 *
 *
 *
 *
 * 整数（按顺序）可以分成以下几个部分：
 *
 *
 * （可选）一个符号字符（'+' 或 '-'）
 * 至少一位数字
 *
 *
 * 部分有效数字列举如下：["2", "0089", "-0.1", "+3.14", "4.", "-.9", "2e10", "-90E3",
 * "3e+7", "+6e-1", "53.5e93", "-123.456e789"]
 *
 * 部分无效数字列举如下：["abc", "1a", "1e", "e3", "99e2.5", "--6", "-+3", "95a54e53"]
 *
 * 给你一个字符串 s ，如果 s 是一个 有效数字 ，请返回 true 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "0"
 * 输出：true
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "e"
 * 输出：false
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = "."
 * 输出：false
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 20
 * s 仅含英文字母（大写和小写），数字（0-9），加号 '+' ，减号 '-' ，或者点 '.' 。
 *
 *
 */
package leetcode

import "testing"

// @lc code=start
func isNumber(s string) bool {
	// 状态机
	stat := unknow
	for _, char := range s {
		stat = dfa65[stat][getCharType(byte(char))]
		if stat == "" {
			return false
		}
	}
	_, ok := dfa65[stat][end65]
	return ok
}

const (
	// 符号
	end65   = "end65"
	point   = "."
	e       = "e" // e E
	digital = "digital"
	sign65  = "sign" // + -

	// 状态
	unknow         = "unkonw"         // 待定
	integer        = "interger"       // 可带符号整数
	integernosign  = "intergernosign" // 整数
	integerrequire = "integerrequire" // 整数,要求有剩余部分
	decimal        = "decimal"        // 小数部分
	decimalrequire = "decimalrequire" // 小数部分,要求有剩余部分
	unknowrequire  = "unknowrequire"  // 要求有剩余部分
	unknowint      = "unknowint"      // 不要求有剩余部分，可以结束
)

var dfa65 = map[string]map[string]string{
	unknow: {
		sign65:  unknowrequire,  // +xxx
		digital: unknowint,      // 2xx
		point:   decimalrequire, // .xxx
	},
	unknowrequire: {
		digital: unknowint,      // +1xxx
		point:   decimalrequire, // +.xxx
	},
	unknowint: {
		digital: unknowint, // [+]12xx
		point:   decimal,   // [+]1.xx
		e:       integer,   // [+]1exx
		end65:   "",
	},
	integer: { // [+]12xxx
		sign65:  integerrequire,
		digital: integernosign,
	},
	integerrequire: {
		digital: integernosign,
	},
	integernosign: { // 12xxx
		digital: integernosign,
		end65:   "",
	},
	decimal: { // .12x[e]xx
		e:       integer,
		digital: decimal,
		end65:   "",
	},
	decimalrequire: {
		digital: decimal,
	},
}

func getCharType(char byte) string {
	if char == '.' {
		return point
	}

	if char == '+' || char == '-' {
		return sign65
	}

	if char == 'e' || char == 'E' {
		return e
	}

	if char >= '0' && char <= '9' {
		return digital
	}

	return ""
}

// @lc code=end

func Test_isNumber(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"x.1", ".", false},
		{"x.2", ".e1", false},
		{"x.3", "4e+", false},
		{"x.4", "+.", false},
		{"0.1", "a", false},
		{"0.2", "1a", false},
		{"0.3", "1e", false},
		{"0.4", "e3", false},
		{"0.5", "99e2.5", false},
		{"0.6", "--6", false},
		{"0.7", "-+3", false},
		{"0.8", "95a54e53", false},

		{"1.1", "2", true},
		{"1.2", "0089", true},
		{"1.3", "-0.1", true},
		{"1.4", "+3.14", true},
		{"1.5", "4.", true},
		{"1.6", "-.9", true},
		{"1.7", "2e10", true},
		{"1.8", "-90E3", true},
		{"1.9", "3e+7", true},
		{"1.10", "+6e-1", true},
		{"1.11", "53.5e93", true},
		{"1.12", "-123.456e789", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNumber(tt.s); got != tt.want {
				t.Errorf("isNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
