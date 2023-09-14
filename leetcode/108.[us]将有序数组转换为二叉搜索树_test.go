/*
- @lc app=leetcode.cn id=108 lang=golang

将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。

本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。
*/
package leetcode

import (
	"testing"
)

// @lc code=start
func sortedArrayToBST(nums []int) *TreeNode {
	// 二分取根
	var toBST func(arr []int, l, r int) *TreeNode
	toBST = func(arr []int, l, r int) *TreeNode {
		if l > r {
			return nil
		}

		mid := l + (r-l)/2
		root := &TreeNode{Val: arr[mid]}
		root.Left = toBST(arr, l, mid-1)
		root.Right = toBST(arr, mid+1, r)
		return root
	}
	return toBST(nums, 0, len(nums)-1)
}

// @lc code=end

func Test_sortedArrayToBST(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want *TreeNode
	}{
		{"0", []int{0}, NewTreeByPyramid([][]any{
			{0},
		})},
		{"odd", []int{0, 1, 2, 3, 4, 5, 6}, NewTreeByPyramid([][]any{
			{3},
			{1, 5},
			{0, 2, 4, 6},
		})},
		{"event", []int{0, 1, 2, 3, 4, 5}, NewTreeByPyramid([][]any{
			{2},
			{0, 4},
			{nil, 1, 3, 5},
		})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedArrayToBST(tt.nums); !got.equal(tt.want) {
				t.Errorf("sortedArrayToBST() = %v, want %v", got.bfsPrefix(), tt.want.bfsPrefix())
			}
		})
	}
}
