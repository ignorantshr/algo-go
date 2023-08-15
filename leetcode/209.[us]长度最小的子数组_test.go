/*
- @lc app=leetcode.cn id=209 lang=golang

给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的 连续 子数组，并返回其长度。如果不存在符合条件的子数组，返回 0。

示例：

输入：s = 7, nums = [2,3,1,2,4,3]
输出：2
解释：子数组 [4,3] 是该条件下的长度最小的子数组。
提示：

1 <= target <= 10^9
1 <= nums.length <= 10^5
1 <= nums[i] <= 10^5
*/
package leetcode

import (
	"math"
	"testing"
)

// @lc code=start
func minSubArrayLen(target int, nums []int) int {
	// 滑动窗口
	sum := 0
	left, right := 0, 0
	size := len(nums)
	minLen := math.MaxInt

	for right < size {
		sum += nums[right]
		right++

		for sum >= target {
			if minLen > right-left {
				minLen = right - left
			}
			sum -= nums[left]
			left++
		}
	}

	if minLen == math.MaxInt {
		return 0
	}
	return minLen
}

// @lc code=end

func Test_minSubArrayLen(t *testing.T) {
	type args struct {
		target int
		nums   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0", args{
			0,
			[]int{},
		}, 0},
		{"1", args{
			1,
			[]int{1},
		}, 1},
		{"1", args{
			1,
			[]int{1, 2},
		}, 1},
		{"2", args{
			7,
			[]int{2, 3, 1, 2, 4, 3},
		}, 2},
		{"2", args{
			8,
			[]int{2, 3, 1, 2, 4, 3},
		}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSubArrayLen(tt.args.target, tt.args.nums); got != tt.want {
				t.Errorf("minSubArrayLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
