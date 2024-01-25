/*
- @lc app=leetcode.cn id=20 lang=golang

给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。
示例 1:

输入: "()"
输出: true
示例 2:

输入: "()[]{}"
输出: true
示例 3:

输入: "(]"
输出: false
示例 4:

输入: "([)]"
输出: false
示例 5:

输入: "{[]}"
输出: true
*/
package leetcode

import "testing"

// @lc code=start

func isValid(s string) bool {
	stack := make([]rune, 0)
	for _, r := range s {
		if len(stack) != 0 && match20(stack[len(stack)-1], r) {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, r)
		}
	}
	return len(stack) == 0
}

func match20(a, b rune) bool {
	if a == '(' && b == ')' {
		return true
	}
	if a == '{' && b == '}' {
		return true
	}
	if a == '[' && b == ']' {
		return true
	}
	return false
}

// @lc code=end

func Test_isValid(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"1", "", true},
		{"1", "{}", true},
		{"1", "{()}", true},
		{"1", "{([])}", true},
		{"1", "{()[]}", true},
		{"2", "{{}", false},
		{"2", "{)}", false},
		{"2", "{[)}", false},
		{"2", "{{})}", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
