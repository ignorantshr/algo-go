/*
 * @lc app=leetcode.cn id=438 lang=golang
 *
 * [438] 找到字符串中所有字母异位词
 *
 * https://leetcode.cn/problems/find-all-anagrams-in-a-string/description/
 *
 * algorithms
 * Medium (54.67%)
 * Likes:    1235
 * Dislikes: 0
 * Total Accepted:    304.7K
 * Total Submissions: 558K
 * Testcase Example:  '"cbaebabacd"\n"abc"'
 *
 * 给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
 *
 * 异位词 指由相同字母重排列形成的字符串（包括相同的字符串）。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: s = "cbaebabacd", p = "abc"
 * 输出: [0,6]
 * 解释:
 * 起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
 * 起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。
 *
 *
 * 示例 2:
 *
 *
 * 输入: s = "abab", p = "ab"
 * 输出: [0,1,2]
 * 解释:
 * 起始索引等于 0 的子串是 "ab", 它是 "ab" 的异位词。
 * 起始索引等于 1 的子串是 "ba", 它是 "ab" 的异位词。
 * 起始索引等于 2 的子串是 "ab", 它是 "ab" 的异位词。
 *
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= s.length, p.length <= 3 * 10^4
 * s 和 p 仅包含小写字母
 *
 *
 */
package leetcode

import (
	"reflect"
	"testing"
)

// @lc code=start
func findAnagrams(s string, p string) []int {
	window := make(map[byte]int)
	need := make(map[byte]int)
	for i := range p {
		need[p[i]]++
	}

	var res []int
	valid := 0
	left, right, size := 0, 0, len(s)

	for right < size {
		c := s[right]
		window[c]++
		if window[c] == need[c] {
			valid++
		}
		right++

		if right-left == len(p) {
			if valid == len(need) {
				res = append(res, left)
			}

			d := s[left]
			if need[d] == window[d] {
				valid--
			}
			window[d]--
			left++
		}
	}

	return res
}

// @lc code=end

func Test_findAnagrams(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{
			"a",
			"a",
		}, []int{0}},
		{"1", args{
			"aa",
			"a",
		}, []int{0, 1}},
		{"1", args{
			"cbaebabacd",
			"abc",
		}, []int{0, 6}},
		{"1", args{
			"abab",
			"ab",
		}, []int{0, 1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findAnagrams(tt.args.s, tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
