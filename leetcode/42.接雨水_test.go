/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 *
 * https://leetcode.cn/problems/trapping-rain-water/description/
 *
 * algorithms
 * Hard (63.15%)
 * Likes:    4911
 * Dislikes: 0
 * Total Accepted:    823.6K
 * Total Submissions: 1.3M
 * Testcase Example:  '[0,1,0,2,1,0,1,3,2,1,2,1]'
 *
 * 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
 * 输出：6
 * 解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
 *
 *
 * 示例 2：
 *
 *
 * 输入：height = [4,2,0,3,2,5]
 * 输出：9
 *
 *
 *
 *
 * 提示：
 *
 *
 * n == height.length
 * 1 <= n <= 2 * 10^4
 * 0 <= height[i] <= 10^5
 *
 *
 */
package leetcode

import "testing"

// @lc code=start
func trap(height []int) int {
	return trapPoint(height)
}

func trapStack(height []int) int {
	// 单调栈
	sum := 0
	stack := make([]int, 0)
	top := -1

	for i, v := range height {
		if top == -1 {
			stack = append(stack, i)
			top++
			continue
		}

		for top >= 0 && v > height[stack[top]] {
			if top >= 1 {
				edge := stack[top-1]
				sum += (min(height[edge], v) - height[stack[top]]) * (i - edge - 1) // heigh * width
			}
			stack = stack[:top]
			top--
		}
		stack = append(stack, i)
		top++
	}

	return sum
}

// 空间复杂度 和 时间复杂度 都最好
func trapPoint(height []int) int {
	// 左右两端双指针，每次移动小的指针，找到比它大的柱子，计算，更新指针位置
	sum := 0
	left, right := 0, len(height)-1
	for left < right {
		if height[left] <= height[right] {
			k := left
			tmp := 0 // 洼地蓄水量
			for left++; left < right && height[left] <= height[k]; left++ {
				tmp += height[k] - height[left]
			}
			sum += tmp
		} else {
			k := right
			tmp := 0
			for right--; left < right && height[right] <= height[k]; right-- {
				tmp += height[k] - height[right]
			}
			sum += tmp
		}
	}

	return sum
}

// 超时
func trapMe(height []int) int {
	// 每次消减一层
	sum := 0
	size := len(height)

	for {
		column := 0
		j := -1 // j: 最靠近 i 的左边柱子
		for i := 0; i < size; {
			sum += i - j - 1
			j = i

			for ; j < size && height[j] <= 0; j++ {
			}
			for ; j+1 < size && height[j+1] > 0; j++ {
			}
			for k := i; k < j; k++ {
				height[k]--
				if height[k] > 0 {
					column++
				}
			}

			for i = j + 2; i < size && height[i] <= 0; i++ {
			}
			for k := j; k < i && k < size; k++ {
				height[k]--
				if height[k] > 0 {
					column++
				}
			}
		}
		if column < 2 {
			break
		}
	}

	return sum
}

// @lc code=end

func Test_trap(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{"0", []int{}, 0},
		{"00", []int{4, 2, 3}, 1},
		{"2", []int{4, 2, 0, 3, 2, 5}, 9},
		{"22", []int{0, 4, 2, 0, 3, 2, 5, 0}, 9},
		{"1", []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
		{"3", []int{1, 2, 3, 4, 5}, 0},
		{"33", []int{3, 2, 1, 0}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trap(tt.height); got != tt.want {
				t.Errorf("trap() = %v, want %v", got, tt.want)
			}
			if got := trapStack(tt.height); got != tt.want {
				t.Errorf("trapStack() = %v, want %v", got, tt.want)
			}
			if got := trapPoint(tt.height); got != tt.want {
				t.Errorf("trapPoint() = %v, want %v", got, tt.want)
			}
			if got := trapMe(tt.height); got != tt.want {
				t.Errorf("trapMe() = %v, want %v", got, tt.want)
			}
		})
	}
}
