/*
 * @lc app=leetcode.cn id=44 lang=golang
 *
 * [44] 通配符匹配
 *
 * https://leetcode.cn/problems/wildcard-matching/description/
 *
 * algorithms
 * Hard (33.84%)
 * Likes:    1113
 * Dislikes: 0
 * Total Accepted:    148.1K
 * Total Submissions: 437.5K
 * Testcase Example:  '"aa"\n"a"'
 *
 * 给你一个输入字符串 (s) 和一个字符模式 (p) ，请你实现一个支持 '?' 和 '*' 匹配规则的通配符匹配：
 *
 *
 * '?' 可以匹配任何单个字符。
 * '*' 可以匹配任意字符序列（包括空字符序列）。
 *
 *
 *
 *
 * 判定匹配成功的充要条件是：字符模式必须能够 完全匹配 输入字符串（而不是部分匹配）。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "aa", p = "a"
 * 输出：false
 * 解释："a" 无法匹配 "aa" 整个字符串。
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "aa", p = "*"
 * 输出：true
 * 解释：'*' 可以匹配任意字符串。
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = "cb", p = "?a"
 * 输出：false
 * 解释：'?' 可以匹配 'c', 但第二个 'a' 无法匹配 'b'。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= s.length, p.length <= 2000
 * s 仅由小写英文字母组成
 * p 仅由小写英文字母、'?' 或 '*' 组成
 *
 *
 */
package leetcode

import (
	"testing"
)

// @lc code=start
func isMatch44(s string, p string) bool {
	m := len(s)
	n := len(p)
	dp := make([][]bool, m+1) // dp[i][j] 代表了 s[:i] 和 p[j] 的匹配情况
	for i := 0; i < m+1; i++ {
		dp[i] = make([]bool, n+1)
	}
	// dp[i][0] = false
	// dp[0][j] = false // 开头连续*号的为真
	for j := 1; j <= n && p[j-1] == '*'; j++ {
		dp[0][j] = true
	}
	dp[0][0] = true

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '?' || s[i-1] == p[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				// 选 * / 不选 *
				dp[i][j] = dp[i][j-1] || dp[i-1][j]
				// } else {
				// 	dp[i][j] = s[i-1] == p[j-1] && dp[i-1][j-1]
			}
		}
	}

	return dp[m][n]
}

// @lc code=end

func Test_isMatch44(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"a", "a?"}, false},
		{"2", args{"aa", "a?"}, true},
		{"3", args{"ab", "a?"}, true},
		{"4", args{"ab", "?b"}, true},
		{"5", args{"ab", "?"}, false},
		{"6", args{"a", "a*"}, true},
		{"6", args{"a", "*a"}, true},
		{"6", args{"a", "*aa"}, false},
		{"7", args{"aa", "a*"}, true},
		{"8", args{"aaa", "a*"}, true},
		{"9", args{"aab", "a*b"}, true},
		{"10", args{"aab", "*"}, true},
		{"11", args{"aab", "*c"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch44(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch44() = %v, want %v", got, tt.want)
			}
		})
	}
}
