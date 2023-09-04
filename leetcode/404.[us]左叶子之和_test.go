/*
- @lc app=leetcode.cn id=404 lang=golang

计算给定二叉树的所有左叶子之和。
*/
package leetcode

import "testing"

// @lc code=start
func sumOfLeftLeaves(root *TreeNode) int {
	// return sumOfLeftLeavesIterate(root)
	return sumOfLeftLeavesDfs(root)
}

func sumOfLeftLeavesDfs(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var leftsum, rightsum int
	if root.Left != nil {
		if root.Left.Left == nil && root.Left.Right == nil {
			leftsum = root.Left.Val
		} else {
			leftsum = sumOfLeftLeaves(root.Left)
		}
	}
	if root.Right != nil {
		rightsum = sumOfLeftLeaves(root.Right)
	}
	return leftsum + rightsum
}

func sumOfLeftLeavesIterate(root *TreeNode) int {
	var sum int
	cur := root
	stack := make([]*TreeNode, 0)
	for cur != nil || len(stack) != 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if top.Left != nil && top.Left.Left == nil && top.Left.Right == nil {
			sum += top.Left.Val
		}
		cur = top.Right
	}
	return sum
}

// @lc code=end

func Test_leftLeafSum(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		{"1", NewTreeByPreOrder([]any{3}), 0},
		{"1", NewTreeByPreOrder([]any{3, nil, 20}), 0},
		{"1", NewTreeByPreOrder([]any{3, 9, 20}), 9},
		{"1", NewTreeByPreOrder([]any{3, 9, 20, nil, nil, 15, 7}), 24},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfLeftLeaves(tt.root); got != tt.want {
				t.Errorf("sumOfLeftLeaves() = %v, want %v", got, tt.want)
			}
		})
	}
}
