/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] 下一个排列
 *
 * https://leetcode.cn/problems/next-permutation/description/
 *
 * algorithms
 * Medium (38.62%)
 * Likes:    2380
 * Dislikes: 0
 * Total Accepted:    473.1K
 * Total Submissions: 1.2M
 * Testcase Example:  '[1,2,3]'
 *
 * 整数数组的一个 排列  就是将其所有成员以序列或线性顺序排列。
 *
 *
 * 例如，arr = [1,2,3] ，以下这些都可以视作 arr 的排列：[1,2,3]、[1,3,2]、[3,1,2]、[2,3,1] 。
 *
 *
 * 整数数组的 下一个排列 是指其整数的下一个字典序更大的排列。更正式地，如果数组的所有排列根据其字典顺序从小到大排列在一个容器中，那么数组的 下一个排列
 * 就是在这个有序容器中排在它后面的那个排列。如果不存在下一个更大的排列，那么这个数组必须重排为字典序最小的排列（即，其元素按升序排列）。
 *
 *
 * 例如，arr = [1,2,3] 的下一个排列是 [1,3,2] 。
 * 类似地，arr = [2,3,1] 的下一个排列是 [3,1,2] 。
 * 而 arr = [3,2,1] 的下一个排列是 [1,2,3] ，因为 [3,2,1] 不存在一个字典序更大的排列。
 *
 *
 * 给你一个整数数组 nums ，找出 nums 的下一个排列。
 *
 * 必须 原地 修改，只允许使用额外常数空间。
 * [1,2,3] [1,3,2] [2,1,3] [2,3,1] [3,1,2] [3,2,1]
 * [1,2,3,4] [1,2,4,3] [1,3,2,4] [1,3,4,2] [1,4,2,3] [1,4,3,2]
 * [1,2,3,4,5] [1,2,3,5,4] [1,2,4,3,5] [1,2,4,5,3] [1,2,5,3,4] [1,2,5,4,3] [1,3,2,4,5] [1,3,2,5,4] [1,3,4,2,5]
 * [34521][35124]
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,2,3]
 * 输出：[1,3,2]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [3,2,1]
 * 输出：[1,2,3]
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [1,1,5]
 * 输出：[1,5,1]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 100
 * 0 <= nums[i] <= 100
 *
 *
 */
package leetcode

import (
	"slices"
	"testing"
)

// @lc code=start
func nextPermutation(nums []int) {
	// 从逆序向后找到第一个逆序的位置 i，从 nums[i:] 中找一个大于 nums[pre] 的数 和 nums[pre] 交换 ，然后重新排序 nums[i:]
	n := len(nums)

	i := n - 1
	for ; i > 0 && nums[i] <= nums[i-1]; i-- {
	}

	reverse := func(start, end int) {
		for j, k := start, end; j < k; j, k = j+1, k-1 {
			nums[k], nums[j] = nums[j], nums[k]
		}
	}

	if i == 0 {
		reverse(0, n-1)
		return
	}

	pre := i - 1
	l := i
	r := n - 1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] > nums[pre] { // 注意是逆序
			l = mid + 1
			// } else if nums[mid] <= nums[pre] {
			// 	r = mid - 1
		} else {
			r = mid - 1
		}
	}

	nums[r], nums[pre] = nums[pre], nums[r]
	reverse(i, n-1)
}

// @lc code=end

func Test_nextPermutation(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"1", []int{1, 2, 3}, []int{1, 3, 2}},
		{"2", []int{1, 3, 2}, []int{2, 1, 3}},
		{"22", []int{2, 3, 1}, []int{3, 1, 2}},
		{"3", []int{3, 2, 1}, []int{1, 2, 3}},
		{"3", []int{1, 5, 1}, []int{5, 1, 1}},
		{"33", []int{5, 1, 1}, []int{1, 1, 5}},
		{"4", []int{1, 2, 5, 4, 3}, []int{1, 3, 2, 4, 5}},
		{"5", []int{1, 3, 2, 4, 5}, []int{1, 3, 2, 5, 4}},
		{"6", []int{1, 5, 4, 3, 2}, []int{2, 1, 3, 4, 5}},
		{"6", []int{2, 5, 4, 3, 1}, []int{3, 1, 2, 4, 5}},
		{"7", []int{5, 4, 7, 5, 3, 2}, []int{5, 5, 2, 3, 4, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nextPermutation(tt.nums)
			if !slices.Equal[[]int](tt.nums, tt.want) {
				t.Errorf("threeSumClosest() = %v, want %v", tt.nums, tt.want)
			}
		})
	}
}
