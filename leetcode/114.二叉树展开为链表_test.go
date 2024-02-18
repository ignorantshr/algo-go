/*
 * @lc app=leetcode.cn id=114 lang=golang
 *
 * [114] 二叉树展开为链表
 *
 * https://leetcode.cn/problems/flatten-binary-tree-to-linked-list/description/
 *
 * algorithms
 * Medium (73.18%)
 * Likes:    1626
 * Dislikes: 0
 * Total Accepted:    425K
 * Total Submissions: 578.8K
 * Testcase Example:  '[1,2,5,3,4,null,6]'
 *
 * 给你二叉树的根结点 root ，请你将它展开为一个单链表：
 *
 *
 * 展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
 * 展开后的单链表应该与二叉树 先序遍历 顺序相同。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [1,2,5,3,4,null,6]
 * 输出：[1,null,2,null,3,null,4,null,5,null,6]
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
 * 输入：root = [0]
 * 输出：[0]
 *
 *
 * 提示：
 *
 * 树中结点数在范围 [0, 2000] 内
 * -100 <= Node.val <= 100
 *
 *
 * 进阶：你可以使用原地算法（O(1) 额外空间）展开这棵树吗？
 *
 */
package leetcode

import (
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
func flatten(root *TreeNode) {
	// flattenStack(root)
	flattenPre = nil
	flatten4(root)
	// flatten3(root)
	// flatten2(root)
	// flattenDFS(root)
}

// 保存右节点信息
func flattenStack(root *TreeNode) {
	if root == nil {
		return
	}

	stack := []*TreeNode{root}
	var pre *TreeNode
	for len(stack) != 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if pre != nil {
			pre.Right = top
			pre.Left = nil
		}
		pre = top

		if top.Right != nil {
			stack = append(stack, top.Right)
		}

		if top.Left != nil {
			stack = append(stack, top.Left)
		}
	}
}

func flatten3(root *TreeNode) {
	for root != nil {
		if root.Left != nil {
			predcessor := root.Left
			for predcessor.Right != nil {
				predcessor = predcessor.Right
			}
			predcessor.Right = root.Right
			root.Right = root.Left
			root.Left = nil
		}
		root = root.Right
	}
}

var flattenPre *TreeNode

// 正序遍历
func flatten2(root *TreeNode) {
	if root == nil {
		return
	}

	if flattenPre != nil {
		flattenPre.Right = root
		flattenPre.Left = nil
	}
	flattenPre = root

	r := root.Right
	flatten2(root.Left)
	flatten2(r)
}

// 倒序遍历
func flatten4(root *TreeNode) {
	if root == nil {
		return
	}

	flatten4(root.Right)
	flatten4(root.Left)
	root.Right = flattenPre
	root.Left = nil
	flattenPre = root
}

func flattenDFS(root *TreeNode) {
	var dfs func(r *TreeNode) *TreeNode
	dfs = func(r *TreeNode) *TreeNode {
		if r == nil {
			return nil
		}

		next := dfs(r.Right)
		r.Right = dfs(r.Left)
		cur := r
		for cur.Right != nil {
			cur = cur.Right
		}
		cur.Right = next
		r.Left = nil
		return r
	}
	dfs(root)
}

// @lc code=end

func Test_flatten(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want *TreeNode
	}{
		{"0", NewTreeByPreOrder(nil), NewTreeByPreOrder(nil)},
		{"1", NewTreeByPreOrder([]any{1}), NewTreeByPreOrder([]any{1})},
		{"2", NewTreeByPreOrder([]any{1, 2, 3}), NewTreeByPreOrder([]any{1, nil, 2, nil, nil, nil, 3})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if flatten(tt.root); !tt.root.equal(tt.want) {
				t.Errorf("flatten() = %v, want %v", tt.root.bfsPrefix(), tt.want.bfsPrefix())
			}
		})
	}
}
