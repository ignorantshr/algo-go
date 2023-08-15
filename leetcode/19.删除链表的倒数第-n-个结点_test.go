/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第 N 个结点
 *
 * https://leetcode.cn/problems/remove-nth-node-from-end-of-list/description/
 *
 * algorithms
 * Medium (45.67%)
 * Likes:    2603
 * Dislikes: 0
 * Total Accepted:    1.2M
 * Total Submissions: 2.6M
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：head = [1,2,3,4,5], n = 2
 * 输出：[1,2,3,5]
 *
 *
 * 示例 2：
 *
 *
 * 输入：head = [1], n = 1
 * 输出：[]
 *
 *
 * 示例 3：
 *
 *
 * 输入：head = [1,2], n = 1
 * 输出：[1]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 链表中结点的数目为 sz
 * 1 <= sz <= 30
 * 0 <= Node.val <= 100
 * 1 <= n <= sz
 *
 *
 *
 *
 * 进阶：你能尝试使用一趟扫描实现吗？
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
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	return removeNthFromEnd1(head, n)
}

func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	// 指针
	// 可以使用虚拟头节点简化边界处理
	fast := head
	var pre *ListNode
	for n > 0 {
		fast = fast.Next
		n--
	}
	for fast != nil {
		fast = fast.Next
		if pre == nil {
			pre = head
		} else {
			pre = pre.Next
		}
	}

	// 因为 1 <= n <= sz，所以 slow 为空时删掉的是根节点
	if pre == nil {
		return head.Next
	}

	pre.Next = pre.Next.Next
	return head
}

// @lc code=end

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"1", args{NewList([]int{1}), 1}, NewList([]int{})},
		{"1", args{NewList([]int{1, 2}), 1}, NewList([]int{1})},
		{"1", args{NewList([]int{1, 2}), 2}, NewList([]int{2})},
		{"1", args{NewList([]int{1, 2, 3, 4, 5}), 2}, NewList([]int{1, 2, 3, 5})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); !tt.want.compare(got) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
