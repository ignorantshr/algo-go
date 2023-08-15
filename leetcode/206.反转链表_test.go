/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 *
 * https://leetcode.cn/problems/reverse-linked-list/description/
 *
 * algorithms
 * Easy (73.53%)
 * Likes:    3238
 * Dislikes: 0
 * Total Accepted:    1.5M
 * Total Submissions: 2.1M
 * Testcase Example:  '[1,2,3,4,5]'
 *
 * 给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：head = [1,2,3,4,5]
 * 输出：[5,4,3,2,1]
 *
 *
 * 示例 2：
 *
 *
 * 输入：head = [1,2]
 * 输出：[2,1]
 *
 *
 * 示例 3：
 *
 *
 * 输入：head = []
 * 输出：[]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 链表中节点的数目范围是 [0, 5000]
 * -5000
 *
 *
 *
 *
 * 进阶：链表可以选用迭代或递归方式完成反转。你能否用两种方法解决这道题？
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
func reverseList(head *ListNode) *ListNode {
	return reverseListDfs3(head)
}

func reverseListDfs3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	last := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}

func reverseListDfs2(head *ListNode) *ListNode {
	dummy := &ListNode{}
	var dfs func(pre, cur *ListNode)
	dfs = func(pre, cur *ListNode) {
		if cur == nil {
			dummy.Next = pre
			return
		}
		dfs(cur, cur.Next)
		cur.Next = pre
	}
	dfs(nil, head)
	return dummy.Next
}

func reverseListDfs1(head *ListNode) *ListNode {
	var dfs func(pre, cur *ListNode) *ListNode
	dfs = func(pre, cur *ListNode) *ListNode {
		if cur == nil {
			return pre
		}
		next := cur.Next
		cur.Next = pre
		return dfs(cur, next)
	}
	return dfs(nil, head)
}

func reverseListIterate(head *ListNode) *ListNode {
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

func Test_reverseList(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		want *ListNode
	}{
		{"1", NewList([]int{}), NewList([]int{})},
		{"1", NewList([]int{1}), NewList([]int{1})},
		{"2", NewList([]int{1, 2}), NewList([]int{2, 1})},
		{"2", NewList([]int{1, 2, 3, 4, 5}), NewList([]int{5, 4, 3, 2, 1})},
		{"2", NewList([]int{1, 3, 7, 2}), NewList([]int{2, 7, 3, 1})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.head); !got.compare(tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
