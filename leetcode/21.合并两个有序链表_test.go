/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 *
 * https://leetcode.cn/problems/merge-two-sorted-lists/description/
 *
 * algorithms
 * Easy (66.18%)
 * Likes:    3183
 * Dislikes: 0
 * Total Accepted:    1.4M
 * Total Submissions: 2.2M
 * Testcase Example:  '[1,2,4]\n[1,3,4]'
 *
 * 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：l1 = [1,2,4], l2 = [1,3,4]
 * 输出：[1,1,2,3,4,4]
 *
 *
 * 示例 2：
 *
 *
 * 输入：l1 = [], l2 = []
 * 输出：[]
 *
 *
 * 示例 3：
 *
 *
 * 输入：l1 = [], l2 = [0]
 * 输出：[0]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 两个链表的节点数目范围是 [0, 50]
 * -100
 * l1 和 l2 均按 非递减顺序 排列
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
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	return mergeTwoLists2(list1, list2)
}

// 递归
func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists2(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists2(list1, list2.Next)
		return list2
	}
}

func mergeTwoLists1(list1 *ListNode, list2 *ListNode) *ListNode {
	virtual := &ListNode{}
	cur := virtual
	a, b := list1, list2
	for a != nil || b != nil {
		if a == nil {
			cur.Next = b
			break
		}
		if b == nil {
			cur.Next = a
			break
		}

		if list1.Val < list2.Val {
			cur.Next = &ListNode{Val: list1.Val}
			a = list1.Next
		} else {
			cur.Next = &ListNode{Val: list2.Val}
			b = list2.Next
		}
		cur = cur.Next
	}

	return virtual.Next
}

// @lc code=end

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		list1 *ListNode
		list2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"1", args{
			NewList([]int{}),
			NewList([]int{}),
		}, NewList([]int{})},
		{"2", args{
			NewList([]int{}),
			NewList([]int{0}),
		}, NewList([]int{0})},
		{"3", args{
			NewList([]int{1, 2, 4}),
			NewList([]int{1, 3, 4}),
		}, NewList([]int{1, 1, 2, 3, 4, 4})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.list1, tt.args.list2); !tt.want.compare(got) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
