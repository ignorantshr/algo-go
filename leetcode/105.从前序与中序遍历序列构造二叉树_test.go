/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
 *
 * https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/
 *
 * algorithms
 * Medium (71.26%)
 * Likes:    2010
 * Dislikes: 0
 * Total Accepted:    508.3K
 * Total Submissions: 713.4K
 * Testcase Example:  '[3,9,20,15,7]\n[9,3,15,20,7]'
 *
 * 给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历， inorder
 * 是同一棵树的中序遍历，请构造二叉树并返回其根节点。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
 * 输出: [3,9,20,null,null,15,7]
 *
 *
 * 示例 2:
 *
 *
 * 输入: preorder = [-1], inorder = [-1]
 * 输出: [-1]
 *
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= preorder.length <= 3000
 * inorder.length == preorder.length
 * -3000 <= preorder[i], inorder[i] <= 3000
 * preorder 和 inorder 均 无重复 元素
 * inorder 均出现在 preorder
 * preorder 保证 为二叉树的前序遍历序列
 * inorder 保证 为二叉树的中序遍历序列
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

// func buildTree(preorder []int, inorder []int) *TreeNode {
func buildTree105(preorder []int, inorder []int) *TreeNode {
	inOrderIdxes := make(map[int]int, len(inorder))
	for i, v := range inorder {
		inOrderIdxes[v] = i
	}
	var build func(preLo, preHi int, inLo, inHi int) *TreeNode
	build = func(preLo, preHi int, inLo, inHi int) *TreeNode {
		if preLo > preHi {
			return nil
		}

		r := &TreeNode{Val: preorder[preLo]}
		idx := inOrderIdxes[preorder[preLo]]
		// preorder [preLo,...,preLo+leftSize,...,preHi]
		// inorder  [inLo,...,idx,...,inHi]
		leftSize := idx - inLo // left tree size
		r.Left = build(preLo+1, preLo+leftSize, inLo, idx-1)
		r.Right = build(preLo+leftSize+1, preHi, idx+1, inHi)
		return r
	}
	return build(0, len(preorder)-1, 0, len(inorder)-1)
}

// @lc code=end

func buildTree105Check(preorder []int, inorder []int) []any {
	t := buildTree105(preorder, inorder)
	return t.bfsPrefix()
}

func Test_buildTree105(t *testing.T) {
	tests := []struct {
		name     string
		preorder []int
		inorder  []int
		want     []any
	}{
		{"1", []int{1, 2, 3}, []int{3, 2, 1}, []any{1, 2, nil, 3, nil}},
		{"1", []int{1, 2}, []int{1, 2}, []any{1, nil, 2}},
		{"1", []int{3, 9, 20}, []int{9, 3, 20}, []any{3, 9, 20}},
		{"1", []int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}, []any{3, 9, 20, nil, nil, 15, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildTree105Check(tt.preorder, tt.inorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
