package leetcode

import (
	"reflect"
	"testing"
)

func lowestCommonAncestor_RV(root, p, q *TreeNode) *TreeNode {
	return lowestCommonAncestorDfs2_RV(root, p, q)
}

func lowestCommonAncestorDfs2_RV(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == q.Val || root.Val == p.Val { // 找到目标，向上返回目标节点
		return root
	}

	left := lowestCommonAncestorDfs2_RV(root.Left, p, q)
	right := lowestCommonAncestorDfs2_RV(root.Right, p, q)

	if left != nil && right != nil { // 双边都有
		return root
	}

	if left != nil {
		return left
	}
	return right
}

func lowestCommonAncestorDfs1_RV(root, p, q *TreeNode) *TreeNode {
	var scan func(r, target *TreeNode, p *[]*TreeNode) bool
	scan = func(r, target *TreeNode, path *[]*TreeNode) bool {
		if r == nil {
			return false
		}
		*path = append(*path, r)
		if r.Val == target.Val {
			return true
		}
		if r.Left != nil {
			if scan(r.Left, target, path) {
				return true
			}
			*path = (*path)[:len(*path)-1]
		}
		if r.Right != nil {
			if scan(r.Right, target, path) {
				return true
			}
			*path = (*path)[:len(*path)-1]
		}
		return false
	}

	ppath := make([]*TreeNode, 0)
	qpath := make([]*TreeNode, 0)
	scan(root, p, &ppath)
	scan(root, q, &qpath)

	if len(ppath) > len(qpath) {
		ppath = ppath[:len(qpath)]
	} else {
		qpath = qpath[:len(ppath)]
	}
	for i := len(ppath) - 1; i >= 0; i-- {
		if ppath[i] == qpath[i] {
			return ppath[i]
		}
	}
	return nil
}

func Test_lowestCommonAncestor_RV(t *testing.T) {
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
			&TreeNode{Val: 8},
		}, 3},
		{"1", args{
			NewTreeByPreOrder([]any{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4}),
			&TreeNode{Val: 5},
			&TreeNode{Val: 4},
		}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestor_RV(tt.args.root, tt.args.p, tt.args.q); got == nil || !reflect.DeepEqual(got.Val, tt.want) {
				t.Errorf("lowestCommonAncestor_RV() = %v, want %v", got, tt.want)
			}
		})
	}
}
