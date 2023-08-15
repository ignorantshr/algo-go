package heap

import (
	"log"
	"strconv"
	"strings"
)

type bigHeap []int

func (h *bigHeap) String() string {
	res := make([]string, 0)
	for _, v := range *h {
		res = append(res, strconv.Itoa(v))
	}
	return strings.Join(res, ",")
}

func NewBigHeap(data []int) *bigHeap {
	h := bigHeap(data)
	return &h
}

func (h *bigHeap) Init() {
	for i := h.Len() - 1; i >= 0; i-- {
		h.down(i)
	}
}

func (h *bigHeap) IsHeaping() bool {
	for i := 0; i < h.Len(); i++ {
		if h.leftIndex(i) < h.Len() && (*h)[i] < (*h)[h.leftIndex(i)] || h.rightIndex(i) < h.Len() && (*h)[i] < (*h)[h.rightIndex(i)] {
			log.Printf("i:%v, %v; left:%v, right:%v",
				i, (*h)[i], h.leftIndex(i), h.rightIndex(i))
			return false
		}
	}
	return true
}

func (h *bigHeap) Len() int {
	return len(*h)
}

func (h *bigHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *bigHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *bigHeap) Push(x int) {
	*h = append(*h, x)
	h.up(h.Len() - 1)
}

func (h *bigHeap) Pop() int {
	if h.Len() == 0 {
		return 0
	}

	e := (*h)[0]
	h.Swap(0, h.Len()-1)
	*h = (*h)[:h.Len()-1]
	if h.Len() > 1 {
		h.down(0)
	}
	return e
}

func (h *bigHeap) Top() int {
	if h.Len() == 0 {
		return 0
	}

	return (*h)[0]
}

func (h *bigHeap) up(i int) {
	for i >= 0 {
		p := h.parentIndex(i)
		if (*h)[p] >= (*h)[i] {
			return
		}
		h.Swap(i, p)
		i = p
	}
}

func (h *bigHeap) down(i int) {
	for {
		l, r, max := h.leftIndex(i), h.rightIndex(i), i
		if l < h.Len() && (*h)[max] < (*h)[l] {
			max = l
		}
		if r < h.Len() && (*h)[max] < (*h)[r] {
			max = r
		}

		if max == i {
			return
		}

		h.Swap(i, max)
		i = max
	}
}

func (h *bigHeap) parentIndex(index int) int {
	return (index - 1) / 2
}

func (h *bigHeap) leftIndex(index int) int {
	return 2*index + 1
}

func (h *bigHeap) rightIndex(index int) int {
	return 2*index + 2
}
