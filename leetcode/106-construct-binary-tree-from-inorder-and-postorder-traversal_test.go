/*
 * @lc app=leetcode.cn id=106 lang=golang
 *
 * [106] 从中序与后序遍历序列构造二叉树
 *
 * https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal/description/
 *
 * algorithms
 * Medium (72.09%)
 * Likes:    1046
 * Dislikes: 0
 * Total Accepted:    292K
 * Total Submissions: 405.1K
 * Testcase Example:  '[9,3,15,20,7]\n[9,15,7,20,3]'
 *
 * 给定两个整数数组 inorder 和 postorder ，其中 inorder 是二叉树的中序遍历， postorder
 * 是同一棵树的后序遍历，请你构造并返回这颗 二叉树 。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入：inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
 * 输出：[3,9,20,null,null,15,7]
 *
 *
 * 示例 2:
 *
 *
 * 输入：inorder = [-1], postorder = [-1]
 * 输出：[-1]
 *
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= inorder.length <= 3000
 * postorder.length == inorder.length
 * -3000 <= inorder[i], postorder[i] <= 3000
 * inorder 和 postorder 都由 不同 的值组成
 * postorder 中每一个值都在 inorder 中
 * inorder 保证是树的中序遍历
 * postorder 保证是树的后序遍历
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
// func buildTree106(inorder []int, postorder []int) *TreeNode {
func buildTree(inorder []int, postorder []int) *TreeNode {
	inOrderIdxes := make(map[int]int, len(inorder))
	for i, v := range inorder {
		inOrderIdxes[v] = i
	}
	// inorder 	 [left,root,right]
	// postorder [left,right,root]
	var build func(inlo, inhi, postlo, posthi int) *TreeNode
	build = func(inlo, inhi, postlo, posthi int) *TreeNode {
		if postlo > posthi {
			return nil
		}
		r := &TreeNode{Val: postorder[posthi]}
		idx := inOrderIdxes[postorder[posthi]]
		rightSize := inhi - idx
		r.Left = build(inlo, idx-1, postlo, posthi-rightSize-1)
		r.Right = build(idx+1, inhi, posthi-rightSize, posthi-1)
		return r
	}
	return build(0, len(inorder)-1, 0, len(postorder)-1)
}

// @lc code=end

func buildTree106Check(preorder []int, inorder []int) []any {
	t := buildTree(preorder, inorder)
	return t.bfsPrefix()
}

func Test_buildTree(t *testing.T) {
	tests := []struct {
		name      string
		inorder   []int
		postorder []int
		want      []any
	}{
		{"1", []int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3}, []any{3, 9, 20, nil, nil, 15, 7}},
		{"1", []int{-1}, []int{-1}, []any{-1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildTree106Check(tt.inorder, tt.postorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
