/*
 * @lc app=leetcode.cn id=61 lang=golang
 *
 * [61] 旋转链表
 *
 * https://leetcode.cn/problems/rotate-list/description/
 *
 * algorithms
 * Medium (41.34%)
 * Likes:    1021
 * Dislikes: 0
 * Total Accepted:    352.7K
 * Total Submissions: 853.1K
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * 给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：head = [1,2,3,4,5], k = 2
 * 输出：[4,5,1,2,3]
 *
 *
 * 示例 2：
 *
 *
 * 输入：head = [0,1,2], k = 4
 * 输出：[2,0,1]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 链表中节点的数目在范围 [0, 500] 内
 * -100 <= Node.val <= 100
 * 0 <= k <= 2 * 10^9
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
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}

	dummy := &ListNode{Next: head}
	tail := dummy
	var pre *ListNode = dummy
	count := -k
	size := 0
	for tail.Next != nil {
		tail = tail.Next
		if count >= 0 {
			pre = pre.Next
		}
		size++
		count++
	}

	if count < 0 {
		count = size - k%size // size 的整数倍
		if count == size {
			count = 0
		}
		pre = dummy
		for count > 0 {
			pre = pre.Next
			count--
		}
	}

	if pre == dummy { // 防止成环
		return head
	}
	newhead := pre.Next
	tail.Next = head
	pre.Next = nil
	return newhead
}

// @lc code=end

func Test_rotateRight(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"0", args{NewList([]int{}), 0}, NewList([]int{})},
		{"0.1", args{NewList([]int{1}), 0}, NewList([]int{1})},
		{"0.2", args{NewList([]int{1, 2}), 0}, NewList([]int{1, 2})},
		{"1", args{NewList([]int{1}), 1}, NewList([]int{1})},
		{"1.1", args{NewList([]int{1, 2}), 1}, NewList([]int{2, 1})},
		{"1.2", args{NewList([]int{1, 2}), 3}, NewList([]int{2, 1})},
		{"1.3", args{NewList([]int{1, 2, 3, 4, 5}), 2}, NewList([]int{4, 5, 1, 2, 3})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateRight(tt.args.head, tt.args.k); !got.compare(tt.want) {
				t.Errorf("rotateRight() = %v, want %v", got, tt.want)
			}
		})
	}
}
