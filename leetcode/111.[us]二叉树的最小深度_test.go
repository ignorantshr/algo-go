/*
- @lc app=leetcode.cn id=111 lang=golang

给定一个二叉树，找出其最小深度。

最小深度是从根节点到最近叶子节点的最短路径上的节点数量。

说明: 叶子节点是指没有子节点的节点。

示例:

给定二叉树 [3,9,20,null,null,15,7],

返回它的最小深度 2.
*/
package leetcode

import "testing"

// @lc code=start
func minDeepth(root *TreeNode) int {
	return minDeepthDfs(root)
}

// 后续遍历/分解子问题
func minDeepthDfs(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left != nil && root.Right == nil {
		return 1 + minDeepthDfs(root.Left)
	}
	if root.Right != nil && root.Left == nil {
		return 1 + minDeepthDfs(root.Right)
	}

	l := minDeepthDfs(root.Left)
	r := minDeepthDfs(root.Right)
	if l < r {
		return 1 + l
	}
	return 1 + r
}

func minDeepthBfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	deep := 0
	for len(queue) > 0 {
		deep++
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if node.Left == nil && node.Right == nil {
				return deep
			}
		}
		queue = queue[size:]
	}
	return deep
}

// @lc code=end

func Test_minDeepth(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		{"1", nil, 0},
		{"1", &TreeNode{Val: 3}, 1},
		{"1", &TreeNode{Val: 3, Left: &TreeNode{Val: 9}}, 2},
		{"1", &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20}}, 2},
		{"1", NewTreeByPreOrder([]any{3, 9, 20, nil, nil, 15, 7}), 2},
		{"1", NewTreeByPreOrder([]any{3, 9, 20, 10, nil, 15, 7}), 3},
		{"1", NewTreeByPreOrder([]any{3, 9, 20, nil, 10, 15, 7}), 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDeepth(tt.root); got != tt.want {
				t.Errorf("minDeepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
