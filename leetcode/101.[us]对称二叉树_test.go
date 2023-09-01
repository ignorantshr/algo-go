/*
- @lc app=leetcode.cn id=101 lang=golang

给定一个二叉树，检查它是否是镜像对称的。
*/
package leetcode

import "testing"

// @lc code=start
func isSymmetric(root *TreeNode) bool {
	// return isSymmetricIterate(root)
	// return isSymmetricBfs(root)
	if root == nil {
		return true
	}
	return isSymmetricDfs(root.Left, root.Right)
}

func isSymmetricDfs(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil || left.Val != right.Val {
		return false
	}

	sameoutside := isSymmetricDfs(left.Left, right.Right)
	if !sameoutside {
		return false
	}
	return isSymmetricDfs(left.Right, right.Left)
}

func isSymmetricIterate(root *TreeNode) bool {
	if root == nil {
		return true
	}
	// 队列、栈 都可以
	quel := make([]*TreeNode, 0)
	quer := make([]*TreeNode, 0)
	quel = append(quel, root.Left)
	quer = append(quer, root.Right)

	for len(quel) > 0 && len(quer) > 0 {
		topl := quel[len(quel)-1]
		quel = quel[:len(quel)-1]
		topr := quer[len(quer)-1]
		quer = quer[:len(quer)-1]

		if topl == nil && topr == nil {
			continue
		}
		if topl == nil || topr == nil || topl.Val != topr.Val {
			return false
		}

		quel = append(quel, topl.Left)
		quel = append(quel, topl.Right)
		quer = append(quer, topr.Right)
		quer = append(quer, topr.Left)
	}
	return true
}

func isSymmetricBfs(root *TreeNode) bool {
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		size := len(queue)
		i := 0
		for j := size - 1; i < j; j-- {
			if queue[i] == nil && queue[j] == nil {
				i++
				continue
			}
			if queue[i] == nil || queue[j] == nil || queue[i].Val != queue[j].Val {
				return false
			}

			queue = append(queue, queue[i].Left)
			queue = append(queue, queue[i].Right)
			i++
		}
		for ; i < size; i++ {
			if queue[i] != nil {
				queue = append(queue, queue[i].Left)
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[size:]
	}
	return true
}

// @lc code=end

func Test_isSymmetric(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want bool
	}{
		{"1", NewTreeByPreOrder([]any{1, 2, 2, 3, nil, nil, 3}), true},
		{"1", NewTreeByPreOrder([]any{1, 2, 2, 3, 4, 4, 3}), true},
		{"2", NewTreeByPreOrder([]any{1, 2, 3}), false},
		{"2", NewTreeByPreOrder([]any{1, 3, 3, 4}), false},
		{"2", NewTreeByPreOrder([]any{1, 2, 2, nil, 3, nil, 3}), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSymmetric(tt.root); got != tt.want {
				t.Errorf("isSymmetric() = %v, want %v", got, tt.want)
			}
		})
	}
}
