/*
- @lc app=leetcode.cn id=530 lang=golang

给你一棵所有节点为非负值的二叉搜索树，请你计算树中任意两节点的差的绝对值的最小值。
*/
package leetcode

import (
	"math"
	"testing"
)

// @lc code=start
func getMinimumDifference(root *TreeNode) int {
	// return getMinimumDifferenceIterate(root)
	return getMinimumDifferenceDfs(root)
}

func getMinimumDifferenceIterate(root *TreeNode) int {
	stack := make([]*TreeNode, 0)
	cur := root
	minv := math.MaxInt
	var pre *TreeNode

	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pre != nil && top.Val-pre.Val < minv {
			minv = top.Val - pre.Val
		}
		pre = top
		cur = top.Right
	}
	return minv
}

func getMinimumDifferenceDfs(root *TreeNode) int {
	// 中序遍历
	var pre *TreeNode
	var traversal func(r *TreeNode)
	minv := math.MaxInt
	traversal = func(r *TreeNode) {
		if r == nil {
			return
		}

		traversal(r.Left)
		if pre != nil && r.Val-pre.Val < minv {
			minv = r.Val - pre.Val
		}
		pre = r
		traversal(r.Right)
	}
	traversal(root)
	return minv
}

func getMinimumDifferenceUgly(root *TreeNode) int {
	// 比较 左子树最大值与根节点的差值 和 右子树最小值与根节点的差值
	if root == nil {
		return -1
	}

	minv := math.MaxInt
	var minabs func(r *TreeNode) int
	minabs = func(r *TreeNode) int {
		if r.Left == nil && r.Right == nil {
			return -1
		}

		leftv := -1
		if r.Left != nil {
			leftv = minabs(r.Left)
		}
		rightv := -1
		if r.Left != nil {
			rightv = minabs(r.Right)
		}
		if minv == 0 {
			return 0
		}

		if rightv != -1 && minv > rightv {
			minv = rightv
		}
		if leftv != -1 && minv > leftv {
			minv = leftv
		}

		left := r.Left
		for ; left != nil && left.Right != nil; left = left.Right {
		}

		right := r.Right
		for ; right != nil && right.Left != nil; right = right.Left {
		}

		if left == nil && right == nil {
			return -1
		}
		if right != nil && minv > abs(right.Val-r.Val) {
			minv = abs(right.Val - r.Val)
		}
		if left != nil && minv > abs(left.Val-r.Val) {
			minv = abs(left.Val - r.Val)
		}
		return minv
	}
	return minabs(root)
}

// @lc code=end

func Test_getMinimumDifference(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		{"1", NewTreeByPreOrder([]any{2, 0, 3}), 1},
		{"1", NewTreeByPreOrder([]any{2, 1, 4}), 1},
		{"1", NewTreeByPreOrder([]any{4, nil, 5}), 1},
		{"1", NewTreeByPreOrder([]any{4, nil, 7, nil, nil, nil, 8}), 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMinimumDifference(tt.root); got != tt.want {
				t.Errorf("getMinimumDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
