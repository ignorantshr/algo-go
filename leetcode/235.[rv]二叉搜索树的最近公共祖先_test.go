package leetcode

import (
	"reflect"
	"testing"
)

func lowestCommonAncestor235RV(root, p, q *TreeNode) *TreeNode {
	cur := root
	for {
		if cur.Val > p.Val && cur.Val > q.Val {
			cur = cur.Left
		} else if cur.Val < p.Val && cur.Val < q.Val {
			cur = cur.Right
		} else {
			return cur
		}
	}
}

func lowestCommonAncestor235RV1(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor235RV(root.Right, p, q)
	}
	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor235RV(root.Left, p, q)
	}

	right := lowestCommonAncestor235RV(root.Right, p, q)
	left := lowestCommonAncestor235RV(root.Left, p, q)

	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}

func Test_lowestCommonAncestor235RV(t *testing.T) {
	type args struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{
			NewTreeByPreOrder([]any{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5}),
			&TreeNode{Val: 3},
			&TreeNode{Val: 5},
		}, 4},
		{"1", args{
			NewTreeByPreOrder([]any{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5}),
			&TreeNode{Val: 2},
			&TreeNode{Val: 8},
		}, 6},
		{"1", args{
			NewTreeByPreOrder([]any{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5}),
			&TreeNode{Val: 2},
			&TreeNode{Val: 4},
		}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestor235RV(tt.args.root, tt.args.p, tt.args.q); got == nil || !reflect.DeepEqual(got.Val, tt.want) {
				t.Errorf("lowestCommonAncestor235RV() = %v, want %v", got, tt.want)
			}
		})
	}
}
