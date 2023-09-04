package leetcode

import (
	"testing"
)

func constructMaximumBinaryTreeRV(nums []int) *TreeNode {
	getMaxIdx := func(left, right int) int {
		max := left
		for left <= right {
			if nums[left] > nums[max] {
				max = left
			}
			left++
		}
		return max
	}

	// [left,right]
	var build func(left, right int) *TreeNode
	build = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}

		rootIdx := getMaxIdx(left, right)
		root := &TreeNode{Val: nums[rootIdx]}
		root.Left = build(left, rootIdx-1)
		root.Right = build(rootIdx+1, right)
		return root
	}
	return build(0, len(nums)-1)
}

func Test_constructMaximumBinaryTreeRV(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want *TreeNode
	}{
		{"1", []int{3}, NewTreeByPreOrder([]any{3})},
		{"1", []int{1, 3, 2}, NewTreeByPreOrder([]any{3, 1, 2})},
		{"1", []int{3, 2, 1, 6, 0, 5}, NewTreeByPreOrder([]any{})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := constructMaximumBinaryTreeRV(tt.nums); got.equal(tt.want) {
				t.Errorf("constructMaximumBinaryTreeRV() = %v, want %v", got, tt.want)
			}
		})
	}
}
