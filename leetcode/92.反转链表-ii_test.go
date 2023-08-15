/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
 *
 * https://leetcode.cn/problems/reverse-linked-list-ii/description/
 *
 * algorithms
 * Medium (55.69%)
 * Likes:    1596
 * Dislikes: 0
 * Total Accepted:    413.2K
 * Total Submissions: 741.9K
 * Testcase Example:  '[1,2,3,4,5]\n2\n4'
 *
 * 给你单链表的头指针 head 和两个整数 left 和 right ，其中 left  。请你反转从位置 left 到位置 right 的链表节点，返回
 * 反转后的链表 。
 *
 *
 * 示例 1：
 *
 *
 * 输入：head = [1,2,3,4,5], left = 2, right = 4
 * 输出：[1,4,3,2,5]
 *
 *
 * 示例 2：
 *
 *
 * 输入：head = [5], left = 1, right = 1
 * 输出：[5]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 链表中节点数目为 n
 * 1 <= n <= 500
 * -500 <= Node.val <= 500
 * 1 <= left <= right <= n
 *
 *
 *
 *
 * 进阶： 你可以使用一趟扫描完成反转吗？
 *
 */
package leetcode

import (
	"testing"
)

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	return reverseBetweenNeedle(head, left, right)
}

func reverseBetweenNeedle(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head

	pre := dummy
	for i := 1; i < left; i++ {
		pre = pre.Next
	}

	cur := pre.Next
	for i := left; i < right; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return dummy.Next
}

func reverseBetweenScan(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head

	pre := dummy
	for i := 1; i < left; i++ {
		pre = pre.Next
	}

	rightNode := pre.Next
	for i := left; i < right; i++ {
		rightNode = rightNode.Next
	}

	successor := rightNode.Next
	leftNode := pre.Next
	pre.Next = nil
	rightNode.Next = nil

	node := reverseListIterate1(leftNode)
	pre.Next = node
	leftNode.Next = successor

	return dummy.Next
}

func reverseListIterate1(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}

// 递归，重点
func reverseBetweenRecursive(head *ListNode, left int, right int) *ListNode {
	// base case
	if left == 1 {
		return reverseN(head, right)
	}

	// 前进到反转的起点触发 base case
	head.Next = reverseBetween(head.Next, left-1, right-1)
	return head
}

var successor92 *ListNode

// 反转前n个节点
func reverseN(head *ListNode, n int) *ListNode {
	if n == 1 {
		successor92 = head.Next
		return head
	}
	// 找到最后一个节点
	last := reverseN(head.Next, n-1)
	// 然后逐个反转前面的节点
	head.Next.Next = head
	head.Next = successor92
	return last
}

// @lc code=end

func Test_reverseBetween(t *testing.T) {
	type args struct {
		head  *ListNode
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"1", args{
			NewList([]int{1}),
			1, 1,
		}, NewList([]int{1})},
		{"1", args{
			NewList([]int{1, 2}),
			2, 2,
		}, NewList([]int{1, 2})},
		{"2", args{
			NewList([]int{1, 2, 3, 4}),
			2, 3,
		}, NewList([]int{1, 3, 2, 4})},
		{"2", args{
			NewList([]int{1, 2, 3, 4}),
			1, 4,
		}, NewList([]int{4, 3, 2, 1})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBetween(tt.args.head, tt.args.left, tt.args.right); !got.compare(tt.want) {
				t.Errorf("reverseBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}
