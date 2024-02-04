/*
 * @lc app=leetcode.cn id=109 lang=golang
 *
 * [109] 有序链表转换二叉搜索树
 *
 * https://leetcode.cn/problems/convert-sorted-list-to-binary-search-tree/description/
 *
 * algorithms
 * Medium (76.45%)
 * Likes:    884
 * Dislikes: 0
 * Total Accepted:    157.2K
 * Total Submissions: 205.4K
 * Testcase Example:  '[-10,-3,0,5,9]'
 *
 * 给定一个单链表的头节点  head ，其中的元素 按升序排序 ，将其转换为高度平衡的二叉搜索树。
 *
 * 本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差不超过 1。
 *
 *
 *
 * 示例 1:
 *
 *
 *
 *
 * 输入: head = [-10,-3,0,5,9]
 * 输出: [0,-3,9,-10,null,5]
 * 解释: 一个可能的答案是[0，-3,9，-10,null,5]，它表示所示的高度平衡的二叉搜索树。
 *
 *
 * 示例 2:
 *
 *
 * 输入: head = []
 * 输出: []
 *
 *
 *
 *
 * 提示:
 *
 *
 * head 中的节点数在[0, 2 * 10^4] 范围内
 * -10^5 <= Node.val <= 10^5
 *
 *
 */
package leetcode

import (
	"testing"
)

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
	return sortedListToBSTInorder(head)
	return sortedListToBSTPoint(head)
	return sortedListToBSTSlice(head)
}

// 分治+中序遍历优化
var globalHead *ListNode

func sortedListToBSTInorder(head *ListNode) *TreeNode {
	l := 0
	for cur := head; cur != nil; cur = cur.Next {
		l++
	}

	globalHead = head
	return buildTree109_2(0, l-1)
}

func buildTree109_2(left, right int) *TreeNode {
	if left > right {
		return nil
	}

	mid := (left + right) / 2
	root := &TreeNode{}
	root.Left = buildTree109_2(left, mid-1)
	root.Val = globalHead.Val
	globalHead = globalHead.Next
	root.Right = buildTree109_2(mid+1, right)
	return root
}

// 分治+双指针
func sortedListToBSTPoint(head *ListNode) *TreeNode {
	return buildTree109(head, nil)
}

func buildTree109(left, right *ListNode) *TreeNode {
	if left == right {
		return nil
	}

	mid := getMidian(left, right)
	node := &TreeNode{Val: mid.Val}
	node.Left = buildTree109(left, mid)
	node.Right = buildTree109(mid.Next, right)
	return node
}

func getMidian(left, right *ListNode) *ListNode {
	slow := left
	fast := left
	for fast != right && fast.Next != right {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

// -----双指针------

// 最好别用
func sortedListToBSTSlice(head *ListNode) *TreeNode {
	nodes := make([]*ListNode, 0)
	for head != nil {
		nodes = append(nodes, head)
		head = head.Next
	}

	var dfs func(start, end int) *TreeNode
	dfs = func(start, end int) *TreeNode {
		if start > end {
			return nil
		}

		mid := (start + end) >> 1
		node := &TreeNode{Val: nodes[mid].Val}
		node.Left = dfs(start, mid-1)
		node.Right = dfs(mid+1, end)
		return node
	}

	return dfs(0, len(nodes)-1)
}

// @lc code=end

func Test_sortedListToBST(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
	}{
		{"0", nil},
		{"2", NewList([]int{-10, -3, 0, 5, 9})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedListToBST(tt.head); !got.isAVL() {
				t.Errorf("sortedListToBST() = %v", got.bfsPrefix())
			}
		})
	}
}
