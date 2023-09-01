package leetcode

import (
	"testing"
)

// 遍历的方法就不写了，只写优化的方法

func countNodesRV(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lh := getHight(root.Left, true)
	lr := getHight(root.Right, false)
	return countNodesRVDfs(root.Left, true, lh) + countNodesRVDfs(root.Right, false, lr) + 1
}

func countNodesRVDfs(root *TreeNode, left bool, last int) int {
	// 递归计算子树的高度
	if root == nil {
		return 0
	}

	var lh, lr int
	if left {
		lh = last - 1
		lr = getHight(root.Right, false)
	} else {
		lr = last - 1
		lh = getHight(root.Left, true)
	}

	if lh == lr {
		return 2 << (lh + 1)
	}

	return 1 + countNodesRVDfs(root.Left, true, lh) + countNodesRVDfs(root.Right, false, lr)
}

func getHight(n *TreeNode, left bool) int {
	deep := 0
	for n != nil {
		deep++
		if left {
			n = n.Left
		} else {
			n = n.Right
		}
	}
	return deep
}

func Test_countNodesRV(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		// {"1", NewTreeByPreOrder([]any{}), 0},
		// {"1", NewTreeByPreOrder([]any{1}), 1},
		{"1", NewTreeByPreOrder([]any{1, 2, 3, 4}), 4},
		{"1", NewTreeByPreOrder([]any{1, 2, 3, 4, 5}), 5},
		{"1", NewTreeByPreOrder([]any{1, 2, 3, 4, 5, 6}), 6},
		{"1", NewTreeByPreOrder([]any{1, 2, 3, 4, 5, 6, 7}), 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countNodesRV(tt.root); got != tt.want {
				t.Errorf("countNodesRV() = %v, want %v", got, tt.want)
			}
		})
	}
}
