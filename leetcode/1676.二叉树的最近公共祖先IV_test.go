package leetcode

import (
	"reflect"
	"testing"
)

/* 给定一组节点（一定在二叉树中），求这一组节点的最近公共祖先。题目保证节点的数字各不相同。 */
func lowestCommonAncestor1676(root *TreeNode, descendants []*TreeNode) *TreeNode {
	offspring := make(map[*TreeNode]bool, len(descendants))
	for _, n := range descendants {
		offspring[n] = true
	}
	var count int

	var dfs func(r *TreeNode) *TreeNode
	dfs = func(r *TreeNode) *TreeNode {
		if r == nil {
			return nil
		}

		left := dfs(r.Left)
		right := dfs(r.Right)

		if offspring[r] {
			count++
			return r
		}

		if left == nil && right == nil {
			return nil
		}

		if left != nil && right != nil {
			return r
		}
		if left != nil {
			return left
		}
		return right
	}
	res := dfs(root)
	if count == len(descendants) {
		return res
	}
	return nil
}

func Test_lowestCommonAncestor1676(t *testing.T) {
	root := buildTree1676()
	type args struct {
		root        *TreeNode
		descendants []*TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"1", args{
			root:        root,
			descendants: []*TreeNode{root.Left.Left, root.Left.Right},
		}, root.Left},
		{"1", args{
			root:        root,
			descendants: []*TreeNode{root.Left.Left, root.Right.Left},
		}, root},
		{"1", args{
			root:        root,
			descendants: []*TreeNode{root.Left.Left, root.Left.Right, root.Right.Left},
		}, root},
		{"1", args{
			root:        root,
			descendants: []*TreeNode{root.Left, root.Left.Right},
		}, root.Left},
		{"1", args{
			root:        root,
			descendants: []*TreeNode{root.Left.Left, root.Left.Right, root.Left, root.Right.Left},
		}, root},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestor1676(tt.args.root, tt.args.descendants); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lowestCommonAncestor1676() = %v, want %v", got, tt.want)
			}
		})
	}
}

func buildTree1676() *TreeNode {
	// 构造树
	/*
			   1
		     /   \
		    2     3
		   / \   /
		  4   5 6
	*/
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}
	return root
}
