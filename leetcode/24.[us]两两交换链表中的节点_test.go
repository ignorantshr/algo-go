/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 *
 * https://leetcode.cn/problems/swap-nodes-in-pairs/description/
 *
 * algorithms
 * Medium (71.59%)
 * Likes:    1977
 * Dislikes: 0
 * Total Accepted:    687.9K
 * Total Submissions: 960.3K
 * Testcase Example:  '[1,2,3,4]'
 *
 * 给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：head = [1,2,3,4]
 * 输出：[2,1,4,3]
 *
 *
 * 示例 2：
 *
 *
 * 输入：head = []
 * 输出：[]
 *
 *
 * 示例 3：
 *
 *
 * 输入：head = [1]
 * 输出：[1]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 链表中节点的数目在范围 [0, 100] 内
 * 0 <= Node.val <= 100
 *
 *
 */
package leetcode

import "testing"

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	return swapPairsDfs(head)
}

func swapPairsDfs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	next := head.Next
	head.Next = swapPairsDfs(next.Next)
	next.Next = head
	return next
}

func swapPairsIterate(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy
	for head != nil && head.Next != nil {
		next := head.Next
		pre.Next = next
		head.Next = next.Next
		next.Next = head

		pre = head
		head = head.Next
	}
	return dummy.Next
}

// @lc code=end

func Test_swapPairs(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		want *ListNode
	}{
		{"0", nil, nil},
		{"1", NewList([]int{1, 2}), NewList([]int{2, 1})},
		{"1", NewList([]int{1, 2, 3}), NewList([]int{2, 1, 3})},
		{"1", NewList([]int{1, 2, 3, 4}), NewList([]int{2, 1, 4, 3})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swapPairs(tt.head); !got.compare(tt.want) {
				t.Errorf("swapPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
