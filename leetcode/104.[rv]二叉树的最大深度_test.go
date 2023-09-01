package leetcode

import "testing"

// 前序遍历
func maxDepthRV2(root *TreeNode) int {
	var res int
	walkRv(root, 0, &res)
	return res
}

func walkRv(root *TreeNode, deep int, res *int) {
	if root == nil {
		*res = max(deep, *res)
		return
	}

	deep++
	walkRv(root.Left, deep, res)
	walkRv(root.Right, deep, res)
}

// 分解子问题
func maxDepthRV1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepthRV1(root.Left), maxDepthRV1(root.Right))
}

func maxDepthRVbfs(root *TreeNode) int {
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
		}
		queue = queue[size:]
	}
	return deep
}

func Test_maxDepthRV(t *testing.T) {
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
			if got := maxDepthRV1(tt.root); got != tt.want {
				t.Errorf("maxDepthRV() = %v, want %v", got, tt.want)
			}
		})
	}
}
