/*
 * @lc app=leetcode.cn id=114 lang=golang
 *
 * [114] 二叉树展开为链表
 *
 * https://leetcode.cn/problems/flatten-binary-tree-to-linked-list/description/
 *
 * algorithms
 * Medium (73.00%)
 * Likes:    1500
 * Dislikes: 0
 * Total Accepted:    367.5K
 * Total Submissions: 503.4K
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
 *
 *
 * 提示：
 *
 *
 * 树中结点数在范围 [0, 2000] 内
 * -100
 *
 *
 *
 *
 * 进阶：你可以使用原地算法（O(1) 额外空间）展开这棵树吗？
 *
 */
package leetcode

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
	flatten2(root)
}

func flatten2(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}

	l := flatten2(node.Left)
	r := flatten2(node.Right)
	if l != nil {
		node.Left = nil
		n := l
		for n.Right != nil {
			n = n.Right
		}
		n.Right = r
		node.Right = l
	} else {
		node.Right = r
	}
	return node
}

var flattenPre *TreeNode

func flatten1(node *TreeNode) {
	if node == nil {
		return
	}

	if flattenPre != nil {
		flattenPre.Right = node
		flattenPre.Left = nil
	}
	flattenPre = node

	r := node.Right // 保留原先的右节点
	flatten1(node.Left)
	flatten1(r)
}

// @lc code=end
