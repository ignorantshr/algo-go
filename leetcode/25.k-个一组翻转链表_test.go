/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 *
 * https://leetcode.cn/problems/reverse-nodes-in-k-group/description/
 *
 * algorithms
 * Hard (67.65%)
 * Likes:    2089
 * Dislikes: 0
 * Total Accepted:    480.8K
 * Total Submissions: 711K
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * 给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。
 *
 * k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
 *
 * 你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：head = [1,2,3,4,5], k = 2
 * 输出：[2,1,4,3,5]
 *
 *
 * 示例 2：
 *
 *
 *
 *
 * 输入：head = [1,2,3,4,5], k = 3
 * 输出：[3,2,1,4,5]
 *
 *
 *
 * 提示：
 *
 *
 * 链表中的节点数目为 n
 * 1 <= k <= n <= 5000
 * 0 <= Node.val <= 1000
 *
 *
 *
 *
 * 进阶：你可以设计一个只用 O(1) 额外内存空间的算法解决此问题吗？
 *
 *
 *
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
func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}

	return reverseKGroupRecurion(head, k)
}

func reverseKGroupRecurion(head *ListNode, k int) *ListNode {
	tail := head
	for n := 0; n < k; n++ {
		if tail == nil {
			return head
		}
		tail = tail.Next
	}

	nHead := reverseRange(head, tail)
	head.Next = reverseKGroupRecurion(tail, k)
	return nHead
}

// reverse [head,tail)
func reverseRange(head, tail *ListNode) *ListNode {
	var pre, nxt *ListNode
	cur := head
	for cur != tail {
		nxt = cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return pre
}

func reverseKGroup1(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy
	for n := 1; head != nil; n++ {
		if n%k == 0 {
			successor := head.Next
			head.Next = nil
			tmpHead := pre.Next
			pre.Next = reverseListIterate25(pre.Next)
			tmpHead.Next = successor

			pre = tmpHead
			head = successor
		} else {
			head = head.Next
		}
	}
	return dummy.Next
}

func reverseListIterate25(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}

// @lc code=end

func Test_reverseKGroup(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"1", args{
			NewList([]int{1}),
			1,
		}, NewList([]int{1})},
		{"1", args{
			NewList([]int{1, 2}),
			2,
		}, NewList([]int{2, 1})},
		{"1", args{
			NewList([]int{1, 2, 3, 4, 5}),
			2,
		}, NewList([]int{2, 1, 4, 3, 5})},
		{"1", args{
			NewList([]int{1, 2, 3, 4, 5}),
			3,
		}, NewList([]int{3, 2, 1, 4, 5})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseKGroup(tt.args.head, tt.args.k); !got.compare(tt.want) {
				t.Errorf("reverseKGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}
