/*
 * @lc app=leetcode.cn id=96 lang=golang
 *
 * [96] 不同的二叉搜索树
 *
 * https://leetcode.cn/problems/unique-binary-search-trees/description/
 *
 * algorithms
 * Medium (70.86%)
 * Likes:    2299
 * Dislikes: 0
 * Total Accepted:    377.7K
 * Total Submissions: 533K
 * Testcase Example:  '3'
 *
 * 给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？返回满足题意的二叉搜索树的种数。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 3
 * 输出：5
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 1
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 *
 *
 */
package leetcode

import "testing"

// @lc code=start
func numTrees(n int) int {
	// 动态规划
	// g(n) = ∑f(i,n), 1 <= i <= n
	// f(i,n) = g(i-1)*g(n-i), 1 <= i <= n
	// g(n) =∑g(i-1)*g(n-i), 1 <= i <= n
	g := make([]int, n+1)
	g[0], g[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			g[i] += g[j-1] * g[i-j]
		}
	}
	return g[n]
}

// @lc code=end

func Test_numTrees(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"1", 1, 1},
		{"1", 2, 2},
		{"1", 3, 5},
		{"1", 4, 14},
		{"1", 19, 1767263190},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numTrees(tt.n); got != tt.want {
				t.Errorf("numTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
