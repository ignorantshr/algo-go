/*
- @lc app=leetcode.cn id=513 lang=golang

给定一个二叉树，在树的最后一行找到最左边的值。
*/
package leetcode

import (
	"math"
	"testing"
)

// @lc code=start
func findBottomLeftValue(root *TreeNode) int {
	// return findBottomLeftValueDfs(root)
	return findBottomLeftValueBfs(root)
}

func findBottomLeftValueDfs(root *TreeNode) int {
	maxDepth := math.MinInt
	res := 0
	var traversal func(*TreeNode, int)
	traversal = func(n *TreeNode, depth int) {
		if n.Left == nil && n.Right == nil {
			if depth > maxDepth {
				maxDepth = depth
				res = n.Val
			}
		}

		if n.Left != nil {
			traversal(n.Left, depth+1)
		}
		if n.Right != nil {
			traversal(n.Right, depth+1)
		}
	}
	traversal(root, 0)
	return res
}

func findBottomLeftValueBfs(root *TreeNode) int {
	var res int
	queue := make([]*TreeNode, 0)
	if root != nil {
		queue = append(queue, root)
	}

	for len(queue) != 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			if i == 0 {
				res = queue[i].Val
			}
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[size:]
	}

	return res
}

// @lc code=end

func Test_findBottomLeftValue(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		{"1", NewTreeByPreOrder([]any{2}), 2},
		{"1", NewTreeByPreOrder([]any{1, 2}), 2},
		{"1", NewTreeByPreOrder([]any{1, nil, 3}), 3},
		{"1", NewTreeByPreOrder([]any{1, 2, 3}), 2},
		{"1", NewTreeByPreOrder([]any{1, 2, 3, nil, 5}), 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findBottomLeftValue(tt.root); got != tt.want {
				t.Errorf("findBottomLeftValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
