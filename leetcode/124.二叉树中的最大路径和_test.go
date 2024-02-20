/*
 * @lc app=leetcode.cn id=124 lang=golang
 *
 * [124] 二叉树中的最大路径和
 *
 * https://leetcode.cn/problems/binary-tree-maximum-path-sum/description/
 *
 * algorithms
 * Hard (45.32%)
 * Likes:    2165
 * Dislikes: 0
 * Total Accepted:    377.6K
 * Total Submissions: 831.4K
 * Testcase Example:  '[1,2,3]'
 *
 * 二叉树中的 路径 被定义为一条节点序列，序列中每对相邻节点之间都存在一条边。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个
 * 节点，且不一定经过根节点。
 *
 * 路径和 是路径中各节点值的总和。
 *
 * 给你一个二叉树的根节点 root ，返回其 最大路径和 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [1,2,3]
 * 输出：6
 * 解释：最优路径是 2 -> 1 -> 3 ，路径和为 2 + 1 + 3 = 6
 *
 * 示例 2：
 *
 *
 * 输入：root = [-10,9,20,null,null,15,7]
 * 输出：42
 * 解释：最优路径是 15 -> 20 -> 7 ，路径和为 15 + 20 + 7 = 42
 *
 * 提示：
 *
 *
 * 树中节点数目范围是 [1, 3 * 10^4]
 * -1000 <= Node.val <= 1000
 *
 *
 */
package leetcode

import (
	"math"
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
func maxPathSum(root *TreeNode) int {
	ans := math.MinInt
	var dfs func(r *TreeNode) int
	dfs = func(r *TreeNode) int {
		if r == nil {
			return 0
		}

		// 只有在最大贡献值大于 0 时，才会选取对应子节点
		left := max(dfs(r.Left), 0)
		right := max(dfs(r.Right), 0)
		ans = max(ans, r.Val+left+right)

		// 只贡献单边
		return r.Val + max(left, right)
	}

	dfs(root)
	return ans
}

func maxPathSum1(root *TreeNode) int {
	ans := math.MinInt

	var dfs func(r *TreeNode) (int, int)
	dfs = func(r *TreeNode) (int, int) {
		if r == nil {
			return math.MinInt, math.MinInt
		}

		ll, lr := dfs(r.Left)
		rl, rr := dfs(r.Right)
		ml := max(ll, lr)
		mr := max(rl, rr)
		ans = max(ans, ml, mr) // 不包含本节点
		if r.Left != nil {
			ans = max(ans, ml+r.Val) // 包含本节点的左单边
			ml = ml + r.Val
		} else {
			ans = max(ans, r.Val)
			ml = r.Val
		}
		if r.Right != nil {
			ans = max(ans, mr+r.Val)
			mr = mr + r.Val
		} else {
			ans = max(ans, r.Val)
			mr = r.Val
		}
		if r.Right != nil && r.Left != nil {
			ans = max(ans, ml+mr-r.Val) // 包含本节点的拐弯节点
		}
		ans = max(ans, r.Val) // 只包含本节点
		ml = max(ml, r.Val)   // 放弃左子树
		mr = max(mr, r.Val)

		return ml, mr
	}

	dfs(root)

	return ans
}

// @lc code=end

func Test_maxPathSum(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		{"x.2", NewTreeByPyramid([][]any{
			{9},
			{6, -3},
			{nil, nil, -6, 2},
			{nil, nil, nil, nil, nil, nil, 2, nil},
			{nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, -6, -6, nil, nil},
		}), 16},
		{"x.1", NewTreeByPreOrder([]any{2, -1, -2}), 2},
		{"1", NewTreeByPreOrder([]any{1}), 1},
		{"-1", NewTreeByPreOrder([]any{-1}), -1},
		{"2.1", NewTreeByPreOrder([]any{1, 2, 3}), 6},
		{"2.5", NewTreeByPreOrder([]any{1, 2, 3, 4, 5}), 11},
		{"2.2", NewTreeByPreOrder([]any{-10, 9, 20, nil, nil, 15, 7}), 42},
		{"2.3", NewTreeByPreOrder([]any{-3, 2, 3}), 3},
		{"2.4", NewTreeByPreOrder([]any{-3, 2, 3, 2}), 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPathSum(tt.root); got != tt.want {
				t.Errorf("maxPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
