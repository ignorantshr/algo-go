/*
- @lc app=leetcode.cn id= lang=golang

请实现一个函数，把字符串 s 中的每个空格替换成"%20"。

示例 1： 输入：s = "We are happy."
输出："We%20are%20happy."
*/
package leetcode

import "testing"

// @lc code=start

func replaceSpace(s string) string {
	count := 0
	for i := range s {
		if s[i] == ' ' {
			count++
		}
	}
	sbytes := make([]byte, len(s)+2*count)
	for i, j := len(sbytes)-1, len(s)-1; i >= 0; j-- {
		if s[j] != ' ' {
			sbytes[i] = s[j]
			i--
		} else {
			sbytes[i] = '0'
			i--
			sbytes[i] = '2'
			i--
			sbytes[i] = '%'
			i--
		}
	}
	return string(sbytes)
}

// @lc code=end

func Test_replaceSpace(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"1", " We are happy. ", "%20We%20are%20happy.%20"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceSpace(tt.s); got != tt.want {
				t.Errorf("replaceSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}
