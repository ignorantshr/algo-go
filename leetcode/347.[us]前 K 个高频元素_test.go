/*
- @lc app=leetcode.cn id=347 lang=golang

给定一个非空的整数数组，返回其中出现频率前 k 高的元素。

示例 1:

输入: nums = [1,1,1,2,2,3], k = 2
输出: [1,2]
示例 2:

输入: nums = [1], k = 1
输出: [1]
提示：

你可以假设给定的 k 总是合理的，且 1 ≤ k ≤ 数组中不相同的元素的个数。
你的算法的时间复杂度必须优于 O(nlogn), n 是数组的大小。

题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的。
你可以按任意顺序返回答案。
*/
package leetcode

import (
	"container/heap"
	"sort"
	"testing"
)

// @lc code=start

func topKFrequent(nums []int, k int) []int {
	return topKFrequentHeap(nums, k)
	// return topKFrequentQuickSort(nums, k)
}

func topKFrequentHeap(nums []int, k int) []int {
	frequency := make(map[int]int, 0)
	for _, v := range nums {
		frequency[v]++
	}
	// 堆，1. 构建小顶堆，弹出小的，保留k个数, 这种比较省空间
	// 2. 构建大顶堆，全部塞入堆中，然后弹出 k 个数
	h := &heap347{}
	for v, fre := range frequency {
		heap.Push(h, [2]int{v, fre})
		if h.Len() == k+1 {
			heap.Pop(h)
		}
	}

	res := make([]int, h.Len())
	for i, v := range h.data {
		res[i] = v[0]
	}
	return res
}

func topKFrequentQuickSort(nums []int, k int) []int {
	frequency := make(map[int]int, 0)
	for _, v := range nums {
		frequency[v]++
	}

	ans := make([]int, 0, len(frequency))
	for v := range frequency {
		ans = append(ans, v)
	}

	sort.Slice(ans, func(i, j int) bool {
		return frequency[ans[i]] > frequency[ans[j]]
	})
	return ans[:k]
}

type heap347 struct {
	data [][2]int // {v, fre}
}

// Len is the number of elements in the collection.
func (h *heap347) Len() int {
	return len(h.data)
}

func (h *heap347) Less(i int, j int) bool {
	return h.data[i][1] < h.data[j][1]
}

// Swap swaps the elements with indexes i and j.
func (h *heap347) Swap(i int, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *heap347) Push(x any) {
	h.data = append(h.data, x.([2]int))
}

func (h *heap347) Pop() any {
	res := h.data[h.Len()-1]
	h.data = h.data[:h.Len()-1]
	return res
}

// @lc code=end

func Test_topKFrequent(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{
			[]int{1},
			1,
		}, []int{1}},
		{"1", args{
			[]int{1, 1, 1, 2, 2, 3},
			2,
		}, []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := topKFrequent(tt.args.nums, tt.args.k); !equalSet(got, tt.want) {
				t.Errorf("topKFrequent() = %v, want %v", got, tt.want)
			}
		})
	}
}
