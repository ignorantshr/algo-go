/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 *
 * https://leetcode.cn/problems/search-in-rotated-sorted-array/description/
 *
 * algorithms
 * Medium (43.95%)
 * Likes:    2833
 * Dislikes: 0
 * Total Accepted:    818.1K
 * Total Submissions: 1.9M
 * Testcase Example:  '[4,5,6,7,0,1,2]\n0'
 *
 * 整数数组 nums 按升序排列，数组中的值 互不相同 。
 *
 * 在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，使数组变为 [nums[k],
 * nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始
 * 计数）。例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。
 *
 * 给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。
 *
 * 你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [4,5,6,7,0,1,2], target = 0
 * 输出：4
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [4,5,6,7,0,1,2], target = 3
 * 输出：-1
 *
 * 示例 3：
 *
 *
 * 输入：nums = [1], target = 0
 * 输出：-1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 5000
 * -10^4 <= nums[i] <= 10^4
 * nums 中的每个值都 独一无二
 * 题目数据保证 nums 在预先未知的某个下标上进行了旋转
 * -10^4 <= target <= 10^4
 *
 *
 */
package leetcode

import "testing"

// @lc code=start
func search(nums []int, target int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		mid := (i + j) >> 1
		if nums[mid] == target {
			return mid
		}

		if nums[i] <= nums[mid] {
			if nums[i] <= target && target < nums[mid] { // 左边有序且 target 位于左边
				j = mid - 1
			} else {
				i = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[j] { // 右边有序且 target 位于右边
				i = mid + 1
			} else {
				j = mid - 1
			}
		}
	}
	return -1
}

func search1(nums []int, target int) int {
	first := nums[0]
	last := nums[len(nums)-1]
	if target < first && target > last {
		return -1
	}

	if target < first {
		// 在右半边
		i, j := 0, len(nums)-1
		for i <= j {
			mid := (i + j) / 2
			if nums[mid] < target {
				i = mid + 1
			} else if nums[mid] > target {
				if first <= last {
					j = mid - 1
				} else {
					if nums[mid] >= first { // mid 在左半边
						// 4 5 m 7 8 2 3
						i = mid + 1
					} else {
						// 4 5 6 7 8 2 m
						j = mid - 1
					}
				}
			} else {
				return mid
			}
		}
		return -1
	} else {
		// 在左半边
		i, j := 0, len(nums)-1
		for i <= j {
			mid := (i + j) / 2
			if nums[mid] > target {
				j = mid - 1
			} else if nums[mid] < target {
				if first <= last {
					i = mid + 1
				} else {
					if nums[mid] >= first { // mid 在左半边
						i = mid + 1
					} else {
						j = mid - 1
					}
				}
			} else {
				return mid
			}
		}
		return -1
	}
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
		want int
	}{
		{"x", args{
			[]int{3, 4, 5, 1, 2},
			4,
		}, 1},
		{"x.1", args{
			[]int{5, 1, 3},
			5,
		}, 0},
		{"1", args{
			[]int{4, 5, 6, 7, 0, 1, 2},
			0,
		}, 4},
		{"3", args{
			[]int{4, 5, 6, 7, 0, 1, 2},
			3,
		}, -1},
		{"3.1", args{
			[]int{4, 5, 6, 7, 0, 1, 2},
			8,
		}, -1},
		{"3.2", args{
			[]int{4, 5, 6, 7, 0, 1, 2},
			-1,
		}, -1},
		{"4", args{
			[]int{4, 5, 6, 7, 0, 1, 2},
			5,
		}, 1},
		{"5", args{
			[]int{4, 5, 6, 7, 0, 1, 2},
			1,
		}, 5},
		{"5", args{
			[]int{4, 5, 6, 7, 0, 1, 2},
			2,
		}, 6},
		{"6", args{
			[]int{3, 1},
			1,
		}, 1},
		{"6.1", args{
			[]int{3, 1},
			3,
		}, 0},
		{"6.2", args{
			[]int{1, 3},
			3,
		}, 1},
		{"6.3", args{
			[]int{1, 3},
			1,
		}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
