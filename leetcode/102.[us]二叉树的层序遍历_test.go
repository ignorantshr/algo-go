/*
- @lc app=leetcode.cn id=102 lang=golang

给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
*/
package leetcode

import (
	"reflect"
	"testing"
)

// @lc code=start
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	res := make([][]int, 0)
	cur := root
	queue := make([]*TreeNode, 0)
	queue = append(queue, cur)

	for len(queue) > 0 {
		size := len(queue)
		tmp := make([]int, 0)
		for i := 0; i < size; i++ {
			tmp = append(tmp, queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		res = append(res, tmp)
		queue = queue[size:]
	}
	return res
}

// @lc code=end

func Test_levelOrder(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want [][]int
	}{
		{"1", NewTreeByPreOrder([]any{1, 2, 3, 4, 5, nil, 8, 6, 7}), [][]int{{1}, {2, 3}, {4, 5, 8}, {6, 7}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrder(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
