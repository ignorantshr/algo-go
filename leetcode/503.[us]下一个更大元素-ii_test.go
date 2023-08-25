/*
 * @lc app=leetcode.cn id=503 lang=golang
 *
 * [503] 下一个更大元素 II
 *
 * https://leetcode.cn/problems/next-greater-element-ii/description/
 *
 * algorithms
 * Medium (66.93%)
 * Likes:    851
 * Dislikes: 0
 * Total Accepted:    210.1K
 * Total Submissions: 313.8K
 * Testcase Example:  '[1,2,1]'
 *
 * 给定一个循环数组 nums （ nums[nums.length - 1] 的下一个元素是 nums[0] ），返回 nums 中每个元素的
 * 下一个更大元素 。
 *
 * 数字 x 的 下一个更大的元素 是按数组遍历顺序，这个数字之后的第一个比它更大的数，这意味着你应该循环地搜索它的下一个更大的数。如果不存在，则输出 -1
 * 。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: nums = [1,2,1]
 * 输出: [2,-1,2]
 * 解释: 第一个 1 的下一个更大的数是 2；
 * 数字 2 找不到下一个更大的数；
 * 第二个 1 的下一个最大的数需要循环搜索，结果也是 2。
 *
 *
 * 示例 2:
 *
 *
 * 输入: nums = [1,2,3,4,3]
 * 输出: [2,3,4,-1,4]
 *
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= nums.length <= 10^4
 * -10^9 <= nums[i] <= 10^9
 *
 *
 */
package leetcode

import (
	"reflect"
	"testing"
)

// @lc code=start
func nextGreaterElements(nums []int) []int {
	// 环形循环
	size := len(nums)
	stack := make([]int, 0)
	res := make([]int, size)
	last := -1

	for i := 2*size - 1; i >= 0; i-- {
		for last >= 0 && stack[last] <= nums[i%size] {
			stack = stack[:last]
			last--
		}

		if i < size {
			if last >= 0 {
				res[i] = stack[last]
			} else {
				res[i] = -1
			}
		}
		stack = append(stack, nums[i%size])
		last++
	}

	return res[:size]
}

func nextGreaterElementsDouble(nums []int) []int {
	// 环形展开就完事了
	tmpNums := append(nums, nums[:len(nums)-1]...)
	stack := make([]int, 0)
	res := make([]int, len(tmpNums))
	last := -1

	for i := len(tmpNums) - 1; i >= 0; i-- {
		for last >= 0 && stack[last] <= tmpNums[i] {
			stack = stack[:last]
			last--
		}

		if last >= 0 {
			res[i] = stack[last]
		} else {
			res[i] = -1
		}
		stack = append(stack, tmpNums[i])
		last++
	}

	return res[:len(nums)]
}

// @lc code=end

func Test_nextGreaterElements(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"1", []int{1, 2}, []int{2, -1}},
		{"1", []int{1, 2, 3}, []int{2, 3, -1}},
		{"1", []int{3, 2, 1}, []int{-1, 3, 3}},
		{"1", []int{1, 2, 1}, []int{2, -1, 2}},
		{"1", []int{1, 2, 3, 4, 3}, []int{2, 3, 4, -1, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreaterElements(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextGreaterElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
