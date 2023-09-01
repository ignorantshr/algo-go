/*
 * @lc app=leetcode.cn id=116 lang=golang
 *
 * [116] 填充每个节点的下一个右侧节点指针
 *
 * https://leetcode.cn/problems/populating-next-right-pointers-in-each-node/description/
 *
 * algorithms
 * Medium (72.60%)
 * Likes:    1017
 * Dislikes: 0
 * Total Accepted:    356.5K
 * Total Submissions: 491K
 * Testcase Example:  '[1,2,3,4,5,6,7]'
 *
 * 给定一个 完美二叉树 ，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下：
 *
 *
 * struct Node {
 * ⁠ int val;
 * ⁠ Node *left;
 * ⁠ Node *right;
 * ⁠ Node *next;
 * }
 *
 * 填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。
 *
 * 初始状态下，所有 next 指针都被设置为 NULL。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：root = [1,2,3,4,5,6,7]
 * 输出：[1,#,2,3,#,4,5,6,7,#]
 * 解释：给定二叉树如图 A 所示，你的函数应该填充它的每个 next 指针，以指向其下一个右侧节点，如图 B
 * 所示。序列化的输出按层序遍历排列，同一层节点由 next 指针连接，'#' 标志着每一层的结束。
 *
 *
 *
 *
 * 示例 2:
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
 * 树中节点的数量在 [0, 2^12 - 1] 范围内
 * -1000 <= node.val <= 1000
 *
 *
 *
 *
 * 进阶：
 *
 *
 * 你只能使用常量级额外空间。
 * 使用递归解题也符合要求，本题中递归程序占用的栈空间不算做额外的空间复杂度。
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
	if root == nil {
		return nil
	}
	connect2(root.Left, root.Right)
	return root
}

// 遍历三叉树
func connect2(n1, n2 *Node) {
	if n1 == nil || n2 == nil {
		return
	}
	n1.Next = n2
	connect2(n1.Left, n1.Right)
	connect2(n2.Left, n2.Right)
	connect2(n1.Right, n2.Left)
}

func connect1(root *Node) *Node {
	nodes := make(map[int][]*Node)
	connectRecurrence(root, 0, nodes)
	for _, ns := range nodes {
		for i := 1; i < len(ns); i++ {
			ns[i-1].Next = ns[i]
		}
	}
	return root
}

func connectRecurrence(n *Node, depth int, nodes map[int][]*Node) {
	if n == nil {
		return
	}

	nodes[depth] = append(nodes[depth], n)
	connectRecurrence(n.Left, depth+1, nodes)
	connectRecurrence(n.Right, depth+1, nodes)
}

// @lc code=end

func Test_connect(t *testing.T) {
	tests := []struct {
		name string
		root *Node
		want *Node
	}{
		{"1", &Node{Val: 1, Left: &Node{Val: 2, Left: &Node{Val: 4}, Right: &Node{Val: 5}}, Right: &Node{Val: 3, Left: &Node{Val: 6}, Right: &Node{Val: 7}}}, &Node{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := connect(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("connect() = %v, want %v", got, tt.want)
			}
		})
	}
}
