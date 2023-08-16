package leetcode

import "testing"

func reverseKGroup25(head *ListNode, k int) *ListNode {
	tail := head
	n := 0
	for n = k; n > 0 && tail != nil; n-- {
		tail = tail.Next
	}

	if n != 0 {
		return head
	}

	nhead := reverse25(head, tail)
	head.Next = reverseKGroup25(tail, k) // 翻转之后 head 变成了末尾节点
	return nhead
}

// [head,tail)
func reverse25(head, tail *ListNode) *ListNode {
	var pre *ListNode
	for head != tail {
		next := head.Next
		head.Next = pre

		pre = head
		head = next
	}
	return pre
}

func Test_reverseKGroup25(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"1", args{
			NewList([]int{1}),
			1,
		}, NewList([]int{1})},
		{"1", args{
			NewList([]int{1, 2}),
			2,
		}, NewList([]int{2, 1})},
		{"1", args{
			NewList([]int{1, 2, 3, 4, 5}),
			2,
		}, NewList([]int{2, 1, 4, 3, 5})},
		{"1", args{
			NewList([]int{1, 2, 3, 4, 5}),
			3,
		}, NewList([]int{3, 2, 1, 4, 5})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseKGroup25(tt.args.head, tt.args.k); !got.compare(tt.want) {
				t.Errorf("reverseKGroup25() = %v, want %v", got, tt.want)
			}
		})
	}
}
