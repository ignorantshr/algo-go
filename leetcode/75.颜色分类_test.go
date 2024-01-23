/*
 * @lc app=leetcode.cn id=75 lang=golang
 *
 * [75] 颜色分类
 *
 * https://leetcode.cn/problems/sort-colors/description/
 *
 * algorithms
 * Medium (60.68%)
 * Likes:    1730
 * Dislikes: 0
 * Total Accepted:    594.9K
 * Total Submissions: 977.9K
 * Testcase Example:  '[2,0,2,1,1,0]'
 *
 * 给定一个包含红色、白色和蓝色、共 n 个元素的数组 nums ，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
 *
 * 我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。
 *
 *
 *
 *
 * 必须在不使用库内置的 sort 函数的情况下解决这个问题。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [2,0,2,1,1,0]
 * 输出：[0,0,1,1,2,2]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [2,0,1]
 * 输出：[0,1,2]
 *
 *
 *
 *
 * 提示：
 *
 *
 * n == nums.length
 * 1 <= n <= 300
 * nums[i] 为 0、1 或 2
 *
 *
 *
 *
 * 进阶：
 *
 *
 * 你能想出一个仅使用常数空间的一趟扫描算法吗？
 *
 *
 */
package leetcode

import "testing"

// @lc code=start
func sortColors(nums []int) {
	sortColors双指针2(nums)
	// sortColors双指针1(nums)
	// sortColors单指针(nums)
	// sortColorsMine(nums)
}

func sortColors双指针2(nums []int) {
	p0, p2 := 0, len(nums)-1 // nums[:p0)=1 nums[p0:i]=1 nums(p2:]=2
	for i := 0; i <= p2; {
		if nums[i] == 2 {
			nums[i], nums[p2] = nums[p2], nums[i]
			p2--
			// 新的 nums[i] 可能仍然是 2，也可能是 0
		} else if nums[i] == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			p0++
			i++
		} else {
			i++
		}
	}
}

func sortColors双指针1(nums []int) {
	p0, p1 := 0, 0 // nums[:p0) nums[p0:p1)
	for i, v := range nums {
		if v == 0 {
			// 00011120
			//    0  1i
			// 00001121
			//    0  1i
			// 00001112
			//    0  1i
			nums[p0], nums[i] = nums[i], nums[p0]
			if p0 < p1 {
				nums[p1], nums[i] = nums[i], nums[p1]
			}
			p0++
			p1++
		} else if v == 1 {
			nums[p1], nums[i] = nums[i], nums[p1]
			p1++
		}
	}
}

func sortColors单指针(nums []int) {
	p := _swapColor(nums, 0)
	_swapColor(nums[p:], 1)
}

func _swapColor(nums []int, target int) int {
	ptr := 0 // nums[:ptr] 都是 target
	for i, v := range nums {
		if v == target {
			nums[ptr], nums[i] = nums[i], nums[ptr]
			ptr++
		}
	}
	return ptr
}

func sortColorsMine(nums []int) {
	// idxN 代表了第一个颜色的位置
	idx1 := -1
	idx2 := -1
	idx3 := -1
	lastidx1 := -1
	lastidx2 := -1
	lastidx3 := -1
	last := -1

	swap := func(idx1, idx2 int) (int, int) {
		if idx1 != -1 && idx2 != -1 && idx1 > idx2 {
			nums[idx1], nums[idx2] = nums[idx2], nums[idx1]
			idx1, idx2 = idx2, idx1
		}
		return idx1, idx2
	}

	for i, v := range nums {
		if v == 0 {
			if idx1 == -1 || last != 0 {
				lastidx1 = idx1
				idx1 = i
			}
		} else if v == 1 {
			if idx2 == -1 || last != 1 {
				lastidx2 = idx2
				idx2 = i
			}
		} else {
			if idx3 == -1 || last != 2 {
				lastidx3 = idx3
				idx3 = i
			}
		}

		// 00011112222
		// 0  1   2
		// 000111122220
		//    1   2   0
		// 000011122221
		//    0   2   1
		// 000011112222
		//    0   1   2
		// 000011112222
		// 0   1   2
		idx1, idx2 = swap(idx1, idx2)
		idx1, idx3 = swap(idx1, idx3)
		idx2, idx3 = swap(idx2, idx3)
		if lastidx3 != idx3 {
			if lastidx3 != -1 {
				if nums[lastidx3] != 2 {
					lastidx3++
				}
				idx3 = lastidx3
			}
		}
		if lastidx2 != idx2 {
			if lastidx2 != -1 {
				if nums[lastidx2] != 1 {
					lastidx2++
				}
				idx2 = lastidx2
			}
		}
		if lastidx1 != idx1 {
			idx1 = 0
		}
	}
}

// @lc code=end

func Test_sortColors(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"1", []int{1}, []int{1}},
		{"1", []int{2, 2, 2}, []int{2, 2, 2}},
		{"2.1", []int{2, 0}, []int{0, 2}},
		{"2.2", []int{1, 0, 0}, []int{0, 0, 1}},
		{"2.3", []int{1, 2, 1}, []int{1, 1, 2}},
		{"3.1", []int{2, 0, 1}, []int{0, 1, 2}},
		{"3.2", []int{2, 0, 2, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2}},
		{"3.3", []int{0, 2, 2, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2}},
		// {"3", []int{}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oa := make([]int, len(tt.nums))
			copy(oa, tt.nums)
			sortColors(oa)
			if !equalSlice[int](oa, tt.want) {
				t.Errorf("sortColors() = %v, want %v", tt.nums, tt.want)
			}

			copy(oa, tt.nums)
			sortColors双指针1(oa)
			if !equalSlice[int](oa, tt.want) {
				t.Errorf("sortColors双指针1() = %v, want %v", tt.nums, tt.want)
			}

			copy(oa, tt.nums)
			sortColors单指针(oa)
			if !equalSlice[int](oa, tt.want) {
				t.Errorf("sortColors单指针() = %v, want %v", tt.nums, tt.want)
			}

			copy(oa, tt.nums)
			sortColorsMine(oa)
			if !equalSlice[int](oa, tt.want) {
				t.Errorf("sortColorsMine() = %v, want %v", tt.nums, tt.want)
			}
		})
	}
}
