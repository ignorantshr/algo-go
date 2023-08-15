/*
 * @lc app=leetcode.cn id=236 lang=golang
 *
 * [236] 二叉树的最近公共祖先
 *
 * https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/description/
 *
 * algorithms
 * Medium (69.74%)
 * Likes:    2348
 * Dislikes: 0
 * Total Accepted:    560K
 * Total Submissions: 803K
 * Testcase Example:  '[3,5,1,6,2,0,8,nil,nil,7,4]\n5\n1'
 *
 * 给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
 *
 * 百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x
 * 的深度尽可能大（一个节点也可以是它自己的祖先）。”
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [3,5,1,6,2,0,8,nil,nil,7,4], p = 5, q = 1
 * 输出：3
 * 解释：节点 5 和节点 1 的最近公共祖先是节点 3 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [3,5,1,6,2,0,8,nil,nil,7,4], p = 5, q = 4
 * 输出：5
 * 解释：节点 5 和节点 4 的最近公共祖先是节点 5 。因为根据定义最近公共祖先节点可以为节点本身。
 *
 *
 * 示例 3：
 *
 *
 * 输入：root = [1,2], p = 1, q = 2
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点数目在范围 [2, 10^5] 内。
 * -10^9
 * 所有 Node.val 互不相同 。
 * p != q
 * p 和 q 均存在于给定的二叉树中。
 *
 *
 */
package leetcode

import (
	"reflect"
	"testing"
)

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return lowestCommonAncestor2(root, p, q)
}

// 后序遍历。判断子树是否存在 p, q，根据情况进行分析判断返回值
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == q.Val || root.Val == p.Val {
		return root
	}

	leftResult := lowestCommonAncestor(root.Left, p, q)
	rightResult := lowestCommonAncestor(root.Right, p, q)

	if leftResult == nil && rightResult == nil { // case 1: 都不在
		return nil
	}
	if leftResult != nil && rightResult != nil { // case 2: 都在，自己就是答案
		return root
	}
	if leftResult != nil { // case 3: 只有一个在，返回自己
		return leftResult
	}
	return rightResult
}

func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	var dfs func(r *TreeNode, path []*TreeNode, res1, res2 *[]*TreeNode)
	dfs = func(r *TreeNode, path []*TreeNode, res1, res2 *[]*TreeNode) {
		if r == nil || (len(*res1) != 0 && len(*res2) != 0) {
			return
		}
		path = append(path, r)
		if r.Val == p.Val {
			*res1 = make([]*TreeNode, len(path))
			copy(*res1, path)
		}
		if r.Val == q.Val {
			*res2 = make([]*TreeNode, len(path))
			copy(*res2, path)
		}
		dfs(r.Left, path, res1, res2)
		dfs(r.Right, path, res1, res2)
	}

	ppath := make([]*TreeNode, 0)
	qpath := make([]*TreeNode, 0)

	dfs(root, []*TreeNode{}, &ppath, &qpath)

	var i int
	l := min(len(ppath), len(qpath))
	for i = 0; i < l; i++ {
		if ppath[i] != qpath[i] {
			return ppath[i-1]
		}
	}
	return ppath[i-1]
}

// func min(a, b int) int {
// 	if a > b {
// 		return b
// 	}
// 	return a
// }

// @lc code=end

func Test_lowestCommonAncestor(t *testing.T) {
	type args struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{
			NewTreeByPreOrder([]any{1, 2}),
			&TreeNode{Val: 2},
			&TreeNode{Val: 1},
		}, 1},
		{"1", args{
			NewTreeByPreOrder([]any{1, 2}),
			&TreeNode{Val: 1},
			&TreeNode{Val: 2},
		}, 1},
		{"1", args{
			NewTreeByPreOrder([]any{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4}),
			&TreeNode{Val: 5},
			&TreeNode{Val: 1},
		}, 3},
		{"1", args{
			NewTreeByPreOrder([]any{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4}),
			&TreeNode{Val: 5},
			&TreeNode{Val: 4},
		}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestor(tt.args.root, tt.args.p, tt.args.q); !reflect.DeepEqual(got.Val, tt.want) {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got.Val, tt.want)
			}
		})
	}
}
