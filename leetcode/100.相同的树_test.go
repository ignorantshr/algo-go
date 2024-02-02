/*
 * @lc app=leetcode.cn id=100 lang=golang
 *
 * [100] 相同的树
 *
 * https://leetcode.cn/problems/same-tree/description/
 *
 * algorithms
 * Easy (60.36%)
 * Likes:    1123
 * Dislikes: 0
 * Total Accepted:    529.4K
 * Total Submissions: 871.1K
 * Testcase Example:  '[1,2,3]\n[1,2,3]'
 *
 * 给你两棵二叉树的根节点 p 和 q ，编写一个函数来检验这两棵树是否相同。
 *
 * 如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：p = [1,2,3], q = [1,2,3]
 * 输出：true
 *
 *
 * 示例 2：
 *
 *
 * 输入：p = [1,2], q = [1,null,2]
 * 输出：false
 *
 *
 * 示例 3：
 *
 *
 * 输入：p = [1,2,1], q = [1,1,2]
 * 输出：false
 *
 *
 *
 *
 * 提示：
 *
 *
 * 两棵树上的节点数目都在范围 [0, 100] 内
 * -10^4
 *
 *
 */
package leetcode

import "testing"

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	// BFS
	queue := make([]*TreeNode, 0)
	queue = append(queue, p)
	queue = append(queue, q)
	for len(queue) != 0 {
		q1 := queue[0]
		p1 := queue[1]
		queue = queue[2:]

		if p1 == nil && q1 == nil {
			continue
		} else if (p1 == nil || q1 == nil) || (p1.Val != q1.Val) {
			return false
		}

		queue = append(queue, p1.Left)
		queue = append(queue, q1.Left)
		queue = append(queue, p1.Right)
		queue = append(queue, q1.Right)
	}

	return true
}

func isSameTreeDFS(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if (p == nil || q == nil) || (p.Val != q.Val) {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// @lc code=end

func Test_isSameTree(t *testing.T) {
	type args struct {
		p *TreeNode
		q *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"0", args{
			NewTreeByPreOrder([]any{}),
			NewTreeByPreOrder([]any{}),
		}, true},
		{"true.1", args{
			NewTreeByPreOrder([]any{}),
			NewTreeByPreOrder([]any{}),
		}, true},
		{"true.2", args{
			NewTreeByPreOrder([]any{1}),
			NewTreeByPreOrder([]any{1}),
		}, true},
		{"true.3", args{
			NewTreeByPreOrder([]any{1, 2, 3}),
			NewTreeByPreOrder([]any{1, 2, 3}),
		}, true},
		{"true.4", args{
			NewTreeByPreOrder([]any{1, nil, 2, 3}),
			NewTreeByPreOrder([]any{1, nil, 2, 3}),
		}, true},
		{"false.1", args{
			NewTreeByPreOrder([]any{1, 2, 3}),
			NewTreeByPreOrder([]any{1, 2}),
		}, false},
		{"false.1", args{
			NewTreeByPreOrder([]any{1, 3}),
			NewTreeByPreOrder([]any{1, 2}),
		}, false},
		{"false.1", args{
			NewTreeByPreOrder([]any{3}),
			NewTreeByPreOrder([]any{2}),
		}, false},
		{"false.1", args{
			NewTreeByPreOrder([]any{}),
			NewTreeByPreOrder([]any{2}),
		}, false},
		{"false.1", args{
			NewTreeByPreOrder([]any{1, nil, 3, 5}),
			NewTreeByPreOrder([]any{1, nil, 5, 3}),
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSameTree(tt.args.p, tt.args.q); got != tt.want {
				t.Errorf("isSameTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
