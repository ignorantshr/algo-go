/*
 * @lc app=leetcode.cn id=72 lang=golang
 *
 * [72] 编辑距离
 *
 * https://leetcode.cn/problems/edit-distance/description/
 *
 * algorithms
 * Hard (62.79%)
 * Likes:    3174
 * Dislikes: 0
 * Total Accepted:    407.7K
 * Total Submissions: 648.9K
 * Testcase Example:  '"horse"\n"ros"'
 *
 * 给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数  。
 *
 * 你可以对一个单词进行如下三种操作：
 *
 *
 * 插入一个字符
 * 删除一个字符
 * 替换一个字符
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：word1 = "horse", word2 = "ros"
 * 输出：3
 * 解释：
 * horse -> rorse (将 'h' 替换为 'r')
 * rorse -> rose (删除 'r')
 * rose -> ros (删除 'e')
 *
 *
 * 示例 2：
 *
 *
 * 输入：word1 = "intention", word2 = "execution"
 * 输出：5
 * 解释：
 * intention -> inention (删除 't')
 * inention -> enention (将 'i' 替换为 'e')
 * enention -> exention (将 'n' 替换为 'x')
 * exention -> exection (将 'n' 替换为 'c')
 * exection -> execution (插入 'u')
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= word1.length, word2.length <= 500
 * word1 和 word2 由小写英文字母组成
 *
 *
 */
package leetcode

import "testing"

// @lc code=start
func minDistance72(word1 string, word2 string) int {
	return minDistanceDp(word1, word2)
	return minDistanceDpDfsMemo(word1, word2)
	return minDistanceDpDfs(word1, word2)
}

// 动态规划，从顶向下的递归形式 优化（2）：从底向上，将 dp 函数改成 dp 数组
func minDistanceDp(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1) // dp[i][j] 是 word[:i1+1] 和 word[:i2+1] 的最小编辑距离
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
		dp[i][0] = i
	}

	for i := range dp[0] {
		dp[0][i] = i
	}

	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				m := min(
					dp[i-1][j]+1,   // del word1[i]
					dp[i-1][j-1]+1, // replace word1[i] with word2[j]
				)
				m = min(
					dp[i][j-1]+1, // insert word2[j] after word1[i]
					m,
				)
				dp[i][j] = m
			}
		}
	}
	return dp[len(word1)][len(word2)]
}

// 动态规划，从顶向下的递归形式 优化（1）：备忘录
func minDistanceDpDfsMemo(word1 string, word2 string) int {
	memo := make([][]int, len(word1))
	for i := range memo {
		memo[i] = make([]int, len(word2))
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(i1, i2 int) int
	dfs = func(i1, i2 int) int {
		if i1 == -1 { // word1 is finished
			return i2 + 1
		}
		if i2 == -1 {
			return i1 + 1
		}

		if word1[i1] == word2[i2] {
			memo[i1][i2] = dfs(i1-1, i2-1)
		} else {
			memo[i1][i2] = min(
				dfs(i1-1, i2)+1, // del word1[i1]
				min(
					dfs(i1-1, i2-1)+1, // replace word1[i1] with word2[i2]
					dfs(i1, i2-1)+1),  // insert word2[i2] after word[i]
			)
		}
		return memo[i1][i2]
	}

	return dfs(len(word1)-1, len(word2)-1)
}

// 动态规划，从顶向下的递归形式
func minDistanceDpDfs(word1 string, word2 string) int {
	var dfs func(i1, i2 int) int
	dfs = func(i1, i2 int) int {
		if i1 == -1 { // word1 is finished
			return i2 + 1
		}
		if i2 == -1 {
			return i1 + 1
		}

		if word1[i1] == word2[i2] {
			return dfs(i1-1, i2-1)
		}
		/* 如何看出这里有重复计算？ */
		/* 观察一下 dfs[i1][i2] 到 dfs[i1-1][i2-1] 是否有多条途径可走 */
		return min(
			dfs(i1-1, i2)+1, // del word1[i1]
			min(
				dfs(i1-1, i2-1)+1, // replace word1[i1] with word2[i2]
				dfs(i1, i2-1)+1),  // insert word2[i2] after word[i]
		)
	}

	return dfs(len(word1)-1, len(word2)-1)
}

// @lc code=end

func Test_minDistance(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0", args{"", ""}, 0},
		{"0", args{"horse", "ros"}, 3},
		{"0", args{"intention", "execution"}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDistance72(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("minDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
