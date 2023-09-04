/*
- @lc app=leetcode.cn id=112 lang=golang

给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。
*/
package leetcode

import (
	"testing"
)

// @lc code=start
func hasPathSum(root *TreeNode, sum int) bool {
	return hasPathSumIterate(root, sum)
	// return hasPathSumDfs(root, sum)
}

func hasPathSumDfs(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return sum == root.Val
	}

	if root.Left != nil {
		if leftExist := hasPathSum(root.Left, sum-root.Val); leftExist {
			return true
		}
	}
	if root.Right != nil {
		return hasPathSum(root.Right, sum-root.Val)
	}
	return false
}

func hasPathSumIterate(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	type pair struct {
		node    *TreeNode
		pathSum int
	}
	stack := make([]*pair, 0)
	stack = append(stack, &pair{root, root.Val})
	for len(stack) != 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if top.node.Left == nil && top.node.Right == nil && top.pathSum == sum {
			return true
		}
		if top.node.Left != nil {
			stack = append(stack, &pair{top.node.Left, top.pathSum + top.node.Left.Val})
		}
		if top.node.Right != nil {
			stack = append(stack, &pair{top.node.Right, top.pathSum + top.node.Right.Val})
		}
	}
	return false
}

// @lc code=end

func Test_hasPathSum(t *testing.T) {
	type args struct {
		root *TreeNode
		sum  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{
			NewTreeByPreOrder([]any{}),
			0,
		}, false},
		{"1", args{
			NewTreeByPreOrder([]any{1}),
			1,
		}, true},
		{"1", args{
			NewTreeByPreOrder([]any{1, 2, 3, 4, 5}),
			7,
		}, true},
		{"1", args{
			NewTreeByPreOrder([]any{1, 2, 3, 4, 5}),
			4,
		}, true},
		{"1", args{
			NewTreeByPreOrder([]any{1, 2, 3, 4, 5}),
			8,
		}, true},
		{"1", args{
			NewTreeByPreOrder([]any{1, 2, 3, 4, 5}),
			9,
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasPathSum(tt.args.root, tt.args.sum); got != tt.want {
				t.Errorf("hasPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
