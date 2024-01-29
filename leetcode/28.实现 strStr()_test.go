/*
- @lc app=leetcode.cn id=28 lang=golang

实现 strStr() 函数。

给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。

示例 1: 输入: haystack = "hello", needle = "ll" 输出: 2

示例 2: 输入: haystack = "aaaaa", needle = "bba" 输出: -1

说明: 当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。 对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与C语言的 strstr() 以及 Java的 indexOf() 定义相符。
*/
package leetcode

import "testing"

// @lc code=start
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return -1
	}

	next := genPrfixTable(needle)
	j := 0
	for i := 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] { // 匹配失败时定位到最长匹配前缀的下一个位置
			j = next[j-1]
		}
		if haystack[i] == needle[j] { // 匹配成功前进一步准备进行下一个字符的匹配
			j++
		}
		if j == len(needle) {
			return i - j + 1
		}
	}
	return -1
}

// 生成前缀表
func genPrfixTable(s string) []int {
	if len(s) == 0 {
		return []int{}
	}
	prefix := make([]int, len(s))
	j := 0
	prefix[0] = j
	for i := 1; i < len(s); i++ {
		for j > 0 && s[i] != s[j] {
			j = prefix[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		prefix[i] = j
	}
	return prefix
}

// @lc code=end

func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{
			"aabaabaaf",
			"",
		}, -1},
		{"1", args{
			"aabaabaaf",
			"abcabcabc",
		}, -1},
		{"1", args{
			"aabaabaaf",
			"aabaaf",
		}, 3},
		{"1", args{
			"hello",
			"ll",
		}, 2},
		{"1", args{
			"aaaaa",
			"bba",
		}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
