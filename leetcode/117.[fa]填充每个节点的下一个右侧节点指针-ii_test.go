/*
 * @lc app=leetcode.cn id=117 lang=golang
 *
 * [117] 填充每个节点的下一个右侧节点指针 II
 *
 * https://leetcode.cn/problems/populating-next-right-pointers-in-each-node-ii/description/
 *
 * algorithms
 * Medium (66.50%)
 * Likes:    828
 * Dislikes: 0
 * Total Accepted:    235.8K
 * Total Submissions: 345.1K
 * Testcase Example:  '[1,2,3,4,5,null,7]'
 *
 * 给定一个二叉树：
 *
 *
 * struct Node {
 * ⁠ int val;
 * ⁠ Node *left;
 * ⁠ Node *right;
 * ⁠ Node *next;
 * }
 *
 * 填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL 。
 *
 * 初始状态下，所有 next 指针都被设置为 NULL 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [1,2,3,4,5,null,7]
 * 输出：[1,#,2,3,#,4,5,7,#]
 * 解释：给定二叉树如图 A 所示，你的函数应该填充它的每个 next 指针，以指向其下一个右侧节点，如图 B 所示。序列化输出按层序遍历顺序（由 next
 * 指针连接），'#' 表示每层的末尾。
 *
 * 示例 2：
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
 * 树中的节点数在范围 [0, 6000] 内
 * -100 <= Node.val <= 100
 *
 *
 * 进阶：
 *
 *
 * 你只能使用常量级额外空间。
 * 使用递归解题也符合要求，本题中递归程序的隐式栈空间不计入额外空间复杂度。
 *
 *
 *
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
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	return connectBFSList(root)
	return connectDFS(root)
	return connectBFS(root)
}

func connectBFSList(root *Node) *Node {
	cur := root
	for cur != nil {
		pre := &Node{}
		zero := pre
		for cur != nil {
			if cur.Left != nil {
				pre.Next = cur.Left
				pre = cur.Left
			}
			if cur.Right != nil {
				pre.Next = cur.Right
				pre = cur.Right
			}
			cur = cur.Next
		}
		cur = zero.Next
	}

	return root
}

func connectDFS(root *Node) *Node {
	heads := make([]*Node, 0)
	var dfs func(r *Node, level int)
	dfs = func(r *Node, level int) {
		if r == nil {
			return
		}

		if len(heads) < level {
			heads = append(heads, r)
		} else {
			heads[level-1].Next = r
			heads[level-1] = r
		}
		dfs(r.Left, level+1)
		dfs(r.Right, level+1)
	}

	dfs(root, 1)
	return root
}

func connectBFS(root *Node) *Node {
	if root == nil {
		return nil
	}

	queue := make([]*Node, 0)
	queue = append(queue, root)
	for size := len(queue); size != 0; size = len(queue) {
		var pre *Node
		for i := 0; i < size; i++ {
			if pre != nil {
				pre.Next = queue[i]
			}
			pre = queue[i]
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[size:]
	}

	return root
}

// @lc code=end

func Test_connect(t *testing.T) {
	tests := []struct {
		name string
		root *Node
		want *Node
	}{
		{"0", nil, nil},
		{"1", &Node{Val: 1}, &Node{Val: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := connect(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("connect() = %v, want %v", got, tt.want)
			}
		})
	}
}
