/*
- @lc app=leetcode.cn id=257 lang=golang

给定一个二叉树，返回所有从根节点到叶子节点的路径。

说明: 叶子节点是指没有子节点的节点。
*/
package leetcode

import (
	"strconv"
	"strings"
	"testing"
)

// @lc code=start
func binaryTreePaths(root *TreeNode) []string {
	return binaryTreePathsIterate(root)
}

// 前序遍历
func binaryTreePathsDfs(root *TreeNode) []string {
	res := make([]string, 0)
	var path func(n *TreeNode, prefix string)
	path = func(n *TreeNode, prefix string) {
		if n == nil {
			return
		}
		if n.Left == nil && n.Right == nil {
			res = append(res, prefix+strconv.Itoa(n.Val))
			return
		}
		prefix += strconv.Itoa(n.Val) + "->"
		path(n.Left, prefix)
		path(n.Right, prefix)
	}
	path(root, "")
	return res
}

// 迭代法
func binaryTreePathsIterate(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}

	res := make([]string, 0)
	stack := make([]*TreeNode, 0)
	path := make([]string, 0)
	stack = append(stack, root)

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		path = append(path, strconv.Itoa(top.Val))

		if top.Left == nil && top.Right == nil {
			res = append(res, strings.Join(path, "->"))
			path = path[:len(path)-1]
			continue
		}
		if top.Left != nil {
			stack = append(stack, top.Left)
		}
		if top.Right != nil {
			stack = append(stack, top.Right)
		}
	}
	return res
}

// @lc code=end

func Test_binaryTreePaths(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want []string
	}{
		{"1", NewTreeByPreOrder([]any{}), []string{}},
		{"1", NewTreeByPreOrder([]any{1}), []string{"1"}},
		{"1", NewTreeByPreOrder([]any{1, 2, 3, nil, 5}), []string{"1->2->5", "1->3"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binaryTreePaths(tt.root); !equalSet(got, tt.want) {
				t.Errorf("binaryTreePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
