package heap

import (
	"container/heap"
	"fmt"
	"math/rand"
	"testing"
)

func TestStdBigHeap(t *testing.T) {
	h := NewStdBigHeap(nil)
	heap.Push(h, 3)
	heap.Push(h, 2)
	heap.Push(h, 7)
	heap.Push(h, 10)
	heap.Push(h, 2343)

	fmt.Println(h)

	fmt.Println(h.Top())
	fmt.Println(heap.Pop(h))
	fmt.Println(h.Top())
	fmt.Println(heap.Pop(h))
	fmt.Println(h.Top())
	fmt.Println(heap.Pop(h))
	fmt.Println(h.Top())
	fmt.Println(heap.Pop(h))
	fmt.Println(h.Top())
	fmt.Println(heap.Pop(h))

	initheap := NewStdBigHeap([]int{1, 4, 5, 2, 6, 7})
	heap.Init(initheap)
	fmt.Println(initheap)
}

func TestBigHeap(t *testing.T) {
	h := NewBigHeap(nil)
	h.Push(3)
	h.Push(2)
	h.Push(7)
	h.Push(10)
	h.Push(2343)

	fmt.Println(h)

	fmt.Println(h.Top())
	fmt.Println(h.Pop())
	fmt.Println(h.Top())
	fmt.Println(h.Pop())
	fmt.Println(h.Top())
	fmt.Println(h.Pop())
	fmt.Println(h.Top())
	fmt.Println(h.Pop())
	fmt.Println(h.Top())
	fmt.Println(h.Pop())
}

func TestBigHeap_Init(t *testing.T) {
	n := 20
	var data []int
	for i := 0; i < n; i++ {
		data = append(data, rand.Intn(10000))
	}
	initheap := NewBigHeap(data)
	initheap.Init()
	fmt.Println(initheap)
	fmt.Println(initheap.IsHeaping())
}
