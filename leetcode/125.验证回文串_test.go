/*
 * @lc app=leetcode.cn id=125 lang=golang
 *
 * [125] 验证回文串
 *
 * https://leetcode.cn/problems/valid-palindrome/description/
 *
 * algorithms
 * Easy (46.58%)
 * Likes:    732
 * Dislikes: 0
 * Total Accepted:    552.7K
 * Total Submissions: 1.2M
 * Testcase Example:  '"A man, a plan, a canal: Panama"'
 *
 * 如果在将所有大写字符转换为小写字符、并移除所有非字母数字字符之后，短语正着读和反着读都一样。则可以认为该短语是一个 回文串 。
 *
 * 字母和数字都属于字母数字字符。
 *
 * 给你一个字符串 s，如果它是 回文串 ，返回 true ；否则，返回 false 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入: s = "A man, a plan, a canal: Panama"
 * 输出：true
 * 解释："amanaplanacanalpanama" 是回文串。
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "race a car"
 * 输出：false
 * 解释："raceacar" 不是回文串。
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = " "
 * 输出：true
 * 解释：在移除非字母数字字符之后，s 是一个空字符串 "" 。
 * 由于空字符串正着反着读都一样，所以是回文串。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 2 * 10^5
 * s 仅由可打印的 ASCII 字符组成
 *
 *
 */
package leetcode

import (
	"strings"
	"testing"
)

// @lc code=start
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	for i, j := 0, len(s)-1; i < j; {
		for i < j && !('0' <= s[i] && s[i] <= '9' || 'A' <= s[i] && s[i] <= 'Z' || 'a' <= s[i] && s[i] <= 'z') {
			i++
		}
		for i < j && !('0' <= s[j] && s[j] <= '9' || 'A' <= s[j] && s[j] <= 'Z' || 'a' <= s[j] && s[j] <= 'z') {
			j--
		}
		if i != j && s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

// @lc code=end

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"x.1", "ab_a", true},
		{"0.1", "", true},
		{"0.2", " ", true},
		{"0.3", "@# %", true},
		{"1.1", "@# 1%", true},
		{"1.2", "@# A%", true},
		{"1.3", "@# %b", true},
		{"1.4", "c@# %", true},
		{"1.5", "@1# %", true},
		{"2.1", "race a car", false},
		{"2.2", "A man, a plan, a canal: Panama", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChar(t *testing.T) {
	t.Log('a', 'z', 'A', 'Z')
}
