package leetcode

import (
	"reflect"
	"testing"
)

/*
		1
	  2   3
	4  5 6 7

	inorder：  4，2，5，1，6，3，7
	postorder：4，5，2，6，7，3，1
*/

func buildTree106RV(inorder []int, postorder []int) *TreeNode {
	// [左中右]
	// [左右中]

	if len(postorder) == 0 {
		return nil
	}
	rootV := postorder[len(postorder)-1]
	root := &TreeNode{Val: postorder[len(postorder)-1]}
	var middleIdx int
	for ; middleIdx < len(inorder); middleIdx++ {
		if inorder[middleIdx] == rootV {
			break
		}
	}
	root.Left = buildTree106RV(inorder[:middleIdx], postorder[:middleIdx])
	root.Right = buildTree106RV(inorder[middleIdx+1:], postorder[middleIdx:len(postorder)-1])
	return root
}

func buildTree106RVCheck(preorder []int, inorder []int) []any {
	t := buildTree106RV(preorder, inorder)
	return t.bfsPrefix()
}

func Test_buildTree106RV(t *testing.T) {
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
			if got := buildTree106RVCheck(tt.inorder, tt.postorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildTree106RV() = %v, want %v", got, tt.want)
			}
		})
	}
}
