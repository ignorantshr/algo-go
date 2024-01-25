/*
 * @lc app=leetcode.cn id=80 lang=golang
 *
 * [80] 删除有序数组中的重复项 II
 *
 * https://leetcode.cn/problems/remove-duplicates-from-sorted-array-ii/description/
 *
 * algorithms
 * Medium (61.43%)
 * Likes:    975
 * Dislikes: 0
 * Total Accepted:    329.4K
 * Total Submissions: 535.6K
 * Testcase Example:  '[1,1,1,2,2,3]'
 *
 * 给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使得出现次数超过两次的元素只出现两次 ，返回删除后数组的新长度。
 *
 * 不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
 *
 *
 *
 * 说明：
 *
 * 为什么返回数值是整数，但输出的答案是数组呢？
 *
 * 请注意，输入数组是以「引用」方式传递的，这意味着在函数里修改输入数组对于调用者是可见的。
 *
 * 你可以想象内部操作如下:
 *
 *
 * // nums 是以“引用”方式传递的。也就是说，不对实参做任何拷贝
 * int len = removeDuplicates(nums);
 *
 * // 在函数里修改输入数组对于调用者是可见的。
 * // 根据你的函数返回的长度, 它会打印出数组中 该长度范围内 的所有元素。
 * for (int i = 0; i < len; i++) {
 * print(nums[i]);
 * }
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,1,1,2,2,3]
 * 输出：5, nums = [1,1,2,2,3]
 * 解释：函数应返回新长度 length = 5, 并且原数组的前五个元素被修改为 1, 1, 2, 2, 3。 不需要考虑数组中超出新长度后面的元素。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [0,0,1,1,1,1,2,3,3]
 * 输出：7, nums = [0,0,1,1,2,3,3]
 * 解释：函数应返回新长度 length = 7, 并且原数组的前七个元素被修改为 0, 0, 1, 1, 2, 3, 3。不需要考虑数组中超出新长度后面的元素。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 3 * 10^4
 * -10^4 <= nums[i] <= 10^4
 * nums 已按升序排列
 *
 *
 */
package leetcode

import "testing"

// @lc code=start
func removeDuplicates(nums []int) int {
	// 双指针
	i := 2
	j := 2
	size := len(nums)
	if size <= 2 {
		return size
	}

	for j < size {
		if nums[i-2] != nums[j] {
			nums[i] = nums[j]
			i++
		}
		j++
	}

	return i
}

// 双指针标记法
func removeDuplicates80_2(nums []int) int {
	i := 0
	j := 0
	size := len(nums)
	special := 99999

	for j < size {
		for j < size && nums[j] == nums[i] {
			j++
		}

		if j-i > 2 {
			for i = i + 2; i < j; i++ {
				nums[i] = special
			}
		} else {
			i = j
			j++
		}
	}

	for i = 0; i < size && nums[i] != special; i++ {
	}
	if i == size {
		return i
	}

	for i < size && nums[i] == special {
		for j = i + 1; j < size && nums[j] == special; j++ {
		}
		if j == size {
			return i
		}
		nums[i] = nums[j]
		nums[j] = special
		i++
	}

	return i
}

// 整体平移
func removeDuplicates80_1(nums []int) int {
	i := 0
	j := 0
	size := len(nums)

	for j < size {
		for j < size && nums[j] == nums[i] {
			j++
		}
		if j == size {
			return min(i+2, j)
		}

		if j-i > 2 {
			for m := j; m < size; m++ {
				nums[i+2+m-j] = nums[m]
			}
			size -= j - i - 2
			i = i + 2
			j = i + 1
		} else {
			i = j
			j++
		}
	}

	return min(i+2, j)
}

// @lc code=end

func Test_removeDuplicates(t *testing.T) {
	tests := []struct {
		name  string
		nums  []int
		after []int
		want  int
	}{
		{"1", []int{1}, []int{1}, 1},
		{"1.1", []int{1, 2, 3}, []int{1, 2, 3}, 3},
		{"2.1", []int{1, 1}, []int{1, 1}, 2},
		{"2.2", []int{1, 1, 1, 3}, []int{1, 1, 3}, 3},
		{"2.3", []int{1, 1, 1, 3, 3, 3, 3, 3}, []int{1, 1, 3, 3}, 4},
		{"2.4", []int{1, 1, 1, 3, 3}, []int{1, 1, 3, 3}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.nums); got != tt.want || !equalSlice[int](tt.nums[:tt.want], tt.after) {
				t.Errorf("removeDuplicates() = %v, after %v, want %v", got, tt.after, tt.want)
			}
		})
	}
}
