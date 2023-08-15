/*
 * @lc app=leetcode.cn id=870 lang=golang
 *
 * [870] 优势洗牌
 *
 * https://leetcode.cn/problems/advantage-shuffle/description/
 *
 * algorithms
 * Medium (50.42%)
 * Likes:    389
 * Dislikes: 0
 * Total Accepted:    66.3K
 * Total Submissions: 131.3K
 * Testcase Example:  '[2,7,11,15]\n[1,10,4,11]'
 *
 * 给定两个长度相等的数组 nums1 和 nums2，nums1 相对于 nums2 的优势可以用满足 nums1[i] > nums2[i] 的索引 i
 * 的数目来描述。
 *
 * 返回 nums1 的任意排列，使其相对于 nums2 的优势最大化。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums1 = [2,7,11,15], nums2 = [1,10,4,11]
 * 输出：[2,11,7,15]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums1 = [12,24,8,32], nums2 = [13,25,32,11]
 * 输出：[24,32,8,12]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums1.length <= 10^5
 * nums2.length == nums1.length
 * 0 <= nums1[i], nums2[i] <= 10^9
 *
 *
 */
package leetcode

import (
	"container/heap"
	"reflect"
	"sort"
	"testing"
)

// @lc code=start
func advantageCount(nums1 []int, nums2 []int) []int {
	size := len(nums1)
	sort.Ints(nums1)
	idx := make([]int, size)
	for i := range nums2 {
		idx[i] = i
	}
	sort.Slice(idx, func(i, j int) bool { return nums2[idx[i]] < nums2[idx[j]] })

	left := 0
	right := size - 1
	res := make([]int, size)
	for i := 0; i < size; i++ {
		if nums1[i] > nums2[idx[left]] {
			// 保留
			res[idx[left]] = nums1[i]
			left++
		} else {
			// 丢弃
			res[idx[right]] = nums1[i]
			right--
		}
	}

	return res
}

func advantageCountHeap(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)

	// 大顶堆
	h := &heap870{make([][]int, 0, len(nums2))}
	for i, v := range nums2 {
		h.Push([]int{i, v})
	}
	heap.Init(h)

	left := 0
	right := len(nums1) - 1
	res := make([]int, len(nums1))
	for h.Len() > 0 {
		pair := heap.Pop(h).([]int)
		idx, val := pair[0], pair[1]
		if nums1[right] > val {
			res[idx] = nums1[right]
			right--
		} else {
			res[idx] = nums1[left]
			left++
		}
	}

	return res
}

type heap870 struct {
	heap [][]int // []pair{idx, val int}
}

func (h *heap870) Len() int {
	return len(h.heap)
}

func (h *heap870) Less(i int, j int) bool {
	return h.heap[i][1] > h.heap[j][1]
}

// Swap swaps the elements with indexes i and j.
func (h *heap870) Swap(i int, j int) {
	h.heap[i], h.heap[j] = h.heap[j], h.heap[i]
}

func (h *heap870) Push(x any) {
	h.heap = append(h.heap, x.([]int))
}

func (h *heap870) Pop() any {
	x := h.heap[h.Len()-1]
	h.heap = h.heap[:h.Len()-1]
	return x
}

// @lc code=end

func Test_advantageCount(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"0", args{
			[]int{5621, 1743, 5532, 3549, 9581},
			[]int{913, 9787, 4121, 5039, 1481},
		}, []int{1743, 9581, 5532, 5621, 3549}},
		{"0", args{
			[]int{},
			[]int{},
		}, []int{}},
		{"1", args{
			[]int{1},
			[]int{1},
		}, []int{1}},
		{"1", args{
			[]int{1},
			[]int{2},
		}, []int{1}},
		{"1", args{
			[]int{3},
			[]int{2},
		}, []int{3}},
		{"1", args{
			[]int{2, 7, 11, 15},
			[]int{1, 10, 4, 11},
		}, []int{2, 11, 7, 15}},
		{"1", args{
			[]int{12, 24, 8, 32},
			[]int{13, 25, 32, 11},
		}, []int{24, 32, 8, 12}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := advantageCount(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("advantageCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
