/*
 * @lc app=leetcode.cn id=538 lang=golang
 *
 * [538] 把二叉搜索树转换为累加树
 *
 * https://leetcode.cn/problems/convert-bst-to-greater-tree/description/
 *
 * algorithms
 * Medium (76.60%)
 * Likes:    920
 * Dislikes: 0
 * Total Accepted:    239.9K
 * Total Submissions: 313.1K
 * Testcase Example:  '[4,1,6,0,2,5,7,nil,nil,nil,3,nil,nil,nil,8]'
 *
 * 给出二叉 搜索 树的根节点，该树的节点值各不相同，请你将其转换为累加树（Greater Sum Tree），使每个节点 node
 * 的新值等于原树中大于或等于 node.val 的值之和。
 *
 * 提醒一下，二叉搜索树满足下列约束条件：
 *
 *
 * 节点的左子树仅包含键 小于 节点键的节点。
 * 节点的右子树仅包含键 大于 节点键的节点。
 * 左右子树也必须是二叉搜索树。
 *
 *
 * 注意：本题和 1038:
 * https://leetcode-cn.com/problems/binary-search-tree-to-greater-sum-tree/
 * 相同
 *
 *
 *
 * 示例 1：
 *
 *
 *
 * 输入：[4,1,6,0,2,5,7,nil,nil,nil,3,nil,nil,nil,8]
 * 输出：[30,36,21,36,35,26,15,nil,nil,nil,33,nil,nil,nil,8]
 *
 *
 * 示例 2：
 *
 * 输入：root = [0,nil,1]
 * 输出：[1,nil,1]
 *
 *
 * 示例 3：
 *
 * 输入：root = [1,0,2]
 * 输出：[3,3,2]
 *
 *
 * 示例 4：
 *
 * 输入：root = [3,2,4,1]
 * 输出：[7,9,4,10]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中的节点数介于 0 和 10^4^ 之间。
 * 每个节点的值介于 -10^4 和 10^4 之间。
 * 树中的所有值 互不相同 。
 * 给定的树为二叉搜索树。
 *
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
func convertBST(root *TreeNode) *TreeNode {
	return convertBSTMirros(root)
}

func convertBST1(root *TreeNode) *TreeNode {
	var sum int
	var traverse func(r *TreeNode)
	traverse = func(r *TreeNode) {
		if r == nil {
			return
		}

		traverse(r.Right)
		sum += r.Val
		r.Val = sum
		traverse(r.Left)
	}

	traverse(root)
	return root
}

// 反向 Mirros 中序遍历
func convertBSTMirros(root *TreeNode) *TreeNode {
	var sum int
	cur := root

	for cur != nil {
		if cur.Right == nil { // 无右子树
			sum += cur.Val
			cur.Val = sum
			cur = cur.Left
		} else {
			predecessor := cur.Right // 找到前驱节点
			for predecessor.Left != nil && predecessor.Left != cur {
				predecessor = predecessor.Left
			}

			if predecessor.Left == nil { // 未访问过，建立线索，以遍历完右子树后回到 cur
				predecessor.Left = cur
				cur = cur.Right
			} else {
				predecessor.Left = nil // 访问过了，恢复树
				sum += cur.Val
				cur.Val = sum
				cur = cur.Left
			}
		}
	}

	return root
}

// @lc code=end

func Test_convertBST(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want []any
	}{
		{"1", NewTreeByPreOrder([]any{0, nil, 1}), []any{1, nil, 1}},
		{"1", NewTreeByPreOrder([]any{1, 0, 2}), []any{3, 3, 2}},
		{"1", NewTreeByPreOrder([]any{3, 2, 4, 1}), []any{7, 9, 4, 10}},
		{"1", NewTreeByPreOrder([]any{4, 1, 6, 0, 2, 5, 7, nil, nil, nil, 3, nil, nil, nil, 8}), []any{30, 36, 21, 36, 35, 26, 15, nil, nil, nil, 33, nil, nil, nil, 8}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertBST(tt.root); !reflect.DeepEqual(got.bfsPrefix(), tt.want) {
				t.Errorf("convertBST() = %v, want %v", got.bfsPrefix(), tt.want)
			}
		})
	}
}
