/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
 *
 * https://leetcode.cn/problems/validate-binary-search-tree/description/
 *
 * algorithms
 * Medium (37.02%)
 * Likes:    2085
 * Dislikes: 0
 * Total Accepted:    743.4K
 * Total Submissions: 2M
 * Testcase Example:  '[2,1,3]'
 *
 * 给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。
 *
 * 有效 二叉搜索树定义如下：
 *
 *
 * 节点的左子树只包含 小于 当前节点的数。
 * 节点的右子树只包含 大于 当前节点的数。
 * 所有左子树和右子树自身必须也是二叉搜索树。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [2,1,3]
 * 输出：true
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [5,1,4,null,null,3,6]
 * 输出：false
 * 解释：根节点的值是 5 ，但是右子节点的值是 4 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点数目范围在[1, 10^4] 内
 * -2^31 <= Node.val <= 2^31 - 1
 *
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
func isValidBST(root *TreeNode) bool {
	// pre98 = nil // 多次调用时每次都初始化
	// return isValidBSTDfs2(root)
	// return isValidBSTIterate2(root)
	return isValidBSTDfs1(root)
}

func isValidBSTIterate2(root *TreeNode) bool {
	// 中序遍历通用写法
	if root == nil {
		return true
	}

	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	var pre *TreeNode
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		if top != nil {
			if top.Right != nil {
				stack = stack[:len(stack)-1]
				stack = append(stack, top.Right)
				stack = append(stack, top)
			}
			stack = append(stack, nil)
			if top.Left != nil {
				stack = append(stack, top.Left)
			}
		} else {
			stack = stack[:len(stack)-1]
			top = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if pre != nil && pre.Val >= top.Val {
				return false
			}
			pre = top
		}
	}
	return true
}

func isValidBSTIterate1(root *TreeNode) bool {
	// 中序遍历是有序的
	if root == nil {
		return true
	}

	stack := make([]*TreeNode, 0)
	cur := root
	var pre *TreeNode
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pre != nil && pre.Val >= top.Val {
			return false
		}
		pre = top
		cur = top.Right
	}
	return true
}

var pre98 *TreeNode

func isValidBSTDfs2(root *TreeNode) bool {
	// 中序遍历是有序的
	if root == nil {
		return true
	}

	if !isValidBSTDfs2(root.Left) {
		return false
	}

	if pre98 == nil || pre98.Val < root.Val {
		pre98 = root
	} else {
		return false
	}

	return isValidBSTDfs2(root.Right)
}

func isValidBSTDfs1(root *TreeNode) bool {
	var isVaild func(r *TreeNode, max, min *TreeNode) bool
	isVaild = func(r, max, min *TreeNode) bool {
		if r == nil {
			return true
		}

		if min != nil && r.Val <= min.Val {
			return false
		}
		if max != nil && r.Val >= max.Val {
			return false
		}
		return isVaild(r.Left, r, min) && isVaild(r.Right, max, r)
	}

	return isVaild(root, nil, nil)
}

// @lc code=end

func Test_isValidBST(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want bool
	}{
		{"1", NewTreeByPreOrder([]any{1}), true},
		{"2", NewTreeByPreOrder([]any{1, 2, 3}), false},
		{"3", NewTreeByPreOrder([]any{2, 1, 3}), true},
		{"4", NewTreeByPreOrder([]any{5, 1, 4, nil, nil, 3, 6}), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
