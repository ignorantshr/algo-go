/*
 * @lc app=leetcode.cn id=107 lang=golang
 *
 * [107] 二叉树的层序遍历 II
 *
 * https://leetcode.cn/problems/binary-tree-level-order-traversal-ii/description/
 *
 * algorithms
 * Medium (72.83%)
 * Likes:    764
 * Dislikes: 0
 * Total Accepted:    307.3K
 * Total Submissions: 420.2K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给你二叉树的根节点 root ，返回其节点值 自底向上的层序遍历 。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [3,9,20,null,null,15,7]
 * 输出：[[15,7],[9,20],[3]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [1]
 * 输出：[[1]]
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
 * 树中节点数目在范围 [0, 2000] 内
 * -1000 <= Node.val <= 1000
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
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	ans := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for size := len(queue); size > 0; size = len(queue) {
		tmp := make([]int, size)
		for i := 0; i < size; i++ {
			cur := queue[i]
			tmp[i] = cur.Val
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		ans = append(ans, tmp)
		queue = queue[size:]
	}

	l := len(ans)
	for i := 0; i < l/2; i++ {
		ans[i], ans[l-i-1] = ans[l-i-1], ans[i]
	}

	return ans
}

// @lc code=end

func Test_levelOrderBottom(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want [][]int
	}{
		{"0", NewTreeByPreOrder([]any{}), nil},
		{"1", NewTreeByPreOrder([]any{1}), [][]int{{1}}},
		{"2", NewTreeByPreOrder([]any{3, 9, 20, nil, nil, 15, 7}), [][]int{{15, 7}, {9, 20}, {3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrderBottom(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrderBottom() = %v, want %v", got, tt.want)
			}
		})
	}
}
