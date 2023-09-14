package leetcode

import (
	"testing"
)

func convertBST_RV(root *TreeNode) *TreeNode {
	var sum int
	var gather func(r *TreeNode)
	gather = func(r *TreeNode) {
		if r == nil {
			return
		}

		gather(r.Right)
		r.Val += sum
		sum = r.Val
		gather(r.Left)
	}

	gather(root)
	return root
}

func Test_convertBST_RV(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want *TreeNode
	}{
		{"1", NewTreeByPreOrder([]any{0, nil, 1}), NewTreeByPyramid([][]any{{1}, {nil, 1}})},
		{"1", NewTreeByPreOrder([]any{1, 0, 2}), NewTreeByPyramid([][]any{{3}, {3, 2}})},
		{"1", NewTreeByPreOrder([]any{3, 2, 4, 1}), NewTreeByPyramid([][]any{{7}, {9, 4}, {10}})},
		{"1",
			NewTreeByPyramid([][]any{{4}, {1, 6}, {0, 2, 5, 7}, {nil, nil, nil, 3, nil, nil, nil, 8}}),
			NewTreeByPyramid([][]any{{30}, {36, 21}, {36, 35, 26, 15}, {nil, nil, nil, 33, nil, nil, nil, 8}}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertBST_RV(tt.root); !got.equal(tt.want) {
				t.Errorf("convertBST_RV() = %v, want %v", got.bfsPrefix(), tt.want.bfsPrefix())
			}
		})
	}
}
