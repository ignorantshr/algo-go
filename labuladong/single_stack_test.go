package labuladong

import (
	"reflect"
	"testing"
)

/* 单调栈 */

// 输入一个数组 nums，请你返回一个等长的结果数组，结果数组中对应索引存储着下一个更大元素，如果没有更大的元素，就存 -1。
func nextGreaterElement(nums []int) []int {
	res := make([]int, len(nums))
	stack := make([]int, 0)
	last := -1 // 栈尾索引

	for i := len(nums) - 1; i >= 0; i-- {
		for last >= 0 && stack[last] <= nums[i] { // 两者 <= 的之间出栈
			stack = stack[:last]
			last--
		}

		if last >= 0 { // 找到了下一个大值
			res[i] = stack[last]
		} else { // 没找到
			res[i] = -1
		}
		stack = append(stack, nums[i])
		last++
	}

	return res
}

func Test_nextGreaterElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"1", []int{2, 1, 2, 4, 3}, []int{4, 2, 4, -1, -1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreaterElement(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextGreaterElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
