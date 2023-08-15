/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
 *
 * https://leetcode.cn/problems/binary-tree-inorder-traversal/description/
 *
 * algorithms
 * Easy (76.22%)
 * Likes:    1865
 * Dislikes: 0
 * Total Accepted:    1.2M
 * Total Submissions: 1.6M
 * Testcase Example:  '[1,null,2,3]'
 *
 * 给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [1,null,2,3]
 * 输出：[1,3,2]
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = []
 * 输出：[]
 *
 *
 * 示例 3：
 *
 *
 * 输入：root = [1]
 * 输出：[1]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点数目在范围 [0, 100] 内
 * -100 <= Node.val <= 100
 *
 *
 *
 *
 * 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
 *
 */
package leetcode

import (
	"reflect"
	"testing"
)

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	return inorderTraversalMirrors(root)
}

func inorderTraversalMirrors(root *TreeNode) []int {
	res := make([]int, 0)
	for root != nil {
		if root.Left != nil {
			predecessor := root.Left
			for predecessor.Right != nil && predecessor.Right != root {
				predecessor = predecessor.Right
			}
			if predecessor.Right == nil {
				predecessor.Right = root
				root = root.Left
			} else {
				res = append(res, root.Val)
				predecessor.Right = nil
				root = root.Right
			}
		} else {
			res = append(res, root.Val)
			root = root.Right
		}
	}

	return res
}

func inorderTraversalIterate(root *TreeNode) []int {
	res := make([]int, 0)
	nodes := make([]*TreeNode, 0)

	for root != nil || len(nodes) > 0 {
		for root != nil {
			nodes = append(nodes, root)
			root = root.Left
		}

		root = nodes[len(nodes)-1]
		nodes = nodes[:len(nodes)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return res
}

func inorderTraversalDfs(root *TreeNode) []int {
	res := make([]int, 0)

	var dfs func(r *TreeNode)
	dfs = func(r *TreeNode) {
		if r == nil {
			return
		}
		dfs(r.Left)
		res = append(res, r.Val)
		dfs(r.Right)
	}
	dfs(root)
	return res
}

// @lc code=end

func Test_inorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{NewTreeByPreOrder([]any{3, 4, nil, nil, 2, nil, 1})}, []int{4, 2, 1, 3}},
		{"1", args{}, []int{}},
		{"1", args{NewTreeByPreOrder([]any{1})}, []int{1}},
		{"1", args{NewTreeByPreOrder([]any{1, nil, 2, nil, nil, 3})}, []int{1, 3, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
