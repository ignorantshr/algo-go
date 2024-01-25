/*
 * @lc app=leetcode.cn id=82 lang=golang
 *
 * [82] 删除排序链表中的重复元素 II
 *
 * https://leetcode.cn/problems/remove-duplicates-from-sorted-list-ii/description/
 *
 * algorithms
 * Medium (53.60%)
 * Likes:    1266
 * Dislikes: 0
 * Total Accepted:    406.3K
 * Total Submissions: 751K
 * Testcase Example:  '[1,2,3,3,4,4,5]'
 *
 * 给定一个已排序的链表的头 head ， 删除原始链表中所有重复数字的节点，只留下不同的数字 。返回 已排序的链表 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：head = [1,2,3,3,4,4,5]
 * 输出：[1,2,5]
 *
 *
 * 示例 2：
 *
 *
 * 输入：head = [1,1,1,2,3]
 * 输出：[2,3]
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

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package leetcode

import (
	"testing"
)

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy

	for pre.Next != nil && pre.Next.Next != nil {
		if pre.Next.Val == pre.Next.Next.Val {
			x := pre.Next.Val
			for pre.Next != nil && pre.Next.Val == x {
				pre.Next = pre.Next.Next
			}
		} else {
			pre = pre.Next
		}
	}

	return dummy.Next
}

func deleteDuplicates82_1(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy
	last := 999

	for head != nil {
		if head.Next != nil && head.Val == head.Next.Val {
			last = head.Val
			pre.Next = head.Next.Next
			head = head.Next.Next
		} else if last == head.Val {
			pre.Next = head.Next
			head = head.Next
		} else {
			pre = head
			head = head.Next
		}
	}

	return dummy.Next
}

// @lc code=end

func Test_deleteDuplicates(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		want *ListNode
	}{
		{"x.1", NewList([]int{1, 1, 1}), NewList([]int{})},
		{"0", NewList([]int{}), NewList([]int{})},
		{"1", NewList([]int{1, 2, 3}), NewList([]int{1, 2, 3})},
		{"1.1", NewList([]int{1, 1, 2, 3}), NewList([]int{2, 3})},
		{"1.2", NewList([]int{1, 1}), NewList([]int{})},
		{"1.3", NewList([]int{1, 1, 2, 3, 3}), NewList([]int{2})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates(tt.head.Clone()); !got.compare(tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v", got, tt.want)
			}

			if got := deleteDuplicates82_1(tt.head.Clone()); !got.compare(tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
