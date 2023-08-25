package labuladong

import (
	"reflect"
	"testing"
)

// 寻找数组中下一个比它大的数，没有则返回 -1
func nextGreaterElementReview(nums []int) []int {
	// 栈，往里放，遇到栈中小于等于它的元素就弹出这个元素，这样就保持了单调性
	// 因为前面的依赖后面的，所以我们从后往前把元素放入栈里面
	stack := make([]int, 0)
	res := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		for len(stack) > 0 && stack[len(stack)-1] <= nums[i] {
			stack = stack[:len(stack)-1] // 弹出小元素
		}
		if len(stack) > 0 {
			res[i] = stack[len(stack)-1]
		} else {
			res[i] = -1
		}
		stack = append(stack, nums[i]) // 入栈
	}
	return res
}

func Test_nextGreaterElementReview(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"1", []int{2, 1, 2, 4, 3}, []int{4, 2, 4, -1, -1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreaterElementReview(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextGreaterElementReview() = %v, want %v", got, tt.want)
			}
		})
	}
}
