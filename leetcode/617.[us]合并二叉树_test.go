/*
- @lc app=leetcode.cn id=617 lang=golang

给定两个二叉树，想象当你将它们中的一个覆盖到另一个上时，两个二叉树的一些节点便会重叠。

你需要将他们合并为一个新的二叉树。合并的规则是如果两个节点重叠，那么将他们的值相加作为节点合并后的新值，否则不为 NULL 的节点将直接作为新二叉树的节点。

注意: 合并必须从两个树的根节点开始。
*/
package leetcode

import (
	"testing"
)

// @lc code=start
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	// return mergeTreesIterate(root1, root2)
	return mergeTreesDfs(root1, root2)
}

func mergeTreesIterate(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root1)
	queue = append(queue, root2)

	for len(queue) > 0 {
		n1 := queue[0]
		queue = queue[1:]
		n2 := queue[0]
		queue = queue[1:]

		n1.Val += n2.Val

		if n1.Left != nil && n2.Left != nil {
			queue = append(queue, n1.Left)
			queue = append(queue, n2.Left)
		}
		if n1.Right != nil && n2.Right != nil {
			queue = append(queue, n1.Right)
			queue = append(queue, n2.Right)
		}

		if n1.Left == nil && n2.Left != nil {
			n1.Left = n2.Left
		}
		if n1.Right == nil && n2.Right != nil {
			n1.Right = n2.Right
		}
	}
	return root1
}

func mergeTreesDfs(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}

	node := &TreeNode{Val: root1.Val + root2.Val}
	node.Left = mergeTrees(root1.Left, root2.Left)
	node.Right = mergeTrees(root1.Right, root2.Right)
	return node
}

// @lc code=end

func Test_mergeTrees(t *testing.T) {
	type args struct {
		root1 *TreeNode
		root2 *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"0", args{nil, nil}, nil},
		{"1", args{NewTreeByPreOrder([]any{1}), nil}, NewTreeByPreOrder([]any{1})},
		{"1", args{nil, NewTreeByPreOrder([]any{1})}, NewTreeByPreOrder([]any{1})},
		{"1", args{NewTreeByPreOrder([]any{1, 2}), NewTreeByPreOrder([]any{1, nil, 3, nil, nil, 4})}, NewTreeByPreOrder([]any{2, 2, 3, nil, nil, 4})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTrees(tt.args.root1, tt.args.root2); !got.equal(tt.want) {
				t.Errorf("mergeTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
