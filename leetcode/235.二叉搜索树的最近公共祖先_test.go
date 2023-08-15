/*
 * @lc app=leetcode.cn id=235 lang=golang
 *
 * [235] 二叉搜索树的最近公共祖先
 *
 * https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-search-tree/description/
 *
 * algorithms
 * Medium (68.36%)
 * Likes:    1110
 * Dislikes: 0
 * Total Accepted:    352.2K
 * Total Submissions: 515K
 * Testcase Example:  '[6,2,8,0,4,7,9,nil,nil,3,5]\n2\n8'
 *
 * 给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。
 *
 * 百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x
 * 的深度尽可能大（一个节点也可以是它自己的祖先）。”
 *
 * 例如，给定如下二叉搜索树:  root = [6,2,8,0,4,7,9,nil,nil,3,5]
 *
 *
 *
 *
 *
 * 示例 1:
 *
 * 输入: root = [6,2,8,0,4,7,9,nil,nil,3,5], p = 2, q = 8
 * 输出: 6
 * 解释: 节点 2 和节点 8 的最近公共祖先是 6。
 *
 *
 * 示例 2:
 *
 * 输入: root = [6,2,8,0,4,7,9,nil,nil,3,5], p = 2, q = 4
 * 输出: 2
 * 解释: 节点 2 和节点 4 的最近公共祖先是 2, 因为根据定义最近公共祖先节点可以为节点本身。
 *
 *
 *
 * 说明:
 *
 *
 * 所有节点的值都是唯一的。
 * p、q 为不同节点且均存在于给定的二叉搜索树中。
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
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor235(root, p, q *TreeNode) *TreeNode {
	// 更好的条件判断
	r := root
	for {
		if p.Val < r.Val && q.Val < r.Val { // all in the left tree
			r = r.Left
		} else if p.Val > r.Val && q.Val > r.Val { // all in the right tree
			r = r.Right
		} else {
			return r
		}
	}
}

// @lc code=end

func Test_lowestCommonAncestor235(t *testing.T) {
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
			NewTreeByPreOrder([]any{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5}),
			&TreeNode{Val: 3},
			&TreeNode{Val: 5},
		}, 4},
		{"1", args{
			NewTreeByPreOrder([]any{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5}),
			&TreeNode{Val: 2},
			&TreeNode{Val: 8},
		}, 6},
		{"1", args{
			NewTreeByPreOrder([]any{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5}),
			&TreeNode{Val: 2},
			&TreeNode{Val: 4},
		}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestor(tt.args.root, tt.args.p, tt.args.q); !reflect.DeepEqual(got.Val, tt.want) {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got.Val, tt.want)
			}
		})
	}
}
