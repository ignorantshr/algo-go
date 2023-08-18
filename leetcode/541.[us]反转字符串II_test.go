/*
- @lc app=leetcode.cn id=541 lang=golang

给定一个字符串 s 和一个整数 k，从字符串开头算起, 每计数至 2k 个字符，就反转这 2k 个字符中的前 k 个字符。

如果剩余字符少于 k 个，则将剩余字符全部反转。

如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样。

示例:

输入: s = "abcdefg", k = 2
输出: "bacdfeg"
*/
package leetcode

import "testing"

// @lc code=start

func reverseStr(s string, k int) string {
	sbytes := []byte(s)
	size := len(sbytes)
	l := 0
	for l+k <= size {
		reverseBytes(sbytes, l, l+k)
		l += 2 * k
	}

	reverseBytes(sbytes, l, size)

	return string(sbytes)
}

// reverse [left,right)
func reverseBytes(s []byte, left, right int) {
	right--
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

// @lc code=end

func Test_reverse541(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"0", args{
			"",
			2,
		}, ""},
		{"1", args{
			"1234567",
			2,
		}, "2134657"},
		{"1", args{
			"1234567",
			3,
		}, "3214567"},
		{"1", args{
			"12345678",
			3,
		}, "32145687"},
		{"1", args{
			"123456789",
			3,
		}, "321456987"},
		{"1", args{
			"1234567890",
			3,
		}, "3214569870"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseStr(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("reverse541() = %v, want %v", got, tt.want)
			}
		})
	}
}
