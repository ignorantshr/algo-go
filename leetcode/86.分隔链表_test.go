/*
 * @lc app=leetcode.cn id=86 lang=golang
 *
 * [86] 分隔链表
 *
 * https://leetcode.cn/problems/partition-list/description/
 *
 * algorithms
 * Medium (64.18%)
 * Likes:    733
 * Dislikes: 0
 * Total Accepted:    224.3K
 * Total Submissions: 349.4K
 * Testcase Example:  '[1,4,3,2,5,2]\n3'
 *
 * 给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。
 *
 * 你应当 保留 两个分区中每个节点的初始相对位置。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：head = [1,4,3,2,5,2], x = 3
 * 输出：[1,2,2,4,3,5]
 *
 *
 * 示例 2：
 *
 *
 * 输入：head = [2,1], x = 2
 * 输出：[1,2]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 链表中节点的数目在范围 [0, 200] 内
 * -100
 * -200
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
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	small, large := &ListNode{}, &ListNode{}
	dummy1, dummy2 := small, large
	for head != nil {
		if head.Val < x {
			small.Next = head
			small = small.Next
		} else {
			large.Next = head
			large = large.Next
		}
		// 断开原链的节点
		node := head.Next
		head.Next = nil
		head = node
	}
	small.Next = dummy2.Next
	return dummy1.Next
}

func partition1(head *ListNode, x int) *ListNode {
	virtual := &ListNode{Next: head}
	slow := virtual
	var fast *ListNode
	for {
		// fmt.Println(head)
		for ; slow.Next != nil && slow.Next.Val < x; slow = slow.Next {
		}
		if fast == nil {
			fast = slow
		}
		for ; fast.Next != nil && fast.Next.Val >= x; fast = fast.Next {
		}

		if fast.Next == nil {
			return virtual.Next
		}

		// [..., slow,(nodes>=x, fast), node<x, ...]
		node := fast.Next
		fast.Next = node.Next
		node.Next = slow.Next
		slow.Next = node
	}
}

// @lc code=end

func Test_partition(t *testing.T) {
	type args struct {
		head *ListNode
		x    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"1", args{NewList([]int{}), 1}, NewList([]int{})},
		{"1", args{NewList([]int{1}), 1}, NewList([]int{1})},
		{"2", args{NewList([]int{1, 2}), 2}, NewList([]int{1, 2})},
		{"2", args{NewList([]int{1, 2}), 0}, NewList([]int{1, 2})},
		{"3", args{NewList([]int{2, 1}), 2}, NewList([]int{1, 2})},
		{"3", args{NewList([]int{2, 1}), 2}, NewList([]int{1, 2})},
		{"4", args{NewList([]int{1, 4, 3, 2, 5, 2}), 3}, NewList([]int{1, 2, 2, 4, 3, 5})},
		{"4", args{NewList([]int{1, 4, 3, 2, 5, 2}), 6}, NewList([]int{1, 4, 3, 2, 5, 2})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.head, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}
