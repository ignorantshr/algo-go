/*
 * @lc app=leetcode.cn id=84 lang=golang
 *
 * [84] 柱状图中最大的矩形
 *
 * https://leetcode.cn/problems/largest-rectangle-in-histogram/description/
 *
 * algorithms
 * Hard (45.22%)
 * Likes:    2651
 * Dislikes: 0
 * Total Accepted:    386.7K
 * Total Submissions: 851.5K
 * Testcase Example:  '[2,1,5,6,2,3]'
 *
 * 给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
 *
 * 求在该柱状图中，能够勾勒出来的矩形的最大面积。
 *
 *
 *
 * 示例 1:
 *
 *
 *
 *
 * 输入：heights = [2,1,5,6,2,3]
 * 输出：10
 * 解释：最大的矩形为图中红色区域，面积为 10
 *
 *
 * 示例 2：
 *
 *
 *
 *
 * 输入： heights = [2,4]
 * 输出： 4
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 0
 *
 *
 */
package leetcode

import "testing"

// @lc code=start
func largestRectangleArea(heights []int) int {
	return largestRectangleAreaStackImprove(heights)
	return largestRectangleAreaStack(heights)
	return largestRectangleAreaEnumHeight(heights)
	return largestRectangleAreaEnumWidth(heights)
}

func largestRectangleAreaStackImprove(heights []int) int {
	// 单调栈 + 常数优化
	size := len(heights)
	lpillars := make([]int, size)
	rpillars := make([]int, size)
	stack := make([]int, 0)

	for i := 0; i < size; i++ {
		rpillars[i] = size
	}

	for i, v := range heights {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= v {
			rpillars[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			lpillars[i] = -1
		} else {
			lpillars[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	ans := 0
	for i, v := range heights {
		ans = max(ans, (rpillars[i]-lpillars[i]-1)*v)
	}

	return ans
}

func largestRectangleAreaStack(heights []int) int {
	// 单调栈
	size := len(heights)
	lpillars := make([]int, size)
	rpillars := make([]int, size)
	stack := make([]int, 0)

	for i, v := range heights {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= v {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			lpillars[i] = -1
		} else {
			lpillars[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	clear(stack)
	for i := size - 1; i >= 0; i-- {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			rpillars[i] = size
		} else {
			rpillars[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	ans := 0
	for i, v := range heights {
		ans = max(ans, (rpillars[i]-lpillars[i]-1)*v)
	}

	return ans
}

// 暴力解法
func largestRectangleAreaEnumHeight(heights []int) int {
	// 枚举高
	size := len(heights)
	ans := 0

	for mid := 0; mid < size; mid++ {
		h := heights[mid]
		// [l,...,r]
		l := mid
		r := mid

		// 找到左边小于 h 的柱子
		for ; l-1 >= 0 && heights[l-1] >= h; l-- {
		}

		for ; r+1 < size && heights[r+1] >= h; r++ {
		}

		ans = max(ans, (r-l+1)*h)
	}

	return ans
}

// 暴力解法
func largestRectangleAreaEnumWidth(heights []int) int {
	// 取决于一个域内最低的柱子
	// 枚举宽
	size := len(heights)
	res := 0

	for i := 0; i < size; i++ {
		lowest := heights[i]
		for j := i; j < size; j++ {
			if heights[j] < lowest {
				lowest = heights[j]
			}
			if res < lowest*(j-i+1) {
				res = lowest * (j - i + 1)
			}
		}
	}
	return res
}

// @lc code=end

func Test_largestRectangleArea(t *testing.T) {
	tests := []struct {
		name    string
		heights []int
		want    int
	}{
		{"x", []int{2, 2, 2}, 6},
		{"0", []int{}, 0},
		{"1", []int{2, 1, 5, 6, 2, 3}, 10},
		{"1", []int{2, 4}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestRectangleArea(tt.heights); got != tt.want {
				t.Errorf("largestRectangleArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
