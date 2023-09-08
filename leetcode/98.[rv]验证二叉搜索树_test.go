package leetcode

import (
	"testing"
)

func isValidBST_RV(root *TreeNode) bool {
	return isValidBSTDfs3_RV(root)
}

func isValidBSTDfs3_RV(root *TreeNode) bool {
	var valid func(r, min, max *TreeNode) bool
	valid = func(r, min, max *TreeNode) bool {
		if r == nil {
			return true
		}

		if min != nil && min.Val >= r.Val {
			return false
		}

		if max != nil && max.Val <= r.Val {
			return false
		}

		return valid(r.Left, min, r) && valid(r.Right, r, max)
	}
	return valid(root, nil, nil)
}

func isValidBSTDfs2_RV(root *TreeNode) bool {
	var valid1 func(r *TreeNode, isLeft bool) (bool, int)
	valid1 = func(r *TreeNode, isLeft bool) (bool, int) {
		if r == nil {
			return true, 0
		}

		var leftMax, rightMin int
		if r.Left != nil {
			var leftRe bool
			leftRe, leftMax = valid1(r.Left, true)
			if !leftRe {
				return false, 0
			}
		}
		if r.Right != nil {
			var rightRe bool
			rightRe, rightMin = valid1(r.Right, false)
			if !rightRe {
				return false, 0
			}
		}

		if (r.Left != nil && leftMax > r.Val) || (r.Right != nil && rightMin < r.Val) {
			return false, 0
		}
		if isLeft {
			if r.Right != nil {
				return true, r.Right.Val
			}
		} else {
			if r.Left != nil {
				return true, r.Left.Val
			}
		}
		return true, r.Val
	}
	is, _ := valid1(root, false)
	return is
}

func Test_isValidBST_RV(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want bool
	}{
		// {"1", NewTreeByPreOrder([]any{1}), true},
		// {"1", NewTreeByPreOrder([]any{1, 2, 3}), false},
		{"1", NewTreeByPreOrder([]any{2, 1, 3}), true},
		{"1", NewTreeByPreOrder([]any{5, 1, 4, nil, nil, 3, 6}), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST_RV(tt.root); got != tt.want {
				t.Errorf("isValidBST_RV() = %v, want %v", got, tt.want)
			}
		})
	}
}
