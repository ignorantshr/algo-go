/*
 * @lc app=leetcode.cn id=1312 lang=golang
 *
 * [1312] 让字符串成为回文串的最少插入次数
 *
 * https://leetcode.cn/problems/minimum-insertion-steps-to-make-a-string-palindrome/description/
 *
 * algorithms
 * Hard (69.16%)
 * Likes:    212
 * Dislikes: 0
 * Total Accepted:    27.5K
 * Total Submissions: 39.5K
 * Testcase Example:  '"zzazz"'
 *
 * 给你一个字符串 s ，每一次操作你都可以在字符串的任意位置插入任意字符。
 *
 * 请你返回让 s 成为回文串的 最少操作次数 。
 *
 * 「回文串」是正读和反读都相同的字符串。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "zzazz"
 * 输出：0
 * 解释：字符串 "zzazz" 已经是回文串了，所以不需要做任何插入操作。
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "mbadm"
 * 输出：2
 * 解释：字符串可变为 "mbdadbm" 或者 "mdbabdm" 。
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = "leetcode"
 * 输出：5
 * 解释：插入 5 个字符后字符串变为 "leetcodocteel" 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 500
 * s 中所有字符都是小写字母。
 *
 *
 */
package leetcode

import "testing"

// @lc code=start
func minInsertions(s string) int {
	// 找到最长回文子序列（516 题），然后用总长度减去子序列的长度即得到答案
	size := len(s)
	dp := make([][]int, size)
	for i := 0; i < size; i++ {
		dp[i] = make([]int, size)
		dp[i][i] = 1
	}

	for i := size - 1; i >= 0; i-- {
		for j := i + 1; j < size; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}

	return size - dp[0][size-1]
}

// @lc code=end

func Test_minInsertions(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"", "a", 0},
		{"", "ab", 1},
		{"", "zzazz", 0},
		{"", "mbadm", 2},
		{"", "leetcode", 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minInsertions(tt.s); got != tt.want {
				t.Errorf("minInsertions() = %v, want %v", got, tt.want)
			}
		})
	}
}
