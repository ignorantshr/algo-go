package leetcode

import (
	"reflect"
	"testing"
)

/*
给定一棵二叉树，再给定两个节点（未必在树中），求这两个节点的最近公共祖先。题目保证节点的数字各不相同。
*/

func lowestCommonAncestor1644(root, p, q *TreeNode) *TreeNode {
	var count int
	var dfs func(root, p, q *TreeNode) *TreeNode
	dfs = func(root, p, q *TreeNode) *TreeNode {
		if root == nil {
			return nil
		}

		left := dfs(root.Left, p, q)
		right := dfs(root.Right, p, q)

		if root.Val == q.Val || root.Val == p.Val {
			count++
			return root
		}

		if left == nil && right == nil {
			return nil
		}

		if left != nil && right != nil {
			return root
		}

		if left == nil {
			return right
		}
		return left
	}

	node := dfs(root, p, q)
	if count == 2 {
		return node
	}
	return nil
}

func Test_lowestCommonAncestor1644(t *testing.T) {
	type args struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		{"1", args{
			NewTreeByPreOrder([]any{1, 2}),
			&TreeNode{Val: 2},
			&TreeNode{Val: 1},
		}, 1},
		{"1", args{
			NewTreeByPreOrder([]any{1, 2}),
			&TreeNode{Val: 1},
			&TreeNode{Val: 2},
		}, 1},
		{"1", args{
			NewTreeByPreOrder([]any{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4}),
			&TreeNode{Val: 5},
			&TreeNode{Val: 1},
		}, 3},
		{"1", args{
			NewTreeByPreOrder([]any{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4}),
			&TreeNode{Val: 5},
			&TreeNode{Val: 4},
		}, 5},
		{"2", args{
			NewTreeByPreOrder([]any{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4}),
			&TreeNode{Val: 5},
			&TreeNode{Val: 9},
		}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestor1644(tt.args.root, tt.args.p, tt.args.q); (got == nil && tt.want != nil) || (got != nil && (nil == tt.want || !reflect.DeepEqual(got.Val, tt.want.(int)))) {
				t.Errorf("lowestCommonAncestor1644() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
