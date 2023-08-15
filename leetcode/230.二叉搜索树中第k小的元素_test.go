/*
 * @lc app=leetcode.cn id=230 lang=golang
 *
 * [230] 二叉搜索树中第K小的元素
 *
 * https://leetcode.cn/problems/kth-smallest-element-in-a-bst/description/
 *
 * algorithms
 * Medium (76.04%)
 * Likes:    752
 * Dislikes: 0
 * Total Accepted:    274.6K
 * Total Submissions: 361.2K
 * Testcase Example:  '[3,1,4,null,2]\n1'
 *
 * 给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 个最小元素（从 1 开始计数）。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [3,1,4,null,2], k = 1
 * 输出：1
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [5,3,6,2,4,null,null,1], k = 3
 * 输出：3
 *
 *
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中的节点数为 n 。
 * 1
 * 0
 *
 *
 *
 *
 * 进阶：如果二叉搜索树经常被修改（插入/删除操作）并且你需要频繁地查找第 k 小的值，你将如何优化算法？
 *
 */
package leetcode

import (
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
func kthSmallest(root *TreeNode, k int) int {
	return kthSmallest1(root, k)
}

func kthSmallest1(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}

	var res int
	var traverse func(r *TreeNode, n *int)
	traverse = func(r *TreeNode, n *int) {
		if r == nil || *n < 0 {
			return
		}

		traverse(r.Left, n)
		*n--
		if *n == 0 {
			res = r.Val
			return
		}
		traverse(r.Right, n)
	}
	traverse(root, &k)
	return res
}

type tree230 struct {
	root      *TreeNode
	nodeCount map[*TreeNode]int
}

func (t *tree230) countNodes(n *TreeNode) int {
	if n == nil {
		return 0
	}
	t.nodeCount[n] = t.countNodes(n.Left) + t.countNodes(n.Right) + 1
	return t.nodeCount[n]
}

func kthSmallest2(root *TreeNode, k int) int {
	t := &tree230{root, make(map[*TreeNode]int)}
	t.countNodes(t.root)
	n := root

	for {
		pre := t.nodeCount[n.Left]
		if pre < k-1 {
			n = n.Right
			k -= pre + 1 // 扣除
		} else if pre == k-1 {
			return n.Val
		} else {
			n = n.Left
		}
	}
}

// @lc code=end

func Test_kthSmallest(t *testing.T) {
	type args struct {
		root *TreeNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{NewTreeByPreOrder([]any{1, 2, 3}), 3}, 3},
		{"1", args{NewTreeByPreOrder([]any{1, nil, 2}), 1}, 1},
		{"1", args{NewTreeByPreOrder([]any{1, nil, 2}), 2}, 2},
		{"1", args{NewTreeByPreOrder([]any{3, 1, 4, nil, 2}), 1}, 1},
		{"2", args{NewTreeByPreOrder([]any{5, 3, 6, 2, 4, nil, nil, 1}), 3}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthSmallest(tt.args.root, tt.args.k); got != tt.want {
				t.Errorf("kthSmallest() = %v, want %v", got, tt.want)
			}
		})
	}
}
