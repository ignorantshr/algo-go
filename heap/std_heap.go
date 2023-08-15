package heap

// https://www.hello-algo.com/chapter_heap/heap/

import (
	"strconv"
	"strings"
)

// 实现标准包 heap.Interface 接口
type stdBigHeap struct {
	data []int
}

func (h *stdBigHeap) String() string {
	res := make([]string, 0)
	for _, v := range h.data {
		res = append(res, strconv.Itoa(v))
	}
	return strings.Join(res, ",")
}

func NewStdBigHeap(data []int) *stdBigHeap {
	return &stdBigHeap{data: data}
}

func (h *stdBigHeap) Len() int {
	return len(h.data)
}

func (h *stdBigHeap) Less(i, j int) bool {
	return h.data[i] > h.data[j]
}

func (h *stdBigHeap) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *stdBigHeap) Push(x interface{}) {
	h.data = append(h.data, x.(int))
}

func (h *stdBigHeap) Pop() interface{} {
	// 注意，待出堆元素会放置在尾部
	e := h.data[h.Len()-1]
	h.data = h.data[:h.Len()-1]
	return e
}

func (h *stdBigHeap) Top() interface{} {
	if h.Len() == 0 {
		return nil
	}

	return h.data[0]
}
