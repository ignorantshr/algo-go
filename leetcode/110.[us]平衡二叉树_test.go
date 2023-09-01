/*
- @lc app=leetcode.cn id=110 lang=golang

给定一个二叉树，判断它是否是高度平衡的二叉树。

本题中，一棵高度平衡二叉树定义为：一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1。

示例 1:

给定二叉树 [3,9,20,nil,nil,15,7]

返回 true 。

示例 2:

给定二叉树 [1,2,2,3,3,nil,nil,4,4]

返回 false 。
*/
package leetcode

import (
	"math"
	"testing"
)

// @lc code=start

func isBalanced(root *TreeNode) bool {
	is, _ := isBalancedDfs(root)
	return is
}

func isBalancedDfs(root *TreeNode) (bool, int) {
	// 递归求子树高，比较
	if root == nil {
		return true, 0
	}

	isl, lh := isBalancedDfs(root.Left)
	isr, lr := isBalancedDfs(root.Right)

	if isl && isr && math.Abs(float64(lh)-float64(lr)) <= 1 {
		return true, max(lh, lr) + 1
	}
	return false, 0
}

// @lc code=end

func Test_isAVL(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want bool
	}{
		{"1", NewTreeByPreOrder([]any{3, 9, 20, nil, nil, 15, 7}), true},
		{"1", NewTreeByPreOrder([]any{1, 2, 2, 3, 3, nil, nil, 4, 4}), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBalanced(tt.root); got != tt.want {
				t.Errorf("isBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
