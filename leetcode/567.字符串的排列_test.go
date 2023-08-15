/*
 * @lc app=leetcode.cn id=567 lang=golang
 *
 * [567] 字符串的排列
 *
 * https://leetcode.cn/problems/permutation-in-string/description/
 *
 * algorithms
 * Medium (44.48%)
 * Likes:    938
 * Dislikes: 0
 * Total Accepted:    267.4K
 * Total Submissions: 600.4K
 * Testcase Example:  '"ab"\n"eidbaooo"'
 *
 * 给你两个字符串 s1 和 s2 ，写一个函数来判断 s2 是否包含 s1 的排列。如果是，返回 true ；否则，返回 false 。
 *
 * 换句话说，s1 的排列之一是 s2 的 子串 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s1 = "ab" s2 = "eidbaooo"
 * 输出：true
 * 解释：s2 包含 s1 的排列之一 ("ba").
 *
 *
 * 示例 2：
 *
 *
 * 输入：s1= "ab" s2 = "eidboaoo"
 * 输出：false
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s1.length, s2.length <= 10^4
 * s1 和 s2 仅包含小写字母
 *
 *
 */
package leetcode

import "testing"

// @lc code=start
func checkInclusion(s1 string, s2 string) bool {
	window := make(map[byte]int)
	need := make(map[byte]int)
	for i := range s1 {
		need[s1[i]]++
	}

	left, right, size := 0, 0, len(s2)
	valid := 0

	for right < size {
		c := s2[right]
		window[c]++
		if need[c] == window[c] {
			valid++
		}
		right++

		if right-left == len(s1) {
			if valid == len(need) {
				return true
			}

			d := s2[left]
			if need[d] == window[d] {
				valid--
			}
			window[d]--
			left++
		}
	}

	return false
}

// @lc code=end

func Test_checkInclusion(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{
			"b",
			"b",
		}, true},
		{"1", args{
			"b",
			"ab",
		}, true},
		{"1", args{
			"ab",
			"ab",
		}, true},
		{"1", args{
			"ab",
			"eidbaooo",
		}, true},
		{"1", args{
			"ab",
			"eidbooo",
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkInclusion(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("checkInclusion() = %v, want %v", got, tt.want)
			}
		})
	}
}
