/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 *
 * https://leetcode.cn/problems/add-two-numbers/description/
 *
 * algorithms
 * Medium (42.85%)
 * Likes:    10367
 * Dislikes: 0
 * Total Accepted:    2M
 * Total Submissions: 4.6M
 * Testcase Example:  '[2,4,3]\n[5,6,4]'
 *
 * 给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
 *
 * 请你将两个数相加，并以相同形式返回一个表示和的链表。
 *
 * 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：l1 = [2,4,3], l2 = [5,6,4]
 * 输出：[7,0,8]
 * 解释：342 + 465 = 807.
 *
 *
 * 示例 2：
 *
 *
 * 输入：l1 = [0], l2 = [0]
 * 输出：[0]
 *
 *
 * 示例 3：
 *
 *
 * 输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
 * 输出：[8,9,9,9,0,0,0,1]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 每个链表中的节点数在范围 [1, 100] 内
 * 0
 * 题目数据保证列表表示的数字不含前导零
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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	cur := new(ListNode)
	pre := cur
	n1 := l1
	n2 := l2
	add := 0
	for n1 != nil || n2 != nil {
		v := add
		if n1 != nil {
			v += n1.Val
			n1 = n1.Next
		}
		if n2 != nil {
			v += n2.Val
			n2 = n2.Next
		}

		add = v / 10
		cur.Next = &ListNode{Val: v % 10}
		cur = cur.Next
	}

	if add != 0 {
		cur.Next = &ListNode{Val: add}
	}

	return pre.Next
}

// @lc code=end

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"0", args{NewList([]int{0}), NewList([]int{0})}, NewList([]int{0})},
		{"1", args{NewList([]int{3}), NewList([]int{2})}, NewList([]int{5})},
		{"2", args{NewList([]int{2, 4, 3}), NewList([]int{5, 6, 4})}, NewList([]int{7, 0, 8})},
		{"3", args{NewList([]int{9, 9, 9, 9, 9, 9, 9}), NewList([]int{9, 9, 9, 9})}, NewList([]int{8, 9, 9, 9, 0, 0, 0, 1})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
