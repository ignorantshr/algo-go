package leetcode

import (
	"reflect"
	"testing"
)

func buildTree105RV(preorder []int, inorder []int) *TreeNode {
	inOrderIdxes := make(map[int]int, len(inorder))
	for i, v := range inorder {
		inOrderIdxes[v] = i
	}

	// preorder  [root,left,right]
	// inorder   [left,root,right]
	var build func(prelo, prehi, inlo, inhi int) *TreeNode
	build = func(prelo, prehi, inlo, inhi int) *TreeNode {
		if inhi < inlo {
			return nil
		}

		root := &TreeNode{Val: preorder[prelo]}
		middleIdx := inOrderIdxes[root.Val]
		long := middleIdx - inlo
		root.Left = build(prelo+1, prelo+long, inlo, middleIdx-1)
		root.Right = build(prelo+long+1, prehi, middleIdx+1, inhi)
		return root
	}
	return build(0, len(preorder)-1, 0, len(inorder)-1)
}

func buildTree105RVCheck(preorder []int, inorder []int) []any {
	t := buildTree105RV(preorder, inorder)
	return t.bfsPrefix()
}

func Test_buildTree105RV(t *testing.T) {
	tests := []struct {
		name     string
		preorder []int
		inorder  []int
		want     []any
	}{
		{"1", []int{1, 2, 3}, []int{3, 2, 1}, []any{1, 2, nil, 3}},
		{"1", []int{1, 2}, []int{1, 2}, []any{1, nil, 2}},
		{"1", []int{3, 9, 20}, []int{9, 3, 20}, []any{3, 9, 20}},
		{"1", []int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}, []any{3, 9, 20, nil, nil, 15, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildTree105RVCheck(tt.preorder, tt.inorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildTree105RV() = %v, want %v", got, tt.want)
			}
		})
	}
}
