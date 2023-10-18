/*
 * @lc app=leetcode.cn id=583 lang=golang
 *
 * [583] 两个字符串的删除操作
 *
 * https://leetcode.cn/problems/delete-operation-for-two-strings/description/
 *
 * algorithms
 * Medium (66.69%)
 * Likes:    629
 * Dislikes: 0
 * Total Accepted:    138.1K
 * Total Submissions: 206.6K
 * Testcase Example:  '"sea"\n"eat"'
 *
 * 给定两个单词 word1 和 word2 ，返回使得 word1 和  word2 相同所需的最小步数。
 *
 * 每步 可以删除任意一个字符串中的一个字符。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入: word1 = "sea", word2 = "eat"
 * 输出: 2
 * 解释: 第一步将 "sea" 变为 "ea" ，第二步将 "eat "变为 "ea"
 *
 *
 * 示例  2:
 *
 *
 * 输入：word1 = "leetcode", word2 = "etco"
 * 输出：4
 *
 *
 *
 *
 * 提示：
 *
 *
 *
 * 1 <= word1.length, word2.length <= 500
 * word1 和 word2 只包含小写英文字母
 *
 *
 */
package leetcode

import "testing"

// @lc code=start
func minDistance(word1 string, word2 string) int {
	return minDistance583DpTable(word1, word2)
	return minDistance583DpDfs(word1, word2)
}

// 动态规划，dp table
func minDistance583DpTable(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1) // word1[0,i) word2[0,j)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
		dp[i][0] = i // word1[0,i) word2[0,0)
	}

	for i := range dp[0] {
		dp[0][i] = i // word1[0,0) word2[0,i)
	}

	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] != word2[j-1] {
				dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1)
			} else {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}
	return dp[len(word1)][len(word2)]
}

// 动态规划，递归，加备忘录
// func minDistance583DpDfsMemo(word1 string, word2 string) int

// 动态规划，递归
func minDistance583DpDfs(word1 string, word2 string) int {
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i == len(word1) {
			return len(word2) - j
		}
		if j == len(word2) {
			return len(word1) - i
		}

		if word1[i] == word2[j] {
			return dfs(i+1, j+1) // don't delete
		} else {
			return min(dfs(i+1, j)+1, dfs(i, j+1)+1) // chose one to delete
		}
	}
	return dfs(0, 0)
}

// @lc code=end

func Test_minDistance583(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{"", ""}, 0},
		{"", args{"sea", ""}, 3},
		{"", args{"sea", "eat"}, 2},
		{"", args{"etco", "leetcode"}, 4},
		{"", args{"etbco", "leetcbode"}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDistance(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("minDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
