/*
- @lc app=leetcode.cn id=450 lang=golang

给定一个二叉搜索树的根节点 root 和一个值 key，删除二叉搜索树中的 key 对应的节点，并保证二叉搜索树的性质不变。返回二叉搜索树（有可能被更新）的根节点的引用。

一般来说，删除节点可分为两个步骤：

首先找到需要删除的节点； 如果找到了，删除它。 说明： 要求算法时间复杂度为 $O(h)$，h 为树的高度。
*/
package leetcode

import (
	"testing"
)

// @lc code=start
func deleteNode(root *TreeNode, key int) *TreeNode {
	// return deleteNodeIterate(root, key)
	return deleteNodeDfs(root, key)
}

func deleteNodeSwapDel(root *TreeNode, key int) *TreeNode {
	for root != nil {
		if root.Val == key {
			if root.Right == nil {
				return root.Left
			}
		}
		if root.Val == key {
			cur := root.Left
			for cur != nil && cur.Right != nil {
				cur = cur.Right
			}
			if cur == nil {
				return root.Right
			}
			cur.Val, root.Val = root.Val, cur.Val
		}
		root.Left = deleteNodeSwapDel(root.Left, key)
		root.Right = deleteNodeSwapDel(root.Right, key)
	}
	return root
}

func deleteNodeDfs(root *TreeNode, key int) *TreeNode {
	var traversal func(cur *TreeNode) *TreeNode
	traversal = func(cur *TreeNode) *TreeNode {
		if cur == nil {
			return nil
		}

		if cur.Val < key {
			cur.Right = traversal(cur.Right)
			return cur
		} else if cur.Val > key {
			cur.Left = traversal(cur.Left)
			return cur
		}

		// cur.Val == key, 找到左树最靠右的节点替换被删除节点，更简便的做法是把左树移动到右树最靠左节点的左子节点
		var p *TreeNode
		n := cur.Left
		for n != nil && n.Right != nil {
			p = n
			n = n.Right
		}

		if n == nil {
			return cur.Right
		}

		if p != nil {
			if n.Left != nil {
				p.Right = n.Left
			} else {
				p.Right = nil
			}
			n.Left = cur.Left
		}
		n.Right = cur.Right
		cur.Left = nil
		cur.Right = nil
		return n
	}
	return traversal(root)
}

func deleteNodeIterate(root *TreeNode, key int) *TreeNode {
	dummy := &TreeNode{Left: root}
	pre := dummy
	cur := root
	for cur != nil && cur.Val != key {
		pre = cur
		if cur.Val < key {
			cur = cur.Right
		} else {
			cur = cur.Left
		}
	}
	if cur == nil {
		return root
	}

	getNextNode := func(r *TreeNode) (*TreeNode, *TreeNode) {
		var pre *TreeNode
		for r != nil && r.Right != nil {
			pre = r
			r = r.Right
		}
		return pre, r
	}

	// 从左树找下一个节点
	/*
			cur
		p		r
		  n
	*/
	if p, n := getNextNode(cur.Left); n != nil {
		if p != nil {
			if n.Left != nil {
				p.Right = n.Left
			} else {
				p.Right = nil
			}
			n.Left = cur.Left
		}
		n.Right = cur.Right
		if pre.Left.Val == cur.Val {
			pre.Left = n
		} else {
			pre.Right = n
		}
		cur.Left = nil
	} else {
		// 直接替换成右树
		if pre.Left.Val == cur.Val {
			pre.Left = cur.Right
		} else {
			pre.Right = cur.Right
		}
	}
	cur.Right = nil
	return dummy.Left
}

// @lc code=end

func Test_deleteNode(t *testing.T) {
	type args struct {
		root *TreeNode
		key  int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"0", args{
			NewTreeByPreOrder([]any{5, 2, 7}),
			0,
		}, NewTreeByPreOrder([]any{5, 2, 7}),
		},
		{"del root", args{
			NewTreeByPreOrder([]any{5, 2, 7}),
			5,
		}, NewTreeByPreOrder([]any{2, nil, 7}),
		},
		{"del root no left", args{
			NewTreeByPreOrder([]any{5, nil, 7}),
			5,
		}, NewTreeByPreOrder([]any{7}),
		},
		{"del root no right", args{
			NewTreeByPreOrder([]any{5, 2}),
			5,
		}, NewTreeByPreOrder([]any{2}),
		},
		{"del left", args{
			NewTreeByPreOrder([]any{5, 2, 7}),
			2,
		}, NewTreeByPreOrder([]any{5, nil, 7}),
		},
		{"del right", args{
			NewTreeByPreOrder([]any{5, 2, 7}),
			7,
		}, NewTreeByPreOrder([]any{5, 2}),
		},
		{"del root level 3", args{
			NewTreeByPreOrder([]any{5, 2, 7, 1, 3, 6, 8}),
			5,
		}, NewTreeByPreOrder([]any{3, 2, 7, 1, nil, 6, 8}),
		},
		{"del left level 3", args{
			NewTreeByPreOrder([]any{7, 4, nil, 2, 6, nil, nil, 1, 3}),
			4,
		}, NewTreeByPreOrder([]any{7, 3, nil, 2, 6, nil, nil, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteNodeIterate(tt.args.root, tt.args.key); !got.equal(tt.want) {
				t.Errorf("deleteNode() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
