/*
- @lc app=leetcode.cn id= lang=golang

字符串的左旋转操作是把字符串前面的若干个字符转移到字符串的尾部。请定义一个函数实现字符串左旋转操作的功能。比如，输入字符串"abcdefg"和数字2，该函数将返回左旋转两位得到的结果"cdefgab"。

示例 1：
输入: s = "abcdefg", k = 2
输出: "cdefgab"

示例 2：
输入: s = "lrloseumgh", k = 6
输出: "umghlrlose"

限制：
1 <= k < s.length <= 10000
*/
package leetcode

import "testing"

// @lc code=start

func leftRotate(s string, k int) string {
	sbytes := []byte(s)
	reverseBytes58(sbytes, 0, k)
	reverseBytes58(sbytes, k, len(sbytes))
	reverseBytes58(sbytes, 0, len(sbytes))
	return string(sbytes)
}

// reverse [left,right)
func reverseBytes58(s []byte, left, right int) {
	right--
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

// @lc code=end

func Test_leftRotate(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{
			"abcdefg",
			2,
		}, "cdefgab"},
		{"1", args{
			"lrloseumgh",
			6,
		}, "umghlrlose"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leftRotate(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("leftRotate() = %v, want %v", got, tt.want)
			}
		})
	}
}
