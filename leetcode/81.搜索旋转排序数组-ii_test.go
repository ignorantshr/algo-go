/*
 * @lc app=leetcode.cn id=81 lang=golang
 *
 * [81] 搜索旋转排序数组 II
 *
 * https://leetcode.cn/problems/search-in-rotated-sorted-array-ii/description/
 *
 * algorithms
 * Medium (40.99%)
 * Likes:    763
 * Dislikes: 0
 * Total Accepted:    213.3K
 * Total Submissions: 519.7K
 * Testcase Example:  '[2,5,6,0,0,1,2]\n0'
 *
 * 已知存在一个按非降序排列的整数数组 nums ，数组中的值不必互不相同。
 *
 * 在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转 ，使数组变为 [nums[k],
 * nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始
 * 计数）。例如， [0,1,2,4,4,4,5,6,6,7] 在下标 5 处经旋转后可能变为 [4,5,6,6,7,0,1,2,4,4] 。
 *
 * 给你 旋转后 的数组 nums 和一个整数 target ，请你编写一个函数来判断给定的目标值是否存在于数组中。如果 nums 中存在这个目标值
 * target ，则返回 true ，否则返回 false 。
 *
 * 你必须尽可能减少整个操作步骤。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [2,5,6,0,0,1,2], target = 0
 * 输出：true
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [2,5,6,0,0,1,2], target = 3
 * 输出：false
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 5000
 * -10^4 <= nums[i] <= 10^4
 * 题目数据保证 nums 在预先未知的某个下标上进行了旋转
 * -10^4 <= target <= 10^4
 *
 *
 *
 *
 * 进阶：
 *
 *
 * 这是 搜索旋转排序数组 的延伸题目，本题中的 nums  可能包含重复元素。
 * 这会影响到程序的时间复杂度吗？会有怎样的影响，为什么？
 *
 *
 *
 *
 */
package leetcode

import (
	"testing"
)

// @lc code=start
func search(nums []int, target int) bool {
	lo, hi := 0, len(nums)-1

	for lo <= hi {
		mid := lo + (hi-lo)>>1
		if nums[mid] == target {
			return true
		}

		if nums[lo] == nums[mid] && nums[mid] == nums[hi] {
			lo++
			hi--
		} else if nums[lo] <= nums[mid] {
			if nums[lo] <= target && target < nums[mid] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[hi] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
	}

	return false
}

// @lc code=end

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"x.2", args{[]int{1, 0, 1, 1, 1}, 0}, true},
		{"x.1", args{[]int{1, 1, 1, 0, 1}, 0}, true},
		{"1.1", args{[]int{4, 5, 6, 6, 7, 0, 1, 2, 4, 4}, 0}, true},
		{"1.2", args{[]int{4, 5, 6, 6, 7, 0, 1, 2, 4, 4}, 1}, true},
		{"1.3", args{[]int{4, 5, 6, 6, 7, 0, 1, 2, 4, 4}, 2}, true},
		{"1.4", args{[]int{4, 5, 6, 6, 7, 0, 1, 2, 4, 4}, 4}, true},
		{"1.5", args{[]int{4, 5, 6, 6, 7, 0, 1, 2, 4, 4}, 5}, true},
		{"1.6", args{[]int{4, 5, 6, 6, 7, 0, 1, 2, 4, 4}, 6}, true},
		{"1.7", args{[]int{4, 5, 6, 6, 7, 0, 1, 2, 4, 4}, 7}, true},
		{"2.1", args{[]int{4, 5, 6, 6, 8, 0, 1, 2, 4, 4}, -1}, false},
		{"2.2", args{[]int{4, 5, 6, 6, 8, 0, 1, 2, 4, 4}, 3}, false},
		{"2.3", args{[]int{4, 5, 6, 6, 8, 0, 1, 2, 4, 4}, 7}, false},
		{"2.4", args{[]int{4, 5, 6, 6, 8, 0, 1, 2, 4, 4}, 9}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
