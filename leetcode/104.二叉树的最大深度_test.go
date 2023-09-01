/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
 *
 * https://leetcode.cn/problems/maximum-depth-of-binary-tree/description/
 *
 * algorithms
 * Easy (77.12%)
 * Likes:    1636
 * Dislikes: 0
 * Total Accepted:    1M
 * Total Submissions: 1.4M
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，找出其最大深度。
 *
 * 二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
 *
 * 说明: 叶子节点是指没有子节点的节点。
 *
 * 示例：
 * 给定二叉树 [3,9,20,null,null,15,7]，
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 * 返回它的最大深度 3 。
 *
 */
package leetcode

import "testing"

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
	return maxDepth1(root)
}

// 前序遍历
func maxDepth2(root *TreeNode) int {
	res := 0
	walk(root, 0, &res)
	return res
}

func walk(root *TreeNode, deep int, res *int) {
	if root == nil {
		*res = max(*res, deep)
		return
	}

	deep++
	walk(root.Left, deep, res)
	walk(root.Right, deep, res)
	// deep--
}

// 分解子问题
func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxl := maxDepth(root.Left)
	maxr := maxDepth(root.Right)
	return max(maxl, maxr) + 1
}

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// @lc code=end

func Test_maxDepth(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		{"1", nil, 0},
		{"1", &TreeNode{Val: 3}, 1},
		{"1", &TreeNode{Val: 3, Left: &TreeNode{Val: 9}}, 2},
		{"1", &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
