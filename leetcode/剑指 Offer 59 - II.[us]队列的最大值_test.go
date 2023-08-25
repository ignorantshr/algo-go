/*
- @lc app=leetcode.cn id= lang=golang

https://leetcode.cn/problems/dui-lie-de-zui-da-zhi-lcof/

请定义一个队列并实现函数 max_value 得到队列里的最大值，要求函数max_value、push_back 和 pop_front 的均摊时间复杂度都是O(1)。

若队列为空，pop_front 和 max_value 需要返回 -1

示例 1：

输入:
["MaxQueue","push_back","push_back","max_value","pop_front","max_value"]
[[],[1],[2],[],[],[]]
输出: [null,null,null,2,1,2]
示例 2：

输入:
["MaxQueue","pop_front","max_value"]
[[],[],[]]
输出: [null,-1,-1]

限制：

1 <= push_back,pop_front,max_value的总操作数 <= 10000
1 <= value <= 10^5
*/
package leetcode

// @lc code=start
type MaxQueue struct {
	queue []int
	real  []int
}

func Constructor() MaxQueue {
	return MaxQueue{
		queue: make([]int, 0),
		real:  make([]int, 0),
	}
}

func (this *MaxQueue) Max_value() int {
	if len(this.queue) == 0 {
		return -1
	}
	return this.queue[0]
}

func (this *MaxQueue) Push_back(value int) {
	for len(this.queue) > 0 && this.queue[len(this.queue)-1] < value {
		this.queue = this.queue[1:]
	}
	this.queue = append(this.queue, value)
	this.real = append(this.real, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.queue) == 0 {
		return -1
	}

	front := this.queue[0]
	if front == this.real[0] {
		this.queue = this.queue[1:]
	}
	this.real = this.real[1:]
	return front
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
// @lc code=end
