/*
- @lc app=leetcode.cn id=239 lang=golang

给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。

返回滑动窗口中的最大值。

进阶：

你能在线性时间复杂度内解决此题吗？

提示：

1 <= nums.length <= 10^5
-10^4 <= nums[i] <= 10^4
1 <= k <= nums.length
*/
package leetcode

import (
	"reflect"
	"testing"
)

// @lc code=start
func maxSlidingWindow(nums []int, k int) []int {
	res := make([]int, 0)
	queue := sortedQueue{make([]int, 0)}

	for i, v := range nums {
		queue.push(v)
		if i >= k-1 {
			res = append(res, queue.front())
			if nums[i+1-k] == queue.front() {
				queue.pop()
			}
		}
	}
	return res
}

// 优先队列
// 保持队列内元素逆序排序
type sortedQueue struct {
	list []int
}

func (s *sortedQueue) push(x int) {
	for len(s.list) > 0 {
		if s.list[len(s.list)-1] >= x {
			break
		}
		s.list = s.list[:len(s.list)-1]
	}
	s.list = append(s.list, x)
}

func (s *sortedQueue) pop() int {
	res := s.list[0]
	s.list = s.list[1:]
	return res
}

func (s *sortedQueue) front() int {
	return s.list[0]
}

// @lc code=end

func Test_maxSlidingWindow(t *testing.T) {
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
			[]int{},
			0,
		}, []int{}},
		{"1", args{
			[]int{1},
			1,
		}, []int{1}},
		{"1", args{
			[]int{1, 2, 3},
			2,
		}, []int{2, 3}},
		{"1", args{
			[]int{1, 2, 3},
			3,
		}, []int{3}},
		{"2", args{
			[]int{1, 3, -1, -3, 5, 3, 6, 7},
			3,
		}, []int{3, 3, 5, 5, 6, 7}},
		{"2", args{
			[]int{1, 3, -1, -3, 5, 3, 6, 7},
			5,
		}, []int{5, 5, 6, 7}},
		{"2", args{
			[]int{5, 4, 3, 2, 1},
			3,
		}, []int{5, 4, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSlidingWindow(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxSlidingWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
