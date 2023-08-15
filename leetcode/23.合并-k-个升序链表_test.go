/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并 K 个升序链表
 *
 * https://leetcode.cn/problems/merge-k-sorted-lists/description/
 *
 * algorithms
 * Hard (57.82%)
 * Likes:    2500
 * Dislikes: 0
 * Total Accepted:    666.2K
 * Total Submissions: 1.2M
 * Testcase Example:  '[[1,4,5],[1,3,4],[2,6]]'
 *
 * 给你一个链表数组，每个链表都已经按升序排列。
 *
 * 请你将所有链表合并到一个升序链表中，返回合并后的链表。
 *
 *
 *
 * 示例 1：
 *
 * 输入：lists = [[1,4,5],[1,3,4],[2,6]]
 * 输出：[1,1,2,3,4,4,5,6]
 * 解释：链表数组如下：
 * [
 * ⁠ 1->4->5,
 * ⁠ 1->3->4,
 * ⁠ 2->6
 * ]
 * 将它们合并到一个有序链表中得到。
 * 1->1->2->3->4->4->5->6
 *
 *
 * 示例 2：
 *
 * 输入：lists = []
 * 输出：[]
 *
 *
 * 示例 3：
 *
 * 输入：lists = [[]]
 * 输出：[]
 *
 *
 *
 *
 * 提示：
 *
 *
 * k == lists.length
 * 0 <= k <= 10^4
 * 0 <= lists[i].length <= 500
 * -10^4 <= lists[i][j] <= 10^4
 * lists[i] 按 升序 排列
 * lists[i].length 的总和不超过 10^4
 *
 *
 */
package leetcode

import (
	"container/heap"
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
func mergeKLists(lists []*ListNode) *ListNode {
	// 优先队列
	queue := make(queue23, 0)

	for _, l := range lists {
		if l != nil { // 注意⚠️
			heap.Push(&queue, l)
		}
	}

	dummy := &ListNode{}
	tail := dummy
	for queue.Len() > 0 {
		node := heap.Pop(&queue).(*ListNode)
		tail.Next = node
		tail = tail.Next
		if node.Next != nil { // 推入下一个节点
			heap.Push(&queue, node.Next)
		}
	}
	return dummy.Next
}

type queue23 []*ListNode

// implement priority queue by heap
// Len is the number of elements in the collection.
func (q *queue23) Len() int {
	return len(*q)
}

// Less reports whether the element with index i
// must sort before the element with index j.
func (q *queue23) Less(i int, j int) bool {
	return (*q)[i].Val < (*q)[j].Val
}

// Swap swaps the elements with indexes i and j.
func (q *queue23) Swap(i int, j int) {
	(*q)[i], (*q)[j] = (*q)[j], (*q)[i]
}

func (q *queue23) Push(x any) {
	*q = append(*q, x.(*ListNode))
}

func (q *queue23) Pop() any {
	if q.Len() == 0 {
		return nil
	}

	x := (*q)[q.Len()-1]
	*q = (*q)[:q.Len()-1]
	return x
}

func mergeKLists1(lists []*ListNode) *ListNode {
	size := len(lists)
	if size == 0 {
		return nil
	}
	if size == 1 {
		return lists[0]
	}

	mid := size / 2
	left := mergeKLists(lists[:mid])
	right := mergeKLists(lists[mid:size])

	dummy := &ListNode{}
	node := dummy
	for left != nil && right != nil {
		if left.Val < right.Val {
			node.Next = &ListNode{Val: left.Val}
			left = left.Next
		} else {
			node.Next = &ListNode{Val: right.Val}
			right = right.Next
		}
		node = node.Next
	}

	if left == nil {
		node.Next = right
	} else {
		node.Next = left
	}

	return dummy.Next
}

// @lc code=end

func Test_mergeKLists(t *testing.T) {
	tests := []struct {
		name  string
		lists []*ListNode
		want  *ListNode
	}{
		{"1", []*ListNode{}, NewList([]int{})},
		{"1", []*ListNode{NewList([]int{})}, NewList([]int{})},
		{"2", []*ListNode{NewList([]int{1})}, NewList([]int{1})},
		{"2", []*ListNode{NewList([]int{1}), NewList([]int{1})}, NewList([]int{1, 1})},
		{"2", []*ListNode{NewList([]int{1, 4, 5}), NewList([]int{1, 3, 4}), NewList([]int{2, 6})}, NewList([]int{1, 1, 2, 3, 4, 4, 5, 6})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeKLists(tt.lists); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeKLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
