package leetcode

import (
	"reflect"
	"testing"
)

/* 给定两个节点（一定在二叉树中），求这两个节点的最近公共祖先。题目保证节点的数字各不相同。 */

// 双指针
func lowestCommonAncestor1650(p, q *PTreeNode) *PTreeNode {
	return lowestCommonAncestor1650Pointer2(p, q)
}

// 双指针
func lowestCommonAncestor1650Pointer2(p, q *PTreeNode) *PTreeNode {
	/*
		l1: m+x
		l2: n+x

		l1 = l2-(n-m)
	*/
	a, b := p, q
	alen, blen := 0, 0
	for a != nil {
		alen++
		a = a.Parent
	}
	for b != nil {
		blen++
		b = b.Parent
	}

	a, b = p, q
	if alen < blen {
		b = q
		for i := 0; i < (blen - alen); i++ {
			b = b.Parent
		}
	} else if alen > blen {
		a = p
		for i := 0; i < (alen - blen); i++ {
			a = a.Parent
		}
	}

	for a != b {
		a = a.Parent
		b = b.Parent
	}

	return a
}

// 双指针
func lowestCommonAncestor1650Pointer1(p, q *PTreeNode) *PTreeNode {
	/*
		l1: m+x
		l2: n+x

		l1+n = l2+m
	*/
	a, b := p, q
	for a != b {
		if a != nil {
			a = a.Parent
		} else {
			a = q
		}

		if b != nil {
			b = b.Parent
		} else {
			b = p
		}
	}
	return a
}

func lowestCommonAncestor1650Hash(p, q *PTreeNode) *PTreeNode {
	pnodes := make(map[*PTreeNode]bool, 0)
	for p != nil {
		pnodes[p] = true
		p = p.Parent
	}

	for q != nil {
		if pnodes[q] {
			return q
		}
		q = q.Parent
	}
	return nil
}

func Test_lowestCommonAncestor1650(t *testing.T) {
	root := buildTree1650()

	type args struct {
		p *PTreeNode
		q *PTreeNode
	}
	tests := []struct {
		name string
		args args
		want *PTreeNode
	}{
		{"1", args{root, root}, root},
		{"1", args{root.Left, root.Right}, root},
		{"1", args{root.Left.Right, root.Right.Left}, root},
		{"1", args{root.Left, root.Left.Left}, root.Left},
		{"1", args{root, root.Left.Right.Right}, root},
		{"1", args{root.Left.Right.Left, root.Left.Right.Right}, root.Left.Right},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestor1650(tt.args.p, tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lowestCommonAncestor1650() = %v, want %v", got, tt.want)
			}
		})
	}
}

func buildTree1650() *PTreeNode {
	root := &PTreeNode{Val: 3}
	root.Left = &PTreeNode{Val: 5}
	root.Right = &PTreeNode{Val: 1}
	root.Left.Left = &PTreeNode{Val: 6}
	root.Left.Right = &PTreeNode{Val: 2}
	root.Right.Left = &PTreeNode{Val: 0}
	root.Right.Right = &PTreeNode{Val: 8}
	root.Left.Right.Left = &PTreeNode{Val: 7}
	root.Left.Right.Right = &PTreeNode{Val: 4}

	root.fillParent(root, nil)
	return root
}
