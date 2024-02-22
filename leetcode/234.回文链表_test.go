/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
 *
 * https://leetcode.cn/problems/palindrome-linked-list/description/
 *
 * algorithms
 * Easy (53.29%)
 * Likes:    1740
 * Dislikes: 0
 * Total Accepted:    611.7K
 * Total Submissions: 1.1M
 * Testcase Example:  '[1,2,2,1]'
 *
 * 给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：head = [1,2,2,1]
 * 输出：true
 *
 *
 * 示例 2：
 *
 *
 * 输入：head = [1,2]
 * 输出：false
 *
 *
 *
 *
 * 提示：
 *
 *
 * 链表中节点数目在范围[1, 10^5] 内
 * 0 <= Node.val <= 9
 *
 *
 *
 *
 * 进阶：你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
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
func isPalindrome234(head *ListNode) bool {
	return isPalindromePointer(head)
}

func isPalindromePointer(head *ListNode) bool {
	// 利用双指针找到链表的中点，然后翻转链表进行比较
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	// 奇数，slow 再向前一步
	if fast != nil {
		slow = slow.Next
	}

	left := head
	right := reverseListIterate234(slow)
	for right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}
	return true
}

func reverseListIterate234(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}

func isPalindromeDfs(head *ListNode) bool {
	// 后序遍历
	left := head
	var dfs func(right *ListNode) bool
	dfs = func(right *ListNode) bool {
		if right == nil {
			return true
		}

		res := dfs(right.Next)
		res = res && left.Val == right.Val
		left = left.Next
		return res
	}
	return dfs(left)
}

func isPalindromeScan(head *ListNode) bool {
	// 遍历
	var vals []int
	for head != nil {
		vals = append(vals, head.Val)
		head = head.Next
	}

	size := len(vals)
	mid := size / 2
	var a, b int
	if size&1 == 1 {
		a, b = mid, mid
	} else {
		a, b = mid-1, mid
	}
	for a >= 0 {
		if vals[a] != vals[b] {
			return false
		}
		a--
		b++
	}
	return true
}

// @lc code=end

func Test_isPalindrome234(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		want bool
	}{
		{"1", NewList([]int{1}), true},
		{"2", NewList([]int{1, 2, 2, 1}), true},
		{"2", NewList([]int{1, 2, 3, 2, 1}), true},
		{"3", NewList([]int{1, 2, 2}), false},
		{"4", NewList([]int{1, 2}), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome234(tt.head); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
