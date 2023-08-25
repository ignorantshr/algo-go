package labuladong

// MonotonicQueue 单调队列的通用实现，可以高效维护最大值和最小值
type MonotonicQueue struct {
	maxQueue []int
	minQueue []int
	real     []int
}

// push 向队尾加入元素
func (q *MonotonicQueue) Push(elem int) {
	q.pushMaxQueue(elem)
	q.pushMinQueue(elem)
	q.real = append(q.real, elem)
}

func (q *MonotonicQueue) pushMaxQueue(elem int) {
	for len(q.maxQueue) > 0 && q.maxQueue[len(q.maxQueue)-1] < elem {
		q.maxQueue = q.maxQueue[:len(q.maxQueue)-1]
	}
	q.maxQueue = append(q.maxQueue, elem)
}

func (q *MonotonicQueue) pushMinQueue(elem int) {
	for len(q.minQueue) > 0 && q.minQueue[len(q.minQueue)-1] > elem {
		q.minQueue = q.minQueue[:len(q.minQueue)-1]
	}
	q.minQueue = append(q.minQueue, elem)
}

// pop 从队头弹出元素，符合先进先出的顺序
func (q *MonotonicQueue) Pop() int {
	front := q.real[0]
	q.real = q.real[1:]

	if front == q.maxQueue[len(q.maxQueue)-1] {
		q.maxQueue = q.maxQueue[1:]
	}

	if front == q.minQueue[len(q.minQueue)-1] {
		q.minQueue = q.minQueue[1:]
	}
	return front
}

// size 返回队列中的元素个数
func (q *MonotonicQueue) Size() int {
	return len(q.real)
}

// max 单调队列特有 API，O(1) 时间计算队列中元素的最大值
func (q *MonotonicQueue) Max() int {
	return q.maxQueue[0]
}

// min 单调队列特有 API，O(1) 时间计算队列中元素的最小值
func (q *MonotonicQueue) Min() int {
	return q.minQueue[0]
}
