/*
 * @lc app=leetcode.cn id=226 lang=golang
 *
 * [226] 翻转二叉树
 *
 * https://leetcode.cn/problems/invert-binary-tree/description/
 *
 * algorithms
 * Easy (79.54%)
 * Likes:    1628
 * Dislikes: 0
 * Total Accepted:    689.8K
 * Total Submissions: 867.2K
 * Testcase Example:  '[4,2,7,1,3,6,9]'
 *
 * 给你一棵二叉树的根节点 root ，翻转这棵二叉树，并返回其根节点。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：root = [4,2,7,1,3,6,9]
 * 输出：[4,7,2,9,6,3,1]
 *
 *
 * 示例 2：
 *
 *
 *
 *
 * 输入：root = [2,1,3]
 * 输出：[2,3,1]
 *
 *
 * 示例 3：
 *
 *
 * 输入：root = []
 * 输出：[]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点数目范围在 [0, 100] 内
 * -100 <= Node.val <= 100
 *
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
func invertTree(root *TreeNode) *TreeNode {
	// invertTreeWalk(root)
	invertTreeIterate(root)
	return root
}

// 迭代法，前序遍历
func invertTreeIterate(root *TreeNode) *TreeNode {
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		top.Left, top.Right = top.Right, top.Left
		if top.Right != nil {
			stack = append(stack, top.Right)
		}
		if top.Left != nil {
			stack = append(stack, top.Left)
		}
	}
	return root
}

// 遍历解法
func invertTreeWalk(root *TreeNode) {
	if root == nil {
		return
	}

	root.Left, root.Right = root.Right, root.Left
	invertTreeWalk(root.Left)
	invertTreeWalk(root.Right)
}

// 分解子问题解法
func invertTree1(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	l := invertTree1(root.Left)
	r := invertTree1(root.Right)
	root.Left = r
	root.Right = l
	return root
}

// @lc code=end
