/*
- @lc app=leetcode.cn id=15 lang=golang
*/
package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

// @lc code=start
func threeSumReview(nums []int) [][]int {
	// 双指针
	size := len(nums)
	res := make([][]int, 0)
	sort.Ints(nums)

	for a := 0; a < size; a++ {
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}

		left, right := a+1, size-1
		for left < right {
			if nums[right] < -nums[a] {
				break
			}
			if nums[left]+nums[right] == -nums[a] {
				res = append(res, []int{nums[a], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				left++
			} else if nums[left]+nums[right] < -nums[a] {
				left++
			} else {
				right--
			}
		}
	}
	return res
}

// @lc code=end

func Test_threeSumReview(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{"1", []int{-1, 0, 1, 2, -1, -4}, [][]int{{0, 1, -1}, {-1, 2, -1}}},
		{"2", []int{0, 1, 1}, [][]int{}},
		{"2", []int{3, 0, -2, -1, 1, 2}, [][]int{{0, -2, 2}, {0, -1, 1}, {3, -2, -1}}},
		{"2", []int{0, 0, 0}, [][]int{{0, 0, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSumReview(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSumReview() = %v, want %v", got, tt.want)
			}
		})
	}
}
