/*
 * @lc app=leetcode.cn id=103 lang=golang
 *
 * [103] 二叉树的锯齿形层序遍历
 *
 * https://leetcode.cn/problems/binary-tree-zigzag-level-order-traversal/description/
 *
 * algorithms
 * Medium (57.78%)
 * Likes:    845
 * Dislikes: 0
 * Total Accepted:    342.1K
 * Total Submissions: 589.6K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给你二叉树的根节点 root ，返回其节点值的 锯齿形层序遍历 。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [3,9,20,null,null,15,7]
 * 输出：[[3],[20,9],[15,7]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [1]
 * 输出：[[1]]
 *
 *
 * 示例 3：
 *
 *
 * 输入：root = []
 * 输出：[]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点数目在范围 [0, 2000] 内
 * -100 <= Node.val <= 100
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
func zigzagLevelOrder(root *TreeNode) [][]int {
	queue := make([]*TreeNode, 0)
	if root != nil {
		queue = append(queue, root)
	}
	ans := make([][]int, 0)
	level := 0

	for len(queue) != 0 {
		size := len(queue)
		ans = append(ans, make([]int, 0))

		for i := 0; i < size; i++ {
			if level&1 == 1 {
				ans[level] = append(ans[level], queue[size-i-1].Val)
			} else {
				ans[level] = append(ans[level], queue[i].Val)
			}
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		// if level&1 == 1 {
		// 	l := len(ans[level])
		// 	for i := 0; i < l>>1; i++ {
		// 		ans[level][i], ans[level][l-i-1] = ans[level][l-i-1], ans[level][i]
		// 	}
		// }
		queue = queue[size:]
		level++
	}

	return ans
}

// @lc code=end

func Test_zigzagLevelOrder(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want [][]int
	}{
		{"0", NewTreeByPreOrder([]any{}), [][]int{}},
		{"1", NewTreeByPreOrder([]any{1}), [][]int{{1}}},
		{"2", NewTreeByPreOrder([]any{3, 9, 20, nil, nil, 15, 7}), [][]int{{3}, {20, 9}, {15, 7}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := zigzagLevelOrder(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("zigzagLevelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
