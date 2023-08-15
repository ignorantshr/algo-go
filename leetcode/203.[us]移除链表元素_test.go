/*
- @lc app=leetcode.cn id=203 lang=golang
https://leetcode.cn/problems/remove-linked-list-elements/

题意：删除链表中等于给定值 val 的所有节点。

示例 1： 输入：head = [1,2,6,3,4,5,6], val = 6 输出：[1,2,3,4,5]

示例 2： 输入：head = [], val = 1 输出：[]

示例 3： 输入：head = [7,7,7,7], val = 7 输出：[]
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
func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}
	cur := dummy
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}

// @lc code=end

func Test_removeElements(t *testing.T) {
	type args struct {
		head *ListNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"0", args{
			nil,
			0,
		}, nil},
		{"1", args{
			NewList([]int{1}),
			1,
		}, nil},
		{"1", args{
			NewList([]int{1}),
			12,
		}, NewList([]int{1})},
		{"2", args{
			NewList([]int{1, 2, 6, 3, 4, 5, 6}),
			6,
		}, NewList([]int{1, 2, 3, 4, 5})},
		{"2", args{
			NewList([]int{1, 1, 1, 1}),
			1,
		}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElements(tt.args.head, tt.args.val); !got.compare(tt.want) {
				t.Errorf("removeElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
