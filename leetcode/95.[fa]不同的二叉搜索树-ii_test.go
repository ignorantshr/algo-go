/*
 * @lc app=leetcode.cn id=95 lang=golang
 *
 * [95] 不同的二叉搜索树 II
 *
 * https://leetcode.cn/problems/unique-binary-search-trees-ii/description/
 *
 * algorithms
 * Medium (73.63%)
 * Likes:    1523
 * Dislikes: 0
 * Total Accepted:    183.4K
 * Total Submissions: 248.5K
 * Testcase Example:  '3'
 *
 * 给你一个整数 n ，请你生成并返回所有由 n 个节点组成且节点值从 1 到 n 互不相同的不同 二叉搜索树 。可以按 任意顺序 返回答案。
 *
 *
 * 示例 1：
 * 输入：n = 3
 * 输出：[[1,null,2,null,3],[1,null,3,2],[2,1,3],[3,1,null,null,2],[3,2,null,1]]
 *
 *
 * 示例 2：
 * 输入：n = 1
 * 输出：[[1]]
 *
 * 提示：
 *
 * 1
 *
 */
package leetcode

import (
	"reflect"
	"testing"
)

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
	// 分解子问题
	var dfs func(start, end int) []*TreeNode
	dfs = func(start, end int) []*TreeNode {
		if start > end {
			return []*TreeNode{nil}
		}

		forest := make([]*TreeNode, 0)
		for i := start; i <= end; i++ {
			lefts := dfs(start, i-1)
			rights := dfs(i+1, end)
			for _, l := range lefts {
				for _, r := range rights {
					forest = append(forest, &TreeNode{i, l, r})
				}
			}
		}
		return forest
	}

	return dfs(1, n)
}

// @lc code=end

func Test_generateTrees(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []*TreeNode
	}{
		{"1", 1, []*TreeNode{NewTreeByPreOrder([]any{1})}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateTrees(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
