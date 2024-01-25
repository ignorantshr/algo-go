/*
 * @lc app=leetcode.cn id=83 lang=golang
 *
 * [83] 删除排序链表中的重复元素
 *
 * https://leetcode.cn/problems/remove-duplicates-from-sorted-list/description/
 *
 * algorithms
 * Easy (53.10%)
 * Likes:    1018
 * Dislikes: 0
 * Total Accepted:    594.4K
 * Total Submissions: 1.1M
 * Testcase Example:  '[1,1,2]'
 *
 * 给定一个已排序的链表的头 head ， 删除所有重复的元素，使每个元素只出现一次 。返回 已排序的链表 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：head = [1,1,2]
 * 输出：[1,2]
 *
 *
 * 示例 2：
 *
 *
 * 输入：head = [1,1,2,3,3]
 * 输出：[1,2,3]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 链表中节点数目在范围 [0, 300] 内
 * -100 <= Node.val <= 100
 * 题目数据保证链表已经按升序 排列
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
func deleteDuplicates83(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	cur := head
	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

func deleteDuplicates83831(head *ListNode) *ListNode {
	// 快慢指针
	if head == nil {
		return nil
	}

	slow, fast := head, head.Next
	for fast != nil {
		if slow.Val != fast.Val {
			// slow = slow.Next
			// slow.Val = fast.Val
			slow.Next = fast
			slow = slow.Next
		}
		fast = fast.Next
	}
	slow.Next = nil
	return head
}

// @lc code=end

func Test_deleteDuplicates83(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		want *ListNode
	}{
		{"1", nil, nil},
		{"1", NewList([]int{}), nil},
		{"1", NewList([]int{1}), NewList([]int{1})},
		{"2", NewList([]int{1, 1}), NewList([]int{1})},
		{"2", NewList([]int{1, 2}), NewList([]int{1, 2})},
		{"2", NewList([]int{1, 1, 2}), NewList([]int{1, 2})},
		{"2", NewList([]int{1, 1, 2}), NewList([]int{1, 2})},
		{"2", NewList([]int{1, 1, 2, 3, 3}), NewList([]int{1, 2, 3})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates83(tt.head); !got.compare(tt.want) {
				t.Errorf("deleteDuplicates83() = %v, want %v", got, tt.want)
			}
		})
	}
}
