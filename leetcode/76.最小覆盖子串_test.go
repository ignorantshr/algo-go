/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 *
 * https://leetcode.cn/problems/minimum-window-substring/description/
 *
 * algorithms
 * Hard (45.19%)
 * Likes:    2588
 * Dislikes: 0
 * Total Accepted:    443.8K
 * Total Submissions: 981.7K
 * Testcase Example:  '"ADOBECODEBANC"\n"ABC"'
 *
 * 给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 ""
 * 。
 *
 *
 *
 * 注意：
 *
 *
 * 对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
 * 如果 s 中存在这样的子串，我们保证它是唯一的答案。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "ADOBECODEBANC", t = "ABC"
 * 输出："BANC"
 * 解释：最小覆盖子串 "BANC" 包含来自字符串 t 的 'A'、'B' 和 'C'。
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "a", t = "a"
 * 输出："a"
 * 解释：整个字符串 s 是最小覆盖子串。
 *
 *
 * 示例 3:
 *
 *
 * 输入: s = "a", t = "aa"
 * 输出: ""
 * 解释: t 中两个字符 'a' 均应包含在 s 的子串中，
 * 因此没有符合条件的子字符串，返回空字符串。
 *
 *
 *
 * 提示：
 *
 *
 * ^m == s.length
 * ^n == t.length
 * 1 <= m, n <= 10^5
 * s 和 t 由英文字母组成
 *
 *
 *
 * 进阶：你能设计一个在 o(m+n) 时间内解决此问题的算法吗？
 */
package leetcode

import (
	"math"
	"testing"
)

/* 滑动窗口算法框架
func minWindow(s string, t string) string {
	// 这里可以根据题目说明选用更合适的数据结构以提高运行速度！
	window := make(map[byte]int) // [left, right)
	left, right, size := 0, 0, len(s)

	for right < size {
		c := s[right]
		window[c]++
		right++
		// 进行窗口内数据的一系列更新

		//debug
		for left < right && window need shrink {
			d := s[left]
			window[d]--
			left++
			// 进行窗口内数据的一系列更新
		}
	}
}
*/

// @lc code=start
func minWindow(s string, t string) string {
	window := make(map[byte]int)
	need := make(map[byte]int)
	for i := range t {
		need[t[i]]++
	}

	validN := 0
	start, length := 0, math.MaxInt32
	left, right, size := 0, 0, len(s)

	for right < size {
		c := s[right]
		if v, ok := need[c]; ok {
			window[c]++
			if v == window[c] { // 之所以不是 >=，是因为对于满足条件的字符只能计算一次有效值
				validN++
			}
		}
		right++

		for left < right && validN == len(need) {
			if length > right-left {
				start = left
				length = right - left
			}
			d := s[left]
			if v, ok := need[d]; ok {
				if v == window[d] {
					validN--
				}
				window[d]--
			}
			left++
		}
	}

	if length == math.MaxInt32 {
		return ""
	}
	return s[start : start+length]
}

// @lc code=end

func Test_minWindow(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// {"1", args{
		// 	"aaa",
		// 	"abc",
		// }, ""},
		// {"1", args{
		// 	"a",
		// 	"a",
		// }, "a"},
		{"1", args{
			"aa",
			"aa",
		}, "aa"},
		{"1", args{
			"a",
			"aa",
		}, ""},
		{"1", args{
			"ADOBECODEBANC",
			"ABC",
		}, "BANC"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minWindow(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("minWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
