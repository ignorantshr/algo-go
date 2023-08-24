/*
- @lc app=leetcode.cn id=225 lang=golang

使用队列实现栈的下列操作：

push(x) -- 元素 x 入栈
pop() -- 移除栈顶元素
top() -- 获取栈顶元素
empty() -- 返回栈是否为空
注意:

你只能使用队列的基本操作-- 也就是 push to back, peek/pop from front, size, 和 is empty 这些操作是合法的。
你所使用的语言也许不支持队列。 你可以使用 list 或者 deque（双端队列）来模拟一个队列 , 只要是标准的队列操作即可。
你可以假设所有操作都是有效的（例如, 对一个空的栈不会调用 pop 或者 top 操作）。
*/
package leetcode

import "testing"

// @lc code=start
type MyStack struct {
	q *queue
}

func Constructor225() MyStack {
	return MyStack{
		q: &queue{make([]int, 0)},
	}
}

func (this *MyStack) Push(x int) {
	this.q.push(x)
}

func (this *MyStack) Pop() int {
	this.move()
	return this.q.pop()
}

func (this *MyStack) Peek() int {
	this.move()
	res := this.q.peek()
	this.Push(this.q.pop())
	return res
}

func (this *MyStack) Empty() bool {
	return this.q.empty()
}

func (this *MyStack) move() {
	// 队列尾部元素放到头部
	for i := this.q.size() - 1; i > 0; i-- {
		this.Push(this.q.pop())
	}
}

// ——————————————————————
type queue struct {
	list []int
}

func (q *queue) push(x int) {
	q.list = append(q.list, x)
}

func (q *queue) pop() int {
	res := q.list[0]
	q.list = q.list[1:]
	return res
}

func (q *queue) peek() int {
	return q.list[0]
}

func (q *queue) size() int {
	return len(q.list)
}

func (q *queue) empty() bool {
	return q.size() == 0
}

// @lc code=end

func Test225(t *testing.T) {
	stack := Constructor225()
	stack.Push(1)
	stack.Push(2)

	if v := stack.Peek(); v != 2 {
		t.Fatalf("Peek: got:%v, want:2\n", v)
	}
	if v := stack.Pop(); v != 2 {
		t.Fatalf("Pop: got:%v, want:2\n", v)
	}
	if v := stack.Empty(); v != false {
		t.Fatalf("Empty: got:%v, want:false\n", v)
	}

	if v := stack.Peek(); v != 1 {
		t.Fatalf("Peek: got:%v, want:1\n", v)
	}
	if v := stack.Pop(); v != 1 {
		t.Fatalf("Pop: got:%v, want:1\n", v)
	}
	if v := stack.Empty(); v != true {
		t.Fatalf("Empty: got:%v, want:true\n", v)
	}
}
