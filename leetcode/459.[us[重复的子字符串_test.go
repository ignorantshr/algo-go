/*
- @lc app=leetcode.cn id=459 lang=golang

给定一个非空的字符串，判断它是否可以由它的一个子串重复多次构成。给定的字符串只含有小写英文字母，并且长度不超过10000。

示例 1:

输入: "abab"
输出: True
解释: 可由子字符串 "ab" 重复两次构成。
示例 2:

输入: "aba"
输出: False
示例 3:

输入: "abcabcabcabc"
输出: True
解释: 可由子字符串 "abc" 重复四次构成。 (或者子字符串 "abcabc" 重复两次构成。)
*/
package leetcode

import (
	"reflect"
	"testing"
)

// @lc code=start
func repeatedSubstringPattern(s string) bool {
	return repeatedSubstringPatternKMP(s)
}

func repeatedSubstringPatternKMP(s string) bool {
	size := len(s)
	if size == 0 {
		return false
	}

	// next[i] 表示 i（包括i）之前最长相等的前后缀长度，前缀不包含最后一个字符，后缀不包含第一个字符
	next := prefixTable(s)
	if next[size-1] != 0 && size%(size-next[size-1]) == 0 {
		return true
	}

	return false
}

// 前缀表
func prefixTable(s string) []int {
	j := 0
	res := make([]int, len(s))
	res[0] = j
	for i := 1; i < len(s); i++ {
		for j > 0 && s[i] != s[j] {
			j = res[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		res[i] = j
	}
	return res
}

func repeatedSubstringPatternSymmetry(s string) bool {
	size := len(s)
	if size&1 == 1 {
		if size == 1 {
			return false
		}
		// 奇数只对应一种情况
		for i := 1; i < size; i++ {
			if s[i] != s[i-1] {
				return false
			}
		}
		return true
	}

	mid := size / 2
	for i := 0; mid < size; i++ {
		if s[i] != s[mid] {
			return false
		}
		mid++
	}
	return true
}

// @lc code=end

func Test_substr(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"1", "aabaaf", false},
		{"1", "abc", false},
		{"1", "b", false},
		{"1", "aa", true},
		{"1", "aaa", true},
		{"1", "awaw", true},
		{"1", "abcabcabcabc", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repeatedSubstringPattern(tt.s); got != tt.want {
				t.Errorf("substr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_prefixTable(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []int
	}{
		{"1", "aabaaf", []int{0, 1, 0, 1, 2, 0}},
		{"1", "abcabc", []int{0, 0, 0, 1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prefixTable(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("prefixTable() = %v, want %v", got, tt.want)
			}
		})
	}
}
