package leetcode

import (
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewList(vals []int) *ListNode {
	pre := &ListNode{}
	cur := pre
	for _, n := range vals {
		cur.Next = &ListNode{
			Val: n,
		}
		cur = cur.Next
	}
	return pre.Next
}

func (l *ListNode) String() string {
	var res strings.Builder
	res.WriteByte('[')
	for a := l; a != nil; a = a.Next {
		res.WriteString(strconv.Itoa(a.Val))
		res.WriteString(", ")
	}
	res.WriteByte(']')
	return res.String()
}

func (l *ListNode) compare(another *ListNode) bool {
	a, b := l, another
	for a != nil && b != nil {
		if a.Val != b.Val {
			return false
		}
		a = a.Next
		b = b.Next
	}

	if a != nil || b != nil {
		return false
	}
	return true
}

func (l *ListNode) Clone() *ListNode {
	dummy := &ListNode{}
	pre := dummy
	cur := l
	for cur != nil {
		pre.Next = &ListNode{
			Val: cur.Val,
		}
		pre = pre.Next
		cur = cur.Next
	}
	return dummy.Next
}

// pos: 环形节点的位置，从 0 开始
func (l *ListNode) buildCycle(pos int) *ListNode {
	if pos < 0 {
		return nil
	}

	var cycleNode *ListNode
	tail := l
	for tail != nil && tail.Next != nil {
		if pos == 0 {
			cycleNode = tail
		}
		pos--
		tail = tail.Next
	}
	if cycleNode != nil {
		tail.Next = cycleNode
	}

	return l
}

// pos: 节点索引，从 0 开始
func (l *ListNode) node(pos int) *ListNode {
	if pos < 0 {
		return nil
	}

	tail := l
	for tail != nil && pos > 0 {
		pos--
		tail = tail.Next
	}

	return tail
}
