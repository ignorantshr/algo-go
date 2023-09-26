/*
- @lc app=leetcode.cn id=131 lang=golang

给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。

返回 s 所有可能的分割方案。

示例: 输入: "aab" 输出: [ ["aa","b"], ["a","a","b"] ]
*/
package leetcode

import (
	"testing"
)

// @lc code=start
func partition(s string) [][]string {
	res := make([][]string, 0)
	path := make([]string, 0)

	isPalindrome := func(start, end int) bool {
		for start < end {
			if s[start] != s[end] {
				return false
			}
			start++
			end--
		}
		return true
	}

	var backtrack func(idx int)
	backtrack = func(idx int) {
		if idx == len(s) {
			tmp := make([]string, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for i := idx; i < len(s); i++ {
			if !isPalindrome(idx, i) {
				continue
			}

			path = append(path, s[idx:i+1])
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
	}

	if len(s) > 0 {
		backtrack(0)
	}
	return res
}

// @lc code=end

func Test_partition131(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want [][]string
	}{
		{"0", "", [][]string{}},
		{"1", "a", [][]string{{"a"}}},
		{"1", "aab", [][]string{{"aa", "b"}, {"a", "a", "b"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.s); !equalSetMatrix(got, tt.want) {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}
