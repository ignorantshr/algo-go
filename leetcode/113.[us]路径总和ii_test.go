/*
- @lc app=leetcode.cn id=113 lang=golang

给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。
*/
package leetcode

import (
	"reflect"
	"testing"
)

// @lc code=start
func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return [][]int{}
	}

	res := make([][]int, 0)
	var traversal func(n *TreeNode, pathSum int, path []int)
	traversal = func(n *TreeNode, pathSum int, path []int) {
		pathSum += n.Val
		if n.Left == nil && n.Right == nil {
			if pathSum == sum {
				tmp := make([]int, len(path))
				copy(tmp, path)
				tmp = append(tmp, n.Val)
				res = append(res, tmp)
			}
			return
		}

		if n.Left != nil {
			path = append(path, n.Val)
			traversal(n.Left, pathSum, path)
			path = path[:len(path)-1]
		}
		if n.Right != nil {
			path = append(path, n.Val)
			traversal(n.Right, pathSum, path)
			// path = path[:len(path)-1]
		}
	}
	traversal(root, 0, []int{})
	return res
}

// @lc code=end

func Test_pathSum(t *testing.T) {
	type args struct {
		root *TreeNode
		sum  int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"0", args{
			NewTreeByPreOrder([]any{}),
			0,
		}, [][]int{}},
		{"1", args{
			NewTreeByPreOrder([]any{1}),
			1,
		}, [][]int{{1}}},
		{"1", args{
			NewTreeByPreOrder([]any{1, 2, 3, nil, 4, 3}),
			7,
		}, [][]int{{1, 2, 4}, {1, 3, 3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pathSum(tt.args.root, tt.args.sum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
