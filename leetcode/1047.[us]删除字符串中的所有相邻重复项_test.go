/*
- @lc app=leetcode.cn id=1047 lang=golang

给出由小写字母组成的字符串 S，重复项删除操作会选择两个相邻且相同的字母，并删除它们。

在 S 上反复执行重复项删除操作，直到无法继续删除。

在完成所有重复项删除操作后返回最终的字符串。答案保证唯一。

示例：

输入："abbaca"
输出："ca"
解释：例如，在 "abbaca" 中，我们可以删除 "bb" 由于两字母相邻且相同，这是此时唯一可以执行删除操作的重复项。之后我们得到字符串 "aaca"，其中又只有 "aa" 可以执行重复项删除操作，所以最后的字符串为 "ca"。
提示：

1 <= S.length <= 20000
S 仅由小写英文字母组成。
*/
package leetcode

import "testing"

// @lc code=start

func removeDuplicates1047(s string) string {
	return removeDuplicates1047Points(s)
}

func removeDuplicates1047Stack(s string) string {
	stack := []rune{}
	for _, r := range s {
		if len(stack) > 0 && stack[len(stack)-1] == r {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, r)
		}
	}
	return string(stack)
}

func removeDuplicates1047Points(s string) string {
	// 双指针模拟栈
	slow, fast := 0, 0
	sbytes := []byte(s)

	for fast < len(s) {
		sbytes[slow] = sbytes[fast] // 入栈
		if slow > 0 && sbytes[slow] == sbytes[slow-1] {
			slow-- // 出栈
		} else {
			slow++
		}
		fast++
	}
	return string(sbytes[:slow])
}

// @lc code=end

func Test_removeDuplicates1047(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"1", "", ""},
		{"1", "aabb", ""},
		{"1", "abba", ""},
		{"1", "abba", ""},
		{"1", "abbca", "aca"},
		{"1", "abbaca", "ca"},
		{"1", "abcddcbda", "ada"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates1047(tt.s); got != tt.want {
				t.Errorf("removeDuplicates1047() = %v, want %v", got, tt.want)
			}
		})
	}
}
