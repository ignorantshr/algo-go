/*
 * @lc app=leetcode.cn id=889 lang=golang
 *
 * [889] 根据前序和后序遍历构造二叉树
 *
 * https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-postorder-traversal/description/
 *
 * algorithms
 * Medium (68.02%)
 * Likes:    316
 * Dislikes: 0
 * Total Accepted:    39.4K
 * Total Submissions: 57.9K
 * Testcase Example:  '[1,2,4,5,3,6,7]\n[4,5,2,6,7,3,1]'
 *
 * 给定两个整数数组，preorder 和 postorder ，其中 preorder 是一个具有 无重复 值的二叉树的前序遍历，postorder
 * 是同一棵树的后序遍历，重构并返回二叉树。
 *
 * 如果存在多个答案，您可以返回其中 任何 一个。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：preorder = [1,2,4,5,3,6,7], postorder = [4,5,2,6,7,3,1]
 * 输出：[1,2,3,4,5,6,7]
 *
 *
 * 示例 2:
 *
 *
 * 输入: preorder = [1], postorder = [1]
 * 输出: [1]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= preorder.length <= 30
 * 1 <= preorder[i] <= preorder.length
 * preorder 中所有值都 不同
 * postorder.length == preorder.length
 * 1 <= postorder[i] <= postorder.length
 * postorder 中所有值都 不同
 * 保证 preorder 和 postorder 是同一棵二叉树的前序遍历和后序遍历
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
func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	// preorder  [root,left,right]
	// postorder [left,right,root]
	postIdxes := make(map[int]int)
	for i, v := range postorder {
		postIdxes[v] = i
	}
	var build func(preStart, preEnd, postStart, postEnd int) *TreeNode
	build = func(preStart int, preEnd int, postStart int, postEnd int) *TreeNode {
		if preStart > preEnd {
			return nil
		}
		if preStart == preEnd {
			return &TreeNode{Val: preorder[preStart]}
		}
		r := &TreeNode{Val: preorder[preStart]}
		idx := postIdxes[preorder[preStart+1]] // idx 确定的是左子树根节点postorder的位置
		leftSize := idx - postStart + 1        // 确定左子树的节点数量
		r.Left = build(preStart+1, preStart+leftSize, postStart, idx)
		r.Right = build(preStart+leftSize+1, preEnd, idx+1, postEnd-1)
		return r
	}
	return build(0, len(preorder)-1, 0, len(postorder)-1)
}

// @lc code=end
func buildTree889Check(preorder []int, inorder []int) []any {
	t := constructFromPrePost(preorder, inorder)
	return t.bfsPrefix()
}

func Test_constructFromPrePost(t *testing.T) {
	tests := []struct {
		name      string
		preorder  []int
		postorder []int
		want      []any
	}{
		{"1", []int{1, 2, 4, 5, 3, 6, 7}, []int{4, 5, 2, 6, 7, 3, 1}, []any{1, 2, 4, 5, 3, 6, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildTree889Check(tt.preorder, tt.postorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("constructFromPrePost() = %v, want %v", got, tt.want)
			}
		})
	}
}
