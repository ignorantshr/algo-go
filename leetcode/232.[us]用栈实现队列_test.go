/*
- @lc app=leetcode.cn id=232 lang=golang

使用栈实现队列的下列操作：

push(x) -- 将一个元素放入队列的尾部。
pop() -- 从队列首部移除元素。
peek() -- 返回队列首部的元素。
empty() -- 返回队列是否为空。

示例:

MyQueue queue = new MyQueue();
queue.push(1);
queue.push(2);
queue.peek();  // 返回 1
queue.pop();   // 返回 1
queue.empty(); // 返回 false
说明:

你只能使用标准的栈操作 -- 也就是只有 push to top, peek/pop from top, size, 和 is empty 操作是合法的。
你所使用的语言也许不支持栈。你可以使用 list 或者 deque（双端队列）来模拟一个栈，只要是标准的栈操作即可。
假设所有操作都是有效的 （例如，一个空的队列不会调用 pop 或者 peek 操作）。
*/
package leetcode

import (
	"testing"
)

// @lc code=start

type MyQueue struct {
	stackIn  *stack //输入栈
	stackOut *stack //输出栈
}

func Constructor232() MyQueue {
	return MyQueue{
		stackIn:  &stack{make([]int, 0)},
		stackOut: &stack{make([]int, 0)},
	}
}

func (this *MyQueue) Push(x int) {
	this.stackIn.push(x)
}

func (this *MyQueue) Pop() int {
	if this.stackOut.empty() {
		this.transfer()
	}

	return this.stackOut.pop()
}

func (this *MyQueue) Peek() int {
	if this.stackOut.empty() {
		this.transfer()
	}
	return this.stackOut.peek()
}

func (this *MyQueue) Empty() bool {
	return this.stackIn.empty() && this.stackOut.empty()
}

func (this *MyQueue) transfer() {
	for i := this.stackIn.size() - 1; i >= 0; i-- {
		this.stackOut.push(this.stackIn.pop())
	}
}

// ——————————————————————
type stack struct {
	list []int
}

func (s *stack) push(x int) {
	s.list = append(s.list, x)
}

func (s *stack) pop() int {
	res := s.list[s.size()-1]
	s.list = s.list[:s.size()-1]
	return res
}

func (s *stack) peek() int {
	return s.list[s.size()-1]
}

func (s *stack) size() int {
	return len(s.list)
}

func (s *stack) empty() bool {
	return s.size() == 0
}

// @lc code=end

func Test232(t *testing.T) {
	queue := Constructor232()
	queue.Push(1)
	queue.Push(2)

	if v := queue.Peek(); v != 1 {
		t.Fatalf("Peek: got:%v, want:1\n", v)
	}
	if v := queue.Pop(); v != 1 {
		t.Fatalf("Pop: got:%v, want:1\n", v)
	}
	if v := queue.Empty(); v != false {
		t.Fatalf("Empty: got:%v, want:false\n", v)
	}

	if v := queue.Peek(); v != 2 {
		t.Fatalf("Peek: got:%v, want:2\n", v)
	}
	if v := queue.Pop(); v != 2 {
		t.Fatalf("Pop: got:%v, want:2\n", v)
	}
	if v := queue.Empty(); v != true {
		t.Fatalf("Empty: got:%v, want:true\n", v)
	}
}
