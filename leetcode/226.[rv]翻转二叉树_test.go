package leetcode

import (
	"reflect"
	"testing"
)

func invertTreeReview(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	r := invertTree(root.Right)
	l := invertTree(root.Left)
	root.Left = r
	root.Right = l
	return root
}

func invertTreeReview2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	invertTreeReview2(root.Left)
	invertTreeReview2(root.Right)
	return root
}

func Test_invertTreeReview(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want *TreeNode
	}{
		{"1", NewTreeByPreOrder([]any{1, 2, 3, 4, 5, 6, 7}), NewTreeByPreOrder([]any{1, 3, 2, 7, 6, 5, 4})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := invertTreeReview2(tt.root); !reflect.DeepEqual(got.bfsPrefix(), tt.want.bfsPrefix()) {
				t.Errorf("invertTreeReview() = %v, want %v", got, tt.want)
			}
		})
	}
}
