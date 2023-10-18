/*
 * @lc app=leetcode.cn id=516 lang=golang
 *
 * [516] 最长回文子序列
 *
 * https://leetcode.cn/problems/longest-palindromic-subsequence/description/
 *
 * algorithms
 * Medium (67.16%)
 * Likes:    1121
 * Dislikes: 0
 * Total Accepted:    203.2K
 * Total Submissions: 302.7K
 * Testcase Example:  '"bbbab"'
 *
 * 给你一个字符串 s ，找出其中最长的回文子序列，并返回该序列的长度。
 *
 * 子序列定义为：不改变剩余字符顺序的情况下，删除某些字符或者不删除任何字符形成的一个序列。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "bbbab"
 * 输出：4
 * 解释：一个可能的最长回文子序列为 "bbbb" 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "cbbd"
 * 输出：2
 * 解释：一个可能的最长回文子序列为 "bb" 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * s 仅由小写英文字母组成
 *
 *
 */
package leetcode

import "testing"

// @lc code=start
func longestPalindromeSubseq(s string) int {
	return longestPalindromeSubseq2(s)
	return longestPalindromeSubseq1(s)
}

func longestPalindromeSubseq2(s string) int {
	dp := make([][]int, len(s)) // dp[i][j] 表示 s[i] 开头 s[j] 结尾 的最长回文子序列
	for i := range dp {
		dp[i] = make([]int, len(s))
		dp[i][i] = 1
	}

	for i := len(s) - 1; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ { // i<=j 才有意义
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][len(s)-1]
}

func longestPalindromeSubseq1(s string) int {
	// 反转 s 得到 p，找出两者的最长公共子序列
	size := len(s)
	p := make([]byte, size)
	for i := size - 1; i >= 0; i-- {
		p[size-1-i] = s[i]
	}

	dp := make([][]int, size+1) // dp[i][j] 表示以 s[i-1]、 p[j-1] 结尾的最长公共子序列
	for i := range dp {
		dp[i] = make([]int, size+1)
	}
	dp[0][0] = 0

	for i := 1; i <= size; i++ {
		for j := 1; j <= size; j++ {
			if s[i-1] == p[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[size][size]
}

// @lc code=end

func Test_longestPalindromeSubseq(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"", "bbbab", 4},
		{"", "cbbd", 2},
		{"", "zzazz", 5},
		{"", "mbadm", 3},
		{"", "leetcode", 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindromeSubseq(tt.s); got != tt.want {
				t.Errorf("longestPalindromeSubseq() = %v, want %v", got, tt.want)
			}
		})
	}
}
