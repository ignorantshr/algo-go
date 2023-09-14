/*
- @lc app=leetcode.cn id=669 lang=golang

给定一个二叉搜索树，同时给定最小边界L 和最大边界 R。通过修剪二叉搜索树，使得所有节点的值在[L, R]中 (R>=L) 。你可能需要改变树的根节点，所以结果应当返回修剪好的二叉搜索树的新的根节点。
*/
package leetcode

import (
	"testing"
)

// @lc code=start
func trimBST(root *TreeNode, low, hight int) *TreeNode {
	return trimBSTIterate(root, low, hight)
	// return trimBSTDfs(root, low, hight)
}

func trimBSTDfs(root *TreeNode, low, hight int) *TreeNode {
	var trim func(r *TreeNode) *TreeNode
	trim = func(r *TreeNode) *TreeNode {
		if r == nil {
			return nil
		}

		if r.Val < low {
			return trim(r.Right)
		}
		if r.Val > hight {
			return trim(r.Left)
		}
		r.Left = trim(r.Left)
		r.Right = trim(r.Right)
		return r
	}

	return trim(root)
}

func trimBSTIterate(root *TreeNode, low, hight int) *TreeNode {

	for root != nil && !(root.Val >= low && root.Val <= hight) {
		if root.Val < low {
			root = root.Right
		}
		if root.Val > hight {
			root = root.Left
		}
	}

	cur := root
	for cur != nil {
		for cur.Left != nil && cur.Left.Val < low {
			cur.Left = cur.Left.Right // 删除左节点和其左子树之后还需要继续判断新上位左节点是否满足条件
		}
		cur = cur.Left
	}

	cur = root
	for cur != nil {
		for cur.Right != nil && cur.Right.Val > hight {
			cur.Right = cur.Right.Left
		}
		cur = cur.Right
	}

	return root
}

// @lc code=end

func Test_trimBST(t *testing.T) {
	type args struct {
		root  *TreeNode
		low   int
		hight int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"0", args{
			NewTreeByPreOrder([]any{}),
			0,
			0,
		}, NewTreeByPreOrder([]any{})},
		{"single", args{
			NewTreeByPreOrder([]any{1}),
			1,
			1,
		}, NewTreeByPreOrder([]any{1})},
		{"范围在根节点左侧", args{
			NewTreeByPreOrder([]any{7, 3, 10, 1, 5}),
			1,
			6,
		}, NewTreeByPreOrder([]any{3, 1, 5})},
		{"范围在根节点左侧", args{
			NewTreeByPreOrder([]any{7, 3, 10, 1, 5}),
			1,
			3,
		}, NewTreeByPreOrder([]any{3, 1})},
		{"范围在根节点左侧", args{
			NewTreeByPreOrder([]any{7, 3, 10, 1, 5}),
			3,
			5,
		}, NewTreeByPreOrder([]any{3, nil, 5})},
		{"范围在根节点右侧", args{
			NewTreeByPreOrder([]any{7, 3, 10, nil, nil, 8, 13, nil, nil, nil, nil, nil, 9, 11, 15}),
			10,
			14,
		}, NewTreeByPreOrder([]any{10, nil, 13, nil, nil, 11})},
		{"范围在根节点右侧", args{
			NewTreeByPreOrder([]any{7, 3, 10, nil, nil, 8, 13, nil, nil, nil, nil, nil, 9, 11, 15}),
			9,
			14,
		}, NewTreeByPreOrder([]any{10, 9, 13, nil, nil, 11})},
		{"hight 等于根节点", args{
			NewTreeByPreOrder([]any{7, 3, 10, 1, 5}),
			3,
			7,
		}, NewTreeByPreOrder([]any{7, 3, nil, nil, 5})},
		{"low 等于根节点", args{
			NewTreeByPreOrder([]any{7, 3, 10, 1, 5}),
			7,
			8,
		}, NewTreeByPreOrder([]any{7})},
		{"范围包含根节点", args{
			NewTreeByPreOrder([]any{7, 3, 10, 1, 5, 8, 11}),
			2,
			9,
		}, NewTreeByPreOrder([]any{7, 3, 8, nil, 5})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trimBST(tt.args.root, tt.args.low, tt.args.hight); !got.equal(tt.want) {
				t.Errorf("trimBST() = %v, want %v", got.bfsPrefix(), tt.want.bfsPrefix())
			}
		})
	}
}
