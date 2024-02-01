/*
 * @lc app=leetcode.cn id=99 lang=golang
 *
 * [99] 恢复二叉搜索树
 *
 * https://leetcode.cn/problems/recover-binary-search-tree/description/
 *
 * algorithms
 * Medium (60.39%)
 * Likes:    933
 * Dislikes: 0
 * Total Accepted:    147.8K
 * Total Submissions: 244.4K
 * Testcase Example:  '[1,3,null,null,2]'
 *
 * 给你二叉搜索树的根节点 root ，该树中的 恰好 两个节点的值被错误地交换。请在不改变其结构的情况下，恢复这棵树 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [1,3,null,null,2]
 * 输出：[3,1,null,null,2]
 * 解释：3 不能是 1 的左孩子，因为 3 > 1 。交换 1 和 3 使二叉搜索树有效。
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [3,1,4,null,null,2]
 * 输出：[2,1,4,null,null,3]
 * 解释：2 不能在 3 的右子树中，因为 2 < 3 。交换 2 和 3 使二叉搜索树有效。
 *
 *
 *
 * 提示：
 *
 *
 * 树上节点的数目在范围 [2, 1000] 内
 * -2^31 <= Node.val <= 2^31 - 1
 *
 *
 *
 *
 * 进阶：使用 O(n) 空间复杂度的解法很容易实现。你能想出一个只使用 O(1) 空间的解决方案吗？
 *
 */
package leetcode

import "testing"

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverTree(root *TreeNode) {
	recoverTreeMirros(root)
	// recoverTreeIterate(root)
	// recoverTreeDfs(root)
}

func recoverTreeMirros(root *TreeNode) {
	var pre *TreeNode
	var first *TreeNode
	var second *TreeNode

	check := func() {
		if pre != nil && pre.Val > root.Val {
			second = root
			if first == nil {
				first = pre
			}
		}
	}

	for root != nil {
		if root.Left == nil {
			check()
			pre = root
			root = root.Right
		} else {
			predcessor := root.Left
			for predcessor.Right != nil && predcessor.Right != root {
				predcessor = predcessor.Right
			}

			if predcessor.Right == root {
				check()
				predcessor.Right = nil
				pre = root
				root = root.Right
			} else {
				predcessor.Right = root
				root = root.Left
			}
		}
	}

	first.Val, second.Val = second.Val, first.Val
}

func recoverTreeIterate(root *TreeNode) {
	stack := make([]*TreeNode, 0)
	var pre *TreeNode
	var first *TreeNode
	var second *TreeNode

	for len(stack) != 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pre != nil {
			if pre.Val > top.Val {
				if first == nil {
					first = pre
					second = top
				} else {
					second = top
					break
				}
			}
		}
		pre = top
		if top.Right != nil {
			root = top.Right
		}
	}

	first.Val, second.Val = second.Val, first.Val
}

func recoverTreeDfs(root *TreeNode) {
	var pre *TreeNode
	var first *TreeNode
	var second *TreeNode

	var dfsInOrder99 func(root *TreeNode)
	dfsInOrder99 = func(root *TreeNode) {
		if root == nil || second != nil {
			return
		}

		dfsInOrder99(root.Left)
		if pre == nil {
			pre = root
		}

		if pre.Val > root.Val {
			if first == nil {
				first = pre
				second = root
			} else {
				second = root
				return
			}
		}

		pre = root
		dfsInOrder99(root.Right)
	}
	dfsInOrder99(root)

	first.Val, second.Val = second.Val, first.Val
}

// @lc code=end

func Test_recoverTree(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want *TreeNode
	}{
		{"", NewTreeByPreOrder([]any{1, 3, nil, nil, 2}), NewTreeByPreOrder([]any{3, 1, nil, nil, 2})},
		{"", NewTreeByPreOrder([]any{3, 1, 4, nil, nil, 2}), NewTreeByPreOrder([]any{2, 1, 4, nil, nil, 3})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recoverTree(tt.root)
			if !tt.want.equal(tt.root) {
				t.Errorf("recoverTree() = %v, want %v", tt.root.bfsPrefix(), tt.want.bfsPrefix())
			}
		})
	}
}
