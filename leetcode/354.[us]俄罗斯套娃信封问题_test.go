/*
 * @lc app=leetcode.cn id=354 lang=golang
 *
 * [354] 俄罗斯套娃信封问题
 *
 * https://leetcode.cn/problems/russian-doll-envelopes/description/
 *
 * algorithms
 * Hard (37.36%)
 * Likes:    959
 * Dislikes: 0
 * Total Accepted:    107.3K
 * Total Submissions: 291.1K
 * Testcase Example:  '[[5,4],[6,4],[6,7],[2,3]]'
 *
 * 给你一个二维整数数组 envelopes ，其中 envelopes[i] = [wi, hi] ，表示第 i 个信封的宽度和高度。
 *
 * 当另一个信封的宽度和高度都比这个信封大的时候，这个信封就可以放进另一个信封里，如同俄罗斯套娃一样。
 *
 * 请计算 最多能有多少个 信封能组成一组“俄罗斯套娃”信封（即可以把一个信封放到另一个信封里面）。
 *
 * 注意：不允许旋转信封。
 *
 *
 * 示例 1：
 *
 *
 * 输入：envelopes = [[5,4],[6,4],[6,7],[2,3]]
 * 输出：3
 * 解释：最多信封的个数为 3, 组合为: [2,3] => [5,4] => [6,7]。
 *
 * 示例 2：
 *
 *
 * 输入：envelopes = [[1,1],[1,1],[1,1]]
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= envelopes.length <= 10^5
 * envelopes[i].length == 2
 * 1 <= wi, hi <= 10^5
 *
 *
 */
package leetcode

import (
	"sort"
	"testing"
)

// @lc code=start
func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] != envelopes[j][0] {
			return envelopes[i][0] < envelopes[j][0]
		}
		return envelopes[i][1] > envelopes[j][1] // 高度倒序，按照高度组成的数组就变成了 300 题的样子
	})

	dp := make([]int, len(envelopes))
	for i := range dp {
		dp[i] = 1
	}

	for i := 0; i < len(envelopes); i++ {
		for j := i - 1; j >= 0; j-- {
			if envelopes[i][0] < envelopes[j][0] && envelopes[i][1] < envelopes[j][1] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	res := 1
	for _, v := range dp {
		res = max(res, v)
	}
	return res
}

func maxEnvelopesMe(envelopes [][]int) int {
	maxwidth := envelopes[0][0]
	maxheigh := envelopes[0][1]
	for _, v := range envelopes {
		maxwidth = max(maxwidth, v[0])
		maxheigh = max(maxheigh, v[1])
	}

	dp := make([][]int, maxwidth+1)    // 【宽，高】为【i，j】的信封最多能套几个信封，最大的是哪个
	dpenv := make([][]int, maxwidth+1) // 【宽，高】为【i，j】的信封能套的最大信封
	for i := range dp {
		dp[i] = make([]int, maxheigh+1)
		dpenv[i] = make([]int, maxheigh+1)
	}

	for i := 1; i <= maxwidth; i++ { // 状态变化
		for j := 1; j <= maxheigh; j++ {
			num := 0
			env := 0
			if dp[i-1][j] < dp[i][j-1] { // 选择子问题的最大解
				num = dp[i][j-1]
				env = dpenv[i][j-1]
			} else {
				num = dp[i-1][j]
				env = dpenv[i-1][j]
			}
			idx := env
			for index, e := range envelopes { // 选择
				if i == e[0] && j == e[1] && (idx == 0 || (i > envelopes[env-1][0] && j > envelopes[env-1][1])) {
					idx = index + 1
					num++
					break
				}
			}
			dpenv[i][j] = idx
			dp[i][j] = num
		}
	}

	return dp[maxwidth][maxheigh]
}

// @lc code=end

func Test_maxEnvelopes(t *testing.T) {
	tests := []struct {
		name      string
		envelopes [][]int
		want      int
	}{
		{"1", [][]int{{5, 4}, {6, 5}, {6, 7}, {2, 3}}, 3},
		{"1", [][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}, 3},
		{"1", [][]int{{1, 1}, {1, 1}, {1, 1}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxEnvelopes(tt.envelopes); got != tt.want {
				t.Errorf("maxEnvelopes() = %v, want %v", got, tt.want)
			}
		})
	}
}
