/*
 * @lc app=leetcode.cn id=41 lang=golang
 *
 * [41] 缺失的第一个正数
 *
 * https://leetcode.cn/problems/first-missing-positive/description/
 *
 * algorithms
 * Hard (43.50%)
 * Likes:    2000
 * Dislikes: 0
 * Total Accepted:    338.4K
 * Total Submissions: 774.6K
 * Testcase Example:  '[1,2,0]'
 *
 * 给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
 * 请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,2,0]
 * 输出：3
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [3,4,-1,1]
 * 输出：2
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [7,8,9,11,12]
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 5 * 105
 * -231 <= nums[i] <= 231 - 1
 *
 *
 */
package leetcode

import (
	"testing"
)

// @lc code=start
func firstMissingPositive(nums []int) int {
	// 交换元素的位置，将元素放到对应的位置上
	size := len(nums)

	// v 放到 v-1 的位置上
	for i := 0; i < size; i++ {
		for nums[i] >= 1 && nums[i] <= size && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
			// 交换之后继续在同一位置判断
		}
	}

	for i, v := range nums {
		if v-1 != i {
			return i + 1
		}
	}

	return size + 1
}

func firstMissingPositiveHash(nums []int) int {
	// 哈希表的替代品
	size := len(nums)

	// 把 <=0 的数变值
	for i := 0; i < size; i++ {
		if nums[i] <= 0 {
			nums[i] = size + 1
		}
	}

	// 将 x-1 的位置标成负数
	for i := 0; i < size; i++ {
		idx := abs(nums[i]) - 1
		if idx < size && nums[idx] > 0 {
			nums[idx] = -nums[idx]
		}
	}

	// 找到第一个没有被标记的位置
	for i, v := range nums {
		if v > 0 {
			return i + 1
		}
	}

	return size + 1
}

// func abs(a int) int {
// 	if a > 0 {
// 		return a
// 	}
// 	return -a
// }

// @lc code=end

func Test_firstMissingPositive(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"1", []int{}, 1},
		{"2", []int{1, 2, 0}, 3},
		{"3", []int{3, 4, -1, 1}, 2},
		{"4", []int{7, 8, 9, 11, 12}, 1},
		{"5", []int{1, 8, 9, 11, 12}, 2},
		{"6", []int{1, 8, 9, 2, 12}, 3},
		{"7", []int{0, 8, 9, 2, 12}, 1},
		{"8", []int{0, 1}, 2},
		{"9", []int{0, 3}, 1},
		{"10", []int{2, 1}, 3},
		{"11", []int{1, 1}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstMissingPositive(tt.nums); got != tt.want {
				t.Errorf("firstMissingPositive() = %v, want %v", got, tt.want)
			}
			if got := firstMissingPositiveHash(tt.nums); got != tt.want {
				t.Errorf("firstMissingPositive() = %v, want %v", got, tt.want)
			}
		})
	}
}
