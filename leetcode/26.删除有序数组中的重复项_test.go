/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除有序数组中的重复项
 *
 * https://leetcode.cn/problems/remove-duplicates-from-sorted-array/description/
 *
 * algorithms
 * Easy (54.80%)
 * Likes:    3243
 * Dislikes: 0
 * Total Accepted:    1.5M
 * Total Submissions: 2.8M
 * Testcase Example:  '[1,1,2]'
 *
 * 给你一个 升序排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持
 * 一致 。然后返回 nums 中唯一元素的个数。
 *
 * 考虑 nums 的唯一元素的数量为 k ，你需要做以下事情确保你的题解可以被通过：
 *
 *
 * 更改数组 nums ，使 nums 的前 k 个元素包含唯一元素，并按照它们最初在 nums 中出现的顺序排列。nums 的其余元素与 nums
 * 的大小不重要。
 * 返回 k 。
 *
 *
 * 判题标准:
 *
 * 系统会用下面的代码来测试你的题解:
 *
 *
 * int[] nums = [...]; // 输入数组
 * int[] expectedNums = [...]; // 长度正确的期望答案
 *
 * int k = removeDuplicates(nums); // 调用
 *
 * assert k == expectedNums.length;
 * for (int i = 0; i < k; i++) {
 * ⁠   assert nums[i] == expectedNums[i];
 * }
 *
 * 如果所有断言都通过，那么您的题解将被 通过。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,1,2]
 * 输出：2, nums = [1,2,_]
 * 解释：函数应该返回新的长度 2 ，并且原数组 nums 的前两个元素被修改为 1, 2 。不需要考虑数组中超出新长度后面的元素。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [0,0,1,1,1,2,2,3,3,4]
 * 输出：5, nums = [0,1,2,3,4]
 * 解释：函数应该返回新的长度 5 ， 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4
 * 。不需要考虑数组中超出新长度后面的元素。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 3 * 10^4
 * -10^4 <= nums[i] <= 10^4
 * nums 已按 升序 排列
 *
 *
 */
package leetcode

import (
	"reflect"
	"testing"
)

// @lc code=start
func removeDuplicates(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}
	slow, fast := 0, 1
	for fast < size {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	return slow + 1
}

func removeDuplicates1(nums []int) int {
	// 遍历，前置
	size := len(nums)
	for i := 0; i < size; i++ {
		j := i + 1
		for ; j < size; j++ {
			if nums[j] != nums[i] {
				break
			}
		}
		if j == size {
			return i + 1
		}

		k := j - i - 1
		if k > 0 {
			for m := i + 1; j < size; m++ {
				nums[m] = nums[j]
				j++
			}
			size -= k
		}
	}
	return size
}

// @lc code=end

func Test_removeDuplicates(t *testing.T) {
	type result struct {
		nums []int
		size int
	}
	tests := []struct {
		name   string
		nums   []int
		result result
	}{
		{"1", []int{}, result{[]int{}, 0}},
		{"1", []int{1}, result{[]int{1}, 1}},
		{"2", []int{1, 1}, result{[]int{1}, 1}},
		{"2", []int{1, 2}, result{[]int{1, 2}, 2}},
		{"2", []int{1, 1, 2}, result{[]int{1, 2}, 2}},
		{"2", []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, result{[]int{0, 1, 2, 3, 4}, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.nums); got != tt.result.size || !reflect.DeepEqual(tt.nums[:tt.result.size], tt.result.nums) {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.result)
			}
		})
	}
}
