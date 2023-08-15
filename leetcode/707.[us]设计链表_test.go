/*
- @lc app=leetcode.cn id=707 lang=golang
在链表类中实现这些功能：

get(index)：获取链表中第 index 个节点的值。如果索引无效，则返回-1。
addAtHead(val)：在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
addAtTail(val)：将值为 val 的节点追加到链表的最后一个元素。
addAtIndex(index,val)：在链表中的第 index 个节点之前添加值为 val  的节点。如果 index 等于链表的长度，则该节点将附加到链表的末尾。如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。
deleteAtIndex(index)：如果索引 index 有效，则删除链表中的第 index 个节点。

index 从 0 开始
*/
package leetcode

import (
	"fmt"
	"testing"
)

// @lc code=start
type MyLinkedList struct {
	head   *ListNode
	tail   *ListNode
	length int
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{
		nil,
		nil,
		0,
	}
}

/*
  - Get the value of the index-th node in the linked list. If the index is
    invalid, return -1.
*/
func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.length {
		return -1
	}

	cur := this.head
	for index--; index >= 0; index-- {
		cur = cur.Next
	}
	return cur.Val
}

/*
  - Add a node of value val before the first element of the linked list. After
    the insertion, the new node will be the first node of the linked list.
*/
func (this *MyLinkedList) AddAtHead(val int) {
	this.head = &ListNode{Next: this.head, Val: val}
	this.length++
	if this.tail == nil {
		this.tail = this.head
	}
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	tail := &ListNode{Val: val}
	this.tail.Next = tail
	this.tail = tail
	this.length++
	if this.head == nil {
		this.head = this.tail
	}
}

/*
  - Add a node of value val before the index-th node in the linked list. If
    index equals to the length of linked list, the node will be appended to the
    end of linked list. If index is greater than the length, the node will not be
    inserted.
*/
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index >= this.length {
		return
	}

	if index < 0 {
		this.AddAtHead(val)
		return
	}

	if index == this.length {
		this.AddAtTail(val)
		return
	}

	dummy := &ListNode{Next: this.head}
	pre := dummy
	for ; index > 0; index-- {
		pre = pre.Next
	}

	pre.Next = &ListNode{Next: pre.Next, Val: val}
	this.length++
	this.head = dummy.Next
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.length {
		return
	}

	dummy := &ListNode{Next: this.head}
	pre := dummy
	for index > 0 {
		pre = pre.Next
		index--
	}

	if this.tail == pre.Next {
		this.tail = pre
	}
	pre.Next = pre.Next.Next
	this.length--
	this.head = dummy.Next
}

// @lc code=end

func (this *MyLinkedList) walk() {
	for cur := this.head; cur != nil; cur = cur.Next {
		fmt.Printf(" %d ->", cur.Val)
	}
	fmt.Printf(" {%v,%v}[%d]\n", this.head, this.tail, this.length)
}

func TestConstructor707(t *testing.T) {
	l := Constructor()
	l.AddAtHead(2)
	l.walk()
	l.AddAtHead(1)
	l.walk()
	l.AddAtTail(3)
	l.walk()
	l.AddAtIndex(-1, -1)
	l.walk()
	l.AddAtIndex(1, 0)
	l.walk()
	l.AddAtIndex(0, -2)
	l.walk()
	l.DeleteAtIndex(0)
	l.walk()
	l.DeleteAtIndex(4)
	l.walk()
	l.DeleteAtIndex(1)
	l.walk()
}
