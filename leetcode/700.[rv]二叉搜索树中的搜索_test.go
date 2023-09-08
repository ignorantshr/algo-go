package leetcode

import (
	"testing"
)

func searchBST_RV(root *TreeNode, val int) *TreeNode {
	return searchBSTIterate_RV(root, val)
}

func searchBSTIterate_RV(root *TreeNode, val int) *TreeNode {
	for root != nil {
		if root.Val == val {
			return root
		}

		if root.Val > val {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return nil
}

func searchBSTDfs_RV(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	}

	if root.Val > val {
		return searchBST_RV(root.Left, val)
	}
	if root.Val < val {
		return searchBST_RV(root.Right, val)
	}
	return nil
}

func Test_searchBST_RV(t *testing.T) {
	type args struct {
		root *TreeNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"1", args{NewTreeByPreOrder([]any{4, 2, 7, 1, 3}), 2}, NewTreeByPreOrder([]any{2, 1, 3})},
		{"1", args{NewTreeByPreOrder([]any{4, 2, 7, 1, 3}), 5}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchBST_RV(tt.args.root, tt.args.val); !got.equal(tt.want) {
				t.Errorf("searchBST_RV() = %v, want %v", got, tt.want)
			}
		})
	}
}
