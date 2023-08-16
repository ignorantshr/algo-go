/*
- @lc app=leetcode.cn id=242 lang=golang
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

示例 1: 输入: s = "anagram", t = "nagaram" 输出: true

示例 2: 输入: s = "rat", t = "car" 输出: false

说明: 你可以假设字符串只包含小写字母。
*/
package leetcode

import "testing"

// @lc code=start
func isAnagram(s string, t string) bool {
	counter := [26]int{}
	for _, r := range s {
		counter[r-'a']++
	}
	for _, r := range t {
		counter[r-'a']--
	}
	return counter == [26]int{}
}

func isAnagram1(s string, t string) bool {
	sm := make(map[rune]int)
	tm := make(map[rune]int)
	for _, r := range s {
		sm[r]++
	}
	for _, r := range t {
		tm[r]++
	}
	if len(sm) != len(tm) {
		return false
	}
	for k, v := range sm {
		if tm[k] != v {
			return false
		}
	}
	return true
}

// @lc code=end

func Test_isAnagram(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"0", args{
			"",
			"",
		}, true},
		{"1", args{
			"anagram",
			"nagaram",
		}, true},
		{"1", args{
			"rat",
			"car",
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagram(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}
